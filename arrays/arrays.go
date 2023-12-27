package arrays

func Sum(numbers []int) (sum int) {
    // range returns the index and the value
    for _, number := range numbers {
        sum += number
    }
    return
}

func SumAll(slices ...[]int) (individualSums []int){
    for _, slice := range slices {
        individualSums = append(individualSums, Sum(slice))
    }
    return 
}
