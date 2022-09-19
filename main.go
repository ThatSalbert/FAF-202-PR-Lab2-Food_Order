package main

import (
	"encoding/json"
	"log"
	"net/http"

	"kitchen/item"

	"github.com/gorilla/mux"
)

var order_list []item.Dh_order_post

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/order", createPost).Methods("POST")
	router.HandleFunc("/order", getPost).Methods("GET")

	log.Println("Listening...")
	http.ListenAndServe(":8000", router)
}

func createPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var order item.Dh_order_post
	_ = json.NewDecoder(r.Body).Decode(&order)
	order_list = append(order_list, order)
	json.NewEncoder(w).Encode(&order)
}

func getPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(order_list)
}
