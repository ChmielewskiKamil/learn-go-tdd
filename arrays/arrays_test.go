package arrays

import (
	"slices"
	"testing"
)

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

func TestSumAll(t *testing.T) {
    got := SumAll([]int{1, 2}, []int{0, 9})
    want := []int{3, 9}

    if !slices.Equal(got, want) {
        t.Errorf("got %v want %v", got, want)
    }
}

func TestSumAllTails(t *testing.T) {
    got := SumAllTails([]int{1, 2}, []int{0, 9})
    want := []int{2, 9}

    if !slices.Equal(got, want) {
        t.Errorf("got %v want %v", got, want)
    }
}
