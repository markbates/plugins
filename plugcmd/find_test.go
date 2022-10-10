package plugcmd

import (
	"context"
	"testing"

	"github.com/markbates/plugins"
)

type mega struct {
	aliases []string
	name    string
	cmdName string
}

func (m mega) PluginName() string {
	if len(m.name) == 0 {
		return "mega"
	}
	return m.name
}

func (m mega) CmdName() string {
	if len(m.cmdName) == 0 {
		return "megacmd"
	}
	return m.cmdName
}

func (m mega) CmdAliases() []string {
	return m.aliases
}

func (m mega) Main(ctx context.Context, root string, args []string) error {
	return nil
}

func Test_Find(t *testing.T) {
	t.Parallel()
	m := mega{
		aliases: []string{"m"},
	}

	plugs := plugins.Plugins{
		background("a"),
		m,
		background("c"),
	}

	table := []string{"mega", "megacmd", "m"}
	for _, tt := range table {
		p := Find(tt, plugs)

		if p == nil {
			t.Fatalf("Expected to find plugin %s", tt)
		}
	}
}

func Test_FindFromArgs(t *testing.T) {
	t.Parallel()
	var m mega

	plugs := plugins.Plugins{
		background("a"),
		m,
		background("c"),
	}

	p := FindFromArgs([]string{"--flag", "mega", "--flag2"}, plugs)

	if p == nil {
		t.Fatalf("Expected to find plugin mega")
	}
}
