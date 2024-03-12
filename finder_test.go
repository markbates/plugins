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

	res, err := f.Find(plugs)
	r.NoError(err)
	r.NotNil(res)
	r.Len(res, 1)

	p := res[0]

	act := p.PluginName()
	r.Equal(exp, act)

	exp = "x/y/z"
	f = Background(path.Base(exp))

	res, err = f.Find(plugs)
	r.NoError(err)
	r.NotNil(res)
	r.Len(res, 1)

	p = res[0]

	act = p.PluginName()
	r.Equal(exp, act)
}

func Test_ByType(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	plugs := Plugins{
		stringPlug("a"),
		intPlug(1),
		stringPlug("b"),
		intPlug(2),
		stringPlug("c"),
		intPlug(3),
	}

	strs := ByType[stringPlug](plugs)
	r.Len(strs, 3)
	r.Equal("a", string(strs[0]))
	r.Equal("b", string(strs[1]))
	r.Equal("c", string(strs[2]))

	ints := ByType[intPlug](plugs)
	r.Len(ints, 3)
	r.Equal(1, int(ints[0]))
	r.Equal(2, int(ints[1]))
	r.Equal(3, int(ints[2]))

}
