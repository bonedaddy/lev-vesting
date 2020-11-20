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
const VestingABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_levTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_dateTimeContract\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"ReleaseAdded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"currentRelease\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"endTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"isReleased\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amountToVest\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"prepareVest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"receiver\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"release\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"releases\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"released\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// VestingBin is the compiled bytecode used for deploying new contracts.
var VestingBin = "0x6080604052600160025534801561001557600080fd5b5060405161074a38038061074a8339818101604052604081101561003857600080fd5b508051602090910151600480546001600160a01b039384166001600160a01b031991821617909155600580549390921692169190911790556106cb8061007f6000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c806386d1a69f1161005b57806386d1a69f1461010e578063b6a9f40f14610116578063f7260d3e14610151578063fcfb9a6d1461017557610088565b80633197cbb61461008d57806369051143146100a757806373962b26146100d557806378e9792514610106575b600080fd5b61009561017d565b60408051918252519081900360200190f35b6100d3600480360360408110156100bd57600080fd5b50803590602001356001600160a01b0316610183565b005b6100f2600480360360208110156100eb57600080fd5b5035610444565b604080519115158252519081900360200190f35b61009561046a565b6100d3610470565b6101336004803603602081101561012c57600080fd5b503561065f565b60408051938452602084019290925282820152519081900360600190f35b610159610680565b604080516001600160a01b039092168252519081900360200190f35b61009561068f565b60015481565b60048054604080516323b872dd60e01b8152339381019390935230602484015260448301859052516001600160a01b03909116916323b872dd9160648083019260209291908290030181600087803b1580156101de57600080fd5b505af11580156101f2573d6000803e3d6000fd5b505050506040513d602081101561020857600080fd5b505161021357600080fd5b600560009054906101000a90046001600160a01b03166001600160a01b031663b3bb8cd46040518163ffffffff1660e01b815260040160206040518083038186803b15801561026157600080fd5b505afa158015610275573d6000803e3d6000fd5b505050506040513d602081101561028b57600080fd5b5051600081905560055460408051634355644d60e01b81526004810193909352600c6024840152516001600160a01b0390911691634355644d916044808301926020929190829003018186803b1580156102e457600080fd5b505afa1580156102f8573d6000803e3d6000fd5b505050506040513d602081101561030e57600080fd5b50516001908155600c8304905b600c811161041f576005546000805460408051634355644d60e01b81526004810192909252602482018590525191926001600160a01b031691634355644d91604480820192602092909190829003018186803b15801561037a57600080fd5b505afa15801561038e573d6000803e3d6000fd5b505050506040513d60208110156103a457600080fd5b505160408051606081018252828152602081810187815260008385018181528882526006845290859020935184559051600184015551600290920191909155815183815291519293507f0abdc1dc461ccb9bb62b1f33678a21e825c8d15e206042615b21155c5b37dfeb92918290030190a15060010161031b565b5050600380546001600160a01b0319166001600160a01b039290921691909117905550565b600081815260066020526040812060020154819060011415610464575060015b92915050565b60005481565b60055460408051632ceee33560e21b815290516000926001600160a01b03169163b3bb8cd4916004808301926020929190829003018186803b1580156104b557600080fd5b505afa1580156104c9573d6000803e3d6000fd5b505050506040513d60208110156104df57600080fd5b505160028054600090815260066020526040902001549091501561053d576040805162461bcd60e51b815260206004820152601060248201526f185b1c9958591e481c995b19585cd95960821b604482015290519081900360640190fd5b600280546000908152600660205260408082206001908401559154815220548110156105b0576040805162461bcd60e51b815260206004820181905260248201527f72656c656173652074696d657374616d70206e6f742079657420706173736564604482015290519081900360640190fd5b60048054600354600254600090815260066020908152604080832060010154815163a9059cbb60e01b81526001600160a01b0395861697810197909752602487015251929093169363a9059cbb93604480830194919391928390030190829087803b15801561061e57600080fd5b505af1158015610632573d6000803e3d6000fd5b505050506040513d602081101561064857600080fd5b505161065357600080fd5b50600280546001019055565b60066020526000908152604090208054600182015460029092015490919083565b6003546001600160a01b031681565b6002548156fea264697066735822122040460ee2fd38c282d44209991b4ee50b32d2a80d055bf33aa88f26790a50f1a364736f6c63430007040033"

