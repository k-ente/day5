package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	switch r.Header.Get("Content_Type") {
	case "application/json":
		var data map[string]interface{}
		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	case "application/x-www-form-urlencoded":
		if err := r.ParseForm(); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return

		}

	case "text/plain":
		var body []byte
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return

		}
		fmt.Print(body)

	}
	fmt.Fprintf(w, "Hello World")
}

func main() {
	http.HandleFunc("/api", handler)
	http.ListenAndServe(":8080", nil)
}

///test
