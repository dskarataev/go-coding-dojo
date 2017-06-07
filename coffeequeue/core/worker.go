package core

import (
	"fmt"
	"time"
)

func RunCoffeeWorker(queue IQueue) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("EMERGENCY PANIC RECOVERED IN COFFEE MACHINE: %#v\n", r)
		}
	}()

	for {
		// get the next order from the queue
		order := queue.Get()

		if order == nil {
			continue
		}

		// wait for preparing coffee
		time.Sleep(time.Second * time.Duration(order.GetDelay()))

		// report about order is done
		fmt.Printf("%s, your %s is ready!\n", order.GetCustomer(), order.GetName())
	}
}
