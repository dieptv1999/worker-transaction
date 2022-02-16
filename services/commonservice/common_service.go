package commonservice

import (
	"context"
	"io/ioutil"
	"log"
	"math/big"
	"strconv"
	"worker-transaction/consts"
	"worker-transaction/models"
	"worker-transaction/restapi/responses"
	"worker-transaction/services/blockchainservice"
)

func GetLatestIndexBlock() *big.Int {
	data, err := ioutil.ReadFile("/home/dieptv/GolandProjects/worker-transaction/last_index.txt")
	if err != nil {
		log.Println(err.Error(), "err.Error() services/commonservice/common_service.go:17")
		return GetCurrentBlock()
	}

	latestBlockIndex, err := strconv.ParseInt(string(data), 10, 64)
	if err != nil {
		log.Println(err.Error(), "err.Error() services/commonservice/common_service.go:24")
		return GetCurrentBlock()
	}

	return big.NewInt(latestBlockIndex)
	return nil
}

func PutLatestIndexBlock(latestBlockIndex *big.Int) {
	err := ioutil.WriteFile("/home/dieptv/GolandProjects/worker-transaction/last_index.txt", []byte(latestBlockIndex.String()), 0644)
	if err != nil {
		log.Println(err.Error(), "err.Error() services/commonservice/common_service.go:35")
	}

	//models.GetBsNftOriginal().BsPutItem(consts.BS_COMMON_VALUE, &generic.TItem{
	//	Key:   []byte(consts.COMMON_LATEST_INDEX_BLOCK),
	//	Value: []byte(latestBlockIndex.String()),
	//})
}

func GetCurrentBlock() *big.Int {
	blockCurrent, err := blockchainservice.GetEthClientService(consts.BSC_CHAIN).BlockNumber(context.Background())
	if err != nil {
		log.Println(err.Error(), "err.Error() services/commonservice/common_service.go:41")
		return big.NewInt(0)
	}

	return big.NewInt(int64(blockCurrent))
}

func GetListCurrency() ([]models.CurrencyContract, int, error) {
	listCurrency, total, err := (&models.CurrencyContract{Chain: consts.BSC_CHAIN}).GetAll()
	if err != nil {
		log.Println(err.Error(), "err.Error() services/commonservice/common_service.go:61")
		return []models.CurrencyContract{}, 0, responses.ErrSystem
	}

	return listCurrency, total, responses.Success
}
