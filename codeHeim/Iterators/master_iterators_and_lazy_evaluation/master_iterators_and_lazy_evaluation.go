package main

import (
	"cmp"
	"fmt"
	"iter"
	"slices"
)

type Order struct {
	OrderID      string
	CustomerName string
	Amount       float64
	Status       string // Status can be "pending", "delivered", or "cancelled"
}

// filter orders and retuern an iterator
// the filter function is generic can use any type of data
func filter[V any](it iter.Seq[V], keep func(V) bool) iter.Seq[V] {
	//the yield will return the next value from the iterator one at a time
	seq := func(yield func(V) bool) {
		for v := range it {
			if keep(v) {
				if !yield(v) {
					break // if yied will return false we break the loop
				}
			}
		}
	}
	return seq //return the seq function
}

// disply prints the orders from the iterator
// The function accept an iterator of type Order
func display(it iter.Seq[Order]) {
	for order := range it {
		fmt.Println("Order ID:", order.OrderID, "Customer Name:", order.CustomerName, "Amount:", order.Amount, "Status:", order.Status)
	}

}

// Play with Iterators and Lazy Evaluation https://tip.golang.org/doc/go1.23#iterators
// Execute go run ./codeHeim/iterators/master_iterators_and_lazy_evaluation/master_iterators_and_lazy_evaluation.go
func main() {
	//Example slice of orders
	orders := []Order{
		{OrderID: "1", CustomerName: "John", Amount: 100.50, Status: "pending"},
		{OrderID: "2", CustomerName: "Doe", Amount: 200, Status: "delivered"},
		{OrderID: "3", CustomerName: "Smith", Amount: 300.75, Status: "cancelled"},
		{OrderID: "4", CustomerName: "Jane", Amount: 400, Status: "pending"},
		{OrderID: "5", CustomerName: "Dane", Amount: 500, Status: "delivered"},
	}

	//Filter orders with amount greater then $300
	highValueOrders := filter(slices.Values(orders), func(order Order) bool {
		return order.Amount > 300
	})
	fmt.Println("Orders with amount > $300:")
	display(highValueOrders)

	//Filter orders with status "delivered"
	deliveredOrders := filter(slices.Values(orders), func(order Order) bool {
		return order.Status == "delivered"
	})
	fmt.Println("Delivered Orders:")
	display(deliveredOrders)

	//Explore Collect() function
	seq := func(yied func(Order) bool) { //The yield function will return the next value from the iterator one at a time to the caller of the function
		for _, order := range orders {
			if order.Amount < 200 {
				if !yied(order) {
					break // if yied will return false we break the loop/iteration
				}
			}
		}
	}
	filteredOreders := slices.Collect(seq)
	fmt.Println("Orders with amount < $200:")
	for _, order := range filteredOreders {
		fmt.Println("Order ID:", order.OrderID, "Customer Name:", order.CustomerName, "Amount:", order.Amount, "Status:", order.Status)
	}

	//Sort orders by amount in descending order using SortedFunc()
	sortFunc := func(a, b Order) int {
		return cmp.Compare(b.Amount, a.Amount) //sort in descending order
	}
	sortedOrders := slices.SortedFunc(slices.Values(orders), sortFunc)
	fmt.Println("Sorted in descending order")
	for _, order := range sortedOrders {
		fmt.Println("Order ID:", order.OrderID, "Customer Name:", order.CustomerName, "Amount:", order.Amount, "Status:", order.Status)
	}

	// Chunk orders into []Order 3 elements at a time
	fmt.Println("\nChunks of orders of size 3:")
	for c := range slices.Chunk(orders, 3) {
		fmt.Println(c)
	}

}
