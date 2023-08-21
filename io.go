package plugins

import (
	"io"
	"os"

	"github.com/markbates/iox"
)

type IO = iox.IO

// IOable can be implemented to return an IO.
type IOable interface {
	Plugin
	Stdio() IO
}

// IOSetable can be implemented to receive an IO.
type IOSetable interface {
	Plugin
	SetStdio(oi IO)
}

type Stdiner interface {
	Plugin
	Stdin() io.Reader
}

type Stdouter interface {
	Plugin
	Stdout() io.Writer
}

type Stderrer interface {
	Plugin
	Stderr() io.Writer
}

type Stdioer interface {
	Plugin
	Stderrer
	Stdiner
	Stdouter
}

// Stdout returns a io.MultiWriter containing all
// plugins that implement Outer. If none are found,
// then os.Stdout is returned
func Stdout(plugs ...Plugin) io.Writer {
	var ins []io.Writer

	for _, p := range plugs {
		if x, ok := p.(Stdouter); ok {
			ins = append(ins, x.Stdout())
		}
	}

	if len(ins) == 0 {
		return os.Stdout
	}
	return io.MultiWriter(ins...)
}

// Stderr returns a io.MultiWriter containing all
// plugins that implement Outer. If none are found,
// then os.Stderr is returned
func Stderr(plugs ...Plugin) io.Writer {
	var ins []io.Writer

	for _, p := range plugs {
		if x, ok := p.(Stderrer); ok {
			ins = append(ins, x.Stderr())
		}
	}

	if len(ins) == 0 {
		return os.Stderr
	}

	return io.MultiWriter(ins...)
}
