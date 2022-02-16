package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"strconv"
	"worker-transaction/appconfig"
	"worker-transaction/consts"
	"worker-transaction/restapi/requests"
	"worker-transaction/services/blockchainservice"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gammazero/workerpool"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/sha3"
)

func main() {
	appconfig.LoadInitConfig()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	r := gin.Default()
	taskPool := workerpool.New(appconfig.Config.NumPoolDepositWithdrawal)

	r.POST("/deposit", func(c *gin.Context) {
		//get request
		var requestBody requests.WithdrawalRequest
		err := c.ShouldBindJSON(&requestBody)
		if err != nil {

			log.Println(err.Error(), "err.Error() main.go:696")
			// c.JSON(http.StatusOK, responses.GetCommonErrorResponse(responses.BadRequest))
			return
		}

		taskPool.Submit(func() {
			client := blockchainservice.GetEthClientService(consts.POLYGON_CHAIN)
			toAddress := common.HexToAddress(requestBody.ToAddress)
			tokenAddress := common.HexToAddress("0x326C977E6efc84E512bB9C30f76E30c160eD06FB")

			//private key of wallet
			privateKey, err := crypto.HexToECDSA("191154ed0a0cf19db8619dee5654365507e71bae18409324b56c5e8faf71d508")
			if err != nil {
				log.Fatal(err)
			}

			publicKey := privateKey.Public()
			publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
			if !ok {
				log.Fatal("error casting public key to ECDSA")
			}

			fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
			nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
			if err != nil {
				log.Fatal(err)
			}

			value := big.NewInt(0) // in wei (0 eth)
			gasPrice, err := client.SuggestGasPrice(context.Background())
			if err != nil {
				log.Fatal(err)
			}

			transferFnSignature := []byte("transfer(address,uint256)")
			hash := sha3.NewLegacyKeccak256()
			hash.Write(transferFnSignature)
			methodID := hash.Sum(nil)[:4]
			fmt.Println(hexutil.Encode(methodID)) // 0xa9059cbb

			//padding address
			paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)
			fmt.Println(hexutil.Encode(paddedAddress))

			amount := new(big.Int)
			amount.SetString(requestBody.Amount, 10)
			//padding amount
			paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
			fmt.Println(hexutil.Encode(paddedAmount))

			var data []byte
			data = append(data, methodID...)
			data = append(data, paddedAddress...)
			data = append(data, paddedAmount...)

			// gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
			// 	To:   &toAddress,
			// 	Data: data,
			// })
			// if err != nil {
			// 	log.Fatal(err)
			// }
			var gasLimit uint64
			gasLimit = 64346
			fmt.Println(gasLimit) // 23256

			tx := types.NewTransaction(nonce, tokenAddress, value, gasLimit, gasPrice, data)

			chainID, err := client.NetworkID(context.Background())
			if err != nil {
				log.Fatal(err)
			}

			signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
			if err != nil {
				log.Fatal(err)
			}

			err = client.SendTransaction(context.Background(), signedTx)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("tx sent: %s", signedTx.Hash().Hex()) //tx sent
		})
	})

	r.POST("/withdrawal", func(c *gin.Context) {
		//get request
		var requestBody requests.WithdrawalRequest
		err := c.ShouldBindJSON(&requestBody)
		if err != nil {

			log.Println(err.Error(), "err.Error() main.go:696")
			// c.JSON(http.StatusOK, responses.GetCommonErrorResponse(responses.BadRequest))
			return
		}

		taskPool.Submit(func() {
			client := blockchainservice.GetEthClientService(consts.POLYGON_CHAIN)
			toAddress := common.HexToAddress(requestBody.ToAddress)
			tokenAddress := common.HexToAddress("0x326C977E6efc84E512bB9C30f76E30c160eD06FB")

			//private key of wallet
			privateKey, err := crypto.HexToECDSA("191154ed0a0cf19db8619dee5654365507e71bae18409324b56c5e8faf71d508")
			if err != nil {
				log.Fatal(err)
			}

			publicKey := privateKey.Public()
			publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
			if !ok {
				log.Fatal("error casting public key to ECDSA")
			}

			fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
			nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
			if err != nil {
				log.Fatal(err)
			}

			gasPrice, err := client.SuggestGasPrice(context.Background())
			if err != nil {
				log.Fatal(err)
			}

			// -> xacs ddinh thong tin chung cua nguoi nhan

			var data []byte
			var value *big.Int

			if requestBody.Coin == "MATIC" {
				value = big.NewInt(0)
				var f float64
				f, _ = strconv.ParseFloat(requestBody.Amount, 10)
				f = f * 1e18
				value.SetInt64(int64(f)) // in wei (with amount eth)
				var decimals int64
				decimals = 18
				value = value.Mul(value, big.NewInt(decimals))
			} else {
				value = big.NewInt(0) // in wei (0 eth)

				transferFnSignature := []byte("transfer(address,uint256)")
				hash := sha3.NewLegacyKeccak256()
				hash.Write(transferFnSignature)
				methodID := hash.Sum(nil)[:4]
				fmt.Println(hexutil.Encode(methodID)) // 0xa9059cbb

				//padding address
				paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)
				fmt.Println(hexutil.Encode(paddedAddress))

				amount := new(big.Int)
				var f float64
				f, _ = strconv.ParseFloat(requestBody.Amount, 10)
				f = f * 1e18
				amount.SetInt64(int64(f))
				var decimals int64
				decimals = 18
				amount = amount.Mul(amount, big.NewInt(decimals))
				//padding amount
				paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
				fmt.Println(hexutil.Encode(paddedAmount))

				data = append(data, methodID...)
				data = append(data, paddedAddress...)
				data = append(data, paddedAmount...)
			}

			// gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
			// 	To:   &toAddress,
			// 	Data: data,
			// })
			// if err != nil {
			// 	log.Fatal(err)
			// }
			var gasLimit uint64
			gasLimit = 64346
			fmt.Println(gasLimit) // 23256

			tx := types.NewTransaction(nonce, tokenAddress, value, gasLimit, gasPrice, data)

			chainID, err := client.NetworkID(context.Background())
			if err != nil {
				log.Fatal(err)
			}

			signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
			if err != nil {
				log.Fatal(err)
			}

			err = client.SendTransaction(context.Background(), signedTx)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("tx sent: %s", signedTx.Hash().Hex()) //tx sent
		})
	})
	r.Run()
}
