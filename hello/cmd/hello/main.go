package main

import (
	"fmt"
	"net/http"
	"hello/cmd/internal"
)

func main() {

	handler := internal.NewHandler()
	http.ListenAndServe(":8080", handler)
}