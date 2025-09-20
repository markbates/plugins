package plugtest

import (
	"os"
	"testing"
	"testing/fstest"

	"github.com/markbates/iox"
	"github.com/markbates/plugins"
	"github.com/stretchr/testify/require"
)

func TestSimple_PluginName(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	s := Simple(42)
	name := s.PluginName()
	r.Equal("plugtest.Simple(42)", name)
}

func TestStringPlugin_PluginName(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	s := StringPlugin("test-plugin")
	name := s.PluginName()
	r.Equal("test-plugin", name)
}

func TestAvailabilityChecker_PluginName(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	ac := AvailabilityChecker(true)
	name := ac.PluginName()
	r.Equal("plugtest.AvailabilityChecker", name)
}

func TestAvailabilityChecker_PluginAvailable(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	acTrue := AvailabilityChecker(true)
	r.True(acTrue.PluginAvailable("/any/path"))

	acFalse := AvailabilityChecker(false)
	r.False(acFalse.PluginAvailable("/any/path"))
}

func TestFeeder_PluginName(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	f := &Feeder{}
	name := f.PluginName()
	r.Equal("*plugtest.Feeder", name)
}

func TestFeeder_PluginFeeder(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	tests := []struct {
		name       string
		feeder     *Feeder
		hasPlugins bool
	}{
		{
			name:       "nil feeder",
			feeder:     nil,
			hasPlugins: false,
		},
		{
			name:       "empty feeder",
			feeder:     &Feeder{},
			hasPlugins: false,
		},
		{
			name: "feeder with plugins",
			feeder: &Feeder{
				Plugins: plugins.Plugins{Simple(1), Simple(2)},
			},
			hasPlugins: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fn := tt.feeder.PluginFeeder()
			r.NotNil(fn)

			result := fn()
			if tt.hasPlugins {
				r.NotNil(result)
				r.Len(result, 2)
			} else {
				r.Nil(result)
			}
		})
	}
}

func TestNeeder_PluginName(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	n := &Needer{}
	name := n.PluginName()
	r.Equal("plugtest.Needer", name)
}

func TestNeeder_WithPlugins(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	tests := []struct {
		name        string
		needer      *Needer
		feederFn    plugins.FeederFn
		expectError bool
		errorMsg    string
	}{
		{
			name:        "nil needer",
			needer:      nil,
			feederFn:    func() plugins.Plugins { return nil },
			expectError: true,
			errorMsg:    "nil Needer",
		},
		{
			name:        "nil feeder function",
			needer:      &Needer{},
			feederFn:    nil,
			expectError: true,
			errorMsg:    "nil FeederFn",
		},
		{
			name:        "success",
			needer:      &Needer{},
			feederFn:    func() plugins.Plugins { return nil },
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.needer.WithPlugins(tt.feederFn)
			
			if tt.expectError {
				r.Error(err)
				if tt.errorMsg != "" {
					r.Contains(err.Error(), tt.errorMsg)
				}
			} else {
				r.NoError(err)
				r.NotNil(tt.needer.Fn)
			}
		})
	}
}

func TestScoper_PluginName(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	s := &Scoper{}
	name := s.PluginName()
	r.Equal("*plugtest.Scoper", name)
}

func TestScoper_ScopedPlugins(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	tests := []struct {
		name     string
		scoper   *Scoper
		expected int
	}{
		{
			name:     "nil scoper",
			scoper:   nil,
			expected: 0,
		},
		{
			name:     "empty scoper",
			scoper:   &Scoper{},
			expected: 0,
		},
		{
			name: "scoper with plugins",
			scoper: &Scoper{
				Plugins: plugins.Plugins{Simple(1), Simple(2)},
			},
			expected: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.scoper.ScopedPlugins()
			if tt.expected == 0 {
				r.Nil(result)
			} else {
				r.Len(result, tt.expected)
			}
		})
	}
}

func TestManager_PluginName(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	m := &Manager{}
	name := m.PluginName()
	r.Equal("*plugtest.Manager", name)
}

func TestManager_WithPlugins(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	tests := []struct {
		name        string
		manager     *Manager
		feederFn    plugins.FeederFn
		expectError bool
		errorMsg    string
	}{
		{
			name:        "nil manager",
			manager:     nil,
			feederFn:    func() plugins.Plugins { return nil },
			expectError: true,
			errorMsg:    "manager cannot be nil",
		},
		{
			name:        "nil feeder function",
			manager:     &Manager{},
			feederFn:    nil,
			expectError: true,
			errorMsg:    "feeder function cannot be nil",
		},
		{
			name:        "success",
			manager:     &Manager{},
			feederFn:    func() plugins.Plugins { return nil },
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.manager.WithPlugins(tt.feederFn)
			
			if tt.expectError {
				r.Error(err)
				if tt.errorMsg != "" {
					r.Contains(err.Error(), tt.errorMsg)
				}
			} else {
				r.NoError(err)
				r.NotNil(tt.manager.FeederFn)
			}
		})
	}
}

