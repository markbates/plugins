package plugins

import (
	"errors"
	"fmt"
	"testing"
)

func Test_Errors(t *testing.T) {
	t.Parallel()

	exp := "boom"
	err := fmt.Errorf(exp)

	p := back("a")
	err = Wrap(p, err)

	var ex Error

	if !errors.As(err, &ex) {
		t.Fatalf("expected %s to be an Error", err)
	}

	_, ok := ex.Plugin.(back)
	if !ok {
		t.Fatalf("expected %T to be a back", ex.Plugin)
	}

	act := ex.Err.Error()

	if act != exp {
		t.Fatalf("expected %q to be %q", act, exp)
	}
}
