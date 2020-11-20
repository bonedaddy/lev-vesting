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
const VestingABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_levTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_dateTimeContract\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"currentRelease\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"endTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isPrepared\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"isReleased\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amountToVest\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"prepareVest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"receiver\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"release\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"releaseAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"releases\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"released\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// VestingBin is the compiled bytecode used for deploying new contracts.
var VestingBin = "0x608060405234801561001057600080fd5b5060405161075d38038061075d8339818101604052604081101561003357600080fd5b508051602090910151600580546001600160a01b039384166001600160a01b031991821617909155600680549390921692169190911790556106e38061007a6000396000f3fe608060405234801561001057600080fd5b506004361061009e5760003560e01c8063939c28e611610066578063939c28e61461012c578063b6a9f40f14610134578063c062dc5f1461016a578063f7260d3e14610172578063fcfb9a6d146101965761009e565b80633197cbb6146100a357806369051143146100bd57806373962b26146100eb57806378e979251461011c57806386d1a69f14610124575b600080fd5b6100ab61019e565b60408051918252519081900360200190f35b6100e9600480360360408110156100d357600080fd5b50803590602001356001600160a01b03166101a4565b005b6101086004803603602081101561010157600080fd5b503561041a565b604080519115158252519081900360200190f35b6100ab610440565b6100e9610446565b61010861063c565b6101516004803603602081101561014a57600080fd5b5035610679565b6040805192835260208301919091528051918290030190f35b6100ab610692565b61017a610698565b604080516001600160a01b039092168252519081900360200190f35b6100ab6106a7565b60015481565b600554604080516323b872dd60e01b81523360048201523060248201526044810185905290516001600160a01b03909216916323b872dd916064808201926020929091908290030181600087803b1580156101fe57600080fd5b505af1158015610212573d6000803e3d6000fd5b505050506040513d602081101561022857600080fd5b505161023357600080fd5b600660009054906101000a90046001600160a01b03166001600160a01b031663b3bb8cd46040518163ffffffff1660e01b815260040160206040518083038186803b15801561028157600080fd5b505afa158015610295573d6000803e3d6000fd5b505050506040513d60208110156102ab57600080fd5b5051600081905560065460408051634355644d60e01b81526004810193909352600c6024840152516001600160a01b0390911691634355644d916044808301926020929190829003018186803b15801561030457600080fd5b505afa158015610318573d6000803e3d6000fd5b505050506040513d602081101561032e57600080fd5b5051600155600c8204600355600480546001600160a01b0319166001600160a01b03831617905560015b600c81116104155760408051808201808352600654600054634355644d60e01b9092526044830191909152606482018490529151909182916001600160a01b0390911690634355644d90608480850191602091818703018186803b1580156103bf57600080fd5b505afa1580156103d3573d6000803e3d6000fd5b505050506040513d60208110156103e957600080fd5b505181526000602091820181905283815260078252604090208251815591015160019182015501610358565b505050565b60008181526007602052604081206001908101548291141561043a575060015b92915050565b60005481565b61044e61063c565b151560011461045c57600080fd5b60065460408051632ceee33560e21b815290516000926001600160a01b03169163b3bb8cd4916004808301926020929190829003018186803b1580156104a157600080fd5b505afa1580156104b5573d6000803e3d6000fd5b505050506040513d60208110156104cb57600080fd5b50516002546000908152600760205260409020600101549091501561052a576040805162461bcd60e51b815260206004820152601060248201526f185b1c9958591e481c995b19585cd95960821b604482015290519081900360640190fd5b6002805460009081526007602052604080822060019081015591548152205481101561059d576040805162461bcd60e51b815260206004820181905260248201527f72656c656173652074696d657374616d70206e6f742079657420706173736564604482015290519081900360640190fd5b600554600480546003546040805163a9059cbb60e01b81526001600160a01b039384169481019490945260248401919091525192169163a9059cbb916044808201926020929091908290030181600087803b1580156105fb57600080fd5b505af115801561060f573d6000803e3d6000fd5b505050506040513d602081101561062557600080fd5b505161063057600080fd5b50600280546001019055565b60045460009081906001600160a01b03161580159061065d57506000600354115b801561066b57506000600254115b15610674575060015b905090565b6007602052600090815260409020805460019091015482565b60035481565b6004546001600160a01b031681565b6002548156fea26469706673582212206bae6b8957da6a2e0a2a3912db32dd349c6cd7c7646449870120dc5983eddb4964736f6c63430007040033"

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
