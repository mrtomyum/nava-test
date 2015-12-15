package models

type Item struct {
	ID       int
	Name     int
	BaseUnit Unit // BaseUnit ratio is always = 1
	StdPrice float32
}

type Unit int

const (
	ea Unit = iota
	pack
	carton
	palate
)

//var UnitTH = [...]string{
//	"ชิ้น",
//	"แพ็ค",
//	"กล่อง",
//	"หีบ",
//	"พาเลท",
//}

type ItemUnit struct {
	Item
	Unit
	Ratio int
}

type Category struct {
	Name   string
	Parent *Category
}

type ItemLoc struct {
	Item
	Location
	Qty int
}
