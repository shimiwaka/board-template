package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	// "github.com/jinzhu/gorm"
	"github.com/shimiwaka/board-template/schema"
	"github.com/shimiwaka/board-template/connection"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func createHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "CREATE!")
	fmt.Fprintln(w, chi.URLParam(r, "email"))

	db := connection.connectDB()

	board := schema.Board{Owner: chi.URLParam(r, "email")}
	db.Create(&board)

	db.Close()
}
