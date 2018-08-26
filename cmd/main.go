package main

import (
	"log"
	"net/http"
	"github.com/shkatovdm/models"

	. "github.com/shkatovdm/tlgstat"
)

func main() {
	db, err := models.Init("postgres://user:pass@localhost/bookstore")
	if err != nil {
		log.Panic(err)
	}
	env := &Env{DB: db}
    router := NewRouter()
    log.Fatal(http.ListenAndServe(":8181", router))
}