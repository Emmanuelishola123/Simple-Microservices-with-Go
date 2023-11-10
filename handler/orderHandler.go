package handler

import (
	"fmt"
	"net/http"
)

type Order struct{}

func (o *Order) getAllOrders(w http.ResponseWriter, r *http.Request) {

	fmt.Println("List all orders")
}

func (o *Order) postNewOrder(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Post order")
}

func (o *Order) GetOrderById(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Get by ID")
}

func (o *Order) UpdateOrderById(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Update")
}

func (o *Order) DeleteOrderById(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Delete")
}
