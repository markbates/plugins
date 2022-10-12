package plugcmd

import (
	"context"
	"flag"
	"io"

	"github.com/markbates/plugins"
)

// Commander is a plugin that is meant to be the beginning of
// a CLI application
type Commander interface {
	Main(ctx context.Context, root string, args []string) error
}

// Aliaser is a command that provides aliases for itself
type Aliaser interface {
	CmdAliases() []string
}

// Namer is a command that provides a different name for the
// command than the name of the Plugin
type Namer interface {
	CmdName() string
}

// SubCommander can be implemented to provide a list of plugins.Plugin
// that can be used as sub-commands of the current Plugin
type SubCommander interface {
	SubCommands() plugins.Plugins
}

// Describer is called by `Print` and can be
// implemented to print a short, single line,
// description of the plugin. `-h`
type Describer interface {
	Description() string
}

type FlagPrinter interface {
	PrintFlags(w io.Writer) error
}

type Flagger interface {
	Flags() (*flag.FlagSet, error)
}

type Hider interface {
	HidePlugin()
}

// UsagePrinter is called by `Print` and can be implemented
// to print a large block of usage information after the
// `Describer` interface is called. This is useful for printing
// flag information, links, and other messages to users.
type UsagePrinter interface {
	PrintUsage(w io.Writer) error
}
