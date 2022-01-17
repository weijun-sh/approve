// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package transactions

import (
	"math/big"
	"strings"

	ethereum "github.com/fsn-dev/fsn-go-sdk/efsn"
	"github.com/fsn-dev/fsn-go-sdk/efsn/accounts/abi"
	"github.com/fsn-dev/fsn-go-sdk/efsn/accounts/abi/bind"
	"github.com/fsn-dev/fsn-go-sdk/efsn/common"
	"github.com/fsn-dev/fsn-go-sdk/efsn/core/types"
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

// CyctokenABI is the input ABI used to generate the binding from.
const CyctokenABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"perShareAmount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"mintInvoker\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"sender\",\"type\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalShares\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"share\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rebaseInvoker\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newInvoker\",\"type\":\"address\"}],\"name\":\"changeRebaseInvoker\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"epoch\",\"type\":\"uint256\"},{\"name\":\"numerator\",\"type\":\"uint256\"},{\"name\":\"denominator\",\"type\":\"uint256\"}],\"name\":\"rebase\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newInvoker\",\"type\":\"address\"}],\"name\":\"changeMintInvoker\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"epoch\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"oldPerShareAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"newPerShareAmount\",\"type\":\"uint256\"}],\"name\":\"Rebase\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"RebaseInvokerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"MintInvokerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// Cyctoken is an auto generated Go binding around an Ethereum contract.
type Cyctoken struct {
	CyctokenCaller     // Read-only binding to the contract
	CyctokenTransactor // Write-only binding to the contract
	CyctokenFilterer   // Log filterer for contract events
}

// CyctokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type CyctokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CyctokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CyctokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CyctokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CyctokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CyctokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CyctokenSession struct {
	Contract     *Cyctoken         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CyctokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CyctokenCallerSession struct {
	Contract *CyctokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// CyctokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CyctokenTransactorSession struct {
	Contract     *CyctokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// CyctokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type CyctokenRaw struct {
	Contract *Cyctoken // Generic contract binding to access the raw methods on
}

// CyctokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CyctokenCallerRaw struct {
	Contract *CyctokenCaller // Generic read-only contract binding to access the raw methods on
}

// CyctokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CyctokenTransactorRaw struct {
	Contract *CyctokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCyctoken creates a new instance of Cyctoken, bound to a specific deployed contract.
func NewCyctoken(address common.Address, backend bind.ContractBackend) (*Cyctoken, error) {
	contract, err := bindCyctoken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Cyctoken{CyctokenCaller: CyctokenCaller{contract: contract}, CyctokenTransactor: CyctokenTransactor{contract: contract}, CyctokenFilterer: CyctokenFilterer{contract: contract}}, nil
}

// NewCyctokenCaller creates a new read-only instance of Cyctoken, bound to a specific deployed contract.
func NewCyctokenCaller(address common.Address, caller bind.ContractCaller) (*CyctokenCaller, error) {
	contract, err := bindCyctoken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CyctokenCaller{contract: contract}, nil
}

// NewCyctokenTransactor creates a new write-only instance of Cyctoken, bound to a specific deployed contract.
func NewCyctokenTransactor(address common.Address, transactor bind.ContractTransactor) (*CyctokenTransactor, error) {
	contract, err := bindCyctoken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CyctokenTransactor{contract: contract}, nil
}

// NewCyctokenFilterer creates a new log filterer instance of Cyctoken, bound to a specific deployed contract.
func NewCyctokenFilterer(address common.Address, filterer bind.ContractFilterer) (*CyctokenFilterer, error) {
	contract, err := bindCyctoken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CyctokenFilterer{contract: contract}, nil
}

