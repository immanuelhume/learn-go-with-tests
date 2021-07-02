package main

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(arrs ...[]int) []int {
	var sums []int
	for _, numbers := range arrs {
		sums = append(sums, Sum(numbers))
	}
	return sums
}

func SumAllTails(arrs ...[]int) []int {
	var tails []int
	for _, numbers := range arrs {
		sum := 0
		if len(numbers) > 0 {
			sum = Sum(numbers[1:])
		}
		tails = append(tails, sum)
	}
	return tails
}
