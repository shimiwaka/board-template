package main

import (
	"fmt"
	"net/http"

	// "github.com/go-chi/chi"
	// "github.com/jinzhu/gorm"
	"github.com/shimiwaka/board-template/schema"
	"github.com/shimiwaka/board-template/connector"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func createHandler(w http.ResponseWriter, r *http.Request) {
	e := r.ParseForm()
	if e != nil {
		panic("error: parse error")
	}

	fmt.Fprintln(w, "CREATE!")
	fmt.Fprintln(w, r.Form.Get("email"))

	db := connector.ConnectDB()

	board := schema.Board{Owner: r.Form.Get("email")}
	db.Create(&board)

	db.Close()
}
