package plugcmd

import (
	"fmt"

	"github.com/markbates/plugins"
)

// Namer is a command that provides a different name for the
// command than the name of the Plugin
type Namer interface {
	CmdName() string
}

var _ Namer = NamerFn(nil)
var _ plugins.Plugin = NamerFn(nil)

// NamerFn is a function that can be used to implement the Namer interface
type NamerFn func() string

func (fn NamerFn) CmdName() string {
	return fn()
}

func (fn NamerFn) PluginName() string {
	return fmt.Sprintf("%T", fn)
}
