package plugins_test

import (
	"testing"

	"github.com/markbates/plugins"
	"github.com/markbates/plugins/plugtest"
	"github.com/stretchr/testify/require"
)

func Test_PlugTest_AvailabilityChecker_Available(t *testing.T) {
	r := require.New(t)

	p := plugtest.AvailabilityChecker(true)
	r.True(p.PluginAvailable(""))
	r.Equal("plugtest.AvailabilityChecker", p.PluginName())
}

func Test_PlugTest_AvailabilityChecker_NotAvailable(t *testing.T) {
	r := require.New(t)

	p := plugtest.AvailabilityChecker(false)
	r.False(p.PluginAvailable(""))
	r.Equal("plugtest.AvailabilityChecker", p.PluginName())
}

func Test_PlugTest_Feeder(t *testing.T) {
	r := require.New(t)

	p := &plugtest.Feeder{}
	r.NotNil(p)
	r.Equal("*plugtest.Feeder", p.PluginName())

	feederFn := p.PluginFeeder()
	r.NotNil(feederFn)
}

func Test_PlugTest_Simple(t *testing.T) {
	r := require.New(t)

	p := plugtest.Simple(42)
	r.Equal("plugtest.Simple(42)", p.PluginName())
}

func Test_PlugTest_StringPlugin(t *testing.T) {
	r := require.New(t)

	p := plugtest.StringPlugin("test")
	r.Equal("test", p.PluginName())
}

func Test_PlugTest_Scoper(t *testing.T) {
	r := require.New(t)

	testPlugins := plugins.Plugins{plugtest.Simple(1), plugtest.Simple(2)}
	p := &plugtest.Scoper{Plugins: testPlugins}
	r.NotNil(p)
	r.Equal("*plugtest.Scoper", p.PluginName())

	scoped := p.ScopedPlugins()
	r.Equal(testPlugins, scoped)
}

func Test_PlugTest_Needer(t *testing.T) {
	r := require.New(t)

	p := &plugtest.Needer{}
	r.NotNil(p)
	r.Equal("plugtest.Needer", p.PluginName())

	// Test WithPlugins
	testFn := func() plugins.Plugins {
		return plugins.Plugins{plugtest.Simple(1)}
	}

	err := p.WithPlugins(testFn)
	r.NoError(err)
}

func Test_PlugTest_Needer_NilReceiver(t *testing.T) {
	r := require.New(t)

	var p *plugtest.Needer

	testFn := func() plugins.Plugins {
		return plugins.Plugins{plugtest.Simple(1)}
	}

	err := p.WithPlugins(testFn)
	r.Error(err)
	r.Contains(err.Error(), "nil Needer")
}

func Test_PlugTest_Needer_NilFeederFn(t *testing.T) {
	r := require.New(t)

	p := &plugtest.Needer{}

	err := p.WithPlugins(nil)
	r.Error(err)
	r.Contains(err.Error(), "nil FeederFn")
}
