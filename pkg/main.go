package main

import (
	"kyubi_go/cmd/server"
	"log"
	"net/http"
)

func main() {

	s := server.New()
	log.Fatal(http.ListenAndServe(":8000", s.Router()))
}
