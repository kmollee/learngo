package main

// about fmt https://golang.org/pkg/fmt/
import "fmt"

func main() {

	const pi float64 = 3.14159265

	// var (
	// 	varA = 2
	// 	varB = 3
	// )

	var myName string = "William"

	fmt.Println(myName + " is is a robot")

	var isOver40 bool = true

	// format print out
	fmt.Printf("%.3f \n", pi)

	// display the data type
	fmt.Printf("%T \n", isOver40)

}
