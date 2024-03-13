package plugtest

import (
	"fmt"
	"io"
	"os"

	"github.com/markbates/iox"
	"github.com/markbates/plugins"
)

var _ plugins.IOSetable = &IO{}
var _ plugins.IOable = &IO{}
var _ plugins.Stderrer = &IO{}
var _ plugins.Stdiner = &IO{}
var _ plugins.Stdouter = &IO{}

type IO struct {
	IO plugins.IO
}

func (i *IO) PluginName() string {
	return fmt.Sprintf("%T", i)
}

func (i *IO) SetStdio(io plugins.IO) error {
	if i == nil {
		return fmt.Errorf("nil IO")
	}

	i.IO = io
	return nil
}

func (i *IO) Stdio() plugins.IO {
	if i == nil {
		return iox.IO{}
	}

	return i.IO
}

func (i *IO) Stderr() io.Writer {
	if i == nil {
		return os.Stderr
	}

	return i.IO.Stderr()
}

func (i *IO) Stdin() io.Reader {
	if i == nil {
		return os.Stdin
	}

	return i.IO.Stdin()
}

func (i *IO) Stdout() io.Writer {
	if i == nil {
		return os.Stdout
	}

	return i.IO.Stdout()
}
