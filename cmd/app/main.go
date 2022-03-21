package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("hello world")

	http.HandleFunc("/", func(rw *http.ResponseWriter, r http.Request) {
		
	})

	err := http.ListenAndServe("127.0.0.1:9000", nil)
	if err != nil {
		log.Fatal("server has been crashed")
	}
}
