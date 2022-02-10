// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package badges

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

// BadgesABI is the input ABI used to generate the binding from.
const BadgesABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"MintAndApproveMarket\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"__FactoryTrustKeysBadges_init_\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"addressTrustKeysAuction\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"addressTrustKeysMarketplace\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ownerAddress\",\"type\":\"address\"}],\"name\":\"getItemsByOwner\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"mintAndApproveMarket\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_total\",\"type\":\"uint256\"}],\"name\":\"mintBatchAndApproveMarket\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"mintToAndApproveMarket\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addressTrustKeysAuction\",\"type\":\"address\"}],\"name\":\"setTrustKeysAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addressTrustKeysMarketplace\",\"type\":\"address\"}],\"name\":\"setTrustKeysMarketplace\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Badges is an auto generated Go binding around an Ethereum contract.
type Badges struct {
	BadgesCaller     // Read-only binding to the contract
	BadgesTransactor // Write-only binding to the contract
	BadgesFilterer   // Log filterer for contract events
}

// BadgesCaller is an auto generated read-only Go binding around an Ethereum contract.
type BadgesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BadgesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BadgesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BadgesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BadgesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BadgesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BadgesSession struct {
	Contract     *Badges           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BadgesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BadgesCallerSession struct {
	Contract *BadgesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BadgesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BadgesTransactorSession struct {
	Contract     *BadgesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BadgesRaw is an auto generated low-level Go binding around an Ethereum contract.
type BadgesRaw struct {
	Contract *Badges // Generic contract binding to access the raw methods on
}

// BadgesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BadgesCallerRaw struct {
	Contract *BadgesCaller // Generic read-only contract binding to access the raw methods on
}

// BadgesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BadgesTransactorRaw struct {
	Contract *BadgesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBadges creates a new instance of Badges, bound to a specific deployed contract.
func NewBadges(address common.Address, backend bind.ContractBackend) (*Badges, error) {
	contract, err := bindBadges(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Badges{BadgesCaller: BadgesCaller{contract: contract}, BadgesTransactor: BadgesTransactor{contract: contract}, BadgesFilterer: BadgesFilterer{contract: contract}}, nil
}

// NewBadgesCaller creates a new read-only instance of Badges, bound to a specific deployed contract.
func NewBadgesCaller(address common.Address, caller bind.ContractCaller) (*BadgesCaller, error) {
	contract, err := bindBadges(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BadgesCaller{contract: contract}, nil
}

// NewBadgesTransactor creates a new write-only instance of Badges, bound to a specific deployed contract.
func NewBadgesTransactor(address common.Address, transactor bind.ContractTransactor) (*BadgesTransactor, error) {
	contract, err := bindBadges(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BadgesTransactor{contract: contract}, nil
}

// NewBadgesFilterer creates a new log filterer instance of Badges, bound to a specific deployed contract.
func NewBadgesFilterer(address common.Address, filterer bind.ContractFilterer) (*BadgesFilterer, error) {
	contract, err := bindBadges(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BadgesFilterer{contract: contract}, nil
}

// bindBadges binds a generic wrapper to an already deployed contract.
func bindBadges(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BadgesABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Badges *BadgesRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Badges.Contract.BadgesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Badges *BadgesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Badges.Contract.BadgesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Badges *BadgesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Badges.Contract.BadgesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Badges *BadgesCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Badges.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Badges *BadgesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Badges.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Badges *BadgesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Badges.Contract.contract.Transact(opts, method, params...)
}

// AddressTrustKeysAuction is a free data retrieval call binding the contract method 0x9316144b.
//
// Solidity: function addressTrustKeysAuction() view returns(address)
func (_Badges *BadgesCaller) AddressTrustKeysAuction(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Badges.contract.Call(opts, &out, "addressTrustKeysAuction")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddressTrustKeysAuction is a free data retrieval call binding the contract method 0x9316144b.
//
// Solidity: function addressTrustKeysAuction() view returns(address)
func (_Badges *BadgesSession) AddressTrustKeysAuction() (common.Address, error) {
	return _Badges.Contract.AddressTrustKeysAuction(&_Badges.CallOpts)
}

// AddressTrustKeysAuction is a free data retrieval call binding the contract method 0x9316144b.
//
// Solidity: function addressTrustKeysAuction() view returns(address)
func (_Badges *BadgesCallerSession) AddressTrustKeysAuction() (common.Address, error) {
	return _Badges.Contract.AddressTrustKeysAuction(&_Badges.CallOpts)
}

// AddressTrustKeysMarketplace is a free data retrieval call binding the contract method 0xe58c3af4.
//
// Solidity: function addressTrustKeysMarketplace() view returns(address)
func (_Badges *BadgesCaller) AddressTrustKeysMarketplace(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Badges.contract.Call(opts, &out, "addressTrustKeysMarketplace")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddressTrustKeysMarketplace is a free data retrieval call binding the contract method 0xe58c3af4.
//
// Solidity: function addressTrustKeysMarketplace() view returns(address)
func (_Badges *BadgesSession) AddressTrustKeysMarketplace() (common.Address, error) {
	return _Badges.Contract.AddressTrustKeysMarketplace(&_Badges.CallOpts)
}

// AddressTrustKeysMarketplace is a free data retrieval call binding the contract method 0xe58c3af4.
//
// Solidity: function addressTrustKeysMarketplace() view returns(address)
func (_Badges *BadgesCallerSession) AddressTrustKeysMarketplace() (common.Address, error) {
	return _Badges.Contract.AddressTrustKeysMarketplace(&_Badges.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Badges *BadgesCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Badges.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Badges *BadgesSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Badges.Contract.BalanceOf(&_Badges.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Badges *BadgesCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Badges.Contract.BalanceOf(&_Badges.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Badges *BadgesCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Badges.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Badges *BadgesSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Badges.Contract.GetApproved(&_Badges.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Badges *BadgesCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Badges.Contract.GetApproved(&_Badges.CallOpts, tokenId)
}

// GetItemsByOwner is a free data retrieval call binding the contract method 0x2c67a8e5.
//
// Solidity: function getItemsByOwner(address ownerAddress) view returns(string[])
func (_Badges *BadgesCaller) GetItemsByOwner(opts *bind.CallOpts, ownerAddress common.Address) ([]string, error) {
	var out []interface{}
	err := _Badges.contract.Call(opts, &out, "getItemsByOwner", ownerAddress)

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// GetItemsByOwner is a free data retrieval call binding the contract method 0x2c67a8e5.
//
// Solidity: function getItemsByOwner(address ownerAddress) view returns(string[])
func (_Badges *BadgesSession) GetItemsByOwner(ownerAddress common.Address) ([]string, error) {
	return _Badges.Contract.GetItemsByOwner(&_Badges.CallOpts, ownerAddress)
}

// GetItemsByOwner is a free data retrieval call binding the contract method 0x2c67a8e5.
//
// Solidity: function getItemsByOwner(address ownerAddress) view returns(string[])
func (_Badges *BadgesCallerSession) GetItemsByOwner(ownerAddress common.Address) ([]string, error) {
	return _Badges.Contract.GetItemsByOwner(&_Badges.CallOpts, ownerAddress)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Badges *BadgesCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Badges.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Badges *BadgesSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Badges.Contract.IsApprovedForAll(&_Badges.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Badges *BadgesCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Badges.Contract.IsApprovedForAll(&_Badges.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Badges *BadgesCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Badges.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Badges *BadgesSession) Name() (string, error) {
	return _Badges.Contract.Name(&_Badges.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Badges *BadgesCallerSession) Name() (string, error) {
	return _Badges.Contract.Name(&_Badges.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Badges *BadgesCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Badges.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Badges *BadgesSession) Owner() (common.Address, error) {
	return _Badges.Contract.Owner(&_Badges.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Badges *BadgesCallerSession) Owner() (common.Address, error) {
	return _Badges.Contract.Owner(&_Badges.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Badges *BadgesCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Badges.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Badges *BadgesSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Badges.Contract.OwnerOf(&_Badges.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Badges *BadgesCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Badges.Contract.OwnerOf(&_Badges.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Badges *BadgesCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Badges.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Badges *BadgesSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Badges.Contract.SupportsInterface(&_Badges.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Badges *BadgesCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Badges.Contract.SupportsInterface(&_Badges.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Badges *BadgesCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Badges.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Badges *BadgesSession) Symbol() (string, error) {
	return _Badges.Contract.Symbol(&_Badges.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Badges *BadgesCallerSession) Symbol() (string, error) {
	return _Badges.Contract.Symbol(&_Badges.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Badges *BadgesCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Badges.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Badges *BadgesSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Badges.Contract.TokenByIndex(&_Badges.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Badges *BadgesCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Badges.Contract.TokenByIndex(&_Badges.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Badges *BadgesCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Badges.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Badges *BadgesSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Badges.Contract.TokenOfOwnerByIndex(&_Badges.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Badges *BadgesCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Badges.Contract.TokenOfOwnerByIndex(&_Badges.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Badges *BadgesCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Badges.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Badges *BadgesSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Badges.Contract.TokenURI(&_Badges.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Badges *BadgesCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Badges.Contract.TokenURI(&_Badges.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Badges *BadgesCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Badges.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Badges *BadgesSession) TotalSupply() (*big.Int, error) {
	return _Badges.Contract.TotalSupply(&_Badges.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Badges *BadgesCallerSession) TotalSupply() (*big.Int, error) {
	return _Badges.Contract.TotalSupply(&_Badges.CallOpts)
}

// FactoryTrustKeysBadgesInit is a paid mutator transaction binding the contract method 0x56d189ab.
//
// Solidity: function __FactoryTrustKeysBadges_init_() returns()
func (_Badges *BadgesTransactor) FactoryTrustKeysBadgesInit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Badges.contract.Transact(opts, "__FactoryTrustKeysBadges_init_")
}

// FactoryTrustKeysBadgesInit is a paid mutator transaction binding the contract method 0x56d189ab.
//
// Solidity: function __FactoryTrustKeysBadges_init_() returns()
func (_Badges *BadgesSession) FactoryTrustKeysBadgesInit() (*types.Transaction, error) {
	return _Badges.Contract.FactoryTrustKeysBadgesInit(&_Badges.TransactOpts)
}

// FactoryTrustKeysBadgesInit is a paid mutator transaction binding the contract method 0x56d189ab.
//
// Solidity: function __FactoryTrustKeysBadges_init_() returns()
func (_Badges *BadgesTransactorSession) FactoryTrustKeysBadgesInit() (*types.Transaction, error) {
	return _Badges.Contract.FactoryTrustKeysBadgesInit(&_Badges.TransactOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Badges *BadgesTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Badges.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Badges *BadgesSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Badges.Contract.Approve(&_Badges.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Badges *BadgesTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Badges.Contract.Approve(&_Badges.TransactOpts, to, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 _tokenId) returns()
func (_Badges *BadgesTransactor) Burn(opts *bind.TransactOpts, _tokenId *big.Int) (*types.Transaction, error) {
	return _Badges.contract.Transact(opts, "burn", _tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 _tokenId) returns()
func (_Badges *BadgesSession) Burn(_tokenId *big.Int) (*types.Transaction, error) {
	return _Badges.Contract.Burn(&_Badges.TransactOpts, _tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 _tokenId) returns()
func (_Badges *BadgesTransactorSession) Burn(_tokenId *big.Int) (*types.Transaction, error) {
	return _Badges.Contract.Burn(&_Badges.TransactOpts, _tokenId)
}

// MintAndApproveMarket is a paid mutator transaction binding the contract method 0x5174e853.
//
// Solidity: function mintAndApproveMarket(string uri) returns()
func (_Badges *BadgesTransactor) MintAndApproveMarket(opts *bind.TransactOpts, uri string) (*types.Transaction, error) {
	return _Badges.contract.Transact(opts, "mintAndApproveMarket", uri)
}

// MintAndApproveMarket is a paid mutator transaction binding the contract method 0x5174e853.
//
// Solidity: function mintAndApproveMarket(string uri) returns()
func (_Badges *BadgesSession) MintAndApproveMarket(uri string) (*types.Transaction, error) {
	return _Badges.Contract.MintAndApproveMarket(&_Badges.TransactOpts, uri)
}

// MintAndApproveMarket is a paid mutator transaction binding the contract method 0x5174e853.
//
// Solidity: function mintAndApproveMarket(string uri) returns()
func (_Badges *BadgesTransactorSession) MintAndApproveMarket(uri string) (*types.Transaction, error) {
	return _Badges.Contract.MintAndApproveMarket(&_Badges.TransactOpts, uri)
}

// MintBatchAndApproveMarket is a paid mutator transaction binding the contract method 0xa94feba3.
//
// Solidity: function mintBatchAndApproveMarket(string uri, uint256 _total) returns()
func (_Badges *BadgesTransactor) MintBatchAndApproveMarket(opts *bind.TransactOpts, uri string, _total *big.Int) (*types.Transaction, error) {
	return _Badges.contract.Transact(opts, "mintBatchAndApproveMarket", uri, _total)
}

// MintBatchAndApproveMarket is a paid mutator transaction binding the contract method 0xa94feba3.
//
// Solidity: function mintBatchAndApproveMarket(string uri, uint256 _total) returns()
func (_Badges *BadgesSession) MintBatchAndApproveMarket(uri string, _total *big.Int) (*types.Transaction, error) {
	return _Badges.Contract.MintBatchAndApproveMarket(&_Badges.TransactOpts, uri, _total)
}

// MintBatchAndApproveMarket is a paid mutator transaction binding the contract method 0xa94feba3.
//
// Solidity: function mintBatchAndApproveMarket(string uri, uint256 _total) returns()
func (_Badges *BadgesTransactorSession) MintBatchAndApproveMarket(uri string, _total *big.Int) (*types.Transaction, error) {
	return _Badges.Contract.MintBatchAndApproveMarket(&_Badges.TransactOpts, uri, _total)
}

// MintToAndApproveMarket is a paid mutator transaction binding the contract method 0x3560027a.
//
// Solidity: function mintToAndApproveMarket(address to, string uri) returns()
func (_Badges *BadgesTransactor) MintToAndApproveMarket(opts *bind.TransactOpts, to common.Address, uri string) (*types.Transaction, error) {
	return _Badges.contract.Transact(opts, "mintToAndApproveMarket", to, uri)
}

// MintToAndApproveMarket is a paid mutator transaction binding the contract method 0x3560027a.
//
// Solidity: function mintToAndApproveMarket(address to, string uri) returns()
func (_Badges *BadgesSession) MintToAndApproveMarket(to common.Address, uri string) (*types.Transaction, error) {
	return _Badges.Contract.MintToAndApproveMarket(&_Badges.TransactOpts, to, uri)
}

// MintToAndApproveMarket is a paid mutator transaction binding the contract method 0x3560027a.
//
// Solidity: function mintToAndApproveMarket(address to, string uri) returns()
func (_Badges *BadgesTransactorSession) MintToAndApproveMarket(to common.Address, uri string) (*types.Transaction, error) {
	return _Badges.Contract.MintToAndApproveMarket(&_Badges.TransactOpts, to, uri)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Badges *BadgesTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Badges.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Badges *BadgesSession) RenounceOwnership() (*types.Transaction, error) {
	return _Badges.Contract.RenounceOwnership(&_Badges.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Badges *BadgesTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Badges.Contract.RenounceOwnership(&_Badges.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Badges *BadgesTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Badges.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Badges *BadgesSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Badges.Contract.SafeTransferFrom(&_Badges.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Badges *BadgesTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Badges.Contract.SafeTransferFrom(&_Badges.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Badges *BadgesTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Badges.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Badges *BadgesSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Badges.Contract.SafeTransferFrom0(&_Badges.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Badges *BadgesTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Badges.Contract.SafeTransferFrom0(&_Badges.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Badges *BadgesTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Badges.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Badges *BadgesSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Badges.Contract.SetApprovalForAll(&_Badges.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Badges *BadgesTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Badges.Contract.SetApprovalForAll(&_Badges.TransactOpts, operator, approved)
}

// SetTrustKeysAuction is a paid mutator transaction binding the contract method 0xddbbf7dc.
//
// Solidity: function setTrustKeysAuction(address _addressTrustKeysAuction) returns()
func (_Badges *BadgesTransactor) SetTrustKeysAuction(opts *bind.TransactOpts, _addressTrustKeysAuction common.Address) (*types.Transaction, error) {
	return _Badges.contract.Transact(opts, "setTrustKeysAuction", _addressTrustKeysAuction)
}

// SetTrustKeysAuction is a paid mutator transaction binding the contract method 0xddbbf7dc.
//
// Solidity: function setTrustKeysAuction(address _addressTrustKeysAuction) returns()
func (_Badges *BadgesSession) SetTrustKeysAuction(_addressTrustKeysAuction common.Address) (*types.Transaction, error) {
	return _Badges.Contract.SetTrustKeysAuction(&_Badges.TransactOpts, _addressTrustKeysAuction)
}

// SetTrustKeysAuction is a paid mutator transaction binding the contract method 0xddbbf7dc.
//
// Solidity: function setTrustKeysAuction(address _addressTrustKeysAuction) returns()
func (_Badges *BadgesTransactorSession) SetTrustKeysAuction(_addressTrustKeysAuction common.Address) (*types.Transaction, error) {
	return _Badges.Contract.SetTrustKeysAuction(&_Badges.TransactOpts, _addressTrustKeysAuction)
}

// SetTrustKeysMarketplace is a paid mutator transaction binding the contract method 0xcdf07bd6.
//
// Solidity: function setTrustKeysMarketplace(address _addressTrustKeysMarketplace) returns()
func (_Badges *BadgesTransactor) SetTrustKeysMarketplace(opts *bind.TransactOpts, _addressTrustKeysMarketplace common.Address) (*types.Transaction, error) {
	return _Badges.contract.Transact(opts, "setTrustKeysMarketplace", _addressTrustKeysMarketplace)
}

// SetTrustKeysMarketplace is a paid mutator transaction binding the contract method 0xcdf07bd6.
//
// Solidity: function setTrustKeysMarketplace(address _addressTrustKeysMarketplace) returns()
func (_Badges *BadgesSession) SetTrustKeysMarketplace(_addressTrustKeysMarketplace common.Address) (*types.Transaction, error) {
	return _Badges.Contract.SetTrustKeysMarketplace(&_Badges.TransactOpts, _addressTrustKeysMarketplace)
}

// SetTrustKeysMarketplace is a paid mutator transaction binding the contract method 0xcdf07bd6.
//
// Solidity: function setTrustKeysMarketplace(address _addressTrustKeysMarketplace) returns()
func (_Badges *BadgesTransactorSession) SetTrustKeysMarketplace(_addressTrustKeysMarketplace common.Address) (*types.Transaction, error) {
	return _Badges.Contract.SetTrustKeysMarketplace(&_Badges.TransactOpts, _addressTrustKeysMarketplace)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Badges *BadgesTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Badges.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Badges *BadgesSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Badges.Contract.TransferFrom(&_Badges.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Badges *BadgesTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Badges.Contract.TransferFrom(&_Badges.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Badges *BadgesTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Badges.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Badges *BadgesSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Badges.Contract.TransferOwnership(&_Badges.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Badges *BadgesTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Badges.Contract.TransferOwnership(&_Badges.TransactOpts, newOwner)
}

// BadgesApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Badges contract.
type BadgesApprovalIterator struct {
	Event *BadgesApproval // Event containing the contract specifics and raw log

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
func (it *BadgesApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BadgesApproval)
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
		it.Event = new(BadgesApproval)
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
func (it *BadgesApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BadgesApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BadgesApproval represents a Approval event raised by the Badges contract.
type BadgesApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Badges *BadgesFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*BadgesApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Badges.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &BadgesApprovalIterator{contract: _Badges.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Badges *BadgesFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *BadgesApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Badges.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BadgesApproval)
				if err := _Badges.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Badges *BadgesFilterer) ParseApproval(log types.Log) (*BadgesApproval, error) {
	event := new(BadgesApproval)
	if err := _Badges.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BadgesApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Badges contract.
type BadgesApprovalForAllIterator struct {
	Event *BadgesApprovalForAll // Event containing the contract specifics and raw log

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
func (it *BadgesApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BadgesApprovalForAll)
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
		it.Event = new(BadgesApprovalForAll)
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
func (it *BadgesApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BadgesApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BadgesApprovalForAll represents a ApprovalForAll event raised by the Badges contract.
type BadgesApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Badges *BadgesFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*BadgesApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Badges.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &BadgesApprovalForAllIterator{contract: _Badges.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Badges *BadgesFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *BadgesApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Badges.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BadgesApprovalForAll)
				if err := _Badges.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Badges *BadgesFilterer) ParseApprovalForAll(log types.Log) (*BadgesApprovalForAll, error) {
	event := new(BadgesApprovalForAll)
	if err := _Badges.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BadgesBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the Badges contract.
type BadgesBurnIterator struct {
	Event *BadgesBurn // Event containing the contract specifics and raw log

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
func (it *BadgesBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BadgesBurn)
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
		it.Event = new(BadgesBurn)
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
func (it *BadgesBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BadgesBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BadgesBurn represents a Burn event raised by the Badges contract.
type BadgesBurn struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0xb90306ad06b2a6ff86ddc9327db583062895ef6540e62dc50add009db5b356eb.
//
// Solidity: event Burn(uint256 tokenId)
func (_Badges *BadgesFilterer) FilterBurn(opts *bind.FilterOpts) (*BadgesBurnIterator, error) {

	logs, sub, err := _Badges.contract.FilterLogs(opts, "Burn")
	if err != nil {
		return nil, err
	}
	return &BadgesBurnIterator{contract: _Badges.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0xb90306ad06b2a6ff86ddc9327db583062895ef6540e62dc50add009db5b356eb.
//
// Solidity: event Burn(uint256 tokenId)
func (_Badges *BadgesFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *BadgesBurn) (event.Subscription, error) {

	logs, sub, err := _Badges.contract.WatchLogs(opts, "Burn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BadgesBurn)
				if err := _Badges.contract.UnpackLog(event, "Burn", log); err != nil {
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

// ParseBurn is a log parse operation binding the contract event 0xb90306ad06b2a6ff86ddc9327db583062895ef6540e62dc50add009db5b356eb.
//
// Solidity: event Burn(uint256 tokenId)
func (_Badges *BadgesFilterer) ParseBurn(log types.Log) (*BadgesBurn, error) {
	event := new(BadgesBurn)
	if err := _Badges.contract.UnpackLog(event, "Burn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BadgesMintAndApproveMarketIterator is returned from FilterMintAndApproveMarket and is used to iterate over the raw logs and unpacked data for MintAndApproveMarket events raised by the Badges contract.
type BadgesMintAndApproveMarketIterator struct {
	Event *BadgesMintAndApproveMarket // Event containing the contract specifics and raw log

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
func (it *BadgesMintAndApproveMarketIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BadgesMintAndApproveMarket)
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
		it.Event = new(BadgesMintAndApproveMarket)
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
func (it *BadgesMintAndApproveMarketIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BadgesMintAndApproveMarketIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BadgesMintAndApproveMarket represents a MintAndApproveMarket event raised by the Badges contract.
type BadgesMintAndApproveMarket struct {
	TokenId *big.Int
	Uri     string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMintAndApproveMarket is a free log retrieval operation binding the contract event 0xd2eb4318b30f1e8a7272281a8b489657e882c4035287a2f7144f2e7b7ea67593.
//
// Solidity: event MintAndApproveMarket(uint256 tokenId, string uri)
func (_Badges *BadgesFilterer) FilterMintAndApproveMarket(opts *bind.FilterOpts) (*BadgesMintAndApproveMarketIterator, error) {

	logs, sub, err := _Badges.contract.FilterLogs(opts, "MintAndApproveMarket")
	if err != nil {
		return nil, err
	}
	return &BadgesMintAndApproveMarketIterator{contract: _Badges.contract, event: "MintAndApproveMarket", logs: logs, sub: sub}, nil
}

// WatchMintAndApproveMarket is a free log subscription operation binding the contract event 0xd2eb4318b30f1e8a7272281a8b489657e882c4035287a2f7144f2e7b7ea67593.
//
// Solidity: event MintAndApproveMarket(uint256 tokenId, string uri)
func (_Badges *BadgesFilterer) WatchMintAndApproveMarket(opts *bind.WatchOpts, sink chan<- *BadgesMintAndApproveMarket) (event.Subscription, error) {

	logs, sub, err := _Badges.contract.WatchLogs(opts, "MintAndApproveMarket")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BadgesMintAndApproveMarket)
				if err := _Badges.contract.UnpackLog(event, "MintAndApproveMarket", log); err != nil {
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

// ParseMintAndApproveMarket is a log parse operation binding the contract event 0xd2eb4318b30f1e8a7272281a8b489657e882c4035287a2f7144f2e7b7ea67593.
//
// Solidity: event MintAndApproveMarket(uint256 tokenId, string uri)
func (_Badges *BadgesFilterer) ParseMintAndApproveMarket(log types.Log) (*BadgesMintAndApproveMarket, error) {
	event := new(BadgesMintAndApproveMarket)
	if err := _Badges.contract.UnpackLog(event, "MintAndApproveMarket", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BadgesOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Badges contract.
type BadgesOwnershipTransferredIterator struct {
	Event *BadgesOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BadgesOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BadgesOwnershipTransferred)
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
		it.Event = new(BadgesOwnershipTransferred)
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
func (it *BadgesOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BadgesOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BadgesOwnershipTransferred represents a OwnershipTransferred event raised by the Badges contract.
type BadgesOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Badges *BadgesFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BadgesOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Badges.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BadgesOwnershipTransferredIterator{contract: _Badges.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Badges *BadgesFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BadgesOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Badges.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BadgesOwnershipTransferred)
				if err := _Badges.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Badges *BadgesFilterer) ParseOwnershipTransferred(log types.Log) (*BadgesOwnershipTransferred, error) {
	event := new(BadgesOwnershipTransferred)
	if err := _Badges.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BadgesTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Badges contract.
type BadgesTransferIterator struct {
	Event *BadgesTransfer // Event containing the contract specifics and raw log

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
func (it *BadgesTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BadgesTransfer)
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
		it.Event = new(BadgesTransfer)
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
func (it *BadgesTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BadgesTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BadgesTransfer represents a Transfer event raised by the Badges contract.
type BadgesTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Badges *BadgesFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*BadgesTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Badges.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &BadgesTransferIterator{contract: _Badges.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Badges *BadgesFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BadgesTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Badges.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BadgesTransfer)
				if err := _Badges.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Badges *BadgesFilterer) ParseTransfer(log types.Log) (*BadgesTransfer, error) {
	event := new(BadgesTransfer)
	if err := _Badges.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
