package plugcmd

import (
	"flag"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_FlaggerFn(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	exp := flag.NewFlagSet("test", flag.ContinueOnError)

	fn := FlaggerFn(func() (*flag.FlagSet, error) {
		return exp, nil
	})

	act, err := fn.Flags()
	r.NoError(err)

	r.Equal(exp, act)

	r.Equal(fmt.Sprintf("%T", fn), fn.PluginName())
}
