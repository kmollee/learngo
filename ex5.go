/*
exmple for maps
*/
package main

import "fmt"

func main() {
	presAge := make(map[string]int)

	presAge["TheodoreRoosevelt"] = 42
	fmt.Println(len(presAge))

	presAge["William"] = 27
	fmt.Println(len(presAge))
}
