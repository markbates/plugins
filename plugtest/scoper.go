package plugtest

import (
	"fmt"

	"github.com/markbates/plugins"
)

var _ plugins.Scoper = &Scoper{}

type Scoper struct {
	Plugins plugins.Plugins
}

func (s *Scoper) PluginName() string {
	return fmt.Sprintf("%T", s)
}

func (s *Scoper) ScopedPlugins() plugins.Plugins {
	if s == nil {
		return nil
	}

	return s.Plugins
}
