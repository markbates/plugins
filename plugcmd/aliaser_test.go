package plugcmd

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_AliaserFn(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	exp := []string{"a", "b", "c"}
	fn := AliaserFn(func() []string {
		return exp
	})

	act := fn.CmdAliases()

	r.Equal(exp, act)

	r.Equal(fmt.Sprintf("%T", fn), fn.PluginName())
}
