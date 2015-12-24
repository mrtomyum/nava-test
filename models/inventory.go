package models

type Category struct {
	Name   string
	Parent *Category
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

type ItemLoc struct {
	Item
	Location
	Qty int
}
