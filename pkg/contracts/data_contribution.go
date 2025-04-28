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

// DataContributionMetaData contains all meta data concerning the DataContribution contract.
var DataContributionMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"recordUsage\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"cid\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"dataset\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerData\",\"inputs\":[{\"name\":\"contributor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"cid\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"dataset\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"DataRegistered\",\"inputs\":[{\"name\":\"contributor\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"cid\",\"type\":\"string\",\"indexed\":true,\"internalType\":\"string\"},{\"name\":\"dataset\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DataUsed\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"cid\",\"type\":\"string\",\"indexed\":true,\"internalType\":\"string\"},{\"name\":\"dataset\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false}]",
}

// DataContributionABI is the input ABI used to generate the binding from.
// Deprecated: Use DataContributionMetaData.ABI instead.
var DataContributionABI = DataContributionMetaData.ABI

// DataContribution is an auto generated Go binding around an Ethereum contract.
type DataContribution struct {
	DataContributionCaller     // Read-only binding to the contract
	DataContributionTransactor // Write-only binding to the contract
	DataContributionFilterer   // Log filterer for contract events
}

// DataContributionCaller is an auto generated read-only Go binding around an Ethereum contract.
type DataContributionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DataContributionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DataContributionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DataContributionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DataContributionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DataContributionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DataContributionSession struct {
	Contract     *DataContribution // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DataContributionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DataContributionCallerSession struct {
	Contract *DataContributionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// DataContributionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DataContributionTransactorSession struct {
	Contract     *DataContributionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// DataContributionRaw is an auto generated low-level Go binding around an Ethereum contract.
type DataContributionRaw struct {
	Contract *DataContribution // Generic contract binding to access the raw methods on
}

// DataContributionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DataContributionCallerRaw struct {
	Contract *DataContributionCaller // Generic read-only contract binding to access the raw methods on
}

// DataContributionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DataContributionTransactorRaw struct {
	Contract *DataContributionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDataContribution creates a new instance of DataContribution, bound to a specific deployed contract.
