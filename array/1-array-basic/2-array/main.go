package main

import "fmt"

func main() {
	/*
		buatlah sebuah fungsi yang membalikan total dari semua elemen array
		contoh array = [4]int{1,2,3,4,5}
		output total = 15
	*/

	var array = [5]int{1, 2, 3, 4, 5}

	total := 0

	for i := 0; i < len(array); i++ {
		total += array[i]
	}

	fmt.Print("Total = ", total)
}
