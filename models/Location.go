package models

//type LocType int
//
//const (
//	Site LocType = iota
//	SubSite
//	Vehicle
//	VMC
//	Store
//)
//
//var LocTypes = [...]string{
//	"Site",
//	"Subsite",
//	"Vehicle",
//	"VMC",
//	"Store",
//}

// Location is parent-child hierarchy data each contain LocType which can be stocking or non stocking location
type Location struct {
	ID int
	Locator
	Parent     *Location
	IsStocking bool
}

type Locator interface {
	Loc()
}

type Machine struct {
	ID      int
	Code    string
	Columns []Location
	MacType string
	Location
}

func (m *Machine) Loc() Location {
	return m.Location
}

type Vehicle struct {
	ID     int
	Code   string
	Driver Person
	Status string
	Location
}

func (v *Vehicle) Loc() Location {
	return v.Location
}

type Store struct {
	Name string
	Location
}

func (s *Store) Loc() Location {
	return s.Location
}
