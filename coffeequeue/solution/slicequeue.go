package solution

import (
	"go-coding-dojo/coffeequeue/core"
	"errors"
	"sync"
)

var (
	NilError = errors.New("Sorry guy, we don't accept nil order")
)

type SliceCoffeeQueue struct {
	sync.RWMutex
	OrderQueue []core.IOrder
}

func (sq *SliceCoffeeQueue) Put(order core.IOrder) error {
	// Prevent nil order
	if nil == order {
		return NilError
	}
	sq.Lock()
	defer sq.Unlock()
	sq.OrderQueue = append(sq.OrderQueue, order)
	return nil
}

func (sq *SliceCoffeeQueue) Get() core.IOrder {

	sq.Lock()
	defer sq.Unlock()
	lenq := len(sq.OrderQueue)

	if lenq > 0 {
		order := sq.OrderQueue[0]
		if lenq > 1 {
			sq.OrderQueue = sq.OrderQueue[1:]
		} else {
			sq.OrderQueue = sq.OrderQueue[0:0]
		}
		return order
	}
	return nil
}

