package plugins_test

import (
	"testing"

	. "github.com/markbates/plugins"
	"github.com/markbates/plugins/plugtest"
	"github.com/stretchr/testify/require"
)

func Test_Plugins_Validate(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		plugins     Plugins
		expectError bool
		errorMsg    string
	}{
		{
			name:        "empty plugins",
			plugins:     Plugins{},
			expectError: true,
			errorMsg:    "no plugins provided",
		},
		{
			name:        "nil plugin",
			plugins:     Plugins{nil},
			expectError: true,
			errorMsg:    "plugin at index 0 is nil",
		},
		{
			name: "duplicate names",
			plugins: Plugins{
				plugtest.StringPlugin("same"),
				plugtest.StringPlugin("same"),
			},
			expectError: true,
			errorMsg:    "duplicate plugin name: same",
		},
		{
			name: "valid plugins",
			plugins: Plugins{
				plugtest.StringPlugin("first"),
				plugtest.StringPlugin("second"),
				plugtest.Simple(1),
			},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			err := tt.plugins.Validate()

			if tt.expectError {
				r.Error(err)
				if tt.errorMsg != "" {
					r.Contains(err.Error(), tt.errorMsg)
				}
			} else {
				r.NoError(err)
			}
		})
	}
}

func Test_Plugins_Names(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	plugs := Plugins{
		plugtest.StringPlugin("first"),
		plugtest.StringPlugin("second"),
		plugtest.Simple(1),
	}

	names := plugs.Names()
	r.Len(names, 3)
	r.Contains(names, "first")
	r.Contains(names, "second")
	r.Contains(names, "plugtest.Simple(1)")
}

func BenchmarkByType(b *testing.B) {
	// Create a large slice of mixed plugins
	plugs := make(Plugins, 1000)
	for i := 0; i < 500; i++ {
		plugs[i] = plugtest.Simple(i)
	}
	for i := 500; i < 1000; i++ {
		plugs[i] = plugtest.StringPlugin("plugin")
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ByType[plugtest.Simple](plugs)
	}
}

func BenchmarkPlugins_Validate(b *testing.B) {
	plugs := make(Plugins, 100)
	for i := 0; i < 100; i++ {
		plugs[i] = plugtest.Simple(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = plugs.Validate()
	}
}

func BenchmarkPlugins_Available(b *testing.B) {
	plugs := make(Plugins, 100)
	for i := 0; i < 100; i++ {
		plugs[i] = plugtest.Simple(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = plugs.Available("/some/path")
	}
}
