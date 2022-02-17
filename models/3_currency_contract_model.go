package models

import (
	"log"
	"worker-transaction/services/databaseservice"
)

//go:generate easytags $GOFILE json,xml

type CurrencyContract struct {
	Id      int    `json:"id" xml:"id"`
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

func (c *CurrencyContract) GetFromId(id int) (*CurrencyContract, error) {
	store := databaseservice.GetDatabase()
	u := &CurrencyContract{}
	rows, err := store.Db.Query("SELECT (id, name, symbol, address, chain, decimal, img) FROM currency WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	err = rows.Scan(&u.Id, &u.Name, &u.Symbol, &u.Address, &u.Chain, &u.Decimal, &u.Img)
	if err != nil {
		log.Println(err.Error(), "err.Error() models/3_currency_contract_model.go:32")
		return nil, err
	}
	return u, nil
}

func (c *CurrencyContract) GetFromSymbol(symbol string) (*CurrencyContract, error) {
	store := databaseservice.GetDatabase()
	u := &CurrencyContract{}
	rows, err := store.Db.Query("SELECT * FROM currency WHERE symbol = $1", symbol)
	if err != nil {
		return nil, err
	}
	if rows.Next() {
		err = rows.Scan(&u.Id, &u.Name, &u.Symbol, &u.Address, &u.Chain, &u.Decimal, &u.Img)
		if err != nil {
			log.Println(err.Error(), "err.Error() models/3_currency_contract_model.go:32")
			return nil, err
		}
	}
	return u, nil
}

func (c *CurrencyContract) PutItem() error {
	store := databaseservice.GetDatabase()
	_, err := store.Db.Exec(
		"UPDATE currency SET name = $2, symbol = $3, address = $4, chain = $5, decimal = $6, img = $7 WHERE currency.id = $1",
		c.Id, c.Name, c.Symbol, c.Address, c.Chain, c.Decimal, c.Img)
	if err != nil {
		return err
	}
	return nil
}

func (c *CurrencyContract) Del() error {
	store := databaseservice.GetDatabase()
	_, err := store.Db.Exec(
		"DELETE FROM currency WHERE id = $1",
		c.Id)
	if err != nil {
		return err
	}
	return nil
}

func (c *CurrencyContract) GetAll() ([]CurrencyContract, int, error) {
	listItem := make([]CurrencyContract, 0)
	store := databaseservice.GetDatabase()
	rows, err := store.Db.Query("SELECT * FROM currency")
	if err != nil {
		return nil, 0, err
	}
	for rows.Next() {
		u := CurrencyContract{}
		err := rows.Scan(&u.Id, &u.Name, &u.Symbol, &u.Address, &u.Chain, &u.Decimal, &u.Img)
		if err != nil {
			return nil, 0, err
		}
		listItem = append(listItem, u)
	}

	return listItem, len(listItem), err
}
