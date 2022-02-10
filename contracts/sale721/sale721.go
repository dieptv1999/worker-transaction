// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package sale721

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// TrustKeysMarketplaceSaleInfo is an auto generated low-level Go binding around an user-defined struct.
type TrustKeysMarketplaceSaleInfo struct {
	Owner         common.Address
	Price         *big.Int
	ContractERC20 common.Address
	Exist         bool
}

// Sale721ABI is the input ABI used to generate the binding from.
const Sale721ABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"contractIERC721Metadata\",\"name\":\"contractNFT\",\"type\":\"address\"}],\"name\":\"Canceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"contractIERC721Metadata\",\"name\":\"contractNFT\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"contractERC20\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"exist\",\"type\":\"bool\"}],\"indexed\":false,\"internalType\":\"structTrustKeysMarketplace.SaleInfo\",\"name\":\"saleInfo\",\"type\":\"tuple\"}],\"name\":\"ListingPrice\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"contractIERC721Metadata\",\"name\":\"contractNFT\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"contractERC20\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"exist\",\"type\":\"bool\"}],\"indexed\":false,\"internalType\":\"structTrustKeysMarketplace.SaleInfo\",\"name\":\"saleInfo\",\"type\":\"tuple\"}],\"name\":\"Sold\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"contractIERC721Metadata\",\"name\":\"contractNFT\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"contractERC20\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"exist\",\"type\":\"bool\"}],\"indexed\":false,\"internalType\":\"structTrustKeysMarketplace.SaleInfo\",\"name\":\"saleInfo\",\"type\":\"tuple\"}],\"name\":\"Updated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC721Metadata\",\"name\":\"_contractNFT\",\"type\":\"address\"}],\"name\":\"BuyToken\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC721Metadata\",\"name\":\"_contractNFT\",\"type\":\"address\"}],\"name\":\"CancelSaleToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC721Metadata\",\"name\":\"_contractNFT\",\"type\":\"address\"}],\"name\":\"EstimatePrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC721Metadata\",\"name\":\"_contractNFT\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"_contractERC20\",\"type\":\"address\"}],\"name\":\"SaleToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC721Metadata\",\"name\":\"_contractNFT\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"_contractERC20\",\"type\":\"address\"}],\"name\":\"UpdatePrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"__TrustKeysMarketplace_init_\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"addressTrustKeysNFTUpgradeable\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"currencies\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeBuy\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeDecimal\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeSale\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"limitCurrency\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"marketplace\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"contractERC20\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"exist\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"royalties\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addressTrustKeysNFT\",\"type\":\"address\"}],\"name\":\"setAddressTrustKeysNFT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_listContractERC20\",\"type\":\"address[]\"}],\"name\":\"setCurrencies\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_contractERC20\",\"type\":\"address\"}],\"name\":\"setCurrency\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_limitCurrency\",\"type\":\"bool\"}],\"name\":\"setLimitCurrency\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_royalties\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_feeBuy\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_feeSale\",\"type\":\"uint32\"}],\"name\":\"updateFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Sale721 is an auto generated Go binding around an Ethereum contract.
type Sale721 struct {
	Sale721Caller     // Read-only binding to the contract
	Sale721Transactor // Write-only binding to the contract
	Sale721Filterer   // Log filterer for contract events
}

// Sale721Caller is an auto generated read-only Go binding around an Ethereum contract.
type Sale721Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Sale721Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Sale721Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Sale721Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Sale721Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Sale721Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Sale721Session struct {
	Contract     *Sale721          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Sale721CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Sale721CallerSession struct {
	Contract *Sale721Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// Sale721TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Sale721TransactorSession struct {
	Contract     *Sale721Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// Sale721Raw is an auto generated low-level Go binding around an Ethereum contract.
type Sale721Raw struct {
	Contract *Sale721 // Generic contract binding to access the raw methods on
}

// Sale721CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Sale721CallerRaw struct {
	Contract *Sale721Caller // Generic read-only contract binding to access the raw methods on
}

// Sale721TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Sale721TransactorRaw struct {
	Contract *Sale721Transactor // Generic write-only contract binding to access the raw methods on
}

