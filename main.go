package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Message struct {
	Text string `json:"text"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	msg := Message{Text: "Hello from DevOps Go App 🚀"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(msg)
}

func main() {
	http.HandleFunc("/", handler)
	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
