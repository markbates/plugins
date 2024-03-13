package plugtest

import (
	"fmt"
	"sync"

	"github.com/markbates/plugins"
)

var _ plugins.Needer = &Needer{}

type Needer struct {
	Fn plugins.FeederFn

	mu sync.Mutex
}

func (n *Needer) PluginName() string {
	return "plugtest.Needer"
}

func (n *Needer) WithPlugins(p plugins.FeederFn) error {
	if n == nil {
		return fmt.Errorf("nil Needer")
	}

	if p == nil {
		return fmt.Errorf("nil FeederFn")
	}

	n.mu.Lock()
	defer n.mu.Unlock()

	n.Fn = p

	return nil
}
