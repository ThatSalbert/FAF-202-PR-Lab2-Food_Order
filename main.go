package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"food_order/item"

	"github.com/gorilla/mux"
)

func SendToRestaurant() {
	if item.OrderList.IsEmpty() == false && item.RestaurantList.IsEmpty() == false {
		time.Sleep(time.Second * 4)
		order := item.OrderList.Dequeue()
		data, err := json.Marshal(item.GenFS_order_post(*order))
		if err != nil {
			log.Fatal(err)
		}
		response, err := http.Post("http://rest1-dh:8080/v2/order", "application/json", bytes.NewBuffer(data))
		fmt.Println("Order: " + string(data) + " sent to the Restaurant's Dining-Hall.")
		if err != nil {
			fmt.Println("Could not make POST request to the Restaurant's Dining-Hall.")
		}
		defer response.Body.Close()

		time.Sleep(time.Second * 4)
	}
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/order", createPost).Methods("POST")
	router.HandleFunc("/order", getPost).Methods("GET")

	router.HandleFunc("/register", registerRGet).Methods("GET")
	router.HandleFunc("/register", registerRPost).Methods("POST")

	router.HandleFunc("/menu", menuGet).Methods("GET")

	http.ListenAndServe(":7000", router)
}

func menuGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(item.Menu_List)
}

func registerRPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var restaurant item.Register_post
	_ = json.NewDecoder(r.Body).Decode(&restaurant)
	item.RestaurantList.Enqueue(restaurant)
	json.NewEncoder(w).Encode(&restaurant)
	fmt.Println("Restaurant " + strconv.Itoa(restaurant.Restaurant_Id) + " has been registered.")
}

func registerRGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(item.RestaurantList.Elements)
}

func createPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var order item.Cl_order_post
	_ = json.NewDecoder(r.Body).Decode(&order)
	item.OrderList.Enqueue(order)
	json.NewEncoder(w).Encode(&order)

	go SendToRestaurant()
}

func getPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(item.OrderList.Elements)
}
