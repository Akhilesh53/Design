package main

import (
	"time"
)

type User struct {
	id       int
	Email    string
	Password string
	Profiles []*Profile
}

type Profile struct {
	id   int
	Name string
	Type string
}

type Netflix struct {
	Videos []*Video
}

type Video struct {
	id          int
	Name        string
	Description string
	Actors      []*Actor
}

type Actor struct {
	id     int
	Name   string
	Videos []*Video
}

// for each user, for each video we need to maintain a status and a last watching timestamp
type UserProfileVideo struct {
	id                    int
	User                  *User
	Profile               *Profile
	Video                 *Video
	LastWatchingTimestamp time.Time
}
