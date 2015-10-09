package main

import (
	"io/ioutil"
	"encoding/json"
	"net/http"
	"log"
	"fmt"
)

func handle(w http.ResponseWriter, r *http.Request) {
	paramsMap := make(map[string]string)
	switch r.Method {
	case "GET":
		query := r.URL.Query()
		for key := range query {
			paramsMap[key] = string(query.Get(key))
		}
		w.Header().Set("Content-Type", "application/json")
		js, err := json.Marshal(paramsMap)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(js)
	case "POST":
		hah, err := ioutil.ReadAll(r.Body);
		if err != nil {
    		fmt.Println(err)
    	}
		w.Header().Set("Content-Type", "application/json")
		w.Write(hah)
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
