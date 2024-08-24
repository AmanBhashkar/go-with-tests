package structsmethodsandinterfaces

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want float64) {
		t.Helper()
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}

	}
	reactangle := Rectangle{10.0, 10.0}
	got := reactangle.Perimeter()
	want := 40.0
	assertCorrectMessage(t, got, want)
}
func TestArea(t *testing.T) {
	areaTests := []struct {
		shape   Shape
		hasArea float64
	}{
		{Rectangle{12, 6}, 72.0},
		{Circle{10}, 314.1592653589793},
		{Triangle{12, 6}, 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.hasArea {
			t.Errorf("got %g want %g", got, tt.hasArea)
		}

	}

}
