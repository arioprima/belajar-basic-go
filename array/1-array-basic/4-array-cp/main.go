package main

func MinArrayCp(nums []int) (int, int) {
	/*
		Deskripsi: Buat sebuah fungsi yang mengembalikan indeks dari elemen terkecil dalam array.
		Contoh Input: [5, 3, 1, 8, 4]
		Contoh Output: Nilai terkecil berada pada index ke-2 dengan nilai 1
	*/

	if len(nums) == 0 {
		return 0, 0
	}

	minArray := nums[0]
	minIndex := 0

	for i, value := range nums {
		if i < len(nums) && value < minArray {
			minArray = value
			minIndex = i
		}
	}

	return minIndex, minArray
}
