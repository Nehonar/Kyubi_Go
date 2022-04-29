package main

import (
	"kyubi_go/api"
)

func main() {
	apiServer := api.NewApi("", 10000)
	apiServer.Run()
}
