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
}
