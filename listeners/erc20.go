package listeners

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gammazero/workerpool"
	"github.com/sirupsen/logrus"
	"log"
	"math/big"
	"strings"
	"sync"
	"time"
	"worker-transaction/appconfig"
	"worker-transaction/consts"
	"worker-transaction/models"
	"worker-transaction/services/blockchainservice"
	"worker-transaction/services/commonservice"
)

type CallbackFunc func()

var (
	m = &sync.RWMutex{}
)

func WatchEventErc20(chain string, fn CallbackFunc) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f, maybe ws interrupt", r)
			fn()
		}
	}()

	log.Println("=== start watch ===")

	wkPool := workerpool.New(appconfig.Config.WkPoolBsc)
	client := blockchainservice.GetEthClientService(chain)
	if client == nil {
		client = blockchainservice.GetNewEthClientService(chain).GetClient()
		if client == nil {
			log.Println("===== Client is nil =====")
			log.Println(client, "client worker/worker.go:62")
			panic("===== Client is nil ===== worker/worker.go:64")
		}
	}

	latestBlockIndex := commonservice.GetLatestIndexBlock()

	mapContract := map[string]string{}

	// get contract erc20 to listen approve
	listContractERC20, total, err := (&models.CurrencyContract{Chain: consts.BSC_CHAIN}).GetAll()
	log.Println("=== total contract erc20 === ", total-1)
	if err != nil {
		log.Println(err.Error(), "err.Error() worker/worker.go:95")
	} else {
		for _, contract := range listContractERC20 {
			if contract.Address != consts.ZERO_ADDRESS {
				mapContract[contract.Address] = ""
			}
		}
	}

	for {
		log.Println("===== " + latestBlockIndex.String() + " =====")
		block, err := client.BlockByNumber(context.Background(), latestBlockIndex)
		if err != nil {
			if !strings.Contains(err.Error(), "not found") {
				if err.Error() == "client is closed" || err.Error() == "connection lost" ||
					strings.Contains(err.Error(), "Bad Gateway") {
					newClient := blockchainservice.GetNewEthClientService(chain)
					if newClient != nil {
						client.Close()
						client = newClient.GetClient()
					}
				}
			}

			log.Println("===== ETH READ TO FAST =====")
			log.Println(latestBlockIndex, "latestBlockIndex worker/worker.go:81")
			time.Sleep(2 * time.Second)
			continue
		}

		logrus.WithFields(logrus.Fields{
			"Block Number": block.Number().String(),
			"Hash":         block.Hash().Hex(),
			"Trans len":    block.Transactions().Len(),
		}).Info("=== header log === workerv2/worker.go:90")

		listTrans := block.Transactions()
		for _, tx := range listTrans {
			txClone := tx
			if tx == nil || tx.To() == nil {
				continue
			}

			m.RLock()
			_, exist := mapContract[fmt.Sprintf("%s", tx.To())]
			m.RUnlock()

			if !exist {
				continue
			}
			wkPool.Submit(func() {
				inputData := common.Bytes2Hex(txClone.Data())
				strHashFunc := inputData[:8]
				switch strHashFunc {
				}
			})

		}
		latestBlockIndex = latestBlockIndex.Add(latestBlockIndex, big.NewInt(1))
		commonservice.PutLatestIndexBlock(latestBlockIndex)
	}
}
