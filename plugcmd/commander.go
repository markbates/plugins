package plugcmd

import "context"

// Commander is a plugin that is meant to be the beginning of
// a CLI application
type Commander interface {
	Main(ctx context.Context, root string, args []string) error
}

type CommanderFn func(ctx context.Context, root string, args []string) error

func (fn CommanderFn) Main(ctx context.Context, root string, args []string) error {
	return fn(ctx, root, args)
}

type NamedCommander interface {
	Commander
	Namer
}
