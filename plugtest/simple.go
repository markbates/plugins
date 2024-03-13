package plugtest

import (
	"fmt"

	"github.com/markbates/plugins"
)

var _ plugins.Plugin = Simple(0)

type Simple int

func (s Simple) PluginName() string {
	return fmt.Sprintf("%T(%d)", s, int(s))
}
