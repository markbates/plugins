package plugcmd

import (
	"fmt"
	"io"

	"github.com/markbates/plugins"
)

// UsagePrinter is called by `Print` and can be implemented
// to print a large block of usage information after the
// `Describer` interface is called. This is useful for printing
// flag information, links, and other messages to users.
type UsagePrinter interface {
	plugins.Plugin
	PrintUsage(w io.Writer) error
}

var _ UsagePrinter = UsagePrinterFn(nil)

// UsagePrinterFn is a function that can be used to implement the UsagePrinter interface
type UsagePrinterFn func(w io.Writer) error

func (fn UsagePrinterFn) PrintUsage(w io.Writer) error {
	return fn(w)
}

func (fn UsagePrinterFn) PluginName() string {
	return fmt.Sprintf("%T", fn)
}
