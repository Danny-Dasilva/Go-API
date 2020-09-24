package main

import (
	"net/http"

	"github.com/Danny-Dasilva/Go-Http/api"
)

func main() {
	srv:= api.NewServer()
	http.ListenAndServe(":8080", srv)
}
