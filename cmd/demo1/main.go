package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/ping", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte("pong\n"))
	})
	http.HandleFunc("/hello", helloHandle)
	http.HandleFunc("/", echoHandle)

	fmt.Println("running on 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func echoHandle(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "path:%s", request.URL.Path)
}

func helloHandle(writer http.ResponseWriter, request *http.Request) {
	for k, v := range request.Header {
		fmt.Fprintf(writer, "Header[%q] = %q\n", k, v)
	}
}
