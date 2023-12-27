package structs

import "testing"

func TestArea(t *testing.T) {
	areaTests := []struct {
		shape Shape
		expectedArea  float64
	}{
        {shape: Rectangle{Width: 2.5, Height: 5.0}, expectedArea: 12.5},
        {shape: Circle{Radius: 5.0}, expectedArea: 78.53981633974483},
        {shape: Triangle{Base: 12, Height: 6}, expectedArea: 36.0},
	}
    
    for _, tt := range areaTests {
        calculatedArea := tt.shape.Area()
        if calculatedArea != tt.expectedArea {
            t.Errorf("Calculated area: %g, while it should be: %g", calculatedArea, tt.expectedArea)
        }
    }
}
