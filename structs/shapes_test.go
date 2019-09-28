package structs

import (
	"math"
	"testing"
)

const float64EqualityThreshold = 0.01

func almostEqual(a, b float64) bool {
	return math.Abs(a-b) <= float64EqualityThreshold
}

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := rectangle.Perimeter()
	want := 40.0

	if !almostEqual(got, want) {
		t.Errorf("got %.2f hasArea %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{"Rectangle", Rectangle{12.0, 6.0}, 72.0},
		{"Circle", Circle{10.0}, 314.16},
		{"Triangle", Triangle{10.0, 10.0}, 50.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if !almostEqual(got, tt.hasArea) {
				t.Errorf("%q got %g hasArea %g", tt.name, got, tt.hasArea)
			}
		})
	}
}
