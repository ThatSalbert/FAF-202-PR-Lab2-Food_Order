package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"kitchen/item"

	"github.com/gorilla/mux"
)

func CookIt() {
	if !item.OrderList.IsEmpty() {
		order := item.OrderList.Dequeue()
		data, err := json.Marshal(order)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Order: " + string(data) + " cooked.")
		time.Sleep(1 * time.Second)

		fmt.Println("Order Generated: " + string(data))
		response, err := http.Post("http://dining-hall:8080/distribution", "application/json", bytes.NewBuffer(data))
		fmt.Println("Order sent to kitchen.")
		if err != nil {
			fmt.Print("Could not make POST request to the kitchen.")
		}
		defer response.Body.Close()

		time.Sleep(1 * time.Second)
	}
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/order", createPost).Methods("POST")
	router.HandleFunc("/order", getPost).Methods("GET")

	http.ListenAndServe(":8000", router)
}

func createPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var order item.Dh_order_post
	_ = json.NewDecoder(r.Body).Decode(&order)
	item.OrderList.Enqueue(order)
	json.NewEncoder(w).Encode(&order)

	go CookIt()
}

func getPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(item.OrderList.Elements)
}
