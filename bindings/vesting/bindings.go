// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package vesting

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

// VestingABI is the input ABI used to generate the binding from.
const VestingABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_levTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_dateTimeContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"TokensReleased\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"currentCycle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"endTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isPrepared\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_cycle\",\"type\":\"uint256\"}],\"name\":\"isReleased\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amountToVest\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"prepare\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"receiver\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"release\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"releaseAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"releases\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"released\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// VestingBin is the compiled bytecode used for deploying new contracts.
var VestingBin = "0x608060405234801561001057600080fd5b5060405161080a38038061080a8339818101604052606081101561003357600080fd5b5080516020820151604090920151600680546001600160a01b03199081166001600160a01b03948516179091556007805482169484169490941790935560058054909316911617905561077f8061008b6000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c80638da5cb5b116100715780638da5cb5b14610137578063939c28e61461015b578063b6a9f40f14610163578063bab2f55214610199578063c062dc5f146101a1578063f7260d3e146101a9576100a9565b80633197cbb6146100ae57806373962b26146100c857806378e97925146100f95780637c0c3e021461010157806386d1a69f1461012f575b600080fd5b6100b66101b1565b60408051918252519081900360200190f35b6100e5600480360360208110156100de57600080fd5b50356101b7565b604080519115158252519081900360200190f35b6100b66101dd565b61012d6004803603604081101561011757600080fd5b50803590602001356001600160a01b03166101e3565b005b61012d610498565b61013f6106c9565b604080516001600160a01b039092168252519081900360200190f35b6100e56106d8565b6101806004803603602081101561017957600080fd5b5035610715565b6040805192835260208301919091528051918290030190f35b6100b661072e565b6100b6610734565b61013f61073a565b60015481565b6000818152600860205260408120600190810154829114156101d7575060015b92915050565b60005481565b6005546001600160a01b031633146101fa57600080fd5b6102026106d8565b1561020c57600080fd5b600654604080516323b872dd60e01b81523360048201523060248201526044810185905290516001600160a01b03909216916323b872dd916064808201926020929091908290030181600087803b15801561026657600080fd5b505af115801561027a573d6000803e3d6000fd5b505050506040513d602081101561029057600080fd5b505161029b57600080fd5b60075460408051632ceee33560e21b815290516000926001600160a01b03169163b3bb8cd4916004808301926020929190829003018186803b1580156102e057600080fd5b505afa1580156102f4573d6000803e3d6000fd5b505050506040513d602081101561030a57600080fd5b505160075460408051634355644d60e01b815260048101849052600c602482015290519293506000926001600160a01b0390921691634355644d91604480820192602092909190829003018186803b15801561036557600080fd5b505afa158015610379573d6000803e3d6000fd5b505050506040513d602081101561038f57600080fd5b5051600c60005260086020527faba08fc86a199da17c4b45fa0fb586714474d27e5e1f0220658b6275561b2e18819055905060015b600b81116104625760075460408051634355644d60e01b8152600481018690526024810184905290516001600160a01b0390921691634355644d91604480820192602092909190829003018186803b15801561041f57600080fd5b505afa158015610433573d6000803e3d6000fd5b505050506040513d602081101561044957600080fd5b50516000828152600860205260409020556001016103c4565b50600c9093046003556000556001918255600480546001600160a01b0319166001600160a01b0392909216919091179055600255565b6104a06106d8565b15156001146104ae57600080fd5b60025460009081526008602052604090206001015415610508576040805162461bcd60e51b815260206004820152601060248201526f185b1c9958591e481c995b19585cd95960821b604482015290519081900360640190fd5b60025460009081526008602090815260408083206001908101556007548151632ceee33560e21b815291516001600160a01b039091169263b3bb8cd49260048082019391829003018186803b15801561056057600080fd5b505afa158015610574573d6000803e3d6000fd5b505050506040513d602081101561058a57600080fd5b50516002546000908152600860205260409020549091508110156105f5576040805162461bcd60e51b815260206004820181905260248201527f72656c656173652074696d657374616d70206e6f742079657420706173736564604482015290519081900360640190fd5b600654600480546003546040805163a9059cbb60e01b81526001600160a01b039384169481019490945260248401919091525192169163a9059cbb916044808201926020929091908290030181600087803b15801561065357600080fd5b505af1158015610667573d6000803e3d6000fd5b505050506040513d602081101561067d57600080fd5b505161068857600080fd5b600c600254101561069d576002805460010190555b6040517f3cf447e1c5a595a578bfe340289a0dad7a7f22771c57c9443ad7eedea1964d7690600090a150565b6005546001600160a01b031681565b60045460009081906001600160a01b0316158015906106f957506000600354115b801561070757506000600254115b15610710575060015b905090565b6008602052600090815260409020805460019091015482565b60025481565b60035481565b6004546001600160a01b03168156fea2646970667358221220587624860927978813ceaa246b835f37199d513605b7a7b695be188116d1f86e64736f6c63430007030033"

