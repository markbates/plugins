// Package plugins provides a flexible and extensible plugin system for Go applications.
//
// The core concept revolves around the Plugin interface, which provides
// a name for identification. Additional interfaces like Needer, Feeder,
// AvailabilityChecker, IOSetable, and FSSetable extend functionality.
//
// # Basic Usage
//
// Create plugins by implementing the Plugin interface and any additional
// specialized interfaces:
//
//	type MyPlugin struct {
//		name string
//	}
//
//	func (p MyPlugin) PluginName() string {
//		return p.name
//	}
//
// Manage a collection of plugins using the Plugins slice type:
//
//	plugs := plugins.Plugins{
//		MyPlugin{name: "plugin1"},
//		MyPlugin{name: "plugin2"},
//	}
//
//	// Find plugins by type
//	myPlugins := plugins.ByType[MyPlugin](plugs)
//
//	// Check availability
//	available := plugs.Available("/some/path")
//
//	// Configure I/O and filesystem
//	err := plugs.SetStdio(someIO)
//	err = plugs.SetFileSystem(someFS)
//
// # Plugin Types
//
// The package provides several specialized plugin interfaces:
//
//   - Plugin: Basic interface that all plugins must implement
//   - Scoper: Plugins that can return scoped plugin collections
//   - Feeder/Needer: Plugin communication and dependency injection
//   - AvailabilityChecker: Runtime availability checking
//   - IOSetable/FSSetable: I/O and filesystem configuration
//
// See the plugcmd subpackage for command-line specific plugin interfaces.
package plugins

import (
	"fmt"
	"io/fs"
	"log/slog"
)

var _ FSSetable = Plugins{}
var _ Feeder = Plugins{}
var _ IOSetable = Plugins{}
var _ Scoper = Plugins{}

// Plugins is a slice of type `Plugin` that provides
// additional useful functionality.
type Plugins []Plugin

// Len is the number of elements in the collection.
func (plugs Plugins) Len() int {
	return len(plugs)
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (plugs Plugins) Less(i int, j int) bool {
	return plugs[i].PluginName() < plugs[j].PluginName()
}

// Swap swaps the elements with indexes i and j.
func (plugs Plugins) Swap(i int, j int) {
	plugs[i], plugs[j] = plugs[j], plugs[i]
}

func (plugs Plugins) PluginFeeder() FeederFn {
	return func() Plugins {
		return plugs
	}
}

// ScopedPlugins implements Scoper, return itself.
func (plugs Plugins) ScopedPlugins() Plugins {
	return plugs
}

// SetStdio for those plugins that implement
// IOSetable.
func (plugs Plugins) SetStdio(io IO) error {
	ios := ByType[IOSetable](plugs)

	for _, p := range ios {
		if err := p.SetStdio(io); err != nil {
			slog.Error("failed to set stdio for plugin",
				"plugin", p.PluginName(),
				"error", err)
			return fmt.Errorf("failed to set stdio for plugin %s: %w", p.PluginName(), err)
		}
	}

	return nil
}

// WithPlugins will call any Needer plugins with the
// Feeder function.
func (plugs Plugins) WithPlugins(fn FeederFn) error {
	if fn == nil {
		return fmt.Errorf("no FeederFn provided")
	}

	needers := ByType[Needer](plugs)

	for _, n := range needers {
		if err := n.WithPlugins(fn); err != nil {
			return fmt.Errorf("failed to set plugins for needer %s: %w", n.PluginName(), err)
		}
	}

	return nil
}

// SetFS for those plugins that implement FSSetable.
func (plugs Plugins) SetFileSystem(fs fs.FS) error {
	if fs == nil {
		return fmt.Errorf("no fs.FS provided")
	}

	fsps := ByType[FSSetable](plugs)

	for _, p := range fsps {
		if err := p.SetFileSystem(fs); err != nil {
			slog.Error("failed to set filesystem for plugin",
				"plugin", p.PluginName(),
				"error", err)
			return fmt.Errorf("failed to set filesystem for plugin %s: %w", p.PluginName(), err)
		}
	}

	return nil
}

// Find plugins using the given Finder.
func (plugs Plugins) Find(fn FinderFn) (Plugins, error) {
	return fn(plugs)
}

// Available will return plugins that are available
// to be used at the given root. By default, all
// plugins are available. The AvailabilityChecker
// interface can be implemented to give the plugin
// the ability to check if it is available.
func (plugs Plugins) Available(root string) Plugins {
	var res Plugins

	for _, p := range plugs {
		ac, ok := p.(AvailabilityChecker)
		if !ok {
			res = append(res, p)
			continue
		}

		if ac.PluginAvailable(root) {
			res = append(res, p)
		}
	}

	return res
}

func (plugs Plugins) PluginName() string {
	return fmt.Sprintf("%T", plugs)
}

// Validate checks the plugins collection for common issues.
// It verifies that no duplicate plugin names exist and that
// all plugins have valid names.
func (plugs Plugins) Validate() error {
	if len(plugs) == 0 {
		return fmt.Errorf("no plugins provided")
	}

	names := make(map[string]bool, len(plugs))
	for i, plugin := range plugs {
		if plugin == nil {
			return fmt.Errorf("plugin at index %d is nil", i)
		}

		name := plugin.PluginName()
		if name == "" {
			return fmt.Errorf("plugin at index %d has empty name", i)
		}

		if names[name] {
			return fmt.Errorf("duplicate plugin name: %s", name)
		}
		names[name] = true
	}
	return nil
}

// Names returns a slice of all plugin names in the collection.
func (plugs Plugins) Names() []string {
	names := make([]string, 0, len(plugs))
	for _, p := range plugs {
		names = append(names, p.PluginName())
	}
	return names
}
