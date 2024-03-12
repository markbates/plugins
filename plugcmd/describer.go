package plugcmd

import (
	"fmt"

	"github.com/markbates/plugins"
)

// Describer is called by `Print` and can be
// implemented to print a short, single line,
// description of the plugin. `-h`
type Describer interface {
	Description() string
}

var _ Describer = DescriberFn(nil)
var _ plugins.Plugin = DescriberFn(nil)

// DescriberFn is a function that can be used to implement the Describer interface
type DescriberFn func() string

func (fn DescriberFn) Description() string {
	return fn()
}

func (fn DescriberFn) PluginName() string {
	return fmt.Sprintf("%T", fn)
}
