// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

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

// BeggingContractDonor is an auto generated low-level Go binding around an user-defined struct.
type BeggingContractDonor struct {
	DonorAddress common.Address
	Amount       *big.Int
}

// BeggingContractMetaData contains all meta data concerning the BeggingContract contract.
var BeggingContractMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"durationDays\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"DEADLINE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"beg\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"donate\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"donations\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDonation\",\"inputs\":[{\"name\":\"donor\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTopDonors\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[3]\",\"internalType\":\"structBeggingContract.Donor[3]\",\"components\":[{\"name\":\"donorAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"topDonors\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"donorAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Donation\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdraw\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]}]",
	Bin: "0x60a060405234801561000f575f5ffd5b50604051611a17380380611a17833981810160405281019061003191906102d3565b815f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036100a2575f6040517f1e4fbdf70000000000000000000000000000000000000000000000000000000081526004016100999190610320565b60405180910390fd5b6100b18161014f60201b60201c565b5060016100d06100c561021060201b60201c565b61023960201b60201c565b5f0181905550600181101580156100e85750605a8111155b610127576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161011e90610393565b60405180910390fd5b620151808161013691906103de565b42610141919061041f565b608081815250505050610452565b5f5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f7f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f005f1b905090565b5f819050919050565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61026f82610246565b9050919050565b61027f81610265565b8114610289575f5ffd5b50565b5f8151905061029a81610276565b92915050565b5f819050919050565b6102b2816102a0565b81146102bc575f5ffd5b50565b5f815190506102cd816102a9565b92915050565b5f5f604083850312156102e9576102e8610242565b5b5f6102f68582860161028c565b9250506020610307858286016102bf565b9150509250929050565b61031a81610265565b82525050565b5f6020820190506103335f830184610311565b92915050565b5f82825260208201905092915050565b7f4475726174696f6e3a20312d39302064617973000000000000000000000000005f82015250565b5f61037d601383610339565b915061038882610349565b602082019050919050565b5f6020820190508181035f8301526103aa81610371565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f6103e8826102a0565b91506103f3836102a0565b9250828202610401816102a0565b91508282048414831517610418576104176103b1565b5b5092915050565b5f610429826102a0565b9150610434836102a0565b925082820190508082111561044c5761044b6103b1565b5b92915050565b60805161159f6104785f395f818161053e0152818161069601526106bc015261159f5ff3fe60806040526004361061009b575f3560e01c80638da5cb5b116100635780638da5cb5b1461016e578063a082c86e14610198578063cc6cb19a146101c2578063d6387ed8146101fe578063ed88c68e14610228578063f2fde38b146102325761009b565b80633ccfd60b1461009f578063410a1d32146100b55780635cbaddab146100f1578063715018a61461012e5780637d780d8214610144575b5f5ffd5b3480156100aa575f5ffd5b506100b361025a565b005b3480156100c0575f5ffd5b506100db60048036038101906100d69190610fb0565b610439565b6040516100e89190610ff3565b60405180910390f35b3480156100fc575f5ffd5b5061011760048036038101906101129190611036565b61047f565b604051610125929190611070565b60405180910390f35b348015610139575f5ffd5b506101426104c5565b005b34801561014f575f5ffd5b506101586104d8565b6040516101659190611107565b60405180910390f35b348015610179575f5ffd5b50610182610515565b60405161018f9190611127565b60405180910390f35b3480156101a3575f5ffd5b506101ac61053c565b6040516101b99190610ff3565b60405180910390f35b3480156101cd575f5ffd5b506101e860048036038101906101e39190610fb0565b610560565b6040516101f59190610ff3565b60405180910390f35b348015610209575f5ffd5b50610212610575565b60405161021f9190611221565b60405180910390f35b610230610624565b005b34801561023d575f5ffd5b5061025860048036038101906102539190610fb0565b610a2e565b005b610262610ab2565b61026a610ad4565b5f4790506102eb6040518060400160405280601b81526020017f5769746864726177696e672066756e647320746f206f776e65723a00000000008152506102af610515565b6040518060400160405280600781526020017f416d6f756e743a0000000000000000000000000000000000000000000000000081525084610b5b565b5f6102f4610515565b73ffffffffffffffffffffffffffffffffffffffff168260405161031790611267565b5f6040518083038185875af1925050503d805f8114610351576040519150601f19603f3d011682016040523d82523d5f602084013e610356565b606091505b505090508061039a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610391906112c5565b60405180910390fd5b6103d86040518060400160405280600681526020017f5858585858580000000000000000000000000000000000000000000000000000815250610bfd565b6103e0610515565b73ffffffffffffffffffffffffffffffffffffffff167f884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364836040516104259190610ff3565b60405180910390a25050610437610c96565b565b5f60015f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20549050919050565b6002816003811061048e575f80fd5b600202015f91509050805f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010154905082565b6104cd610ad4565b6104d65f610cb0565b565b60606040518060400160405280601a81526020017f506c656173652067697665206d6520736f6d65206d6f6e657921000000000000815250905090565b5f5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b7f000000000000000000000000000000000000000000000000000000000000000081565b6001602052805f5260405f205f915090505481565b61057d610eed565b6002600380602002604051908101604052809291905f905b8282101561061b578382600202016040518060400160405290815f82015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160018201548152505081526020019060010190610595565b50505050905090565b6106ba6040518060400160405280601281526020017f526563656976656420646f6e6174696f6e3a0000000000000000000000000000815250426040518060400160405280600981526020017f444541444c494e453a00000000000000000000000000000000000000000000008152507f0000000000000000000000000000000000000000000000000000000000000000610d71565b7f000000000000000000000000000000000000000000000000000000000000000042111561071d576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107149061132d565b60405180910390fd5b5f341161075f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161075690611395565b60405180910390fd5b3460015f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546107ab91906113e0565b925050819055503373ffffffffffffffffffffffffffffffffffffffff167f5d8bc849764969eb1bcc6d0a2f55999d0167c1ccec240a4f39cf664ca9c4148e346040516107f89190610ff3565b60405180910390a25f5f90505b6003811015610a2b576002816003811061082257610821611413565b5b600202016001015460015f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20541115610a1e575f6001600361087d9190611440565b90505b818111156109475760026001826108979190611440565b600381106108a8576108a7611413565b5b60020201600282600381106108c0576108bf611413565b5b600202015f82015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16815f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060018201548160010155905050808061093f90611473565b915050610880565b5060405180604001604052803373ffffffffffffffffffffffffffffffffffffffff16815260200160015f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054815250600282600381106109c4576109c3611413565b5b600202015f820151815f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010155905050610a2b565b8080600101915050610805565b50565b610a36610ad4565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610aa6575f6040517f1e4fbdf7000000000000000000000000000000000000000000000000000000008152600401610a9d9190611127565b60405180910390fd5b610aaf81610cb0565b50565b610aba610e13565b6002610acc610ac7610e54565b610e7d565b5f0181905550565b610adc610e86565b73ffffffffffffffffffffffffffffffffffffffff16610afa610515565b73ffffffffffffffffffffffffffffffffffffffff1614610b5957610b1d610e86565b6040517f118cdaa7000000000000000000000000000000000000000000000000000000008152600401610b509190611127565b60405180910390fd5b565b610bf784848484604051602401610b75949392919061149a565b6040516020818303038152906040527f91d1112e000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050610e8d565b50505050565b610c9381604051602401610c119190611107565b6040516020818303038152906040527f41304fac000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050610e8d565b50565b6001610ca8610ca3610e54565b610e7d565b5f0181905550565b5f5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b610e0d84848484604051602401610d8b94939291906114eb565b6040516020818303038152906040527fc67ea9d1000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050610e8d565b50505050565b610e1b610ea7565b15610e52576040517f3ee5aeb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b5f7f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f005f1b905090565b5f819050919050565b5f33905090565b610ea481610e9c610ec3610ee2565b63ffffffff16565b50565b5f6002610eba610eb5610e54565b610e7d565b5f015414905090565b5f6a636f6e736f6c652e6c6f6790505f5f835160208501845afa505050565b610f1a819050919050565b60405180606001604052806003905b610f04610f24565b815260200190600190039081610efc5790505090565b610f2261153c565b565b60405180604001604052805f73ffffffffffffffffffffffffffffffffffffffff1681526020015f81525090565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610f7f82610f56565b9050919050565b610f8f81610f75565b8114610f99575f5ffd5b50565b5f81359050610faa81610f86565b92915050565b5f60208284031215610fc557610fc4610f52565b5b5f610fd284828501610f9c565b91505092915050565b5f819050919050565b610fed81610fdb565b82525050565b5f6020820190506110065f830184610fe4565b92915050565b61101581610fdb565b811461101f575f5ffd5b50565b5f813590506110308161100c565b92915050565b5f6020828403121561104b5761104a610f52565b5b5f61105884828501611022565b91505092915050565b61106a81610f75565b82525050565b5f6040820190506110835f830185611061565b6110906020830184610fe4565b9392505050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f6110d982611097565b6110e381856110a1565b93506110f38185602086016110b1565b6110fc816110bf565b840191505092915050565b5f6020820190508181035f83015261111f81846110cf565b905092915050565b5f60208201905061113a5f830184611061565b92915050565b5f60039050919050565b5f81905092915050565b5f819050919050565b61116681610f75565b82525050565b61117581610fdb565b82525050565b604082015f82015161118f5f85018261115d565b5060208201516111a2602085018261116c565b50505050565b5f6111b3838361117b565b60408301905092915050565b5f602082019050919050565b6111d481611140565b6111de818461114a565b92506111e982611154565b805f5b8381101561121957815161120087826111a8565b965061120b836111bf565b9250506001810190506111ec565b505050505050565b5f60c0820190506112345f8301846111cb565b92915050565b5f81905092915050565b50565b5f6112525f8361123a565b915061125d82611244565b5f82019050919050565b5f61127182611247565b9150819050919050565b7f4661696c656420746f20776974686472617720457468657200000000000000005f82015250565b5f6112af6018836110a1565b91506112ba8261127b565b602082019050919050565b5f6020820190508181035f8301526112dc816112a3565b9050919050565b7f46756e6472616973696e6720656e6465640000000000000000000000000000005f82015250565b5f6113176011836110a1565b9150611322826112e3565b602082019050919050565b5f6020820190508181035f8301526113448161130b565b9050919050565b7f446f6e6174696f6e206d7573742062652067726561746572207468616e2030005f82015250565b5f61137f601f836110a1565b915061138a8261134b565b602082019050919050565b5f6020820190508181035f8301526113ac81611373565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f6113ea82610fdb565b91506113f583610fdb565b925082820190508082111561140d5761140c6113b3565b5b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f61144a82610fdb565b915061145583610fdb565b925082820390508181111561146d5761146c6113b3565b5b92915050565b5f61147d82610fdb565b91505f820361148f5761148e6113b3565b5b600182039050919050565b5f6080820190508181035f8301526114b281876110cf565b90506114c16020830186611061565b81810360408301526114d381856110cf565b90506114e26060830184610fe4565b95945050505050565b5f6080820190508181035f83015261150381876110cf565b90506115126020830186610fe4565b818103604083015261152481856110cf565b90506115336060830184610fe4565b95945050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52605160045260245ffdfea2646970667358221220e0e37da7ff35443d5f663005c8e35feacbe30fb29203e14dbe68b11e7a554c8364736f6c63430008230033",
}

