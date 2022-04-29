package api

import (
	"fmt"
	"log"
	"net/http"
)

type Api struct {
	host   string
	port   int
	logger *log.Logger
}

func NewApi(host string, port int) *Api {
	http.HandleFunc("/ping", ping)
	http.HandleFunc("/login", login)
	return &Api{
		host: host,
		port: port,
	}
}

func (api *Api) Run() error {
	return http.ListenAndServe(fmt.Sprintf("%s:%d", api.host, api.port), nil)
}
