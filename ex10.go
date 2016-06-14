/*
about use http, build simple server
https://golang.org/pkg/net/http/
*/

package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Earth\n")

}

func handler2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Earth2\n")

}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/earth", handler2)
	http.ListenAndServe(":8080", nil)
}