// BeggingContractABI is the input ABI used to generate the binding from.
// Deprecated: Use BeggingContractMetaData.ABI instead.
var BeggingContractABI = BeggingContractMetaData.ABI

// BeggingContractBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BeggingContractMetaData.Bin instead.
var BeggingContractBin = BeggingContractMetaData.Bin

// DeployBeggingContract deploys a new Ethereum contract, binding an instance of BeggingContract to it.
func DeployBeggingContract(auth *bind.TransactOpts, backend bind.ContractBackend, initialOwner common.Address, durationDays *big.Int) (common.Address, *types.Transaction, *BeggingContract, error) {
	parsed, err := BeggingContractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BeggingContractBin), backend, initialOwner, durationDays)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BeggingContract{BeggingContractCaller: BeggingContractCaller{contract: contract}, BeggingContractTransactor: BeggingContractTransactor{contract: contract}, BeggingContractFilterer: BeggingContractFilterer{contract: contract}}, nil
}

// BeggingContract is an auto generated Go binding around an Ethereum contract.
type BeggingContract struct {
	BeggingContractCaller     // Read-only binding to the contract
	BeggingContractTransactor // Write-only binding to the contract
	BeggingContractFilterer   // Log filterer for contract events
}

// BeggingContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type BeggingContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BeggingContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BeggingContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BeggingContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BeggingContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BeggingContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BeggingContractSession struct {
	Contract     *BeggingContract  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BeggingContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BeggingContractCallerSession struct {
	Contract *BeggingContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// BeggingContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BeggingContractTransactorSession struct {
	Contract     *BeggingContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// BeggingContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type BeggingContractRaw struct {
	Contract *BeggingContract // Generic contract binding to access the raw methods on
}

// BeggingContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BeggingContractCallerRaw struct {
	Contract *BeggingContractCaller // Generic read-only contract binding to access the raw methods on
}

// BeggingContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BeggingContractTransactorRaw struct {
	Contract *BeggingContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBeggingContract creates a new instance of BeggingContract, bound to a specific deployed contract.
func NewBeggingContract(address common.Address, backend bind.ContractBackend) (*BeggingContract, error) {
	contract, err := bindBeggingContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BeggingContract{BeggingContractCaller: BeggingContractCaller{contract: contract}, BeggingContractTransactor: BeggingContractTransactor{contract: contract}, BeggingContractFilterer: BeggingContractFilterer{contract: contract}}, nil
}

// NewBeggingContractCaller creates a new read-only instance of BeggingContract, bound to a specific deployed contract.
func NewBeggingContractCaller(address common.Address, caller bind.ContractCaller) (*BeggingContractCaller, error) {
	contract, err := bindBeggingContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BeggingContractCaller{contract: contract}, nil
}

// NewBeggingContractTransactor creates a new write-only instance of BeggingContract, bound to a specific deployed contract.
func NewBeggingContractTransactor(address common.Address, transactor bind.ContractTransactor) (*BeggingContractTransactor, error) {
	contract, err := bindBeggingContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BeggingContractTransactor{contract: contract}, nil
}

