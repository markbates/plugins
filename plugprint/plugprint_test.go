package plugprint

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/markbates/plugins"
	"github.com/stretchr/testify/require"
)

type cmd struct {
	name string
}

func (c cmd) Description() string {
	return fmt.Sprintf("%s Description", c.name)
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
	s := fmt.Sprintf("Usage of %s:", c.name)
	_, err := w.Write([]byte(s))
	return err
}

func (c cmd) SubCommands() plugins.Plugins {
	return plugins.Plugins{
		cmd{name: fmt.Sprintf("%s sub1", c.name)},
		cmd{name: fmt.Sprintf("%s sub2", c.name)},
		cmd{name: fmt.Sprintf("%s sub3", c.name)},
	}
}

func (cmd) PrintFlags(w io.Writer) error {
	_, err := w.Write([]byte(`My Flags`))
	return err
}

func Test_Print(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	c := cmd{
		name: "main",
	}

	bb := &bytes.Buffer{}

	err := Print(bb, c)
	r.NoError(err)

	exp := `main Description

$ main
------
github.com/markbates/plugins/plugprint.cmd

Aliases:
main1, main2

Usage of main:
Flags:
My Flags
Available Commands:
  Command    Description
  -------    -----------
  main sub1  main sub1 Description
  main sub2  main sub2 Description
  main sub3  main sub3 Description

Using Plugins:
  Name  Description       Type
  ----  -----------       ----
  main  main Description  github.com/markbates/plugins/plugprint.cmd`

	act := bb.String()
	act = strings.TrimSpace(act)

	// fmt.Println(act)
	r.Equal(exp, act)
}
