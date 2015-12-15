package models
import (
	"errors"
)

type LocType int

const (
	Site LocType = iota
	SubSite
	Vehicle
	VMC
	Store
)

var LocTypes = [...]string{
	"Site",
	"Subsite",
	"Vehicle",
	"VMC",
	"Store",
}

// Location is parent-child hierarchy data each contain LocType which can be stocking or non stocking location
type Loc struct {
	ID int
	LocType
	Parent  *Loc
	IsStocking bool
}

func (l Loc) New(lt LocType, pr *Loc) error{
	if lt && pr == nil {
		return errors.New("โปรดระบุประเภทที่ Location และ Parent Loc")
	}
	l.LocType = lt
	l.Parent = pr
	return nil
}

type Machine struct {
	ID      int
	Code    string
	Columns []Loc
	Loc
	MacType string
}

type Vehicle struct {
	ID int
	Code string
	Driver Person
	Status string
}