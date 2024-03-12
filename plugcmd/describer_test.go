package plugcmd

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_DescribeFn(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	exp := "description"
	fn := DescriberFn(func() string {
		return exp
	})

	act := fn.Description()

	r.Equal(exp, act)

	r.Equal(fmt.Sprintf("%T", fn), fn.PluginName())
}
