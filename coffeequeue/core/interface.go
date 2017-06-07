package core

//go:generate mockery -name IOrder -output mock -outpkg mock
//go:generate mockery -name IQueue -output mock -outpkg mock

type IOrder interface {
	GetName() string
	GetCustomer() string
	GetDelay() int
}

type IQueue interface {
	Put(IOrder) error
	Get() IOrder
}
