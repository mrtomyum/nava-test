package models_test

import (
	m "github.com/mrtomyum/nava-test/models"
	"testing"
)

func TestNewMachine(t *testing.T) {
	if mc := new(m.Machine); mc == nil {
		t.Error("Expect instant new Machine but no new object ")
	}
}

func TestNewLoc(t *testing.T) {
	l := new(m.Location)
	if l == nil {
		t.Error("Expect new Site instant")
	}
}

func TestNewItem(t *testing.T) {
	if i := new(m.Item); i == nil {
		t.Error("Expect new instant of Item but nil")
	}
}

//func TestItemReceive(t *testing.T) {
//	i := new(m.Item)
//	items := []m.Item{
//		{
//			i.ID:       1,
//			i.Name:     "Coke can 325ml.",
//			i.BaseUnit: 1,
//			i.StdPrice: 15,
//		},
//		{
//			i.ID:       2,
//			i.Name:     "Mirinda Red can 325ml",
//			i.BaseUnit: 1,
//			i.StdPrice: 15,
//		},
//	}
//	for a := range items {
//		fmt.Printf("ID: %d\n" ,a.ID)
//		fmt.Printf("name: %s\n" ,a.Name)
//		fmt.Printf("price: %d\n" ,a.StdPrice)
//		fmt.Printf("Unit: %d\n" ,a.BaseUnit)
//	}
//}
