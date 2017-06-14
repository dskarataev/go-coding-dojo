package solution_test

import (
	"testing"
	"go-coding-dojo/coffeequeue/solution"
	"go-coding-dojo/coffeequeue/core"
)

func TestSliceCoffeeQueue_Get_EmptyQueue(t *testing.T) {
	sliceQueue := solution.SliceCoffeeQueue{}
	order := sliceQueue.Get()

	if order != nil {
		t.Fatalf("Test failed. Order not nil")
	}
}

func TestSliceCoffeeQueue_Get_SingleElement(t *testing.T) {
	sliceQueue := solution.SliceCoffeeQueue{}

	firstOrder := core.Order{
		Name: "t1",
		Customer: "Minh",
		Delay: 10,
	}

	sliceQueue.OrderQueue = []core.IOrder{
		firstOrder,
	}

	order := sliceQueue.Get()

	if order != firstOrder {
		t.Fatalf("Test failed. Order not nil")
	}

	if len(sliceQueue.OrderQueue) > 0 {
		t.Fatalf("Test failed. Queue not empty")
	}
}

func TestSliceCoffeeQueue_Get_MultiElements(t *testing.T) {
	sliceQueue := solution.SliceCoffeeQueue{}

	firstOrder := core.Order{
		Name: "t1",
		Customer: "Minh",
		Delay: 10,
	}

	secondOrder := firstOrder
	secondOrder.Name = "t2"

	sliceQueue.OrderQueue = []core.IOrder{
		firstOrder,
		secondOrder,
	}

	order := sliceQueue.Get()

	if order != firstOrder {
		t.Fatalf("Test failed. Order not nil")
	}
}

func TestSliceCoffeeQueue_Put(t *testing.T) {
	sliceQueue := solution.SliceCoffeeQueue{}
	firstOrder := core.Order{
		Name: "t1",
		Customer: "Minh",
		Delay: 10,
	}

	sliceQueue.Put(firstOrder)

	expectedQueue := solution.SliceCoffeeQueue{}
	expectedQueue.OrderQueue = []core.IOrder{
		firstOrder,
	}

	for i, order := range expectedQueue.OrderQueue {
		if order.GetName() != sliceQueue.OrderQueue[i].GetName() {
			t.Fatalf("Test failed")
		}
		if order.GetDelay() != sliceQueue.OrderQueue[i].GetDelay() {
			t.Fatalf("Test failed")
		}
		if order.GetCustomer() != sliceQueue.OrderQueue[i].GetCustomer() {
			t.Fatalf("Test failed")
		}
	}
}

func TestSliceCoffeeQueue_Put_Two(t *testing.T) {
	sliceQueue := solution.SliceCoffeeQueue{}
	firstOrder := core.Order{
		Name: "t1",
		Customer: "Minh",
		Delay: 10,
	}
	secondOrder := firstOrder
	secondOrder.Name = "t2"

	sliceQueue.Put(firstOrder)
	sliceQueue.Put(secondOrder)

	expectedQueue := solution.SliceCoffeeQueue{}
	expectedQueue.OrderQueue = []core.IOrder{
		firstOrder,
		secondOrder,
	}

	for i, order := range expectedQueue.OrderQueue {
		if order.GetName() != sliceQueue.OrderQueue[i].GetName() {
			t.Fatalf("Test failed")
		}
		if order.GetDelay() != sliceQueue.OrderQueue[i].GetDelay() {
			t.Fatalf("Test failed")
		}
		if order.GetCustomer() != sliceQueue.OrderQueue[i].GetCustomer() {
			t.Fatalf("Test failed")
		}
	}
}

func TestSliceCoffeeQueue_Put_NilOrder(t *testing.T) {
	sliceQueue := solution.SliceCoffeeQueue{}

	nilorder := core.IOrder(nil)

	error := sliceQueue.Put(nilorder)

	if solution.NilError != error {
		t.Fatalf("FAIL")
	}

	if len(sliceQueue.OrderQueue) > 0 {
		t.Fatalf("FAIL")
	}
}

