package tests

import (
	"fmt"
	"log"
	"testing"
	"worker-transaction/models"
)

func TestGetFromSymbol(t *testing.T) {
	var currency *models.CurrencyContract
	var err error
	currency, err = currency.GetFromSymbol("LINK")
	if err != nil {
		log.Println(err.Error(), "err.Error() main.go:148")
		t.Fail()
	}

	currency_test := &models.CurrencyContract{
		Id:      1,
		Name:    "ChainLink",
		Symbol:  "Link",
		Address: "0x326C977E6efc84E512bB9C30f76E30c160eD06FB",
		Chain:   "POLYGON",
		Decimal: 18,
		Img:     "https://s2.coinmarketcap.com/static/img/coins/64x64/12389.png",
	}

	if fmt.Sprint(currency) == fmt.Sprint(currency_test) {

	} else {
		t.Errorf("ket qua khong chinh xacs %s, %s, %d %d %s", currency.Address, currency.Chain, currency.Id, currency.Decimal, currency.Name)
		t.Fail()
	}
}
