package plugcmd

import (
	"fmt"

	"github.com/markbates/plugins"
)

// SubCommander can be implemented to provide a list of plugins.Plugin
// that can be used as sub-commands of the current Plugin
type SubCommander interface {
	plugins.Plugin
	SubCommands() []Commander
}

var _ SubCommander = SubCommanderFn(nil)

// SubCommanderFn is a function that can be used to implement the SubCommander interface
type SubCommanderFn func() []Commander

func (fn SubCommanderFn) SubCommands() []Commander {
	return fn()
}

func (fn SubCommanderFn) PluginName() string {
	return fmt.Sprintf("%T", fn)
}
