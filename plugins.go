package plugins

import (
	"io/fs"
	"sort"

	"github.com/markbates/iox"
)

var _ iox.IOSetable = Plugins{}
var _ Needer = Plugins{}
var _ Scoper = Plugins{}
var _ sort.Interface = Plugins{}

// Plugins is a slice of type `Plugin` that provides
// additional useful functionality.
type Plugins []Plugin

// Len is the number of elements in the collection.
func (p Plugins) Len() int {
	return len(p)
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (p Plugins) Less(i int, j int) bool {
	return p[i].PluginName() < p[j].PluginName()
}

// Swap swaps the elements with indexes i and j.
func (p Plugins) Swap(i int, j int) {
	p[i], p[j] = p[j], p[i]
}

// ScopedPlugins implements Scoper, return itself.
func (plugs Plugins) ScopedPlugins() Plugins {
	return plugs
}

// SetStdio for those plugins that implement
// IOSetable.
func (plugs Plugins) SetStdio(io IO) {
	for _, p := range plugs {
		if ioable, ok := p.(iox.IOSetable); ok {
			ioable.SetStdio(io)
		}
	}
}

// WithPlugins will call any Needer plugins with the
// Feeder function.
func (plugs Plugins) WithPlugins(fn Feeder) {
	for _, p := range plugs {
		if needer, ok := p.(Needer); ok {
			needer.WithPlugins(fn)
		}
	}
}

// SetFS for those plugins that implement FSSetable.
func (plugs Plugins) SetFileSystem(fs fs.FS) {
	for _, p := range plugs {
		if fss, ok := p.(FSSetable); ok {
			fss.SetFileSystem(fs)
		}
	}
}

// Find plugins using the given Finder.
func (plugs Plugins) Find(fn FinderFn) Plugins {
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

// Sort will sort a slice of plugins, by wrapping them
// in Plugins and using it's sort.Interface implementation.
func Sort(plugs Plugins) {
	sort.Sort(Plugins(plugs))
}
