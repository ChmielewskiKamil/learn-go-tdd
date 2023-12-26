package arrays

import "testing"

func TestSum(t *testing.T) {
	t.Run("sum of 5 numbers", func(t *testing.T) {
        numbers := []int{1, 1, 1, 1, 1}

        want := 5
        got := Sum(numbers)

        // %v placeholder prints the inputs in a readable format
        if got != want {
            t.Errorf("got %d want %d given, %v", got, want, numbers)
        }
    })

    t.Run("sum of 3 numbers", func(t *testing.T) {
        numbers := []int{1, 2, 3}

        want := 6
        got := Sum(numbers)

        if got != want {
            t.Errorf("got %d want %d given, %v", got, want, numbers)
        }
    })
}