// DeployVesting deploys a new Ethereum contract, binding an instance of Vesting to it.
func DeployVesting(auth *bind.TransactOpts, backend bind.ContractBackend, _levTokenAddress common.Address, _dateTimeContract common.Address) (common.Address, *types.Transaction, *Vesting, error) {
	parsed, err := abi.JSON(strings.NewReader(VestingABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(VestingBin), backend, _levTokenAddress, _dateTimeContract)
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

// CurrentRelease is a free data retrieval call binding the contract method 0xfcfb9a6d.
//
// Solidity: function currentRelease() view returns(uint256)
func (_Vesting *VestingCaller) CurrentRelease(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vesting.contract.Call(opts, &out, "currentRelease")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentRelease is a free data retrieval call binding the contract method 0xfcfb9a6d.
//
// Solidity: function currentRelease() view returns(uint256)
func (_Vesting *VestingSession) CurrentRelease() (*big.Int, error) {
	return _Vesting.Contract.CurrentRelease(&_Vesting.CallOpts)
}

// CurrentRelease is a free data retrieval call binding the contract method 0xfcfb9a6d.
//
// Solidity: function currentRelease() view returns(uint256)
func (_Vesting *VestingCallerSession) CurrentRelease() (*big.Int, error) {
	return _Vesting.Contract.CurrentRelease(&_Vesting.CallOpts)
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

// IsReleased is a free data retrieval call binding the contract method 0x73962b26.
//
// Solidity: function isReleased(uint256 _index) view returns(bool)
func (_Vesting *VestingCaller) IsReleased(opts *bind.CallOpts, _index *big.Int) (bool, error) {
	var out []interface{}
	err := _Vesting.contract.Call(opts, &out, "isReleased", _index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsReleased is a free data retrieval call binding the contract method 0x73962b26.
//
// Solidity: function isReleased(uint256 _index) view returns(bool)
func (_Vesting *VestingSession) IsReleased(_index *big.Int) (bool, error) {
	return _Vesting.Contract.IsReleased(&_Vesting.CallOpts, _index)
}

// IsReleased is a free data retrieval call binding the contract method 0x73962b26.
//
// Solidity: function isReleased(uint256 _index) view returns(bool)
func (_Vesting *VestingCallerSession) IsReleased(_index *big.Int) (bool, error) {
	return _Vesting.Contract.IsReleased(&_Vesting.CallOpts, _index)
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

// Releases is a free data retrieval call binding the contract method 0xb6a9f40f.
//
// Solidity: function releases(uint256 ) view returns(uint256 timestamp, uint256 amount, uint256 released)
func (_Vesting *VestingCaller) Releases(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Timestamp *big.Int
	Amount    *big.Int
	Released  *big.Int
}, error) {
	var out []interface{}
	err := _Vesting.contract.Call(opts, &out, "releases", arg0)

	outstruct := new(struct {
		Timestamp *big.Int
		Amount    *big.Int
		Released  *big.Int
	})

	outstruct.Timestamp = out[0].(*big.Int)
	outstruct.Amount = out[1].(*big.Int)
	outstruct.Released = out[2].(*big.Int)

	return *outstruct, err

}

// Releases is a free data retrieval call binding the contract method 0xb6a9f40f.
//
// Solidity: function releases(uint256 ) view returns(uint256 timestamp, uint256 amount, uint256 released)
func (_Vesting *VestingSession) Releases(arg0 *big.Int) (struct {
	Timestamp *big.Int
	Amount    *big.Int
	Released  *big.Int
}, error) {
	return _Vesting.Contract.Releases(&_Vesting.CallOpts, arg0)
}

// Releases is a free data retrieval call binding the contract method 0xb6a9f40f.
//
// Solidity: function releases(uint256 ) view returns(uint256 timestamp, uint256 amount, uint256 released)
func (_Vesting *VestingCallerSession) Releases(arg0 *big.Int) (struct {
	Timestamp *big.Int
	Amount    *big.Int
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

// PrepareVest is a paid mutator transaction binding the contract method 0x69051143.
//
// Solidity: function prepareVest(uint256 _amountToVest, address _receiver) returns()
func (_Vesting *VestingTransactor) PrepareVest(opts *bind.TransactOpts, _amountToVest *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Vesting.contract.Transact(opts, "prepareVest", _amountToVest, _receiver)
}

// PrepareVest is a paid mutator transaction binding the contract method 0x69051143.
//
// Solidity: function prepareVest(uint256 _amountToVest, address _receiver) returns()
func (_Vesting *VestingSession) PrepareVest(_amountToVest *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Vesting.Contract.PrepareVest(&_Vesting.TransactOpts, _amountToVest, _receiver)
}

// PrepareVest is a paid mutator transaction binding the contract method 0x69051143.
//
// Solidity: function prepareVest(uint256 _amountToVest, address _receiver) returns()
func (_Vesting *VestingTransactorSession) PrepareVest(_amountToVest *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Vesting.Contract.PrepareVest(&_Vesting.TransactOpts, _amountToVest, _receiver)
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

// VestingReleaseAddedIterator is returned from FilterReleaseAdded and is used to iterate over the raw logs and unpacked data for ReleaseAdded events raised by the Vesting contract.
type VestingReleaseAddedIterator struct {
	Event *VestingReleaseAdded // Event containing the contract specifics and raw log

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
func (it *VestingReleaseAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VestingReleaseAdded)
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
		it.Event = new(VestingReleaseAdded)
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
func (it *VestingReleaseAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VestingReleaseAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VestingReleaseAdded represents a ReleaseAdded event raised by the Vesting contract.
type VestingReleaseAdded struct {
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterReleaseAdded is a free log retrieval operation binding the contract event 0x0abdc1dc461ccb9bb62b1f33678a21e825c8d15e206042615b21155c5b37dfeb.
//
// Solidity: event ReleaseAdded(uint256 timestamp)
func (_Vesting *VestingFilterer) FilterReleaseAdded(opts *bind.FilterOpts) (*VestingReleaseAddedIterator, error) {

	logs, sub, err := _Vesting.contract.FilterLogs(opts, "ReleaseAdded")
	if err != nil {
		return nil, err
	}
	return &VestingReleaseAddedIterator{contract: _Vesting.contract, event: "ReleaseAdded", logs: logs, sub: sub}, nil
}

// WatchReleaseAdded is a free log subscription operation binding the contract event 0x0abdc1dc461ccb9bb62b1f33678a21e825c8d15e206042615b21155c5b37dfeb.
//
// Solidity: event ReleaseAdded(uint256 timestamp)
func (_Vesting *VestingFilterer) WatchReleaseAdded(opts *bind.WatchOpts, sink chan<- *VestingReleaseAdded) (event.Subscription, error) {

	logs, sub, err := _Vesting.contract.WatchLogs(opts, "ReleaseAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VestingReleaseAdded)
				if err := _Vesting.contract.UnpackLog(event, "ReleaseAdded", log); err != nil {
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

// ParseReleaseAdded is a log parse operation binding the contract event 0x0abdc1dc461ccb9bb62b1f33678a21e825c8d15e206042615b21155c5b37dfeb.
//
// Solidity: event ReleaseAdded(uint256 timestamp)
func (_Vesting *VestingFilterer) ParseReleaseAdded(log types.Log) (*VestingReleaseAdded, error) {
	event := new(VestingReleaseAdded)
	if err := _Vesting.contract.UnpackLog(event, "ReleaseAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}