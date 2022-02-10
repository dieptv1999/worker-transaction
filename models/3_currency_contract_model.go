package models

import (
	"worker-transaction/services/databaseservice"
)

//go:generate easytags $GOFILE json,xml

type CurrencyContract struct {
	Id      string `json:"id" xml:"id"`
	Name    string `json:"name" xml:"name"`
	Symbol  string `json:"symbol" xml:"symbol"`
	Address string `json:"address" xml:"address"`
	Chain   string `json:"chain" xml:"chain"`
	Decimal int    `json:"decimal" xml:"decimal"`
	Img     string `json:"img" xml:"img"`
}

func (c *CurrencyContract) String() string {
	return c.Address
}

func (c *CurrencyContract) GetFromKey(key int) (*CurrencyContract, error) {
	db := databaseservice.Database
	u := &CurrencyContract{}
	rows, err := db.Query("SELECT (name, symbol, address, chain, decimal, img) FROM currency WHERE id = $1", key)
	if err != nil {
		return nil, err
	}
	rows.Scan(&u.Name, &u.Symbol, &u.Address, &u.Chain, &u.Decimal, &u.Img)
	return u, nil
}

func (c *CurrencyContract) PutItem() error {
	db := databaseservice.Database
	_, err := db.Exec(
		"UPDATE currency SET name = $2, symbol = $3, address = $4, chain = $5, decimal = $6, img = $7 WHERE currency.id = $1",
		c.Id, c.Name, c.Symbol, c.Address, c.Chain, c.Decimal, c.Img)
	if err != nil {
		return err
	}
	return nil
}

func (c *CurrencyContract) Del() error {
	db := databaseservice.Database
	_, err := db.Exec(
		"DELETE FROM currency WHERE id = $1",
		c.Id)
	if err != nil {
		return err
	}
	return nil
}

func (c *CurrencyContract) GetAll() ([]CurrencyContract, int, error) {
	listItem := make([]CurrencyContract, 0)
	db := databaseservice.Database
	rows, err := db.Query("SELECT * FROM currency")
	if err != nil {
		return nil, 0, err
	}
	for rows.Next() {
		u := CurrencyContract{}
		err := rows.Scan(&u.Name, &u.Symbol, &u.Address, &u.Chain, &u.Decimal, &u.Img)
		if err != nil {
			return nil, 0, err
		}
		listItem = append(listItem, u)
	}

	return listItem, len(listItem), err
}
