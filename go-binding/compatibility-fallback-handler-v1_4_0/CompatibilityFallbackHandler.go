// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package compatibilityfallbackhandler

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

// CompatibilityfallbackhandlerMetaData contains all meta data concerning the Compatibilityfallbackhandler contract.
var CompatibilityfallbackhandlerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractSafe\",\"name\":\"safe\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"encodeMessageDataForSafe\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"getMessageHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractSafe\",\"name\":\"safe\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"getMessageHashForSafe\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getModules\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_dataHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"isValidSignature\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"isValidSignature\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"targetContract\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"calldataPayload\",\"type\":\"bytes\"}],\"name\":\"simulate\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"response\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"tokensReceived\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
}

// CompatibilityfallbackhandlerABI is the input ABI used to generate the binding from.
// Deprecated: Use CompatibilityfallbackhandlerMetaData.ABI instead.
var CompatibilityfallbackhandlerABI = CompatibilityfallbackhandlerMetaData.ABI

// Compatibilityfallbackhandler is an auto generated Go binding around an Ethereum contract.
type Compatibilityfallbackhandler struct {
	CompatibilityfallbackhandlerCaller     // Read-only binding to the contract
	CompatibilityfallbackhandlerTransactor // Write-only binding to the contract
	CompatibilityfallbackhandlerFilterer   // Log filterer for contract events
}

// CompatibilityfallbackhandlerCaller is an auto generated read-only Go binding around an Ethereum contract.
type CompatibilityfallbackhandlerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CompatibilityfallbackhandlerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CompatibilityfallbackhandlerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CompatibilityfallbackhandlerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CompatibilityfallbackhandlerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CompatibilityfallbackhandlerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CompatibilityfallbackhandlerSession struct {
	Contract     *Compatibilityfallbackhandler // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                 // Call options to use throughout this session
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// CompatibilityfallbackhandlerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CompatibilityfallbackhandlerCallerSession struct {
	Contract *CompatibilityfallbackhandlerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                       // Call options to use throughout this session
}

// CompatibilityfallbackhandlerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CompatibilityfallbackhandlerTransactorSession struct {
	Contract     *CompatibilityfallbackhandlerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                       // Transaction auth options to use throughout this session
}

// CompatibilityfallbackhandlerRaw is an auto generated low-level Go binding around an Ethereum contract.
type CompatibilityfallbackhandlerRaw struct {
	Contract *Compatibilityfallbackhandler // Generic contract binding to access the raw methods on
}

// CompatibilityfallbackhandlerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CompatibilityfallbackhandlerCallerRaw struct {
	Contract *CompatibilityfallbackhandlerCaller // Generic read-only contract binding to access the raw methods on
}

// CompatibilityfallbackhandlerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CompatibilityfallbackhandlerTransactorRaw struct {
	Contract *CompatibilityfallbackhandlerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCompatibilityfallbackhandler creates a new instance of Compatibilityfallbackhandler, bound to a specific deployed contract.
