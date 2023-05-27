package main

import (
	"fmt"
	"net/http"
)

func createHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "CREATE!")
}
