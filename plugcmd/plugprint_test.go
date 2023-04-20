package plugcmd

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/markbates/plugins"
	"github.com/stretchr/testify/require"
)

type stringPlugin string

func (s stringPlugin) PluginName() string {
	return string(s)
}

func (s stringPlugin) Description() string {
	return "string/plugin"
}

type cmd struct {
	name  string
	desc  string
	plugs plugins.Plugins
}

func (c cmd) ScopedPlugins() plugins.Plugins {
	return c.plugs
}

func (c cmd) Main(ctx context.Context, root string, args []string) error {
	return nil
}

func (c cmd) Description() string {
	if len(c.desc) > 0 {
		return c.desc
	}

	return fmt.Sprintf("Description of %s", c.name)
}

func (c cmd) PluginName() string {
	return c.name
}

func (c cmd) CmdAliases() []string {
	return []string{
		fmt.Sprintf("%s1", c.name),
		fmt.Sprintf("%s2", c.name),
	}
}

func (c cmd) PrintUsage(w io.Writer) error {
	s := fmt.Sprintf("This is how to use %s.", c.name)
	_, err := w.Write([]byte(s))
	return err
}

func (c cmd) SubCommands() []Commander {
	return []Commander{
		cmd{name: fmt.Sprintf("%s sub1", c.name)},
		cmd{name: fmt.Sprintf("%s sub2", c.name)},
		cmd{name: fmt.Sprintf("%s sub3", c.name)},
	}
}

func (cmd) PrintFlags(w io.Writer) error {
	_, err := w.Write([]byte("My Flags\n"))
	return err
}

func Test_Print(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	c := cmd{
		name: "main",
		plugs: plugins.Plugins{
			stringPlugin("one"),
			stringPlugin("two"),
		},
	}

	bb := &bytes.Buffer{}

	err := Print(bb, c)
	r.NoError(err)

	exp := `$ main
------
github.com/markbates/plugins/plugcmd.cmd

This is how to use main.

Aliases:
main1, main2

Flags:
My Flags

Available Commands:
  Command    Description
  -------    -----------
  main sub1  Description of main sub1
  main sub2  Description of main sub2
  main sub3  Description of main sub3

Using Plugins:
  Name  Description    Type
  ----  -----------    ----
  one   string/plugin  github.com/markbates/plugins/plugcmd.stringPlugin
  two   string/plugin  github.com/markbates/plugins/plugcmd.stringPlugin`

	exp = strings.TrimSpace(exp)

	act := bb.String()
	act = strings.TrimSpace(act)

	// fmt.Println(act)

	r.Equal(exp, act)
}
