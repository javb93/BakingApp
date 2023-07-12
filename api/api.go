package api

import (
	"fmt"
	"net/http"
)

func HandleGet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `{"message": "Hello, world!"}`)
}
