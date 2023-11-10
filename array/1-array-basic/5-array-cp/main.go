package main

func EvenNumbers(nums []int) ([]int, int) {
	evenNumbers := make([]int, len(nums))
	countEven := 0

	for _, value := range nums {
		if value%2 == 0 {
			evenNumbers[countEven] = value
			countEven++
		}
	}

	return evenNumbers[:countEven], countEven
}
