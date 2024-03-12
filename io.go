package plugins

import (
	"io"
	"os"

	"github.com/markbates/iox"
)

type IO = iox.IO

// IOable can be implemented to return an IO.
type IOable = iox.IOable

// IOSetable can be implemented to receive an IO.
type IOSetable = iox.IOSetable

type Stderrer = iox.Stderrer
type Stdiner = iox.Stdiner
type Stdioer = iox.Stdioer
type Stdouter = iox.Stdouter

// Stdout returns a io.MultiWriter containing all
// plugins that implement Outer. If none are found,
// then os.Stdout is returned
func Stdout(plugs ...Plugin) io.Writer {
	if len(plugs) == 0 {
		return os.Stdout
	}

	var ins []io.Writer

	outs := ByType[Stdouter](plugs)
	for _, p := range outs {
		ins = append(ins, p.Stdout())
	}

	return io.MultiWriter(ins...)
}

// Stderr returns a io.MultiWriter containing all
// plugins that implement Outer. If none are found,
// then os.Stderr is returned
func Stderr(plugs ...Plugin) io.Writer {
	if len(plugs) == 0 {
		return os.Stderr
	}

	var ins []io.Writer

	outs := ByType[Stderrer](plugs)
	for _, p := range outs {
		ins = append(ins, p.Stderr())
	}

	return io.MultiWriter(ins...)
}
