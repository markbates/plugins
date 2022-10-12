package plugins

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

// IOable can be implemented to return an IO.
type IOable interface {
	Stdio() IO
}

// IOSetable can be implemented to receive an IO.
type IOSetable interface {
	SetStdio(oi IO)
}

// IO represents the standard input, output, and error stream.
type IO struct {
	In  io.Reader `json:"in,omitempty"`  // standard input
	Out io.Writer `json:"out,omitempty"` // standard output
	Err io.Writer `json:"err,omitempty"` // standard error
}

// Stdout returns IO.In.
// Defaults to os.Stdout.
func (oi IO) Stdout() io.Writer {
	if oi.Out == nil {
		return os.Stdout
	}

	return oi.Out
}

// Stderr returns IO.Err.
// Defaults to os.Stderr.
func (oi IO) Stderr() io.Writer {
	if oi.Err == nil {
		return os.Stderr
	}

	return oi.Err
}

// Stdin returns IO.In.
// Defaults to os.Stdin.
func (oi IO) Stdin() io.Reader {
	if oi.In == nil {
		return os.Stdin
	}

	return oi.In
}

func (oi IO) String() string {
	b, _ := json.MarshalIndent(oi, "", "  ")
	return string(b)
}

func (oi IO) MarshalJSON() ([]byte, error) {
	m := map[string]string{
		"stdin":  fmt.Sprintf("%T", oi.Stdin()),
		"stdout": fmt.Sprintf("%T", oi.Stdout()),
		"stderr": fmt.Sprintf("%T", oi.Stderr()),
	}

	return json.Marshal(m)
}

// Stdout returns a io.MultiWriter containing all
// plugins that implement Outer. If none are found,
// then os.Stdout is returned
func Stdout(plugs ...Plugin) io.Writer {
	var ins []io.Writer

	type outer interface {
		Stdout() io.Writer
	}

	for _, p := range plugs {
		if x, ok := p.(outer); ok {
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

	type outer interface {
		Stderr() io.Writer
	}

	for _, p := range plugs {
		if x, ok := p.(outer); ok {
			ins = append(ins, x.Stderr())
		}
	}

	if len(ins) == 0 {
		return os.Stderr
	}

	return io.MultiWriter(ins...)
}
