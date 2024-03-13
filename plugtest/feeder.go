package plugtest

import (
	"fmt"

	"github.com/markbates/plugins"
)

var _ plugins.Feeder = &Feeder{}

type Feeder struct {
	Plugins plugins.Plugins
}

func (f *Feeder) PluginName() string {
	return fmt.Sprintf("%T", f)
}

func (f *Feeder) PluginFeeder() plugins.FeederFn {
	if f == nil {
		return func() plugins.Plugins {
			return nil
		}
	}

	return func() plugins.Plugins {
		return f.Plugins
	}
}
