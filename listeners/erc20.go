package listeners

import (
	"context"
	"fmt"
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

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gammazero/workerpool"
	"github.com/sirupsen/logrus"
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

	log.Println("=== start watch ===", chain)

	wkPool := workerpool.New(appconfig.Config.WkPoolBsc)
	client := blockchainservice.GetEthClientService(chain)
	if client == nil {
		fmt.Println("client nil")
		client = blockchainservice.GetNewEthClientService(chain).GetClient()
		if client == nil {
			log.Println("===== Client is nil =====")
			log.Println(client, "client worker/worker.go:62")
			panic("===== Client is nil ===== worker/worker.go:64")
		}
	}

	latestBlockIndex := commonservice.GetLatestIndexBlock()
	fmt.Println(latestBlockIndex, "latestBlockIndex")

	mapContract := map[string]string{}

	// get contract erc20 to listen approve
	listContractERC20, total, err := (&models.CurrencyContract{Chain: chain}).GetAll()
	fmt.Println(listContractERC20, "listContractERC20")
	log.Println("=== total contract erc20 === ", total-1)
	if err != nil {
		log.Println(err.Error(), "err.Error() erc20.go:59")
	} else {
		for _, contract := range listContractERC20 {
			if contract.Address != consts.ZERO_ADDRESS {
				mapContract[contract.Address] = ""
			}
		}
	}

	methodTransfer := common.Bytes2Hex(crypto.Keccak256([]byte("transfer(address,uint256)"))[:4])

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
			log.Println(latestBlockIndex, "latestBlockIndex erc20.go:84")
			time.Sleep(2 * time.Second)
			continue
		}

		// logrus.WithFields(logrus.Fields{
		// 	"Block Number": block.Number().String(),
		// 	"Hash":         block.Hash().Hex(),
		// 	"Trans len":    block.Transactions().Len(),
		// }).Info("=== header log === erc20.go:93")

		listTrans := block.Transactions()
		for _, tx := range listTrans {
			txClone := tx
			if tx == nil || tx.To() == nil {
				continue
			}

			if len(txClone.Data()) == 0 {
				wkPool.Submit(func() {
					TransferCoinbase(txClone, chain)
				})
			} else {
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
					case methodTransfer:
						//fmt.Println(txClone.Hash().Hex(), txClone.Value().String())
						//fmt.Println(txClone.To().Hex())
						//fmt.Println(txClone.Data())
						TransferTokenERC20(txClone, chain)
					}
				})
			}
		}
		latestBlockIndex = latestBlockIndex.Add(latestBlockIndex, big.NewInt(1))
		commonservice.PutLatestIndexBlock(latestBlockIndex)
	}
}

func TransferTokenERC20(trans *types.Transaction, chain string) {
	log.Println("=== Start Transfer ERC20 === erc20.go 138")
	defer log.Println("=== End Transfer ERC20 === erc20.go 139")
	log.Println(trans.Hash(), "trans erc20.go:140")

	tx, err := blockchainservice.GetEthClientService(chain).TransactionReceipt(context.Background(), trans.Hash())
	if err != nil {
		tx, err = blockchainservice.GetNewEthClientService(chain).TransactionReceipt(context.Background(), trans.Hash())
		if err != nil {
			log.Println(err.Error(), "err.Error() cmd/worker/erc20/erc20.go:27")
			return
		}
	}

	// transfer fail
	if tx.Status == 0 {
		return
	}

	toAddress := "0x" + common.Bytes2Hex(trans.Data()[4:36])
	value, ok := big.NewInt(0).SetString(common.Bytes2Hex(trans.Data()[36:68]), 16)
	if !ok {
		fmt.Println("erc20.go 159 parse value data error")
	}
	logrus.WithFields(logrus.Fields{
		"Address": toAddress,
		"Value":   value.String(),
	}).Info("=== print data transfer erc20 === erc20.go:164")
}

func TransferCoinbase(trans *types.Transaction, chain string) {
	log.Println("=== Start Transfer Coinbase === erc20.go 169")
	defer log.Println("=== End Transfer Coinbase === erc20.go 170")
	log.Println(trans.Hash(), "trans erc20.go:170")

	toAddress := trans.To().Hex()
	value := trans.Value()

	logrus.WithFields(logrus.Fields{
		"Address": toAddress,
		"Value":   value.String(),
	}).Info("=== print data transfer conbase === erc20.go:177")
}
