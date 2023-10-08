package models

type IAudi interface {
	GetAudiName() string
	GetSeatStructure() []ISeat
	GetCapacity() int

	SetAudiName(string) IAudi
	SetSeatStructure([]ISeat) IAudi
	SetCapacity(int) IAudi
}

type Audi1 struct {
	name     string
	seats    []ISeat
	capacity int
}

func (a *Audi1) GetAudiName() string {
	return a.name
}

func (a *Audi1) GetSeatStructure() []ISeat {
	return a.seats
}

func (a *Audi1) GetCapacity() int {
	return a.capacity
}

func (a *Audi1) SetAudiName(name string) IAudi {
	a.name = name
	return a
}

func (a *Audi1) SetSeatStructure(seats []ISeat) IAudi {
	a.seats = seats
	return a
}

func (a *Audi1) SetCapacity(capacity int) IAudi {
	a.capacity = capacity
	return a
}
