package main

import "fmt"

func main() {
	// array
	var favNum [5]float64
	favNum[0] = 1.
	favNum[1] = 2.
	favNum[2] = 3.
	favNum[3] = 4.
	favNum[4] = 5.

	// array index
	fmt.Println(favNum[3])

	// assign
	favNum2 := [5]float64{6, 7, 8, 9, 10}

	// similar to python's enumerate
	for i, value := range favNum2 {
		fmt.Println(value, i)
	}

	// similar to python's enumerate
	for _, value := range favNum2 {
		fmt.Println(value)
	}
}
