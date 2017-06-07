package test

import (
	"go-coding-dojo/coffeequeue/core"
	"go-coding-dojo/coffeequeue/solution"

	"fmt"
	"testing"
)

func BenchmarkQueue(b *testing.B) {
	queue := solution.NewQueue()
	order := core.NewOrder(core.OrderNameEspresso, "Bench Customer")

	var err error

	for n := 0; n < b.N; n++ {
		err = queue.Put(order)

		if err != nil {
			panic(fmt.Sprintf("error in queue: %#v", err))
		}
	}

	for n := 0; n < b.N; n++ {
		queue.Get()
	}
}
