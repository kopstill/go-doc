//go:build ignore

package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	s := r.URL.Path[1:]
	if s == "" {
		fmt.Fprintf(w, "Please input something in your path!")
	} else {
		fmt.Fprintf(w, "Hi there, I love %s!", s)
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/view/", viewHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