func NewDataContribution(address common.Address, backend bind.ContractBackend) (*DataContribution, error) {
	contract, err := bindDataContribution(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DataContribution{DataContributionCaller: DataContributionCaller{contract: contract}, DataContributionTransactor: DataContributionTransactor{contract: contract}, DataContributionFilterer: DataContributionFilterer{contract: contract}}, nil
}

// NewDataContributionCaller creates a new read-only instance of DataContribution, bound to a specific deployed contract.
func NewDataContributionCaller(address common.Address, caller bind.ContractCaller) (*DataContributionCaller, error) {
	contract, err := bindDataContribution(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DataContributionCaller{contract: contract}, nil
}

// NewDataContributionTransactor creates a new write-only instance of DataContribution, bound to a specific deployed contract.
func NewDataContributionTransactor(address common.Address, transactor bind.ContractTransactor) (*DataContributionTransactor, error) {
	contract, err := bindDataContribution(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DataContributionTransactor{contract: contract}, nil
}

// NewDataContributionFilterer creates a new log filterer instance of DataContribution, bound to a specific deployed contract.
func NewDataContributionFilterer(address common.Address, filterer bind.ContractFilterer) (*DataContributionFilterer, error) {
	contract, err := bindDataContribution(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DataContributionFilterer{contract: contract}, nil
}

// bindDataContribution binds a generic wrapper to an already deployed contract.
func bindDataContribution(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DataContributionMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DataContribution *DataContributionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DataContribution.Contract.DataContributionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DataContribution *DataContributionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DataContribution.Contract.DataContributionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DataContribution *DataContributionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DataContribution.Contract.DataContributionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DataContribution *DataContributionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DataContribution.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DataContribution *DataContributionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DataContribution.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DataContribution *DataContributionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DataContribution.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DataContribution *DataContributionCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DataContribution.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DataContribution *DataContributionSession) Owner() (common.Address, error) {
	return _DataContribution.Contract.Owner(&_DataContribution.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DataContribution *DataContributionCallerSession) Owner() (common.Address, error) {
	return _DataContribution.Contract.Owner(&_DataContribution.CallOpts)
}

// RecordUsage is a paid mutator transaction binding the contract method 0xf7770429.
//
// Solidity: function recordUsage(address user, string cid, string dataset) returns()
func (_DataContribution *DataContributionTransactor) RecordUsage(opts *bind.TransactOpts, user common.Address, cid string, dataset string) (*types.Transaction, error) {
	return _DataContribution.contract.Transact(opts, "recordUsage", user, cid, dataset)
}

// RecordUsage is a paid mutator transaction binding the contract method 0xf7770429.
//
// Solidity: function recordUsage(address user, string cid, string dataset) returns()
func (_DataContribution *DataContributionSession) RecordUsage(user common.Address, cid string, dataset string) (*types.Transaction, error) {
	return _DataContribution.Contract.RecordUsage(&_DataContribution.TransactOpts, user, cid, dataset)
}

// RecordUsage is a paid mutator transaction binding the contract method 0xf7770429.
//
// Solidity: function recordUsage(address user, string cid, string dataset) returns()
func (_DataContribution *DataContributionTransactorSession) RecordUsage(user common.Address, cid string, dataset string) (*types.Transaction, error) {
	return _DataContribution.Contract.RecordUsage(&_DataContribution.TransactOpts, user, cid, dataset)
}

// RegisterData is a paid mutator transaction binding the contract method 0x322bc451.
//
// Solidity: function registerData(address contributor, string cid, string dataset) returns()
func (_DataContribution *DataContributionTransactor) RegisterData(opts *bind.TransactOpts, contributor common.Address, cid string, dataset string) (*types.Transaction, error) {
	return _DataContribution.contract.Transact(opts, "registerData", contributor, cid, dataset)
}

// RegisterData is a paid mutator transaction binding the contract method 0x322bc451.
//
// Solidity: function registerData(address contributor, string cid, string dataset) returns()
func (_DataContribution *DataContributionSession) RegisterData(contributor common.Address, cid string, dataset string) (*types.Transaction, error) {
	return _DataContribution.Contract.RegisterData(&_DataContribution.TransactOpts, contributor, cid, dataset)
}

// RegisterData is a paid mutator transaction binding the contract method 0x322bc451.
//
// Solidity: function registerData(address contributor, string cid, string dataset) returns()
func (_DataContribution *DataContributionTransactorSession) RegisterData(contributor common.Address, cid string, dataset string) (*types.Transaction, error) {
	return _DataContribution.Contract.RegisterData(&_DataContribution.TransactOpts, contributor, cid, dataset)
}

// DataContributionDataRegisteredIterator is returned from FilterDataRegistered and is used to iterate over the raw logs and unpacked data for DataRegistered events raised by the DataContribution contract.
type DataContributionDataRegisteredIterator struct {
	Event *DataContributionDataRegistered // Event containing the contract specifics and raw log

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
func (it *DataContributionDataRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DataContributionDataRegistered)
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
		it.Event = new(DataContributionDataRegistered)
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
func (it *DataContributionDataRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DataContributionDataRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DataContributionDataRegistered represents a DataRegistered event raised by the DataContribution contract.
type DataContributionDataRegistered struct {
	Contributor common.Address
	Cid         common.Hash
	Dataset     string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDataRegistered is a free log retrieval operation binding the contract event 0xa6f8e79ff0a6a2c15964df60aae8a2c02dd61b6a6c1773d5069970ab3ed007aa.
//
// Solidity: event DataRegistered(address indexed contributor, string indexed cid, string dataset)
func (_DataContribution *DataContributionFilterer) FilterDataRegistered(opts *bind.FilterOpts, contributor []common.Address, cid []string) (*DataContributionDataRegisteredIterator, error) {

	var contributorRule []interface{}
	for _, contributorItem := range contributor {
		contributorRule = append(contributorRule, contributorItem)
	}
	var cidRule []interface{}
	for _, cidItem := range cid {
		cidRule = append(cidRule, cidItem)
	}

	logs, sub, err := _DataContribution.contract.FilterLogs(opts, "DataRegistered", contributorRule, cidRule)
	if err != nil {
		return nil, err
	}
	return &DataContributionDataRegisteredIterator{contract: _DataContribution.contract, event: "DataRegistered", logs: logs, sub: sub}, nil
}

// WatchDataRegistered is a free log subscription operation binding the contract event 0xa6f8e79ff0a6a2c15964df60aae8a2c02dd61b6a6c1773d5069970ab3ed007aa.
//
// Solidity: event DataRegistered(address indexed contributor, string indexed cid, string dataset)
func (_DataContribution *DataContributionFilterer) WatchDataRegistered(opts *bind.WatchOpts, sink chan<- *DataContributionDataRegistered, contributor []common.Address, cid []string) (event.Subscription, error) {

	var contributorRule []interface{}
	for _, contributorItem := range contributor {
		contributorRule = append(contributorRule, contributorItem)
	}
	var cidRule []interface{}
	for _, cidItem := range cid {
		cidRule = append(cidRule, cidItem)
	}

	logs, sub, err := _DataContribution.contract.WatchLogs(opts, "DataRegistered", contributorRule, cidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DataContributionDataRegistered)
				if err := _DataContribution.contract.UnpackLog(event, "DataRegistered", log); err != nil {
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

// ParseDataRegistered is a log parse operation binding the contract event 0xa6f8e79ff0a6a2c15964df60aae8a2c02dd61b6a6c1773d5069970ab3ed007aa.
//
// Solidity: event DataRegistered(address indexed contributor, string indexed cid, string dataset)
func (_DataContribution *DataContributionFilterer) ParseDataRegistered(log types.Log) (*DataContributionDataRegistered, error) {
	event := new(DataContributionDataRegistered)
	if err := _DataContribution.contract.UnpackLog(event, "DataRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DataContributionDataUsedIterator is returned from FilterDataUsed and is used to iterate over the raw logs and unpacked data for DataUsed events raised by the DataContribution contract.
type DataContributionDataUsedIterator struct {
	Event *DataContributionDataUsed // Event containing the contract specifics and raw log

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
func (it *DataContributionDataUsedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DataContributionDataUsed)
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
		it.Event = new(DataContributionDataUsed)
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
func (it *DataContributionDataUsedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DataContributionDataUsedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DataContributionDataUsed represents a DataUsed event raised by the DataContribution contract.
type DataContributionDataUsed struct {
	User    common.Address
	Cid     common.Hash
	Dataset string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDataUsed is a free log retrieval operation binding the contract event 0x71e939dd060653a92275bc37cafacad6ec4aef056145b37275894e92283a3bba.
//
// Solidity: event DataUsed(address indexed user, string indexed cid, string dataset)
func (_DataContribution *DataContributionFilterer) FilterDataUsed(opts *bind.FilterOpts, user []common.Address, cid []string) (*DataContributionDataUsedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var cidRule []interface{}
	for _, cidItem := range cid {
		cidRule = append(cidRule, cidItem)
	}

	logs, sub, err := _DataContribution.contract.FilterLogs(opts, "DataUsed", userRule, cidRule)
	if err != nil {
		return nil, err
	}
	return &DataContributionDataUsedIterator{contract: _DataContribution.contract, event: "DataUsed", logs: logs, sub: sub}, nil
}

// WatchDataUsed is a free log subscription operation binding the contract event 0x71e939dd060653a92275bc37cafacad6ec4aef056145b37275894e92283a3bba.
//
// Solidity: event DataUsed(address indexed user, string indexed cid, string dataset)
func (_DataContribution *DataContributionFilterer) WatchDataUsed(opts *bind.WatchOpts, sink chan<- *DataContributionDataUsed, user []common.Address, cid []string) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var cidRule []interface{}
	for _, cidItem := range cid {
		cidRule = append(cidRule, cidItem)
	}

	logs, sub, err := _DataContribution.contract.WatchLogs(opts, "DataUsed", userRule, cidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DataContributionDataUsed)
				if err := _DataContribution.contract.UnpackLog(event, "DataUsed", log); err != nil {
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

// ParseDataUsed is a log parse operation binding the contract event 0x71e939dd060653a92275bc37cafacad6ec4aef056145b37275894e92283a3bba.
//
// Solidity: event DataUsed(address indexed user, string indexed cid, string dataset)
func (_DataContribution *DataContributionFilterer) ParseDataUsed(log types.Log) (*DataContributionDataUsed, error) {
	event := new(DataContributionDataUsed)
	if err := _DataContribution.contract.UnpackLog(event, "DataUsed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
