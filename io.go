package plugins

import (
	"io"
	"os"
	"slices"

	"github.com/markbates/iox"
)

type IO = iox.IO

// IOable can be implemented to return an IO.
type IOable interface {
	Plugin
	iox.IOable
}

// IOSetable can be implemented to receive an IO.
type IOSetable interface {
	Plugin
	iox.IOSetable
}

type Stderrer interface {
	Plugin
	iox.Stderrer
}

type Stdiner interface {
	Plugin
	iox.Stdiner
}

type Stdioer interface {
	Plugin
	iox.Stdioer
}

type Stdouter interface {
	Plugin
	iox.Stdouter
}

// Stdout returns a io.MultiWriter containing all
// plugins that implement Outer. If none are found,
// then os.Stdout is returned
func Stdout(plugs ...Plugin) io.Writer {
	if len(plugs) == 0 {
		return os.Stdout
	}

	outs := ByType[Stdouter](plugs)
	writers := make([]io.Writer, 0, len(outs))
	for _, p := range outs {
		writers = append(writers, p.Stdout())
	}

	// Use slices.Compact to remove any nil writers
	writers = slices.DeleteFunc(writers, func(w io.Writer) bool {
		return w == nil
	})

	if len(writers) == 0 {
		return os.Stdout
	}

	return io.MultiWriter(writers...)
}

// Stderr returns a io.MultiWriter containing all
// plugins that implement Outer. If none are found,
// then os.Stderr is returned
func Stderr(plugs ...Plugin) io.Writer {
	if len(plugs) == 0 {
		return os.Stderr
	}

	outs := ByType[Stderrer](plugs)
	writers := make([]io.Writer, 0, len(outs))
	for _, p := range outs {
		writers = append(writers, p.Stderr())
	}

	// Use slices.Compact to remove any nil writers
	writers = slices.DeleteFunc(writers, func(w io.Writer) bool {
		return w == nil
	})

	if len(writers) == 0 {
		return os.Stderr
	}

	return io.MultiWriter(writers...)
}
