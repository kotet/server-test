package main

import (
	"net/http"

	"github.com/gobuffalo/packr"
	"github.com/kotet/server-test/internal/handler"
)

func main() {
	box := packr.NewBox("./assets")
	http.Handle("/", http.FileServer(box))
	http.HandleFunc("/api", handler.Handler)
	println("Serving HTTP on http://localhost:8080 ...")
	http.ListenAndServe(":8080", nil)
}
