package structs

import "testing"

func TestArea(t *testing.T) {
    t.Run("rectangle area", func(t *testing.T){
        got := Area(2.5, 5.0)
        want := 12.5

        if got != want {
            t.Errorf("got %.2f want %.2f", got, want)
        }
    })
}
