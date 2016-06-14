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
	// range example https://gobyexample.com/range
	for i, value := range favNum2 {
		fmt.Println(value, i)
	}

	// similar to python's enumerate
	for _, value := range favNum2 {
		fmt.Println(value)
	}

	// slice
	numSlice := []int{5, 4, 3, 2, 1}
	numSlice2 := numSlice[3:5]

	fmt.Println("numSlice2[0] = ", numSlice2[0])
	numSlice3 := numSlice[:2]

	fmt.Println("numSlice3 = ", numSlice3)
	// make([]int, 0, 10) allocates a slice of length 0 and capacity 10.
	numSlice4 := make([]int, 5, 10)
	fmt.Println("before append", numSlice4)

	/*
		func append(slice []Type, elems ...Type) []Type

		The append built-in function appends elements to the end of a slice. If it has sufficient capacity, the destination is resliced to accommodate the new elements. If it does not, a new underlying array will be allocated. Append returns the updated slice. It is therefore necessary to store the result of append, often in the variable holding the slice itself:
	*/
	numSlice4 = append(numSlice4, 0, -1)
	fmt.Println("after append", numSlice4)

}
