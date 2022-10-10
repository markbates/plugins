package plugins

import (
	"path"
	"testing"

	"github.com/stretchr/testify/require"
)

type plug string

func (p plug) PluginName() string {
	return string(p)
}

func Test_Finder(t *testing.T) {
	r := require.New(t)

	plugs := Plugins{
		plug("a"),
		plug("b"),
		plug("c"),
		plug("x/y/z"),
	}

	exp := "b"
	f := Background(exp)

	res := f.Find(plugs)
	r.NotNil(res)
	r.Len(res, 1)

	p := res[0]

	act := p.PluginName()
	r.Equal(exp, act)

	exp = "x/y/z"
	f = Background(path.Base(exp))

	res = f.Find(plugs)
	r.NotNil(res)
	r.Len(res, 1)

	p = res[0]

	act = p.PluginName()
	r.Equal(exp, act)
}
