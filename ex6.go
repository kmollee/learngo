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

/*
about defer function
*/
func printOne() { fmt.Println(1) }
func printTwo() { fmt.Println(2) }

func safeDiv(num1 int, num2 int) int {
	defer func() {
		// recover catch the error msg
		fmt.Println(recover())
	}()
	solution := num1 / num2
	return solution
}

/*
pointer apply
function take a pointer, as parameter
*/
func changeXVal(xPrt *int) {
	*xPrt = 2
}

func main() {
	listNums := []float64{1, 2, 3, 4, 5}
	fmt.Println("Sum :", addThemUp(listNums))

	num1, num2 := next2Values(5)
	fmt.Println(num1, num2)

	fmt.Println(subtractThem(1, 2, 3, 4, 5))

	fmt.Println(factorial(3))
	// defer function will wait the main
	// function execute finish
	// then execute (like stack, LIFO)
	defer printOne()
	defer printTwo()
	printOne()

	fmt.Println(safeDiv(3, 0))

	// pointer examine
	x := 0
	fmt.Println("Now x =", x)
	// there should pass a refer, as pointer
	changeXVal(&x)
	fmt.Println("after change x, x =", x)
	fmt.Println("the address of x is", &x)

	/*
		func new(Type) *Type
		The new built-in function allocates memory. The first argument is a type, not a value,
		and the value returned is a pointer to a newly allocated zero value of that type.
	*/
	yPrt := new(int)
	changeXVal(yPrt)
	fmt.Println("now yPrt point to ", *yPrt)
}
