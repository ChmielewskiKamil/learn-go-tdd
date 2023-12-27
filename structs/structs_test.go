package structs

import "testing"

func TestArea(t *testing.T) {
	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}

	t.Run("rectangle area", func(t *testing.T) {
		rectangle := Rectangle{Width: 2.5, Height: 5.0}
		want := 12.5
		checkArea(t, rectangle, want)
	})

	t.Run("circle area", func(t *testing.T) {
		circle := Circle{Radius: 5.0}
		want := 78.53981633974483
        checkArea(t, circle, want)
	})
}
