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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"executionCounter\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"executions\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"scientist\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"cid\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"dataset\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hasVoted\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isCommitteeMember\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"resolve\",\"inputs\":[{\"name\":\"cid\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"executionId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setCommitteeMember\",\"inputs\":[{\"name\":\"member\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"approved\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setVotingDuration\",\"inputs\":[{\"name\":\"duration\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submitAlgorithm\",\"inputs\":[{\"name\":\"executionId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"scientist\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"cid\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"dataset\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"vote\",\"inputs\":[{\"name\":\"cid\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"approve\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"votes\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"yesVotes\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"noVotes\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"resolved\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"approved\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"startTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"endTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"votingDuration\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"AlgorithmResolved\",\"inputs\":[{\"name\":\"executionId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"cid\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"approved\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CommitteeMemberUpdated\",\"inputs\":[{\"name\":\"member\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"approved\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ExecutionSubmitted\",\"inputs\":[{\"name\":\"executionId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"cid\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"startTime\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"endTime\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"VoteCasted\",\"inputs\":[{\"name\":\"member\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"cid\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"approved\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"voteTime\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x60806040526203f4806001553480156015575f5ffd5b50335f5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506119e4806100625f395ff3fe608060405234801561000f575f5ffd5b50600436106100b2575f3560e01c8063aadc3b721161006f578063aadc3b721461017b578063b3608e8a146101ab578063c6a50809146101c7578063e636d84b146101e5578063f76c922914610215578063f9c2af8614610248576100b2565b8063132002fc146100b657806327d3410c146100d45780632b38cd96146100f05780635bcfadb51461012557806375ecf3b1146101415780638da5cb5b1461015d575b5f5ffd5b6100be610264565b6040516100cb9190610d3a565b60405180910390f35b6100ee60048036038101906100e99190610df1565b61026a565b005b61010a60048036038101906101059190610e81565b610595565b60405161011c96959493929190610ebb565b60405180910390f35b61013f600480360381019061013a9190610f44565b6105e6565b005b61015b60048036038101906101569190610fc9565b61067e565b005b610165610906565b604051610172919061107b565b60405180910390f35b61019560048036038101906101909190611094565b61092a565b6040516101a291906110d2565b60405180910390f35b6101c560048036038101906101c091906110eb565b610954565b005b6101cf610a88565b6040516101dc9190610d3a565b60405180910390f35b6101ff60048036038101906101fa9190611129565b610a8e565b60405161020c91906110d2565b60405180910390f35b61022f600480360381019061022a9190610f44565b610aab565b60405161023f94939291906111c4565b60405180910390f35b610262600480360381019061025d9190611215565b610c02565b005b60015481565b60045f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff166102f3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102ea906112bc565b60405180910390fd5b5f8383604051610304929190611316565b604051809103902090505f60065f8381526020019081526020015f2090505f816003015403610368576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161035f90611378565b60405180910390fd5b806002015f9054906101000a900460ff16156103b9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103b0906113e0565b60405180910390fd5b8060040154421115610400576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103f790611448565b60405180910390fd5b60035f8381526020019081526020015f205f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff1615610499576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610490906114b0565b60405180910390fd5b600160035f8481526020019081526020015f205f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff021916908315150217905550821561052057805f015f815480929190610516906114fb565b919050555061053a565b806001015f815480929190610534906114fb565b91905055505b3373ffffffffffffffffffffffffffffffffffffffff167f5128f8fec9b40ca4b18527897904755cac057c02f346595fc6efd08c3bbf1cd786868642604051610586949392919061156e565b60405180910390a25050505050565b6006602052805f5260405f205f91509050805f015490806001015490806002015f9054906101000a900460ff16908060020160019054906101000a900460ff16908060030154908060040154905086565b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610674576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161066b906115f6565b60405180910390fd5b8060018190555050565b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461070c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610703906115f6565b60405180910390fd5b5f848460405161071d929190611316565b604051809103902090505f42905060405180608001604052808981526020018873ffffffffffffffffffffffffffffffffffffffff16815260200187878080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f82011690508083019250505050505050815260200185858080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f8201169050808301925050505050505081525060055f8a81526020019081526020015f205f820151815f01556020820151816001015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550604082015181600201908161085c919061183e565b506060820151816003019081610872919061183e565b509050505f60065f8481526020019081526020015f2090505f8160030154036108b557818160030181905550600154826108ac919061190d565b81600401819055505b887f63ef969e1ddfb70ebfc0b7094e9d170057766c20457c8e1a9a8abc310f7871a18888846003015485600401546040516108f39493929190611940565b60405180910390a2505050505050505050565b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6003602052815f5260405f20602052805f5260405f205f915091509054906101000a900460ff1681565b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146109e2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109d9906115f6565b60405180910390fd5b8060045f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff0219169083151502179055508173ffffffffffffffffffffffffffffffffffffffff167f9dcafe810104fa98c663da717c190f40294d51f77f4d41183a01be882b14af0382604051610a7c91906110d2565b60405180910390a25050565b60025481565b6004602052805f5260405f205f915054906101000a900460ff1681565b6005602052805f5260405f205f91509050805f015490806001015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806002018054610af59061166e565b80601f0160208091040260200160405190810160405280929190818152602001828054610b219061166e565b8015610b6c5780601f10610b4357610100808354040283529160200191610b6c565b820191905f5260205f20905b815481529060010190602001808311610b4f57829003601f168201915b505050505090806003018054610b819061166e565b80601f0160208091040260200160405190810160405280929190818152602001828054610bad9061166e565b8015610bf85780601f10610bcf57610100808354040283529160200191610bf8565b820191905f5260205f20905b815481529060010190602001808311610bdb57829003601f168201915b5050505050905084565b5f8383604051610c13929190611316565b604051809103902090505f60065f8381526020019081526020015f2090505f816003015403610c77576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c6e90611378565b60405180910390fd5b806002015f9054906101000a900460ff16610cce576001816002015f6101000a81548160ff0219169083151502179055508060010154815f0154118160020160016101000a81548160ff0219169083151502179055505b827f4581a0d468af4916140747e3232ba054684a03025d3aa7698679cb0d653413f886868460020160019054906101000a900460ff16604051610d139392919061197e565b60405180910390a25050505050565b5f819050919050565b610d3481610d22565b82525050565b5f602082019050610d4d5f830184610d2b565b92915050565b5f5ffd5b5f5ffd5b5f5ffd5b5f5ffd5b5f5ffd5b5f5f83601f840112610d7c57610d7b610d5b565b5b8235905067ffffffffffffffff811115610d9957610d98610d5f565b5b602083019150836001820283011115610db557610db4610d63565b5b9250929050565b5f8115159050919050565b610dd081610dbc565b8114610dda575f5ffd5b50565b5f81359050610deb81610dc7565b92915050565b5f5f5f60408486031215610e0857610e07610d53565b5b5f84013567ffffffffffffffff811115610e2557610e24610d57565b5b610e3186828701610d67565b93509350506020610e4486828701610ddd565b9150509250925092565b5f819050919050565b610e6081610e4e565b8114610e6a575f5ffd5b50565b5f81359050610e7b81610e57565b92915050565b5f60208284031215610e9657610e95610d53565b5b5f610ea384828501610e6d565b91505092915050565b610eb581610dbc565b82525050565b5f60c082019050610ece5f830189610d2b565b610edb6020830188610d2b565b610ee86040830187610eac565b610ef56060830186610eac565b610f026080830185610d2b565b610f0f60a0830184610d2b565b979650505050505050565b610f2381610d22565b8114610f2d575f5ffd5b50565b5f81359050610f3e81610f1a565b92915050565b5f60208284031215610f5957610f58610d53565b5b5f610f6684828501610f30565b91505092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610f9882610f6f565b9050919050565b610fa881610f8e565b8114610fb2575f5ffd5b50565b5f81359050610fc381610f9f565b92915050565b5f5f5f5f5f5f60808789031215610fe357610fe2610d53565b5b5f610ff089828a01610f30565b965050602061100189828a01610fb5565b955050604087013567ffffffffffffffff81111561102257611021610d57565b5b61102e89828a01610d67565b9450945050606087013567ffffffffffffffff81111561105157611050610d57565b5b61105d89828a01610d67565b92509250509295509295509295565b61107581610f8e565b82525050565b5f60208201905061108e5f83018461106c565b92915050565b5f5f604083850312156110aa576110a9610d53565b5b5f6110b785828601610e6d565b92505060206110c885828601610fb5565b9150509250929050565b5f6020820190506110e55f830184610eac565b92915050565b5f5f6040838503121561110157611100610d53565b5b5f61110e85828601610fb5565b925050602061111f85828601610ddd565b9150509250929050565b5f6020828403121561113e5761113d610d53565b5b5f61114b84828501610fb5565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f61119682611154565b6111a0818561115e565b93506111b081856020860161116e565b6111b98161117c565b840191505092915050565b5f6080820190506111d75f830187610d2b565b6111e4602083018661106c565b81810360408301526111f6818561118c565b9050818103606083015261120a818461118c565b905095945050505050565b5f5f5f6040848603121561122c5761122b610d53565b5b5f84013567ffffffffffffffff81111561124957611248610d57565b5b61125586828701610d67565b9350935050602061126886828701610f30565b9150509250925092565b7f4e6f7420636f6d6d6974746565206d656d6265720000000000000000000000005f82015250565b5f6112a660148361115e565b91506112b182611272565b602082019050919050565b5f6020820190508181035f8301526112d38161129a565b9050919050565b5f81905092915050565b828183375f83830152505050565b5f6112fd83856112da565b935061130a8385846112e4565b82840190509392505050565b5f6113228284866112f2565b91508190509392505050565b7f416c676f726974686d206e6f7420666f756e64000000000000000000000000005f82015250565b5f61136260138361115e565b915061136d8261132e565b602082019050919050565b5f6020820190508181035f83015261138f81611356565b9050919050565b7f416c7265616479207265736f6c766564000000000000000000000000000000005f82015250565b5f6113ca60108361115e565b91506113d582611396565b602082019050919050565b5f6020820190508181035f8301526113f7816113be565b9050919050565b7f566f74696e672068617320656e646564000000000000000000000000000000005f82015250565b5f61143260108361115e565b915061143d826113fe565b602082019050919050565b5f6020820190508181035f83015261145f81611426565b9050919050565b7f416c726561647920766f746564000000000000000000000000000000000000005f82015250565b5f61149a600d8361115e565b91506114a582611466565b602082019050919050565b5f6020820190508181035f8301526114c78161148e565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61150582610d22565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611537576115366114ce565b5b600182019050919050565b5f61154d838561115e565b935061155a8385846112e4565b6115638361117c565b840190509392505050565b5f6060820190508181035f830152611587818688611542565b90506115966020830185610eac565b6115a36040830184610d2b565b95945050505050565b7f4e6f74206f776e657200000000000000000000000000000000000000000000005f82015250565b5f6115e060098361115e565b91506115eb826115ac565b602082019050919050565b5f6020820190508181035f83015261160d816115d4565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061168557607f821691505b60208210810361169857611697611641565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026116fa7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826116bf565b61170486836116bf565b95508019841693508086168417925050509392505050565b5f819050919050565b5f61173f61173a61173584610d22565b61171c565b610d22565b9050919050565b5f819050919050565b61175883611725565b61176c61176482611746565b8484546116cb565b825550505050565b5f5f905090565b611783611774565b61178e81848461174f565b505050565b5b818110156117b1576117a65f8261177b565b600181019050611794565b5050565b601f8211156117f6576117c78161169e565b6117d0846116b0565b810160208510156117df578190505b6117f36117eb856116b0565b830182611793565b50505b505050565b5f82821c905092915050565b5f6118165f19846008026117fb565b1980831691505092915050565b5f61182e8383611807565b9150826002028217905092915050565b61184782611154565b67ffffffffffffffff8111156118605761185f611614565b5b61186a825461166e565b6118758282856117b5565b5f60209050601f8311600181146118a6575f8415611894578287015190505b61189e8582611823565b865550611905565b601f1984166118b48661169e565b5f5b828110156118db578489015182556001820191506020850194506020810190506118b6565b868310156118f857848901516118f4601f891682611807565b8355505b6001600288020188555050505b505050505050565b5f61191782610d22565b915061192283610d22565b925082820190508082111561193a576119396114ce565b5b92915050565b5f6060820190508181035f830152611959818688611542565b90506119686020830185610d2b565b6119756040830184610d2b565b95945050505050565b5f6040820190508181035f830152611997818587611542565b90506119a66020830184610eac565b94935050505056fea2646970667358221220bf8abf83ff6aa8b77f02217fcbe76a4f6b7b27732a7f18852edbfa2c03760f0164736f6c634300081c0033",
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

