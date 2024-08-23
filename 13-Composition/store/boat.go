package store

type Boat struct {
	*Product
	Capacity int
	Motorized bool
}

func NewBoat(name string, price float64, capac int, motor bool) *Boat {
	return &Boat{NewProduct(name, "Water", price), capac, motor}
}