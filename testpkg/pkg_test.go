package testpkg

import "testing"

func TestFunction(t *testing.T) {
	if Function() != "function" {
		t.Fatal("NO FUNCTION FOR YOU")
	}
}
