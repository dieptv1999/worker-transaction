// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package auction721

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

// TrustKeysAuctionAuctionInfo is an auto generated low-level Go binding around an user-defined struct.
type TrustKeysAuctionAuctionInfo struct {
	Owner         common.Address
	Winner        common.Address
	Price         *big.Int
	BidStep       *big.Int
	ContractERC20 common.Address
	StartAt       uint64
	EndAt         uint64
	CountBid      uint32
	Exist         bool
}

// Auction721ABI is the input ABI used to generate the binding from.
const Auction721ABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"contractIERC721Metadata\",\"name\":\"contractNFT\",\"type\":\"address\"}],\"name\":\"AuctionCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"contractIERC721Metadata\",\"name\":\"contractNFT\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bidStep\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"contractERC20\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"startAt\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"endAt\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"countBid\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"exist\",\"type\":\"bool\"}],\"indexed\":false,\"internalType\":\"structTrustKeysAuction.AuctionInfo\",\"name\":\"auctionInfo\",\"type\":\"tuple\"}],\"name\":\"AuctionCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"contractIERC721Metadata\",\"name\":\"contractNFT\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bidStep\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"contractERC20\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"startAt\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"endAt\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"countBid\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"exist\",\"type\":\"bool\"}],\"indexed\":false,\"internalType\":\"structTrustKeysAuction.AuctionInfo\",\"name\":\"auctionInfo\",\"type\":\"tuple\"}],\"name\":\"AuctionUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"contractIERC721Metadata\",\"name\":\"contractNFT\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bidStep\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"contractERC20\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"startAt\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"endAt\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"countBid\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"exist\",\"type\":\"bool\"}],\"indexed\":false,\"internalType\":\"structTrustKeysAuction.AuctionInfo\",\"name\":\"auctionInfo\",\"type\":\"tuple\"}],\"name\":\"BidPlaced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"contractIERC721Metadata\",\"name\":\"contractNFT\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bidStep\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"contractERC20\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"startAt\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"endAt\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"countBid\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"exist\",\"type\":\"bool\"}],\"indexed\":false,\"internalType\":\"structTrustKeysAuction.AuctionInfo\",\"name\":\"auctionInfo\",\"type\":\"tuple\"}],\"name\":\"Claim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC721Metadata\",\"name\":\"_contractNFT\",\"type\":\"address\"}],\"name\":\"CancelAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC721Metadata\",\"name\":\"_contractNFT\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_bidStep\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"_contractERC20\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_startAfter\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_duration\",\"type\":\"uint64\"}],\"name\":\"CreateAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC721Metadata\",\"name\":\"_contractNFT\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_bidStep\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"_contractERC20\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_startAfter\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_duration\",\"type\":\"uint64\"}],\"name\":\"EditAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC721Metadata\",\"name\":\"_contractNFT\",\"type\":\"address\"}],\"name\":\"EndAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC721Metadata\",\"name\":\"_contractNFT\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"PlaceBid\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"__TrustKeysAuction_init_\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"addressTrustKeysNFTUpgradeable\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"auctionMarketplace\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bidStep\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"contractERC20\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"startAt\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"endAt\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"countBid\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"exist\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC721Metadata\",\"name\":\"_contractNFT\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"canBid\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"currencies\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeAuction\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeDecimal\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"limitCurrency\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"royalties\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addressTrustKeysNFT\",\"type\":\"address\"}],\"name\":\"setAddressTrustKeysNFT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_listContractERC20\",\"type\":\"address[]\"}],\"name\":\"setCurrencies\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_contractERC20\",\"type\":\"address\"}],\"name\":\"setCurrency\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_limitCurrency\",\"type\":\"bool\"}],\"name\":\"setLimitCurrency\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_royalties\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_feeAuction\",\"type\":\"uint32\"}],\"name\":\"updateFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Auction721 is an auto generated Go binding around an Ethereum contract.
type Auction721 struct {
	Auction721Caller     // Read-only binding to the contract
	Auction721Transactor // Write-only binding to the contract
	Auction721Filterer   // Log filterer for contract events
}

// Auction721Caller is an auto generated read-only Go binding around an Ethereum contract.
type Auction721Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Auction721Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Auction721Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Auction721Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Auction721Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Auction721Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Auction721Session struct {
	Contract     *Auction721       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Auction721CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Auction721CallerSession struct {
	Contract *Auction721Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// Auction721TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Auction721TransactorSession struct {
	Contract     *Auction721Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// Auction721Raw is an auto generated low-level Go binding around an Ethereum contract.
type Auction721Raw struct {
	Contract *Auction721 // Generic contract binding to access the raw methods on
}

// Auction721CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Auction721CallerRaw struct {
	Contract *Auction721Caller // Generic read-only contract binding to access the raw methods on
}

// Auction721TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Auction721TransactorRaw struct {
	Contract *Auction721Transactor // Generic write-only contract binding to access the raw methods on
}

