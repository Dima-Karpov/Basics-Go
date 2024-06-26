package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type testMessage struct {
	Message string `json:"message"`
}

func displayTest(w http.ResponseWriter, r *http.Request) {
	data := testMessage{"A test message"}
	b, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(b))
}

func main() {
	http.HandleFunc("/api/v1/test", displayTest)
	http.ListenAndServe(":8888", nil)
}
