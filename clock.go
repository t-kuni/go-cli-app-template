package main

import "time"

type Clocker interface {
	Now() time.Time
}

type Clock struct{}

func (clock Clock) Now() time.Time {
	return time.Now()
}
