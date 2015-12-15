package models

import (
	"time"
)

type Counter struct {
	TimeStamp time.Time
	Qty       int
	Machine
	Column Location
}

type Move struct {
	TimeStamp time.Time
	Qty       int
	From      Location
	To        Location
}

type Fulfill struct {
	TimeStamp time.Time
	Qty       int
	From      Location
	To        Location
}

type Stock struct {
	TimeStamp time.Time
	Qty       int
	Operator
}

func (i *ItemLoc) Move(q int, from Location, to Location) (err error) {
	return err
}
