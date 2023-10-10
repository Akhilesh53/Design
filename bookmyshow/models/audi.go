package models

type IAudi interface {
	GetAudiName() string
	GetSeatStructure() []ISeat
	GetCapacity() int

	SetAudiName(string) IAudi
	SetSeatStructure([]ISeat) IAudi
	SetCapacity(int) IAudi
}

type Audi struct {
	name     string
	seats    []ISeat
	capacity int
	// we can have list of shows here too
}

func (a *Audi) GetAudiName() string {
	return a.name
}

func (a *Audi) GetSeatStructure() []ISeat {
	return a.seats
}

func (a *Audi) GetCapacity() int {
	return a.capacity
}

func (a *Audi) SetAudiName(name string) IAudi {
	a.name = name
	return a
}

func (a *Audi) SetSeatStructure(seats []ISeat) IAudi {
	a.seats = seats
	return a
}

func (a *Audi) SetCapacity(capacity int) IAudi {
	a.capacity = capacity
	return a
}
