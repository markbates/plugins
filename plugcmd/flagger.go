package plugcmd

import "flag"

type Flagger interface {
	Commander
	Flags() (*flag.FlagSet, error)
}

type FlaggerFn func() (*flag.FlagSet, error)

func (fn FlaggerFn) Flags() (*flag.FlagSet, error) {
	return fn()
}
