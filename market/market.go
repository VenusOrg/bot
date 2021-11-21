// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package market

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
	Bin: "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122099b0682def54149baffcf245b5f8e4e70f7e6d10bbb01290df0ae9a2f3f9463464736f6c634300060c0033",
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

// SafeMathMetaData contains all meta data concerning the SafeMath contract.
var SafeMathMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122007befcfde2ca95ab9f77fcc1b9506af41c7c186be8cf4f01e4f58a725b5b9ee364736f6c634300060c0033",
}

// SafeMathABI is the input ABI used to generate the binding from.
// Deprecated: Use SafeMathMetaData.ABI instead.
var SafeMathABI = SafeMathMetaData.ABI

// SafeMathBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SafeMathMetaData.Bin instead.
var SafeMathBin = SafeMathMetaData.Bin

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := SafeMathMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}

// TransferHelperMetaData contains all meta data concerning the TransferHelper contract.
var TransferHelperMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212205440f997a139c15474e6fb293ea928c9a8db539fd1c676cb38826fc6490e4f6c64736f6c634300060c0033",
}

// TransferHelperABI is the input ABI used to generate the binding from.
// Deprecated: Use TransferHelperMetaData.ABI instead.
var TransferHelperABI = TransferHelperMetaData.ABI

// TransferHelperBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TransferHelperMetaData.Bin instead.
var TransferHelperBin = TransferHelperMetaData.Bin

// DeployTransferHelper deploys a new Ethereum contract, binding an instance of TransferHelper to it.
func DeployTransferHelper(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TransferHelper, error) {
	parsed, err := TransferHelperMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TransferHelperBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TransferHelper{TransferHelperCaller: TransferHelperCaller{contract: contract}, TransferHelperTransactor: TransferHelperTransactor{contract: contract}, TransferHelperFilterer: TransferHelperFilterer{contract: contract}}, nil
}

// TransferHelper is an auto generated Go binding around an Ethereum contract.
type TransferHelper struct {
	TransferHelperCaller     // Read-only binding to the contract
	TransferHelperTransactor // Write-only binding to the contract
	TransferHelperFilterer   // Log filterer for contract events
}

