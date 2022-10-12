package plugins

import (
	"path"
)

// Finder can be implemented to find plugins
type Finder interface {
	Find(plugs Plugins) Plugins
}

// FinderFn is a function that can be used
// to find plugins. It implements the Finder
// interface.
type FinderFn func(plugs Plugins) Plugins

// Find plugins using the underlying function.
func (f FinderFn) Find(plugs Plugins) Plugins {
	return f(plugs)
}

// Background finder that will search for a
// plugin based on the plugin's name.
func Background(name string) Finder {
	fn := func(plugs Plugins) Plugins {
		var res Plugins

		for _, p := range plugs {
			if name == p.PluginName() {
				res = append(res, p)
				continue
			}

			if name == path.Base(p.PluginName()) {
				res = append(res, p)
				continue
			}
		}

		return res
	}
	return FinderFn(fn)
}

// ByType finder can be used to find plugins
// by their type.
func ByType[T any](plugs Plugins) []T {
	var res []T

	for _, p := range plugs {
		if plug, ok := p.(T); ok {
			res = append(res, plug)
		}
	}

	return res

}
