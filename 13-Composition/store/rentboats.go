package store

type Crew struct {
	Captain, FirstOfficer string
}

type RentalBoat struct {
	*Boat
	*Crew
}

func NewRentalBoat(name string, price float64, capacity int, 
					motor, crewed bool, captain, firstOfficer string) *RentalBoat {
	return &RentalBoat{NewBoat(name, price, capacity, motor), &Crew{captain, firstOfficer}}
}