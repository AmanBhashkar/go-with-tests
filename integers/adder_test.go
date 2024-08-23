package integers

import (
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 3)
	expect := 5
	if sum != expect {
		t.Errorf("expected %d but got %d", expect, sum)
	}
}
