package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/time", getTime)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
