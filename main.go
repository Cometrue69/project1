package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

var task string

type requestBody struct {
	Task string `json:"task"`
}

func SetTaskHandler(w http.ResponseWriter, r *http.Request) {
	var req requestBody
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	task = req.Task
	w.WriteHeader(http.StatusOK)
}

func GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello, %s", task)
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/task", SetTaskHandler).Methods("POST")
	router.HandleFunc("/api/task", GetTaskHandler).Methods("GET")

	http.ListenAndServe(":8080", router)
}
