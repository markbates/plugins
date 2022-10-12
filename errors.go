package plugins

import (
	"bytes"
	"errors"
	"fmt"
	"path"
	"reflect"
	"strings"
)

// Wrap wraps an error with the plugin
// that caused the error.
func Wrap(p Plugin, err error) error {
	if err == nil {
		return nil
	}

	return Error{
		Plugin: p,
		Err:    err,
	}
}

// Error is used to wrap an error
// with the plugin that caused it.
type Error struct {
	Plugin Plugin
	Err    error
}

// Error returns a formatted error message
// that includes the `type` name of the plugin
// the plugin's name, and the underlying error.
func (e Error) Error() string {

	p := e.Plugin
	err := e.Err

	return fmt.Sprintf("%s (%s):\n\t%s", typeName(p), p.PluginName(), err)
}

func (e Error) Unwrap() error {
	type unwrapper interface {
		Unwrap() error
	}
	if _, ok := e.Err.(unwrapper); ok {
		return errors.Unwrap(e.Err)
	}

	return e.Err
}

func (e Error) As(target any) bool {
	ex, ok := target.(*Error)
	if !ok {
		// if the target is not an Error,
		// pass the underlying error up the chain
		// by calling errors.As with the underlying error
		// and the target error
		return errors.As(e.Err, target)
	}

	// set the target to the current error
	(*ex) = e
	return true
}

func (e Error) Is(target error) bool {
	if _, ok := target.(Error); ok {
		// return true if target is Error
		return true
	}

	// if not, pass the underlying error up the chain
	// by calling errors.Is with the underlying error
	// and the target error
	return errors.Is(e.Err, target)
}

func typeName(p Plugin) string {
	rv := reflect.Indirect(reflect.ValueOf(p))
	rt := reflect.TypeOf(rv.Interface())

	bb := &bytes.Buffer{}

	t := fmt.Sprintf("%T", p)
	if strings.HasPrefix(t, "*") {
		fmt.Fprint(bb, "*")
		t = strings.TrimPrefix(t, "*")
	}
	fmt.Fprintf(bb, "%s/", path.Dir(rt.PkgPath()))
	fmt.Fprint(bb, t)

	return bb.String()
}
