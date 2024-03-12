package plugcmd

import (
	"context"
	"fmt"
	"io"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_CommanderFn(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	fn := CommanderFn(func(ctx context.Context, root string, args []string) error {
		return io.EOF
	})

	err := fn.Main(context.Background(), "", []string{})
	r.Equal(io.EOF, err)

	r.Equal(fmt.Sprintf("%T", fn), fn.PluginName())
}
