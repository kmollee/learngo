/*
examples for function
*/
package main

import "fmt"

// sum up all element of numbers
func addThemUp(numbers []float64) float64 {
	sum := 0.0

	for _, val := range numbers {
		sum += val
	}
	return sum
}

// return next 2 value of number
func next2Values(number int) (int, int) {
	return number + 1, number + 2
}

// undefine number of function
// subtract
func subtractThem(args ...int) int {
	finalValue := 0
	for _, val := range args {
		finalValue -= val
	}
	return finalValue
}

/*
recursive function about factorial
*/
func factorial(num int) int {
	if num == 0 {
		return 1
	}
	return num * factorial(num-1)
}

func main() {
	listNums := []float64{1, 2, 3, 4, 5}
	fmt.Println("Sum :", addThemUp(listNums))

	num1, num2 := next2Values(5)
	fmt.Println(num1, num2)

	fmt.Println(subtractThem(1, 2, 3, 4, 5))

	fmt.Println(factorial(3))
}