// NewBeggingContractFilterer creates a new log filterer instance of BeggingContract, bound to a specific deployed contract.
func NewBeggingContractFilterer(address common.Address, filterer bind.ContractFilterer) (*BeggingContractFilterer, error) {
	contract, err := bindBeggingContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BeggingContractFilterer{contract: contract}, nil
}

// bindBeggingContract binds a generic wrapper to an already deployed contract.
func bindBeggingContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BeggingContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BeggingContract *BeggingContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BeggingContract.Contract.BeggingContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BeggingContract *BeggingContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BeggingContract.Contract.BeggingContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BeggingContract *BeggingContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BeggingContract.Contract.BeggingContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BeggingContract *BeggingContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BeggingContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BeggingContract *BeggingContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BeggingContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BeggingContract *BeggingContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BeggingContract.Contract.contract.Transact(opts, method, params...)
}

// DEADLINE is a free data retrieval call binding the contract method 0xa082c86e.
//
// Solidity: function DEADLINE() view returns(uint256)
func (_BeggingContract *BeggingContractCaller) DEADLINE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BeggingContract.contract.Call(opts, &out, "DEADLINE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DEADLINE is a free data retrieval call binding the contract method 0xa082c86e.
//
// Solidity: function DEADLINE() view returns(uint256)
func (_BeggingContract *BeggingContractSession) DEADLINE() (*big.Int, error) {
	return _BeggingContract.Contract.DEADLINE(&_BeggingContract.CallOpts)
}

// DEADLINE is a free data retrieval call binding the contract method 0xa082c86e.
//
// Solidity: function DEADLINE() view returns(uint256)
func (_BeggingContract *BeggingContractCallerSession) DEADLINE() (*big.Int, error) {
	return _BeggingContract.Contract.DEADLINE(&_BeggingContract.CallOpts)
}

// Beg is a free data retrieval call binding the contract method 0x7d780d82.
//
// Solidity: function beg() pure returns(string)
func (_BeggingContract *BeggingContractCaller) Beg(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BeggingContract.contract.Call(opts, &out, "beg")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Beg is a free data retrieval call binding the contract method 0x7d780d82.
//
// Solidity: function beg() pure returns(string)
func (_BeggingContract *BeggingContractSession) Beg() (string, error) {
	return _BeggingContract.Contract.Beg(&_BeggingContract.CallOpts)
}

// Beg is a free data retrieval call binding the contract method 0x7d780d82.
//
// Solidity: function beg() pure returns(string)
func (_BeggingContract *BeggingContractCallerSession) Beg() (string, error) {
	return _BeggingContract.Contract.Beg(&_BeggingContract.CallOpts)
}

// Donations is a free data retrieval call binding the contract method 0xcc6cb19a.
//
// Solidity: function donations(address ) view returns(uint256)
func (_BeggingContract *BeggingContractCaller) Donations(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BeggingContract.contract.Call(opts, &out, "donations", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Donations is a free data retrieval call binding the contract method 0xcc6cb19a.
//
// Solidity: function donations(address ) view returns(uint256)
func (_BeggingContract *BeggingContractSession) Donations(arg0 common.Address) (*big.Int, error) {
	return _BeggingContract.Contract.Donations(&_BeggingContract.CallOpts, arg0)
}

// Donations is a free data retrieval call binding the contract method 0xcc6cb19a.
//
// Solidity: function donations(address ) view returns(uint256)
func (_BeggingContract *BeggingContractCallerSession) Donations(arg0 common.Address) (*big.Int, error) {
	return _BeggingContract.Contract.Donations(&_BeggingContract.CallOpts, arg0)
}

// GetDonation is a free data retrieval call binding the contract method 0x410a1d32.
//
// Solidity: function getDonation(address donor) view returns(uint256)
func (_BeggingContract *BeggingContractCaller) GetDonation(opts *bind.CallOpts, donor common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BeggingContract.contract.Call(opts, &out, "getDonation", donor)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDonation is a free data retrieval call binding the contract method 0x410a1d32.
//
// Solidity: function getDonation(address donor) view returns(uint256)
func (_BeggingContract *BeggingContractSession) GetDonation(donor common.Address) (*big.Int, error) {
	return _BeggingContract.Contract.GetDonation(&_BeggingContract.CallOpts, donor)
}

// GetDonation is a free data retrieval call binding the contract method 0x410a1d32.
//
// Solidity: function getDonation(address donor) view returns(uint256)
func (_BeggingContract *BeggingContractCallerSession) GetDonation(donor common.Address) (*big.Int, error) {
	return _BeggingContract.Contract.GetDonation(&_BeggingContract.CallOpts, donor)
}

// GetTopDonors is a free data retrieval call binding the contract method 0xd6387ed8.
//
// Solidity: function getTopDonors() view returns((address,uint256)[3])
func (_BeggingContract *BeggingContractCaller) GetTopDonors(opts *bind.CallOpts) ([3]BeggingContractDonor, error) {
	var out []interface{}
	err := _BeggingContract.contract.Call(opts, &out, "getTopDonors")

	if err != nil {
		return *new([3]BeggingContractDonor), err
	}

	out0 := *abi.ConvertType(out[0], new([3]BeggingContractDonor)).(*[3]BeggingContractDonor)

	return out0, err

}

// GetTopDonors is a free data retrieval call binding the contract method 0xd6387ed8.
//
// Solidity: function getTopDonors() view returns((address,uint256)[3])
func (_BeggingContract *BeggingContractSession) GetTopDonors() ([3]BeggingContractDonor, error) {
	return _BeggingContract.Contract.GetTopDonors(&_BeggingContract.CallOpts)
}

// GetTopDonors is a free data retrieval call binding the contract method 0xd6387ed8.
//
// Solidity: function getTopDonors() view returns((address,uint256)[3])
func (_BeggingContract *BeggingContractCallerSession) GetTopDonors() ([3]BeggingContractDonor, error) {
	return _BeggingContract.Contract.GetTopDonors(&_BeggingContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BeggingContract *BeggingContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BeggingContract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BeggingContract *BeggingContractSession) Owner() (common.Address, error) {
	return _BeggingContract.Contract.Owner(&_BeggingContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BeggingContract *BeggingContractCallerSession) Owner() (common.Address, error) {
	return _BeggingContract.Contract.Owner(&_BeggingContract.CallOpts)
}

// TopDonors is a free data retrieval call binding the contract method 0x5cbaddab.
//
// Solidity: function topDonors(uint256 ) view returns(address donorAddress, uint256 amount)
func (_BeggingContract *BeggingContractCaller) TopDonors(opts *bind.CallOpts, arg0 *big.Int) (struct {
	DonorAddress common.Address
	Amount       *big.Int
}, error) {
	var out []interface{}
	err := _BeggingContract.contract.Call(opts, &out, "topDonors", arg0)

	outstruct := new(struct {
		DonorAddress common.Address
		Amount       *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.DonorAddress = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Amount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// TopDonors is a free data retrieval call binding the contract method 0x5cbaddab.
//
// Solidity: function topDonors(uint256 ) view returns(address donorAddress, uint256 amount)
func (_BeggingContract *BeggingContractSession) TopDonors(arg0 *big.Int) (struct {
	DonorAddress common.Address
	Amount       *big.Int
}, error) {
	return _BeggingContract.Contract.TopDonors(&_BeggingContract.CallOpts, arg0)
}

// TopDonors is a free data retrieval call binding the contract method 0x5cbaddab.
//
// Solidity: function topDonors(uint256 ) view returns(address donorAddress, uint256 amount)
func (_BeggingContract *BeggingContractCallerSession) TopDonors(arg0 *big.Int) (struct {
	DonorAddress common.Address
	Amount       *big.Int
}, error) {
	return _BeggingContract.Contract.TopDonors(&_BeggingContract.CallOpts, arg0)
}

// Donate is a paid mutator transaction binding the contract method 0xed88c68e.
//
// Solidity: function donate() payable returns()
func (_BeggingContract *BeggingContractTransactor) Donate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BeggingContract.contract.Transact(opts, "donate")
}

// Donate is a paid mutator transaction binding the contract method 0xed88c68e.
//
// Solidity: function donate() payable returns()
func (_BeggingContract *BeggingContractSession) Donate() (*types.Transaction, error) {
	return _BeggingContract.Contract.Donate(&_BeggingContract.TransactOpts)
}

// Donate is a paid mutator transaction binding the contract method 0xed88c68e.
//
// Solidity: function donate() payable returns()
func (_BeggingContract *BeggingContractTransactorSession) Donate() (*types.Transaction, error) {
	return _BeggingContract.Contract.Donate(&_BeggingContract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BeggingContract *BeggingContractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BeggingContract.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BeggingContract *BeggingContractSession) RenounceOwnership() (*types.Transaction, error) {
	return _BeggingContract.Contract.RenounceOwnership(&_BeggingContract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BeggingContract *BeggingContractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _BeggingContract.Contract.RenounceOwnership(&_BeggingContract.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BeggingContract *BeggingContractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BeggingContract.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BeggingContract *BeggingContractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BeggingContract.Contract.TransferOwnership(&_BeggingContract.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BeggingContract *BeggingContractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BeggingContract.Contract.TransferOwnership(&_BeggingContract.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_BeggingContract *BeggingContractTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BeggingContract.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_BeggingContract *BeggingContractSession) Withdraw() (*types.Transaction, error) {
	return _BeggingContract.Contract.Withdraw(&_BeggingContract.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_BeggingContract *BeggingContractTransactorSession) Withdraw() (*types.Transaction, error) {
	return _BeggingContract.Contract.Withdraw(&_BeggingContract.TransactOpts)
}

// BeggingContractDonationIterator is returned from FilterDonation and is used to iterate over the raw logs and unpacked data for Donation events raised by the BeggingContract contract.
type BeggingContractDonationIterator struct {
	Event *BeggingContractDonation // Event containing the contract specifics and raw log

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
func (it *BeggingContractDonationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BeggingContractDonation)
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
		it.Event = new(BeggingContractDonation)
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
func (it *BeggingContractDonationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BeggingContractDonationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BeggingContractDonation represents a Donation event raised by the BeggingContract contract.
type BeggingContractDonation struct {
	From   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDonation is a free log retrieval operation binding the contract event 0x5d8bc849764969eb1bcc6d0a2f55999d0167c1ccec240a4f39cf664ca9c4148e.
//
// Solidity: event Donation(address indexed from, uint256 amount)
func (_BeggingContract *BeggingContractFilterer) FilterDonation(opts *bind.FilterOpts, from []common.Address) (*BeggingContractDonationIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _BeggingContract.contract.FilterLogs(opts, "Donation", fromRule)
	if err != nil {
		return nil, err
	}
	return &BeggingContractDonationIterator{contract: _BeggingContract.contract, event: "Donation", logs: logs, sub: sub}, nil
}

// WatchDonation is a free log subscription operation binding the contract event 0x5d8bc849764969eb1bcc6d0a2f55999d0167c1ccec240a4f39cf664ca9c4148e.
//
// Solidity: event Donation(address indexed from, uint256 amount)
func (_BeggingContract *BeggingContractFilterer) WatchDonation(opts *bind.WatchOpts, sink chan<- *BeggingContractDonation, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _BeggingContract.contract.WatchLogs(opts, "Donation", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BeggingContractDonation)
				if err := _BeggingContract.contract.UnpackLog(event, "Donation", log); err != nil {
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

// ParseDonation is a log parse operation binding the contract event 0x5d8bc849764969eb1bcc6d0a2f55999d0167c1ccec240a4f39cf664ca9c4148e.
//
// Solidity: event Donation(address indexed from, uint256 amount)
func (_BeggingContract *BeggingContractFilterer) ParseDonation(log types.Log) (*BeggingContractDonation, error) {
	event := new(BeggingContractDonation)
	if err := _BeggingContract.contract.UnpackLog(event, "Donation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BeggingContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BeggingContract contract.
type BeggingContractOwnershipTransferredIterator struct {
	Event *BeggingContractOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BeggingContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BeggingContractOwnershipTransferred)
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
		it.Event = new(BeggingContractOwnershipTransferred)
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
func (it *BeggingContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BeggingContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BeggingContractOwnershipTransferred represents a OwnershipTransferred event raised by the BeggingContract contract.
type BeggingContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BeggingContract *BeggingContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BeggingContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BeggingContract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BeggingContractOwnershipTransferredIterator{contract: _BeggingContract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BeggingContract *BeggingContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BeggingContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BeggingContract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BeggingContractOwnershipTransferred)
				if err := _BeggingContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_BeggingContract *BeggingContractFilterer) ParseOwnershipTransferred(log types.Log) (*BeggingContractOwnershipTransferred, error) {
	event := new(BeggingContractOwnershipTransferred)
	if err := _BeggingContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BeggingContractWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the BeggingContract contract.
type BeggingContractWithdrawIterator struct {
	Event *BeggingContractWithdraw // Event containing the contract specifics and raw log

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
func (it *BeggingContractWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BeggingContractWithdraw)
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
		it.Event = new(BeggingContractWithdraw)
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
func (it *BeggingContractWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BeggingContractWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BeggingContractWithdraw represents a Withdraw event raised by the BeggingContract contract.
type BeggingContractWithdraw struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed to, uint256 amount)
func (_BeggingContract *BeggingContractFilterer) FilterWithdraw(opts *bind.FilterOpts, to []common.Address) (*BeggingContractWithdrawIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BeggingContract.contract.FilterLogs(opts, "Withdraw", toRule)
	if err != nil {
		return nil, err
	}
	return &BeggingContractWithdrawIterator{contract: _BeggingContract.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed to, uint256 amount)
func (_BeggingContract *BeggingContractFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *BeggingContractWithdraw, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BeggingContract.contract.WatchLogs(opts, "Withdraw", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BeggingContractWithdraw)
				if err := _BeggingContract.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed to, uint256 amount)
func (_BeggingContract *BeggingContractFilterer) ParseWithdraw(log types.Log) (*BeggingContractWithdraw, error) {
	event := new(BeggingContractWithdraw)
	if err := _BeggingContract.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
