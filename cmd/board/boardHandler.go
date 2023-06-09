package main

import (
	"fmt"
	"net/http"

	"github.com/shimiwaka/board-template/schema"
	"github.com/shimiwaka/board-template/connector"
	"github.com/go-chi/chi"
)

func boardHandler(w http.ResponseWriter, r *http.Request) {
	db := connector.ConnectDB()

	board := schema.Board{}
	db.First(&board, "token = ?", chi.URLParam(r, "boardToken"))

	db.Close()

	if board.Owner == "" {
		fmt.Fprintln(w, "{\"error\": \"invalid token\"}")
	} else {
		fmt.Fprintf(w, "{\"owner\": \"%s\"}", board.Owner)
	}
}
