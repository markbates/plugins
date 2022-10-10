package plugins

import (
	"path"
)

type Finder interface {
	Find(plugs Plugins) Plugins
}

type FinderFn func(plugs Plugins) Plugins

func (f FinderFn) Find(plugs Plugins) Plugins {
	return f(plugs)
}

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

func ByType[T any](plugs Plugins) []T {
	var res []T

	for _, p := range plugs {
		if plug, ok := p.(T); ok {
			res = append(res, plug)
		}
	}

	return res

}