func TestManager_PluginFeeder(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	tests := []struct {
		name        string
		manager     *Manager
		expectsNil  bool
	}{
		{
			name:        "nil manager",
			manager:     nil,
			expectsNil:  true,
		},
		{
			name:        "manager with nil feeder",
			manager:     &Manager{},
			expectsNil:  true,
		},
		{
			name: "manager with feeder",
			manager: &Manager{
				FeederFn: func() plugins.Plugins { 
					return plugins.Plugins{Simple(1)} 
				},
			},
			expectsNil: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fn := tt.manager.PluginFeeder()
			r.NotNil(fn)
			
			result := fn()
			if tt.expectsNil {
				r.Nil(result)
			} else {
				r.NotNil(result)
			}
		})
	}
}

func TestManager_ScopedPlugins(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	m := &Manager{}
	result := m.ScopedPlugins()
	r.Nil(result)
}

func TestIO_PluginName(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	io := &IO{}
	name := io.PluginName()
	r.Equal("*plugtest.IO", name)
}

func TestIO_SetStdio(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	tests := []struct {
		name        string
		io          *IO
		stdio       plugins.IO
		expectError bool
		errorMsg    string
	}{
		{
			name:        "nil IO",
			io:          nil,
			stdio:       iox.IO{},
			expectError: true,
			errorMsg:    "nil IO",
		},
		{
			name:        "success",
			io:          &IO{},
			stdio:       iox.IO{},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.io.SetStdio(tt.stdio)
			
			if tt.expectError {
				r.Error(err)
				if tt.errorMsg != "" {
					r.Contains(err.Error(), tt.errorMsg)
				}
			} else {
				r.NoError(err)
				r.Equal(tt.stdio, tt.io.IO)
			}
		})
	}
}

func TestIO_Stdio(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	// Test nil IO
	var nilIO *IO
	stdio := nilIO.Stdio()
	r.NotNil(stdio)

	// Test IO with configured stdio
	io := &IO{}
	expectedIO := iox.IO{}
	io.IO = expectedIO
	result := io.Stdio()
	r.Equal(expectedIO, result)
}

func TestIO_Stderr(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	// Test nil IO
	var nilIO *IO
	stderr := nilIO.Stderr()
	r.Equal(os.Stderr, stderr)

	// Test configured IO
	io := &IO{}
	buffer := iox.Buffer{}
	io.IO = buffer.IO()
	result := io.Stderr()
	r.NotNil(result)
}

func TestIO_Stdin(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	// Test nil IO
	var nilIO *IO
	stdin := nilIO.Stdin()
	r.Equal(os.Stdin, stdin)

	// Test configured IO
	io := &IO{}
	buffer := iox.Buffer{}
	io.IO = buffer.IO()
	result := io.Stdin()
	r.NotNil(result)
}

func TestIO_Stdout(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	// Test nil IO
	var nilIO *IO
	stdout := nilIO.Stdout()
	r.Equal(os.Stdout, stdout)

	// Test configured IO
	io := &IO{}
	buffer := iox.Buffer{}
	io.IO = buffer.IO()
	result := io.Stdout()
	r.NotNil(result)
}

func TestFSable_PluginName(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	fs := &FSable{}
	name := fs.PluginName()
	r.Equal("*plugtest.FSable", name)
}

func TestFSable_FileSystem(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	tests := []struct {
		name        string
		fsable      *FSable
		expectError bool
		errorMsg    string
	}{
		{
			name:        "nil FSable",
			fsable:      nil,
			expectError: true,
			errorMsg:    "FSable instance cannot be nil",
		},
		{
			name:        "FSable with nil filesystem",
			fsable:      &FSable{},
			expectError: true,
			errorMsg:    "no filesystem configured",
		},
		{
			name: "FSable with filesystem",
			fsable: &FSable{
				FS: fstest.MapFS{
					"file.txt": &fstest.MapFile{Data: []byte("test")},
				},
			},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs, err := tt.fsable.FileSystem()
			
			if tt.expectError {
				r.Error(err)
				if tt.errorMsg != "" {
					r.Contains(err.Error(), tt.errorMsg)
				}
				r.Nil(fs)
			} else {
				r.NoError(err)
				r.NotNil(fs)
			}
		})
	}
}

func TestFSable_SetFileSystem(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	tests := []struct {
		name        string
		fsable      *FSable
		fs          *fstest.MapFS
		expectError bool
		errorMsg    string
	}{
		{
			name:        "nil FSable",
			fsable:      nil,
			fs:          &fstest.MapFS{},
			expectError: true,
			errorMsg:    "FSable instance cannot be nil",
		},
		{
			name:        "nil filesystem",
			fsable:      &FSable{},
			fs:          nil,
			expectError: true,
			errorMsg:    "filesystem cannot be nil",
		},
		{
			name:        "success",
			fsable:      &FSable{},
			fs:          &fstest.MapFS{"file.txt": &fstest.MapFile{Data: []byte("test")}},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var err error
			if tt.fs != nil {
				err = tt.fsable.SetFileSystem(*tt.fs)
			} else {
				err = tt.fsable.SetFileSystem(nil)
			}
			
			if tt.expectError {
				r.Error(err)
				if tt.errorMsg != "" {
					r.Contains(err.Error(), tt.errorMsg)
				}
			} else {
				r.NoError(err)
				r.NotNil(tt.fsable.FS)
			}
		})
	}
}