// bindCyctoken binds a generic wrapper to an already deployed contract.
func bindCyctoken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CyctokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Cyctoken *CyctokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Cyctoken.Contract.CyctokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Cyctoken *CyctokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cyctoken.Contract.CyctokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Cyctoken *CyctokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Cyctoken.Contract.CyctokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Cyctoken *CyctokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Cyctoken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Cyctoken *CyctokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cyctoken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Cyctoken *CyctokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Cyctoken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Cyctoken *CyctokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Cyctoken.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Cyctoken *CyctokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Cyctoken.Contract.Allowance(&_Cyctoken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Cyctoken *CyctokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Cyctoken.Contract.Allowance(&_Cyctoken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Cyctoken *CyctokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Cyctoken.contract.Call(opts, out, "balanceOf", account)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Cyctoken *CyctokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Cyctoken.Contract.BalanceOf(&_Cyctoken.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Cyctoken *CyctokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Cyctoken.Contract.BalanceOf(&_Cyctoken.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Cyctoken *CyctokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Cyctoken.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Cyctoken *CyctokenSession) Decimals() (uint8, error) {
	return _Cyctoken.Contract.Decimals(&_Cyctoken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Cyctoken *CyctokenCallerSession) Decimals() (uint8, error) {
	return _Cyctoken.Contract.Decimals(&_Cyctoken.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Cyctoken *CyctokenCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Cyctoken.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Cyctoken *CyctokenSession) IsOwner() (bool, error) {
	return _Cyctoken.Contract.IsOwner(&_Cyctoken.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Cyctoken *CyctokenCallerSession) IsOwner() (bool, error) {
	return _Cyctoken.Contract.IsOwner(&_Cyctoken.CallOpts)
}

// MintInvoker is a free data retrieval call binding the contract method 0x0ec31d73.
//
// Solidity: function mintInvoker() view returns(address)
func (_Cyctoken *CyctokenCaller) MintInvoker(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Cyctoken.contract.Call(opts, out, "mintInvoker")
	return *ret0, err
}

// MintInvoker is a free data retrieval call binding the contract method 0x0ec31d73.
//
// Solidity: function mintInvoker() view returns(address)
func (_Cyctoken *CyctokenSession) MintInvoker() (common.Address, error) {
	return _Cyctoken.Contract.MintInvoker(&_Cyctoken.CallOpts)
}

// MintInvoker is a free data retrieval call binding the contract method 0x0ec31d73.
//
// Solidity: function mintInvoker() view returns(address)
func (_Cyctoken *CyctokenCallerSession) MintInvoker() (common.Address, error) {
	return _Cyctoken.Contract.MintInvoker(&_Cyctoken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Cyctoken *CyctokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Cyctoken.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Cyctoken *CyctokenSession) Name() (string, error) {
	return _Cyctoken.Contract.Name(&_Cyctoken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Cyctoken *CyctokenCallerSession) Name() (string, error) {
	return _Cyctoken.Contract.Name(&_Cyctoken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Cyctoken *CyctokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Cyctoken.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Cyctoken *CyctokenSession) Owner() (common.Address, error) {
	return _Cyctoken.Contract.Owner(&_Cyctoken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Cyctoken *CyctokenCallerSession) Owner() (common.Address, error) {
	return _Cyctoken.Contract.Owner(&_Cyctoken.CallOpts)
}

// PerShareAmount is a free data retrieval call binding the contract method 0x0552803e.
//
// Solidity: function perShareAmount() view returns(uint256)
func (_Cyctoken *CyctokenCaller) PerShareAmount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Cyctoken.contract.Call(opts, out, "perShareAmount")
	return *ret0, err
}

// PerShareAmount is a free data retrieval call binding the contract method 0x0552803e.
//
// Solidity: function perShareAmount() view returns(uint256)
func (_Cyctoken *CyctokenSession) PerShareAmount() (*big.Int, error) {
	return _Cyctoken.Contract.PerShareAmount(&_Cyctoken.CallOpts)
}

// PerShareAmount is a free data retrieval call binding the contract method 0x0552803e.
//
// Solidity: function perShareAmount() view returns(uint256)
func (_Cyctoken *CyctokenCallerSession) PerShareAmount() (*big.Int, error) {
	return _Cyctoken.Contract.PerShareAmount(&_Cyctoken.CallOpts)
}

// RebaseInvoker is a free data retrieval call binding the contract method 0x5d45d09d.
//
// Solidity: function rebaseInvoker() view returns(address)
func (_Cyctoken *CyctokenCaller) RebaseInvoker(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Cyctoken.contract.Call(opts, out, "rebaseInvoker")
	return *ret0, err
}

// RebaseInvoker is a free data retrieval call binding the contract method 0x5d45d09d.
//
// Solidity: function rebaseInvoker() view returns(address)
func (_Cyctoken *CyctokenSession) RebaseInvoker() (common.Address, error) {
	return _Cyctoken.Contract.RebaseInvoker(&_Cyctoken.CallOpts)
}

// RebaseInvoker is a free data retrieval call binding the contract method 0x5d45d09d.
//
// Solidity: function rebaseInvoker() view returns(address)
func (_Cyctoken *CyctokenCallerSession) RebaseInvoker() (common.Address, error) {
	return _Cyctoken.Contract.RebaseInvoker(&_Cyctoken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Cyctoken *CyctokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Cyctoken.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Cyctoken *CyctokenSession) Symbol() (string, error) {
	return _Cyctoken.Contract.Symbol(&_Cyctoken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Cyctoken *CyctokenCallerSession) Symbol() (string, error) {
	return _Cyctoken.Contract.Symbol(&_Cyctoken.CallOpts)
}

// TotalShares is a free data retrieval call binding the contract method 0x3a98ef39.
//
// Solidity: function totalShares() view returns(uint256)
func (_Cyctoken *CyctokenCaller) TotalShares(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Cyctoken.contract.Call(opts, out, "totalShares")
	return *ret0, err
}

// TotalShares is a free data retrieval call binding the contract method 0x3a98ef39.
//
// Solidity: function totalShares() view returns(uint256)
func (_Cyctoken *CyctokenSession) TotalShares() (*big.Int, error) {
	return _Cyctoken.Contract.TotalShares(&_Cyctoken.CallOpts)
}

// TotalShares is a free data retrieval call binding the contract method 0x3a98ef39.
//
// Solidity: function totalShares() view returns(uint256)
func (_Cyctoken *CyctokenCallerSession) TotalShares() (*big.Int, error) {
	return _Cyctoken.Contract.TotalShares(&_Cyctoken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Cyctoken *CyctokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Cyctoken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Cyctoken *CyctokenSession) TotalSupply() (*big.Int, error) {
	return _Cyctoken.Contract.TotalSupply(&_Cyctoken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Cyctoken *CyctokenCallerSession) TotalSupply() (*big.Int, error) {
	return _Cyctoken.Contract.TotalSupply(&_Cyctoken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Cyctoken *CyctokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Cyctoken.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Cyctoken *CyctokenSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Cyctoken.Contract.Approve(&_Cyctoken.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Cyctoken *CyctokenTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Cyctoken.Contract.Approve(&_Cyctoken.TransactOpts, spender, amount)
}

// ChangeMintInvoker is a paid mutator transaction binding the contract method 0xe6eec40b.
//
// Solidity: function changeMintInvoker(address newInvoker) returns()
func (_Cyctoken *CyctokenTransactor) ChangeMintInvoker(opts *bind.TransactOpts, newInvoker common.Address) (*types.Transaction, error) {
	return _Cyctoken.contract.Transact(opts, "changeMintInvoker", newInvoker)
}

// ChangeMintInvoker is a paid mutator transaction binding the contract method 0xe6eec40b.
//
// Solidity: function changeMintInvoker(address newInvoker) returns()
func (_Cyctoken *CyctokenSession) ChangeMintInvoker(newInvoker common.Address) (*types.Transaction, error) {
	return _Cyctoken.Contract.ChangeMintInvoker(&_Cyctoken.TransactOpts, newInvoker)
}

// ChangeMintInvoker is a paid mutator transaction binding the contract method 0xe6eec40b.
//
// Solidity: function changeMintInvoker(address newInvoker) returns()
func (_Cyctoken *CyctokenTransactorSession) ChangeMintInvoker(newInvoker common.Address) (*types.Transaction, error) {
	return _Cyctoken.Contract.ChangeMintInvoker(&_Cyctoken.TransactOpts, newInvoker)
}

// ChangeRebaseInvoker is a paid mutator transaction binding the contract method 0x8e88de11.
//
// Solidity: function changeRebaseInvoker(address newInvoker) returns()
func (_Cyctoken *CyctokenTransactor) ChangeRebaseInvoker(opts *bind.TransactOpts, newInvoker common.Address) (*types.Transaction, error) {
	return _Cyctoken.contract.Transact(opts, "changeRebaseInvoker", newInvoker)
}

// ChangeRebaseInvoker is a paid mutator transaction binding the contract method 0x8e88de11.
//
// Solidity: function changeRebaseInvoker(address newInvoker) returns()
func (_Cyctoken *CyctokenSession) ChangeRebaseInvoker(newInvoker common.Address) (*types.Transaction, error) {
	return _Cyctoken.Contract.ChangeRebaseInvoker(&_Cyctoken.TransactOpts, newInvoker)
}

// ChangeRebaseInvoker is a paid mutator transaction binding the contract method 0x8e88de11.
//
// Solidity: function changeRebaseInvoker(address newInvoker) returns()
func (_Cyctoken *CyctokenTransactorSession) ChangeRebaseInvoker(newInvoker common.Address) (*types.Transaction, error) {
	return _Cyctoken.Contract.ChangeRebaseInvoker(&_Cyctoken.TransactOpts, newInvoker)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 share) returns()
func (_Cyctoken *CyctokenTransactor) Mint(opts *bind.TransactOpts, to common.Address, share *big.Int) (*types.Transaction, error) {
	return _Cyctoken.contract.Transact(opts, "mint", to, share)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 share) returns()
func (_Cyctoken *CyctokenSession) Mint(to common.Address, share *big.Int) (*types.Transaction, error) {
	return _Cyctoken.Contract.Mint(&_Cyctoken.TransactOpts, to, share)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 share) returns()
func (_Cyctoken *CyctokenTransactorSession) Mint(to common.Address, share *big.Int) (*types.Transaction, error) {
	return _Cyctoken.Contract.Mint(&_Cyctoken.TransactOpts, to, share)
}

// Rebase is a paid mutator transaction binding the contract method 0xa48daa4f.
//
// Solidity: function rebase(uint256 epoch, uint256 numerator, uint256 denominator) returns(uint256)
func (_Cyctoken *CyctokenTransactor) Rebase(opts *bind.TransactOpts, epoch *big.Int, numerator *big.Int, denominator *big.Int) (*types.Transaction, error) {
	return _Cyctoken.contract.Transact(opts, "rebase", epoch, numerator, denominator)
}

// Rebase is a paid mutator transaction binding the contract method 0xa48daa4f.
//
// Solidity: function rebase(uint256 epoch, uint256 numerator, uint256 denominator) returns(uint256)
func (_Cyctoken *CyctokenSession) Rebase(epoch *big.Int, numerator *big.Int, denominator *big.Int) (*types.Transaction, error) {
	return _Cyctoken.Contract.Rebase(&_Cyctoken.TransactOpts, epoch, numerator, denominator)
}

// Rebase is a paid mutator transaction binding the contract method 0xa48daa4f.
//
// Solidity: function rebase(uint256 epoch, uint256 numerator, uint256 denominator) returns(uint256)
func (_Cyctoken *CyctokenTransactorSession) Rebase(epoch *big.Int, numerator *big.Int, denominator *big.Int) (*types.Transaction, error) {
	return _Cyctoken.Contract.Rebase(&_Cyctoken.TransactOpts, epoch, numerator, denominator)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Cyctoken *CyctokenTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cyctoken.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Cyctoken *CyctokenSession) RenounceOwnership() (*types.Transaction, error) {
	return _Cyctoken.Contract.RenounceOwnership(&_Cyctoken.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Cyctoken *CyctokenTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Cyctoken.Contract.RenounceOwnership(&_Cyctoken.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Cyctoken *CyctokenTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Cyctoken.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Cyctoken *CyctokenSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Cyctoken.Contract.Transfer(&_Cyctoken.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Cyctoken *CyctokenTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Cyctoken.Contract.Transfer(&_Cyctoken.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Cyctoken *CyctokenTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Cyctoken.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Cyctoken *CyctokenSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Cyctoken.Contract.TransferFrom(&_Cyctoken.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Cyctoken *CyctokenTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Cyctoken.Contract.TransferFrom(&_Cyctoken.TransactOpts, sender, recipient, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Cyctoken *CyctokenTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Cyctoken.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Cyctoken *CyctokenSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Cyctoken.Contract.TransferOwnership(&_Cyctoken.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Cyctoken *CyctokenTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Cyctoken.Contract.TransferOwnership(&_Cyctoken.TransactOpts, newOwner)
}

// CyctokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Cyctoken contract.
type CyctokenApprovalIterator struct {
	Event *CyctokenApproval // Event containing the contract specifics and raw log

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
func (it *CyctokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CyctokenApproval)
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
		it.Event = new(CyctokenApproval)
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
func (it *CyctokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CyctokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CyctokenApproval represents a Approval event raised by the Cyctoken contract.
type CyctokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_Cyctoken *CyctokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*CyctokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Cyctoken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &CyctokenApprovalIterator{contract: _Cyctoken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_Cyctoken *CyctokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *CyctokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Cyctoken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CyctokenApproval)
				if err := _Cyctoken.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_Cyctoken *CyctokenFilterer) ParseApproval(log types.Log) (*CyctokenApproval, error) {
	event := new(CyctokenApproval)
	if err := _Cyctoken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CyctokenMintInvokerChangedIterator is returned from FilterMintInvokerChanged and is used to iterate over the raw logs and unpacked data for MintInvokerChanged events raised by the Cyctoken contract.
type CyctokenMintInvokerChangedIterator struct {
	Event *CyctokenMintInvokerChanged // Event containing the contract specifics and raw log

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
func (it *CyctokenMintInvokerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CyctokenMintInvokerChanged)
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
		it.Event = new(CyctokenMintInvokerChanged)
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
func (it *CyctokenMintInvokerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CyctokenMintInvokerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CyctokenMintInvokerChanged represents a MintInvokerChanged event raised by the Cyctoken contract.
type CyctokenMintInvokerChanged struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterMintInvokerChanged is a free log retrieval operation binding the contract event 0xfa319a129fc873627781c98bb447f65d0fc7eeca75f5ca8224de0e2feaffa129.
//
// Solidity: event MintInvokerChanged(address indexed previousOwner, address indexed newOwner)
func (_Cyctoken *CyctokenFilterer) FilterMintInvokerChanged(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*CyctokenMintInvokerChangedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Cyctoken.contract.FilterLogs(opts, "MintInvokerChanged", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CyctokenMintInvokerChangedIterator{contract: _Cyctoken.contract, event: "MintInvokerChanged", logs: logs, sub: sub}, nil
}

// WatchMintInvokerChanged is a free log subscription operation binding the contract event 0xfa319a129fc873627781c98bb447f65d0fc7eeca75f5ca8224de0e2feaffa129.
//
// Solidity: event MintInvokerChanged(address indexed previousOwner, address indexed newOwner)
func (_Cyctoken *CyctokenFilterer) WatchMintInvokerChanged(opts *bind.WatchOpts, sink chan<- *CyctokenMintInvokerChanged, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Cyctoken.contract.WatchLogs(opts, "MintInvokerChanged", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CyctokenMintInvokerChanged)
				if err := _Cyctoken.contract.UnpackLog(event, "MintInvokerChanged", log); err != nil {
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

// ParseMintInvokerChanged is a log parse operation binding the contract event 0xfa319a129fc873627781c98bb447f65d0fc7eeca75f5ca8224de0e2feaffa129.
//
// Solidity: event MintInvokerChanged(address indexed previousOwner, address indexed newOwner)
func (_Cyctoken *CyctokenFilterer) ParseMintInvokerChanged(log types.Log) (*CyctokenMintInvokerChanged, error) {
	event := new(CyctokenMintInvokerChanged)
	if err := _Cyctoken.contract.UnpackLog(event, "MintInvokerChanged", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CyctokenOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Cyctoken contract.
type CyctokenOwnershipTransferredIterator struct {
	Event *CyctokenOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *CyctokenOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CyctokenOwnershipTransferred)
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
		it.Event = new(CyctokenOwnershipTransferred)
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
func (it *CyctokenOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CyctokenOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CyctokenOwnershipTransferred represents a OwnershipTransferred event raised by the Cyctoken contract.
type CyctokenOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Cyctoken *CyctokenFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*CyctokenOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Cyctoken.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CyctokenOwnershipTransferredIterator{contract: _Cyctoken.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Cyctoken *CyctokenFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CyctokenOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Cyctoken.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CyctokenOwnershipTransferred)
				if err := _Cyctoken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Cyctoken *CyctokenFilterer) ParseOwnershipTransferred(log types.Log) (*CyctokenOwnershipTransferred, error) {
	event := new(CyctokenOwnershipTransferred)
	if err := _Cyctoken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CyctokenRebaseIterator is returned from FilterRebase and is used to iterate over the raw logs and unpacked data for Rebase events raised by the Cyctoken contract.
type CyctokenRebaseIterator struct {
	Event *CyctokenRebase // Event containing the contract specifics and raw log

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
func (it *CyctokenRebaseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CyctokenRebase)
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
		it.Event = new(CyctokenRebase)
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
func (it *CyctokenRebaseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CyctokenRebaseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CyctokenRebase represents a Rebase event raised by the Cyctoken contract.
type CyctokenRebase struct {
	Epoch             *big.Int
	OldPerShareAmount *big.Int
	NewPerShareAmount *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRebase is a free log retrieval operation binding the contract event 0xc6642d24d84e7f3d36ca39f5cce10e75639d9b158d5193aa350e2f900653e4c0.
//
// Solidity: event Rebase(uint256 indexed epoch, uint256 oldPerShareAmount, uint256 newPerShareAmount)
func (_Cyctoken *CyctokenFilterer) FilterRebase(opts *bind.FilterOpts, epoch []*big.Int) (*CyctokenRebaseIterator, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _Cyctoken.contract.FilterLogs(opts, "Rebase", epochRule)
	if err != nil {
		return nil, err
	}
	return &CyctokenRebaseIterator{contract: _Cyctoken.contract, event: "Rebase", logs: logs, sub: sub}, nil
}

// WatchRebase is a free log subscription operation binding the contract event 0xc6642d24d84e7f3d36ca39f5cce10e75639d9b158d5193aa350e2f900653e4c0.
//
// Solidity: event Rebase(uint256 indexed epoch, uint256 oldPerShareAmount, uint256 newPerShareAmount)
func (_Cyctoken *CyctokenFilterer) WatchRebase(opts *bind.WatchOpts, sink chan<- *CyctokenRebase, epoch []*big.Int) (event.Subscription, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _Cyctoken.contract.WatchLogs(opts, "Rebase", epochRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CyctokenRebase)
				if err := _Cyctoken.contract.UnpackLog(event, "Rebase", log); err != nil {
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

// ParseRebase is a log parse operation binding the contract event 0xc6642d24d84e7f3d36ca39f5cce10e75639d9b158d5193aa350e2f900653e4c0.
//
// Solidity: event Rebase(uint256 indexed epoch, uint256 oldPerShareAmount, uint256 newPerShareAmount)
func (_Cyctoken *CyctokenFilterer) ParseRebase(log types.Log) (*CyctokenRebase, error) {
	event := new(CyctokenRebase)
	if err := _Cyctoken.contract.UnpackLog(event, "Rebase", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CyctokenRebaseInvokerChangedIterator is returned from FilterRebaseInvokerChanged and is used to iterate over the raw logs and unpacked data for RebaseInvokerChanged events raised by the Cyctoken contract.
type CyctokenRebaseInvokerChangedIterator struct {
	Event *CyctokenRebaseInvokerChanged // Event containing the contract specifics and raw log

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
func (it *CyctokenRebaseInvokerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CyctokenRebaseInvokerChanged)
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
		it.Event = new(CyctokenRebaseInvokerChanged)
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
func (it *CyctokenRebaseInvokerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CyctokenRebaseInvokerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CyctokenRebaseInvokerChanged represents a RebaseInvokerChanged event raised by the Cyctoken contract.
type CyctokenRebaseInvokerChanged struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRebaseInvokerChanged is a free log retrieval operation binding the contract event 0x7b8be1078fc15fa9c531c3600c578a2ce6df5480eb02b7e5ef04e0898091bcea.
//
// Solidity: event RebaseInvokerChanged(address indexed previousOwner, address indexed newOwner)
func (_Cyctoken *CyctokenFilterer) FilterRebaseInvokerChanged(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*CyctokenRebaseInvokerChangedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Cyctoken.contract.FilterLogs(opts, "RebaseInvokerChanged", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CyctokenRebaseInvokerChangedIterator{contract: _Cyctoken.contract, event: "RebaseInvokerChanged", logs: logs, sub: sub}, nil
}

// WatchRebaseInvokerChanged is a free log subscription operation binding the contract event 0x7b8be1078fc15fa9c531c3600c578a2ce6df5480eb02b7e5ef04e0898091bcea.
//
// Solidity: event RebaseInvokerChanged(address indexed previousOwner, address indexed newOwner)
func (_Cyctoken *CyctokenFilterer) WatchRebaseInvokerChanged(opts *bind.WatchOpts, sink chan<- *CyctokenRebaseInvokerChanged, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Cyctoken.contract.WatchLogs(opts, "RebaseInvokerChanged", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CyctokenRebaseInvokerChanged)
				if err := _Cyctoken.contract.UnpackLog(event, "RebaseInvokerChanged", log); err != nil {
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

// ParseRebaseInvokerChanged is a log parse operation binding the contract event 0x7b8be1078fc15fa9c531c3600c578a2ce6df5480eb02b7e5ef04e0898091bcea.
//
// Solidity: event RebaseInvokerChanged(address indexed previousOwner, address indexed newOwner)
func (_Cyctoken *CyctokenFilterer) ParseRebaseInvokerChanged(log types.Log) (*CyctokenRebaseInvokerChanged, error) {
	event := new(CyctokenRebaseInvokerChanged)
	if err := _Cyctoken.contract.UnpackLog(event, "RebaseInvokerChanged", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CyctokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Cyctoken contract.
type CyctokenTransferIterator struct {
	Event *CyctokenTransfer // Event containing the contract specifics and raw log

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
func (it *CyctokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CyctokenTransfer)
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
		it.Event = new(CyctokenTransfer)
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
func (it *CyctokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CyctokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CyctokenTransfer represents a Transfer event raised by the Cyctoken contract.
type CyctokenTransfer struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_Cyctoken *CyctokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CyctokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Cyctoken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CyctokenTransferIterator{contract: _Cyctoken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_Cyctoken *CyctokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *CyctokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Cyctoken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CyctokenTransfer)
				if err := _Cyctoken.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_Cyctoken *CyctokenFilterer) ParseTransfer(log types.Log) (*CyctokenTransfer, error) {
	event := new(CyctokenTransfer)
	if err := _Cyctoken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}
