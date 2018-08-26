package main

import (
	"log"
	"net/http"
	"github.com/shkatovdm/models"
	. "github.com/shkatovdm/tlgstat"
	"context"
)

type ContextInjector struct {
	ctx context.Context
	h   http.Handler
}

func main() {
	db, err := models.Init("postgres://tlgstat:tlgstat@localhost/tlgstat")
	if err != nil {
		log.Panic(err)
	}
	context.WithValue(context.Background(), "db", db)
    router := NewRouter()
    log.Fatal(http.ListenAndServe(":8181", router))
}