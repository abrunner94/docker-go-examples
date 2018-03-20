package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/test", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Hello world!")
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
