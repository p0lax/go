package main

import (
	"net/http"
	"io"
	"log"
)

func handle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	jsonStr := `[{"Name":"Adam","Age":36,"Job":"CEO"},
		  {"Name":"Eve","Age":34,"Job":"CFO"},
		  {"Name":"Mike","Age":38,"Job":"COO"}]`
	switch r.Method {
	case "GET":
		r.Write([]byte(jsonStr))
		// Serve the resource.
	case "POST":
	// Create a new record.
	case "PUT":
	// Update an existing record.
	case "DELETE":
	// Remove the record.
	default:
		panic("Can't detect API method!")
	}
}

func main() {
	log.Println("Listenng...")
	http.HandleFunc("/", handle)
	http.ListenAndServe(":9000", nil)
}
