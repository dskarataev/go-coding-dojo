package solution

import (
	"go-coding-dojo/coffeequeue/core"
	"errors"
)

var (
	NilError = errors.New("Sorry guy, we don't accept nil order")
)

type SliceCoffeeQueue struct {
	OrderQueue []core.IOrder
}

func (sq *SliceCoffeeQueue) Put(order core.IOrder) error {
	// Prevent nil order
	if nil == order {
		return NilError
	}
	sq.OrderQueue = append(sq.OrderQueue, order)
	return nil
}

func (sq *SliceCoffeeQueue) Get() core.IOrder {
	if len(sq.OrderQueue) > 0 {
		order := sq.OrderQueue[0]

		if len(sq.OrderQueue) > 1 {
			sq.OrderQueue = sq.OrderQueue[1:]
		} else {
			sq.OrderQueue = make([]core.IOrder, 0, 100)
		}

		return order
	}
	return nil
}

