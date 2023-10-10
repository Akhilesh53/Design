package models

type IActor interface {
	GetActorName() string
	SetActorName(string) IActor
}

type Cast struct {
	Name string
	// further details of the actor
}

func (c *Cast) GetActorName() string {
	return c.Name
}

func (c *Cast) SetActorName(name string) IActor {
	c.Name = name
	return c
}
