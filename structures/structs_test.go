package main

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{2.0, 2.0}
	got := rectangle.Perimeter()
	want := 8.0

	if got != want {
		t.Errorf("got %.2f, want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, want: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, want: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 10, Height: 10}, want: 50},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.want {
				t.Errorf(" %#v got %.2f, want %.2f", tt.shape, got, tt.want)
			}
		})
	}

}
