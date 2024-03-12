package plugcmd

import (
	"context"
	"fmt"

	"github.com/markbates/plugins"
)

// Commander is a plugin that is meant to be the beginning of
// a CLI application
type Commander interface {
	plugins.Plugin
	Main(ctx context.Context, root string, args []string) error
}

var _ Commander = CommanderFn(nil)

// CommanderFn is a function that can be used to implement the Commander interface
type CommanderFn func(ctx context.Context, root string, args []string) error

func (fn CommanderFn) Main(ctx context.Context, root string, args []string) error {
	return fn(ctx, root, args)
}

func (fn CommanderFn) PluginName() string {
	return fmt.Sprintf("%T", fn)
}

type NamedCommander interface {
	Commander
	Namer
}
