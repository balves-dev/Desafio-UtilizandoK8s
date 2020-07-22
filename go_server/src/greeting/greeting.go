package greeting

import (
	"fmt"
	"net/http"
)

func Greeting(w http.ResponseWriter, message string) {
	fmt.Fprint(w, message)
}
