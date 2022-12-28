package shapes

import (
	"testing"
)

func TestPerimeter(t *testing.T) {

	checkPerimeter := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Perimeter()
		if got != want {
			t.Errorf("got %g, want %g", got, want)
		}
	}

	t.Run("rectangle", func(t *testing.T) {
		rectangle := Rectangle{12, 6}
		checkPerimeter(t, rectangle, 36.0)
	})

	t.Run("circle", func(t *testing.T) {
		circle := Circle{10}
		got := circle.Perimeter()
		want := 62.83185307179586
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}

	})
}

func TestArea0(t *testing.T) {
	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %g, want %g", got, want)
		}
	}
	t.Run("rectangle", func(t *testing.T) {
		rectangle := Rectangle{12, 6}
		checkArea(t, rectangle, 72.0)
	})

	t.Run("circle", func(t *testing.T) {
		circle := Circle{10}
		checkArea(t, circle, 314.1592653589793)
	})

}

// ----- Table Driven Tests -----
func TestArea(t *testing.T) {
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{12, 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{12, 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.hasArea {
			t.Errorf("got %g want %g", got, tt.hasArea)
		}
	}
}
