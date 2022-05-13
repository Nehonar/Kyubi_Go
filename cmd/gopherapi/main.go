package main

import (
	"kyubi_go/pkg/server"
	"log"
	"net/http"
	//"github.com/friendsofgo/gopher-api/pkg/server"
)

func main() {
	s := server.New()
	log.Fatal(http.ListenAndServe(":8080", s.Router()))
}