// NewSale721 creates a new instance of Sale721, bound to a specific deployed contract.
func NewSale721(address common.Address, backend bind.ContractBackend) (*Sale721, error) {
	contract, err := bindSale721(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Sale721{Sale721Caller: Sale721Caller{contract: contract}, Sale721Transactor: Sale721Transactor{contract: contract}, Sale721Filterer: Sale721Filterer{contract: contract}}, nil
}

// NewSale721Caller creates a new read-only instance of Sale721, bound to a specific deployed contract.
func NewSale721Caller(address common.Address, caller bind.ContractCaller) (*Sale721Caller, error) {
	contract, err := bindSale721(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Sale721Caller{contract: contract}, nil
}

// NewSale721Transactor creates a new write-only instance of Sale721, bound to a specific deployed contract.
func NewSale721Transactor(address common.Address, transactor bind.ContractTransactor) (*Sale721Transactor, error) {
	contract, err := bindSale721(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Sale721Transactor{contract: contract}, nil
}

// NewSale721Filterer creates a new log filterer instance of Sale721, bound to a specific deployed contract.
func NewSale721Filterer(address common.Address, filterer bind.ContractFilterer) (*Sale721Filterer, error) {
	contract, err := bindSale721(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Sale721Filterer{contract: contract}, nil
}

// bindSale721 binds a generic wrapper to an already deployed contract.
func bindSale721(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Sale721ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sale721 *Sale721Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Sale721.Contract.Sale721Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sale721 *Sale721Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sale721.Contract.Sale721Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sale721 *Sale721Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sale721.Contract.Sale721Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sale721 *Sale721CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Sale721.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sale721 *Sale721TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sale721.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sale721 *Sale721TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sale721.Contract.contract.Transact(opts, method, params...)
}

// EstimatePrice is a free data retrieval call binding the contract method 0xf8dec8e4.
//
// Solidity: function EstimatePrice(uint256 _tokenId, address _contractNFT) view returns(uint256, address)
func (_Sale721 *Sale721Caller) EstimatePrice(opts *bind.CallOpts, _tokenId *big.Int, _contractNFT common.Address) (*big.Int, common.Address, error) {
	var out []interface{}
	err := _Sale721.contract.Call(opts, &out, "EstimatePrice", _tokenId, _contractNFT)

	if err != nil {
		return *new(*big.Int), *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return out0, out1, err

}

// EstimatePrice is a free data retrieval call binding the contract method 0xf8dec8e4.
//
// Solidity: function EstimatePrice(uint256 _tokenId, address _contractNFT) view returns(uint256, address)
func (_Sale721 *Sale721Session) EstimatePrice(_tokenId *big.Int, _contractNFT common.Address) (*big.Int, common.Address, error) {
	return _Sale721.Contract.EstimatePrice(&_Sale721.CallOpts, _tokenId, _contractNFT)
}

// EstimatePrice is a free data retrieval call binding the contract method 0xf8dec8e4.
//
// Solidity: function EstimatePrice(uint256 _tokenId, address _contractNFT) view returns(uint256, address)
func (_Sale721 *Sale721CallerSession) EstimatePrice(_tokenId *big.Int, _contractNFT common.Address) (*big.Int, common.Address, error) {
	return _Sale721.Contract.EstimatePrice(&_Sale721.CallOpts, _tokenId, _contractNFT)
}

// AddressTrustKeysNFTUpgradeable is a free data retrieval call binding the contract method 0x947beddb.
//
// Solidity: function addressTrustKeysNFTUpgradeable() view returns(address)
func (_Sale721 *Sale721Caller) AddressTrustKeysNFTUpgradeable(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Sale721.contract.Call(opts, &out, "addressTrustKeysNFTUpgradeable")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddressTrustKeysNFTUpgradeable is a free data retrieval call binding the contract method 0x947beddb.
//
// Solidity: function addressTrustKeysNFTUpgradeable() view returns(address)
func (_Sale721 *Sale721Session) AddressTrustKeysNFTUpgradeable() (common.Address, error) {
	return _Sale721.Contract.AddressTrustKeysNFTUpgradeable(&_Sale721.CallOpts)
}

// AddressTrustKeysNFTUpgradeable is a free data retrieval call binding the contract method 0x947beddb.
//
// Solidity: function addressTrustKeysNFTUpgradeable() view returns(address)
func (_Sale721 *Sale721CallerSession) AddressTrustKeysNFTUpgradeable() (common.Address, error) {
	return _Sale721.Contract.AddressTrustKeysNFTUpgradeable(&_Sale721.CallOpts)
}

// Currencies is a free data retrieval call binding the contract method 0x6036cba3.
//
// Solidity: function currencies(address ) view returns(bool)
func (_Sale721 *Sale721Caller) Currencies(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Sale721.contract.Call(opts, &out, "currencies", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Currencies is a free data retrieval call binding the contract method 0x6036cba3.
//
// Solidity: function currencies(address ) view returns(bool)
func (_Sale721 *Sale721Session) Currencies(arg0 common.Address) (bool, error) {
	return _Sale721.Contract.Currencies(&_Sale721.CallOpts, arg0)
}

// Currencies is a free data retrieval call binding the contract method 0x6036cba3.
//
// Solidity: function currencies(address ) view returns(bool)
func (_Sale721 *Sale721CallerSession) Currencies(arg0 common.Address) (bool, error) {
	return _Sale721.Contract.Currencies(&_Sale721.CallOpts, arg0)
}

// FeeBuy is a free data retrieval call binding the contract method 0xc5b0fe9d.
//
// Solidity: function feeBuy() view returns(uint32)
func (_Sale721 *Sale721Caller) FeeBuy(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Sale721.contract.Call(opts, &out, "feeBuy")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// FeeBuy is a free data retrieval call binding the contract method 0xc5b0fe9d.
//
// Solidity: function feeBuy() view returns(uint32)
func (_Sale721 *Sale721Session) FeeBuy() (uint32, error) {
	return _Sale721.Contract.FeeBuy(&_Sale721.CallOpts)
}

// FeeBuy is a free data retrieval call binding the contract method 0xc5b0fe9d.
//
// Solidity: function feeBuy() view returns(uint32)
func (_Sale721 *Sale721CallerSession) FeeBuy() (uint32, error) {
	return _Sale721.Contract.FeeBuy(&_Sale721.CallOpts)
}

// FeeDecimal is a free data retrieval call binding the contract method 0x7e11b31f.
//
// Solidity: function feeDecimal() view returns(uint32)
func (_Sale721 *Sale721Caller) FeeDecimal(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Sale721.contract.Call(opts, &out, "feeDecimal")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// FeeDecimal is a free data retrieval call binding the contract method 0x7e11b31f.
//
// Solidity: function feeDecimal() view returns(uint32)
func (_Sale721 *Sale721Session) FeeDecimal() (uint32, error) {
	return _Sale721.Contract.FeeDecimal(&_Sale721.CallOpts)
}

// FeeDecimal is a free data retrieval call binding the contract method 0x7e11b31f.
//
// Solidity: function feeDecimal() view returns(uint32)
func (_Sale721 *Sale721CallerSession) FeeDecimal() (uint32, error) {
	return _Sale721.Contract.FeeDecimal(&_Sale721.CallOpts)
}

// FeeSale is a free data retrieval call binding the contract method 0xe9c41dd1.
//
// Solidity: function feeSale() view returns(uint32)
func (_Sale721 *Sale721Caller) FeeSale(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Sale721.contract.Call(opts, &out, "feeSale")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// FeeSale is a free data retrieval call binding the contract method 0xe9c41dd1.
//
// Solidity: function feeSale() view returns(uint32)
func (_Sale721 *Sale721Session) FeeSale() (uint32, error) {
	return _Sale721.Contract.FeeSale(&_Sale721.CallOpts)
}

// FeeSale is a free data retrieval call binding the contract method 0xe9c41dd1.
//
// Solidity: function feeSale() view returns(uint32)
func (_Sale721 *Sale721CallerSession) FeeSale() (uint32, error) {
	return _Sale721.Contract.FeeSale(&_Sale721.CallOpts)
}

// LimitCurrency is a free data retrieval call binding the contract method 0x89e167e8.
//
// Solidity: function limitCurrency() view returns(bool)
func (_Sale721 *Sale721Caller) LimitCurrency(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Sale721.contract.Call(opts, &out, "limitCurrency")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// LimitCurrency is a free data retrieval call binding the contract method 0x89e167e8.
//
// Solidity: function limitCurrency() view returns(bool)
func (_Sale721 *Sale721Session) LimitCurrency() (bool, error) {
	return _Sale721.Contract.LimitCurrency(&_Sale721.CallOpts)
}

// LimitCurrency is a free data retrieval call binding the contract method 0x89e167e8.
//
// Solidity: function limitCurrency() view returns(bool)
func (_Sale721 *Sale721CallerSession) LimitCurrency() (bool, error) {
	return _Sale721.Contract.LimitCurrency(&_Sale721.CallOpts)
}

// Marketplace is a free data retrieval call binding the contract method 0xcc0ef311.
//
// Solidity: function marketplace(address , uint256 ) view returns(address owner, uint256 price, address contractERC20, bool exist)
func (_Sale721 *Sale721Caller) Marketplace(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	Owner         common.Address
	Price         *big.Int
	ContractERC20 common.Address
	Exist         bool
}, error) {
	var out []interface{}
	err := _Sale721.contract.Call(opts, &out, "marketplace", arg0, arg1)

	outstruct := new(struct {
		Owner         common.Address
		Price         *big.Int
		ContractERC20 common.Address
		Exist         bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Owner = out[0].(common.Address)
	outstruct.Price = out[1].(*big.Int)
	outstruct.ContractERC20 = out[2].(common.Address)
	outstruct.Exist = out[3].(bool)

	return *outstruct, err

}

// Marketplace is a free data retrieval call binding the contract method 0xcc0ef311.
//
// Solidity: function marketplace(address , uint256 ) view returns(address owner, uint256 price, address contractERC20, bool exist)
func (_Sale721 *Sale721Session) Marketplace(arg0 common.Address, arg1 *big.Int) (struct {
	Owner         common.Address
	Price         *big.Int
	ContractERC20 common.Address
	Exist         bool
}, error) {
	return _Sale721.Contract.Marketplace(&_Sale721.CallOpts, arg0, arg1)
}

// Marketplace is a free data retrieval call binding the contract method 0xcc0ef311.
//
// Solidity: function marketplace(address , uint256 ) view returns(address owner, uint256 price, address contractERC20, bool exist)
func (_Sale721 *Sale721CallerSession) Marketplace(arg0 common.Address, arg1 *big.Int) (struct {
	Owner         common.Address
	Price         *big.Int
	ContractERC20 common.Address
	Exist         bool
}, error) {
	return _Sale721.Contract.Marketplace(&_Sale721.CallOpts, arg0, arg1)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Sale721 *Sale721Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Sale721.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Sale721 *Sale721Session) Owner() (common.Address, error) {
	return _Sale721.Contract.Owner(&_Sale721.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Sale721 *Sale721CallerSession) Owner() (common.Address, error) {
	return _Sale721.Contract.Owner(&_Sale721.CallOpts)
}

// Royalties is a free data retrieval call binding the contract method 0xf053dc5c.
//
// Solidity: function royalties() view returns(uint32)
func (_Sale721 *Sale721Caller) Royalties(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Sale721.contract.Call(opts, &out, "royalties")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// Royalties is a free data retrieval call binding the contract method 0xf053dc5c.
//
// Solidity: function royalties() view returns(uint32)
func (_Sale721 *Sale721Session) Royalties() (uint32, error) {
	return _Sale721.Contract.Royalties(&_Sale721.CallOpts)
}

// Royalties is a free data retrieval call binding the contract method 0xf053dc5c.
//
// Solidity: function royalties() view returns(uint32)
func (_Sale721 *Sale721CallerSession) Royalties() (uint32, error) {
	return _Sale721.Contract.Royalties(&_Sale721.CallOpts)
}

// BuyToken is a paid mutator transaction binding the contract method 0xbbb0d9c2.
//
// Solidity: function BuyToken(uint256 _tokenId, address _contractNFT) payable returns()
func (_Sale721 *Sale721Transactor) BuyToken(opts *bind.TransactOpts, _tokenId *big.Int, _contractNFT common.Address) (*types.Transaction, error) {
	return _Sale721.contract.Transact(opts, "BuyToken", _tokenId, _contractNFT)
}

// BuyToken is a paid mutator transaction binding the contract method 0xbbb0d9c2.
//
// Solidity: function BuyToken(uint256 _tokenId, address _contractNFT) payable returns()
func (_Sale721 *Sale721Session) BuyToken(_tokenId *big.Int, _contractNFT common.Address) (*types.Transaction, error) {
	return _Sale721.Contract.BuyToken(&_Sale721.TransactOpts, _tokenId, _contractNFT)
}

// BuyToken is a paid mutator transaction binding the contract method 0xbbb0d9c2.
//
// Solidity: function BuyToken(uint256 _tokenId, address _contractNFT) payable returns()
func (_Sale721 *Sale721TransactorSession) BuyToken(_tokenId *big.Int, _contractNFT common.Address) (*types.Transaction, error) {
	return _Sale721.Contract.BuyToken(&_Sale721.TransactOpts, _tokenId, _contractNFT)
}

// CancelSaleToken is a paid mutator transaction binding the contract method 0x1056258c.
//
// Solidity: function CancelSaleToken(uint256 _tokenId, address _contractNFT) returns()
func (_Sale721 *Sale721Transactor) CancelSaleToken(opts *bind.TransactOpts, _tokenId *big.Int, _contractNFT common.Address) (*types.Transaction, error) {
	return _Sale721.contract.Transact(opts, "CancelSaleToken", _tokenId, _contractNFT)
}

// CancelSaleToken is a paid mutator transaction binding the contract method 0x1056258c.
//
// Solidity: function CancelSaleToken(uint256 _tokenId, address _contractNFT) returns()
func (_Sale721 *Sale721Session) CancelSaleToken(_tokenId *big.Int, _contractNFT common.Address) (*types.Transaction, error) {
	return _Sale721.Contract.CancelSaleToken(&_Sale721.TransactOpts, _tokenId, _contractNFT)
}

// CancelSaleToken is a paid mutator transaction binding the contract method 0x1056258c.
//
// Solidity: function CancelSaleToken(uint256 _tokenId, address _contractNFT) returns()
func (_Sale721 *Sale721TransactorSession) CancelSaleToken(_tokenId *big.Int, _contractNFT common.Address) (*types.Transaction, error) {
	return _Sale721.Contract.CancelSaleToken(&_Sale721.TransactOpts, _tokenId, _contractNFT)
}

// SaleToken is a paid mutator transaction binding the contract method 0x3b2f4114.
//
// Solidity: function SaleToken(uint256 _tokenId, address _contractNFT, uint256 _price, address _contractERC20) returns()
func (_Sale721 *Sale721Transactor) SaleToken(opts *bind.TransactOpts, _tokenId *big.Int, _contractNFT common.Address, _price *big.Int, _contractERC20 common.Address) (*types.Transaction, error) {
	return _Sale721.contract.Transact(opts, "SaleToken", _tokenId, _contractNFT, _price, _contractERC20)
}

// SaleToken is a paid mutator transaction binding the contract method 0x3b2f4114.
//
// Solidity: function SaleToken(uint256 _tokenId, address _contractNFT, uint256 _price, address _contractERC20) returns()
func (_Sale721 *Sale721Session) SaleToken(_tokenId *big.Int, _contractNFT common.Address, _price *big.Int, _contractERC20 common.Address) (*types.Transaction, error) {
	return _Sale721.Contract.SaleToken(&_Sale721.TransactOpts, _tokenId, _contractNFT, _price, _contractERC20)
}

// SaleToken is a paid mutator transaction binding the contract method 0x3b2f4114.
//
// Solidity: function SaleToken(uint256 _tokenId, address _contractNFT, uint256 _price, address _contractERC20) returns()
func (_Sale721 *Sale721TransactorSession) SaleToken(_tokenId *big.Int, _contractNFT common.Address, _price *big.Int, _contractERC20 common.Address) (*types.Transaction, error) {
	return _Sale721.Contract.SaleToken(&_Sale721.TransactOpts, _tokenId, _contractNFT, _price, _contractERC20)
}

// UpdatePrice is a paid mutator transaction binding the contract method 0x4c9f1072.
//
// Solidity: function UpdatePrice(uint256 _tokenId, address _contractNFT, uint256 _price, address _contractERC20) returns()
func (_Sale721 *Sale721Transactor) UpdatePrice(opts *bind.TransactOpts, _tokenId *big.Int, _contractNFT common.Address, _price *big.Int, _contractERC20 common.Address) (*types.Transaction, error) {
	return _Sale721.contract.Transact(opts, "UpdatePrice", _tokenId, _contractNFT, _price, _contractERC20)
}

// UpdatePrice is a paid mutator transaction binding the contract method 0x4c9f1072.
//
// Solidity: function UpdatePrice(uint256 _tokenId, address _contractNFT, uint256 _price, address _contractERC20) returns()
func (_Sale721 *Sale721Session) UpdatePrice(_tokenId *big.Int, _contractNFT common.Address, _price *big.Int, _contractERC20 common.Address) (*types.Transaction, error) {
	return _Sale721.Contract.UpdatePrice(&_Sale721.TransactOpts, _tokenId, _contractNFT, _price, _contractERC20)
}

// UpdatePrice is a paid mutator transaction binding the contract method 0x4c9f1072.
//
// Solidity: function UpdatePrice(uint256 _tokenId, address _contractNFT, uint256 _price, address _contractERC20) returns()
func (_Sale721 *Sale721TransactorSession) UpdatePrice(_tokenId *big.Int, _contractNFT common.Address, _price *big.Int, _contractERC20 common.Address) (*types.Transaction, error) {
	return _Sale721.Contract.UpdatePrice(&_Sale721.TransactOpts, _tokenId, _contractNFT, _price, _contractERC20)
}

// TrustKeysMarketplaceInit is a paid mutator transaction binding the contract method 0x8fbe91d5.
//
// Solidity: function __TrustKeysMarketplace_init_() returns()
func (_Sale721 *Sale721Transactor) TrustKeysMarketplaceInit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sale721.contract.Transact(opts, "__TrustKeysMarketplace_init_")
}

// TrustKeysMarketplaceInit is a paid mutator transaction binding the contract method 0x8fbe91d5.
//
// Solidity: function __TrustKeysMarketplace_init_() returns()
func (_Sale721 *Sale721Session) TrustKeysMarketplaceInit() (*types.Transaction, error) {
	return _Sale721.Contract.TrustKeysMarketplaceInit(&_Sale721.TransactOpts)
}

// TrustKeysMarketplaceInit is a paid mutator transaction binding the contract method 0x8fbe91d5.
//
// Solidity: function __TrustKeysMarketplace_init_() returns()
func (_Sale721 *Sale721TransactorSession) TrustKeysMarketplaceInit() (*types.Transaction, error) {
	return _Sale721.Contract.TrustKeysMarketplaceInit(&_Sale721.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Sale721 *Sale721Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sale721.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Sale721 *Sale721Session) RenounceOwnership() (*types.Transaction, error) {
	return _Sale721.Contract.RenounceOwnership(&_Sale721.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Sale721 *Sale721TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Sale721.Contract.RenounceOwnership(&_Sale721.TransactOpts)
}

// SetAddressTrustKeysNFT is a paid mutator transaction binding the contract method 0xbdb4af85.
//
// Solidity: function setAddressTrustKeysNFT(address _addressTrustKeysNFT) returns()
func (_Sale721 *Sale721Transactor) SetAddressTrustKeysNFT(opts *bind.TransactOpts, _addressTrustKeysNFT common.Address) (*types.Transaction, error) {
	return _Sale721.contract.Transact(opts, "setAddressTrustKeysNFT", _addressTrustKeysNFT)
}

// SetAddressTrustKeysNFT is a paid mutator transaction binding the contract method 0xbdb4af85.
//
// Solidity: function setAddressTrustKeysNFT(address _addressTrustKeysNFT) returns()
func (_Sale721 *Sale721Session) SetAddressTrustKeysNFT(_addressTrustKeysNFT common.Address) (*types.Transaction, error) {
	return _Sale721.Contract.SetAddressTrustKeysNFT(&_Sale721.TransactOpts, _addressTrustKeysNFT)
}

// SetAddressTrustKeysNFT is a paid mutator transaction binding the contract method 0xbdb4af85.
//
// Solidity: function setAddressTrustKeysNFT(address _addressTrustKeysNFT) returns()
func (_Sale721 *Sale721TransactorSession) SetAddressTrustKeysNFT(_addressTrustKeysNFT common.Address) (*types.Transaction, error) {
	return _Sale721.Contract.SetAddressTrustKeysNFT(&_Sale721.TransactOpts, _addressTrustKeysNFT)
}

// SetCurrencies is a paid mutator transaction binding the contract method 0x8d34f97f.
//
// Solidity: function setCurrencies(address[] _listContractERC20) returns()
func (_Sale721 *Sale721Transactor) SetCurrencies(opts *bind.TransactOpts, _listContractERC20 []common.Address) (*types.Transaction, error) {
	return _Sale721.contract.Transact(opts, "setCurrencies", _listContractERC20)
}

// SetCurrencies is a paid mutator transaction binding the contract method 0x8d34f97f.
//
// Solidity: function setCurrencies(address[] _listContractERC20) returns()
func (_Sale721 *Sale721Session) SetCurrencies(_listContractERC20 []common.Address) (*types.Transaction, error) {
	return _Sale721.Contract.SetCurrencies(&_Sale721.TransactOpts, _listContractERC20)
}

// SetCurrencies is a paid mutator transaction binding the contract method 0x8d34f97f.
//
// Solidity: function setCurrencies(address[] _listContractERC20) returns()
func (_Sale721 *Sale721TransactorSession) SetCurrencies(_listContractERC20 []common.Address) (*types.Transaction, error) {
	return _Sale721.Contract.SetCurrencies(&_Sale721.TransactOpts, _listContractERC20)
}

// SetCurrency is a paid mutator transaction binding the contract method 0x2f84c391.
//
// Solidity: function setCurrency(address _contractERC20) returns()
func (_Sale721 *Sale721Transactor) SetCurrency(opts *bind.TransactOpts, _contractERC20 common.Address) (*types.Transaction, error) {
	return _Sale721.contract.Transact(opts, "setCurrency", _contractERC20)
}

// SetCurrency is a paid mutator transaction binding the contract method 0x2f84c391.
//
// Solidity: function setCurrency(address _contractERC20) returns()
func (_Sale721 *Sale721Session) SetCurrency(_contractERC20 common.Address) (*types.Transaction, error) {
	return _Sale721.Contract.SetCurrency(&_Sale721.TransactOpts, _contractERC20)
}

// SetCurrency is a paid mutator transaction binding the contract method 0x2f84c391.
//
// Solidity: function setCurrency(address _contractERC20) returns()
func (_Sale721 *Sale721TransactorSession) SetCurrency(_contractERC20 common.Address) (*types.Transaction, error) {
	return _Sale721.Contract.SetCurrency(&_Sale721.TransactOpts, _contractERC20)
}

// SetLimitCurrency is a paid mutator transaction binding the contract method 0x7ea90cd6.
//
// Solidity: function setLimitCurrency(bool _limitCurrency) returns()
func (_Sale721 *Sale721Transactor) SetLimitCurrency(opts *bind.TransactOpts, _limitCurrency bool) (*types.Transaction, error) {
	return _Sale721.contract.Transact(opts, "setLimitCurrency", _limitCurrency)
}

// SetLimitCurrency is a paid mutator transaction binding the contract method 0x7ea90cd6.
//
// Solidity: function setLimitCurrency(bool _limitCurrency) returns()
func (_Sale721 *Sale721Session) SetLimitCurrency(_limitCurrency bool) (*types.Transaction, error) {
	return _Sale721.Contract.SetLimitCurrency(&_Sale721.TransactOpts, _limitCurrency)
}

// SetLimitCurrency is a paid mutator transaction binding the contract method 0x7ea90cd6.
//
// Solidity: function setLimitCurrency(bool _limitCurrency) returns()
func (_Sale721 *Sale721TransactorSession) SetLimitCurrency(_limitCurrency bool) (*types.Transaction, error) {
	return _Sale721.Contract.SetLimitCurrency(&_Sale721.TransactOpts, _limitCurrency)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Sale721 *Sale721Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Sale721.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Sale721 *Sale721Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Sale721.Contract.TransferOwnership(&_Sale721.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Sale721 *Sale721TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Sale721.Contract.TransferOwnership(&_Sale721.TransactOpts, newOwner)
}

// UpdateFees is a paid mutator transaction binding the contract method 0x92d0da99.
//
// Solidity: function updateFees(uint32 _royalties, uint32 _feeBuy, uint32 _feeSale) returns()
func (_Sale721 *Sale721Transactor) UpdateFees(opts *bind.TransactOpts, _royalties uint32, _feeBuy uint32, _feeSale uint32) (*types.Transaction, error) {
	return _Sale721.contract.Transact(opts, "updateFees", _royalties, _feeBuy, _feeSale)
}

// UpdateFees is a paid mutator transaction binding the contract method 0x92d0da99.
//
// Solidity: function updateFees(uint32 _royalties, uint32 _feeBuy, uint32 _feeSale) returns()
func (_Sale721 *Sale721Session) UpdateFees(_royalties uint32, _feeBuy uint32, _feeSale uint32) (*types.Transaction, error) {
	return _Sale721.Contract.UpdateFees(&_Sale721.TransactOpts, _royalties, _feeBuy, _feeSale)
}

// UpdateFees is a paid mutator transaction binding the contract method 0x92d0da99.
//
// Solidity: function updateFees(uint32 _royalties, uint32 _feeBuy, uint32 _feeSale) returns()
func (_Sale721 *Sale721TransactorSession) UpdateFees(_royalties uint32, _feeBuy uint32, _feeSale uint32) (*types.Transaction, error) {
	return _Sale721.Contract.UpdateFees(&_Sale721.TransactOpts, _royalties, _feeBuy, _feeSale)
}

// Sale721CanceledIterator is returned from FilterCanceled and is used to iterate over the raw logs and unpacked data for Canceled events raised by the Sale721 contract.
type Sale721CanceledIterator struct {
	Event *Sale721Canceled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Sale721CanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Sale721Canceled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Sale721Canceled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Sale721CanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Sale721CanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Sale721Canceled represents a Canceled event raised by the Sale721 contract.
type Sale721Canceled struct {
	Seller      common.Address
	TokenId     *big.Int
	ContractNFT common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterCanceled is a free log retrieval operation binding the contract event 0x09d3dbe43fafc7dcb2770048ced91d2880093a47f9342e248e5db274b24f330b.
//
// Solidity: event Canceled(address seller, uint256 tokenId, address contractNFT)
func (_Sale721 *Sale721Filterer) FilterCanceled(opts *bind.FilterOpts) (*Sale721CanceledIterator, error) {

	logs, sub, err := _Sale721.contract.FilterLogs(opts, "Canceled")
	if err != nil {
		return nil, err
	}
	return &Sale721CanceledIterator{contract: _Sale721.contract, event: "Canceled", logs: logs, sub: sub}, nil
}

// WatchCanceled is a free log subscription operation binding the contract event 0x09d3dbe43fafc7dcb2770048ced91d2880093a47f9342e248e5db274b24f330b.
//
// Solidity: event Canceled(address seller, uint256 tokenId, address contractNFT)
func (_Sale721 *Sale721Filterer) WatchCanceled(opts *bind.WatchOpts, sink chan<- *Sale721Canceled) (event.Subscription, error) {

	logs, sub, err := _Sale721.contract.WatchLogs(opts, "Canceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Sale721Canceled)
				if err := _Sale721.contract.UnpackLog(event, "Canceled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCanceled is a log parse operation binding the contract event 0x09d3dbe43fafc7dcb2770048ced91d2880093a47f9342e248e5db274b24f330b.
//
// Solidity: event Canceled(address seller, uint256 tokenId, address contractNFT)
func (_Sale721 *Sale721Filterer) ParseCanceled(log types.Log) (*Sale721Canceled, error) {
	event := new(Sale721Canceled)
	if err := _Sale721.contract.UnpackLog(event, "Canceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Sale721ListingPriceIterator is returned from FilterListingPrice and is used to iterate over the raw logs and unpacked data for ListingPrice events raised by the Sale721 contract.
type Sale721ListingPriceIterator struct {
	Event *Sale721ListingPrice // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Sale721ListingPriceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Sale721ListingPrice)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Sale721ListingPrice)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Sale721ListingPriceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Sale721ListingPriceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Sale721ListingPrice represents a ListingPrice event raised by the Sale721 contract.
type Sale721ListingPrice struct {
	Seller      common.Address
	TokenId     *big.Int
	ContractNFT common.Address
	SaleInfo    TrustKeysMarketplaceSaleInfo
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterListingPrice is a free log retrieval operation binding the contract event 0xd8b26715f92245dca1b93c2ebf8820a11688e0d81fd8070389108bd8717ab144.
//
// Solidity: event ListingPrice(address seller, uint256 tokenId, address contractNFT, (address,uint256,address,bool) saleInfo)
func (_Sale721 *Sale721Filterer) FilterListingPrice(opts *bind.FilterOpts) (*Sale721ListingPriceIterator, error) {

	logs, sub, err := _Sale721.contract.FilterLogs(opts, "ListingPrice")
	if err != nil {
		return nil, err
	}
	return &Sale721ListingPriceIterator{contract: _Sale721.contract, event: "ListingPrice", logs: logs, sub: sub}, nil
}

// WatchListingPrice is a free log subscription operation binding the contract event 0xd8b26715f92245dca1b93c2ebf8820a11688e0d81fd8070389108bd8717ab144.
//
// Solidity: event ListingPrice(address seller, uint256 tokenId, address contractNFT, (address,uint256,address,bool) saleInfo)
func (_Sale721 *Sale721Filterer) WatchListingPrice(opts *bind.WatchOpts, sink chan<- *Sale721ListingPrice) (event.Subscription, error) {

	logs, sub, err := _Sale721.contract.WatchLogs(opts, "ListingPrice")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Sale721ListingPrice)
				if err := _Sale721.contract.UnpackLog(event, "ListingPrice", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseListingPrice is a log parse operation binding the contract event 0xd8b26715f92245dca1b93c2ebf8820a11688e0d81fd8070389108bd8717ab144.
//
// Solidity: event ListingPrice(address seller, uint256 tokenId, address contractNFT, (address,uint256,address,bool) saleInfo)
func (_Sale721 *Sale721Filterer) ParseListingPrice(log types.Log) (*Sale721ListingPrice, error) {
	event := new(Sale721ListingPrice)
	if err := _Sale721.contract.UnpackLog(event, "ListingPrice", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Sale721OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Sale721 contract.
type Sale721OwnershipTransferredIterator struct {
	Event *Sale721OwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Sale721OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Sale721OwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Sale721OwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Sale721OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Sale721OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Sale721OwnershipTransferred represents a OwnershipTransferred event raised by the Sale721 contract.
type Sale721OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Sale721 *Sale721Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*Sale721OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Sale721.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &Sale721OwnershipTransferredIterator{contract: _Sale721.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Sale721 *Sale721Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *Sale721OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Sale721.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Sale721OwnershipTransferred)
				if err := _Sale721.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Sale721 *Sale721Filterer) ParseOwnershipTransferred(log types.Log) (*Sale721OwnershipTransferred, error) {
	event := new(Sale721OwnershipTransferred)
	if err := _Sale721.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Sale721SoldIterator is returned from FilterSold and is used to iterate over the raw logs and unpacked data for Sold events raised by the Sale721 contract.
type Sale721SoldIterator struct {
	Event *Sale721Sold // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Sale721SoldIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Sale721Sold)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Sale721Sold)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Sale721SoldIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Sale721SoldIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Sale721Sold represents a Sold event raised by the Sale721 contract.
type Sale721Sold struct {
	Buyer       common.Address
	TokenId     *big.Int
	ContractNFT common.Address
	SaleInfo    TrustKeysMarketplaceSaleInfo
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSold is a free log retrieval operation binding the contract event 0x9fc333e5ae724c1f60c1c2b270767a0ba02a0c30ac6811729735cbd6e2e41f3e.
//
// Solidity: event Sold(address buyer, uint256 tokenId, address contractNFT, (address,uint256,address,bool) saleInfo)
func (_Sale721 *Sale721Filterer) FilterSold(opts *bind.FilterOpts) (*Sale721SoldIterator, error) {

	logs, sub, err := _Sale721.contract.FilterLogs(opts, "Sold")
	if err != nil {
		return nil, err
	}
	return &Sale721SoldIterator{contract: _Sale721.contract, event: "Sold", logs: logs, sub: sub}, nil
}

// WatchSold is a free log subscription operation binding the contract event 0x9fc333e5ae724c1f60c1c2b270767a0ba02a0c30ac6811729735cbd6e2e41f3e.
//
// Solidity: event Sold(address buyer, uint256 tokenId, address contractNFT, (address,uint256,address,bool) saleInfo)
func (_Sale721 *Sale721Filterer) WatchSold(opts *bind.WatchOpts, sink chan<- *Sale721Sold) (event.Subscription, error) {

	logs, sub, err := _Sale721.contract.WatchLogs(opts, "Sold")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Sale721Sold)
				if err := _Sale721.contract.UnpackLog(event, "Sold", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSold is a log parse operation binding the contract event 0x9fc333e5ae724c1f60c1c2b270767a0ba02a0c30ac6811729735cbd6e2e41f3e.
//
// Solidity: event Sold(address buyer, uint256 tokenId, address contractNFT, (address,uint256,address,bool) saleInfo)
func (_Sale721 *Sale721Filterer) ParseSold(log types.Log) (*Sale721Sold, error) {
	event := new(Sale721Sold)
	if err := _Sale721.contract.UnpackLog(event, "Sold", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Sale721UpdatedIterator is returned from FilterUpdated and is used to iterate over the raw logs and unpacked data for Updated events raised by the Sale721 contract.
type Sale721UpdatedIterator struct {
	Event *Sale721Updated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Sale721UpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Sale721Updated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Sale721Updated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Sale721UpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Sale721UpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Sale721Updated represents a Updated event raised by the Sale721 contract.
type Sale721Updated struct {
	Seller      common.Address
	TokenId     *big.Int
	ContractNFT common.Address
	SaleInfo    TrustKeysMarketplaceSaleInfo
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUpdated is a free log retrieval operation binding the contract event 0xf4fb2f09d10240c86197fcb3020353cbb8dff4545d4d43658a05876dcabbc77b.
//
// Solidity: event Updated(address seller, uint256 tokenId, address contractNFT, (address,uint256,address,bool) saleInfo)
func (_Sale721 *Sale721Filterer) FilterUpdated(opts *bind.FilterOpts) (*Sale721UpdatedIterator, error) {

	logs, sub, err := _Sale721.contract.FilterLogs(opts, "Updated")
	if err != nil {
		return nil, err
	}
	return &Sale721UpdatedIterator{contract: _Sale721.contract, event: "Updated", logs: logs, sub: sub}, nil
}

// WatchUpdated is a free log subscription operation binding the contract event 0xf4fb2f09d10240c86197fcb3020353cbb8dff4545d4d43658a05876dcabbc77b.
//
// Solidity: event Updated(address seller, uint256 tokenId, address contractNFT, (address,uint256,address,bool) saleInfo)
func (_Sale721 *Sale721Filterer) WatchUpdated(opts *bind.WatchOpts, sink chan<- *Sale721Updated) (event.Subscription, error) {

	logs, sub, err := _Sale721.contract.WatchLogs(opts, "Updated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Sale721Updated)
				if err := _Sale721.contract.UnpackLog(event, "Updated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdated is a log parse operation binding the contract event 0xf4fb2f09d10240c86197fcb3020353cbb8dff4545d4d43658a05876dcabbc77b.
//
// Solidity: event Updated(address seller, uint256 tokenId, address contractNFT, (address,uint256,address,bool) saleInfo)
func (_Sale721 *Sale721Filterer) ParseUpdated(log types.Log) (*Sale721Updated, error) {
	event := new(Sale721Updated)
	if err := _Sale721.contract.UnpackLog(event, "Updated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
