package plugcmd

import (
	"flag"
	"fmt"

	"github.com/markbates/plugins"
)

type Flagger interface {
	Flags() (*flag.FlagSet, error)
}

var _ Flagger = FlaggerFn(nil)
var _ plugins.Plugin = FlaggerFn(nil)

// FlaggerFn is a function that can be used to implement the Flagger interface
type FlaggerFn func() (*flag.FlagSet, error)

func (fn FlaggerFn) Flags() (*flag.FlagSet, error) {
	return fn()
}

func (fn FlaggerFn) PluginName() string {
	return fmt.Sprintf("%T", fn)
}
