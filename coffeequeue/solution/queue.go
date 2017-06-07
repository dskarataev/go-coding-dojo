package solution

import (
	"go-coding-dojo/coffeequeue/core"
)

//type LzdQueue struct{
//	CreatedAt time.Time
//	CoffeeName string
//	CustomerName string
//	Next *LzdQueue
//}
//
//
//
//type IntHeap []LzdQueue
//
//func (h IntHeap) Len() int           { return len(h) }
//func (h IntHeap) Less(i, j int) bool { return h[i].CreatedAt < h[j].CreatedAt }
//func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
//
//
//func (dq LzdQueue) Put(order core.IOrder) error {
//	return nil
//}
//
//func (dq LzdQueue) Get() core.IOrder {
//	fakeOrder := core.Order{
//		Name:     core.OrderNameEspresso,
//		Customer: "Dummy Customer",
//		Delay:    10,
//	}
//
//	return fakeOrder
//}
//
//func New() core.IQueue {
//	return IntHeap{}
//
//}

type Pointer struct {
	Order core.IOrder
	Ptr   *Pointer
}

//var firstOrder *Pointer
//var lastOrder *Pointer


type CoffeeQueue struct {
	FirstOrder *Pointer
	LastOrder *Pointer
}

func (q *CoffeeQueue) Put(order core.IOrder) error {
	newEl := &Pointer{
		Order: order,
	}

	if q.FirstOrder == nil || q.LastOrder == nil {
		q.FirstOrder = newEl
		q.LastOrder = newEl
		return nil
	}

	q.LastOrder.Ptr = newEl
	q.LastOrder = newEl
	return nil

	//
	//if q.LastOrder == nil {
	//	newEl.Ptr = q.LastOrder
	//}
	//q.LastOrder = newEl
	//
	//if q.FirstOrder == nil {
	//	q.FirstOrder = q.LastOrder
	//}


	//if q.LastOrder != nil {
	//	q.LastOrder.Ptr = &newEl
	//}

	//return nil
}

func (q *CoffeeQueue) Get() core.IOrder {
	if q.FirstOrder == nil {
		return nil
	}
	returnOrder := q.FirstOrder.Order
	q.FirstOrder = q.FirstOrder.Ptr
	return returnOrder
}

