// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// AlgorithmReviewMetaData contains all meta data concerning the AlgorithmReview contract.
var AlgorithmReviewMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"algorithmCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"algorithms\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"scientist\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"cid\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"dataset\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"startTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"endTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"yesVotes\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"noVotes\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"resolved\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hasVoted\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isCommitteeMember\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"resolve\",\"inputs\":[{\"name\":\"algoId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setCommitteeMember\",\"inputs\":[{\"name\":\"member\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"approved\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submitAlgorithm\",\"inputs\":[{\"name\":\"scientist\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"cid\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"dataset\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"vote\",\"inputs\":[{\"name\":\"algoId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"approve\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AlgorithmResolved\",\"inputs\":[{\"name\":\"algoId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"approved\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AlgorithmSubmitted\",\"inputs\":[{\"name\":\"algoId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"scientist\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"cid\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"VoteCasted\",\"inputs\":[{\"name\":\"algoId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"voter\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"approved\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false}]",
}

// AlgorithmReviewABI is the input ABI used to generate the binding from.
// Deprecated: Use AlgorithmReviewMetaData.ABI instead.
var AlgorithmReviewABI = AlgorithmReviewMetaData.ABI

// AlgorithmReview is an auto generated Go binding around an Ethereum contract.
type AlgorithmReview struct {
	AlgorithmReviewCaller     // Read-only binding to the contract
	AlgorithmReviewTransactor // Write-only binding to the contract
	AlgorithmReviewFilterer   // Log filterer for contract events
}

