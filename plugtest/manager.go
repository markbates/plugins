package plugtest

import (
	"fmt"
	"sync"

	"github.com/markbates/plugins"
)

var _ plugins.Needer = &Manager{}
var _ plugins.Feeder = &Manager{}
var _ plugins.Scoper = &Manager{}

type Manager struct {
	FeederFn plugins.FeederFn

	mu sync.RWMutex
}

func (m *Manager) PluginName() string {
	return fmt.Sprintf("%T", m)
}

func (m *Manager) WithPlugins(fn plugins.FeederFn) error {
	if m == nil {
		return fmt.Errorf("manager cannot be nil")
	}

	if fn == nil {
		return fmt.Errorf("feeder function cannot be nil")
	}

	m.mu.Lock()
	defer m.mu.Unlock()

	m.FeederFn = fn
	return nil
}

func (m *Manager) PluginFeeder() plugins.FeederFn {
	fn := func() plugins.Plugins {
		return nil
	}

	if m == nil {
		return fn
	}

	m.mu.RLock()
	defer m.mu.RUnlock()

	if m.FeederFn == nil {
		return fn
	}

	return m.FeederFn
}

func (m *Manager) ScopedPlugins() plugins.Plugins {
	return m.PluginFeeder()()
}
