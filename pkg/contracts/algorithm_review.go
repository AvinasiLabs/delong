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
	Bin: "0x6080604052348015600e575f5ffd5b50335f5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506115c08061005b5f395ff3fe608060405234801561000f575f5ffd5b5060043610610091575f3560e01c8063b0f1649811610064578063b0f164981461011b578063b3608e8a14610152578063c9d27afe1461016e578063e636d84b1461018a578063f7e7d12b146101ba57610091565b806343859632146100955780634f896d4f146100c55780638da5cb5b146100e1578063aac9c595146100ff575b5f5ffd5b6100af60048036038101906100aa9190610b75565b6101d8565b6040516100bc9190610bcd565b60405180910390f35b6100df60048036038101906100da9190610be6565b610202565b005b6100e9610314565b6040516100f69190610c20565b60405180910390f35b61011960048036038101906101149190610c9a565b610338565b005b61013560048036038101906101309190610be6565b6105e9565b604051610149989796959493929190610daa565b60405180910390f35b61016c60048036038101906101679190610e5e565b610764565b005b61018860048036038101906101839190610e9c565b61084a565b005b6101a4600480360381019061019f9190610eda565b610abd565b6040516101b19190610bcd565b60405180910390f35b6101c2610ada565b6040516101cf9190610f05565b60405180910390f35b6004602052815f5260405f20602052805f5260405f205f915091509054906101000a900460ff1681565b5f60025f8381526020019081526020015f2090508060040154421161025c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161025390610f68565b60405180910390fd5b806007015f9054906101000a900460ff16156102ad576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102a490610fd0565b60405180910390fd5b6001816007015f6101000a81548160ff0219169083151502179055505f81600601548260050154119050827fa7e74acaa53738c48ef6f2d293749b18bd9e42ea235b93c110b142efcb76f13a826040516103079190610bcd565b60405180910390a2505050565b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146103c6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103bd90611038565b60405180910390fd5b5f60015f8154809291906103d990611083565b9190505590506040518061010001604052808773ffffffffffffffffffffffffffffffffffffffff16815260200186868080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f82011690508083019250505050505050815260200184848080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f8201169050808301925050505050505081526020014281526020016203f480426104ae91906110ca565b81526020015f81526020015f81526020015f151581525060025f8381526020019081526020015f205f820151815f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550602082015181600101908161052f9190611327565b5060408201518160020190816105459190611327565b50606082015181600301556080820151816004015560a0820151816005015560c0820151816006015560e0820151816007015f6101000a81548160ff0219169083151502179055509050508573ffffffffffffffffffffffffffffffffffffffff16817f1b4a6be3ab457b503db821aa6606fadb33508bf1df903c08a4d7af9890dbb1c287876040516105d9929190611430565b60405180910390a3505050505050565b6002602052805f5260405f205f91509050805f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169080600101805461062d90611157565b80601f016020809104026020016040519081016040528092919081815260200182805461065990611157565b80156106a45780601f1061067b576101008083540402835291602001916106a4565b820191905f5260205f20905b81548152906001019060200180831161068757829003601f168201915b5050505050908060020180546106b990611157565b80601f01602080910402602001604051908101604052809291908181526020018280546106e590611157565b80156107305780601f1061070757610100808354040283529160200191610730565b820191905f5260205f20905b81548152906001019060200180831161071357829003601f168201915b505050505090806003015490806004015490806005015490806006015490806007015f9054906101000a900460ff16905088565b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146107f2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107e990611038565b60405180910390fd5b8060035f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff0219169083151502179055505050565b60035f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff166108d3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108ca9061149c565b60405180910390fd5b5f60025f8481526020019081526020015f209050806004015442111561092e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161092590611504565b60405180910390fd5b60045f8481526020019081526020015f205f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff16156109c7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109be9061156c565b60405180910390fd5b600160045f8581526020019081526020015f205f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff0219169083151502179055508115610a4f57806005015f815480929190610a4590611083565b9190505550610a69565b806006015f815480929190610a6390611083565b91905055505b3373ffffffffffffffffffffffffffffffffffffffff16837f5aaa9aad7433112662b9e5ae23b96ed62b00035f413ab908c55607284e0804e284604051610ab09190610bcd565b60405180910390a3505050565b6003602052805f5260405f205f915054906101000a900460ff1681565b60015481565b5f5ffd5b5f5ffd5b5f819050919050565b610afa81610ae8565b8114610b04575f5ffd5b50565b5f81359050610b1581610af1565b92915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610b4482610b1b565b9050919050565b610b5481610b3a565b8114610b5e575f5ffd5b50565b5f81359050610b6f81610b4b565b92915050565b5f5f60408385031215610b8b57610b8a610ae0565b5b5f610b9885828601610b07565b9250506020610ba985828601610b61565b9150509250929050565b5f8115159050919050565b610bc781610bb3565b82525050565b5f602082019050610be05f830184610bbe565b92915050565b5f60208284031215610bfb57610bfa610ae0565b5b5f610c0884828501610b07565b91505092915050565b610c1a81610b3a565b82525050565b5f602082019050610c335f830184610c11565b92915050565b5f5ffd5b5f5ffd5b5f5ffd5b5f5f83601f840112610c5a57610c59610c39565b5b8235905067ffffffffffffffff811115610c7757610c76610c3d565b5b602083019150836001820283011115610c9357610c92610c41565b5b9250929050565b5f5f5f5f5f60608688031215610cb357610cb2610ae0565b5b5f610cc088828901610b61565b955050602086013567ffffffffffffffff811115610ce157610ce0610ae4565b5b610ced88828901610c45565b9450945050604086013567ffffffffffffffff811115610d1057610d0f610ae4565b5b610d1c88828901610c45565b92509250509295509295909350565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f610d6d82610d2b565b610d778185610d35565b9350610d87818560208601610d45565b610d9081610d53565b840191505092915050565b610da481610ae8565b82525050565b5f61010082019050610dbe5f83018b610c11565b8181036020830152610dd0818a610d63565b90508181036040830152610de48189610d63565b9050610df36060830188610d9b565b610e006080830187610d9b565b610e0d60a0830186610d9b565b610e1a60c0830185610d9b565b610e2760e0830184610bbe565b9998505050505050505050565b610e3d81610bb3565b8114610e47575f5ffd5b50565b5f81359050610e5881610e34565b92915050565b5f5f60408385031215610e7457610e73610ae0565b5b5f610e8185828601610b61565b9250506020610e9285828601610e4a565b9150509250929050565b5f5f60408385031215610eb257610eb1610ae0565b5b5f610ebf85828601610b07565b9250506020610ed085828601610e4a565b9150509250929050565b5f60208284031215610eef57610eee610ae0565b5b5f610efc84828501610b61565b91505092915050565b5f602082019050610f185f830184610d9b565b92915050565b7f566f74696e67206e6f742079657420656e6465640000000000000000000000005f82015250565b5f610f52601483610d35565b9150610f5d82610f1e565b602082019050919050565b5f6020820190508181035f830152610f7f81610f46565b9050919050565b7f416c7265616479207265736f6c766564000000000000000000000000000000005f82015250565b5f610fba601083610d35565b9150610fc582610f86565b602082019050919050565b5f6020820190508181035f830152610fe781610fae565b9050919050565b7f4e6f74206f776e657200000000000000000000000000000000000000000000005f82015250565b5f611022600983610d35565b915061102d82610fee565b602082019050919050565b5f6020820190508181035f83015261104f81611016565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61108d82610ae8565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036110bf576110be611056565b5b600182019050919050565b5f6110d482610ae8565b91506110df83610ae8565b92508282019050808211156110f7576110f6611056565b5b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061116e57607f821691505b6020821081036111815761118061112a565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026111e37fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826111a8565b6111ed86836111a8565b95508019841693508086168417925050509392505050565b5f819050919050565b5f61122861122361121e84610ae8565b611205565b610ae8565b9050919050565b5f819050919050565b6112418361120e565b61125561124d8261122f565b8484546111b4565b825550505050565b5f5f905090565b61126c61125d565b611277818484611238565b505050565b5b8181101561129a5761128f5f82611264565b60018101905061127d565b5050565b601f8211156112df576112b081611187565b6112b984611199565b810160208510156112c8578190505b6112dc6112d485611199565b83018261127c565b50505b505050565b5f82821c905092915050565b5f6112ff5f19846008026112e4565b1980831691505092915050565b5f61131783836112f0565b9150826002028217905092915050565b61133082610d2b565b67ffffffffffffffff811115611349576113486110fd565b5b6113538254611157565b61135e82828561129e565b5f60209050601f83116001811461138f575f841561137d578287015190505b611387858261130c565b8655506113ee565b601f19841661139d86611187565b5f5b828110156113c45784890151825560018201915060208501945060208101905061139f565b868310156113e157848901516113dd601f8916826112f0565b8355505b6001600288020188555050505b505050505050565b828183375f83830152505050565b5f61140f8385610d35565b935061141c8385846113f6565b61142583610d53565b840190509392505050565b5f6020820190508181035f830152611449818486611404565b90509392505050565b7f4e6f7420636f6d6d6974746565206d656d6265720000000000000000000000005f82015250565b5f611486601483610d35565b915061149182611452565b602082019050919050565b5f6020820190508181035f8301526114b38161147a565b9050919050565b7f566f74696e6720656e64656400000000000000000000000000000000000000005f82015250565b5f6114ee600c83610d35565b91506114f9826114ba565b602082019050919050565b5f6020820190508181035f83015261151b816114e2565b9050919050565b7f416c726561647920766f746564000000000000000000000000000000000000005f82015250565b5f611556600d83610d35565b915061156182611522565b602082019050919050565b5f6020820190508181035f8301526115838161154a565b905091905056fea26469706673582212208ad5c6b7f3fa79a00ae06a6a1b9661f5f5862a58ff96af874ae9f1657640acac64736f6c634300081c0033",
}

// AlgorithmReviewABI is the input ABI used to generate the binding from.
// Deprecated: Use AlgorithmReviewMetaData.ABI instead.
var AlgorithmReviewABI = AlgorithmReviewMetaData.ABI

// AlgorithmReviewBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AlgorithmReviewMetaData.Bin instead.
var AlgorithmReviewBin = AlgorithmReviewMetaData.Bin

// DeployAlgorithmReview deploys a new Ethereum contract, binding an instance of AlgorithmReview to it.
func DeployAlgorithmReview(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AlgorithmReview, error) {
	parsed, err := AlgorithmReviewMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AlgorithmReviewBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AlgorithmReview{AlgorithmReviewCaller: AlgorithmReviewCaller{contract: contract}, AlgorithmReviewTransactor: AlgorithmReviewTransactor{contract: contract}, AlgorithmReviewFilterer: AlgorithmReviewFilterer{contract: contract}}, nil
}

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
