package main

import (
	"greeting/greeting"
	"net/http"
	"fmt"
)

func handler(w http.ResponseWriter, r *http.Request) { 
	fmt.Fprint(w, greeting.Greeting("Code.education Rocks!"))

}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8000", nil)
}
