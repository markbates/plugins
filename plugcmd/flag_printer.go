package plugcmd

import (
	"fmt"
	"io"

	"github.com/markbates/plugins"
)

type FlagPrinter interface {
	plugins.Plugin
	PrintFlags(w io.Writer) error
}

var _ FlagPrinter = FlagPrinterFn(nil)

// FlagPrinterFn is a function that can be used to implement the FlagPrinter interface
type FlagPrinterFn func(w io.Writer) error

func (fn FlagPrinterFn) PrintFlags(w io.Writer) error {
	return fn(w)
}

func (fn FlagPrinterFn) PluginName() string {
	return fmt.Sprintf("%T", fn)
}
