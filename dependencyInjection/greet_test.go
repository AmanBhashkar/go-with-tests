package dependencyinjection

import (
	"bytes"
	"testing"
)

func TestGree(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Aman")
	got := buffer.String()
	want := "Hello, Aman"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
