package main

func main() {
	array := []int{4, 3, 5, 8, 1}

	minValue := array[0]

	for i, value := range array {
		if i < len(array) && value < minValue {
			minValue = value
		} else if len(array) == 0 {
			minValue = 0
		}
	}

	println(minValue)

}
