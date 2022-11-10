package plugcmd

import "github.com/markbates/plugins"

// SubCommander can be implemented to provide a list of plugins.Plugin
// that can be used as sub-commands of the current Plugin
type SubCommander interface {
	Commander
	SubCommands() plugins.Plugins
}

type SubCommanderFn func() plugins.Plugins

func (fn SubCommanderFn) SubCommands() plugins.Plugins {
	return fn()
}
