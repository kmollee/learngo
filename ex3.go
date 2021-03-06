package main

import "fmt"

func main() {
	// about colon equals operator
	// http://stackoverflow.com/questions/16521472/assignment-operator-in-go-language
	i := 1
	// only init and condition
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	for j := 0; j < 5; j++ {
		fmt.Println(j)
	}

	// if condition
	// only execute one, even serverl condition is true
	yourAge := 18
	if yourAge >= 18 {
		fmt.Println("You can drive")
	} else if yourAge >= 20 {
		fmt.Println("You can vote")
	} else {
		fmt.Println("You are too young")
	}

	switch yourAge {
	case 16:
		fmt.Println("Go drive")
	case 18:
		fmt.Println("Go vote")
	default:
		fmt.Println("Go have fun")

	}
}
