package main

import (
	"go-coding-dojo/coffeequeue/core"
	"go-coding-dojo/coffeequeue/solution"
)

func main() {
	// Implement your queue and use it instead of DummyQueue inside solution.NewQueue()
	queue := solution.NewQueue()

	go core.RunCoffeeWorker(queue)
	core.StartServer(queue)
}
