package main

//import (
//	"github.com/mrtomyum/nava-test/models"
//	"log"
//	"net/http"
//	"fmt"
//)
//
//type Env struct {
//	db models.DataStore
//}
//
//func main() {
//	db, err := models.NewDB("postgres://postgres:postgres@itu.nopadol.com/npdb")
//	if err != nil {
//		log.Panic(err)
//	}
//
//	env := &Env{db}
//
//	http.HandleFunc("/items", env.itemIndex)
//	http.ListenAndServe(":8080", nil)
//}
//
//func (env *Env) itemIndex(w http.ResponseWriter, r *http.Request) {
//	items, _ := env.db.AllItem()
//	for _, i := range items {
//		fmt.Fprintf(w,"ID: %d ,Name: %s,Price: %d", i.ID, i.Name, i.StdPrice )
//	}
//}
import (
	"github.com/mrtomyum/nava-test/models"
	//	"testing"
	"fmt"
	"log"
	"net/http"
)

type Env struct {
	db models.DataStore
}

func main() {
	db, err := models.NewDB("postgres://postgres:postgres@itu.nopadol.com/npdb")
	if err != nil {
		log.Println(err)
	}
	fmt.Println("start NewDB")
	e := &Env{db}
	http.HandleFunc("/items", e.itemIndex)
	http.HandleFunc("/locations", e.locationIndex)
	http.ListenAndServe(":8080", nil)
}

func (e *Env) itemIndex(w http.ResponseWriter, r *http.Request) {
	items, err := e.db.ImportItem()
	if err != nil {
		log.Fatal(err)
	}
	//	fmt.Println("Start For Loop.", items)
	for _, i := range items {
		fmt.Fprintf(w, "ID:%4d | NameEN: %-30s | NameTH: %-50s| UnitStock: %-10s | Price: %3v\n", i.ID, i.NameEN, i.NameTH, i.UnitStock, i.SellPrice)
	}
}

func (e *Env) locationIndex(w http.ResponseWriter, r *http.Request) {
	locs, err := e.db.ImportLocation()
	if err != nil {
		log.Fatal(err)
	}
	for _, l := range locs {
		fmt.Fprintf(w, "ID:%4d\n", l.ID)
	}
}
