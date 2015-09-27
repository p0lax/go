package main

import (
	"encoding/json"
	"net/http"
	"log"
)

type Test struct{
	a string
}

func handle(w http.ResponseWriter, r *http.Request) {
	test := Test{a:"1111"}
	js, err := json.Marshal(test)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	switch r.Method {
	case "GET":
		w.Header().Set("Content-Type", "application/json")
		r.Write(js)
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
