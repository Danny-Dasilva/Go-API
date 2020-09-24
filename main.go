package main

import (
	"net/http"

	"https://github.com/Danny-Dasilva/Go-Http/api"
)

func main() {
	srv:= api.NewServer()
	http.ListenAndServe(":8080", srv)
}
