package plugins_test

import (
	"path"
	"testing"

	. "github.com/markbates/plugins"
	"github.com/markbates/plugins/plugtest"
	"github.com/stretchr/testify/require"
)

func Test_Finder(t *testing.T) {
	r := require.New(t)

	plugs := Plugins{
		plugtest.StringPlugin("a"),
		plugtest.StringPlugin("b"),
		plugtest.StringPlugin("c"),
		plugtest.StringPlugin("x/y/z"),
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
		plugtest.StringPlugin("a"),
		plugtest.Simple(1),
		plugtest.StringPlugin("b"),
		plugtest.Simple(2),
		plugtest.StringPlugin("c"),
		plugtest.Simple(3),
	}

	strs := ByType[plugtest.StringPlugin](plugs)
	r.Len(strs, 3)
	r.Equal("a", string(strs[0]))
	r.Equal("b", string(strs[1]))
	r.Equal("c", string(strs[2]))

	ints := ByType[plugtest.Simple](plugs)
	r.Len(ints, 3)
	r.Equal(1, int(ints[0]))
	r.Equal(2, int(ints[1]))
	r.Equal(3, int(ints[2]))

}
