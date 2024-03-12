package plugins

import (
	"fmt"
	"io/fs"
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
			return err
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
			return err
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
			return err
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
