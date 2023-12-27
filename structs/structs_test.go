package structs

import "testing"

func TestArea(t *testing.T) {
    t.Run("rectangle area", func(t *testing.T){
        rectangle := Rectangle{Width: 2.5, Height: 5.0}
        got := rectangle.Area()
        want := 12.5

        if got != want {
            t.Errorf("got %.2f want %.2f", got, want)
        }
    })

    t.Run("circle area", func(t *testing.T) {
        circle := Circle{Radius: 5.0}
        got := circle.Area()
        want := 78.53981633974483

        if got != want {
            t.Errorf("got %g want %g", got, want)
        }
    })
}
