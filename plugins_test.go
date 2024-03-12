package plugins

import (
	"sort"
	"strings"
	"testing"
	"testing/fstest"

	"github.com/markbates/iox"
	"github.com/stretchr/testify/require"
)

func Test_Plugins_Sort(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	plugs := Plugins{
		stringPlug("b"),
		stringPlug("c"),
		stringPlug("a"),
	}

	sort.Sort(plugs)

	exp := []string{"a", "b", "c"}
	var act []string
	for _, p := range plugs {
		act = append(act, p.PluginName())
	}

	a := strings.Join(act, ",")
	e := strings.Join(exp, ",")
	r.Equal(e, a)
}

func Test_Plugins_PluginFeeder(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	plugs := Plugins{
		stringPlug("a"),
		stringPlug("b"),
		stringPlug("c"),
	}

	fn, err := plugs.PluginFeeder()
	r.NoError(err)

	act, err := fn()
	r.NoError(err)

	r.Equal(plugs, act)
}

func Test_Plugins_ScopedPlugins(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	plugs := Plugins{
		stringPlug("a"),
		stringPlug("b"),
		stringPlug("c"),
	}

	act, err := plugs.ScopedPlugins()
	r.NoError(err)

	r.Equal(plugs, act)
}

func Test_Plugins_SetStdio(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	plugs := Plugins{
		&ioPlug{},
		&ioPlug{},
		&ioPlug{},
	}

	io := iox.IO{}
	err := plugs.SetStdio(io)
	r.NoError(err)

	for _, p := range plugs {
		ip, ok := p.(*ioPlug)
		r.True(ok)
		r.Equal(io, ip.IO)
	}
}

func Test_Plugins_WithPlugins(t *testing.T) {
	t.Parallel()

	t.Run("nil FeederFn", func(t *testing.T) {
		r := require.New(t)

		plugs := Plugins{
			&needFeedPlug{},
		}

		err := plugs.WithPlugins(nil)
		r.Error(err)

	})
	t.Run("success", func(t *testing.T) {
		r := require.New(t)
		plugs := Plugins{
			stringPlug("a"),
			&needFeedPlug{},
			&needFeedPlug{},
			&needFeedPlug{},
		}

		fn := func() (Plugins, error) {
			return plugs, nil
		}

		err := plugs.WithPlugins(fn)
		r.NoError(err)

		for _, p := range plugs {
			ip, ok := p.(*needFeedPlug)
			if !ok {
				continue
			}

			fn, err := ip.PluginFeeder()
			r.NoError(err)

			act, err := fn()
			r.NoError(err)

			r.Equal(plugs, act)
		}
	})
}

func Test_Plugins_SetFileSystem(t *testing.T) {
	t.Parallel()

	t.Run("nil FS", func(t *testing.T) {
		r := require.New(t)

		plugs := Plugins{
			&fsPlug{},
		}

		err := plugs.SetFileSystem(nil)
		r.Error(err)

	})

	t.Run("success", func(t *testing.T) {
		r := require.New(t)

		plugs := Plugins{
			stringPlug("a"),
			&fsPlug{},
			&fsPlug{},
		}

		fs := fstest.MapFS{}
		err := plugs.SetFileSystem(fs)
		r.NoError(err)

		for _, p := range plugs {
			ip, ok := p.(*fsPlug)
			if !ok {
				continue
			}

			r.Equal(fs, ip.FS)
		}
	})
}

func Test_Plugins_Find(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	plugs := Plugins{
		stringPlug("a"),
		stringPlug("b"),
	}

	fn := func(plugs Plugins) (Plugins, error) {
		return plugs, nil
	}

	act, err := plugs.Find(fn)
	r.NoError(err)

	r.Equal(plugs, act)

}

func Test_Plugins_Available(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	plugs := Plugins{
		intPlug(1),
		availPlug("a/b/c/d.md"),
		availPlug("a/b/e.md"),
		availPlug("x/y/z.md"),
	}

	act := plugs.Available("a/b")
	r.Len(act, 3)

	ip, ok := act[0].(intPlug)
	r.True(ok)
	r.Equal(1, int(ip))

	ap, ok := act[1].(availPlug)
	r.True(ok)
	r.Equal("a/b/c/d.md", string(ap))

	ap, ok = act[2].(availPlug)
	r.True(ok)
	r.Equal("a/b/e.md", string(ap))

}
