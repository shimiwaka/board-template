package main

import (
	"net/http"
	// "net/http/cgi"
	"os"

	"github.com/go-chi/chi"
)

func main() {
	rootPath := os.Getenv("SCRIPT_NAME")

    r := chi.NewRouter()
    r.Get(rootPath + "/", pingHandler)
	r.Post(rootPath + "/create", createHandler)

    http.ListenAndServe(":9999", r)
	// cgi.Serve(r)
}
