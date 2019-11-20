package main

import (
	"log"
	"net/http"
)

func hello(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Hello, Go Web!"))
}

func main() {
	http.HandleFunc("/hello", hello)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
