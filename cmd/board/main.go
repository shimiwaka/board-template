package main

import (
	"net/http"
	// "net/http/cgi"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

func main() {
	rootPath := os.Getenv("SCRIPT_NAME")

    r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"}}))

    r.Get(rootPath + "/", pingHandler)
	r.Post(rootPath + "/create", createHandler)
	r.Post(rootPath + "/forget", forgetHandler)

	r.Route(rootPath + "/board", func(r chi.Router) {
		r.Get("/{boardToken}", boardHandler)
	  })

    http.ListenAndServe(":9999", r)
	// cgi.Serve(r)
}
