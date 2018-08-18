package main

import (
	"log"
	"net/http"

	. "github.com/shkatovdm/tlgstat"
)

func main() {
    router := NewRouter()
    log.Fatal(http.ListenAndServe(":8181", router))
}