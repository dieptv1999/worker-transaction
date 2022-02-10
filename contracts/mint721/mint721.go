// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mint721

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

// Mint721ABI is the input ABI used to generate the binding from.
const Mint721ABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"MintAndApproveMarket\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"__FactoryTrustKeysNFT_init_\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"addressTrustKeysAuction\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"addressTrustKeysMarketplace\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getCreatorByTokenId\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"mintAndApproveMarket\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addressTrustKeysAuction\",\"type\":\"address\"}],\"name\":\"setTrustKeysAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addressTrustKeysMarketplace\",\"type\":\"address\"}],\"name\":\"setTrustKeysMarketplace\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Mint721 is an auto generated Go binding around an Ethereum contract.
type Mint721 struct {
	Mint721Caller     // Read-only binding to the contract
	Mint721Transactor // Write-only binding to the contract
	Mint721Filterer   // Log filterer for contract events
}

// Mint721Caller is an auto generated read-only Go binding around an Ethereum contract.
type Mint721Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Mint721Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Mint721Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Mint721Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Mint721Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Mint721Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Mint721Session struct {
	Contract     *Mint721          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Mint721CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Mint721CallerSession struct {
	Contract *Mint721Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// Mint721TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Mint721TransactorSession struct {
	Contract     *Mint721Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// Mint721Raw is an auto generated low-level Go binding around an Ethereum contract.
type Mint721Raw struct {
	Contract *Mint721 // Generic contract binding to access the raw methods on
}

// Mint721CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Mint721CallerRaw struct {
	Contract *Mint721Caller // Generic read-only contract binding to access the raw methods on
}

// Mint721TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Mint721TransactorRaw struct {
	Contract *Mint721Transactor // Generic write-only contract binding to access the raw methods on
}

