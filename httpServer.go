package main

import (
	"io/ioutil"
	"encoding/json"
	"net/http"
	"log"
	"fmt"
)

type SuccessResponse struct  {
	code int32
	message string
	status bool
}

type Request struct {
	isSuccess bool
}

func handle(w http.ResponseWriter, r *http.Request) {
	paramsMap := make(map[string]string)
	switch r.Method {
	case "GET":
		query := r.URL.Query()
		for key := range query {
			paramsMap[key] = string(query.Get(key))
		}
		w.Header().Set("Content-Type", "application/json")
		data, err := json.Marshal(paramsMap)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(data)
	case "POST":
		data, err := ioutil.ReadAll(r.Body);
		if err != nil {
    		fmt.Println(err)
    	}
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	case "PUT":
		var data Request
		if err := json.Unmarshal(r.Body, &data); err != nil {
			panic(err);
		}
		fmt.Println(data)
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
