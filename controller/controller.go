// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package controller

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
)

// AddressMetaData contains all meta data concerning the Address contract.
var AddressMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122091878a8a7eda7bf6a23bcb6736830e7e73a2998d9b8ead2ffe02236659425e0364736f6c634300060c0033",
}

// AddressABI is the input ABI used to generate the binding from.
// Deprecated: Use AddressMetaData.ABI instead.
var AddressABI = AddressMetaData.ABI

// AddressBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AddressMetaData.Bin instead.
var AddressBin = AddressMetaData.Bin

// DeployAddress deploys a new Ethereum contract, binding an instance of Address to it.
func DeployAddress(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Address, error) {
	parsed, err := AddressMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AddressBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Address{AddressCaller: AddressCaller{contract: contract}, AddressTransactor: AddressTransactor{contract: contract}, AddressFilterer: AddressFilterer{contract: contract}}, nil
}

// Address is an auto generated Go binding around an Ethereum contract.
type Address struct {
	AddressCaller     // Read-only binding to the contract
	AddressTransactor // Write-only binding to the contract
	AddressFilterer   // Log filterer for contract events
}

// AddressCaller is an auto generated read-only Go binding around an Ethereum contract.
type AddressCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AddressTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AddressFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AddressSession struct {
	Contract     *Address          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AddressCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AddressCallerSession struct {
	Contract *AddressCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// AddressTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AddressTransactorSession struct {
	Contract     *AddressTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// AddressRaw is an auto generated low-level Go binding around an Ethereum contract.
type AddressRaw struct {
	Contract *Address // Generic contract binding to access the raw methods on
}

// AddressCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AddressCallerRaw struct {
	Contract *AddressCaller // Generic read-only contract binding to access the raw methods on
}

// AddressTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AddressTransactorRaw struct {
	Contract *AddressTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAddress creates a new instance of Address, bound to a specific deployed contract.
func NewAddress(address common.Address, backend bind.ContractBackend) (*Address, error) {
	contract, err := bindAddress(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Address{AddressCaller: AddressCaller{contract: contract}, AddressTransactor: AddressTransactor{contract: contract}, AddressFilterer: AddressFilterer{contract: contract}}, nil
}

// NewAddressCaller creates a new read-only instance of Address, bound to a specific deployed contract.
func NewAddressCaller(address common.Address, caller bind.ContractCaller) (*AddressCaller, error) {
	contract, err := bindAddress(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AddressCaller{contract: contract}, nil
}

// NewAddressTransactor creates a new write-only instance of Address, bound to a specific deployed contract.
func NewAddressTransactor(address common.Address, transactor bind.ContractTransactor) (*AddressTransactor, error) {
	contract, err := bindAddress(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AddressTransactor{contract: contract}, nil
}

// NewAddressFilterer creates a new log filterer instance of Address, bound to a specific deployed contract.
func NewAddressFilterer(address common.Address, filterer bind.ContractFilterer) (*AddressFilterer, error) {
	contract, err := bindAddress(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AddressFilterer{contract: contract}, nil
}

// bindAddress binds a generic wrapper to an already deployed contract.
func bindAddress(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AddressABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Address *AddressRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Address.Contract.AddressCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Address *AddressRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Address.Contract.AddressTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Address *AddressRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Address.Contract.AddressTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Address *AddressCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Address.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Address *AddressTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Address.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Address *AddressTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Address.Contract.contract.Transact(opts, method, params...)
}

// ControllerMetaData contains all meta data concerning the Controller contract.
var ControllerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"ManagerList\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"MarketList\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_cAddr\",\"type\":\"address\"}],\"name\":\"addAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_contractAddr\",\"type\":\"address\"}],\"name\":\"addContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addrM\",\"type\":\"address\"}],\"name\":\"addManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addrM\",\"type\":\"address\"}],\"name\":\"addMarket\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"addrList\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"delContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addrM\",\"type\":\"address\"}],\"name\":\"delManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addrM\",\"type\":\"address\"}],\"name\":\"delMarket\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"getAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_mAddr\",\"type\":\"address\"}],\"name\":\"isManager\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_mAddr\",\"type\":\"address\"}],\"name\":\"isMarket\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"1d1261e8": "ManagerList(address)",
		"e74cb409": "MarketList(address)",
		"ceb35b0f": "addAddress(string,address)",
		"bf5b6016": "addContract(string,address)",
		"2d06177a": "addManager(address)",
		"93e30633": "addMarket(address)",
		"92791661": "addrList(string)",
		"a6f9dae1": "changeOwner(address)",
		"7def3104": "delContract(string)",
		"a65eacdc": "delManager(address)",
		"829fe8da": "delMarket(address)",
		"d502db97": "getAddr(string)",
		"f3ae2415": "isManager(address)",
		"6ec934da": "isMarket(address)",
		"8da5cb5b": "owner()",
	},
	Bin: "0x608060405234801561001057600080fd5b50604051610baa380380610baa8339818101604052602081101561003357600080fd5b5051600080546001600160a01b039092166001600160a01b0319909216919091179055610b45806100656000396000f3fe608060405234801561001057600080fd5b50600436106100f55760003560e01c806393e3063311610097578063ceb35b0f11610066578063ceb35b0f14610435578063d502db97146104e4578063e74cb40914610552578063f3ae241514610578576100f5565b806393e3063314610314578063a65eacdc1461033a578063a6f9dae114610360578063bf5b601614610386576100f5565b80637def3104116100d35780637def310414610182578063829fe8da146102265780638da5cb5b1461024c5780639279166114610270576100f5565b80631d1261e8146100fa5780632d06177a146101345780636ec934da1461015c575b600080fd5b6101206004803603602081101561011057600080fd5b50356001600160a01b031661059e565b604080519115158252519081900360200190f35b61015a6004803603602081101561014a57600080fd5b50356001600160a01b03166105b3565b005b6101206004803603602081101561017257600080fd5b50356001600160a01b0316610623565b61015a6004803603602081101561019857600080fd5b810190602081018135600160201b8111156101b257600080fd5b8201836020820111156101c457600080fd5b803590602001918460018302840111600160201b831117156101e557600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610641945050505050565b61015a6004803603602081101561023c57600080fd5b50356001600160a01b0316610710565b61025461077d565b604080516001600160a01b039092168252519081900360200190f35b6102546004803603602081101561028657600080fd5b810190602081018135600160201b8111156102a057600080fd5b8201836020820111156102b257600080fd5b803590602001918460018302840111600160201b831117156102d357600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955061078c945050505050565b61015a6004803603602081101561032a57600080fd5b50356001600160a01b03166107b2565b61015a6004803603602081101561035057600080fd5b50356001600160a01b0316610822565b61015a6004803603602081101561037657600080fd5b50356001600160a01b031661088f565b61015a6004803603604081101561039c57600080fd5b810190602081018135600160201b8111156103b657600080fd5b8201836020820111156103c857600080fd5b803590602001918460018302840111600160201b831117156103e957600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550505090356001600160a01b031691506108fd9050565b61015a6004803603604081101561044b57600080fd5b810190602081018135600160201b81111561046557600080fd5b82018360208201111561047757600080fd5b803590602001918460018302840111600160201b8311171561049857600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550505090356001600160a01b03169150610a1a9050565b610254600480360360208110156104fa57600080fd5b810190602081018135600160201b81111561051457600080fd5b82018360208201111561052657600080fd5b803590602001918460018302840111600160201b8311171561054757600080fd5b509092509050610a66565b6101206004803603602081101561056857600080fd5b50356001600160a01b0316610aa0565b6101206004803603602081101561058e57600080fd5b50356001600160a01b0316610ab5565b60026020526000908152604090205460ff1681565b6000546001600160a01b031633146105ff576040805162461bcd60e51b815260206004820152600a6024820152693737ba1039b2ba3a32b960b11b604482015290519081900360640190fd5b6001600160a01b03166000908152600260205260409020805460ff19166001179055565b6001600160a01b031660009081526003602052604090205460ff1690565b6000546001600160a01b0316331461068d576040805162461bcd60e51b815260206004820152600a6024820152693737ba1039b2ba3a32b960b11b604482015290519081900360640190fd5b60006001826040518082805190602001908083835b602083106106c15780518252601f1990920191602091820191016106a2565b51815160209384036101000a6000190180199092169116179052920194855250604051938490030190922080546001600160a01b0319166001600160a01b039490941693909317909255505050565b6000546001600160a01b0316331461075c576040805162461bcd60e51b815260206004820152600a6024820152693737ba1039b2ba3a32b960b11b604482015290519081900360640190fd5b6001600160a01b03166000908152600360205260409020805460ff19169055565b6000546001600160a01b031681565b80516020818301810180516001825292820191909301209152546001600160a01b031681565b6000546001600160a01b031633146107fe576040805162461bcd60e51b815260206004820152600a6024820152693737ba1039b2ba3a32b960b11b604482015290519081900360640190fd5b6001600160a01b03166000908152600360205260409020805460ff19166001179055565b6000546001600160a01b0316331461086e576040805162461bcd60e51b815260206004820152600a6024820152693737ba1039b2ba3a32b960b11b604482015290519081900360640190fd5b6001600160a01b03166000908152600260205260409020805460ff19169055565b6000546001600160a01b031633146108db576040805162461bcd60e51b815260206004820152600a6024820152693737ba1039b2ba3a32b960b11b604482015290519081900360640190fd5b600080546001600160a01b0319166001600160a01b0392909216919091179055565b6000546001600160a01b03163314610949576040805162461bcd60e51b815260206004820152600a6024820152693737ba1039b2ba3a32b960b11b604482015290519081900360640190fd5b61095281610ad3565b610997576040805162461bcd60e51b81526020600482015260116024820152703737ba1031b7b73a3930b1ba1020b2323960791b604482015290519081900360640190fd5b806001836040518082805190602001908083835b602083106109ca5780518252601f1990920191602091820191016109ab565b51815160209384036101000a6000190180199092169116179052920194855250604051938490030190922080546001600160a01b0319166001600160a01b03949094169390931790925550505050565b6000546001600160a01b03163314610997576040805162461bcd60e51b815260206004820152600a6024820152693737ba1039b2ba3a32b960b11b604482015290519081900360640190fd5b6000600183836040518083838082843791909101948552505060405192839003602001909220546001600160a01b03169250505092915050565b60036020526000908152604090205460ff1681565b6001600160a01b031660009081526002602052604090205460ff1690565b6000813f7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a4708115801590610b075750808214155b94935050505056fea2646970667358221220c6dbeda81cd5999f6f9e354db847d90e7c70a9308b82924c58c4e077ebc7542064736f6c634300060c0033",
}

// ControllerABI is the input ABI used to generate the binding from.
// Deprecated: Use ControllerMetaData.ABI instead.
var ControllerABI = ControllerMetaData.ABI

// Deprecated: Use ControllerMetaData.Sigs instead.
// ControllerFuncSigs maps the 4-byte function signature to its string representation.
var ControllerFuncSigs = ControllerMetaData.Sigs

// ControllerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ControllerMetaData.Bin instead.
var ControllerBin = ControllerMetaData.Bin

// DeployController deploys a new Ethereum contract, binding an instance of Controller to it.
func DeployController(auth *bind.TransactOpts, backend bind.ContractBackend, _owner common.Address) (common.Address, *types.Transaction, *Controller, error) {
	parsed, err := ControllerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ControllerBin), backend, _owner)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Controller{ControllerCaller: ControllerCaller{contract: contract}, ControllerTransactor: ControllerTransactor{contract: contract}, ControllerFilterer: ControllerFilterer{contract: contract}}, nil
}

// Controller is an auto generated Go binding around an Ethereum contract.
type Controller struct {
	ControllerCaller     // Read-only binding to the contract
	ControllerTransactor // Write-only binding to the contract
	ControllerFilterer   // Log filterer for contract events
}

// ControllerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ControllerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ControllerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ControllerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ControllerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ControllerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ControllerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ControllerSession struct {
	Contract     *Controller       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ControllerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ControllerCallerSession struct {
	Contract *ControllerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ControllerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ControllerTransactorSession struct {
	Contract     *ControllerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ControllerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ControllerRaw struct {
	Contract *Controller // Generic contract binding to access the raw methods on
}

// ControllerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ControllerCallerRaw struct {
	Contract *ControllerCaller // Generic read-only contract binding to access the raw methods on
}

// ControllerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ControllerTransactorRaw struct {
	Contract *ControllerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewController creates a new instance of Controller, bound to a specific deployed contract.
func NewController(address common.Address, backend bind.ContractBackend) (*Controller, error) {
	contract, err := bindController(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Controller{ControllerCaller: ControllerCaller{contract: contract}, ControllerTransactor: ControllerTransactor{contract: contract}, ControllerFilterer: ControllerFilterer{contract: contract}}, nil
}

// NewControllerCaller creates a new read-only instance of Controller, bound to a specific deployed contract.
func NewControllerCaller(address common.Address, caller bind.ContractCaller) (*ControllerCaller, error) {
	contract, err := bindController(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ControllerCaller{contract: contract}, nil
}

// NewControllerTransactor creates a new write-only instance of Controller, bound to a specific deployed contract.
func NewControllerTransactor(address common.Address, transactor bind.ContractTransactor) (*ControllerTransactor, error) {
	contract, err := bindController(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ControllerTransactor{contract: contract}, nil
}

// NewControllerFilterer creates a new log filterer instance of Controller, bound to a specific deployed contract.
func NewControllerFilterer(address common.Address, filterer bind.ContractFilterer) (*ControllerFilterer, error) {
	contract, err := bindController(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ControllerFilterer{contract: contract}, nil
}

// bindController binds a generic wrapper to an already deployed contract.
func bindController(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ControllerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Controller *ControllerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Controller.Contract.ControllerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Controller *ControllerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Controller.Contract.ControllerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Controller *ControllerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Controller.Contract.ControllerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Controller *ControllerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Controller.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Controller *ControllerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Controller.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Controller *ControllerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Controller.Contract.contract.Transact(opts, method, params...)
}

// ManagerList is a free data retrieval call binding the contract method 0x1d1261e8.
//
// Solidity: function ManagerList(address ) view returns(bool)
func (_Controller *ControllerCaller) ManagerList(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "ManagerList", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ManagerList is a free data retrieval call binding the contract method 0x1d1261e8.
//
// Solidity: function ManagerList(address ) view returns(bool)
func (_Controller *ControllerSession) ManagerList(arg0 common.Address) (bool, error) {
	return _Controller.Contract.ManagerList(&_Controller.CallOpts, arg0)
}

// ManagerList is a free data retrieval call binding the contract method 0x1d1261e8.
//
// Solidity: function ManagerList(address ) view returns(bool)
func (_Controller *ControllerCallerSession) ManagerList(arg0 common.Address) (bool, error) {
	return _Controller.Contract.ManagerList(&_Controller.CallOpts, arg0)
}

// MarketList is a free data retrieval call binding the contract method 0xe74cb409.
//
// Solidity: function MarketList(address ) view returns(bool)
func (_Controller *ControllerCaller) MarketList(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "MarketList", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MarketList is a free data retrieval call binding the contract method 0xe74cb409.
//
// Solidity: function MarketList(address ) view returns(bool)
func (_Controller *ControllerSession) MarketList(arg0 common.Address) (bool, error) {
	return _Controller.Contract.MarketList(&_Controller.CallOpts, arg0)
}

// MarketList is a free data retrieval call binding the contract method 0xe74cb409.
//
// Solidity: function MarketList(address ) view returns(bool)
func (_Controller *ControllerCallerSession) MarketList(arg0 common.Address) (bool, error) {
	return _Controller.Contract.MarketList(&_Controller.CallOpts, arg0)
}

// AddrList is a free data retrieval call binding the contract method 0x92791661.
//
// Solidity: function addrList(string ) view returns(address)
func (_Controller *ControllerCaller) AddrList(opts *bind.CallOpts, arg0 string) (common.Address, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "addrList", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddrList is a free data retrieval call binding the contract method 0x92791661.
//
// Solidity: function addrList(string ) view returns(address)
func (_Controller *ControllerSession) AddrList(arg0 string) (common.Address, error) {
	return _Controller.Contract.AddrList(&_Controller.CallOpts, arg0)
}

// AddrList is a free data retrieval call binding the contract method 0x92791661.
//
// Solidity: function addrList(string ) view returns(address)
func (_Controller *ControllerCallerSession) AddrList(arg0 string) (common.Address, error) {
	return _Controller.Contract.AddrList(&_Controller.CallOpts, arg0)
}

// GetAddr is a free data retrieval call binding the contract method 0xd502db97.
//
// Solidity: function getAddr(string _name) view returns(address)
func (_Controller *ControllerCaller) GetAddr(opts *bind.CallOpts, _name string) (common.Address, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "getAddr", _name)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddr is a free data retrieval call binding the contract method 0xd502db97.
//
// Solidity: function getAddr(string _name) view returns(address)
func (_Controller *ControllerSession) GetAddr(_name string) (common.Address, error) {
	return _Controller.Contract.GetAddr(&_Controller.CallOpts, _name)
}

// GetAddr is a free data retrieval call binding the contract method 0xd502db97.
//
// Solidity: function getAddr(string _name) view returns(address)
func (_Controller *ControllerCallerSession) GetAddr(_name string) (common.Address, error) {
	return _Controller.Contract.GetAddr(&_Controller.CallOpts, _name)
}

// IsManager is a free data retrieval call binding the contract method 0xf3ae2415.
//
// Solidity: function isManager(address _mAddr) view returns(bool)
func (_Controller *ControllerCaller) IsManager(opts *bind.CallOpts, _mAddr common.Address) (bool, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "isManager", _mAddr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsManager is a free data retrieval call binding the contract method 0xf3ae2415.
//
// Solidity: function isManager(address _mAddr) view returns(bool)
func (_Controller *ControllerSession) IsManager(_mAddr common.Address) (bool, error) {
	return _Controller.Contract.IsManager(&_Controller.CallOpts, _mAddr)
}

// IsManager is a free data retrieval call binding the contract method 0xf3ae2415.
//
// Solidity: function isManager(address _mAddr) view returns(bool)
func (_Controller *ControllerCallerSession) IsManager(_mAddr common.Address) (bool, error) {
	return _Controller.Contract.IsManager(&_Controller.CallOpts, _mAddr)
}

// IsMarket is a free data retrieval call binding the contract method 0x6ec934da.
//
// Solidity: function isMarket(address _mAddr) view returns(bool)
func (_Controller *ControllerCaller) IsMarket(opts *bind.CallOpts, _mAddr common.Address) (bool, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "isMarket", _mAddr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMarket is a free data retrieval call binding the contract method 0x6ec934da.
//
// Solidity: function isMarket(address _mAddr) view returns(bool)
func (_Controller *ControllerSession) IsMarket(_mAddr common.Address) (bool, error) {
	return _Controller.Contract.IsMarket(&_Controller.CallOpts, _mAddr)
}

// IsMarket is a free data retrieval call binding the contract method 0x6ec934da.
//
// Solidity: function isMarket(address _mAddr) view returns(bool)
func (_Controller *ControllerCallerSession) IsMarket(_mAddr common.Address) (bool, error) {
	return _Controller.Contract.IsMarket(&_Controller.CallOpts, _mAddr)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Controller *ControllerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Controller *ControllerSession) Owner() (common.Address, error) {
	return _Controller.Contract.Owner(&_Controller.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Controller *ControllerCallerSession) Owner() (common.Address, error) {
	return _Controller.Contract.Owner(&_Controller.CallOpts)
}

// AddAddress is a paid mutator transaction binding the contract method 0xceb35b0f.
//
// Solidity: function addAddress(string _name, address _cAddr) returns()
func (_Controller *ControllerTransactor) AddAddress(opts *bind.TransactOpts, _name string, _cAddr common.Address) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "addAddress", _name, _cAddr)
}

// AddAddress is a paid mutator transaction binding the contract method 0xceb35b0f.
//
// Solidity: function addAddress(string _name, address _cAddr) returns()
func (_Controller *ControllerSession) AddAddress(_name string, _cAddr common.Address) (*types.Transaction, error) {
	return _Controller.Contract.AddAddress(&_Controller.TransactOpts, _name, _cAddr)
}

// AddAddress is a paid mutator transaction binding the contract method 0xceb35b0f.
//
// Solidity: function addAddress(string _name, address _cAddr) returns()
func (_Controller *ControllerTransactorSession) AddAddress(_name string, _cAddr common.Address) (*types.Transaction, error) {
	return _Controller.Contract.AddAddress(&_Controller.TransactOpts, _name, _cAddr)
}

// AddContract is a paid mutator transaction binding the contract method 0xbf5b6016.
//
// Solidity: function addContract(string _name, address _contractAddr) returns()
func (_Controller *ControllerTransactor) AddContract(opts *bind.TransactOpts, _name string, _contractAddr common.Address) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "addContract", _name, _contractAddr)
}

// AddContract is a paid mutator transaction binding the contract method 0xbf5b6016.
//
// Solidity: function addContract(string _name, address _contractAddr) returns()
func (_Controller *ControllerSession) AddContract(_name string, _contractAddr common.Address) (*types.Transaction, error) {
	return _Controller.Contract.AddContract(&_Controller.TransactOpts, _name, _contractAddr)
}

// AddContract is a paid mutator transaction binding the contract method 0xbf5b6016.
//
// Solidity: function addContract(string _name, address _contractAddr) returns()
func (_Controller *ControllerTransactorSession) AddContract(_name string, _contractAddr common.Address) (*types.Transaction, error) {
	return _Controller.Contract.AddContract(&_Controller.TransactOpts, _name, _contractAddr)
}

// AddManager is a paid mutator transaction binding the contract method 0x2d06177a.
//
// Solidity: function addManager(address _addrM) returns()
func (_Controller *ControllerTransactor) AddManager(opts *bind.TransactOpts, _addrM common.Address) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "addManager", _addrM)
}

// AddManager is a paid mutator transaction binding the contract method 0x2d06177a.
//
// Solidity: function addManager(address _addrM) returns()
func (_Controller *ControllerSession) AddManager(_addrM common.Address) (*types.Transaction, error) {
	return _Controller.Contract.AddManager(&_Controller.TransactOpts, _addrM)
}

// AddManager is a paid mutator transaction binding the contract method 0x2d06177a.
//
// Solidity: function addManager(address _addrM) returns()
func (_Controller *ControllerTransactorSession) AddManager(_addrM common.Address) (*types.Transaction, error) {
	return _Controller.Contract.AddManager(&_Controller.TransactOpts, _addrM)
}

// AddMarket is a paid mutator transaction binding the contract method 0x93e30633.
//
// Solidity: function addMarket(address _addrM) returns()
func (_Controller *ControllerTransactor) AddMarket(opts *bind.TransactOpts, _addrM common.Address) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "addMarket", _addrM)
}

// AddMarket is a paid mutator transaction binding the contract method 0x93e30633.
//
// Solidity: function addMarket(address _addrM) returns()
func (_Controller *ControllerSession) AddMarket(_addrM common.Address) (*types.Transaction, error) {
	return _Controller.Contract.AddMarket(&_Controller.TransactOpts, _addrM)
}

// AddMarket is a paid mutator transaction binding the contract method 0x93e30633.
//
// Solidity: function addMarket(address _addrM) returns()
func (_Controller *ControllerTransactorSession) AddMarket(_addrM common.Address) (*types.Transaction, error) {
	return _Controller.Contract.AddMarket(&_Controller.TransactOpts, _addrM)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address _owner) returns()
func (_Controller *ControllerTransactor) ChangeOwner(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "changeOwner", _owner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address _owner) returns()
func (_Controller *ControllerSession) ChangeOwner(_owner common.Address) (*types.Transaction, error) {
	return _Controller.Contract.ChangeOwner(&_Controller.TransactOpts, _owner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address _owner) returns()
func (_Controller *ControllerTransactorSession) ChangeOwner(_owner common.Address) (*types.Transaction, error) {
	return _Controller.Contract.ChangeOwner(&_Controller.TransactOpts, _owner)
}

// DelContract is a paid mutator transaction binding the contract method 0x7def3104.
//
// Solidity: function delContract(string _name) returns()
func (_Controller *ControllerTransactor) DelContract(opts *bind.TransactOpts, _name string) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "delContract", _name)
}

// DelContract is a paid mutator transaction binding the contract method 0x7def3104.
//
// Solidity: function delContract(string _name) returns()
func (_Controller *ControllerSession) DelContract(_name string) (*types.Transaction, error) {
	return _Controller.Contract.DelContract(&_Controller.TransactOpts, _name)
}

// DelContract is a paid mutator transaction binding the contract method 0x7def3104.
//
// Solidity: function delContract(string _name) returns()
func (_Controller *ControllerTransactorSession) DelContract(_name string) (*types.Transaction, error) {
	return _Controller.Contract.DelContract(&_Controller.TransactOpts, _name)
}

// DelManager is a paid mutator transaction binding the contract method 0xa65eacdc.
//
// Solidity: function delManager(address _addrM) returns()
func (_Controller *ControllerTransactor) DelManager(opts *bind.TransactOpts, _addrM common.Address) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "delManager", _addrM)
}

// DelManager is a paid mutator transaction binding the contract method 0xa65eacdc.
//
// Solidity: function delManager(address _addrM) returns()
func (_Controller *ControllerSession) DelManager(_addrM common.Address) (*types.Transaction, error) {
	return _Controller.Contract.DelManager(&_Controller.TransactOpts, _addrM)
}

// DelManager is a paid mutator transaction binding the contract method 0xa65eacdc.
//
// Solidity: function delManager(address _addrM) returns()
func (_Controller *ControllerTransactorSession) DelManager(_addrM common.Address) (*types.Transaction, error) {
	return _Controller.Contract.DelManager(&_Controller.TransactOpts, _addrM)
}

// DelMarket is a paid mutator transaction binding the contract method 0x829fe8da.
//
// Solidity: function delMarket(address _addrM) returns()
func (_Controller *ControllerTransactor) DelMarket(opts *bind.TransactOpts, _addrM common.Address) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "delMarket", _addrM)
}

// DelMarket is a paid mutator transaction binding the contract method 0x829fe8da.
//
// Solidity: function delMarket(address _addrM) returns()
func (_Controller *ControllerSession) DelMarket(_addrM common.Address) (*types.Transaction, error) {
	return _Controller.Contract.DelMarket(&_Controller.TransactOpts, _addrM)
}

// DelMarket is a paid mutator transaction binding the contract method 0x829fe8da.
//
// Solidity: function delMarket(address _addrM) returns()
func (_Controller *ControllerTransactorSession) DelMarket(_addrM common.Address) (*types.Transaction, error) {
	return _Controller.Contract.DelMarket(&_Controller.TransactOpts, _addrM)
}

