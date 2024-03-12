package plugcmd

import (
	"flag"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_SetToSlice(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	var help bool
	var name string

	flags := flag.NewFlagSet("test", flag.ContinueOnError)
	flags.BoolVar(&help, "help", false, "help")
	flags.StringVar(&name, "name", "", "name")

	slice := SetToSlice(flags)

	exp := 2
	act := len(slice)
	r.Equal(exp, act)
}

func Test_CleanFlags(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	p := NamerFn(func() string {
		return "foo"
	})

	flags := Clean(p, []*flag.Flag{
		{Name: "my-flag"},
	})

	r.Len(flags, 1)

	f := flags[0]

	exp := "foo-my-flag"
	act := f.Name
	r.Equal(exp, act)
}

func Test_CleanFlagSet(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	p := stringPlug("foo")

	set := flag.NewFlagSet("", flag.ContinueOnError)
	set.String("my-flag", "", "")
	flags := CleanSet(p, set)

	r.Len(flags, 1)

	f := flags[0]

	exp := "foo-my-flag"
	act := f.Name
	r.Equal(exp, act)
}
