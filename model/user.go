package model

import "time"

type User struct {
	Id            int
	Name          string
	Presence      time.Time
	Absence       time.Time
	Late_presence bool
}
