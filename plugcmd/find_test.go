package plugcmd

import (
	"context"
	"testing"

	"github.com/markbates/plugins"
	"github.com/stretchr/testify/require"
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
	r := require.New(t)

	p := Find("mega", plugins.Plugins{})
	r.Nil(p)

	m := mega{
		aliases: []string{"m"},
	}

	plugs := plugins.Plugins{
		stringPlug("a"),
		m,
		stringPlug("c"),
	}

	table := []string{"mega", "megacmd", "m"}
	for _, tt := range table {
		p = Find(tt, plugs)

		r.NotNil(p)
	}
}

func Test_FindFromArgs(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	p := FindFromArgs([]string{}, plugins.Plugins{})
	r.Nil(p)

	var m mega
	plugs := plugins.Plugins{
		stringPlug("a"),
		m,
		stringPlug("c"),
	}

	p = FindFromArgs([]string{"--flag", "mega", "--flag2"}, plugs)
	r.NotNil(p)
	r.Equal(m, p)
}
