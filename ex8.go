/*
this example is about lib's usage method
*/
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {
	// lib string
	sampString := "Hello World"
	fmt.Println(strings.Contains(sampString, "lo"))
	fmt.Println(strings.Index(sampString, "lo"))
	fmt.Println(strings.Count(sampString, "lo"))
	fmt.Println(strings.Replace(sampString, "l", "L", 3))
	csvString := "1,2,3,4,5,6"
	fmt.Println(strings.Split(csvString, ","))
	listOfLetter := []string{"b", "c", "a"}
	// use sort, sort string array
	sort.Strings(listOfLetter)
	fmt.Println("Letters:", listOfLetter)
	listOfNums := strings.Join([]string{"3", "2", "1"}, ",")
	fmt.Println("list Of numbers:", listOfNums)

	// write file
	file, err := os.Create("samp.txt")

	if err != nil {
		log.Fatal(err)
	}
	file.WriteString("This is random txt")
	file.Close()

	// read file
	stream, err := ioutil.ReadFile("samp.txt")
	if err != nil {
		log.Fatal(err)
	}
	readString := string(stream)
	fmt.Println("content of samp.txt", readString)

}
