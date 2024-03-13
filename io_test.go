package plugins_test

import (
	"os"
	"testing"

	"github.com/markbates/iox"
	. "github.com/markbates/plugins"
	"github.com/markbates/plugins/plugtest"
	"github.com/stretchr/testify/require"
)

func Test_Stdout(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	w := Stdout()
	r.NotNil(w)
	r.Equal(os.Stdout, w)

	bf1 := iox.Buffer{}
	bf2 := iox.Buffer{}
	bf3 := iox.Buffer{}

	plugs := Plugins{
		&plugtest.IO{IO: bf1.IO()},
		&plugtest.IO{IO: bf2.IO()},
		&plugtest.IO{IO: bf3.IO()},
	}

	w = Stdout(plugs...)
	_, err := w.Write([]byte("hi"))
	r.NoError(err)

	r.Equal("hi", bf1.Out.String())
	r.Equal("hi", bf2.Out.String())
	r.Equal("hi", bf3.Out.String())
}

func Test_Stderr(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	w := Stderr()
	r.NotNil(w)
	r.Equal(os.Stderr, w)

	bf1 := iox.Buffer{}
	bf2 := iox.Buffer{}
	bf3 := iox.Buffer{}

	plugs := Plugins{
		&plugtest.IO{IO: bf1.IO()},
		&plugtest.IO{IO: bf2.IO()},
		&plugtest.IO{IO: bf3.IO()},
	}

	w = Stderr(plugs...)
	_, err := w.Write([]byte("hi"))
	r.NoError(err)

	r.Equal("hi", bf1.Err.String())
	r.Equal("hi", bf2.Err.String())
	r.Equal("hi", bf3.Err.String())
}
