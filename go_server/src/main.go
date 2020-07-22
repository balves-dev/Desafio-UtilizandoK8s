package main

// package greeting

// package /.app.go

import (
	"net/http"
)

// func greeting(message string) byte {
// 	fmt.Fprint(w, message)
// }

func handler(w http.ResponseWriter, r *http.Request) {
	// greeting("Code.education Rocks!")
	// greeting.greeting("Code.education Rocks!")

}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8000", nil)
}