// TransferHelperCaller is an auto generated read-only Go binding around an Ethereum contract.
type TransferHelperCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransferHelperTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TransferHelperTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransferHelperFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TransferHelperFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransferHelperSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TransferHelperSession struct {
	Contract     *TransferHelper   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TransferHelperCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TransferHelperCallerSession struct {
	Contract *TransferHelperCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// TransferHelperTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TransferHelperTransactorSession struct {
	Contract     *TransferHelperTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// TransferHelperRaw is an auto generated low-level Go binding around an Ethereum contract.
type TransferHelperRaw struct {
	Contract *TransferHelper // Generic contract binding to access the raw methods on
}

// TransferHelperCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TransferHelperCallerRaw struct {
	Contract *TransferHelperCaller // Generic read-only contract binding to access the raw methods on
}

// TransferHelperTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TransferHelperTransactorRaw struct {
	Contract *TransferHelperTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTransferHelper creates a new instance of TransferHelper, bound to a specific deployed contract.
func NewTransferHelper(address common.Address, backend bind.ContractBackend) (*TransferHelper, error) {
	contract, err := bindTransferHelper(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TransferHelper{TransferHelperCaller: TransferHelperCaller{contract: contract}, TransferHelperTransactor: TransferHelperTransactor{contract: contract}, TransferHelperFilterer: TransferHelperFilterer{contract: contract}}, nil
}

// NewTransferHelperCaller creates a new read-only instance of TransferHelper, bound to a specific deployed contract.
func NewTransferHelperCaller(address common.Address, caller bind.ContractCaller) (*TransferHelperCaller, error) {
	contract, err := bindTransferHelper(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TransferHelperCaller{contract: contract}, nil
}

// NewTransferHelperTransactor creates a new write-only instance of TransferHelper, bound to a specific deployed contract.
func NewTransferHelperTransactor(address common.Address, transactor bind.ContractTransactor) (*TransferHelperTransactor, error) {
	contract, err := bindTransferHelper(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TransferHelperTransactor{contract: contract}, nil
}

// NewTransferHelperFilterer creates a new log filterer instance of TransferHelper, bound to a specific deployed contract.
func NewTransferHelperFilterer(address common.Address, filterer bind.ContractFilterer) (*TransferHelperFilterer, error) {
	contract, err := bindTransferHelper(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TransferHelperFilterer{contract: contract}, nil
}

// bindTransferHelper binds a generic wrapper to an already deployed contract.
func bindTransferHelper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TransferHelperABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TransferHelper *TransferHelperRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TransferHelper.Contract.TransferHelperCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TransferHelper *TransferHelperRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TransferHelper.Contract.TransferHelperTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TransferHelper *TransferHelperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TransferHelper.Contract.TransferHelperTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TransferHelper *TransferHelperCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TransferHelper.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TransferHelper *TransferHelperTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TransferHelper.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TransferHelper *TransferHelperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TransferHelper.Contract.contract.Transact(opts, method, params...)
}

// VenusControllerMetaData contains all meta data concerning the VenusController contract.
var VenusControllerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"getAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_mAddr\",\"type\":\"address\"}],\"name\":\"isManager\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"d502db97": "getAddr(string)",
		"f3ae2415": "isManager(address)",
	},
}

// VenusControllerABI is the input ABI used to generate the binding from.
// Deprecated: Use VenusControllerMetaData.ABI instead.
var VenusControllerABI = VenusControllerMetaData.ABI

// Deprecated: Use VenusControllerMetaData.Sigs instead.
// VenusControllerFuncSigs maps the 4-byte function signature to its string representation.
var VenusControllerFuncSigs = VenusControllerMetaData.Sigs

// VenusController is an auto generated Go binding around an Ethereum contract.
type VenusController struct {
	VenusControllerCaller     // Read-only binding to the contract
	VenusControllerTransactor // Write-only binding to the contract
	VenusControllerFilterer   // Log filterer for contract events
}

// VenusControllerCaller is an auto generated read-only Go binding around an Ethereum contract.
type VenusControllerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VenusControllerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VenusControllerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VenusControllerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VenusControllerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VenusControllerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VenusControllerSession struct {
	Contract     *VenusController  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VenusControllerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VenusControllerCallerSession struct {
	Contract *VenusControllerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// VenusControllerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VenusControllerTransactorSession struct {
	Contract     *VenusControllerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// VenusControllerRaw is an auto generated low-level Go binding around an Ethereum contract.
type VenusControllerRaw struct {
	Contract *VenusController // Generic contract binding to access the raw methods on
}

// VenusControllerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VenusControllerCallerRaw struct {
	Contract *VenusControllerCaller // Generic read-only contract binding to access the raw methods on
}

// VenusControllerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VenusControllerTransactorRaw struct {
	Contract *VenusControllerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVenusController creates a new instance of VenusController, bound to a specific deployed contract.
func NewVenusController(address common.Address, backend bind.ContractBackend) (*VenusController, error) {
	contract, err := bindVenusController(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VenusController{VenusControllerCaller: VenusControllerCaller{contract: contract}, VenusControllerTransactor: VenusControllerTransactor{contract: contract}, VenusControllerFilterer: VenusControllerFilterer{contract: contract}}, nil
}

// NewVenusControllerCaller creates a new read-only instance of VenusController, bound to a specific deployed contract.
func NewVenusControllerCaller(address common.Address, caller bind.ContractCaller) (*VenusControllerCaller, error) {
	contract, err := bindVenusController(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VenusControllerCaller{contract: contract}, nil
}

// NewVenusControllerTransactor creates a new write-only instance of VenusController, bound to a specific deployed contract.
func NewVenusControllerTransactor(address common.Address, transactor bind.ContractTransactor) (*VenusControllerTransactor, error) {
	contract, err := bindVenusController(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VenusControllerTransactor{contract: contract}, nil
}

// NewVenusControllerFilterer creates a new log filterer instance of VenusController, bound to a specific deployed contract.
func NewVenusControllerFilterer(address common.Address, filterer bind.ContractFilterer) (*VenusControllerFilterer, error) {
	contract, err := bindVenusController(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VenusControllerFilterer{contract: contract}, nil
}

// bindVenusController binds a generic wrapper to an already deployed contract.
func bindVenusController(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VenusControllerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VenusController *VenusControllerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VenusController.Contract.VenusControllerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VenusController *VenusControllerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VenusController.Contract.VenusControllerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VenusController *VenusControllerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VenusController.Contract.VenusControllerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VenusController *VenusControllerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VenusController.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VenusController *VenusControllerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VenusController.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VenusController *VenusControllerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VenusController.Contract.contract.Transact(opts, method, params...)
}

// GetAddr is a free data retrieval call binding the contract method 0xd502db97.
//
// Solidity: function getAddr(string _name) view returns(address)
func (_VenusController *VenusControllerCaller) GetAddr(opts *bind.CallOpts, _name string) (common.Address, error) {
	var out []interface{}
	err := _VenusController.contract.Call(opts, &out, "getAddr", _name)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddr is a free data retrieval call binding the contract method 0xd502db97.
//
// Solidity: function getAddr(string _name) view returns(address)
func (_VenusController *VenusControllerSession) GetAddr(_name string) (common.Address, error) {
	return _VenusController.Contract.GetAddr(&_VenusController.CallOpts, _name)
}

// GetAddr is a free data retrieval call binding the contract method 0xd502db97.
//
// Solidity: function getAddr(string _name) view returns(address)
func (_VenusController *VenusControllerCallerSession) GetAddr(_name string) (common.Address, error) {
	return _VenusController.Contract.GetAddr(&_VenusController.CallOpts, _name)
}

// IsManager is a free data retrieval call binding the contract method 0xf3ae2415.
//
// Solidity: function isManager(address _mAddr) view returns(bool)
func (_VenusController *VenusControllerCaller) IsManager(opts *bind.CallOpts, _mAddr common.Address) (bool, error) {
	var out []interface{}
	err := _VenusController.contract.Call(opts, &out, "isManager", _mAddr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsManager is a free data retrieval call binding the contract method 0xf3ae2415.
//
// Solidity: function isManager(address _mAddr) view returns(bool)
func (_VenusController *VenusControllerSession) IsManager(_mAddr common.Address) (bool, error) {
	return _VenusController.Contract.IsManager(&_VenusController.CallOpts, _mAddr)
}

// IsManager is a free data retrieval call binding the contract method 0xf3ae2415.
//
// Solidity: function isManager(address _mAddr) view returns(bool)
func (_VenusController *VenusControllerCallerSession) IsManager(_mAddr common.Address) (bool, error) {
	return _VenusController.Contract.IsManager(&_VenusController.CallOpts, _mAddr)
}

// VenusMarketMetaData contains all meta data concerning the VenusMarket contract.
var VenusMarketMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractVenusController\",\"name\":\"_addrc\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_oid\",\"type\":\"uint256\"}],\"name\":\"CancelOrder\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_oid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_receiveToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_cycle\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_endTime\",\"type\":\"uint256\"}],\"name\":\"CreateOrder\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"_canSell\",\"type\":\"bool\"}],\"name\":\"SetBlackUser\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"addrc\",\"outputs\":[{\"internalType\":\"contractVenusController\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"batchSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"blackUser\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_oid\",\"type\":\"uint256\"}],\"name\":\"cancelOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_receiveToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_cycle\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_endTime\",\"type\":\"uint256\"}],\"name\":\"createOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minCycle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"nameAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"orderID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"orderIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"orders\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"reciveToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"stepTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cycle\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_batch\",\"type\":\"uint256\"}],\"name\":\"setBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_canUse\",\"type\":\"bool\"}],\"name\":\"setBlackUser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_feeRate\",\"type\":\"uint256\"}],\"name\":\"setFeeRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"succIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_start\",\"type\":\"uint256\"}],\"name\":\"trigTask\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userToOrder\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"913fcfe0": "addrc()",
		"f4daaba1": "batchSize()",
		"06926173": "blackUser(address)",
		"514fcac7": "cancelOrder(uint256)",
		"328efc47": "createOrder(uint256,address,uint256,uint256)",
		"978bbdb9": "feeRate()",
		"a9aa92e8": "minCycle()",
		"a310d29c": "nameAddr(string)",
		"81203d03": "orderID()",
		"864d1260": "orderIndex(uint256)",
		"a85c38ef": "orders(uint256)",
		"b76060f7": "setBatch(uint256)",
		"79e3001a": "setBlackUser(address,bool)",
		"45596e2e": "setFeeRate(uint256)",
		"9c23547d": "succIndex()",
		"d5e7ea62": "trigTask(uint256)",
		"6f53f770": "userToOrder(address)",
	},
	Bin: "0x608060405261271060025562015180600755606460085534801561002257600080fd5b50604051610f31380380610f318339818101604052602081101561004557600080fd5b5051600080546001600160a01b039092166001600160a01b0319909216919091179055610eba806100776000396000f3fe608060405234801561001057600080fd5b506004361061010b5760003560e01c8063913fcfe0116100a2578063a85c38ef11610071578063a85c38ef14610323578063a9aa92e81461039b578063b76060f7146103a3578063d5e7ea62146103c0578063f4daaba1146103dd5761010b565b8063913fcfe014610249578063978bbdb91461026d5780639c23547d14610275578063a310d29c1461027d5761010b565b80636f53f770116100de5780636f53f770146101be57806379e3001a146101f657806381203d0314610224578063864d12601461022c5761010b565b80630692617314610110578063328efc471461014a57806345596e2e14610184578063514fcac7146101a1575b600080fd5b6101366004803603602081101561012657600080fd5b50356001600160a01b03166103e5565b604080519115158252519081900360200190f35b6101826004803603608081101561016057600080fd5b508035906001600160a01b0360208201351690604081013590606001356103fa565b005b6101826004803603602081101561019a57600080fd5b5035610700565b610182600480360360208110156101b757600080fd5b50356107bb565b6101e4600480360360208110156101d457600080fd5b50356001600160a01b0316610943565b60408051918252519081900360200190f35b6101826004803603604081101561020c57600080fd5b506001600160a01b0381351690602001351515610955565b6101e4610a6f565b6101e46004803603602081101561024257600080fd5b5035610a75565b610251610a93565b604080516001600160a01b039092168252519081900360200190f35b6101e4610aa2565b6101e4610aa8565b6102516004803603602081101561029357600080fd5b8101906020810181356401000000008111156102ae57600080fd5b8201836020820111156102c057600080fd5b803590602001918460018302840111640100000000831117156102e257600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610aae945050505050565b6103406004803603602081101561033957600080fd5b5035610b8a565b60405180896001600160a01b03168152602001888152602001878152602001866001600160a01b0316815260200185815260200184815260200183815260200182151581526020019850505050505050505060405180910390f35b6101e4610bdf565b610182600480360360208110156103b957600080fd5b5035610be5565b610182600480360360208110156103d657600080fd5b5035610ca0565b6101e4610e42565b60066020526000908152604090205460ff1681565b3360009081526006602052604090205460ff1661045e576040805162461bcd60e51b815260206004820152601b60248201527f54686973206163636f756e74206973206e6f7420616c6c6f7765640000000000604482015290519081900360640190fd5b336000908152600560205260409020546104b8576040805162461bcd60e51b81526020600482015260166024820152753ab9b2b91030b63932b0b23c903430b99037b93232b960511b604482015290519081900360640190fd5b6104c183610e48565b61050b576040805162461bcd60e51b81526020600482015260166024820152756e6f74206120636f6e7472616374206164647265737360501b604482015290519081900360640190fd5b600260008154809291906001019190505550604051806101000160405280336001600160a01b031681526020016004805490508152602001858152602001846001600160a01b031681526020014281526020018281526020018381526020016001151581525060036000600254815260200190815260200160002060008201518160000160006101000a8154816001600160a01b0302191690836001600160a01b03160217905550602082015181600101556040820151816002015560608201518160030160006101000a8154816001600160a01b0302191690836001600160a01b031602179055506080820151816004015560a0820151816005015560c0820151816006015560e08201518160070160006101000a81548160ff0219169083151502179055509050506004600254908060018154018082558091505060019003906000526020600020016000909190919091505560025460056000336001600160a01b03166001600160a01b031681526020019081526020016000208190555060075482101561069c5760078290555b600254604080519182523360208301528181018690526001600160a01b03851660608301526080820184905260a08201839052517f23df57ff35565965f751a83073b22102fcf2bf4fb084463c0c5b68ddc9a5b1289181900360c00190a150505050565b6000546040805163f3ae241560e01b815233600482015290516001600160a01b039092169163f3ae241591602480820192602092909190829003018186803b15801561074b57600080fd5b505afa15801561075f573d6000803e3d6000fd5b505050506040513d602081101561077557600080fd5b50516107b6576040805162461bcd60e51b815260206004820152600b60248201526a37b7363ca6b0b730b3b2b960a91b604482015290519081900360640190fd5b600155565b3360009081526006602052604090205460ff1661081f576040805162461bcd60e51b815260206004820152601b60248201527f54686973206163636f756e74206973206e6f7420616c6c6f7765640000000000604482015290519081900360640190fd5b60008181526003602052604090206007015460ff1661087a576040805162461bcd60e51b81526020600482015260126024820152711bdc99195c881a5cc81b9bdd08195e1a5cdd60721b604482015290519081900360640190fd5b6000818152600360205260409020546001600160a01b031633146108e5576040805162461bcd60e51b815260206004820152601960248201527f6f72646572206f776e657220646f65736e2774206d6174636800000000000000604482015290519081900360640190fd5b6000818152600360209081526040808320600701805460ff191690553383526005825280832092909255815183815291517f935c9ad2f1fda9d7eae0d2a512f1521cb7190ee06165414e722366a65975fb6b9281900390910190a150565b60056020526000908152604090205481565b6000546040805163f3ae241560e01b815233600482015290516001600160a01b039092169163f3ae241591602480820192602092909190829003018186803b1580156109a057600080fd5b505afa1580156109b4573d6000803e3d6000fd5b505050506040513d60208110156109ca57600080fd5b5051610a0b576040805162461bcd60e51b815260206004820152600b60248201526a37b7363ca6b0b730b3b2b960a91b604482015290519081900360640190fd5b6001600160a01b038216600081815260066020908152604091829020805460ff191685151590811790915582519384529083015280517f268e635d8b3370e1ab4560005f02806d3efb314b2aca5cea4a68f4f44eb3d9839281900390910190a15050565b60025481565b60048181548110610a8257fe5b600091825260209091200154905081565b6000546001600160a01b031681565b60015481565b60095481565b6000805460405163d502db9760e01b81526020600482018181528551602484015285516001600160a01b039094169363d502db979387938392604490920191908501908083838b5b83811015610b0e578181015183820152602001610af6565b50505050905090810190601f168015610b3b5780820380516001836020036101000a031916815260200191505b509250505060206040518083038186803b158015610b5857600080fd5b505afa158015610b6c573d6000803e3d6000fd5b505050506040513d6020811015610b8257600080fd5b505192915050565b6003602081905260009182526040909120805460018201546002830154938301546004840154600585015460068601546007909601546001600160a01b03958616979496949590931693919290919060ff1688565b60075481565b6000546040805163f3ae241560e01b815233600482015290516001600160a01b039092169163f3ae241591602480820192602092909190829003018186803b158015610c3057600080fd5b505afa158015610c44573d6000803e3d6000fd5b505050506040513d6020811015610c5a57600080fd5b5051610c9b576040805162461bcd60e51b815260206004820152600b60248201526a37b7363ca6b0b730b3b2b960a91b604482015290519081900360640190fd5b600855565b6000546040805163f3ae241560e01b815233600482015290516001600160a01b039092169163f3ae241591602480820192602092909190829003018186803b158015610ceb57600080fd5b505afa158015610cff573d6000803e3d6000fd5b505050506040513d6020811015610d1557600080fd5b5051610d56576040805162461bcd60e51b815260206004820152600b60248201526a37b7363ca6b0b730b3b2b960a91b604482015290519081900360640190fd5b6004548110610d6457610e3f565b60085460045490820190811115610d7a57506004545b815b81811015610e3c578060098190555060006003600060048481548110610d9e57fe5b906000526020600020015481526020019081526020016000209050806005015481600601548260040154011180610dd85750428160040154115b15610ded57600701805460ff19169055610e34565b42816006015482600401540110610e3257805460038201546002830154610e21926001600160a01b03908116921690610e3c565b600681015460048201805490910190555b505b600101610d7c565b50505b50565b60085481565b6000813f7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a4708115801590610e7c5750808214155b94935050505056fea26469706673582212201cd69e5fd8f491d60000ee06eebfac47edd84a47e5ac6851b7baa4ecb91422f764736f6c634300060c0033",
}

// VenusMarketABI is the input ABI used to generate the binding from.
// Deprecated: Use VenusMarketMetaData.ABI instead.
var VenusMarketABI = VenusMarketMetaData.ABI

// Deprecated: Use VenusMarketMetaData.Sigs instead.
// VenusMarketFuncSigs maps the 4-byte function signature to its string representation.
var VenusMarketFuncSigs = VenusMarketMetaData.Sigs

// VenusMarketBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use VenusMarketMetaData.Bin instead.
var VenusMarketBin = VenusMarketMetaData.Bin

// DeployVenusMarket deploys a new Ethereum contract, binding an instance of VenusMarket to it.
func DeployVenusMarket(auth *bind.TransactOpts, backend bind.ContractBackend, _addrc common.Address) (common.Address, *types.Transaction, *VenusMarket, error) {
	parsed, err := VenusMarketMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VenusMarketBin), backend, _addrc)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &VenusMarket{VenusMarketCaller: VenusMarketCaller{contract: contract}, VenusMarketTransactor: VenusMarketTransactor{contract: contract}, VenusMarketFilterer: VenusMarketFilterer{contract: contract}}, nil
}

// VenusMarket is an auto generated Go binding around an Ethereum contract.
type VenusMarket struct {
	VenusMarketCaller     // Read-only binding to the contract
	VenusMarketTransactor // Write-only binding to the contract
	VenusMarketFilterer   // Log filterer for contract events
}

// VenusMarketCaller is an auto generated read-only Go binding around an Ethereum contract.
type VenusMarketCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VenusMarketTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VenusMarketTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VenusMarketFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VenusMarketFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VenusMarketSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VenusMarketSession struct {
	Contract     *VenusMarket      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VenusMarketCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VenusMarketCallerSession struct {
	Contract *VenusMarketCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// VenusMarketTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VenusMarketTransactorSession struct {
	Contract     *VenusMarketTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// VenusMarketRaw is an auto generated low-level Go binding around an Ethereum contract.
type VenusMarketRaw struct {
	Contract *VenusMarket // Generic contract binding to access the raw methods on
}

// VenusMarketCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VenusMarketCallerRaw struct {
	Contract *VenusMarketCaller // Generic read-only contract binding to access the raw methods on
}

// VenusMarketTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VenusMarketTransactorRaw struct {
	Contract *VenusMarketTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVenusMarket creates a new instance of VenusMarket, bound to a specific deployed contract.
func NewVenusMarket(address common.Address, backend bind.ContractBackend) (*VenusMarket, error) {
	contract, err := bindVenusMarket(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VenusMarket{VenusMarketCaller: VenusMarketCaller{contract: contract}, VenusMarketTransactor: VenusMarketTransactor{contract: contract}, VenusMarketFilterer: VenusMarketFilterer{contract: contract}}, nil
}

// NewVenusMarketCaller creates a new read-only instance of VenusMarket, bound to a specific deployed contract.
func NewVenusMarketCaller(address common.Address, caller bind.ContractCaller) (*VenusMarketCaller, error) {
	contract, err := bindVenusMarket(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VenusMarketCaller{contract: contract}, nil
}

// NewVenusMarketTransactor creates a new write-only instance of VenusMarket, bound to a specific deployed contract.
func NewVenusMarketTransactor(address common.Address, transactor bind.ContractTransactor) (*VenusMarketTransactor, error) {
	contract, err := bindVenusMarket(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VenusMarketTransactor{contract: contract}, nil
}

// NewVenusMarketFilterer creates a new log filterer instance of VenusMarket, bound to a specific deployed contract.
func NewVenusMarketFilterer(address common.Address, filterer bind.ContractFilterer) (*VenusMarketFilterer, error) {
	contract, err := bindVenusMarket(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VenusMarketFilterer{contract: contract}, nil
}

// bindVenusMarket binds a generic wrapper to an already deployed contract.
func bindVenusMarket(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VenusMarketABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VenusMarket *VenusMarketRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VenusMarket.Contract.VenusMarketCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VenusMarket *VenusMarketRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VenusMarket.Contract.VenusMarketTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VenusMarket *VenusMarketRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VenusMarket.Contract.VenusMarketTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VenusMarket *VenusMarketCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VenusMarket.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VenusMarket *VenusMarketTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VenusMarket.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VenusMarket *VenusMarketTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VenusMarket.Contract.contract.Transact(opts, method, params...)
}

// Addrc is a free data retrieval call binding the contract method 0x913fcfe0.
//
// Solidity: function addrc() view returns(address)
func (_VenusMarket *VenusMarketCaller) Addrc(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VenusMarket.contract.Call(opts, &out, "addrc")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Addrc is a free data retrieval call binding the contract method 0x913fcfe0.
//
// Solidity: function addrc() view returns(address)
func (_VenusMarket *VenusMarketSession) Addrc() (common.Address, error) {
	return _VenusMarket.Contract.Addrc(&_VenusMarket.CallOpts)
}

// Addrc is a free data retrieval call binding the contract method 0x913fcfe0.
//
// Solidity: function addrc() view returns(address)
func (_VenusMarket *VenusMarketCallerSession) Addrc() (common.Address, error) {
	return _VenusMarket.Contract.Addrc(&_VenusMarket.CallOpts)
}

// BatchSize is a free data retrieval call binding the contract method 0xf4daaba1.
//
// Solidity: function batchSize() view returns(uint256)
func (_VenusMarket *VenusMarketCaller) BatchSize(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VenusMarket.contract.Call(opts, &out, "batchSize")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BatchSize is a free data retrieval call binding the contract method 0xf4daaba1.
//
// Solidity: function batchSize() view returns(uint256)
func (_VenusMarket *VenusMarketSession) BatchSize() (*big.Int, error) {
	return _VenusMarket.Contract.BatchSize(&_VenusMarket.CallOpts)
}

// BatchSize is a free data retrieval call binding the contract method 0xf4daaba1.
//
// Solidity: function batchSize() view returns(uint256)
func (_VenusMarket *VenusMarketCallerSession) BatchSize() (*big.Int, error) {
	return _VenusMarket.Contract.BatchSize(&_VenusMarket.CallOpts)
}

// BlackUser is a free data retrieval call binding the contract method 0x06926173.
//
// Solidity: function blackUser(address ) view returns(bool)
func (_VenusMarket *VenusMarketCaller) BlackUser(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _VenusMarket.contract.Call(opts, &out, "blackUser", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BlackUser is a free data retrieval call binding the contract method 0x06926173.
//
// Solidity: function blackUser(address ) view returns(bool)
func (_VenusMarket *VenusMarketSession) BlackUser(arg0 common.Address) (bool, error) {
	return _VenusMarket.Contract.BlackUser(&_VenusMarket.CallOpts, arg0)
}

// BlackUser is a free data retrieval call binding the contract method 0x06926173.
//
// Solidity: function blackUser(address ) view returns(bool)
func (_VenusMarket *VenusMarketCallerSession) BlackUser(arg0 common.Address) (bool, error) {
	return _VenusMarket.Contract.BlackUser(&_VenusMarket.CallOpts, arg0)
}

// FeeRate is a free data retrieval call binding the contract method 0x978bbdb9.
//
// Solidity: function feeRate() view returns(uint256)
func (_VenusMarket *VenusMarketCaller) FeeRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VenusMarket.contract.Call(opts, &out, "feeRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeRate is a free data retrieval call binding the contract method 0x978bbdb9.
//
// Solidity: function feeRate() view returns(uint256)
func (_VenusMarket *VenusMarketSession) FeeRate() (*big.Int, error) {
	return _VenusMarket.Contract.FeeRate(&_VenusMarket.CallOpts)
}

// FeeRate is a free data retrieval call binding the contract method 0x978bbdb9.
//
// Solidity: function feeRate() view returns(uint256)
func (_VenusMarket *VenusMarketCallerSession) FeeRate() (*big.Int, error) {
	return _VenusMarket.Contract.FeeRate(&_VenusMarket.CallOpts)
}

// MinCycle is a free data retrieval call binding the contract method 0xa9aa92e8.
//
// Solidity: function minCycle() view returns(uint256)
func (_VenusMarket *VenusMarketCaller) MinCycle(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VenusMarket.contract.Call(opts, &out, "minCycle")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinCycle is a free data retrieval call binding the contract method 0xa9aa92e8.
//
// Solidity: function minCycle() view returns(uint256)
func (_VenusMarket *VenusMarketSession) MinCycle() (*big.Int, error) {
	return _VenusMarket.Contract.MinCycle(&_VenusMarket.CallOpts)
}

// MinCycle is a free data retrieval call binding the contract method 0xa9aa92e8.
//
// Solidity: function minCycle() view returns(uint256)
func (_VenusMarket *VenusMarketCallerSession) MinCycle() (*big.Int, error) {
	return _VenusMarket.Contract.MinCycle(&_VenusMarket.CallOpts)
}

// NameAddr is a free data retrieval call binding the contract method 0xa310d29c.
//
// Solidity: function nameAddr(string _name) view returns(address)
func (_VenusMarket *VenusMarketCaller) NameAddr(opts *bind.CallOpts, _name string) (common.Address, error) {
	var out []interface{}
	err := _VenusMarket.contract.Call(opts, &out, "nameAddr", _name)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NameAddr is a free data retrieval call binding the contract method 0xa310d29c.
//
// Solidity: function nameAddr(string _name) view returns(address)
func (_VenusMarket *VenusMarketSession) NameAddr(_name string) (common.Address, error) {
	return _VenusMarket.Contract.NameAddr(&_VenusMarket.CallOpts, _name)
}

// NameAddr is a free data retrieval call binding the contract method 0xa310d29c.
//
// Solidity: function nameAddr(string _name) view returns(address)
func (_VenusMarket *VenusMarketCallerSession) NameAddr(_name string) (common.Address, error) {
	return _VenusMarket.Contract.NameAddr(&_VenusMarket.CallOpts, _name)
}

// OrderID is a free data retrieval call binding the contract method 0x81203d03.
//
// Solidity: function orderID() view returns(uint256)
func (_VenusMarket *VenusMarketCaller) OrderID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VenusMarket.contract.Call(opts, &out, "orderID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OrderID is a free data retrieval call binding the contract method 0x81203d03.
//
// Solidity: function orderID() view returns(uint256)
func (_VenusMarket *VenusMarketSession) OrderID() (*big.Int, error) {
	return _VenusMarket.Contract.OrderID(&_VenusMarket.CallOpts)
}

// OrderID is a free data retrieval call binding the contract method 0x81203d03.
//
// Solidity: function orderID() view returns(uint256)
func (_VenusMarket *VenusMarketCallerSession) OrderID() (*big.Int, error) {
	return _VenusMarket.Contract.OrderID(&_VenusMarket.CallOpts)
}

// OrderIndex is a free data retrieval call binding the contract method 0x864d1260.
//
// Solidity: function orderIndex(uint256 ) view returns(uint256)
func (_VenusMarket *VenusMarketCaller) OrderIndex(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _VenusMarket.contract.Call(opts, &out, "orderIndex", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OrderIndex is a free data retrieval call binding the contract method 0x864d1260.
//
// Solidity: function orderIndex(uint256 ) view returns(uint256)
func (_VenusMarket *VenusMarketSession) OrderIndex(arg0 *big.Int) (*big.Int, error) {
	return _VenusMarket.Contract.OrderIndex(&_VenusMarket.CallOpts, arg0)
}

// OrderIndex is a free data retrieval call binding the contract method 0x864d1260.
//
// Solidity: function orderIndex(uint256 ) view returns(uint256)
func (_VenusMarket *VenusMarketCallerSession) OrderIndex(arg0 *big.Int) (*big.Int, error) {
	return _VenusMarket.Contract.OrderIndex(&_VenusMarket.CallOpts, arg0)
}

// Orders is a free data retrieval call binding the contract method 0xa85c38ef.
//
// Solidity: function orders(uint256 ) view returns(address owner, uint256 index, uint256 price, address reciveToken, uint256 stepTime, uint256 endTime, uint256 cycle, bool status)
func (_VenusMarket *VenusMarketCaller) Orders(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Owner       common.Address
	Index       *big.Int
	Price       *big.Int
	ReciveToken common.Address
	StepTime    *big.Int
	EndTime     *big.Int
	Cycle       *big.Int
	Status      bool
}, error) {
	var out []interface{}
	err := _VenusMarket.contract.Call(opts, &out, "orders", arg0)

	outstruct := new(struct {
		Owner       common.Address
		Index       *big.Int
		Price       *big.Int
		ReciveToken common.Address
		StepTime    *big.Int
		EndTime     *big.Int
		Cycle       *big.Int
		Status      bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Owner = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Index = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Price = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.ReciveToken = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.StepTime = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.EndTime = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.Cycle = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.Status = *abi.ConvertType(out[7], new(bool)).(*bool)

	return *outstruct, err

}

// Orders is a free data retrieval call binding the contract method 0xa85c38ef.
//
// Solidity: function orders(uint256 ) view returns(address owner, uint256 index, uint256 price, address reciveToken, uint256 stepTime, uint256 endTime, uint256 cycle, bool status)
func (_VenusMarket *VenusMarketSession) Orders(arg0 *big.Int) (struct {
	Owner       common.Address
	Index       *big.Int
	Price       *big.Int
	ReciveToken common.Address
	StepTime    *big.Int
	EndTime     *big.Int
	Cycle       *big.Int
	Status      bool
}, error) {
	return _VenusMarket.Contract.Orders(&_VenusMarket.CallOpts, arg0)
}

// Orders is a free data retrieval call binding the contract method 0xa85c38ef.
//
// Solidity: function orders(uint256 ) view returns(address owner, uint256 index, uint256 price, address reciveToken, uint256 stepTime, uint256 endTime, uint256 cycle, bool status)
func (_VenusMarket *VenusMarketCallerSession) Orders(arg0 *big.Int) (struct {
	Owner       common.Address
	Index       *big.Int
	Price       *big.Int
	ReciveToken common.Address
	StepTime    *big.Int
	EndTime     *big.Int
	Cycle       *big.Int
	Status      bool
}, error) {
	return _VenusMarket.Contract.Orders(&_VenusMarket.CallOpts, arg0)
}

// SuccIndex is a free data retrieval call binding the contract method 0x9c23547d.
//
// Solidity: function succIndex() view returns(uint256)
func (_VenusMarket *VenusMarketCaller) SuccIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VenusMarket.contract.Call(opts, &out, "succIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SuccIndex is a free data retrieval call binding the contract method 0x9c23547d.
//
// Solidity: function succIndex() view returns(uint256)
func (_VenusMarket *VenusMarketSession) SuccIndex() (*big.Int, error) {
	return _VenusMarket.Contract.SuccIndex(&_VenusMarket.CallOpts)
}

// SuccIndex is a free data retrieval call binding the contract method 0x9c23547d.
//
// Solidity: function succIndex() view returns(uint256)
func (_VenusMarket *VenusMarketCallerSession) SuccIndex() (*big.Int, error) {
	return _VenusMarket.Contract.SuccIndex(&_VenusMarket.CallOpts)
}

// UserToOrder is a free data retrieval call binding the contract method 0x6f53f770.
//
// Solidity: function userToOrder(address ) view returns(uint256)
func (_VenusMarket *VenusMarketCaller) UserToOrder(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VenusMarket.contract.Call(opts, &out, "userToOrder", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserToOrder is a free data retrieval call binding the contract method 0x6f53f770.
//
// Solidity: function userToOrder(address ) view returns(uint256)
func (_VenusMarket *VenusMarketSession) UserToOrder(arg0 common.Address) (*big.Int, error) {
	return _VenusMarket.Contract.UserToOrder(&_VenusMarket.CallOpts, arg0)
}

// UserToOrder is a free data retrieval call binding the contract method 0x6f53f770.
//
// Solidity: function userToOrder(address ) view returns(uint256)
func (_VenusMarket *VenusMarketCallerSession) UserToOrder(arg0 common.Address) (*big.Int, error) {
	return _VenusMarket.Contract.UserToOrder(&_VenusMarket.CallOpts, arg0)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x514fcac7.
//
// Solidity: function cancelOrder(uint256 _oid) returns()
func (_VenusMarket *VenusMarketTransactor) CancelOrder(opts *bind.TransactOpts, _oid *big.Int) (*types.Transaction, error) {
	return _VenusMarket.contract.Transact(opts, "cancelOrder", _oid)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x514fcac7.
//
// Solidity: function cancelOrder(uint256 _oid) returns()
func (_VenusMarket *VenusMarketSession) CancelOrder(_oid *big.Int) (*types.Transaction, error) {
	return _VenusMarket.Contract.CancelOrder(&_VenusMarket.TransactOpts, _oid)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x514fcac7.
//
// Solidity: function cancelOrder(uint256 _oid) returns()
func (_VenusMarket *VenusMarketTransactorSession) CancelOrder(_oid *big.Int) (*types.Transaction, error) {
	return _VenusMarket.Contract.CancelOrder(&_VenusMarket.TransactOpts, _oid)
}

// CreateOrder is a paid mutator transaction binding the contract method 0x328efc47.
//
// Solidity: function createOrder(uint256 _price, address _receiveToken, uint256 _cycle, uint256 _endTime) returns()
func (_VenusMarket *VenusMarketTransactor) CreateOrder(opts *bind.TransactOpts, _price *big.Int, _receiveToken common.Address, _cycle *big.Int, _endTime *big.Int) (*types.Transaction, error) {
	return _VenusMarket.contract.Transact(opts, "createOrder", _price, _receiveToken, _cycle, _endTime)
}

// CreateOrder is a paid mutator transaction binding the contract method 0x328efc47.
//
// Solidity: function createOrder(uint256 _price, address _receiveToken, uint256 _cycle, uint256 _endTime) returns()
func (_VenusMarket *VenusMarketSession) CreateOrder(_price *big.Int, _receiveToken common.Address, _cycle *big.Int, _endTime *big.Int) (*types.Transaction, error) {
	return _VenusMarket.Contract.CreateOrder(&_VenusMarket.TransactOpts, _price, _receiveToken, _cycle, _endTime)
}

// CreateOrder is a paid mutator transaction binding the contract method 0x328efc47.
//
// Solidity: function createOrder(uint256 _price, address _receiveToken, uint256 _cycle, uint256 _endTime) returns()
func (_VenusMarket *VenusMarketTransactorSession) CreateOrder(_price *big.Int, _receiveToken common.Address, _cycle *big.Int, _endTime *big.Int) (*types.Transaction, error) {
	return _VenusMarket.Contract.CreateOrder(&_VenusMarket.TransactOpts, _price, _receiveToken, _cycle, _endTime)
}

// SetBatch is a paid mutator transaction binding the contract method 0xb76060f7.
//
// Solidity: function setBatch(uint256 _batch) returns()
func (_VenusMarket *VenusMarketTransactor) SetBatch(opts *bind.TransactOpts, _batch *big.Int) (*types.Transaction, error) {
	return _VenusMarket.contract.Transact(opts, "setBatch", _batch)
}

// SetBatch is a paid mutator transaction binding the contract method 0xb76060f7.
//
// Solidity: function setBatch(uint256 _batch) returns()
func (_VenusMarket *VenusMarketSession) SetBatch(_batch *big.Int) (*types.Transaction, error) {
	return _VenusMarket.Contract.SetBatch(&_VenusMarket.TransactOpts, _batch)
}

// SetBatch is a paid mutator transaction binding the contract method 0xb76060f7.
//
// Solidity: function setBatch(uint256 _batch) returns()
func (_VenusMarket *VenusMarketTransactorSession) SetBatch(_batch *big.Int) (*types.Transaction, error) {
	return _VenusMarket.Contract.SetBatch(&_VenusMarket.TransactOpts, _batch)
}

// SetBlackUser is a paid mutator transaction binding the contract method 0x79e3001a.
//
// Solidity: function setBlackUser(address _user, bool _canUse) returns()
func (_VenusMarket *VenusMarketTransactor) SetBlackUser(opts *bind.TransactOpts, _user common.Address, _canUse bool) (*types.Transaction, error) {
	return _VenusMarket.contract.Transact(opts, "setBlackUser", _user, _canUse)
}

// SetBlackUser is a paid mutator transaction binding the contract method 0x79e3001a.
//
// Solidity: function setBlackUser(address _user, bool _canUse) returns()
func (_VenusMarket *VenusMarketSession) SetBlackUser(_user common.Address, _canUse bool) (*types.Transaction, error) {
	return _VenusMarket.Contract.SetBlackUser(&_VenusMarket.TransactOpts, _user, _canUse)
}

// SetBlackUser is a paid mutator transaction binding the contract method 0x79e3001a.
//
// Solidity: function setBlackUser(address _user, bool _canUse) returns()
func (_VenusMarket *VenusMarketTransactorSession) SetBlackUser(_user common.Address, _canUse bool) (*types.Transaction, error) {
	return _VenusMarket.Contract.SetBlackUser(&_VenusMarket.TransactOpts, _user, _canUse)
}

// SetFeeRate is a paid mutator transaction binding the contract method 0x45596e2e.
//
// Solidity: function setFeeRate(uint256 _feeRate) returns()
func (_VenusMarket *VenusMarketTransactor) SetFeeRate(opts *bind.TransactOpts, _feeRate *big.Int) (*types.Transaction, error) {
	return _VenusMarket.contract.Transact(opts, "setFeeRate", _feeRate)
}

// SetFeeRate is a paid mutator transaction binding the contract method 0x45596e2e.
//
// Solidity: function setFeeRate(uint256 _feeRate) returns()
func (_VenusMarket *VenusMarketSession) SetFeeRate(_feeRate *big.Int) (*types.Transaction, error) {
	return _VenusMarket.Contract.SetFeeRate(&_VenusMarket.TransactOpts, _feeRate)
}

// SetFeeRate is a paid mutator transaction binding the contract method 0x45596e2e.
//
// Solidity: function setFeeRate(uint256 _feeRate) returns()
func (_VenusMarket *VenusMarketTransactorSession) SetFeeRate(_feeRate *big.Int) (*types.Transaction, error) {
	return _VenusMarket.Contract.SetFeeRate(&_VenusMarket.TransactOpts, _feeRate)
}

// TrigTask is a paid mutator transaction binding the contract method 0xd5e7ea62.
//
// Solidity: function trigTask(uint256 _start) returns()
func (_VenusMarket *VenusMarketTransactor) TrigTask(opts *bind.TransactOpts, _start *big.Int) (*types.Transaction, error) {
	return _VenusMarket.contract.Transact(opts, "trigTask", _start)
}

// TrigTask is a paid mutator transaction binding the contract method 0xd5e7ea62.
//
// Solidity: function trigTask(uint256 _start) returns()
func (_VenusMarket *VenusMarketSession) TrigTask(_start *big.Int) (*types.Transaction, error) {
	return _VenusMarket.Contract.TrigTask(&_VenusMarket.TransactOpts, _start)
}

// TrigTask is a paid mutator transaction binding the contract method 0xd5e7ea62.
//
// Solidity: function trigTask(uint256 _start) returns()
func (_VenusMarket *VenusMarketTransactorSession) TrigTask(_start *big.Int) (*types.Transaction, error) {
	return _VenusMarket.Contract.TrigTask(&_VenusMarket.TransactOpts, _start)
}

// VenusMarketCancelOrderIterator is returned from FilterCancelOrder and is used to iterate over the raw logs and unpacked data for CancelOrder events raised by the VenusMarket contract.
type VenusMarketCancelOrderIterator struct {
	Event *VenusMarketCancelOrder // Event containing the contract specifics and raw log

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
func (it *VenusMarketCancelOrderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VenusMarketCancelOrder)
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
		it.Event = new(VenusMarketCancelOrder)
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
func (it *VenusMarketCancelOrderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VenusMarketCancelOrderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VenusMarketCancelOrder represents a CancelOrder event raised by the VenusMarket contract.
type VenusMarketCancelOrder struct {
	Oid *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterCancelOrder is a free log retrieval operation binding the contract event 0x935c9ad2f1fda9d7eae0d2a512f1521cb7190ee06165414e722366a65975fb6b.
//
// Solidity: event CancelOrder(uint256 _oid)
func (_VenusMarket *VenusMarketFilterer) FilterCancelOrder(opts *bind.FilterOpts) (*VenusMarketCancelOrderIterator, error) {

	logs, sub, err := _VenusMarket.contract.FilterLogs(opts, "CancelOrder")
	if err != nil {
		return nil, err
	}
	return &VenusMarketCancelOrderIterator{contract: _VenusMarket.contract, event: "CancelOrder", logs: logs, sub: sub}, nil
}

// WatchCancelOrder is a free log subscription operation binding the contract event 0x935c9ad2f1fda9d7eae0d2a512f1521cb7190ee06165414e722366a65975fb6b.
//
// Solidity: event CancelOrder(uint256 _oid)
func (_VenusMarket *VenusMarketFilterer) WatchCancelOrder(opts *bind.WatchOpts, sink chan<- *VenusMarketCancelOrder) (event.Subscription, error) {

	logs, sub, err := _VenusMarket.contract.WatchLogs(opts, "CancelOrder")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VenusMarketCancelOrder)
				if err := _VenusMarket.contract.UnpackLog(event, "CancelOrder", log); err != nil {
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

// ParseCancelOrder is a log parse operation binding the contract event 0x935c9ad2f1fda9d7eae0d2a512f1521cb7190ee06165414e722366a65975fb6b.
//
// Solidity: event CancelOrder(uint256 _oid)
func (_VenusMarket *VenusMarketFilterer) ParseCancelOrder(log types.Log) (*VenusMarketCancelOrder, error) {
	event := new(VenusMarketCancelOrder)
	if err := _VenusMarket.contract.UnpackLog(event, "CancelOrder", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VenusMarketCreateOrderIterator is returned from FilterCreateOrder and is used to iterate over the raw logs and unpacked data for CreateOrder events raised by the VenusMarket contract.
type VenusMarketCreateOrderIterator struct {
	Event *VenusMarketCreateOrder // Event containing the contract specifics and raw log

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
func (it *VenusMarketCreateOrderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VenusMarketCreateOrder)
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
		it.Event = new(VenusMarketCreateOrder)
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
func (it *VenusMarketCreateOrderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VenusMarketCreateOrderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VenusMarketCreateOrder represents a CreateOrder event raised by the VenusMarket contract.
type VenusMarketCreateOrder struct {
	Oid          *big.Int
	User         common.Address
	Price        *big.Int
	ReceiveToken common.Address
	Cycle        *big.Int
	EndTime      *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterCreateOrder is a free log retrieval operation binding the contract event 0x23df57ff35565965f751a83073b22102fcf2bf4fb084463c0c5b68ddc9a5b128.
//
// Solidity: event CreateOrder(uint256 _oid, address user, uint256 _price, address _receiveToken, uint256 _cycle, uint256 _endTime)
func (_VenusMarket *VenusMarketFilterer) FilterCreateOrder(opts *bind.FilterOpts) (*VenusMarketCreateOrderIterator, error) {

	logs, sub, err := _VenusMarket.contract.FilterLogs(opts, "CreateOrder")
	if err != nil {
		return nil, err
	}
	return &VenusMarketCreateOrderIterator{contract: _VenusMarket.contract, event: "CreateOrder", logs: logs, sub: sub}, nil
}

// WatchCreateOrder is a free log subscription operation binding the contract event 0x23df57ff35565965f751a83073b22102fcf2bf4fb084463c0c5b68ddc9a5b128.
//
// Solidity: event CreateOrder(uint256 _oid, address user, uint256 _price, address _receiveToken, uint256 _cycle, uint256 _endTime)
func (_VenusMarket *VenusMarketFilterer) WatchCreateOrder(opts *bind.WatchOpts, sink chan<- *VenusMarketCreateOrder) (event.Subscription, error) {

	logs, sub, err := _VenusMarket.contract.WatchLogs(opts, "CreateOrder")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VenusMarketCreateOrder)
				if err := _VenusMarket.contract.UnpackLog(event, "CreateOrder", log); err != nil {
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

// ParseCreateOrder is a log parse operation binding the contract event 0x23df57ff35565965f751a83073b22102fcf2bf4fb084463c0c5b68ddc9a5b128.
//
// Solidity: event CreateOrder(uint256 _oid, address user, uint256 _price, address _receiveToken, uint256 _cycle, uint256 _endTime)
func (_VenusMarket *VenusMarketFilterer) ParseCreateOrder(log types.Log) (*VenusMarketCreateOrder, error) {
	event := new(VenusMarketCreateOrder)
	if err := _VenusMarket.contract.UnpackLog(event, "CreateOrder", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VenusMarketSetBlackUserIterator is returned from FilterSetBlackUser and is used to iterate over the raw logs and unpacked data for SetBlackUser events raised by the VenusMarket contract.
type VenusMarketSetBlackUserIterator struct {
	Event *VenusMarketSetBlackUser // Event containing the contract specifics and raw log

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
func (it *VenusMarketSetBlackUserIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VenusMarketSetBlackUser)
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
		it.Event = new(VenusMarketSetBlackUser)
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
func (it *VenusMarketSetBlackUserIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VenusMarketSetBlackUserIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VenusMarketSetBlackUser represents a SetBlackUser event raised by the VenusMarket contract.
type VenusMarketSetBlackUser struct {
	User    common.Address
	CanSell bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSetBlackUser is a free log retrieval operation binding the contract event 0x268e635d8b3370e1ab4560005f02806d3efb314b2aca5cea4a68f4f44eb3d983.
//
// Solidity: event SetBlackUser(address _user, bool _canSell)
func (_VenusMarket *VenusMarketFilterer) FilterSetBlackUser(opts *bind.FilterOpts) (*VenusMarketSetBlackUserIterator, error) {

	logs, sub, err := _VenusMarket.contract.FilterLogs(opts, "SetBlackUser")
	if err != nil {
		return nil, err
	}
	return &VenusMarketSetBlackUserIterator{contract: _VenusMarket.contract, event: "SetBlackUser", logs: logs, sub: sub}, nil
}

// WatchSetBlackUser is a free log subscription operation binding the contract event 0x268e635d8b3370e1ab4560005f02806d3efb314b2aca5cea4a68f4f44eb3d983.
//
// Solidity: event SetBlackUser(address _user, bool _canSell)
func (_VenusMarket *VenusMarketFilterer) WatchSetBlackUser(opts *bind.WatchOpts, sink chan<- *VenusMarketSetBlackUser) (event.Subscription, error) {

	logs, sub, err := _VenusMarket.contract.WatchLogs(opts, "SetBlackUser")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VenusMarketSetBlackUser)
				if err := _VenusMarket.contract.UnpackLog(event, "SetBlackUser", log); err != nil {
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

// ParseSetBlackUser is a log parse operation binding the contract event 0x268e635d8b3370e1ab4560005f02806d3efb314b2aca5cea4a68f4f44eb3d983.
//
// Solidity: event SetBlackUser(address _user, bool _canSell)
func (_VenusMarket *VenusMarketFilterer) ParseSetBlackUser(log types.Log) (*VenusMarketSetBlackUser, error) {
	event := new(VenusMarketSetBlackUser)
	if err := _VenusMarket.contract.UnpackLog(event, "SetBlackUser", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

