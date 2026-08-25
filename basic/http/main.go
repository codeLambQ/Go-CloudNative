package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hello go http")

	})
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}
}
