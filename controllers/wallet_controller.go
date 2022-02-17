package controllers

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math"
	"math/big"
	"net/http"
	"strconv"
	"worker-transaction/consts"
	"worker-transaction/models"
	"worker-transaction/restapi/requests"
	"worker-transaction/restapi/responses"
	"worker-transaction/services/blockchainservice"
	"worker-transaction/services/responseservice"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/sha3"
)

// @Title Withdrawal
// @Summary Withdrawal from wallet
// @Description Withdrawal from wallet
// @Tags user
// @Accept  json
// @Produce  json
// @Param sessionkey    header  string  true    "session key"
// @Param	body	body	requests.WithdrawalRequest	true	"body withdrawal"
// @Success 200	{object} responses.ResponseCommonSingle
// @Failure 403 {pubKey} not found
// @router /withdrawal [post]
func Withdrawal(c *gin.Context) {

	//get request
	var requestBody requests.WithdrawalRequest
	err := c.ShouldBindJSON(&requestBody)
	if err != nil {

		log.Println(err.Error(), "err.Error() main.go:696")
		c.JSON(http.StatusOK, responseservice.GetCommonErrorResponse(responses.BadRequest))
		return
	}

	client := blockchainservice.GetEthClientService(consts.POLYGON_CHAIN)
	toAddress := common.HexToAddress(requestBody.ToAddress)

	//private key of wallet
	privateKey, err := crypto.HexToECDSA("191154ed0a0cf19db8619dee5654365507e71bae18409324b56c5e8faf71d508")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
		c.JSON(http.StatusOK, responseservice.GetCommonErrorResponse(responses.BadRequest))
		return
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusOK, responseservice.GetCommonErrorResponse(responses.BadRequest))
		return
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusOK, responseservice.GetCommonErrorResponse(responses.BadRequest))
		return
	}

	var gasLimit uint64
	gasLimit = 40000
	fmt.Println(gasLimit) // 23256

	// -> xacs ddinh thong tin chung cua nguoi nhan

	var data []byte
	var value *big.Int
	var tx *types.Transaction

	if requestBody.Coin == "MATIC" {
		value = big.NewInt(0)
		var f float64
		f, _ = strconv.ParseFloat(requestBody.Amount, 10)
		f = f * 1e18
		value.SetInt64(int64(f)) // in wei (with amount eth)
		if requestBody.Type == consts.TRANS_TYPE_EIP1559 {
			tx = types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)
		} else {
			tx = types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)
		}
	} else {
		value = big.NewInt(0) // in wei (0 eth)

		var currency *models.CurrencyContract
		currency, err = currency.GetFromSymbol(requestBody.Coin)
		if err != nil {
			log.Println(err.Error(), "err.Error() main.go:148")
			c.JSON(http.StatusOK, responseservice.GetCommonErrorResponse(responses.BadRequest))
			return
		}

		log.Println(currency.Address, "Variable controllers/wallet_controller.go:113")

		tokenAddress := common.HexToAddress(currency.Address)

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
		f = f * math.Pow(10, float64(currency.Decimal))
		amount.SetInt64(int64(f))
		// var decimals int64
		// decimals = 18
		// amount = amount.Mul(amount, big.NewInt(decimals))
		//padding amount
		paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
		fmt.Println(hexutil.Encode(paddedAmount))

		data = append(data, methodID...)
		data = append(data, paddedAddress...)
		data = append(data, paddedAmount...)

		tx = types.NewTransaction(nonce, tokenAddress, value, gasLimit, gasPrice, data)
	}

	// gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
	// 	To:   &toAddress,
	// 	Data: data,
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusOK, responseservice.GetCommonErrorResponse(responses.BadRequest))
		return
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusOK, responseservice.GetCommonErrorResponse(responses.BadRequest))
		return
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusOK, responseservice.GetCommonErrorResponse(responses.BadRequest))
		return
	}

	fmt.Printf("tx sent: %s\n", signedTx.Hash().Hex()) //tx sent
	c.JSON(http.StatusOK, responseservice.GetCommonSucceedResponse())
}
