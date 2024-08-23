package arrayAndSlices

func Sum(numbers []int) int {
	result := 0
	// for i := 0; i < len(numbers); i++ {
	// 	result += numbers[i]
	// }
	for _, number := range numbers {
		result += number
	}

	return result
}

func SumAll(numbersToSum ...[]int) []int {
	// lengthOfNumbers := len(numbersToSum)
	// sums := make([]int, lengthOfNumbers)
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {

			tails := numbers[1:]
			sums = append(sums, Sum(tails))
		}
	}
	return sums
}
