package plugcmd

import (
	"fmt"

	"github.com/markbates/plugins"
)

// Aliaser is a command that provides aliases for itself
type Aliaser interface {
	CmdAliases() []string
}

var _ Aliaser = AliaserFn(nil)
var _ plugins.Plugin = AliaserFn(nil)

// AliaserFn is a function that can be used to implement the Aliaser interface
type AliaserFn func() []string

func (fn AliaserFn) CmdAliases() []string {
	return fn()
}

func (fn AliaserFn) PluginName() string {
	return fmt.Sprintf("%T", fn)
}
