package main

import (
	"fmt"
	"net/http"
)

func boardHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World!")
}
