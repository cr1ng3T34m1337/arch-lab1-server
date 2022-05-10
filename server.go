package main

import (
	"http",
	"log"
)

func main() {
	http.HandleFunc("/time", getTime)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
