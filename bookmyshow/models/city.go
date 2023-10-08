package models

type ICity interface {
	GetCityName() string
	GetListOfCinemaHalls() []ICinema

	SetCityName(cityName string) ICity
	SetListOfCinemaHalls([]ICinema) ICity
}

type DelhiCity struct {
	name    string
	cinemas []ICinema
}

func (c *DelhiCity) GetCityName() string {
	return c.name
}

func (c *DelhiCity) GetListOfCinemaHalls() []ICinema {
	return c.cinemas
}

func (c *DelhiCity) SetCityName(name string) ICity {
	c.name = name
	return c
}

func (c *DelhiCity) SetListOfCinemaHalls(cinemas []ICinema) ICity {
	c.cinemas = cinemas
	return c
}
