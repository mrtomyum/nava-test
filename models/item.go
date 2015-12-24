package models

import (
	"database/sql"
	"github.com/shopspring/decimal"
	"log"
	"strings"
)

type Item struct {
	ID        int
	Code      string
	NameTH    string
	NameEN    string
	SellPrice decimal.Decimal
	UnitStock string
	UnitBuy   string
	UnitSell  string
	TypeID    int
	BrandID   int
	//	BaseUnit Unit // BaseUnit ratio is always = 1
}

func (db *DB) ImportItem() ([]*Item, error) {
	rows, err := db.Query(
		"SELECT itemID, code, nameTH, nameEN, lastSalePrice, defStkUnitCode, defBuyUnitCode, defSaleUnitCode, itemTypeID, itemBrandID FROM nava.tb_sv_item ORDER BY itemID")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := make([]*Item, 0)

	var (
		id, typeID, brandID                    sql.NullInt64
		code, nTH, nEN, p, uStock, uBuy, uSell sql.NullString
	)

	for rows.Next() {
		i := new(Item)
		err := rows.Scan(&id, &code, &nTH, &nEN, &p, &uStock, &uBuy, &uSell, &typeID, &brandID)
		if err != nil {
			return nil, err
		}
		if id.Valid {
			i.ID = int(id.Int64)
		}
		if code.Valid {
			i.Code = strings.Trim(code.String, " ")
		}
		if nTH.Valid {
			i.NameTH = strings.Trim(nTH.String, " ")
		}
		if nEN.Valid {
			i.NameEN = strings.Trim(nEN.String, " ")
		}
		if p.Valid {
			i.SellPrice, err = decimal.NewFromString(p.String)
			if err != nil {
				log.Fatal("after decimal.NewFromString()", err)
			}
		}
		if uStock.Valid {
			i.UnitStock = strings.Trim(uStock.String, " ")
		}
		if uBuy.Valid {
			i.UnitBuy = strings.Trim(uBuy.String, " ")
		}
		if uSell.Valid {
			i.UnitSell = strings.Trim(uSell.String, " ")
		}
		if typeID.Valid {
			i.TypeID = int(typeID.Int64)
		}
		if brandID.Valid {
			i.BrandID = int(brandID.Int64)
		}

		items = append(items, i)
		//		fmt.Printf("Append Item ID: %d Name: %s Unit: %s Price: %v\n", i.ID, i.Name, i.UOM, i.StdPrice)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
