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
	for _, p := range plugs {
		c, ok := p.(Commander)
		if !ok {
			continue
		}
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
