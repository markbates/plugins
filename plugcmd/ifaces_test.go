package plugcmd

import (
	"context"

	"github.com/markbates/plugins"
)

type stringPlug string

func (b stringPlug) PluginName() string {
	return string(b)
}

var _ plugins.Plugin = aliaser{}
var _ Aliaser = aliaser{}

type aliaser []string

func (a aliaser) PluginName() string {
	return "aliaser"
}

func (a aliaser) CmdAliases() []string {
	return []string(a)
}

func (a aliaser) Main(ctx context.Context, root string, args []string) error {
	return nil
}

var _ plugins.Plugin = namer("")
var _ Namer = namer("")

type namer string

func (namer) PluginName() string {
	return "namer"
}

func (n namer) CmdName() string {
	return string(n)
}

func (n namer) Main(ctx context.Context, root string, args []string) error {
	return nil
}

var _ plugins.Plugin = commander(nil)
var _ Commander = commander(nil)

func (c commander) PluginName() string {
	return "commander"
}

type commander func(ctx context.Context, root string, args []string) error

func (c commander) Main(ctx context.Context, root string, args []string) error {
	return c(ctx, root, args)
}
