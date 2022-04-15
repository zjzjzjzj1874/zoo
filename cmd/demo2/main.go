package main

import (
	"fmt"
	"log"
	"net/http"
)

type Engine struct{}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		fmt.Fprintf(w, "request path:%s", req.URL.Path)
	case "/hello":
		fmt.Fprintf(w, "hello %s", "world")
	case "/ping":
		fmt.Fprintf(w, "%s", "pong")
	default:
		fmt.Fprintf(w, "404 not found:%s", req.URL.Path)
	}
}

func main() {
	engine := &Engine{}

	fmt.Println("running on 8081......")
	log.Fatal(http.ListenAndServe(":8081", engine))
}
