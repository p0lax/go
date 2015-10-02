package main

import (
	"encoding/json"
	"net/http"
	"log"
)

func handle(w http.ResponseWriter, r *http.Request) {
	var testMap [string]string
	switch r.Method {
	case "GET":
		query := r.URL.Query()
		for key := range query {
			log.Println(key, query.Get(key));
			testMap[key] = query.Get(key)
		}
		w.Header().Set("Content-Type", "application/json")
		js, err := json.Marshal(testMap)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(js)
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
