package test

import (
	"go-coding-dojo/coffeequeue/core"
	"go-coding-dojo/coffeequeue/solution"

	"fmt"
	"testing"

	//. "gopkg.in/check.v1"
)

const (
	Count int = 1000
)

// TODO: check on next step
//func (extraTestSuite) TestQueueRaceCondition(c *C) {
//	queue := solution.NewQueue()
//	order := core.NewOrder(core.OrderNameEspresso, "Test Customer")
//
//	go queue.Put(order)
//	go queue.Get()
//}

func BenchmarkQueuePut(b *testing.B) {
	queue := solution.NewQueue()
	order := core.NewOrder(core.OrderNameEspresso, "Bench Customer")

	var err error

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		for i := 0; i < 1000; i++ {
			err = queue.Put(order)

			if err != nil {
				panic(fmt.Sprintf("error in queue: %#v", err))
			}
		}
	}
}

func BenchmarkQueueGet(b *testing.B) {
	queue := solution.NewQueue()
	order := core.NewOrder(core.OrderNameEspresso, "Bench Customer")

	var err error

	// fill the queue before benchmarking
	for j := 0; j < 1000; j++ {
		err = queue.Put(order)

		if err != nil {
			panic(fmt.Sprintf("error in queue: %#v", err))
		}
	}

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		for i := 0; i < 1000; i++ {
			queue.Get()
		}
	}
}
