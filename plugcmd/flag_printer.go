package plugcmd

import "io"

type FlagPrinter interface {
	PrintFlags(w io.Writer) error
}

type FlagPrinterFn func(w io.Writer) error

func (fn FlagPrinterFn) PrintFlags(w io.Writer) error {
	return fn(w)
}
