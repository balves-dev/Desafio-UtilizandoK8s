package main

import (
	"greeting/greeting"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	greeting.Greeting(w, "<b> Code.education Rocks! </b>")
	// greeting.greeting("Code.education Rocks!")

}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8000", nil)
}
