package main

import (
	"log"
	"net/http"
	"test_facade/handlers"
)

func main() {
	handlers.InitHandlers()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
