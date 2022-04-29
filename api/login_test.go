package api

import (
	"net/http"
	"testing"
)

func TestLogin(t *testing.T) {
	req, err := http.NewRequest("GET", "", nil)
	if err != nil {
		t.Fatalf("Error creating request %s", err.Error())
	}
	var w http.ResponseWriter
	login(w, req)

}
