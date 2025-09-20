package plugins_test

import (
	"errors"
	"io/fs"
	"testing"
	"testing/fstest"

	"github.com/markbates/iox"
	. "github.com/markbates/plugins"
	"github.com/markbates/plugins/plugtest"
	"github.com/stretchr/testify/require"
)

func Test_Plugins_SetStdio_Error(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	// Create a plugin that will fail SetStdio
	failingPlugin := &FailingIOPlugin{}

	plugs := Plugins{failingPlugin}

	// Create a valid IO
	io := iox.Discard()

	// Try to set stdio - should fail
	err := plugs.SetStdio(io)
	r.Error(err)
	r.Contains(err.Error(), "failed to set stdio for plugin")
}

// Test helper plugins

type FailingIOPlugin struct{}

func (p *FailingIOPlugin) PluginName() string {
	return "failing-io-plugin"
}

func (p *FailingIOPlugin) SetStdio(io IO) error {
	return errors.New("simulated IO error")
}

func Test_Plugins_SetFileSystem_Error(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	// Test with nil filesystem
	plugs := Plugins{&plugtest.FSable{}}
	err := plugs.SetFileSystem(nil)
	r.Error(err)
	r.Contains(err.Error(), "no fs.FS provided")
}

func Test_Plugins_SetFileSystem_PluginError(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	// Create a filesystem that will cause an error
	fsys := fstest.MapFS{
		"test.txt": &fstest.MapFile{Data: []byte("test")},
	}

	// Create a failing plugin
	failingPlugin := &FailingFSPlugin{}
	plugs := Plugins{failingPlugin}

	err := plugs.SetFileSystem(fsys)
	r.Error(err)
	r.Contains(err.Error(), "failed to set filesystem for plugin")
}

func Test_Plugins_WithPlugins_Error(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	// Test with nil FeederFn
	plugs := Plugins{&plugtest.Manager{}}
	err := plugs.WithPlugins(nil)
	r.Error(err)
	r.Contains(err.Error(), "no FeederFn provided")
}

func Test_Plugins_WithPlugins_PluginError(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	// Create a failing needer plugin
	failingPlugin := &FailingNeederPlugin{}
	plugs := Plugins{failingPlugin}

	fn := func() Plugins { return plugs }
	err := plugs.WithPlugins(fn)
	r.Error(err)
	r.Contains(err.Error(), "failed to set plugins for needer")
}

func Test_Plugins_Available_WithChecker(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	// Create plugins with availability checkers
	available := &AvailablePlugin{available: true}
	unavailable := &AvailablePlugin{available: false}
	normal := plugtest.Simple(1)

	plugs := Plugins{available, unavailable, normal}

	result := plugs.Available("/some/path")
	r.Len(result, 2) // available plugin + normal plugin
	r.Contains(result, available)
	r.Contains(result, normal)
	r.NotContains(result, unavailable)
}

type FailingFSPlugin struct{}

func (p *FailingFSPlugin) PluginName() string {
	return "failing-fs-plugin"
}

func (p *FailingFSPlugin) SetFileSystem(fs fs.FS) error {
	return errors.New("simulated filesystem error")
}

type FailingNeederPlugin struct{}

func (p *FailingNeederPlugin) PluginName() string {
	return "failing-needer-plugin"
}

func (p *FailingNeederPlugin) WithPlugins(fn FeederFn) error {
	return errors.New("simulated needer error")
}

type AvailablePlugin struct {
	available bool
}

func (p *AvailablePlugin) PluginName() string {
	return "available-plugin"
}

func (p *AvailablePlugin) PluginAvailable(root string) bool {
	return p.available
}
