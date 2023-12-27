package structs

import "testing"

func TestArea(t *testing.T) {
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{Width: 2.5, Height: 5.0}, 12.5},
        {Circle{Radius: 5.0}, 78.53981633974483},
	}
    
    for _, tt := range areaTests {
        got := tt.shape.Area()
        if got != tt.want {
            t.Errorf("got %g want %g", got, tt.want)
        }
    }
}
