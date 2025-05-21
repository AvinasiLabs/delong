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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"algorithmCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"algorithms\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"scientist\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"cid\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"dataset\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"startTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"endTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"yesVotes\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"noVotes\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"resolved\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hasVoted\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isCommitteeMember\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"resolve\",\"inputs\":[{\"name\":\"algoId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setCommitteeMember\",\"inputs\":[{\"name\":\"member\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"approved\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setVotingDuration\",\"inputs\":[{\"name\":\"duration\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submitAlgorithm\",\"inputs\":[{\"name\":\"scientist\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"cid\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"dataset\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"vote\",\"inputs\":[{\"name\":\"algoId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"approve\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"votingDuration\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"AlgorithmResolved\",\"inputs\":[{\"name\":\"algoId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"approved\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AlgorithmSubmitted\",\"inputs\":[{\"name\":\"algoId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"scientist\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"cid\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"startTime\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"endTime\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CommitteeMemberUpdated\",\"inputs\":[{\"name\":\"member\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"approved\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"VoteCasted\",\"inputs\":[{\"name\":\"algoId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"voter\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"approved\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false}]",
	Bin: "0x60806040526203f4806002553480156015575f5ffd5b50335f5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550611728806100625f395ff3fe608060405234801561000f575f5ffd5b50600436106100a7575f3560e01c8063aac9c5951161006f578063aac9c5951461014f578063b0f164981461016b578063b3608e8a146101a2578063c9d27afe146101be578063e636d84b146101da578063f7e7d12b1461020a576100a7565b8063132002fc146100ab57806343859632146100c95780634f896d4f146100f95780635bcfadb5146101155780638da5cb5b14610131575b5f5ffd5b6100b3610228565b6040516100c09190610c44565b60405180910390f35b6100e360048036038101906100de9190610ce9565b61022e565b6040516100f09190610d41565b60405180910390f35b610113600480360381019061010e9190610d5a565b610258565b005b61012f600480360381019061012a9190610d5a565b61036a565b005b610139610402565b6040516101469190610d94565b60405180910390f35b61016960048036038101906101649190610e0e565b610426565b005b61018560048036038101906101809190610d5a565b6106e7565b604051610199989796959493929190610f0f565b60405180910390f35b6101bc60048036038101906101b79190610fc3565b610862565b005b6101d860048036038101906101d39190611001565b610996565b005b6101f460048036038101906101ef919061103f565b610c09565b6040516102019190610d41565b60405180910390f35b610212610c26565b60405161021f9190610c44565b60405180910390f35b60025481565b6005602052815f5260405f20602052805f5260405f205f915091509054906101000a900460ff1681565b5f60035f8381526020019081526020015f209050806004015442116102b2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102a9906110b4565b60405180910390fd5b806007015f9054906101000a900460ff1615610303576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102fa9061111c565b60405180910390fd5b6001816007015f6101000a81548160ff0219169083151502179055505f81600601548260050154119050827fa7e74acaa53738c48ef6f2d293749b18bd9e42ea235b93c110b142efcb76f13a8260405161035d9190610d41565b60405180910390a2505050565b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146103f8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103ef90611184565b60405180910390fd5b8060028190555050565b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146104b4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104ab90611184565b60405180910390fd5b5f60015f8154809291906104c7906111cf565b9190505590506040518061010001604052808773ffffffffffffffffffffffffffffffffffffffff16815260200186868080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f82011690508083019250505050505050815260200184848080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f8201169050808301925050505050505081526020014281526020016002544261059b9190611216565b81526020015f81526020015f81526020015f151581525060035f8381526020019081526020015f205f820151815f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550602082015181600101908161061c9190611473565b5060408201518160020190816106329190611473565b50606082015181600301556080820151816004015560a0820151816005015560c0820151816006015560e0820151816007015f6101000a81548160ff0219169083151502179055509050508573ffffffffffffffffffffffffffffffffffffffff16817fdb7a6630fcaedcf5c4aeb345f20542f8134fc7c9a06800df7a5e38b5f28597c4878742600254426106c79190611216565b6040516106d7949392919061157c565b60405180910390a3505050505050565b6003602052805f5260405f205f91509050805f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169080600101805461072b906112a3565b80601f0160208091040260200160405190810160405280929190818152602001828054610757906112a3565b80156107a25780601f10610779576101008083540402835291602001916107a2565b820191905f5260205f20905b81548152906001019060200180831161078557829003601f168201915b5050505050908060020180546107b7906112a3565b80601f01602080910402602001604051908101604052809291908181526020018280546107e3906112a3565b801561082e5780601f106108055761010080835404028352916020019161082e565b820191905f5260205f20905b81548152906001019060200180831161081157829003601f168201915b505050505090806003015490806004015490806005015490806006015490806007015f9054906101000a900460ff16905088565b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146108f0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108e790611184565b60405180910390fd5b8060045f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff0219169083151502179055508173ffffffffffffffffffffffffffffffffffffffff167f9dcafe810104fa98c663da717c190f40294d51f77f4d41183a01be882b14af038260405161098a9190610d41565b60405180910390a25050565b60045f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff16610a1f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a1690611604565b60405180910390fd5b5f60035f8481526020019081526020015f2090508060040154421115610a7a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a719061166c565b60405180910390fd5b60055f8481526020019081526020015f205f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff1615610b13576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b0a906116d4565b60405180910390fd5b600160055f8581526020019081526020015f205f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff0219169083151502179055508115610b9b57806005015f815480929190610b91906111cf565b9190505550610bb5565b806006015f815480929190610baf906111cf565b91905055505b3373ffffffffffffffffffffffffffffffffffffffff16837f5aaa9aad7433112662b9e5ae23b96ed62b00035f413ab908c55607284e0804e284604051610bfc9190610d41565b60405180910390a3505050565b6004602052805f5260405f205f915054906101000a900460ff1681565b60015481565b5f819050919050565b610c3e81610c2c565b82525050565b5f602082019050610c575f830184610c35565b92915050565b5f5ffd5b5f5ffd5b610c6e81610c2c565b8114610c78575f5ffd5b50565b5f81359050610c8981610c65565b92915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610cb882610c8f565b9050919050565b610cc881610cae565b8114610cd2575f5ffd5b50565b5f81359050610ce381610cbf565b92915050565b5f5f60408385031215610cff57610cfe610c5d565b5b5f610d0c85828601610c7b565b9250506020610d1d85828601610cd5565b9150509250929050565b5f8115159050919050565b610d3b81610d27565b82525050565b5f602082019050610d545f830184610d32565b92915050565b5f60208284031215610d6f57610d6e610c5d565b5b5f610d7c84828501610c7b565b91505092915050565b610d8e81610cae565b82525050565b5f602082019050610da75f830184610d85565b92915050565b5f5ffd5b5f5ffd5b5f5ffd5b5f5f83601f840112610dce57610dcd610dad565b5b8235905067ffffffffffffffff811115610deb57610dea610db1565b5b602083019150836001820283011115610e0757610e06610db5565b5b9250929050565b5f5f5f5f5f60608688031215610e2757610e26610c5d565b5b5f610e3488828901610cd5565b955050602086013567ffffffffffffffff811115610e5557610e54610c61565b5b610e6188828901610db9565b9450945050604086013567ffffffffffffffff811115610e8457610e83610c61565b5b610e9088828901610db9565b92509250509295509295909350565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f610ee182610e9f565b610eeb8185610ea9565b9350610efb818560208601610eb9565b610f0481610ec7565b840191505092915050565b5f61010082019050610f235f83018b610d85565b8181036020830152610f35818a610ed7565b90508181036040830152610f498189610ed7565b9050610f586060830188610c35565b610f656080830187610c35565b610f7260a0830186610c35565b610f7f60c0830185610c35565b610f8c60e0830184610d32565b9998505050505050505050565b610fa281610d27565b8114610fac575f5ffd5b50565b5f81359050610fbd81610f99565b92915050565b5f5f60408385031215610fd957610fd8610c5d565b5b5f610fe685828601610cd5565b9250506020610ff785828601610faf565b9150509250929050565b5f5f6040838503121561101757611016610c5d565b5b5f61102485828601610c7b565b925050602061103585828601610faf565b9150509250929050565b5f6020828403121561105457611053610c5d565b5b5f61106184828501610cd5565b91505092915050565b7f566f74696e67206e6f742079657420656e6465640000000000000000000000005f82015250565b5f61109e601483610ea9565b91506110a98261106a565b602082019050919050565b5f6020820190508181035f8301526110cb81611092565b9050919050565b7f416c7265616479207265736f6c766564000000000000000000000000000000005f82015250565b5f611106601083610ea9565b9150611111826110d2565b602082019050919050565b5f6020820190508181035f830152611133816110fa565b9050919050565b7f4e6f74206f776e657200000000000000000000000000000000000000000000005f82015250565b5f61116e600983610ea9565b91506111798261113a565b602082019050919050565b5f6020820190508181035f83015261119b81611162565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f6111d982610c2c565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361120b5761120a6111a2565b5b600182019050919050565b5f61122082610c2c565b915061122b83610c2c565b9250828201905080821115611243576112426111a2565b5b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60028204905060018216806112ba57607f821691505b6020821081036112cd576112cc611276565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f6008830261132f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826112f4565b61133986836112f4565b95508019841693508086168417925050509392505050565b5f819050919050565b5f61137461136f61136a84610c2c565b611351565b610c2c565b9050919050565b5f819050919050565b61138d8361135a565b6113a16113998261137b565b848454611300565b825550505050565b5f5f905090565b6113b86113a9565b6113c3818484611384565b505050565b5b818110156113e6576113db5f826113b0565b6001810190506113c9565b5050565b601f82111561142b576113fc816112d3565b611405846112e5565b81016020851015611414578190505b611428611420856112e5565b8301826113c8565b50505b505050565b5f82821c905092915050565b5f61144b5f1984600802611430565b1980831691505092915050565b5f611463838361143c565b9150826002028217905092915050565b61147c82610e9f565b67ffffffffffffffff81111561149557611494611249565b5b61149f82546112a3565b6114aa8282856113ea565b5f60209050601f8311600181146114db575f84156114c9578287015190505b6114d38582611458565b86555061153a565b601f1984166114e9866112d3565b5f5b82811015611510578489015182556001820191506020850194506020810190506114eb565b8683101561152d5784890151611529601f89168261143c565b8355505b6001600288020188555050505b505050505050565b828183375f83830152505050565b5f61155b8385610ea9565b9350611568838584611542565b61157183610ec7565b840190509392505050565b5f6060820190508181035f830152611595818688611550565b90506115a46020830185610c35565b6115b16040830184610c35565b95945050505050565b7f4e6f7420636f6d6d6974746565206d656d6265720000000000000000000000005f82015250565b5f6115ee601483610ea9565b91506115f9826115ba565b602082019050919050565b5f6020820190508181035f83015261161b816115e2565b9050919050565b7f566f74696e6720656e64656400000000000000000000000000000000000000005f82015250565b5f611656600c83610ea9565b915061166182611622565b602082019050919050565b5f6020820190508181035f8301526116838161164a565b9050919050565b7f416c726561647920766f746564000000000000000000000000000000000000005f82015250565b5f6116be600d83610ea9565b91506116c98261168a565b602082019050919050565b5f6020820190508181035f8301526116eb816116b2565b905091905056fea26469706673582212202d666e57014b54c0e7acd1ba6dd8366ad6a811bb4fc717dffff06b95fa9d98eb64736f6c634300081c0033",
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

