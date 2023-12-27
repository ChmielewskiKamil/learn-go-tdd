package structs

import "testing"

func TestArea(t *testing.T) {
	areaTests := []struct {
		testName     string
		shape        Shape
		expectedArea float64
	}{
		{testName: "Rectangle", shape: Rectangle{Width: 2.5, Height: 5.0}, expectedArea: 12.5},
		{testName: "Circle", shape: Circle{Radius: 5.0}, expectedArea: 78.53981633974483},
		{testName: "Triangle", shape: Triangle{Base: 12, Height: 6}, expectedArea: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.testName, func(t *testing.T) {
			calculatedArea := tt.shape.Area()
			if calculatedArea != tt.expectedArea {
				t.Errorf("Calculated area: %g, while it should be: %g", calculatedArea, tt.expectedArea)
			}
		})
	}
}
