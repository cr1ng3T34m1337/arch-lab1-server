package main

import (
	"fmt"
	"net/http"
	"time"
)

func getTime(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		fmt.Fprintf(w, "Only GET requests permitted")
		return
	}
	current_time := time.Now().Format(time.RFC3339)
	response_json, err := stringify(current_time)
	if err != nil {
		fmt.Printf("JSON parsing error: %v", err)
	} else {
		fmt.Fprint(w, string(response_json))
	}
}