// VotingDuration is a free data retrieval call binding the contract method 0x132002fc.
//
// Solidity: function votingDuration() view returns(uint256)
func (_AlgorithmReview *AlgorithmReviewCaller) VotingDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AlgorithmReview.contract.Call(opts, &out, "votingDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VotingDuration is a free data retrieval call binding the contract method 0x132002fc.
//
// Solidity: function votingDuration() view returns(uint256)
func (_AlgorithmReview *AlgorithmReviewSession) VotingDuration() (*big.Int, error) {
	return _AlgorithmReview.Contract.VotingDuration(&_AlgorithmReview.CallOpts)
}

// VotingDuration is a free data retrieval call binding the contract method 0x132002fc.
//
// Solidity: function votingDuration() view returns(uint256)
func (_AlgorithmReview *AlgorithmReviewCallerSession) VotingDuration() (*big.Int, error) {
	return _AlgorithmReview.Contract.VotingDuration(&_AlgorithmReview.CallOpts)
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

// SetVotingDuration is a paid mutator transaction binding the contract method 0x5bcfadb5.
//
// Solidity: function setVotingDuration(uint256 duration) returns()
func (_AlgorithmReview *AlgorithmReviewTransactor) SetVotingDuration(opts *bind.TransactOpts, duration *big.Int) (*types.Transaction, error) {
	return _AlgorithmReview.contract.Transact(opts, "setVotingDuration", duration)
}

// SetVotingDuration is a paid mutator transaction binding the contract method 0x5bcfadb5.
//
// Solidity: function setVotingDuration(uint256 duration) returns()
func (_AlgorithmReview *AlgorithmReviewSession) SetVotingDuration(duration *big.Int) (*types.Transaction, error) {
	return _AlgorithmReview.Contract.SetVotingDuration(&_AlgorithmReview.TransactOpts, duration)
}

// SetVotingDuration is a paid mutator transaction binding the contract method 0x5bcfadb5.
//
// Solidity: function setVotingDuration(uint256 duration) returns()
func (_AlgorithmReview *AlgorithmReviewTransactorSession) SetVotingDuration(duration *big.Int) (*types.Transaction, error) {
	return _AlgorithmReview.Contract.SetVotingDuration(&_AlgorithmReview.TransactOpts, duration)
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
	StartTime *big.Int
	EndTime   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAlgorithmSubmitted is a free log retrieval operation binding the contract event 0xdb7a6630fcaedcf5c4aeb345f20542f8134fc7c9a06800df7a5e38b5f28597c4.
//
// Solidity: event AlgorithmSubmitted(uint256 indexed algoId, address indexed scientist, string cid, uint256 startTime, uint256 endTime)
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

// WatchAlgorithmSubmitted is a free log subscription operation binding the contract event 0xdb7a6630fcaedcf5c4aeb345f20542f8134fc7c9a06800df7a5e38b5f28597c4.
//
// Solidity: event AlgorithmSubmitted(uint256 indexed algoId, address indexed scientist, string cid, uint256 startTime, uint256 endTime)
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

// ParseAlgorithmSubmitted is a log parse operation binding the contract event 0xdb7a6630fcaedcf5c4aeb345f20542f8134fc7c9a06800df7a5e38b5f28597c4.
//
// Solidity: event AlgorithmSubmitted(uint256 indexed algoId, address indexed scientist, string cid, uint256 startTime, uint256 endTime)
func (_AlgorithmReview *AlgorithmReviewFilterer) ParseAlgorithmSubmitted(log types.Log) (*AlgorithmReviewAlgorithmSubmitted, error) {
	event := new(AlgorithmReviewAlgorithmSubmitted)
	if err := _AlgorithmReview.contract.UnpackLog(event, "AlgorithmSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AlgorithmReviewCommitteeMemberUpdatedIterator is returned from FilterCommitteeMemberUpdated and is used to iterate over the raw logs and unpacked data for CommitteeMemberUpdated events raised by the AlgorithmReview contract.
type AlgorithmReviewCommitteeMemberUpdatedIterator struct {
	Event *AlgorithmReviewCommitteeMemberUpdated // Event containing the contract specifics and raw log

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
func (it *AlgorithmReviewCommitteeMemberUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AlgorithmReviewCommitteeMemberUpdated)
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
		it.Event = new(AlgorithmReviewCommitteeMemberUpdated)
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
func (it *AlgorithmReviewCommitteeMemberUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AlgorithmReviewCommitteeMemberUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AlgorithmReviewCommitteeMemberUpdated represents a CommitteeMemberUpdated event raised by the AlgorithmReview contract.
type AlgorithmReviewCommitteeMemberUpdated struct {
	Member   common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCommitteeMemberUpdated is a free log retrieval operation binding the contract event 0x9dcafe810104fa98c663da717c190f40294d51f77f4d41183a01be882b14af03.
//
// Solidity: event CommitteeMemberUpdated(address indexed member, bool approved)
func (_AlgorithmReview *AlgorithmReviewFilterer) FilterCommitteeMemberUpdated(opts *bind.FilterOpts, member []common.Address) (*AlgorithmReviewCommitteeMemberUpdatedIterator, error) {

	var memberRule []interface{}
	for _, memberItem := range member {
		memberRule = append(memberRule, memberItem)
	}

	logs, sub, err := _AlgorithmReview.contract.FilterLogs(opts, "CommitteeMemberUpdated", memberRule)
	if err != nil {
		return nil, err
	}
	return &AlgorithmReviewCommitteeMemberUpdatedIterator{contract: _AlgorithmReview.contract, event: "CommitteeMemberUpdated", logs: logs, sub: sub}, nil
}

// WatchCommitteeMemberUpdated is a free log subscription operation binding the contract event 0x9dcafe810104fa98c663da717c190f40294d51f77f4d41183a01be882b14af03.
//
// Solidity: event CommitteeMemberUpdated(address indexed member, bool approved)
func (_AlgorithmReview *AlgorithmReviewFilterer) WatchCommitteeMemberUpdated(opts *bind.WatchOpts, sink chan<- *AlgorithmReviewCommitteeMemberUpdated, member []common.Address) (event.Subscription, error) {

	var memberRule []interface{}
	for _, memberItem := range member {
		memberRule = append(memberRule, memberItem)
	}

	logs, sub, err := _AlgorithmReview.contract.WatchLogs(opts, "CommitteeMemberUpdated", memberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AlgorithmReviewCommitteeMemberUpdated)
				if err := _AlgorithmReview.contract.UnpackLog(event, "CommitteeMemberUpdated", log); err != nil {
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

// ParseCommitteeMemberUpdated is a log parse operation binding the contract event 0x9dcafe810104fa98c663da717c190f40294d51f77f4d41183a01be882b14af03.
//
// Solidity: event CommitteeMemberUpdated(address indexed member, bool approved)
func (_AlgorithmReview *AlgorithmReviewFilterer) ParseCommitteeMemberUpdated(log types.Log) (*AlgorithmReviewCommitteeMemberUpdated, error) {
	event := new(AlgorithmReviewCommitteeMemberUpdated)
	if err := _AlgorithmReview.contract.UnpackLog(event, "CommitteeMemberUpdated", log); err != nil {
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
