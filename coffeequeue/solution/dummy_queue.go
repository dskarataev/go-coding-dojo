package solution

import (
	"go-coding-dojo/coffeequeue/core"
)

type DummyQueue struct{}

func (dq DummyQueue) Put(order core.IOrder) error {
	return nil
}

func (dq DummyQueue) Get() core.IOrder {
	fakeOrder := core.Order{
		Name:     core.OrderNameEspresso,
		Customer: "Dummy Customer",
		Delay:    10,
	}

	return fakeOrder
}
