package plugcmd

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_NamerFn(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	exp := "name"
	fn := NamerFn(func() string {
		return exp
	})

	act := fn.CmdName()

	r.Equal(exp, act)

	r.Equal(fmt.Sprintf("%T", fn), fn.PluginName())

}
