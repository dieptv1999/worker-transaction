package blockchainservice

import (
	"encoding/hex"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"log"
	"math/big"
	"strings"
	"sync"
	"worker-transaction/consts"
	"worker-transaction/contracts/erc20"
	"worker-transaction/helps"
	"worker-transaction/restapi/responses"
)

var (
	mintERC20ABI     abi.ABI
	mintERC20ABIOnce sync.Once
)

func GetERC20ABI() abi.ABI {
	mintERC20ABIOnce.Do(func() {
		var err error
		mintERC20ABI, err = abi.JSON(strings.NewReader(erc20.Erc20ABI))
		if err != nil {
			log.Println(err.Error(), "err.Error() can't load abi services/blockchainservice/erc20_service.go:21")
		}
	})
	return mintERC20ABI
}

func GetInfoErc20(address common.Address) (decimal int, symbol, name string, err error) {
	erc20Cli, err := erc20.NewErc20(address, GetEthClientService(consts.BSC_CHAIN))
	if err != nil {
		log.Println(err.Error(), "err services/blockchainservice/erc20_service.go:32")
		return 0, "", "", err
	}
	uDecimal, err := erc20Cli.Decimals(&bind.CallOpts{})
	if err != nil {
		log.Println(err.Error(), "err.Error() services/blockchainservice/erc20_service.go:37")
		return 0, "", "", err
	}
	symbol, err = erc20Cli.Symbol(&bind.CallOpts{})
	if err != nil {
		log.Println(err.Error(), "err.Error() services/blockchainservice/erc20_service.go:42")
		return int(uDecimal), "", "", err
	}

	name, err = erc20Cli.Name(&bind.CallOpts{})
	if err != nil {
		log.Println(err.Error(), "err.Error() services/blockchainservice/erc20_service.go:48")
		return int(uDecimal), symbol, "", err
	}

	return int(uDecimal), symbol, name, nil
}

func GetDecimalErc20(address common.Address) (int64, error) {
	erc20Cli, err := erc20.NewErc20(address, GetEthClientService(consts.BSC_CHAIN))
	if err != nil {
		log.Println(err.Error(), "err services/blockchainservice/erc20_service.go:58")
		return int64(0), err
	}
	uDecimal, err := erc20Cli.Decimals(&bind.CallOpts{})
	if err != nil {
		log.Println(err.Error(), "err.Error() services/blockchainservice/erc20_service.go:63")
		return int64(0), err
	}

	return int64(uDecimal), nil
}

func approveTokenERC20(userBlockchainAddress string, toAddress, approveTo common.Address, price *big.Int) (r *responses.DataApprove, err error) {
	fromAddress := common.HexToAddress(helps.VerifyContractAddress(userBlockchainAddress))

	input, err := GetERC20ABI().Pack("approve",
		approveTo,
		price,
	)
	if err != nil {
		log.Println(err.Error(), "err.Error() services/blockchainservice/erc20_service.go:89")
		return &responses.DataApprove{}, responses.MethodNotAllowed
	}

	gasLimit, gasPrice, err := getGas(input, nil, fromAddress, approveTo)
	if err != nil {
		log.Println(err.Error(), "err.Error() services/blockchainservice/erc20_service.go:95")
		return &responses.DataApprove{}, responses.MethodNotAllowed
	}

	return &responses.DataApprove{
		From:     userBlockchainAddress,
		To:       toAddress.Hex(),
		Gas:      gasLimit,
		GasPrice: gasPrice,
		Data:     hex.EncodeToString(input),
	}, responses.Success
}

func GetDataApproveERC20IfNeed(userBlockchainAddress string, toAddress, marketContract common.Address, price *big.Int) (r *responses.DataApprove, err error) {
	erc20Instance, err := erc20.NewErc20(toAddress, GetEthClientService(consts.BSC_CHAIN))
	if err != nil {
		log.Println(err.Error(), "err.Error() services/blockchainservice/erc20_service.go:106")
		return nil, err
	}

	allowance, err := erc20Instance.Allowance(&bind.CallOpts{}, common.HexToAddress(userBlockchainAddress), marketContract)
	if err != nil {
		log.Println(err.Error(), "err.Error() services/blockchainservice/erc20_service.go:112")
		return approveTokenERC20(userBlockchainAddress, toAddress, marketContract, price)
	}

	if price.Cmp(allowance) > 0 {
		return approveTokenERC20(userBlockchainAddress, toAddress, marketContract, price)
	}

	return nil, nil
}
