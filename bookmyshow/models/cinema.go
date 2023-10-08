package models

type ICinema interface {
	GetCinemaName() string
	GetCinemaAddress() string
	GetAuditoriumList() []IAudi
	GetCinemaCity() ICity

	SetCinemaName(string) ICinema
	SetCinemaAddress(string) ICinema
	SetAuditoriumList([]IAudi) ICinema
	SetCinemaCity(ICity) ICinema
}

type PVRCinema struct {
	cinemaName     string
	cinemaAddress  string
	city           ICity
	auditoriumList []IAudi
}

func (c *PVRCinema) GetCinemaName() string {
	return c.cinemaName
}

func (c *PVRCinema) GetCinemaAddress() string {
	return c.cinemaAddress
}

func (c *PVRCinema) GetAuditoriumList() []IAudi {
	return c.auditoriumList
}

func (c *PVRCinema) GetCinemaCity() ICity {
	return c.city
}

func (c *PVRCinema) SetCinemaName(name string) ICinema {
	c.cinemaName = name
	return c
}

func (c *PVRCinema) SetCinemaAddress(address string) ICinema {
	c.cinemaAddress = address
	return c
}

func (c *PVRCinema) SetAuditoriumList(audis []IAudi) ICinema {
	c.auditoriumList = audis
	return c
}

func (c *PVRCinema) SetCinemaCity(city ICity) ICinema {
	c.city = city
	return c
}
