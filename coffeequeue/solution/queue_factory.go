package solution

import (
	"go-coding-dojo/coffeequeue/core"
)

func NewQueue() core.IQueue {
	// TODO: change DummyQueue here to your queue implementation
	//return &CoffeeQueue{}
	return &ChannelQueue{
		channel: make(chan core.IOrder, 2000),
	}
}
