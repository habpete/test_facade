package handlers

import (
	"fmt"
	"net/http"
)

func StartPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}

func InitHandlers() {
	http.HandleFunc("/", StartPage)
}