// NewMint721 creates a new instance of Mint721, bound to a specific deployed contract.
func NewMint721(address common.Address, backend bind.ContractBackend) (*Mint721, error) {
	contract, err := bindMint721(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Mint721{Mint721Caller: Mint721Caller{contract: contract}, Mint721Transactor: Mint721Transactor{contract: contract}, Mint721Filterer: Mint721Filterer{contract: contract}}, nil
}

// NewMint721Caller creates a new read-only instance of Mint721, bound to a specific deployed contract.
func NewMint721Caller(address common.Address, caller bind.ContractCaller) (*Mint721Caller, error) {
	contract, err := bindMint721(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Mint721Caller{contract: contract}, nil
}

// NewMint721Transactor creates a new write-only instance of Mint721, bound to a specific deployed contract.
func NewMint721Transactor(address common.Address, transactor bind.ContractTransactor) (*Mint721Transactor, error) {
	contract, err := bindMint721(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Mint721Transactor{contract: contract}, nil
}

// NewMint721Filterer creates a new log filterer instance of Mint721, bound to a specific deployed contract.
func NewMint721Filterer(address common.Address, filterer bind.ContractFilterer) (*Mint721Filterer, error) {
	contract, err := bindMint721(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Mint721Filterer{contract: contract}, nil
}

// bindMint721 binds a generic wrapper to an already deployed contract.
func bindMint721(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Mint721ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mint721 *Mint721Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Mint721.Contract.Mint721Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mint721 *Mint721Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mint721.Contract.Mint721Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mint721 *Mint721Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mint721.Contract.Mint721Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mint721 *Mint721CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Mint721.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mint721 *Mint721TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mint721.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mint721 *Mint721TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mint721.Contract.contract.Transact(opts, method, params...)
}

// AddressTrustKeysAuction is a free data retrieval call binding the contract method 0x9316144b.
//
// Solidity: function addressTrustKeysAuction() view returns(address)
func (_Mint721 *Mint721Caller) AddressTrustKeysAuction(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Mint721.contract.Call(opts, &out, "addressTrustKeysAuction")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddressTrustKeysAuction is a free data retrieval call binding the contract method 0x9316144b.
//
// Solidity: function addressTrustKeysAuction() view returns(address)
func (_Mint721 *Mint721Session) AddressTrustKeysAuction() (common.Address, error) {
	return _Mint721.Contract.AddressTrustKeysAuction(&_Mint721.CallOpts)
}

// AddressTrustKeysAuction is a free data retrieval call binding the contract method 0x9316144b.
//
// Solidity: function addressTrustKeysAuction() view returns(address)
func (_Mint721 *Mint721CallerSession) AddressTrustKeysAuction() (common.Address, error) {
	return _Mint721.Contract.AddressTrustKeysAuction(&_Mint721.CallOpts)
}

// AddressTrustKeysMarketplace is a free data retrieval call binding the contract method 0xe58c3af4.
//
// Solidity: function addressTrustKeysMarketplace() view returns(address)
func (_Mint721 *Mint721Caller) AddressTrustKeysMarketplace(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Mint721.contract.Call(opts, &out, "addressTrustKeysMarketplace")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddressTrustKeysMarketplace is a free data retrieval call binding the contract method 0xe58c3af4.
//
// Solidity: function addressTrustKeysMarketplace() view returns(address)
func (_Mint721 *Mint721Session) AddressTrustKeysMarketplace() (common.Address, error) {
	return _Mint721.Contract.AddressTrustKeysMarketplace(&_Mint721.CallOpts)
}

// AddressTrustKeysMarketplace is a free data retrieval call binding the contract method 0xe58c3af4.
//
// Solidity: function addressTrustKeysMarketplace() view returns(address)
func (_Mint721 *Mint721CallerSession) AddressTrustKeysMarketplace() (common.Address, error) {
	return _Mint721.Contract.AddressTrustKeysMarketplace(&_Mint721.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Mint721 *Mint721Caller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Mint721.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Mint721 *Mint721Session) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Mint721.Contract.BalanceOf(&_Mint721.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Mint721 *Mint721CallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Mint721.Contract.BalanceOf(&_Mint721.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Mint721 *Mint721Caller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Mint721.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Mint721 *Mint721Session) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Mint721.Contract.GetApproved(&_Mint721.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Mint721 *Mint721CallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Mint721.Contract.GetApproved(&_Mint721.CallOpts, tokenId)
}

// GetCreatorByTokenId is a free data retrieval call binding the contract method 0x3fd3257b.
//
// Solidity: function getCreatorByTokenId(uint256 tokenId) view returns(address)
func (_Mint721 *Mint721Caller) GetCreatorByTokenId(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Mint721.contract.Call(opts, &out, "getCreatorByTokenId", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetCreatorByTokenId is a free data retrieval call binding the contract method 0x3fd3257b.
//
// Solidity: function getCreatorByTokenId(uint256 tokenId) view returns(address)
func (_Mint721 *Mint721Session) GetCreatorByTokenId(tokenId *big.Int) (common.Address, error) {
	return _Mint721.Contract.GetCreatorByTokenId(&_Mint721.CallOpts, tokenId)
}

// GetCreatorByTokenId is a free data retrieval call binding the contract method 0x3fd3257b.
//
// Solidity: function getCreatorByTokenId(uint256 tokenId) view returns(address)
func (_Mint721 *Mint721CallerSession) GetCreatorByTokenId(tokenId *big.Int) (common.Address, error) {
	return _Mint721.Contract.GetCreatorByTokenId(&_Mint721.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Mint721 *Mint721Caller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Mint721.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Mint721 *Mint721Session) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Mint721.Contract.IsApprovedForAll(&_Mint721.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Mint721 *Mint721CallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Mint721.Contract.IsApprovedForAll(&_Mint721.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Mint721 *Mint721Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Mint721.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Mint721 *Mint721Session) Name() (string, error) {
	return _Mint721.Contract.Name(&_Mint721.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Mint721 *Mint721CallerSession) Name() (string, error) {
	return _Mint721.Contract.Name(&_Mint721.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Mint721 *Mint721Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Mint721.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Mint721 *Mint721Session) Owner() (common.Address, error) {
	return _Mint721.Contract.Owner(&_Mint721.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Mint721 *Mint721CallerSession) Owner() (common.Address, error) {
	return _Mint721.Contract.Owner(&_Mint721.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Mint721 *Mint721Caller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Mint721.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Mint721 *Mint721Session) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Mint721.Contract.OwnerOf(&_Mint721.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Mint721 *Mint721CallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Mint721.Contract.OwnerOf(&_Mint721.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Mint721 *Mint721Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Mint721.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Mint721 *Mint721Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Mint721.Contract.SupportsInterface(&_Mint721.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Mint721 *Mint721CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Mint721.Contract.SupportsInterface(&_Mint721.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Mint721 *Mint721Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Mint721.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Mint721 *Mint721Session) Symbol() (string, error) {
	return _Mint721.Contract.Symbol(&_Mint721.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Mint721 *Mint721CallerSession) Symbol() (string, error) {
	return _Mint721.Contract.Symbol(&_Mint721.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Mint721 *Mint721Caller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Mint721.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Mint721 *Mint721Session) TokenURI(tokenId *big.Int) (string, error) {
	return _Mint721.Contract.TokenURI(&_Mint721.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Mint721 *Mint721CallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Mint721.Contract.TokenURI(&_Mint721.CallOpts, tokenId)
}

// FactoryTrustKeysNFTInit is a paid mutator transaction binding the contract method 0xe7502fff.
//
// Solidity: function __FactoryTrustKeysNFT_init_() returns()
func (_Mint721 *Mint721Transactor) FactoryTrustKeysNFTInit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mint721.contract.Transact(opts, "__FactoryTrustKeysNFT_init_")
}

// FactoryTrustKeysNFTInit is a paid mutator transaction binding the contract method 0xe7502fff.
//
// Solidity: function __FactoryTrustKeysNFT_init_() returns()
func (_Mint721 *Mint721Session) FactoryTrustKeysNFTInit() (*types.Transaction, error) {
	return _Mint721.Contract.FactoryTrustKeysNFTInit(&_Mint721.TransactOpts)
}

// FactoryTrustKeysNFTInit is a paid mutator transaction binding the contract method 0xe7502fff.
//
// Solidity: function __FactoryTrustKeysNFT_init_() returns()
func (_Mint721 *Mint721TransactorSession) FactoryTrustKeysNFTInit() (*types.Transaction, error) {
	return _Mint721.Contract.FactoryTrustKeysNFTInit(&_Mint721.TransactOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Mint721 *Mint721Transactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Mint721.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Mint721 *Mint721Session) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Mint721.Contract.Approve(&_Mint721.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Mint721 *Mint721TransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Mint721.Contract.Approve(&_Mint721.TransactOpts, to, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 _tokenId) returns()
func (_Mint721 *Mint721Transactor) Burn(opts *bind.TransactOpts, _tokenId *big.Int) (*types.Transaction, error) {
	return _Mint721.contract.Transact(opts, "burn", _tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 _tokenId) returns()
func (_Mint721 *Mint721Session) Burn(_tokenId *big.Int) (*types.Transaction, error) {
	return _Mint721.Contract.Burn(&_Mint721.TransactOpts, _tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 _tokenId) returns()
func (_Mint721 *Mint721TransactorSession) Burn(_tokenId *big.Int) (*types.Transaction, error) {
	return _Mint721.Contract.Burn(&_Mint721.TransactOpts, _tokenId)
}

// MintAndApproveMarket is a paid mutator transaction binding the contract method 0x5174e853.
//
// Solidity: function mintAndApproveMarket(string uri) returns()
func (_Mint721 *Mint721Transactor) MintAndApproveMarket(opts *bind.TransactOpts, uri string) (*types.Transaction, error) {
	return _Mint721.contract.Transact(opts, "mintAndApproveMarket", uri)
}

// MintAndApproveMarket is a paid mutator transaction binding the contract method 0x5174e853.
//
// Solidity: function mintAndApproveMarket(string uri) returns()
func (_Mint721 *Mint721Session) MintAndApproveMarket(uri string) (*types.Transaction, error) {
	return _Mint721.Contract.MintAndApproveMarket(&_Mint721.TransactOpts, uri)
}

// MintAndApproveMarket is a paid mutator transaction binding the contract method 0x5174e853.
//
// Solidity: function mintAndApproveMarket(string uri) returns()
func (_Mint721 *Mint721TransactorSession) MintAndApproveMarket(uri string) (*types.Transaction, error) {
	return _Mint721.Contract.MintAndApproveMarket(&_Mint721.TransactOpts, uri)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Mint721 *Mint721Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mint721.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Mint721 *Mint721Session) RenounceOwnership() (*types.Transaction, error) {
	return _Mint721.Contract.RenounceOwnership(&_Mint721.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Mint721 *Mint721TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Mint721.Contract.RenounceOwnership(&_Mint721.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Mint721 *Mint721Transactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Mint721.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Mint721 *Mint721Session) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Mint721.Contract.SafeTransferFrom(&_Mint721.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Mint721 *Mint721TransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Mint721.Contract.SafeTransferFrom(&_Mint721.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Mint721 *Mint721Transactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Mint721.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Mint721 *Mint721Session) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Mint721.Contract.SafeTransferFrom0(&_Mint721.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Mint721 *Mint721TransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Mint721.Contract.SafeTransferFrom0(&_Mint721.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Mint721 *Mint721Transactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Mint721.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Mint721 *Mint721Session) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Mint721.Contract.SetApprovalForAll(&_Mint721.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Mint721 *Mint721TransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Mint721.Contract.SetApprovalForAll(&_Mint721.TransactOpts, operator, approved)
}

// SetTrustKeysAuction is a paid mutator transaction binding the contract method 0xddbbf7dc.
//
// Solidity: function setTrustKeysAuction(address _addressTrustKeysAuction) returns()
func (_Mint721 *Mint721Transactor) SetTrustKeysAuction(opts *bind.TransactOpts, _addressTrustKeysAuction common.Address) (*types.Transaction, error) {
	return _Mint721.contract.Transact(opts, "setTrustKeysAuction", _addressTrustKeysAuction)
}

// SetTrustKeysAuction is a paid mutator transaction binding the contract method 0xddbbf7dc.
//
// Solidity: function setTrustKeysAuction(address _addressTrustKeysAuction) returns()
func (_Mint721 *Mint721Session) SetTrustKeysAuction(_addressTrustKeysAuction common.Address) (*types.Transaction, error) {
	return _Mint721.Contract.SetTrustKeysAuction(&_Mint721.TransactOpts, _addressTrustKeysAuction)
}

// SetTrustKeysAuction is a paid mutator transaction binding the contract method 0xddbbf7dc.
//
// Solidity: function setTrustKeysAuction(address _addressTrustKeysAuction) returns()
func (_Mint721 *Mint721TransactorSession) SetTrustKeysAuction(_addressTrustKeysAuction common.Address) (*types.Transaction, error) {
	return _Mint721.Contract.SetTrustKeysAuction(&_Mint721.TransactOpts, _addressTrustKeysAuction)
}

// SetTrustKeysMarketplace is a paid mutator transaction binding the contract method 0xcdf07bd6.
//
// Solidity: function setTrustKeysMarketplace(address _addressTrustKeysMarketplace) returns()
func (_Mint721 *Mint721Transactor) SetTrustKeysMarketplace(opts *bind.TransactOpts, _addressTrustKeysMarketplace common.Address) (*types.Transaction, error) {
	return _Mint721.contract.Transact(opts, "setTrustKeysMarketplace", _addressTrustKeysMarketplace)
}

// SetTrustKeysMarketplace is a paid mutator transaction binding the contract method 0xcdf07bd6.
//
// Solidity: function setTrustKeysMarketplace(address _addressTrustKeysMarketplace) returns()
func (_Mint721 *Mint721Session) SetTrustKeysMarketplace(_addressTrustKeysMarketplace common.Address) (*types.Transaction, error) {
	return _Mint721.Contract.SetTrustKeysMarketplace(&_Mint721.TransactOpts, _addressTrustKeysMarketplace)
}

// SetTrustKeysMarketplace is a paid mutator transaction binding the contract method 0xcdf07bd6.
//
// Solidity: function setTrustKeysMarketplace(address _addressTrustKeysMarketplace) returns()
func (_Mint721 *Mint721TransactorSession) SetTrustKeysMarketplace(_addressTrustKeysMarketplace common.Address) (*types.Transaction, error) {
	return _Mint721.Contract.SetTrustKeysMarketplace(&_Mint721.TransactOpts, _addressTrustKeysMarketplace)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Mint721 *Mint721Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Mint721.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Mint721 *Mint721Session) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Mint721.Contract.TransferFrom(&_Mint721.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Mint721 *Mint721TransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Mint721.Contract.TransferFrom(&_Mint721.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Mint721 *Mint721Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Mint721.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Mint721 *Mint721Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Mint721.Contract.TransferOwnership(&_Mint721.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Mint721 *Mint721TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Mint721.Contract.TransferOwnership(&_Mint721.TransactOpts, newOwner)
}

// Mint721ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Mint721 contract.
type Mint721ApprovalIterator struct {
	Event *Mint721Approval // Event containing the contract specifics and raw log

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
func (it *Mint721ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Mint721Approval)
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
		it.Event = new(Mint721Approval)
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
func (it *Mint721ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Mint721ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Mint721Approval represents a Approval event raised by the Mint721 contract.
type Mint721Approval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Mint721 *Mint721Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*Mint721ApprovalIterator, error) {

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

	logs, sub, err := _Mint721.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &Mint721ApprovalIterator{contract: _Mint721.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Mint721 *Mint721Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *Mint721Approval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Mint721.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Mint721Approval)
				if err := _Mint721.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Mint721 *Mint721Filterer) ParseApproval(log types.Log) (*Mint721Approval, error) {
	event := new(Mint721Approval)
	if err := _Mint721.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Mint721ApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Mint721 contract.
type Mint721ApprovalForAllIterator struct {
	Event *Mint721ApprovalForAll // Event containing the contract specifics and raw log

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
func (it *Mint721ApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Mint721ApprovalForAll)
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
		it.Event = new(Mint721ApprovalForAll)
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
func (it *Mint721ApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Mint721ApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Mint721ApprovalForAll represents a ApprovalForAll event raised by the Mint721 contract.
type Mint721ApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Mint721 *Mint721Filterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*Mint721ApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Mint721.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &Mint721ApprovalForAllIterator{contract: _Mint721.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Mint721 *Mint721Filterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *Mint721ApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Mint721.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Mint721ApprovalForAll)
				if err := _Mint721.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_Mint721 *Mint721Filterer) ParseApprovalForAll(log types.Log) (*Mint721ApprovalForAll, error) {
	event := new(Mint721ApprovalForAll)
	if err := _Mint721.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Mint721BurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the Mint721 contract.
type Mint721BurnIterator struct {
	Event *Mint721Burn // Event containing the contract specifics and raw log

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
func (it *Mint721BurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Mint721Burn)
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
		it.Event = new(Mint721Burn)
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
func (it *Mint721BurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Mint721BurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Mint721Burn represents a Burn event raised by the Mint721 contract.
type Mint721Burn struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0xb90306ad06b2a6ff86ddc9327db583062895ef6540e62dc50add009db5b356eb.
//
// Solidity: event Burn(uint256 tokenId)
func (_Mint721 *Mint721Filterer) FilterBurn(opts *bind.FilterOpts) (*Mint721BurnIterator, error) {

	logs, sub, err := _Mint721.contract.FilterLogs(opts, "Burn")
	if err != nil {
		return nil, err
	}
	return &Mint721BurnIterator{contract: _Mint721.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0xb90306ad06b2a6ff86ddc9327db583062895ef6540e62dc50add009db5b356eb.
//
// Solidity: event Burn(uint256 tokenId)
func (_Mint721 *Mint721Filterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *Mint721Burn) (event.Subscription, error) {

	logs, sub, err := _Mint721.contract.WatchLogs(opts, "Burn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Mint721Burn)
				if err := _Mint721.contract.UnpackLog(event, "Burn", log); err != nil {
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
func (_Mint721 *Mint721Filterer) ParseBurn(log types.Log) (*Mint721Burn, error) {
	event := new(Mint721Burn)
	if err := _Mint721.contract.UnpackLog(event, "Burn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Mint721MintAndApproveMarketIterator is returned from FilterMintAndApproveMarket and is used to iterate over the raw logs and unpacked data for MintAndApproveMarket events raised by the Mint721 contract.
type Mint721MintAndApproveMarketIterator struct {
	Event *Mint721MintAndApproveMarket // Event containing the contract specifics and raw log

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
func (it *Mint721MintAndApproveMarketIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Mint721MintAndApproveMarket)
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
		it.Event = new(Mint721MintAndApproveMarket)
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
func (it *Mint721MintAndApproveMarketIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Mint721MintAndApproveMarketIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Mint721MintAndApproveMarket represents a MintAndApproveMarket event raised by the Mint721 contract.
type Mint721MintAndApproveMarket struct {
	TokenId *big.Int
	Uri     string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMintAndApproveMarket is a free log retrieval operation binding the contract event 0xd2eb4318b30f1e8a7272281a8b489657e882c4035287a2f7144f2e7b7ea67593.
//
// Solidity: event MintAndApproveMarket(uint256 tokenId, string uri)
func (_Mint721 *Mint721Filterer) FilterMintAndApproveMarket(opts *bind.FilterOpts) (*Mint721MintAndApproveMarketIterator, error) {

	logs, sub, err := _Mint721.contract.FilterLogs(opts, "MintAndApproveMarket")
	if err != nil {
		return nil, err
	}
	return &Mint721MintAndApproveMarketIterator{contract: _Mint721.contract, event: "MintAndApproveMarket", logs: logs, sub: sub}, nil
}

// WatchMintAndApproveMarket is a free log subscription operation binding the contract event 0xd2eb4318b30f1e8a7272281a8b489657e882c4035287a2f7144f2e7b7ea67593.
//
// Solidity: event MintAndApproveMarket(uint256 tokenId, string uri)
func (_Mint721 *Mint721Filterer) WatchMintAndApproveMarket(opts *bind.WatchOpts, sink chan<- *Mint721MintAndApproveMarket) (event.Subscription, error) {

	logs, sub, err := _Mint721.contract.WatchLogs(opts, "MintAndApproveMarket")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Mint721MintAndApproveMarket)
				if err := _Mint721.contract.UnpackLog(event, "MintAndApproveMarket", log); err != nil {
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
func (_Mint721 *Mint721Filterer) ParseMintAndApproveMarket(log types.Log) (*Mint721MintAndApproveMarket, error) {
	event := new(Mint721MintAndApproveMarket)
	if err := _Mint721.contract.UnpackLog(event, "MintAndApproveMarket", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Mint721OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Mint721 contract.
type Mint721OwnershipTransferredIterator struct {
	Event *Mint721OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *Mint721OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Mint721OwnershipTransferred)
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
		it.Event = new(Mint721OwnershipTransferred)
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
func (it *Mint721OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Mint721OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Mint721OwnershipTransferred represents a OwnershipTransferred event raised by the Mint721 contract.
type Mint721OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Mint721 *Mint721Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*Mint721OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Mint721.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &Mint721OwnershipTransferredIterator{contract: _Mint721.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Mint721 *Mint721Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *Mint721OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Mint721.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Mint721OwnershipTransferred)
				if err := _Mint721.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Mint721 *Mint721Filterer) ParseOwnershipTransferred(log types.Log) (*Mint721OwnershipTransferred, error) {
	event := new(Mint721OwnershipTransferred)
	if err := _Mint721.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Mint721TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Mint721 contract.
type Mint721TransferIterator struct {
	Event *Mint721Transfer // Event containing the contract specifics and raw log

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
func (it *Mint721TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Mint721Transfer)
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
		it.Event = new(Mint721Transfer)
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
func (it *Mint721TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Mint721TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Mint721Transfer represents a Transfer event raised by the Mint721 contract.
type Mint721Transfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Mint721 *Mint721Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*Mint721TransferIterator, error) {

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

	logs, sub, err := _Mint721.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &Mint721TransferIterator{contract: _Mint721.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Mint721 *Mint721Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *Mint721Transfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Mint721.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Mint721Transfer)
				if err := _Mint721.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Mint721 *Mint721Filterer) ParseTransfer(log types.Log) (*Mint721Transfer, error) {
	event := new(Mint721Transfer)
	if err := _Mint721.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
