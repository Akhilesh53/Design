package models

import "time"

type IShow interface {
}

type Show struct {
	Movie  IMovie
	Time   time.Time
	Cinema ICinema
	Audi   IAudi
	Date time.Weekday
}