// NewAuction721 creates a new instance of Auction721, bound to a specific deployed contract.
func NewAuction721(address common.Address, backend bind.ContractBackend) (*Auction721, error) {
	contract, err := bindAuction721(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Auction721{Auction721Caller: Auction721Caller{contract: contract}, Auction721Transactor: Auction721Transactor{contract: contract}, Auction721Filterer: Auction721Filterer{contract: contract}}, nil
}

// NewAuction721Caller creates a new read-only instance of Auction721, bound to a specific deployed contract.
func NewAuction721Caller(address common.Address, caller bind.ContractCaller) (*Auction721Caller, error) {
	contract, err := bindAuction721(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Auction721Caller{contract: contract}, nil
}

// NewAuction721Transactor creates a new write-only instance of Auction721, bound to a specific deployed contract.
func NewAuction721Transactor(address common.Address, transactor bind.ContractTransactor) (*Auction721Transactor, error) {
	contract, err := bindAuction721(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Auction721Transactor{contract: contract}, nil
}

// NewAuction721Filterer creates a new log filterer instance of Auction721, bound to a specific deployed contract.
func NewAuction721Filterer(address common.Address, filterer bind.ContractFilterer) (*Auction721Filterer, error) {
	contract, err := bindAuction721(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Auction721Filterer{contract: contract}, nil
}

// bindAuction721 binds a generic wrapper to an already deployed contract.
func bindAuction721(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Auction721ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Auction721 *Auction721Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Auction721.Contract.Auction721Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Auction721 *Auction721Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Auction721.Contract.Auction721Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Auction721 *Auction721Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Auction721.Contract.Auction721Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Auction721 *Auction721CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Auction721.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Auction721 *Auction721TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Auction721.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Auction721 *Auction721TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Auction721.Contract.contract.Transact(opts, method, params...)
}

// AddressTrustKeysNFTUpgradeable is a free data retrieval call binding the contract method 0x947beddb.
//
// Solidity: function addressTrustKeysNFTUpgradeable() view returns(address)
func (_Auction721 *Auction721Caller) AddressTrustKeysNFTUpgradeable(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Auction721.contract.Call(opts, &out, "addressTrustKeysNFTUpgradeable")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddressTrustKeysNFTUpgradeable is a free data retrieval call binding the contract method 0x947beddb.
//
// Solidity: function addressTrustKeysNFTUpgradeable() view returns(address)
func (_Auction721 *Auction721Session) AddressTrustKeysNFTUpgradeable() (common.Address, error) {
	return _Auction721.Contract.AddressTrustKeysNFTUpgradeable(&_Auction721.CallOpts)
}

// AddressTrustKeysNFTUpgradeable is a free data retrieval call binding the contract method 0x947beddb.
//
// Solidity: function addressTrustKeysNFTUpgradeable() view returns(address)
func (_Auction721 *Auction721CallerSession) AddressTrustKeysNFTUpgradeable() (common.Address, error) {
	return _Auction721.Contract.AddressTrustKeysNFTUpgradeable(&_Auction721.CallOpts)
}

// AuctionMarketplace is a free data retrieval call binding the contract method 0x2c46b0af.
//
// Solidity: function auctionMarketplace(address , uint256 ) view returns(address owner, address winner, uint256 price, uint256 bidStep, address contractERC20, uint64 startAt, uint64 endAt, uint32 countBid, bool exist)
func (_Auction721 *Auction721Caller) AuctionMarketplace(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	Owner         common.Address
	Winner        common.Address
	Price         *big.Int
	BidStep       *big.Int
	ContractERC20 common.Address
	StartAt       uint64
	EndAt         uint64
	CountBid      uint32
	Exist         bool
}, error) {
	var out []interface{}
	err := _Auction721.contract.Call(opts, &out, "auctionMarketplace", arg0, arg1)

	outstruct := new(struct {
		Owner         common.Address
		Winner        common.Address
		Price         *big.Int
		BidStep       *big.Int
		ContractERC20 common.Address
		StartAt       uint64
		EndAt         uint64
		CountBid      uint32
		Exist         bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Owner = out[0].(common.Address)
	outstruct.Winner = out[1].(common.Address)
	outstruct.Price = out[2].(*big.Int)
	outstruct.BidStep = out[3].(*big.Int)
	outstruct.ContractERC20 = out[4].(common.Address)
	outstruct.StartAt = out[5].(uint64)
	outstruct.EndAt = out[6].(uint64)
	outstruct.CountBid = out[7].(uint32)
	outstruct.Exist = out[8].(bool)

	return *outstruct, err

}

// AuctionMarketplace is a free data retrieval call binding the contract method 0x2c46b0af.
//
// Solidity: function auctionMarketplace(address , uint256 ) view returns(address owner, address winner, uint256 price, uint256 bidStep, address contractERC20, uint64 startAt, uint64 endAt, uint32 countBid, bool exist)
func (_Auction721 *Auction721Session) AuctionMarketplace(arg0 common.Address, arg1 *big.Int) (struct {
	Owner         common.Address
	Winner        common.Address
	Price         *big.Int
	BidStep       *big.Int
	ContractERC20 common.Address
	StartAt       uint64
	EndAt         uint64
	CountBid      uint32
	Exist         bool
}, error) {
	return _Auction721.Contract.AuctionMarketplace(&_Auction721.CallOpts, arg0, arg1)
}

// AuctionMarketplace is a free data retrieval call binding the contract method 0x2c46b0af.
//
// Solidity: function auctionMarketplace(address , uint256 ) view returns(address owner, address winner, uint256 price, uint256 bidStep, address contractERC20, uint64 startAt, uint64 endAt, uint32 countBid, bool exist)
func (_Auction721 *Auction721CallerSession) AuctionMarketplace(arg0 common.Address, arg1 *big.Int) (struct {
	Owner         common.Address
	Winner        common.Address
	Price         *big.Int
	BidStep       *big.Int
	ContractERC20 common.Address
	StartAt       uint64
	EndAt         uint64
	CountBid      uint32
	Exist         bool
}, error) {
	return _Auction721.Contract.AuctionMarketplace(&_Auction721.CallOpts, arg0, arg1)
}

// CanBid is a free data retrieval call binding the contract method 0xe6f836ed.
//
// Solidity: function canBid(uint256 _tokenId, address _contractNFT, uint256 amount) view returns()
func (_Auction721 *Auction721Caller) CanBid(opts *bind.CallOpts, _tokenId *big.Int, _contractNFT common.Address, amount *big.Int) error {
	var out []interface{}
	err := _Auction721.contract.Call(opts, &out, "canBid", _tokenId, _contractNFT, amount)

	if err != nil {
		return err
	}

	return err

}

// CanBid is a free data retrieval call binding the contract method 0xe6f836ed.
//
// Solidity: function canBid(uint256 _tokenId, address _contractNFT, uint256 amount) view returns()
func (_Auction721 *Auction721Session) CanBid(_tokenId *big.Int, _contractNFT common.Address, amount *big.Int) error {
	return _Auction721.Contract.CanBid(&_Auction721.CallOpts, _tokenId, _contractNFT, amount)
}

// CanBid is a free data retrieval call binding the contract method 0xe6f836ed.
//
// Solidity: function canBid(uint256 _tokenId, address _contractNFT, uint256 amount) view returns()
func (_Auction721 *Auction721CallerSession) CanBid(_tokenId *big.Int, _contractNFT common.Address, amount *big.Int) error {
	return _Auction721.Contract.CanBid(&_Auction721.CallOpts, _tokenId, _contractNFT, amount)
}

// Currencies is a free data retrieval call binding the contract method 0x6036cba3.
//
// Solidity: function currencies(address ) view returns(bool)
func (_Auction721 *Auction721Caller) Currencies(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Auction721.contract.Call(opts, &out, "currencies", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Currencies is a free data retrieval call binding the contract method 0x6036cba3.
//
// Solidity: function currencies(address ) view returns(bool)
func (_Auction721 *Auction721Session) Currencies(arg0 common.Address) (bool, error) {
	return _Auction721.Contract.Currencies(&_Auction721.CallOpts, arg0)
}

// Currencies is a free data retrieval call binding the contract method 0x6036cba3.
//
// Solidity: function currencies(address ) view returns(bool)
func (_Auction721 *Auction721CallerSession) Currencies(arg0 common.Address) (bool, error) {
	return _Auction721.Contract.Currencies(&_Auction721.CallOpts, arg0)
}

// FeeAuction is a free data retrieval call binding the contract method 0xda82140c.
//
// Solidity: function feeAuction() view returns(uint32)
func (_Auction721 *Auction721Caller) FeeAuction(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Auction721.contract.Call(opts, &out, "feeAuction")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// FeeAuction is a free data retrieval call binding the contract method 0xda82140c.
//
// Solidity: function feeAuction() view returns(uint32)
func (_Auction721 *Auction721Session) FeeAuction() (uint32, error) {
	return _Auction721.Contract.FeeAuction(&_Auction721.CallOpts)
}

// FeeAuction is a free data retrieval call binding the contract method 0xda82140c.
//
// Solidity: function feeAuction() view returns(uint32)
func (_Auction721 *Auction721CallerSession) FeeAuction() (uint32, error) {
	return _Auction721.Contract.FeeAuction(&_Auction721.CallOpts)
}

// FeeDecimal is a free data retrieval call binding the contract method 0x7e11b31f.
//
// Solidity: function feeDecimal() view returns(uint32)
func (_Auction721 *Auction721Caller) FeeDecimal(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Auction721.contract.Call(opts, &out, "feeDecimal")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// FeeDecimal is a free data retrieval call binding the contract method 0x7e11b31f.
//
// Solidity: function feeDecimal() view returns(uint32)
func (_Auction721 *Auction721Session) FeeDecimal() (uint32, error) {
	return _Auction721.Contract.FeeDecimal(&_Auction721.CallOpts)
}

// FeeDecimal is a free data retrieval call binding the contract method 0x7e11b31f.
//
// Solidity: function feeDecimal() view returns(uint32)
func (_Auction721 *Auction721CallerSession) FeeDecimal() (uint32, error) {
	return _Auction721.Contract.FeeDecimal(&_Auction721.CallOpts)
}

// LimitCurrency is a free data retrieval call binding the contract method 0x89e167e8.
//
// Solidity: function limitCurrency() view returns(bool)
func (_Auction721 *Auction721Caller) LimitCurrency(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Auction721.contract.Call(opts, &out, "limitCurrency")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// LimitCurrency is a free data retrieval call binding the contract method 0x89e167e8.
//
// Solidity: function limitCurrency() view returns(bool)
func (_Auction721 *Auction721Session) LimitCurrency() (bool, error) {
	return _Auction721.Contract.LimitCurrency(&_Auction721.CallOpts)
}

// LimitCurrency is a free data retrieval call binding the contract method 0x89e167e8.
//
// Solidity: function limitCurrency() view returns(bool)
func (_Auction721 *Auction721CallerSession) LimitCurrency() (bool, error) {
	return _Auction721.Contract.LimitCurrency(&_Auction721.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Auction721 *Auction721Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Auction721.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Auction721 *Auction721Session) Owner() (common.Address, error) {
	return _Auction721.Contract.Owner(&_Auction721.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Auction721 *Auction721CallerSession) Owner() (common.Address, error) {
	return _Auction721.Contract.Owner(&_Auction721.CallOpts)
}

// Royalties is a free data retrieval call binding the contract method 0xf053dc5c.
//
// Solidity: function royalties() view returns(uint32)
func (_Auction721 *Auction721Caller) Royalties(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Auction721.contract.Call(opts, &out, "royalties")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// Royalties is a free data retrieval call binding the contract method 0xf053dc5c.
//
// Solidity: function royalties() view returns(uint32)
func (_Auction721 *Auction721Session) Royalties() (uint32, error) {
	return _Auction721.Contract.Royalties(&_Auction721.CallOpts)
}

// Royalties is a free data retrieval call binding the contract method 0xf053dc5c.
//
// Solidity: function royalties() view returns(uint32)
func (_Auction721 *Auction721CallerSession) Royalties() (uint32, error) {
	return _Auction721.Contract.Royalties(&_Auction721.CallOpts)
}

// CancelAuction is a paid mutator transaction binding the contract method 0x2a6445a4.
//
// Solidity: function CancelAuction(uint256 _tokenId, address _contractNFT) returns()
func (_Auction721 *Auction721Transactor) CancelAuction(opts *bind.TransactOpts, _tokenId *big.Int, _contractNFT common.Address) (*types.Transaction, error) {
	return _Auction721.contract.Transact(opts, "CancelAuction", _tokenId, _contractNFT)
}

// CancelAuction is a paid mutator transaction binding the contract method 0x2a6445a4.
//
// Solidity: function CancelAuction(uint256 _tokenId, address _contractNFT) returns()
func (_Auction721 *Auction721Session) CancelAuction(_tokenId *big.Int, _contractNFT common.Address) (*types.Transaction, error) {
	return _Auction721.Contract.CancelAuction(&_Auction721.TransactOpts, _tokenId, _contractNFT)
}

// CancelAuction is a paid mutator transaction binding the contract method 0x2a6445a4.
//
// Solidity: function CancelAuction(uint256 _tokenId, address _contractNFT) returns()
func (_Auction721 *Auction721TransactorSession) CancelAuction(_tokenId *big.Int, _contractNFT common.Address) (*types.Transaction, error) {
	return _Auction721.Contract.CancelAuction(&_Auction721.TransactOpts, _tokenId, _contractNFT)
}

// CreateAuction is a paid mutator transaction binding the contract method 0xf7ef93c1.
//
// Solidity: function CreateAuction(uint256 _tokenId, address _contractNFT, uint256 _price, uint256 _bidStep, address _contractERC20, uint64 _startAfter, uint64 _duration) returns()
func (_Auction721 *Auction721Transactor) CreateAuction(opts *bind.TransactOpts, _tokenId *big.Int, _contractNFT common.Address, _price *big.Int, _bidStep *big.Int, _contractERC20 common.Address, _startAfter uint64, _duration uint64) (*types.Transaction, error) {
	return _Auction721.contract.Transact(opts, "CreateAuction", _tokenId, _contractNFT, _price, _bidStep, _contractERC20, _startAfter, _duration)
}

// CreateAuction is a paid mutator transaction binding the contract method 0xf7ef93c1.
//
// Solidity: function CreateAuction(uint256 _tokenId, address _contractNFT, uint256 _price, uint256 _bidStep, address _contractERC20, uint64 _startAfter, uint64 _duration) returns()
func (_Auction721 *Auction721Session) CreateAuction(_tokenId *big.Int, _contractNFT common.Address, _price *big.Int, _bidStep *big.Int, _contractERC20 common.Address, _startAfter uint64, _duration uint64) (*types.Transaction, error) {
	return _Auction721.Contract.CreateAuction(&_Auction721.TransactOpts, _tokenId, _contractNFT, _price, _bidStep, _contractERC20, _startAfter, _duration)
}

// CreateAuction is a paid mutator transaction binding the contract method 0xf7ef93c1.
//
// Solidity: function CreateAuction(uint256 _tokenId, address _contractNFT, uint256 _price, uint256 _bidStep, address _contractERC20, uint64 _startAfter, uint64 _duration) returns()
func (_Auction721 *Auction721TransactorSession) CreateAuction(_tokenId *big.Int, _contractNFT common.Address, _price *big.Int, _bidStep *big.Int, _contractERC20 common.Address, _startAfter uint64, _duration uint64) (*types.Transaction, error) {
	return _Auction721.Contract.CreateAuction(&_Auction721.TransactOpts, _tokenId, _contractNFT, _price, _bidStep, _contractERC20, _startAfter, _duration)
}

// EditAuction is a paid mutator transaction binding the contract method 0x7675caec.
//
// Solidity: function EditAuction(uint256 _tokenId, address _contractNFT, uint256 _price, uint256 _bidStep, address _contractERC20, uint64 _startAfter, uint64 _duration) returns()
func (_Auction721 *Auction721Transactor) EditAuction(opts *bind.TransactOpts, _tokenId *big.Int, _contractNFT common.Address, _price *big.Int, _bidStep *big.Int, _contractERC20 common.Address, _startAfter uint64, _duration uint64) (*types.Transaction, error) {
	return _Auction721.contract.Transact(opts, "EditAuction", _tokenId, _contractNFT, _price, _bidStep, _contractERC20, _startAfter, _duration)
}

// EditAuction is a paid mutator transaction binding the contract method 0x7675caec.
//
// Solidity: function EditAuction(uint256 _tokenId, address _contractNFT, uint256 _price, uint256 _bidStep, address _contractERC20, uint64 _startAfter, uint64 _duration) returns()
func (_Auction721 *Auction721Session) EditAuction(_tokenId *big.Int, _contractNFT common.Address, _price *big.Int, _bidStep *big.Int, _contractERC20 common.Address, _startAfter uint64, _duration uint64) (*types.Transaction, error) {
	return _Auction721.Contract.EditAuction(&_Auction721.TransactOpts, _tokenId, _contractNFT, _price, _bidStep, _contractERC20, _startAfter, _duration)
}

// EditAuction is a paid mutator transaction binding the contract method 0x7675caec.
//
// Solidity: function EditAuction(uint256 _tokenId, address _contractNFT, uint256 _price, uint256 _bidStep, address _contractERC20, uint64 _startAfter, uint64 _duration) returns()
func (_Auction721 *Auction721TransactorSession) EditAuction(_tokenId *big.Int, _contractNFT common.Address, _price *big.Int, _bidStep *big.Int, _contractERC20 common.Address, _startAfter uint64, _duration uint64) (*types.Transaction, error) {
	return _Auction721.Contract.EditAuction(&_Auction721.TransactOpts, _tokenId, _contractNFT, _price, _bidStep, _contractERC20, _startAfter, _duration)
}

// EndAuction is a paid mutator transaction binding the contract method 0xbd77fbe3.
//
// Solidity: function EndAuction(uint256 _tokenId, address _contractNFT) returns()
func (_Auction721 *Auction721Transactor) EndAuction(opts *bind.TransactOpts, _tokenId *big.Int, _contractNFT common.Address) (*types.Transaction, error) {
	return _Auction721.contract.Transact(opts, "EndAuction", _tokenId, _contractNFT)
}

// EndAuction is a paid mutator transaction binding the contract method 0xbd77fbe3.
//
// Solidity: function EndAuction(uint256 _tokenId, address _contractNFT) returns()
func (_Auction721 *Auction721Session) EndAuction(_tokenId *big.Int, _contractNFT common.Address) (*types.Transaction, error) {
	return _Auction721.Contract.EndAuction(&_Auction721.TransactOpts, _tokenId, _contractNFT)
}

// EndAuction is a paid mutator transaction binding the contract method 0xbd77fbe3.
//
// Solidity: function EndAuction(uint256 _tokenId, address _contractNFT) returns()
func (_Auction721 *Auction721TransactorSession) EndAuction(_tokenId *big.Int, _contractNFT common.Address) (*types.Transaction, error) {
	return _Auction721.Contract.EndAuction(&_Auction721.TransactOpts, _tokenId, _contractNFT)
}

// PlaceBid is a paid mutator transaction binding the contract method 0x842d2ac3.
//
// Solidity: function PlaceBid(uint256 _tokenId, address _contractNFT, uint256 amount) payable returns()
func (_Auction721 *Auction721Transactor) PlaceBid(opts *bind.TransactOpts, _tokenId *big.Int, _contractNFT common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Auction721.contract.Transact(opts, "PlaceBid", _tokenId, _contractNFT, amount)
}

// PlaceBid is a paid mutator transaction binding the contract method 0x842d2ac3.
//
// Solidity: function PlaceBid(uint256 _tokenId, address _contractNFT, uint256 amount) payable returns()
func (_Auction721 *Auction721Session) PlaceBid(_tokenId *big.Int, _contractNFT common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Auction721.Contract.PlaceBid(&_Auction721.TransactOpts, _tokenId, _contractNFT, amount)
}

// PlaceBid is a paid mutator transaction binding the contract method 0x842d2ac3.
//
// Solidity: function PlaceBid(uint256 _tokenId, address _contractNFT, uint256 amount) payable returns()
func (_Auction721 *Auction721TransactorSession) PlaceBid(_tokenId *big.Int, _contractNFT common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Auction721.Contract.PlaceBid(&_Auction721.TransactOpts, _tokenId, _contractNFT, amount)
}

// TrustKeysAuctionInit is a paid mutator transaction binding the contract method 0x3cf6dbdf.
//
// Solidity: function __TrustKeysAuction_init_() returns()
func (_Auction721 *Auction721Transactor) TrustKeysAuctionInit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Auction721.contract.Transact(opts, "__TrustKeysAuction_init_")
}

// TrustKeysAuctionInit is a paid mutator transaction binding the contract method 0x3cf6dbdf.
//
// Solidity: function __TrustKeysAuction_init_() returns()
func (_Auction721 *Auction721Session) TrustKeysAuctionInit() (*types.Transaction, error) {
	return _Auction721.Contract.TrustKeysAuctionInit(&_Auction721.TransactOpts)
}

// TrustKeysAuctionInit is a paid mutator transaction binding the contract method 0x3cf6dbdf.
//
// Solidity: function __TrustKeysAuction_init_() returns()
func (_Auction721 *Auction721TransactorSession) TrustKeysAuctionInit() (*types.Transaction, error) {
	return _Auction721.Contract.TrustKeysAuctionInit(&_Auction721.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Auction721 *Auction721Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Auction721.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Auction721 *Auction721Session) RenounceOwnership() (*types.Transaction, error) {
	return _Auction721.Contract.RenounceOwnership(&_Auction721.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Auction721 *Auction721TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Auction721.Contract.RenounceOwnership(&_Auction721.TransactOpts)
}

// SetAddressTrustKeysNFT is a paid mutator transaction binding the contract method 0xbdb4af85.
//
// Solidity: function setAddressTrustKeysNFT(address _addressTrustKeysNFT) returns()
func (_Auction721 *Auction721Transactor) SetAddressTrustKeysNFT(opts *bind.TransactOpts, _addressTrustKeysNFT common.Address) (*types.Transaction, error) {
	return _Auction721.contract.Transact(opts, "setAddressTrustKeysNFT", _addressTrustKeysNFT)
}

// SetAddressTrustKeysNFT is a paid mutator transaction binding the contract method 0xbdb4af85.
//
// Solidity: function setAddressTrustKeysNFT(address _addressTrustKeysNFT) returns()
func (_Auction721 *Auction721Session) SetAddressTrustKeysNFT(_addressTrustKeysNFT common.Address) (*types.Transaction, error) {
	return _Auction721.Contract.SetAddressTrustKeysNFT(&_Auction721.TransactOpts, _addressTrustKeysNFT)
}

// SetAddressTrustKeysNFT is a paid mutator transaction binding the contract method 0xbdb4af85.
//
// Solidity: function setAddressTrustKeysNFT(address _addressTrustKeysNFT) returns()
func (_Auction721 *Auction721TransactorSession) SetAddressTrustKeysNFT(_addressTrustKeysNFT common.Address) (*types.Transaction, error) {
	return _Auction721.Contract.SetAddressTrustKeysNFT(&_Auction721.TransactOpts, _addressTrustKeysNFT)
}

// SetCurrencies is a paid mutator transaction binding the contract method 0x8d34f97f.
//
// Solidity: function setCurrencies(address[] _listContractERC20) returns()
func (_Auction721 *Auction721Transactor) SetCurrencies(opts *bind.TransactOpts, _listContractERC20 []common.Address) (*types.Transaction, error) {
	return _Auction721.contract.Transact(opts, "setCurrencies", _listContractERC20)
}

// SetCurrencies is a paid mutator transaction binding the contract method 0x8d34f97f.
//
// Solidity: function setCurrencies(address[] _listContractERC20) returns()
func (_Auction721 *Auction721Session) SetCurrencies(_listContractERC20 []common.Address) (*types.Transaction, error) {
	return _Auction721.Contract.SetCurrencies(&_Auction721.TransactOpts, _listContractERC20)
}

// SetCurrencies is a paid mutator transaction binding the contract method 0x8d34f97f.
//
// Solidity: function setCurrencies(address[] _listContractERC20) returns()
func (_Auction721 *Auction721TransactorSession) SetCurrencies(_listContractERC20 []common.Address) (*types.Transaction, error) {
	return _Auction721.Contract.SetCurrencies(&_Auction721.TransactOpts, _listContractERC20)
}

// SetCurrency is a paid mutator transaction binding the contract method 0x2f84c391.
//
// Solidity: function setCurrency(address _contractERC20) returns()
func (_Auction721 *Auction721Transactor) SetCurrency(opts *bind.TransactOpts, _contractERC20 common.Address) (*types.Transaction, error) {
	return _Auction721.contract.Transact(opts, "setCurrency", _contractERC20)
}

// SetCurrency is a paid mutator transaction binding the contract method 0x2f84c391.
//
// Solidity: function setCurrency(address _contractERC20) returns()
func (_Auction721 *Auction721Session) SetCurrency(_contractERC20 common.Address) (*types.Transaction, error) {
	return _Auction721.Contract.SetCurrency(&_Auction721.TransactOpts, _contractERC20)
}

// SetCurrency is a paid mutator transaction binding the contract method 0x2f84c391.
//
// Solidity: function setCurrency(address _contractERC20) returns()
func (_Auction721 *Auction721TransactorSession) SetCurrency(_contractERC20 common.Address) (*types.Transaction, error) {
	return _Auction721.Contract.SetCurrency(&_Auction721.TransactOpts, _contractERC20)
}

// SetLimitCurrency is a paid mutator transaction binding the contract method 0x7ea90cd6.
//
// Solidity: function setLimitCurrency(bool _limitCurrency) returns()
func (_Auction721 *Auction721Transactor) SetLimitCurrency(opts *bind.TransactOpts, _limitCurrency bool) (*types.Transaction, error) {
	return _Auction721.contract.Transact(opts, "setLimitCurrency", _limitCurrency)
}

// SetLimitCurrency is a paid mutator transaction binding the contract method 0x7ea90cd6.
//
// Solidity: function setLimitCurrency(bool _limitCurrency) returns()
func (_Auction721 *Auction721Session) SetLimitCurrency(_limitCurrency bool) (*types.Transaction, error) {
	return _Auction721.Contract.SetLimitCurrency(&_Auction721.TransactOpts, _limitCurrency)
}

// SetLimitCurrency is a paid mutator transaction binding the contract method 0x7ea90cd6.
//
// Solidity: function setLimitCurrency(bool _limitCurrency) returns()
func (_Auction721 *Auction721TransactorSession) SetLimitCurrency(_limitCurrency bool) (*types.Transaction, error) {
	return _Auction721.Contract.SetLimitCurrency(&_Auction721.TransactOpts, _limitCurrency)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Auction721 *Auction721Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Auction721.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Auction721 *Auction721Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Auction721.Contract.TransferOwnership(&_Auction721.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Auction721 *Auction721TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Auction721.Contract.TransferOwnership(&_Auction721.TransactOpts, newOwner)
}

// UpdateFees is a paid mutator transaction binding the contract method 0xf454add3.
//
// Solidity: function updateFees(uint32 _royalties, uint32 _feeAuction) returns()
func (_Auction721 *Auction721Transactor) UpdateFees(opts *bind.TransactOpts, _royalties uint32, _feeAuction uint32) (*types.Transaction, error) {
	return _Auction721.contract.Transact(opts, "updateFees", _royalties, _feeAuction)
}

// UpdateFees is a paid mutator transaction binding the contract method 0xf454add3.
//
// Solidity: function updateFees(uint32 _royalties, uint32 _feeAuction) returns()
func (_Auction721 *Auction721Session) UpdateFees(_royalties uint32, _feeAuction uint32) (*types.Transaction, error) {
	return _Auction721.Contract.UpdateFees(&_Auction721.TransactOpts, _royalties, _feeAuction)
}

// UpdateFees is a paid mutator transaction binding the contract method 0xf454add3.
//
// Solidity: function updateFees(uint32 _royalties, uint32 _feeAuction) returns()
func (_Auction721 *Auction721TransactorSession) UpdateFees(_royalties uint32, _feeAuction uint32) (*types.Transaction, error) {
	return _Auction721.Contract.UpdateFees(&_Auction721.TransactOpts, _royalties, _feeAuction)
}

// Auction721AuctionCanceledIterator is returned from FilterAuctionCanceled and is used to iterate over the raw logs and unpacked data for AuctionCanceled events raised by the Auction721 contract.
type Auction721AuctionCanceledIterator struct {
	Event *Auction721AuctionCanceled // Event containing the contract specifics and raw log

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
func (it *Auction721AuctionCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Auction721AuctionCanceled)
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
		it.Event = new(Auction721AuctionCanceled)
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
func (it *Auction721AuctionCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Auction721AuctionCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Auction721AuctionCanceled represents a AuctionCanceled event raised by the Auction721 contract.
type Auction721AuctionCanceled struct {
	Seller      common.Address
	TokenId     *big.Int
	ContractNFT common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAuctionCanceled is a free log retrieval operation binding the contract event 0x9692087c146fbd306ff4e4facdd560f7b260a1f7a0278cd22c5c2c8f986559c7.
//
// Solidity: event AuctionCanceled(address seller, uint256 tokenId, address contractNFT)
func (_Auction721 *Auction721Filterer) FilterAuctionCanceled(opts *bind.FilterOpts) (*Auction721AuctionCanceledIterator, error) {

	logs, sub, err := _Auction721.contract.FilterLogs(opts, "AuctionCanceled")
	if err != nil {
		return nil, err
	}
	return &Auction721AuctionCanceledIterator{contract: _Auction721.contract, event: "AuctionCanceled", logs: logs, sub: sub}, nil
}

// WatchAuctionCanceled is a free log subscription operation binding the contract event 0x9692087c146fbd306ff4e4facdd560f7b260a1f7a0278cd22c5c2c8f986559c7.
//
// Solidity: event AuctionCanceled(address seller, uint256 tokenId, address contractNFT)
func (_Auction721 *Auction721Filterer) WatchAuctionCanceled(opts *bind.WatchOpts, sink chan<- *Auction721AuctionCanceled) (event.Subscription, error) {

	logs, sub, err := _Auction721.contract.WatchLogs(opts, "AuctionCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Auction721AuctionCanceled)
				if err := _Auction721.contract.UnpackLog(event, "AuctionCanceled", log); err != nil {
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

// ParseAuctionCanceled is a log parse operation binding the contract event 0x9692087c146fbd306ff4e4facdd560f7b260a1f7a0278cd22c5c2c8f986559c7.
//
// Solidity: event AuctionCanceled(address seller, uint256 tokenId, address contractNFT)
func (_Auction721 *Auction721Filterer) ParseAuctionCanceled(log types.Log) (*Auction721AuctionCanceled, error) {
	event := new(Auction721AuctionCanceled)
	if err := _Auction721.contract.UnpackLog(event, "AuctionCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Auction721AuctionCreatedIterator is returned from FilterAuctionCreated and is used to iterate over the raw logs and unpacked data for AuctionCreated events raised by the Auction721 contract.
type Auction721AuctionCreatedIterator struct {
	Event *Auction721AuctionCreated // Event containing the contract specifics and raw log

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
func (it *Auction721AuctionCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Auction721AuctionCreated)
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
		it.Event = new(Auction721AuctionCreated)
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
func (it *Auction721AuctionCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Auction721AuctionCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Auction721AuctionCreated represents a AuctionCreated event raised by the Auction721 contract.
type Auction721AuctionCreated struct {
	Seller      common.Address
	TokenId     *big.Int
	ContractNFT common.Address
	AuctionInfo TrustKeysAuctionAuctionInfo
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAuctionCreated is a free log retrieval operation binding the contract event 0x253adfe4b5bc99a5ac4c704a93178c2310b0a746f18fcdd153d4dea11ca3cf09.
//
// Solidity: event AuctionCreated(address seller, uint256 tokenId, address contractNFT, (address,address,uint256,uint256,address,uint64,uint64,uint32,bool) auctionInfo)
func (_Auction721 *Auction721Filterer) FilterAuctionCreated(opts *bind.FilterOpts) (*Auction721AuctionCreatedIterator, error) {

	logs, sub, err := _Auction721.contract.FilterLogs(opts, "AuctionCreated")
	if err != nil {
		return nil, err
	}
	return &Auction721AuctionCreatedIterator{contract: _Auction721.contract, event: "AuctionCreated", logs: logs, sub: sub}, nil
}

// WatchAuctionCreated is a free log subscription operation binding the contract event 0x253adfe4b5bc99a5ac4c704a93178c2310b0a746f18fcdd153d4dea11ca3cf09.
//
// Solidity: event AuctionCreated(address seller, uint256 tokenId, address contractNFT, (address,address,uint256,uint256,address,uint64,uint64,uint32,bool) auctionInfo)
func (_Auction721 *Auction721Filterer) WatchAuctionCreated(opts *bind.WatchOpts, sink chan<- *Auction721AuctionCreated) (event.Subscription, error) {

	logs, sub, err := _Auction721.contract.WatchLogs(opts, "AuctionCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Auction721AuctionCreated)
				if err := _Auction721.contract.UnpackLog(event, "AuctionCreated", log); err != nil {
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

// ParseAuctionCreated is a log parse operation binding the contract event 0x253adfe4b5bc99a5ac4c704a93178c2310b0a746f18fcdd153d4dea11ca3cf09.
//
// Solidity: event AuctionCreated(address seller, uint256 tokenId, address contractNFT, (address,address,uint256,uint256,address,uint64,uint64,uint32,bool) auctionInfo)
func (_Auction721 *Auction721Filterer) ParseAuctionCreated(log types.Log) (*Auction721AuctionCreated, error) {
	event := new(Auction721AuctionCreated)
	if err := _Auction721.contract.UnpackLog(event, "AuctionCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Auction721AuctionUpdatedIterator is returned from FilterAuctionUpdated and is used to iterate over the raw logs and unpacked data for AuctionUpdated events raised by the Auction721 contract.
type Auction721AuctionUpdatedIterator struct {
	Event *Auction721AuctionUpdated // Event containing the contract specifics and raw log

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
func (it *Auction721AuctionUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Auction721AuctionUpdated)
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
		it.Event = new(Auction721AuctionUpdated)
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
func (it *Auction721AuctionUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Auction721AuctionUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Auction721AuctionUpdated represents a AuctionUpdated event raised by the Auction721 contract.
type Auction721AuctionUpdated struct {
	Seller      common.Address
	TokenId     *big.Int
	ContractNFT common.Address
	AuctionInfo TrustKeysAuctionAuctionInfo
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAuctionUpdated is a free log retrieval operation binding the contract event 0x0bfb7dd0e0bfe874a0dcd5987dd009b1f1b902b5b3c9bd53494f07463a7687d8.
//
// Solidity: event AuctionUpdated(address seller, uint256 tokenId, address contractNFT, (address,address,uint256,uint256,address,uint64,uint64,uint32,bool) auctionInfo)
func (_Auction721 *Auction721Filterer) FilterAuctionUpdated(opts *bind.FilterOpts) (*Auction721AuctionUpdatedIterator, error) {

	logs, sub, err := _Auction721.contract.FilterLogs(opts, "AuctionUpdated")
	if err != nil {
		return nil, err
	}
	return &Auction721AuctionUpdatedIterator{contract: _Auction721.contract, event: "AuctionUpdated", logs: logs, sub: sub}, nil
}

// WatchAuctionUpdated is a free log subscription operation binding the contract event 0x0bfb7dd0e0bfe874a0dcd5987dd009b1f1b902b5b3c9bd53494f07463a7687d8.
//
// Solidity: event AuctionUpdated(address seller, uint256 tokenId, address contractNFT, (address,address,uint256,uint256,address,uint64,uint64,uint32,bool) auctionInfo)
func (_Auction721 *Auction721Filterer) WatchAuctionUpdated(opts *bind.WatchOpts, sink chan<- *Auction721AuctionUpdated) (event.Subscription, error) {

	logs, sub, err := _Auction721.contract.WatchLogs(opts, "AuctionUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Auction721AuctionUpdated)
				if err := _Auction721.contract.UnpackLog(event, "AuctionUpdated", log); err != nil {
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

// ParseAuctionUpdated is a log parse operation binding the contract event 0x0bfb7dd0e0bfe874a0dcd5987dd009b1f1b902b5b3c9bd53494f07463a7687d8.
//
// Solidity: event AuctionUpdated(address seller, uint256 tokenId, address contractNFT, (address,address,uint256,uint256,address,uint64,uint64,uint32,bool) auctionInfo)
func (_Auction721 *Auction721Filterer) ParseAuctionUpdated(log types.Log) (*Auction721AuctionUpdated, error) {
	event := new(Auction721AuctionUpdated)
	if err := _Auction721.contract.UnpackLog(event, "AuctionUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Auction721BidPlacedIterator is returned from FilterBidPlaced and is used to iterate over the raw logs and unpacked data for BidPlaced events raised by the Auction721 contract.
type Auction721BidPlacedIterator struct {
	Event *Auction721BidPlaced // Event containing the contract specifics and raw log

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
func (it *Auction721BidPlacedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Auction721BidPlaced)
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
		it.Event = new(Auction721BidPlaced)
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
func (it *Auction721BidPlacedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Auction721BidPlacedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Auction721BidPlaced represents a BidPlaced event raised by the Auction721 contract.
type Auction721BidPlaced struct {
	Bidder      common.Address
	TokenId     *big.Int
	ContractNFT common.Address
	AuctionInfo TrustKeysAuctionAuctionInfo
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBidPlaced is a free log retrieval operation binding the contract event 0x448e954bb46c38d6bf3c785db486089602781d873aa8e3bbcceaa07b08381025.
//
// Solidity: event BidPlaced(address bidder, uint256 tokenId, address contractNFT, (address,address,uint256,uint256,address,uint64,uint64,uint32,bool) auctionInfo)
func (_Auction721 *Auction721Filterer) FilterBidPlaced(opts *bind.FilterOpts) (*Auction721BidPlacedIterator, error) {

	logs, sub, err := _Auction721.contract.FilterLogs(opts, "BidPlaced")
	if err != nil {
		return nil, err
	}
	return &Auction721BidPlacedIterator{contract: _Auction721.contract, event: "BidPlaced", logs: logs, sub: sub}, nil
}

// WatchBidPlaced is a free log subscription operation binding the contract event 0x448e954bb46c38d6bf3c785db486089602781d873aa8e3bbcceaa07b08381025.
//
// Solidity: event BidPlaced(address bidder, uint256 tokenId, address contractNFT, (address,address,uint256,uint256,address,uint64,uint64,uint32,bool) auctionInfo)
func (_Auction721 *Auction721Filterer) WatchBidPlaced(opts *bind.WatchOpts, sink chan<- *Auction721BidPlaced) (event.Subscription, error) {

	logs, sub, err := _Auction721.contract.WatchLogs(opts, "BidPlaced")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Auction721BidPlaced)
				if err := _Auction721.contract.UnpackLog(event, "BidPlaced", log); err != nil {
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

// ParseBidPlaced is a log parse operation binding the contract event 0x448e954bb46c38d6bf3c785db486089602781d873aa8e3bbcceaa07b08381025.
//
// Solidity: event BidPlaced(address bidder, uint256 tokenId, address contractNFT, (address,address,uint256,uint256,address,uint64,uint64,uint32,bool) auctionInfo)
func (_Auction721 *Auction721Filterer) ParseBidPlaced(log types.Log) (*Auction721BidPlaced, error) {
	event := new(Auction721BidPlaced)
	if err := _Auction721.contract.UnpackLog(event, "BidPlaced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Auction721ClaimIterator is returned from FilterClaim and is used to iterate over the raw logs and unpacked data for Claim events raised by the Auction721 contract.
type Auction721ClaimIterator struct {
	Event *Auction721Claim // Event containing the contract specifics and raw log

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
func (it *Auction721ClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Auction721Claim)
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
		it.Event = new(Auction721Claim)
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
func (it *Auction721ClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Auction721ClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Auction721Claim represents a Claim event raised by the Auction721 contract.
type Auction721Claim struct {
	Winner      common.Address
	TokenId     *big.Int
	ContractNFT common.Address
	AuctionInfo TrustKeysAuctionAuctionInfo
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterClaim is a free log retrieval operation binding the contract event 0x1fc77468e4411f3077a0299d5c62b09fef07ee6fe22433cf86ca7dfb37aa5622.
//
// Solidity: event Claim(address winner, uint256 tokenId, address contractNFT, (address,address,uint256,uint256,address,uint64,uint64,uint32,bool) auctionInfo)
func (_Auction721 *Auction721Filterer) FilterClaim(opts *bind.FilterOpts) (*Auction721ClaimIterator, error) {

	logs, sub, err := _Auction721.contract.FilterLogs(opts, "Claim")
	if err != nil {
		return nil, err
	}
	return &Auction721ClaimIterator{contract: _Auction721.contract, event: "Claim", logs: logs, sub: sub}, nil
}

// WatchClaim is a free log subscription operation binding the contract event 0x1fc77468e4411f3077a0299d5c62b09fef07ee6fe22433cf86ca7dfb37aa5622.
//
// Solidity: event Claim(address winner, uint256 tokenId, address contractNFT, (address,address,uint256,uint256,address,uint64,uint64,uint32,bool) auctionInfo)
func (_Auction721 *Auction721Filterer) WatchClaim(opts *bind.WatchOpts, sink chan<- *Auction721Claim) (event.Subscription, error) {

	logs, sub, err := _Auction721.contract.WatchLogs(opts, "Claim")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Auction721Claim)
				if err := _Auction721.contract.UnpackLog(event, "Claim", log); err != nil {
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

// ParseClaim is a log parse operation binding the contract event 0x1fc77468e4411f3077a0299d5c62b09fef07ee6fe22433cf86ca7dfb37aa5622.
//
// Solidity: event Claim(address winner, uint256 tokenId, address contractNFT, (address,address,uint256,uint256,address,uint64,uint64,uint32,bool) auctionInfo)
func (_Auction721 *Auction721Filterer) ParseClaim(log types.Log) (*Auction721Claim, error) {
	event := new(Auction721Claim)
	if err := _Auction721.contract.UnpackLog(event, "Claim", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Auction721OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Auction721 contract.
type Auction721OwnershipTransferredIterator struct {
	Event *Auction721OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *Auction721OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Auction721OwnershipTransferred)
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
		it.Event = new(Auction721OwnershipTransferred)
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
func (it *Auction721OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Auction721OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Auction721OwnershipTransferred represents a OwnershipTransferred event raised by the Auction721 contract.
type Auction721OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Auction721 *Auction721Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*Auction721OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Auction721.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &Auction721OwnershipTransferredIterator{contract: _Auction721.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Auction721 *Auction721Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *Auction721OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Auction721.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Auction721OwnershipTransferred)
				if err := _Auction721.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Auction721 *Auction721Filterer) ParseOwnershipTransferred(log types.Log) (*Auction721OwnershipTransferred, error) {
	event := new(Auction721OwnershipTransferred)
	if err := _Auction721.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
