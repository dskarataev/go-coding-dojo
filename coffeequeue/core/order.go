package core

const (
	OrderNameEspresso   = "Espresso"
	OrderNameCappuccino = "Cappuccino"
	OrderNameLatte      = "Latte"
)

var (
	MapOrderNameToDelay = map[string]int{
		OrderNameEspresso:   2,
		OrderNameCappuccino: 3,
		OrderNameLatte:      5,
	}
)

type Order struct {
	Name     string `json:"name"`
	Customer string `json:"customer"`
	Delay    int
}

func NewOrder(name string, customer string) IOrder {
	return Order{
		Name:     name,
		Customer: customer,
		Delay:    MapOrderNameToDelay[name],
	}
}

func (o Order) GetName() string {
	return o.Name
}

func (o Order) GetCustomer() string {
	return o.Customer
}

func (o Order) GetDelay() int {
	return o.Delay
}