// AlgorithmReviewCaller is an auto generated read-only Go binding around an Ethereum contract.
type AlgorithmReviewCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AlgorithmReviewTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AlgorithmReviewTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AlgorithmReviewFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AlgorithmReviewFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AlgorithmReviewSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AlgorithmReviewSession struct {
	Contract     *AlgorithmReview  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AlgorithmReviewCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AlgorithmReviewCallerSession struct {
	Contract *AlgorithmReviewCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// AlgorithmReviewTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AlgorithmReviewTransactorSession struct {
	Contract     *AlgorithmReviewTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// AlgorithmReviewRaw is an auto generated low-level Go binding around an Ethereum contract.
type AlgorithmReviewRaw struct {
	Contract *AlgorithmReview // Generic contract binding to access the raw methods on
}

// AlgorithmReviewCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AlgorithmReviewCallerRaw struct {
	Contract *AlgorithmReviewCaller // Generic read-only contract binding to access the raw methods on
}

// AlgorithmReviewTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AlgorithmReviewTransactorRaw struct {
	Contract *AlgorithmReviewTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAlgorithmReview creates a new instance of AlgorithmReview, bound to a specific deployed contract.
func NewAlgorithmReview(address common.Address, backend bind.ContractBackend) (*AlgorithmReview, error) {
	contract, err := bindAlgorithmReview(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AlgorithmReview{AlgorithmReviewCaller: AlgorithmReviewCaller{contract: contract}, AlgorithmReviewTransactor: AlgorithmReviewTransactor{contract: contract}, AlgorithmReviewFilterer: AlgorithmReviewFilterer{contract: contract}}, nil
}

// NewAlgorithmReviewCaller creates a new read-only instance of AlgorithmReview, bound to a specific deployed contract.
func NewAlgorithmReviewCaller(address common.Address, caller bind.ContractCaller) (*AlgorithmReviewCaller, error) {
	contract, err := bindAlgorithmReview(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AlgorithmReviewCaller{contract: contract}, nil
}

// NewAlgorithmReviewTransactor creates a new write-only instance of AlgorithmReview, bound to a specific deployed contract.
func NewAlgorithmReviewTransactor(address common.Address, transactor bind.ContractTransactor) (*AlgorithmReviewTransactor, error) {
	contract, err := bindAlgorithmReview(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AlgorithmReviewTransactor{contract: contract}, nil
}

// NewAlgorithmReviewFilterer creates a new log filterer instance of AlgorithmReview, bound to a specific deployed contract.
func NewAlgorithmReviewFilterer(address common.Address, filterer bind.ContractFilterer) (*AlgorithmReviewFilterer, error) {
	contract, err := bindAlgorithmReview(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AlgorithmReviewFilterer{contract: contract}, nil
}

// bindAlgorithmReview binds a generic wrapper to an already deployed contract.
func bindAlgorithmReview(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AlgorithmReviewMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AlgorithmReview *AlgorithmReviewRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AlgorithmReview.Contract.AlgorithmReviewCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AlgorithmReview *AlgorithmReviewRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AlgorithmReview.Contract.AlgorithmReviewTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AlgorithmReview *AlgorithmReviewRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AlgorithmReview.Contract.AlgorithmReviewTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AlgorithmReview *AlgorithmReviewCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AlgorithmReview.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AlgorithmReview *AlgorithmReviewTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AlgorithmReview.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AlgorithmReview *AlgorithmReviewTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AlgorithmReview.Contract.contract.Transact(opts, method, params...)
}

// AlgorithmCount is a free data retrieval call binding the contract method 0xf7e7d12b.
//
// Solidity: function algorithmCount() view returns(uint256)
func (_AlgorithmReview *AlgorithmReviewCaller) AlgorithmCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AlgorithmReview.contract.Call(opts, &out, "algorithmCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AlgorithmCount is a free data retrieval call binding the contract method 0xf7e7d12b.
//
// Solidity: function algorithmCount() view returns(uint256)
func (_AlgorithmReview *AlgorithmReviewSession) AlgorithmCount() (*big.Int, error) {
	return _AlgorithmReview.Contract.AlgorithmCount(&_AlgorithmReview.CallOpts)
}

// AlgorithmCount is a free data retrieval call binding the contract method 0xf7e7d12b.
//
// Solidity: function algorithmCount() view returns(uint256)
func (_AlgorithmReview *AlgorithmReviewCallerSession) AlgorithmCount() (*big.Int, error) {
	return _AlgorithmReview.Contract.AlgorithmCount(&_AlgorithmReview.CallOpts)
}

// Algorithms is a free data retrieval call binding the contract method 0xb0f16498.
//
// Solidity: function algorithms(uint256 ) view returns(address scientist, string cid, string dataset, uint256 startTime, uint256 endTime, uint256 yesVotes, uint256 noVotes, bool resolved)
func (_AlgorithmReview *AlgorithmReviewCaller) Algorithms(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Scientist common.Address
	Cid       string
	Dataset   string
	StartTime *big.Int
	EndTime   *big.Int
	YesVotes  *big.Int
	NoVotes   *big.Int
	Resolved  bool
}, error) {
	var out []interface{}
	err := _AlgorithmReview.contract.Call(opts, &out, "algorithms", arg0)

	outstruct := new(struct {
		Scientist common.Address
		Cid       string
		Dataset   string
		StartTime *big.Int
		EndTime   *big.Int
		YesVotes  *big.Int
		NoVotes   *big.Int
		Resolved  bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Scientist = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Cid = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Dataset = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.StartTime = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.EndTime = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.YesVotes = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.NoVotes = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.Resolved = *abi.ConvertType(out[7], new(bool)).(*bool)

	return *outstruct, err

}

// Algorithms is a free data retrieval call binding the contract method 0xb0f16498.
//
// Solidity: function algorithms(uint256 ) view returns(address scientist, string cid, string dataset, uint256 startTime, uint256 endTime, uint256 yesVotes, uint256 noVotes, bool resolved)
func (_AlgorithmReview *AlgorithmReviewSession) Algorithms(arg0 *big.Int) (struct {
	Scientist common.Address
	Cid       string
	Dataset   string
	StartTime *big.Int
	EndTime   *big.Int
	YesVotes  *big.Int
	NoVotes   *big.Int
	Resolved  bool
}, error) {
	return _AlgorithmReview.Contract.Algorithms(&_AlgorithmReview.CallOpts, arg0)
}

// Algorithms is a free data retrieval call binding the contract method 0xb0f16498.
//
// Solidity: function algorithms(uint256 ) view returns(address scientist, string cid, string dataset, uint256 startTime, uint256 endTime, uint256 yesVotes, uint256 noVotes, bool resolved)
func (_AlgorithmReview *AlgorithmReviewCallerSession) Algorithms(arg0 *big.Int) (struct {
	Scientist common.Address
	Cid       string
	Dataset   string
	StartTime *big.Int
	EndTime   *big.Int
	YesVotes  *big.Int
	NoVotes   *big.Int
	Resolved  bool
}, error) {
	return _AlgorithmReview.Contract.Algorithms(&_AlgorithmReview.CallOpts, arg0)
}

// HasVoted is a free data retrieval call binding the contract method 0x43859632.
//
// Solidity: function hasVoted(uint256 , address ) view returns(bool)
func (_AlgorithmReview *AlgorithmReviewCaller) HasVoted(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _AlgorithmReview.contract.Call(opts, &out, "hasVoted", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasVoted is a free data retrieval call binding the contract method 0x43859632.
//
// Solidity: function hasVoted(uint256 , address ) view returns(bool)
func (_AlgorithmReview *AlgorithmReviewSession) HasVoted(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _AlgorithmReview.Contract.HasVoted(&_AlgorithmReview.CallOpts, arg0, arg1)
}

// HasVoted is a free data retrieval call binding the contract method 0x43859632.
//
// Solidity: function hasVoted(uint256 , address ) view returns(bool)
func (_AlgorithmReview *AlgorithmReviewCallerSession) HasVoted(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _AlgorithmReview.Contract.HasVoted(&_AlgorithmReview.CallOpts, arg0, arg1)
}

// IsCommitteeMember is a free data retrieval call binding the contract method 0xe636d84b.
//
// Solidity: function isCommitteeMember(address ) view returns(bool)
func (_AlgorithmReview *AlgorithmReviewCaller) IsCommitteeMember(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _AlgorithmReview.contract.Call(opts, &out, "isCommitteeMember", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsCommitteeMember is a free data retrieval call binding the contract method 0xe636d84b.
//
// Solidity: function isCommitteeMember(address ) view returns(bool)
func (_AlgorithmReview *AlgorithmReviewSession) IsCommitteeMember(arg0 common.Address) (bool, error) {
	return _AlgorithmReview.Contract.IsCommitteeMember(&_AlgorithmReview.CallOpts, arg0)
}

// IsCommitteeMember is a free data retrieval call binding the contract method 0xe636d84b.
//
// Solidity: function isCommitteeMember(address ) view returns(bool)
func (_AlgorithmReview *AlgorithmReviewCallerSession) IsCommitteeMember(arg0 common.Address) (bool, error) {
	return _AlgorithmReview.Contract.IsCommitteeMember(&_AlgorithmReview.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AlgorithmReview *AlgorithmReviewCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AlgorithmReview.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AlgorithmReview *AlgorithmReviewSession) Owner() (common.Address, error) {
	return _AlgorithmReview.Contract.Owner(&_AlgorithmReview.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AlgorithmReview *AlgorithmReviewCallerSession) Owner() (common.Address, error) {
	return _AlgorithmReview.Contract.Owner(&_AlgorithmReview.CallOpts)
}

// Resolve is a paid mutator transaction binding the contract method 0x4f896d4f.
//
// Solidity: function resolve(uint256 algoId) returns()
func (_AlgorithmReview *AlgorithmReviewTransactor) Resolve(opts *bind.TransactOpts, algoId *big.Int) (*types.Transaction, error) {
	return _AlgorithmReview.contract.Transact(opts, "resolve", algoId)
}

// Resolve is a paid mutator transaction binding the contract method 0x4f896d4f.
//
// Solidity: function resolve(uint256 algoId) returns()
func (_AlgorithmReview *AlgorithmReviewSession) Resolve(algoId *big.Int) (*types.Transaction, error) {
	return _AlgorithmReview.Contract.Resolve(&_AlgorithmReview.TransactOpts, algoId)
}

// Resolve is a paid mutator transaction binding the contract method 0x4f896d4f.
//
// Solidity: function resolve(uint256 algoId) returns()
func (_AlgorithmReview *AlgorithmReviewTransactorSession) Resolve(algoId *big.Int) (*types.Transaction, error) {
	return _AlgorithmReview.Contract.Resolve(&_AlgorithmReview.TransactOpts, algoId)
}

// SetCommitteeMember is a paid mutator transaction binding the contract method 0xb3608e8a.
//
// Solidity: function setCommitteeMember(address member, bool approved) returns()
func (_AlgorithmReview *AlgorithmReviewTransactor) SetCommitteeMember(opts *bind.TransactOpts, member common.Address, approved bool) (*types.Transaction, error) {
	return _AlgorithmReview.contract.Transact(opts, "setCommitteeMember", member, approved)
}

// SetCommitteeMember is a paid mutator transaction binding the contract method 0xb3608e8a.
//
// Solidity: function setCommitteeMember(address member, bool approved) returns()
func (_AlgorithmReview *AlgorithmReviewSession) SetCommitteeMember(member common.Address, approved bool) (*types.Transaction, error) {
	return _AlgorithmReview.Contract.SetCommitteeMember(&_AlgorithmReview.TransactOpts, member, approved)
}

// SetCommitteeMember is a paid mutator transaction binding the contract method 0xb3608e8a.
//
// Solidity: function setCommitteeMember(address member, bool approved) returns()
func (_AlgorithmReview *AlgorithmReviewTransactorSession) SetCommitteeMember(member common.Address, approved bool) (*types.Transaction, error) {
	return _AlgorithmReview.Contract.SetCommitteeMember(&_AlgorithmReview.TransactOpts, member, approved)
}

// SubmitAlgorithm is a paid mutator transaction binding the contract method 0xaac9c595.
//
// Solidity: function submitAlgorithm(address scientist, string cid, string dataset) returns()
func (_AlgorithmReview *AlgorithmReviewTransactor) SubmitAlgorithm(opts *bind.TransactOpts, scientist common.Address, cid string, dataset string) (*types.Transaction, error) {
	return _AlgorithmReview.contract.Transact(opts, "submitAlgorithm", scientist, cid, dataset)
}

// SubmitAlgorithm is a paid mutator transaction binding the contract method 0xaac9c595.
//
// Solidity: function submitAlgorithm(address scientist, string cid, string dataset) returns()
func (_AlgorithmReview *AlgorithmReviewSession) SubmitAlgorithm(scientist common.Address, cid string, dataset string) (*types.Transaction, error) {
	return _AlgorithmReview.Contract.SubmitAlgorithm(&_AlgorithmReview.TransactOpts, scientist, cid, dataset)
}

// SubmitAlgorithm is a paid mutator transaction binding the contract method 0xaac9c595.
//
// Solidity: function submitAlgorithm(address scientist, string cid, string dataset) returns()
func (_AlgorithmReview *AlgorithmReviewTransactorSession) SubmitAlgorithm(scientist common.Address, cid string, dataset string) (*types.Transaction, error) {
	return _AlgorithmReview.Contract.SubmitAlgorithm(&_AlgorithmReview.TransactOpts, scientist, cid, dataset)
}

// Vote is a paid mutator transaction binding the contract method 0xc9d27afe.
//
// Solidity: function vote(uint256 algoId, bool approve) returns()
func (_AlgorithmReview *AlgorithmReviewTransactor) Vote(opts *bind.TransactOpts, algoId *big.Int, approve bool) (*types.Transaction, error) {
	return _AlgorithmReview.contract.Transact(opts, "vote", algoId, approve)
}

// Vote is a paid mutator transaction binding the contract method 0xc9d27afe.
//
// Solidity: function vote(uint256 algoId, bool approve) returns()
func (_AlgorithmReview *AlgorithmReviewSession) Vote(algoId *big.Int, approve bool) (*types.Transaction, error) {
	return _AlgorithmReview.Contract.Vote(&_AlgorithmReview.TransactOpts, algoId, approve)
}

// Vote is a paid mutator transaction binding the contract method 0xc9d27afe.
//
// Solidity: function vote(uint256 algoId, bool approve) returns()
func (_AlgorithmReview *AlgorithmReviewTransactorSession) Vote(algoId *big.Int, approve bool) (*types.Transaction, error) {
	return _AlgorithmReview.Contract.Vote(&_AlgorithmReview.TransactOpts, algoId, approve)
}

// AlgorithmReviewAlgorithmResolvedIterator is returned from FilterAlgorithmResolved and is used to iterate over the raw logs and unpacked data for AlgorithmResolved events raised by the AlgorithmReview contract.
type AlgorithmReviewAlgorithmResolvedIterator struct {
	Event *AlgorithmReviewAlgorithmResolved // Event containing the contract specifics and raw log

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
func (it *AlgorithmReviewAlgorithmResolvedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AlgorithmReviewAlgorithmResolved)
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
		it.Event = new(AlgorithmReviewAlgorithmResolved)
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
func (it *AlgorithmReviewAlgorithmResolvedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AlgorithmReviewAlgorithmResolvedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AlgorithmReviewAlgorithmResolved represents a AlgorithmResolved event raised by the AlgorithmReview contract.
type AlgorithmReviewAlgorithmResolved struct {
	AlgoId   *big.Int
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAlgorithmResolved is a free log retrieval operation binding the contract event 0xa7e74acaa53738c48ef6f2d293749b18bd9e42ea235b93c110b142efcb76f13a.
//
// Solidity: event AlgorithmResolved(uint256 indexed algoId, bool approved)
func (_AlgorithmReview *AlgorithmReviewFilterer) FilterAlgorithmResolved(opts *bind.FilterOpts, algoId []*big.Int) (*AlgorithmReviewAlgorithmResolvedIterator, error) {

	var algoIdRule []interface{}
	for _, algoIdItem := range algoId {
		algoIdRule = append(algoIdRule, algoIdItem)
	}

	logs, sub, err := _AlgorithmReview.contract.FilterLogs(opts, "AlgorithmResolved", algoIdRule)
	if err != nil {
		return nil, err
	}
	return &AlgorithmReviewAlgorithmResolvedIterator{contract: _AlgorithmReview.contract, event: "AlgorithmResolved", logs: logs, sub: sub}, nil
}

// WatchAlgorithmResolved is a free log subscription operation binding the contract event 0xa7e74acaa53738c48ef6f2d293749b18bd9e42ea235b93c110b142efcb76f13a.
//
// Solidity: event AlgorithmResolved(uint256 indexed algoId, bool approved)
func (_AlgorithmReview *AlgorithmReviewFilterer) WatchAlgorithmResolved(opts *bind.WatchOpts, sink chan<- *AlgorithmReviewAlgorithmResolved, algoId []*big.Int) (event.Subscription, error) {

	var algoIdRule []interface{}
	for _, algoIdItem := range algoId {
		algoIdRule = append(algoIdRule, algoIdItem)
	}

	logs, sub, err := _AlgorithmReview.contract.WatchLogs(opts, "AlgorithmResolved", algoIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AlgorithmReviewAlgorithmResolved)
				if err := _AlgorithmReview.contract.UnpackLog(event, "AlgorithmResolved", log); err != nil {
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

// ParseAlgorithmResolved is a log parse operation binding the contract event 0xa7e74acaa53738c48ef6f2d293749b18bd9e42ea235b93c110b142efcb76f13a.
//
// Solidity: event AlgorithmResolved(uint256 indexed algoId, bool approved)
func (_AlgorithmReview *AlgorithmReviewFilterer) ParseAlgorithmResolved(log types.Log) (*AlgorithmReviewAlgorithmResolved, error) {
	event := new(AlgorithmReviewAlgorithmResolved)
	if err := _AlgorithmReview.contract.UnpackLog(event, "AlgorithmResolved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AlgorithmReviewAlgorithmSubmittedIterator is returned from FilterAlgorithmSubmitted and is used to iterate over the raw logs and unpacked data for AlgorithmSubmitted events raised by the AlgorithmReview contract.
type AlgorithmReviewAlgorithmSubmittedIterator struct {
	Event *AlgorithmReviewAlgorithmSubmitted // Event containing the contract specifics and raw log

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
func (it *AlgorithmReviewAlgorithmSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AlgorithmReviewAlgorithmSubmitted)
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
		it.Event = new(AlgorithmReviewAlgorithmSubmitted)
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
func (it *AlgorithmReviewAlgorithmSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AlgorithmReviewAlgorithmSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AlgorithmReviewAlgorithmSubmitted represents a AlgorithmSubmitted event raised by the AlgorithmReview contract.
type AlgorithmReviewAlgorithmSubmitted struct {
	AlgoId    *big.Int
	Scientist common.Address
	Cid       string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAlgorithmSubmitted is a free log retrieval operation binding the contract event 0x1b4a6be3ab457b503db821aa6606fadb33508bf1df903c08a4d7af9890dbb1c2.
//
// Solidity: event AlgorithmSubmitted(uint256 indexed algoId, address indexed scientist, string cid)
func (_AlgorithmReview *AlgorithmReviewFilterer) FilterAlgorithmSubmitted(opts *bind.FilterOpts, algoId []*big.Int, scientist []common.Address) (*AlgorithmReviewAlgorithmSubmittedIterator, error) {

	var algoIdRule []interface{}
	for _, algoIdItem := range algoId {
		algoIdRule = append(algoIdRule, algoIdItem)
	}
	var scientistRule []interface{}
	for _, scientistItem := range scientist {
		scientistRule = append(scientistRule, scientistItem)
	}

	logs, sub, err := _AlgorithmReview.contract.FilterLogs(opts, "AlgorithmSubmitted", algoIdRule, scientistRule)
	if err != nil {
		return nil, err
	}
	return &AlgorithmReviewAlgorithmSubmittedIterator{contract: _AlgorithmReview.contract, event: "AlgorithmSubmitted", logs: logs, sub: sub}, nil
}

// WatchAlgorithmSubmitted is a free log subscription operation binding the contract event 0x1b4a6be3ab457b503db821aa6606fadb33508bf1df903c08a4d7af9890dbb1c2.
//
// Solidity: event AlgorithmSubmitted(uint256 indexed algoId, address indexed scientist, string cid)
func (_AlgorithmReview *AlgorithmReviewFilterer) WatchAlgorithmSubmitted(opts *bind.WatchOpts, sink chan<- *AlgorithmReviewAlgorithmSubmitted, algoId []*big.Int, scientist []common.Address) (event.Subscription, error) {

	var algoIdRule []interface{}
	for _, algoIdItem := range algoId {
		algoIdRule = append(algoIdRule, algoIdItem)
	}
	var scientistRule []interface{}
	for _, scientistItem := range scientist {
		scientistRule = append(scientistRule, scientistItem)
	}

	logs, sub, err := _AlgorithmReview.contract.WatchLogs(opts, "AlgorithmSubmitted", algoIdRule, scientistRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AlgorithmReviewAlgorithmSubmitted)
				if err := _AlgorithmReview.contract.UnpackLog(event, "AlgorithmSubmitted", log); err != nil {
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

// ParseAlgorithmSubmitted is a log parse operation binding the contract event 0x1b4a6be3ab457b503db821aa6606fadb33508bf1df903c08a4d7af9890dbb1c2.
//
// Solidity: event AlgorithmSubmitted(uint256 indexed algoId, address indexed scientist, string cid)
func (_AlgorithmReview *AlgorithmReviewFilterer) ParseAlgorithmSubmitted(log types.Log) (*AlgorithmReviewAlgorithmSubmitted, error) {
	event := new(AlgorithmReviewAlgorithmSubmitted)
	if err := _AlgorithmReview.contract.UnpackLog(event, "AlgorithmSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AlgorithmReviewVoteCastedIterator is returned from FilterVoteCasted and is used to iterate over the raw logs and unpacked data for VoteCasted events raised by the AlgorithmReview contract.
type AlgorithmReviewVoteCastedIterator struct {
	Event *AlgorithmReviewVoteCasted // Event containing the contract specifics and raw log

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
func (it *AlgorithmReviewVoteCastedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AlgorithmReviewVoteCasted)
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
		it.Event = new(AlgorithmReviewVoteCasted)
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
func (it *AlgorithmReviewVoteCastedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AlgorithmReviewVoteCastedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AlgorithmReviewVoteCasted represents a VoteCasted event raised by the AlgorithmReview contract.
type AlgorithmReviewVoteCasted struct {
	AlgoId   *big.Int
	Voter    common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterVoteCasted is a free log retrieval operation binding the contract event 0x5aaa9aad7433112662b9e5ae23b96ed62b00035f413ab908c55607284e0804e2.
//
// Solidity: event VoteCasted(uint256 indexed algoId, address indexed voter, bool approved)
func (_AlgorithmReview *AlgorithmReviewFilterer) FilterVoteCasted(opts *bind.FilterOpts, algoId []*big.Int, voter []common.Address) (*AlgorithmReviewVoteCastedIterator, error) {

	var algoIdRule []interface{}
	for _, algoIdItem := range algoId {
		algoIdRule = append(algoIdRule, algoIdItem)
	}
	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _AlgorithmReview.contract.FilterLogs(opts, "VoteCasted", algoIdRule, voterRule)
	if err != nil {
		return nil, err
	}
	return &AlgorithmReviewVoteCastedIterator{contract: _AlgorithmReview.contract, event: "VoteCasted", logs: logs, sub: sub}, nil
}

// WatchVoteCasted is a free log subscription operation binding the contract event 0x5aaa9aad7433112662b9e5ae23b96ed62b00035f413ab908c55607284e0804e2.
//
// Solidity: event VoteCasted(uint256 indexed algoId, address indexed voter, bool approved)
func (_AlgorithmReview *AlgorithmReviewFilterer) WatchVoteCasted(opts *bind.WatchOpts, sink chan<- *AlgorithmReviewVoteCasted, algoId []*big.Int, voter []common.Address) (event.Subscription, error) {

	var algoIdRule []interface{}
	for _, algoIdItem := range algoId {
		algoIdRule = append(algoIdRule, algoIdItem)
	}
	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _AlgorithmReview.contract.WatchLogs(opts, "VoteCasted", algoIdRule, voterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AlgorithmReviewVoteCasted)
				if err := _AlgorithmReview.contract.UnpackLog(event, "VoteCasted", log); err != nil {
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

// ParseVoteCasted is a log parse operation binding the contract event 0x5aaa9aad7433112662b9e5ae23b96ed62b00035f413ab908c55607284e0804e2.
//
// Solidity: event VoteCasted(uint256 indexed algoId, address indexed voter, bool approved)
func (_AlgorithmReview *AlgorithmReviewFilterer) ParseVoteCasted(log types.Log) (*AlgorithmReviewVoteCasted, error) {
	event := new(AlgorithmReviewVoteCasted)
	if err := _AlgorithmReview.contract.UnpackLog(event, "VoteCasted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
