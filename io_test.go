package plugins

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_StdIO(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	s := IO{}
	r.Equal(os.Stdin, s.Stdin())
	r.Equal(os.Stdout, s.Stdout())
	r.Equal(os.Stderr, s.Stderr())

	var out bytes.Buffer
	var err bytes.Buffer
	var in bytes.Reader

	s.In = &in
	r.Equal(&in, s.Stdin())

	s.Out = &out
	r.Equal(&out, s.Stdout())

	s.Err = &err
	r.Equal(&err, s.Stderr())
}

func Test_IO_String(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	oi := IO{}

	act := oi.String()
	exp := `{
  "stderr": "*os.File",
  "stdin": "*os.File",
  "stdout": "*os.File"
}`
	r.Equal(exp, act)

	var out bytes.Buffer
	var err bytes.Buffer
	var in bytes.Reader

	oi.In = &in
	oi.Out = &out
	oi.Err = &err

	act = oi.String()

	exp = `{
  "stderr": "*bytes.Buffer",
  "stdin": "*bytes.Reader",
  "stdout": "*bytes.Buffer"
}`
	r.Equal(exp, act)

}