// ExecutionCounter is a free data retrieval call binding the contract method 0xc6a50809.
//
// Solidity: function executionCounter() view returns(uint256)
func (_AlgorithmReview *AlgorithmReviewCaller) ExecutionCounter(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AlgorithmReview.contract.Call(opts, &out, "executionCounter")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExecutionCounter is a free data retrieval call binding the contract method 0xc6a50809.
//
// Solidity: function executionCounter() view returns(uint256)
func (_AlgorithmReview *AlgorithmReviewSession) ExecutionCounter() (*big.Int, error) {
	return _AlgorithmReview.Contract.ExecutionCounter(&_AlgorithmReview.CallOpts)
}

// ExecutionCounter is a free data retrieval call binding the contract method 0xc6a50809.
//
// Solidity: function executionCounter() view returns(uint256)
func (_AlgorithmReview *AlgorithmReviewCallerSession) ExecutionCounter() (*big.Int, error) {
	return _AlgorithmReview.Contract.ExecutionCounter(&_AlgorithmReview.CallOpts)
}

// Executions is a free data retrieval call binding the contract method 0xf76c9229.
//
// Solidity: function executions(uint256 ) view returns(uint256 id, address scientist, string cid, string dataset)
func (_AlgorithmReview *AlgorithmReviewCaller) Executions(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Id        *big.Int
	Scientist common.Address
	Cid       string
	Dataset   string
}, error) {
	var out []interface{}
	err := _AlgorithmReview.contract.Call(opts, &out, "executions", arg0)

	outstruct := new(struct {
		Id        *big.Int
		Scientist common.Address
		Cid       string
		Dataset   string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Scientist = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Cid = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Dataset = *abi.ConvertType(out[3], new(string)).(*string)

	return *outstruct, err

}

// Executions is a free data retrieval call binding the contract method 0xf76c9229.
//
// Solidity: function executions(uint256 ) view returns(uint256 id, address scientist, string cid, string dataset)
func (_AlgorithmReview *AlgorithmReviewSession) Executions(arg0 *big.Int) (struct {
	Id        *big.Int
	Scientist common.Address
	Cid       string
	Dataset   string
}, error) {
	return _AlgorithmReview.Contract.Executions(&_AlgorithmReview.CallOpts, arg0)
}

// Executions is a free data retrieval call binding the contract method 0xf76c9229.
//
// Solidity: function executions(uint256 ) view returns(uint256 id, address scientist, string cid, string dataset)
func (_AlgorithmReview *AlgorithmReviewCallerSession) Executions(arg0 *big.Int) (struct {
	Id        *big.Int
	Scientist common.Address
	Cid       string
	Dataset   string
}, error) {
	return _AlgorithmReview.Contract.Executions(&_AlgorithmReview.CallOpts, arg0)
}

// HasVoted is a free data retrieval call binding the contract method 0xaadc3b72.
//
// Solidity: function hasVoted(bytes32 , address ) view returns(bool)
func (_AlgorithmReview *AlgorithmReviewCaller) HasVoted(opts *bind.CallOpts, arg0 [32]byte, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _AlgorithmReview.contract.Call(opts, &out, "hasVoted", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasVoted is a free data retrieval call binding the contract method 0xaadc3b72.
//
// Solidity: function hasVoted(bytes32 , address ) view returns(bool)
func (_AlgorithmReview *AlgorithmReviewSession) HasVoted(arg0 [32]byte, arg1 common.Address) (bool, error) {
	return _AlgorithmReview.Contract.HasVoted(&_AlgorithmReview.CallOpts, arg0, arg1)
}

// HasVoted is a free data retrieval call binding the contract method 0xaadc3b72.
//
// Solidity: function hasVoted(bytes32 , address ) view returns(bool)
func (_AlgorithmReview *AlgorithmReviewCallerSession) HasVoted(arg0 [32]byte, arg1 common.Address) (bool, error) {
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

// Votes is a free data retrieval call binding the contract method 0x2b38cd96.
//
// Solidity: function votes(bytes32 ) view returns(uint256 yesVotes, uint256 noVotes, bool resolved, bool approved, uint256 startTime, uint256 endTime)
func (_AlgorithmReview *AlgorithmReviewCaller) Votes(opts *bind.CallOpts, arg0 [32]byte) (struct {
	YesVotes  *big.Int
	NoVotes   *big.Int
	Resolved  bool
	Approved  bool
	StartTime *big.Int
	EndTime   *big.Int
}, error) {
	var out []interface{}
	err := _AlgorithmReview.contract.Call(opts, &out, "votes", arg0)

	outstruct := new(struct {
		YesVotes  *big.Int
		NoVotes   *big.Int
		Resolved  bool
		Approved  bool
		StartTime *big.Int
		EndTime   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.YesVotes = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.NoVotes = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Resolved = *abi.ConvertType(out[2], new(bool)).(*bool)
	outstruct.Approved = *abi.ConvertType(out[3], new(bool)).(*bool)
	outstruct.StartTime = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.EndTime = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Votes is a free data retrieval call binding the contract method 0x2b38cd96.
//
// Solidity: function votes(bytes32 ) view returns(uint256 yesVotes, uint256 noVotes, bool resolved, bool approved, uint256 startTime, uint256 endTime)
func (_AlgorithmReview *AlgorithmReviewSession) Votes(arg0 [32]byte) (struct {
	YesVotes  *big.Int
	NoVotes   *big.Int
	Resolved  bool
	Approved  bool
	StartTime *big.Int
	EndTime   *big.Int
}, error) {
	return _AlgorithmReview.Contract.Votes(&_AlgorithmReview.CallOpts, arg0)
}

// Votes is a free data retrieval call binding the contract method 0x2b38cd96.
//
// Solidity: function votes(bytes32 ) view returns(uint256 yesVotes, uint256 noVotes, bool resolved, bool approved, uint256 startTime, uint256 endTime)
func (_AlgorithmReview *AlgorithmReviewCallerSession) Votes(arg0 [32]byte) (struct {
	YesVotes  *big.Int
	NoVotes   *big.Int
	Resolved  bool
	Approved  bool
	StartTime *big.Int
	EndTime   *big.Int
}, error) {
	return _AlgorithmReview.Contract.Votes(&_AlgorithmReview.CallOpts, arg0)
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

// Resolve is a paid mutator transaction binding the contract method 0xf9c2af86.
//
// Solidity: function resolve(string cid, uint256 executionId) returns()
func (_AlgorithmReview *AlgorithmReviewTransactor) Resolve(opts *bind.TransactOpts, cid string, executionId *big.Int) (*types.Transaction, error) {
	return _AlgorithmReview.contract.Transact(opts, "resolve", cid, executionId)
}

// Resolve is a paid mutator transaction binding the contract method 0xf9c2af86.
//
// Solidity: function resolve(string cid, uint256 executionId) returns()
func (_AlgorithmReview *AlgorithmReviewSession) Resolve(cid string, executionId *big.Int) (*types.Transaction, error) {
	return _AlgorithmReview.Contract.Resolve(&_AlgorithmReview.TransactOpts, cid, executionId)
}

// Resolve is a paid mutator transaction binding the contract method 0xf9c2af86.
//
// Solidity: function resolve(string cid, uint256 executionId) returns()
func (_AlgorithmReview *AlgorithmReviewTransactorSession) Resolve(cid string, executionId *big.Int) (*types.Transaction, error) {
	return _AlgorithmReview.Contract.Resolve(&_AlgorithmReview.TransactOpts, cid, executionId)
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

// SubmitAlgorithm is a paid mutator transaction binding the contract method 0x75ecf3b1.
//
// Solidity: function submitAlgorithm(uint256 executionId, address scientist, string cid, string dataset) returns()
func (_AlgorithmReview *AlgorithmReviewTransactor) SubmitAlgorithm(opts *bind.TransactOpts, executionId *big.Int, scientist common.Address, cid string, dataset string) (*types.Transaction, error) {
	return _AlgorithmReview.contract.Transact(opts, "submitAlgorithm", executionId, scientist, cid, dataset)
}

// SubmitAlgorithm is a paid mutator transaction binding the contract method 0x75ecf3b1.
//
// Solidity: function submitAlgorithm(uint256 executionId, address scientist, string cid, string dataset) returns()
func (_AlgorithmReview *AlgorithmReviewSession) SubmitAlgorithm(executionId *big.Int, scientist common.Address, cid string, dataset string) (*types.Transaction, error) {
	return _AlgorithmReview.Contract.SubmitAlgorithm(&_AlgorithmReview.TransactOpts, executionId, scientist, cid, dataset)
}

// SubmitAlgorithm is a paid mutator transaction binding the contract method 0x75ecf3b1.
//
// Solidity: function submitAlgorithm(uint256 executionId, address scientist, string cid, string dataset) returns()
func (_AlgorithmReview *AlgorithmReviewTransactorSession) SubmitAlgorithm(executionId *big.Int, scientist common.Address, cid string, dataset string) (*types.Transaction, error) {
	return _AlgorithmReview.Contract.SubmitAlgorithm(&_AlgorithmReview.TransactOpts, executionId, scientist, cid, dataset)
}

// Vote is a paid mutator transaction binding the contract method 0x27d3410c.
//
// Solidity: function vote(string cid, bool approve) returns()
func (_AlgorithmReview *AlgorithmReviewTransactor) Vote(opts *bind.TransactOpts, cid string, approve bool) (*types.Transaction, error) {
	return _AlgorithmReview.contract.Transact(opts, "vote", cid, approve)
}

// Vote is a paid mutator transaction binding the contract method 0x27d3410c.
//
// Solidity: function vote(string cid, bool approve) returns()
func (_AlgorithmReview *AlgorithmReviewSession) Vote(cid string, approve bool) (*types.Transaction, error) {
	return _AlgorithmReview.Contract.Vote(&_AlgorithmReview.TransactOpts, cid, approve)
}

// Vote is a paid mutator transaction binding the contract method 0x27d3410c.
//
// Solidity: function vote(string cid, bool approve) returns()
func (_AlgorithmReview *AlgorithmReviewTransactorSession) Vote(cid string, approve bool) (*types.Transaction, error) {
	return _AlgorithmReview.Contract.Vote(&_AlgorithmReview.TransactOpts, cid, approve)
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
	ExecutionId *big.Int
	Cid         string
	Approved    bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAlgorithmResolved is a free log retrieval operation binding the contract event 0x4581a0d468af4916140747e3232ba054684a03025d3aa7698679cb0d653413f8.
//
// Solidity: event AlgorithmResolved(uint256 indexed executionId, string cid, bool approved)
func (_AlgorithmReview *AlgorithmReviewFilterer) FilterAlgorithmResolved(opts *bind.FilterOpts, executionId []*big.Int) (*AlgorithmReviewAlgorithmResolvedIterator, error) {

	var executionIdRule []interface{}
	for _, executionIdItem := range executionId {
		executionIdRule = append(executionIdRule, executionIdItem)
	}

	logs, sub, err := _AlgorithmReview.contract.FilterLogs(opts, "AlgorithmResolved", executionIdRule)
	if err != nil {
		return nil, err
	}
	return &AlgorithmReviewAlgorithmResolvedIterator{contract: _AlgorithmReview.contract, event: "AlgorithmResolved", logs: logs, sub: sub}, nil
}

// WatchAlgorithmResolved is a free log subscription operation binding the contract event 0x4581a0d468af4916140747e3232ba054684a03025d3aa7698679cb0d653413f8.
//
// Solidity: event AlgorithmResolved(uint256 indexed executionId, string cid, bool approved)
func (_AlgorithmReview *AlgorithmReviewFilterer) WatchAlgorithmResolved(opts *bind.WatchOpts, sink chan<- *AlgorithmReviewAlgorithmResolved, executionId []*big.Int) (event.Subscription, error) {

	var executionIdRule []interface{}
	for _, executionIdItem := range executionId {
		executionIdRule = append(executionIdRule, executionIdItem)
	}

	logs, sub, err := _AlgorithmReview.contract.WatchLogs(opts, "AlgorithmResolved", executionIdRule)
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

// ParseAlgorithmResolved is a log parse operation binding the contract event 0x4581a0d468af4916140747e3232ba054684a03025d3aa7698679cb0d653413f8.
//
// Solidity: event AlgorithmResolved(uint256 indexed executionId, string cid, bool approved)
func (_AlgorithmReview *AlgorithmReviewFilterer) ParseAlgorithmResolved(log types.Log) (*AlgorithmReviewAlgorithmResolved, error) {
	event := new(AlgorithmReviewAlgorithmResolved)
	if err := _AlgorithmReview.contract.UnpackLog(event, "AlgorithmResolved", log); err != nil {
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

// AlgorithmReviewExecutionSubmittedIterator is returned from FilterExecutionSubmitted and is used to iterate over the raw logs and unpacked data for ExecutionSubmitted events raised by the AlgorithmReview contract.
type AlgorithmReviewExecutionSubmittedIterator struct {
	Event *AlgorithmReviewExecutionSubmitted // Event containing the contract specifics and raw log

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
func (it *AlgorithmReviewExecutionSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AlgorithmReviewExecutionSubmitted)
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
		it.Event = new(AlgorithmReviewExecutionSubmitted)
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
func (it *AlgorithmReviewExecutionSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AlgorithmReviewExecutionSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AlgorithmReviewExecutionSubmitted represents a ExecutionSubmitted event raised by the AlgorithmReview contract.
type AlgorithmReviewExecutionSubmitted struct {
	ExecutionId *big.Int
	Cid         string
	StartTime   *big.Int
	EndTime     *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterExecutionSubmitted is a free log retrieval operation binding the contract event 0x63ef969e1ddfb70ebfc0b7094e9d170057766c20457c8e1a9a8abc310f7871a1.
//
// Solidity: event ExecutionSubmitted(uint256 indexed executionId, string cid, uint256 startTime, uint256 endTime)
func (_AlgorithmReview *AlgorithmReviewFilterer) FilterExecutionSubmitted(opts *bind.FilterOpts, executionId []*big.Int) (*AlgorithmReviewExecutionSubmittedIterator, error) {

	var executionIdRule []interface{}
	for _, executionIdItem := range executionId {
		executionIdRule = append(executionIdRule, executionIdItem)
	}

	logs, sub, err := _AlgorithmReview.contract.FilterLogs(opts, "ExecutionSubmitted", executionIdRule)
	if err != nil {
		return nil, err
	}
	return &AlgorithmReviewExecutionSubmittedIterator{contract: _AlgorithmReview.contract, event: "ExecutionSubmitted", logs: logs, sub: sub}, nil
}

// WatchExecutionSubmitted is a free log subscription operation binding the contract event 0x63ef969e1ddfb70ebfc0b7094e9d170057766c20457c8e1a9a8abc310f7871a1.
//
// Solidity: event ExecutionSubmitted(uint256 indexed executionId, string cid, uint256 startTime, uint256 endTime)
func (_AlgorithmReview *AlgorithmReviewFilterer) WatchExecutionSubmitted(opts *bind.WatchOpts, sink chan<- *AlgorithmReviewExecutionSubmitted, executionId []*big.Int) (event.Subscription, error) {

	var executionIdRule []interface{}
	for _, executionIdItem := range executionId {
		executionIdRule = append(executionIdRule, executionIdItem)
	}

	logs, sub, err := _AlgorithmReview.contract.WatchLogs(opts, "ExecutionSubmitted", executionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AlgorithmReviewExecutionSubmitted)
				if err := _AlgorithmReview.contract.UnpackLog(event, "ExecutionSubmitted", log); err != nil {
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

// ParseExecutionSubmitted is a log parse operation binding the contract event 0x63ef969e1ddfb70ebfc0b7094e9d170057766c20457c8e1a9a8abc310f7871a1.
//
// Solidity: event ExecutionSubmitted(uint256 indexed executionId, string cid, uint256 startTime, uint256 endTime)
func (_AlgorithmReview *AlgorithmReviewFilterer) ParseExecutionSubmitted(log types.Log) (*AlgorithmReviewExecutionSubmitted, error) {
	event := new(AlgorithmReviewExecutionSubmitted)
	if err := _AlgorithmReview.contract.UnpackLog(event, "ExecutionSubmitted", log); err != nil {
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
	Member   common.Address
	Cid      string
	Approved bool
	VoteTime *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterVoteCasted is a free log retrieval operation binding the contract event 0x5128f8fec9b40ca4b18527897904755cac057c02f346595fc6efd08c3bbf1cd7.
//
// Solidity: event VoteCasted(address indexed member, string cid, bool approved, uint256 voteTime)
func (_AlgorithmReview *AlgorithmReviewFilterer) FilterVoteCasted(opts *bind.FilterOpts, member []common.Address) (*AlgorithmReviewVoteCastedIterator, error) {

	var memberRule []interface{}
	for _, memberItem := range member {
		memberRule = append(memberRule, memberItem)
	}

	logs, sub, err := _AlgorithmReview.contract.FilterLogs(opts, "VoteCasted", memberRule)
	if err != nil {
		return nil, err
	}
	return &AlgorithmReviewVoteCastedIterator{contract: _AlgorithmReview.contract, event: "VoteCasted", logs: logs, sub: sub}, nil
}

// WatchVoteCasted is a free log subscription operation binding the contract event 0x5128f8fec9b40ca4b18527897904755cac057c02f346595fc6efd08c3bbf1cd7.
//
// Solidity: event VoteCasted(address indexed member, string cid, bool approved, uint256 voteTime)
func (_AlgorithmReview *AlgorithmReviewFilterer) WatchVoteCasted(opts *bind.WatchOpts, sink chan<- *AlgorithmReviewVoteCasted, member []common.Address) (event.Subscription, error) {

	var memberRule []interface{}
	for _, memberItem := range member {
		memberRule = append(memberRule, memberItem)
	}

	logs, sub, err := _AlgorithmReview.contract.WatchLogs(opts, "VoteCasted", memberRule)
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

// ParseVoteCasted is a log parse operation binding the contract event 0x5128f8fec9b40ca4b18527897904755cac057c02f346595fc6efd08c3bbf1cd7.
//
// Solidity: event VoteCasted(address indexed member, string cid, bool approved, uint256 voteTime)
func (_AlgorithmReview *AlgorithmReviewFilterer) ParseVoteCasted(log types.Log) (*AlgorithmReviewVoteCasted, error) {
	event := new(AlgorithmReviewVoteCasted)
	if err := _AlgorithmReview.contract.UnpackLog(event, "VoteCasted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
