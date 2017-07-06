package solution

import (
	"go-coding-dojo/coffeequeue/core"
	"time"
	"errors"
)

type ChannelQueue struct{
	channel chan core.IOrder
}

func (dq ChannelQueue) Put(order core.IOrder) error {
	// Prevent nil order
	if nil == order {
		return NilError
	}

	select {
	case dq.channel <- order:
		return nil
	case <-time.After(10000 * time.Millisecond):
		return errors.New("Queue is full")
	default:

	}

	return nil
}

func (dq ChannelQueue) Get() core.IOrder {
	select {
	case ret := <-dq.channel:
		return ret
	case <-time.After(1000 * time.Millisecond):
		return nil
	}
	return nil
}
