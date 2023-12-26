package arrays

func Sum(numbers []int) (sum int) {
    // range returns the index and the value
    for _, number := range numbers {
        sum += number
    }
    return
}
