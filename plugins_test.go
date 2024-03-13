package plugins_test

import (
	"sort"
	"testing"
	"testing/fstest"

	"github.com/markbates/iox"
	. "github.com/markbates/plugins"
	"github.com/markbates/plugins/plugtest"
	"github.com/stretchr/testify/require"
)

func Test_Plugins_Sort(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	plugs := Plugins{
		plugtest.Simple(2),
		plugtest.Simple(1),
	}

	sort.Sort(plugs)

	r.Equal(1, int(plugs[0].(plugtest.Simple)))
	r.Equal(2, int(plugs[1].(plugtest.Simple)))
}

func Test_Plugins_PluginFeeder(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	plugs := Plugins{
		plugtest.Simple(1),
	}

	fn := plugs.PluginFeeder()
	r.NotNil(fn)

	act := fn()
	r.Equal(plugs, act)
}

func Test_Plugins_ScopedPlugins(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	plugs := Plugins{
		plugtest.Simple(1),
	}

	act := plugs.ScopedPlugins()

	r.Equal(plugs, act)
}

func Test_Plugins_SetStdio(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	plugs := Plugins{
		&plugtest.IO{},
		&plugtest.IO{},
	}

	io := iox.IO{}
	err := plugs.SetStdio(io)
	r.NoError(err)

	for _, p := range plugs {
		ip, ok := p.(*plugtest.IO)
		r.True(ok)
		r.Equal(io, ip.IO)
	}
}

func Test_Plugins_WithPlugins(t *testing.T) {
	t.Parallel()

	t.Run("nil FeederFn", func(t *testing.T) {
		r := require.New(t)

		plugs := Plugins{
			&plugtest.Needer{},
		}

		err := plugs.WithPlugins(nil)
		r.Error(err)

	})
	t.Run("success", func(t *testing.T) {
		r := require.New(t)
		plugs := Plugins{
			plugtest.Simple(1),
			&plugtest.Needer{},
			&plugtest.Needer{},
			&plugtest.Needer{},
		}

		fn := func() Plugins {
			return plugs
		}

		err := plugs.WithPlugins(fn)
		r.NoError(err)

		for _, p := range plugs {
			ip, ok := p.(*plugtest.Needer)
			if !ok {
				continue
			}

			fn := ip.Fn
			r.NotNil(fn)

			act := fn()

			r.Equal(plugs, act)
		}
	})
}

func Test_Plugins_SetFileSystem(t *testing.T) {
	t.Parallel()

	t.Run("nil FS", func(t *testing.T) {
		r := require.New(t)

		plugs := Plugins{
			&plugtest.FSable{},
		}

		err := plugs.SetFileSystem(nil)
		r.Error(err)

	})

	t.Run("success", func(t *testing.T) {
		r := require.New(t)

		plugs := Plugins{
			plugtest.Simple(1),
			&plugtest.FSable{},
			&plugtest.FSable{},
		}

		fs := fstest.MapFS{}
		err := plugs.SetFileSystem(fs)
		r.NoError(err)

		for _, p := range plugs {
			ip, ok := p.(*plugtest.FSable)
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
		plugtest.Simple(1),
		plugtest.Simple(2),
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
		plugtest.AvailabilityChecker(true),
		plugtest.AvailabilityChecker(false),
		plugtest.AvailabilityChecker(true),
		plugtest.AvailabilityChecker(false),
	}

	act := plugs.Available("a/b")
	r.Len(act, 2)
}
