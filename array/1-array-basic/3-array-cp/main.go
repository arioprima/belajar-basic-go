package main

/*
	Array 1 Dimensi adalah kumpulan data yang memiliki tipe data yang sama
	Pada kali ini kita akan menjumpai array 1 dimensi
	contoh:
	input = [1,2,3,4]
	output = [1,3,6,10]

	penjelasan:
	1 + 2 = 3
	3 + 3 = 6
	6 + 4 = 10
*/

func RunningSum(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	runningSum := make([]int, len(nums))

	runningSum[0] = nums[0]

	for i, value := range nums[1:] {
		runningSum[i+1] = runningSum[i] + value
	}

	return runningSum
}