// DeployVesting deploys a new Ethereum contract, binding an instance of Vesting to it.
func DeployVesting(auth *bind.TransactOpts, backend bind.ContractBackend, _levTokenAddress common.Address, _dateTimeContract common.Address, _owner common.Address) (common.Address, *types.Transaction, *Vesting, error) {
	parsed, err := abi.JSON(strings.NewReader(VestingABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(VestingBin), backend, _levTokenAddress, _dateTimeContract, _owner)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Vesting{VestingCaller: VestingCaller{contract: contract}, VestingTransactor: VestingTransactor{contract: contract}, VestingFilterer: VestingFilterer{contract: contract}}, nil
}

// Vesting is an auto generated Go binding around an Ethereum contract.
type Vesting struct {
	VestingCaller     // Read-only binding to the contract
	VestingTransactor // Write-only binding to the contract
	VestingFilterer   // Log filterer for contract events
}

// VestingCaller is an auto generated read-only Go binding around an Ethereum contract.
type VestingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VestingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VestingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VestingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VestingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VestingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VestingSession struct {
	Contract     *Vesting          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VestingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VestingCallerSession struct {
	Contract *VestingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// VestingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VestingTransactorSession struct {
	Contract     *VestingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// VestingRaw is an auto generated low-level Go binding around an Ethereum contract.
type VestingRaw struct {
	Contract *Vesting // Generic contract binding to access the raw methods on
}

// VestingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VestingCallerRaw struct {
	Contract *VestingCaller // Generic read-only contract binding to access the raw methods on
}

// VestingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VestingTransactorRaw struct {
	Contract *VestingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVesting creates a new instance of Vesting, bound to a specific deployed contract.
func NewVesting(address common.Address, backend bind.ContractBackend) (*Vesting, error) {
	contract, err := bindVesting(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Vesting{VestingCaller: VestingCaller{contract: contract}, VestingTransactor: VestingTransactor{contract: contract}, VestingFilterer: VestingFilterer{contract: contract}}, nil
}

// NewVestingCaller creates a new read-only instance of Vesting, bound to a specific deployed contract.
func NewVestingCaller(address common.Address, caller bind.ContractCaller) (*VestingCaller, error) {
	contract, err := bindVesting(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VestingCaller{contract: contract}, nil
}

// NewVestingTransactor creates a new write-only instance of Vesting, bound to a specific deployed contract.
func NewVestingTransactor(address common.Address, transactor bind.ContractTransactor) (*VestingTransactor, error) {
	contract, err := bindVesting(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VestingTransactor{contract: contract}, nil
}

// NewVestingFilterer creates a new log filterer instance of Vesting, bound to a specific deployed contract.
func NewVestingFilterer(address common.Address, filterer bind.ContractFilterer) (*VestingFilterer, error) {
	contract, err := bindVesting(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VestingFilterer{contract: contract}, nil
}

// bindVesting binds a generic wrapper to an already deployed contract.
func bindVesting(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VestingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vesting *VestingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vesting.Contract.VestingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vesting *VestingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vesting.Contract.VestingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vesting *VestingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vesting.Contract.VestingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vesting *VestingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vesting.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vesting *VestingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vesting.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vesting *VestingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vesting.Contract.contract.Transact(opts, method, params...)
}

// CurrentCycle is a free data retrieval call binding the contract method 0xbab2f552.
//
// Solidity: function currentCycle() view returns(uint256)
func (_Vesting *VestingCaller) CurrentCycle(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vesting.contract.Call(opts, &out, "currentCycle")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentCycle is a free data retrieval call binding the contract method 0xbab2f552.
//
// Solidity: function currentCycle() view returns(uint256)
func (_Vesting *VestingSession) CurrentCycle() (*big.Int, error) {
	return _Vesting.Contract.CurrentCycle(&_Vesting.CallOpts)
}

// CurrentCycle is a free data retrieval call binding the contract method 0xbab2f552.
//
// Solidity: function currentCycle() view returns(uint256)
func (_Vesting *VestingCallerSession) CurrentCycle() (*big.Int, error) {
	return _Vesting.Contract.CurrentCycle(&_Vesting.CallOpts)
}

// EndTime is a free data retrieval call binding the contract method 0x3197cbb6.
//
// Solidity: function endTime() view returns(uint256)
func (_Vesting *VestingCaller) EndTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vesting.contract.Call(opts, &out, "endTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EndTime is a free data retrieval call binding the contract method 0x3197cbb6.
//
// Solidity: function endTime() view returns(uint256)
func (_Vesting *VestingSession) EndTime() (*big.Int, error) {
	return _Vesting.Contract.EndTime(&_Vesting.CallOpts)
}

// EndTime is a free data retrieval call binding the contract method 0x3197cbb6.
//
// Solidity: function endTime() view returns(uint256)
func (_Vesting *VestingCallerSession) EndTime() (*big.Int, error) {
	return _Vesting.Contract.EndTime(&_Vesting.CallOpts)
}

// IsPrepared is a free data retrieval call binding the contract method 0x939c28e6.
//
// Solidity: function isPrepared() view returns(bool)
func (_Vesting *VestingCaller) IsPrepared(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Vesting.contract.Call(opts, &out, "isPrepared")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPrepared is a free data retrieval call binding the contract method 0x939c28e6.
//
// Solidity: function isPrepared() view returns(bool)
func (_Vesting *VestingSession) IsPrepared() (bool, error) {
	return _Vesting.Contract.IsPrepared(&_Vesting.CallOpts)
}

// IsPrepared is a free data retrieval call binding the contract method 0x939c28e6.
//
// Solidity: function isPrepared() view returns(bool)
func (_Vesting *VestingCallerSession) IsPrepared() (bool, error) {
	return _Vesting.Contract.IsPrepared(&_Vesting.CallOpts)
}

// IsReleased is a free data retrieval call binding the contract method 0x73962b26.
//
// Solidity: function isReleased(uint256 _cycle) view returns(bool)
func (_Vesting *VestingCaller) IsReleased(opts *bind.CallOpts, _cycle *big.Int) (bool, error) {
	var out []interface{}
	err := _Vesting.contract.Call(opts, &out, "isReleased", _cycle)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsReleased is a free data retrieval call binding the contract method 0x73962b26.
//
// Solidity: function isReleased(uint256 _cycle) view returns(bool)
func (_Vesting *VestingSession) IsReleased(_cycle *big.Int) (bool, error) {
	return _Vesting.Contract.IsReleased(&_Vesting.CallOpts, _cycle)
}

// IsReleased is a free data retrieval call binding the contract method 0x73962b26.
//
// Solidity: function isReleased(uint256 _cycle) view returns(bool)
func (_Vesting *VestingCallerSession) IsReleased(_cycle *big.Int) (bool, error) {
	return _Vesting.Contract.IsReleased(&_Vesting.CallOpts, _cycle)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Vesting *VestingCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Vesting.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Vesting *VestingSession) Owner() (common.Address, error) {
	return _Vesting.Contract.Owner(&_Vesting.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Vesting *VestingCallerSession) Owner() (common.Address, error) {
	return _Vesting.Contract.Owner(&_Vesting.CallOpts)
}

// Receiver is a free data retrieval call binding the contract method 0xf7260d3e.
//
// Solidity: function receiver() view returns(address)
func (_Vesting *VestingCaller) Receiver(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Vesting.contract.Call(opts, &out, "receiver")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Receiver is a free data retrieval call binding the contract method 0xf7260d3e.
//
// Solidity: function receiver() view returns(address)
func (_Vesting *VestingSession) Receiver() (common.Address, error) {
	return _Vesting.Contract.Receiver(&_Vesting.CallOpts)
}

// Receiver is a free data retrieval call binding the contract method 0xf7260d3e.
//
// Solidity: function receiver() view returns(address)
func (_Vesting *VestingCallerSession) Receiver() (common.Address, error) {
	return _Vesting.Contract.Receiver(&_Vesting.CallOpts)
}

// ReleaseAmount is a free data retrieval call binding the contract method 0xc062dc5f.
//
// Solidity: function releaseAmount() view returns(uint256)
func (_Vesting *VestingCaller) ReleaseAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vesting.contract.Call(opts, &out, "releaseAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReleaseAmount is a free data retrieval call binding the contract method 0xc062dc5f.
//
// Solidity: function releaseAmount() view returns(uint256)
func (_Vesting *VestingSession) ReleaseAmount() (*big.Int, error) {
	return _Vesting.Contract.ReleaseAmount(&_Vesting.CallOpts)
}

// ReleaseAmount is a free data retrieval call binding the contract method 0xc062dc5f.
//
// Solidity: function releaseAmount() view returns(uint256)
func (_Vesting *VestingCallerSession) ReleaseAmount() (*big.Int, error) {
	return _Vesting.Contract.ReleaseAmount(&_Vesting.CallOpts)
}

// Releases is a free data retrieval call binding the contract method 0xb6a9f40f.
//
// Solidity: function releases(uint256 ) view returns(uint256 timestamp, uint256 released)
func (_Vesting *VestingCaller) Releases(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Timestamp *big.Int
	Released  *big.Int
}, error) {
	var out []interface{}
	err := _Vesting.contract.Call(opts, &out, "releases", arg0)

	outstruct := new(struct {
		Timestamp *big.Int
		Released  *big.Int
	})

	outstruct.Timestamp = out[0].(*big.Int)
	outstruct.Released = out[1].(*big.Int)

	return *outstruct, err

}

// Releases is a free data retrieval call binding the contract method 0xb6a9f40f.
//
// Solidity: function releases(uint256 ) view returns(uint256 timestamp, uint256 released)
func (_Vesting *VestingSession) Releases(arg0 *big.Int) (struct {
	Timestamp *big.Int
	Released  *big.Int
}, error) {
	return _Vesting.Contract.Releases(&_Vesting.CallOpts, arg0)
}

// Releases is a free data retrieval call binding the contract method 0xb6a9f40f.
//
// Solidity: function releases(uint256 ) view returns(uint256 timestamp, uint256 released)
func (_Vesting *VestingCallerSession) Releases(arg0 *big.Int) (struct {
	Timestamp *big.Int
	Released  *big.Int
}, error) {
	return _Vesting.Contract.Releases(&_Vesting.CallOpts, arg0)
}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() view returns(uint256)
func (_Vesting *VestingCaller) StartTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vesting.contract.Call(opts, &out, "startTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() view returns(uint256)
func (_Vesting *VestingSession) StartTime() (*big.Int, error) {
	return _Vesting.Contract.StartTime(&_Vesting.CallOpts)
}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() view returns(uint256)
func (_Vesting *VestingCallerSession) StartTime() (*big.Int, error) {
	return _Vesting.Contract.StartTime(&_Vesting.CallOpts)
}

// Prepare is a paid mutator transaction binding the contract method 0x7c0c3e02.
//
// Solidity: function prepare(uint256 _amountToVest, address _receiver) returns()
func (_Vesting *VestingTransactor) Prepare(opts *bind.TransactOpts, _amountToVest *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Vesting.contract.Transact(opts, "prepare", _amountToVest, _receiver)
}

// Prepare is a paid mutator transaction binding the contract method 0x7c0c3e02.
//
// Solidity: function prepare(uint256 _amountToVest, address _receiver) returns()
func (_Vesting *VestingSession) Prepare(_amountToVest *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Vesting.Contract.Prepare(&_Vesting.TransactOpts, _amountToVest, _receiver)
}

// Prepare is a paid mutator transaction binding the contract method 0x7c0c3e02.
//
// Solidity: function prepare(uint256 _amountToVest, address _receiver) returns()
func (_Vesting *VestingTransactorSession) Prepare(_amountToVest *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Vesting.Contract.Prepare(&_Vesting.TransactOpts, _amountToVest, _receiver)
}

// Release is a paid mutator transaction binding the contract method 0x86d1a69f.
//
// Solidity: function release() returns()
func (_Vesting *VestingTransactor) Release(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vesting.contract.Transact(opts, "release")
}

// Release is a paid mutator transaction binding the contract method 0x86d1a69f.
//
// Solidity: function release() returns()
func (_Vesting *VestingSession) Release() (*types.Transaction, error) {
	return _Vesting.Contract.Release(&_Vesting.TransactOpts)
}

// Release is a paid mutator transaction binding the contract method 0x86d1a69f.
//
// Solidity: function release() returns()
func (_Vesting *VestingTransactorSession) Release() (*types.Transaction, error) {
	return _Vesting.Contract.Release(&_Vesting.TransactOpts)
}

// VestingTokensReleasedIterator is returned from FilterTokensReleased and is used to iterate over the raw logs and unpacked data for TokensReleased events raised by the Vesting contract.
type VestingTokensReleasedIterator struct {
	Event *VestingTokensReleased // Event containing the contract specifics and raw log

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
func (it *VestingTokensReleasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VestingTokensReleased)
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
		it.Event = new(VestingTokensReleased)
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
func (it *VestingTokensReleasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VestingTokensReleasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VestingTokensReleased represents a TokensReleased event raised by the Vesting contract.
type VestingTokensReleased struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterTokensReleased is a free log retrieval operation binding the contract event 0x3cf447e1c5a595a578bfe340289a0dad7a7f22771c57c9443ad7eedea1964d76.
//
// Solidity: event TokensReleased()
func (_Vesting *VestingFilterer) FilterTokensReleased(opts *bind.FilterOpts) (*VestingTokensReleasedIterator, error) {

	logs, sub, err := _Vesting.contract.FilterLogs(opts, "TokensReleased")
	if err != nil {
		return nil, err
	}
	return &VestingTokensReleasedIterator{contract: _Vesting.contract, event: "TokensReleased", logs: logs, sub: sub}, nil
}

// WatchTokensReleased is a free log subscription operation binding the contract event 0x3cf447e1c5a595a578bfe340289a0dad7a7f22771c57c9443ad7eedea1964d76.
//
// Solidity: event TokensReleased()
func (_Vesting *VestingFilterer) WatchTokensReleased(opts *bind.WatchOpts, sink chan<- *VestingTokensReleased) (event.Subscription, error) {

	logs, sub, err := _Vesting.contract.WatchLogs(opts, "TokensReleased")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VestingTokensReleased)
				if err := _Vesting.contract.UnpackLog(event, "TokensReleased", log); err != nil {
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

// ParseTokensReleased is a log parse operation binding the contract event 0x3cf447e1c5a595a578bfe340289a0dad7a7f22771c57c9443ad7eedea1964d76.
//
// Solidity: event TokensReleased()
func (_Vesting *VestingFilterer) ParseTokensReleased(log types.Log) (*VestingTokensReleased, error) {
	event := new(VestingTokensReleased)
	if err := _Vesting.contract.UnpackLog(event, "TokensReleased", log); err != nil {
		return nil, err
	}
	return event, nil
}
