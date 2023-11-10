package main

import (
	"fmt"
)

/*
	Buat sebuah fungsi yang menghitung jumlah elemen genap dalam array.
	input = [1,2,3,4]
	output = 2
*/

func main() {
	array := []int{1, 2, 3, 4}

	count := 0

	for _, value := range array {
		if value%2 == 0 {
			count++
		}
	}

	fmt.Printf("total jumlah bilangan genap pada array adalah %d", count)
}
