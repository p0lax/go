package main

import (
	"encoding/json"
	"net/http"
	"log"
)

func handle(w http.ResponseWriter, r *http.Request) {
	testMap := make(map[string]string)
	switch r.Method {
	case "GET":
		query := r.URL.Query()
		for key := range query {
			testMap[key] = string(query.Get(key))
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
		decoder := json.NewDecoder(r.Body)
		w.Header().Set("Content-Type", "application/json")
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