func NewCompatibilityfallbackhandler(address common.Address, backend bind.ContractBackend) (*Compatibilityfallbackhandler, error) {
	contract, err := bindCompatibilityfallbackhandler(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Compatibilityfallbackhandler{CompatibilityfallbackhandlerCaller: CompatibilityfallbackhandlerCaller{contract: contract}, CompatibilityfallbackhandlerTransactor: CompatibilityfallbackhandlerTransactor{contract: contract}, CompatibilityfallbackhandlerFilterer: CompatibilityfallbackhandlerFilterer{contract: contract}}, nil
}

// NewCompatibilityfallbackhandlerCaller creates a new read-only instance of Compatibilityfallbackhandler, bound to a specific deployed contract.
func NewCompatibilityfallbackhandlerCaller(address common.Address, caller bind.ContractCaller) (*CompatibilityfallbackhandlerCaller, error) {
	contract, err := bindCompatibilityfallbackhandler(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CompatibilityfallbackhandlerCaller{contract: contract}, nil
}

// NewCompatibilityfallbackhandlerTransactor creates a new write-only instance of Compatibilityfallbackhandler, bound to a specific deployed contract.
func NewCompatibilityfallbackhandlerTransactor(address common.Address, transactor bind.ContractTransactor) (*CompatibilityfallbackhandlerTransactor, error) {
	contract, err := bindCompatibilityfallbackhandler(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CompatibilityfallbackhandlerTransactor{contract: contract}, nil
}

// NewCompatibilityfallbackhandlerFilterer creates a new log filterer instance of Compatibilityfallbackhandler, bound to a specific deployed contract.
func NewCompatibilityfallbackhandlerFilterer(address common.Address, filterer bind.ContractFilterer) (*CompatibilityfallbackhandlerFilterer, error) {
	contract, err := bindCompatibilityfallbackhandler(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CompatibilityfallbackhandlerFilterer{contract: contract}, nil
}

// bindCompatibilityfallbackhandler binds a generic wrapper to an already deployed contract.
func bindCompatibilityfallbackhandler(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CompatibilityfallbackhandlerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Compatibilityfallbackhandler *CompatibilityfallbackhandlerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Compatibilityfallbackhandler.Contract.CompatibilityfallbackhandlerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Compatibilityfallbackhandler *CompatibilityfallbackhandlerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Compatibilityfallbackhandler.Contract.CompatibilityfallbackhandlerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Compatibilityfallbackhandler *CompatibilityfallbackhandlerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Compatibilityfallbackhandler.Contract.CompatibilityfallbackhandlerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Compatibilityfallbackhandler *CompatibilityfallbackhandlerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Compatibilityfallbackhandler.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Compatibilityfallbackhandler *CompatibilityfallbackhandlerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Compatibilityfallbackhandler.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Compatibilityfallbackhandler *CompatibilityfallbackhandlerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Compatibilityfallbackhandler.Contract.contract.Transact(opts, method, params...)
}

// EncodeMessageDataForSafe is a free data retrieval call binding the contract method 0x23031640.
//
// Solidity: function encodeMessageDataForSafe(address safe, bytes message) view returns(bytes)
func (_Compatibilityfallbackhandler *CompatibilityfallbackhandlerCaller) EncodeMessageDataForSafe(opts *bind.CallOpts, safe common.Address, message []byte) ([]byte, error) {
	var out []interface{}
	err := _Compatibilityfallbackhandler.contract.Call(opts, &out, "encodeMessageDataForSafe", safe, message)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// EncodeMessageDataForSafe is a free data retrieval call binding the contract method 0x23031640.
//
// Solidity: function encodeMessageDataForSafe(address safe, bytes message) view returns(bytes)
func (_Compatibilityfallbackhandler *CompatibilityfallbackhandlerSession) EncodeMessageDataForSafe(safe common.Address, message []byte) ([]byte, error) {
	return _Compatibilityfallbackhandler.Contract.EncodeMessageDataForSafe(&_Compatibilityfallbackhandler.CallOpts, safe, message)
}

// EncodeMessageDataForSafe is a free data retrieval call binding the contract method 0x23031640.
//
// Solidity: function encodeMessageDataForSafe(address safe, bytes message) view returns(bytes)
func (_Compatibilityfallbackhandler *CompatibilityfallbackhandlerCallerSession) EncodeMessageDataForSafe(safe common.Address, message []byte) ([]byte, error) {
	return _Compatibilityfallbackhandler.Contract.EncodeMessageDataForSafe(&_Compatibilityfallbackhandler.CallOpts, safe, message)
}

// GetMessageHash is a free data retrieval call binding the contract method 0x0a1028c4.
//
// Solidity: function getMessageHash(bytes message) view returns(bytes32)
func (_Compatibilityfallbackhandler *CompatibilityfallbackhandlerCaller) GetMessageHash(opts *bind.CallOpts, message []byte) ([32]byte, error) {
	var out []interface{}
	err := _Compatibilityfallbackhandler.contract.Call(opts, &out, "getMessageHash", message)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetMessageHash is a free data retrieval call binding the contract method 0x0a1028c4.
//
// Solidity: function getMessageHash(bytes message) view returns(bytes32)
func (_Compatibilityfallbackhandler *CompatibilityfallbackhandlerSession) GetMessageHash(message []byte) ([32]byte, error) {
	return _Compatibilityfallbackhandler.Contract.GetMessageHash(&_Compatibilityfallbackhandler.CallOpts, message)
}

// GetMessageHash is a free data retrieval call binding the contract method 0x0a1028c4.
//
// Solidity: function getMessageHash(bytes message) view returns(bytes32)
func (_Compatibilityfallbackhandler *CompatibilityfallbackhandlerCallerSession) GetMessageHash(message []byte) ([32]byte, error) {
	return _Compatibilityfallbackhandler.Contract.GetMessageHash(&_Compatibilityfallbackhandler.CallOpts, message)
}

// GetMessageHashForSafe is a free data retrieval call binding the contract method 0x6ac24784.
//
// Solidity: function getMessageHashForSafe(address safe, bytes message) view returns(bytes32)
func (_Compatibilityfallbackhandler *CompatibilityfallbackhandlerCaller) GetMessageHashForSafe(opts *bind.CallOpts, safe common.Address, message []byte) ([32]byte, error) {
	var out []interface{}
	err := _Compatibilityfallbackhandler.contract.Call(opts, &out, "getMessageHashForSafe", safe, message)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetMessageHashForSafe is a free data retrieval call binding the contract method 0x6ac24784.
//
// Solidity: function getMessageHashForSafe(address safe, bytes message) view returns(bytes32)
func (_Compatibilityfallbackhandler *CompatibilityfallbackhandlerSession) GetMessageHashForSafe(safe common.Address, message []byte) ([32]byte, error) {
	return _Compatibilityfallbackhandler.Contract.GetMessageHashForSafe(&_Compatibilityfallbackhandler.CallOpts, safe, message)
}

// GetMessageHashForSafe is a free data retrieval call binding the contract method 0x6ac24784.
//
// Solidity: function getMessageHashForSafe(address safe, bytes message) view returns(bytes32)
func (_Compatibilityfallbackhandler *CompatibilityfallbackhandlerCallerSession) GetMessageHashForSafe(safe common.Address, message []byte) ([32]byte, error) {
	return _Compatibilityfallbackhandler.Contract.GetMessageHashForSafe(&_Compatibilityfallbackhandler.CallOpts, safe, message)
}

// GetModules is a free data retrieval call binding the contract method 0xb2494df3.
//
// Solidity: function getModules() view returns(address[])
func (_Compatibilityfallbackhandler *CompatibilityfallbackhandlerCaller) GetModules(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Compatibilityfallbackhandler.contract.Call(opts, &out, "getModules")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetModules is a free data retrieval call binding the contract method 0xb2494df3.
//
// Solidity: function getModules() view returns(address[])
func (_Compatibilityfallbackhandler *CompatibilityfallbackhandlerSession) GetModules() ([]common.Address, error) {
	return _Compatibilityfallbackhandler.Contract.GetModules(&_Compatibilityfallbackhandler.CallOpts)
}

// GetModules is a free data retrieval call binding the contract method 0xb2494df3.
//
// Solidity: function getModules() view returns(address[])
func (_Compatibilityfallbackhandler *CompatibilityfallbackhandlerCallerSession) GetModules() ([]common.Address, error) {
	return _Compatibilityfallbackhandler.Contract.GetModules(&_Compatibilityfallbackhandler.CallOpts)
}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 _dataHash, bytes _signature) view returns(bytes4)
func (_Compatibilityfallbackhandler *CompatibilityfallbackhandlerCaller) IsValidSignature(opts *bind.CallOpts, _dataHash [32]byte, _signature []byte) ([4]byte, error) {
	var out []interface{}
	err := _Compatibilityfallbackhandler.contract.Call(opts, &out, "isValidSignature", _dataHash, _signature)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 _dataHash, bytes _signature) view returns(bytes4)
func (_Compatibilityfallbackhandler *CompatibilityfallbackhandlerSession) IsValidSignature(_dataHash [32]byte, _signature []byte) ([4]byte, error) {
	return _Compatibilityfallbackhandler.Contract.IsValidSignature(&_Compatibilityfallbackhandler.CallOpts, _dataHash, _signature)
}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 _dataHash, bytes _signature) view returns(bytes4)
func (_Compatibilityfallbackhandler *CompatibilityfallbackhandlerCallerSession) IsValidSignature(_dataHash [32]byte, _signature []byte) ([4]byte, error) {
	return _Compatibilityfallbackhandler.Contract.IsValidSignature(&_Compatibilityfallbackhandler.CallOpts, _dataHash, _signature)
}

// IsValidSignature0 is a free data retrieval call binding the contract method 0x20c13b0b.
//
// Solidity: function isValidSignature(bytes _data, bytes _signature) view returns(bytes4)
func (_Compatibilityfallbackhandler *CompatibilityfallbackhandlerCaller) IsValidSignature0(opts *bind.CallOpts, _data []byte, _signature []byte) ([4]byte, error) {
	var out []interface{}
	err := _Compatibilityfallbackhandler.contract.Call(opts, &out, "isValidSignature0", _data, _signature)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// IsValidSignature0 is a free data retrieval call binding the contract method 0x20c13b0b.
//
// Solidity: function isValidSignature(bytes _data, bytes _signature) view returns(bytes4)
func (_Compatibilityfallbackhandler *CompatibilityfallbackhandlerSession) IsValidSignature0(_data []byte, _signature []byte) ([4]byte, error) {
	return _Compatibilityfallbackhandler.Contract.IsValidSignature0(&_Compatibilityfallbackhandler.CallOpts, _data, _signature)
}

// IsValidSignature0 is a free data retrieval call binding the contract method 0x20c13b0b.
//
// Solidity: function isValidSignature(bytes _data, bytes _signature) view returns(bytes4)
func (_Compatibilityfallbackhandler *CompatibilityfallbackhandlerCallerSession) IsValidSignature0(_data []byte, _signature []byte) ([4]byte, error) {
	return _Compatibilityfallbackhandler.Contract.IsValidSignature0(&_Compatibilityfallbackhandler.CallOpts, _data, _signature)
}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) pure returns(bytes4)
func (_Compatibilityfallbackhandler *CompatibilityfallbackhandlerCaller) OnERC1155BatchReceived(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([4]byte, error) {
	var out []interface{}
	err := _Compatibilityfallbackhandler.contract.Call(opts, &out, "onERC1155BatchReceived", arg0, arg1, arg2, arg3, arg4)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) pure returns(bytes4)
func (_Compatibilityfallbackhandler *CompatibilityfallbackhandlerSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([4]byte, error) {
	return _Compatibilityfallbackhandler.Contract.OnERC1155BatchReceived(&_Compatibilityfallbackhandler.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) pure returns(bytes4)
func (_Compatibilityfallbackhandler *CompatibilityfallbackhandlerCallerSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([4]byte, error) {
	return _Compatibilityfallbackhandler.Contract.OnERC1155BatchReceived(&_Compatibilityfallbackhandler.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) pure returns(bytes4)
func (_Compatibilityfallbackhandler *CompatibilityfallbackhandlerCaller) OnERC1155Received(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	var out []interface{}
	err := _Compatibilityfallbackhandler.contract.Call(opts, &out, "onERC1155Received", arg0, arg1, arg2, arg3, arg4)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) pure returns(bytes4)
func (_Compatibilityfallbackhandler *CompatibilityfallbackhandlerSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	return _Compatibilityfallbackhandler.Contract.OnERC1155Received(&_Compatibilityfallbackhandler.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) pure returns(bytes4)
func (_Compatibilityfallbackhandler *CompatibilityfallbackhandlerCallerSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	return _Compatibilityfallbackhandler.Contract.OnERC1155Received(&_Compatibilityfallbackhandler.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_Compatibilityfallbackhandler *CompatibilityfallbackhandlerCaller) OnERC721Received(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	var out []interface{}
	err := _Compatibilityfallbackhandler.contract.Call(opts, &out, "onERC721Received", arg0, arg1, arg2, arg3)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_Compatibilityfallbackhandler *CompatibilityfallbackhandlerSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _Compatibilityfallbackhandler.Contract.OnERC721Received(&_Compatibilityfallbackhandler.CallOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_Compatibilityfallbackhandler *CompatibilityfallbackhandlerCallerSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _Compatibilityfallbackhandler.Contract.OnERC721Received(&_Compatibilityfallbackhandler.CallOpts, arg0, arg1, arg2, arg3)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Compatibilityfallbackhandler *CompatibilityfallbackhandlerCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Compatibilityfallbackhandler.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Compatibilityfallbackhandler *CompatibilityfallbackhandlerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Compatibilityfallbackhandler.Contract.SupportsInterface(&_Compatibilityfallbackhandler.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Compatibilityfallbackhandler *CompatibilityfallbackhandlerCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Compatibilityfallbackhandler.Contract.SupportsInterface(&_Compatibilityfallbackhandler.CallOpts, interfaceId)
}

// TokensReceived is a free data retrieval call binding the contract method 0x0023de29.
//
// Solidity: function tokensReceived(address , address , address , uint256 , bytes , bytes ) pure returns()
func (_Compatibilityfallbackhandler *CompatibilityfallbackhandlerCaller) TokensReceived(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 common.Address, arg3 *big.Int, arg4 []byte, arg5 []byte) error {
	var out []interface{}
	err := _Compatibilityfallbackhandler.contract.Call(opts, &out, "tokensReceived", arg0, arg1, arg2, arg3, arg4, arg5)

	if err != nil {
		return err
	}

	return err

}

// TokensReceived is a free data retrieval call binding the contract method 0x0023de29.
//
// Solidity: function tokensReceived(address , address , address , uint256 , bytes , bytes ) pure returns()
func (_Compatibilityfallbackhandler *CompatibilityfallbackhandlerSession) TokensReceived(arg0 common.Address, arg1 common.Address, arg2 common.Address, arg3 *big.Int, arg4 []byte, arg5 []byte) error {
	return _Compatibilityfallbackhandler.Contract.TokensReceived(&_Compatibilityfallbackhandler.CallOpts, arg0, arg1, arg2, arg3, arg4, arg5)
}

// TokensReceived is a free data retrieval call binding the contract method 0x0023de29.
//
// Solidity: function tokensReceived(address , address , address , uint256 , bytes , bytes ) pure returns()
func (_Compatibilityfallbackhandler *CompatibilityfallbackhandlerCallerSession) TokensReceived(arg0 common.Address, arg1 common.Address, arg2 common.Address, arg3 *big.Int, arg4 []byte, arg5 []byte) error {
	return _Compatibilityfallbackhandler.Contract.TokensReceived(&_Compatibilityfallbackhandler.CallOpts, arg0, arg1, arg2, arg3, arg4, arg5)
}

// Simulate is a paid mutator transaction binding the contract method 0xbd61951d.
//
// Solidity: function simulate(address targetContract, bytes calldataPayload) returns(bytes response)
func (_Compatibilityfallbackhandler *CompatibilityfallbackhandlerTransactor) Simulate(opts *bind.TransactOpts, targetContract common.Address, calldataPayload []byte) (*types.Transaction, error) {
	return _Compatibilityfallbackhandler.contract.Transact(opts, "simulate", targetContract, calldataPayload)
}

// Simulate is a paid mutator transaction binding the contract method 0xbd61951d.
//
// Solidity: function simulate(address targetContract, bytes calldataPayload) returns(bytes response)
func (_Compatibilityfallbackhandler *CompatibilityfallbackhandlerSession) Simulate(targetContract common.Address, calldataPayload []byte) (*types.Transaction, error) {
	return _Compatibilityfallbackhandler.Contract.Simulate(&_Compatibilityfallbackhandler.TransactOpts, targetContract, calldataPayload)
}

// Simulate is a paid mutator transaction binding the contract method 0xbd61951d.
//
// Solidity: function simulate(address targetContract, bytes calldataPayload) returns(bytes response)
func (_Compatibilityfallbackhandler *CompatibilityfallbackhandlerTransactorSession) Simulate(targetContract common.Address, calldataPayload []byte) (*types.Transaction, error) {
	return _Compatibilityfallbackhandler.Contract.Simulate(&_Compatibilityfallbackhandler.TransactOpts, targetContract, calldataPayload)
}
