package models_test

import (
	//	"fmt"
	"github.com/mrtomyum/nava-test/models"
	//	"log"
	//	"net/http"
	"fmt"
	"testing"
)

type Env struct {
	db models.DataStore
}

func Test_InitDB(t *testing.T) {
	db, err := models.NewDB("postgres://postgres:postgres@itu.nopadol.com/npdb")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("start NewDB")
	e := &Env{db}
	e.itemIndex()
}

func (e *Env) itemIndex() {
	items, _ := e.db.AllItem()
	for _, i := range items {
		fmt.Printf("ID: %d ,Name: %s,Price: %d", i.ID, i.Name, i.SellPrice)
	}
}
