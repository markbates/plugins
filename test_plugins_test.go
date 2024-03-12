package plugins

import (
	"fmt"
	"io/fs"
	"strings"

	"github.com/markbates/iox"
)

var _ Plugin = stringPlug("")

type stringPlug string

func (b stringPlug) PluginName() string {
	return string(b)
}

type intPlug int

func (b intPlug) PluginName() string {
	return fmt.Sprintf("%T", b)
}

var _ IOSetable = &ioPlug{}
var _ IOable = &ioPlug{}
var _ Plugin = &ioPlug{}
var _ Stdioer = &ioPlug{}

type ioPlug struct {
	iox.IO
}

func (i *ioPlug) SetStdio(oi iox.IO) error {
	i.IO = oi
	return nil
}

func (i *ioPlug) Stdio() iox.IO {
	return i.IO
}

func (i *ioPlug) PluginName() string {
	return "ioPlug"
}

var _ Feeder = &needFeedPlug{}
var _ Needer = &needFeedPlug{}
var _ Plugin = &needFeedPlug{}

type needFeedPlug struct {
	Fn FeederFn
}

func (n *needFeedPlug) WithPlugins(fn FeederFn) error {
	if n == nil {
		return fmt.Errorf("no needFeedPlug")
	}

	if fn == nil {
		return fmt.Errorf("no FeederFn")
	}

	n.Fn = fn
	return nil
}

func (n *needFeedPlug) PluginFeeder() (FeederFn, error) {
	if n == nil || n.Fn == nil {
		return nil, fmt.Errorf("no FeederFn set")
	}

	return n.Fn, nil
}

func (n *needFeedPlug) PluginName() string {
	return "neederPlug"
}

var _ FSSetable = &fsPlug{}
var _ FSable = &fsPlug{}
var _ Plugin = &fsPlug{}

type fsPlug struct {
	fs.FS
}

func (f *fsPlug) SetFileSystem(fs fs.FS) error {
	if f == nil {
		return fmt.Errorf("no fsPlug")
	}

	if fs == nil {
		return fmt.Errorf("no fs.FS")
	}

	f.FS = fs
	return nil
}

func (f *fsPlug) FileSystem() (fs.FS, error) {
	if f == nil {
		return nil, fmt.Errorf("no fsPlug")
	}

	if f.FS == nil {
		return nil, fmt.Errorf("no fs.FS")
	}

	return f.FS, nil
}

func (f *fsPlug) PluginName() string {
	return "fsPlug"
}

var _ AvailabilityChecker = availPlug("")
var _ Plugin = availPlug("")

type availPlug string

func (a availPlug) PluginAvailable(root string) bool {
	return strings.HasPrefix(string(a), root)
}

func (a availPlug) PluginName() string {
	return "availPlug"
}
