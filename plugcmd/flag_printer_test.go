package plugcmd

import (
	"bytes"
	"fmt"
	"io"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_FlagPrinterFn(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	exp := "fun with flags"
	fn := FlagPrinterFn(func(w io.Writer) error {
		_, err := w.Write([]byte(exp))
		if err != nil {
			return err
		}
		return nil
	})

	bb := &bytes.Buffer{}

	err := fn.PrintFlags(bb)
	r.NoError(err)

	act := bb.String()
	r.Equal(exp, act)

	r.Equal(fmt.Sprintf("%T", fn), fn.PluginName())
}
