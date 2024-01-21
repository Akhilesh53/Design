package entities

import "time"

type Ticket struct {
	ticketId    uint16
	parkingSpot *ParkingSpot
	entryTime   time.Time
	vehicle     *Vehicle
	gate        *Gate
}

// auto generated ticket id
func NewTicket() *Ticket {
	return &Ticket{}
}

func (t *Ticket) ForVehicle(vehicle *Vehicle) *Ticket {
	t.vehicle = vehicle
	return t
}

func (t *Ticket) HavingEntryTime(entryTime time.Time) *Ticket {
	t.entryTime = entryTime
	return t
}

func (t *Ticket) WithParkingSpot(parkingSpot *ParkingSpot) *Ticket {
	t.parkingSpot = parkingSpot
	return t
}

func (t *Ticket) FromGate(gate *Gate) *Ticket {
	t.gate = gate
	return t
}

func (t *Ticket) Build() *Ticket {
	return t
}
