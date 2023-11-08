package main

import "fmt"

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

func main() {
	// Implementasikan kode di sini
	var array = [5]int{1, 2, 3, 4, 5}
	jumlahArray := make([]int, len(array))

	jumlahArray[0] = array[0]

	for i := 1; i < len(array); i++ {
		jumlahArray[i] = jumlahArray[i-1] + array[i]
	}

	fmt.Println(jumlahArray)
}
