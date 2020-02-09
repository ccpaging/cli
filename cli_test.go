package cli

import (
	"testing"
)

func mustPanic(t *testing.T, text string, fn func()) {
	defer func() {
		state := recover()
		if state == nil {
			t.Errorf(`case "%s" didn't panic`, text)
		}
	}()

	fn()
}

func TestNew(t *testing.T) {
	a := NewApp("smth")
	if a.Name != "smth" {
		t.Errorf("actual app name (%s) doesn't match passed (smth)", a.Name)
	}
}
