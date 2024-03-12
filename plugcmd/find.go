package plugcmd

import (
	"path"
	"strings"

	"github.com/markbates/plugins"
)

// FindFromArgs uses the first arg that does not begin with `-`
// as the name argument for Find
func FindFromArgs(args []string, plugs plugins.Plugins) Commander {
	for _, a := range args {
		if strings.HasPrefix(a, "-") {
			continue
		}
		return Find(a, plugs)
	}

	return nil
}

// Find wraps the other cmd finders into a mega finder for cmds
func Find(name string, plugs plugins.Plugins) Commander {
	if len(plugs) == 0 {
		return nil
	}

	cmds := plugins.ByType[Commander](plugs)

	for _, c := range cmds {
		if n, ok := c.(Namer); ok {
			if n.CmdName() == name {
				return c
			}
		}

		if n, ok := c.(Aliaser); ok {
			for _, a := range n.CmdAliases() {
				if a == name {
					return c
				}
			}
		}

		if name == path.Base(c.PluginName()) {
			return c
		}
	}

	return nil
}
