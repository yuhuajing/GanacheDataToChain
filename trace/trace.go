// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package trace

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

// AgriTraOrgInfo is an auto generated low-level Go binding around an user-defined struct.
type AgriTraOrgInfo struct {
	Number     string
	Workamount string
	Persion    string
	Workmethod string
	Worktime   string
	Remarks    string
}

// TraceMetaData contains all meta data concerning the Trace contract.
var TraceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"number\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"workamount\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"persion\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"workmethod\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"worktime\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"remarks\",\"type\":\"string\"}],\"name\":\"addorupdateData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"tracedata\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"number\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"workamount\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"persion\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"workmethod\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"worktime\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"remarks\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"number\",\"type\":\"string\"}],\"name\":\"tracedataByNum\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"number\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"workamount\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"persion\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"workmethod\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"worktime\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"remarks\",\"type\":\"string\"}],\"internalType\":\"structAgriTra.OrgInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50611181806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c80631bb2bfea146100465780635c7e9495146100765780638a205085146100ab575b600080fd5b610060600480360381019061005b9190610a70565b6100c7565b60405161006d9190610be4565b60405180910390f35b610090600480360381019061008b9190610a70565b61046b565b6040516100a296959493929190610c50565b60405180910390f35b6100c560048036038101906100c09190610cdb565b6107ed565b005b6100cf6108e0565b6000826040516100df9190610e4c565b90815260200160405180910390206040518060c001604052908160008201805461010890610e92565b80601f016020809104026020016040519081016040528092919081815260200182805461013490610e92565b80156101815780601f1061015657610100808354040283529160200191610181565b820191906000526020600020905b81548152906001019060200180831161016457829003601f168201915b5050505050815260200160018201805461019a90610e92565b80601f01602080910402602001604051908101604052809291908181526020018280546101c690610e92565b80156102135780601f106101e857610100808354040283529160200191610213565b820191906000526020600020905b8154815290600101906020018083116101f657829003601f168201915b5050505050815260200160028201805461022c90610e92565b80601f016020809104026020016040519081016040528092919081815260200182805461025890610e92565b80156102a55780601f1061027a576101008083540402835291602001916102a5565b820191906000526020600020905b81548152906001019060200180831161028857829003601f168201915b505050505081526020016003820180546102be90610e92565b80601f01602080910402602001604051908101604052809291908181526020018280546102ea90610e92565b80156103375780601f1061030c57610100808354040283529160200191610337565b820191906000526020600020905b81548152906001019060200180831161031a57829003601f168201915b5050505050815260200160048201805461035090610e92565b80601f016020809104026020016040519081016040528092919081815260200182805461037c90610e92565b80156103c95780601f1061039e576101008083540402835291602001916103c9565b820191906000526020600020905b8154815290600101906020018083116103ac57829003601f168201915b505050505081526020016005820180546103e290610e92565b80601f016020809104026020016040519081016040528092919081815260200182805461040e90610e92565b801561045b5780601f106104305761010080835404028352916020019161045b565b820191906000526020600020905b81548152906001019060200180831161043e57829003601f168201915b5050505050815250509050919050565b6000818051602081018201805184825260208301602085012081835280955050505050506000915090508060000180546104a490610e92565b80601f01602080910402602001604051908101604052809291908181526020018280546104d090610e92565b801561051d5780601f106104f25761010080835404028352916020019161051d565b820191906000526020600020905b81548152906001019060200180831161050057829003601f168201915b50505050509080600101805461053290610e92565b80601f016020809104026020016040519081016040528092919081815260200182805461055e90610e92565b80156105ab5780601f10610580576101008083540402835291602001916105ab565b820191906000526020600020905b81548152906001019060200180831161058e57829003601f168201915b5050505050908060020180546105c090610e92565b80601f01602080910402602001604051908101604052809291908181526020018280546105ec90610e92565b80156106395780601f1061060e57610100808354040283529160200191610639565b820191906000526020600020905b81548152906001019060200180831161061c57829003601f168201915b50505050509080600301805461064e90610e92565b80601f016020809104026020016040519081016040528092919081815260200182805461067a90610e92565b80156106c75780601f1061069c576101008083540402835291602001916106c7565b820191906000526020600020905b8154815290600101906020018083116106aa57829003601f168201915b5050505050908060040180546106dc90610e92565b80601f016020809104026020016040519081016040528092919081815260200182805461070890610e92565b80156107555780601f1061072a57610100808354040283529160200191610755565b820191906000526020600020905b81548152906001019060200180831161073857829003601f168201915b50505050509080600501805461076a90610e92565b80601f016020809104026020016040519081016040528092919081815260200182805461079690610e92565b80156107e35780601f106107b8576101008083540402835291602001916107e3565b820191906000526020600020905b8154815290600101906020018083116107c657829003601f168201915b5050505050905086565b846000876040516107fe9190610e4c565b9081526020016040518091039020600101908161081b9190611079565b508360008760405161082d9190610e4c565b9081526020016040518091039020600201908161084a9190611079565b508260008760405161085c9190610e4c565b908152602001604051809103902060030190816108799190611079565b508160008760405161088b9190610e4c565b908152602001604051809103902060040190816108a89190611079565b50806000876040516108ba9190610e4c565b908152602001604051809103902060050190816108d79190611079565b50505050505050565b6040518060c001604052806060815260200160608152602001606081526020016060815260200160608152602001606081525090565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b61097d82610934565b810181811067ffffffffffffffff8211171561099c5761099b610945565b5b80604052505050565b60006109af610916565b90506109bb8282610974565b919050565b600067ffffffffffffffff8211156109db576109da610945565b5b6109e482610934565b9050602081019050919050565b82818337600083830152505050565b6000610a13610a0e846109c0565b6109a5565b905082815260208101848484011115610a2f57610a2e61092f565b5b610a3a8482856109f1565b509392505050565b600082601f830112610a5757610a5661092a565b5b8135610a67848260208601610a00565b91505092915050565b600060208284031215610a8657610a85610920565b5b600082013567ffffffffffffffff811115610aa457610aa3610925565b5b610ab084828501610a42565b91505092915050565b600081519050919050565b600082825260208201905092915050565b60005b83811015610af3578082015181840152602081019050610ad8565b60008484015250505050565b6000610b0a82610ab9565b610b148185610ac4565b9350610b24818560208601610ad5565b610b2d81610934565b840191505092915050565b600060c0830160008301518482036000860152610b558282610aff565b91505060208301518482036020860152610b6f8282610aff565b91505060408301518482036040860152610b898282610aff565b91505060608301518482036060860152610ba38282610aff565b91505060808301518482036080860152610bbd8282610aff565b91505060a083015184820360a0860152610bd78282610aff565b9150508091505092915050565b60006020820190508181036000830152610bfe8184610b38565b905092915050565b600082825260208201905092915050565b6000610c2282610ab9565b610c2c8185610c06565b9350610c3c818560208601610ad5565b610c4581610934565b840191505092915050565b600060c0820190508181036000830152610c6a8189610c17565b90508181036020830152610c7e8188610c17565b90508181036040830152610c928187610c17565b90508181036060830152610ca68186610c17565b90508181036080830152610cba8185610c17565b905081810360a0830152610cce8184610c17565b9050979650505050505050565b60008060008060008060c08789031215610cf857610cf7610920565b5b600087013567ffffffffffffffff811115610d1657610d15610925565b5b610d2289828a01610a42565b965050602087013567ffffffffffffffff811115610d4357610d42610925565b5b610d4f89828a01610a42565b955050604087013567ffffffffffffffff811115610d7057610d6f610925565b5b610d7c89828a01610a42565b945050606087013567ffffffffffffffff811115610d9d57610d9c610925565b5b610da989828a01610a42565b935050608087013567ffffffffffffffff811115610dca57610dc9610925565b5b610dd689828a01610a42565b92505060a087013567ffffffffffffffff811115610df757610df6610925565b5b610e0389828a01610a42565b9150509295509295509295565b600081905092915050565b6000610e2682610ab9565b610e308185610e10565b9350610e40818560208601610ad5565b80840191505092915050565b6000610e588284610e1b565b915081905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680610eaa57607f821691505b602082108103610ebd57610ebc610e63565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302610f257fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610ee8565b610f2f8683610ee8565b95508019841693508086168417925050509392505050565b6000819050919050565b6000819050919050565b6000610f76610f71610f6c84610f47565b610f51565b610f47565b9050919050565b6000819050919050565b610f9083610f5b565b610fa4610f9c82610f7d565b848454610ef5565b825550505050565b600090565b610fb9610fac565b610fc4818484610f87565b505050565b5b81811015610fe857610fdd600082610fb1565b600181019050610fca565b5050565b601f82111561102d57610ffe81610ec3565b61100784610ed8565b81016020851015611016578190505b61102a61102285610ed8565b830182610fc9565b50505b505050565b600082821c905092915050565b600061105060001984600802611032565b1980831691505092915050565b6000611069838361103f565b9150826002028217905092915050565b61108282610ab9565b67ffffffffffffffff81111561109b5761109a610945565b5b6110a58254610e92565b6110b0828285610fec565b600060209050601f8311600181146110e357600084156110d1578287015190505b6110db858261105d565b865550611143565b601f1984166110f186610ec3565b60005b82811015611119578489015182556001820191506020850194506020810190506110f4565b868310156111365784890151611132601f89168261103f565b8355505b6001600288020188555050505b50505050505056fea26469706673582212202e2dd85f4c0e24a4f4a4744abfdecf8a3e606758394bc95d200b81f2b758ad9c64736f6c63430008180033",
}

// TraceABI is the input ABI used to generate the binding from.
// Deprecated: Use TraceMetaData.ABI instead.
var TraceABI = TraceMetaData.ABI

// TraceBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TraceMetaData.Bin instead.
var TraceBin = TraceMetaData.Bin

// DeployTrace deploys a new Ethereum contract, binding an instance of Trace to it.
func DeployTrace(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Trace, error) {
	parsed, err := TraceMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TraceBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Trace{TraceCaller: TraceCaller{contract: contract}, TraceTransactor: TraceTransactor{contract: contract}, TraceFilterer: TraceFilterer{contract: contract}}, nil
}

// Trace is an auto generated Go binding around an Ethereum contract.
type Trace struct {
	TraceCaller     // Read-only binding to the contract
	TraceTransactor // Write-only binding to the contract
	TraceFilterer   // Log filterer for contract events
}

// TraceCaller is an auto generated read-only Go binding around an Ethereum contract.
type TraceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TraceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TraceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TraceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TraceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TraceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TraceSession struct {
	Contract     *Trace            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TraceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TraceCallerSession struct {
	Contract *TraceCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TraceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TraceTransactorSession struct {
	Contract     *TraceTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TraceRaw is an auto generated low-level Go binding around an Ethereum contract.
type TraceRaw struct {
	Contract *Trace // Generic contract binding to access the raw methods on
}

// TraceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TraceCallerRaw struct {
	Contract *TraceCaller // Generic read-only contract binding to access the raw methods on
}

// TraceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TraceTransactorRaw struct {
	Contract *TraceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTrace creates a new instance of Trace, bound to a specific deployed contract.
func NewTrace(address common.Address, backend bind.ContractBackend) (*Trace, error) {
	contract, err := bindTrace(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Trace{TraceCaller: TraceCaller{contract: contract}, TraceTransactor: TraceTransactor{contract: contract}, TraceFilterer: TraceFilterer{contract: contract}}, nil
}

// NewTraceCaller creates a new read-only instance of Trace, bound to a specific deployed contract.
func NewTraceCaller(address common.Address, caller bind.ContractCaller) (*TraceCaller, error) {
	contract, err := bindTrace(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TraceCaller{contract: contract}, nil
}

// NewTraceTransactor creates a new write-only instance of Trace, bound to a specific deployed contract.
func NewTraceTransactor(address common.Address, transactor bind.ContractTransactor) (*TraceTransactor, error) {
	contract, err := bindTrace(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TraceTransactor{contract: contract}, nil
}

// NewTraceFilterer creates a new log filterer instance of Trace, bound to a specific deployed contract.
func NewTraceFilterer(address common.Address, filterer bind.ContractFilterer) (*TraceFilterer, error) {
	contract, err := bindTrace(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TraceFilterer{contract: contract}, nil
}

// bindTrace binds a generic wrapper to an already deployed contract.
func bindTrace(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TraceMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Trace *TraceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Trace.Contract.TraceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Trace *TraceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Trace.Contract.TraceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Trace *TraceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Trace.Contract.TraceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Trace *TraceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Trace.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Trace *TraceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Trace.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Trace *TraceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Trace.Contract.contract.Transact(opts, method, params...)
}

// Tracedata is a free data retrieval call binding the contract method 0x5c7e9495.
//
// Solidity: function tracedata(string ) view returns(string number, string workamount, string persion, string workmethod, string worktime, string remarks)
func (_Trace *TraceCaller) Tracedata(opts *bind.CallOpts, arg0 string) (struct {
	Number     string
	Workamount string
	Persion    string
	Workmethod string
	Worktime   string
	Remarks    string
}, error) {
	var out []interface{}
	err := _Trace.contract.Call(opts, &out, "tracedata", arg0)

	outstruct := new(struct {
		Number     string
		Workamount string
		Persion    string
		Workmethod string
		Worktime   string
		Remarks    string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Number = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Workamount = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Persion = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Workmethod = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.Worktime = *abi.ConvertType(out[4], new(string)).(*string)
	outstruct.Remarks = *abi.ConvertType(out[5], new(string)).(*string)

	return *outstruct, err

}

// Tracedata is a free data retrieval call binding the contract method 0x5c7e9495.
//
// Solidity: function tracedata(string ) view returns(string number, string workamount, string persion, string workmethod, string worktime, string remarks)
func (_Trace *TraceSession) Tracedata(arg0 string) (struct {
	Number     string
	Workamount string
	Persion    string
	Workmethod string
	Worktime   string
	Remarks    string
}, error) {
	return _Trace.Contract.Tracedata(&_Trace.CallOpts, arg0)
}

// Tracedata is a free data retrieval call binding the contract method 0x5c7e9495.
//
// Solidity: function tracedata(string ) view returns(string number, string workamount, string persion, string workmethod, string worktime, string remarks)
func (_Trace *TraceCallerSession) Tracedata(arg0 string) (struct {
	Number     string
	Workamount string
	Persion    string
	Workmethod string
	Worktime   string
	Remarks    string
}, error) {
	return _Trace.Contract.Tracedata(&_Trace.CallOpts, arg0)
}

// TracedataByNum is a free data retrieval call binding the contract method 0x1bb2bfea.
//
// Solidity: function tracedataByNum(string number) view returns((string,string,string,string,string,string))
func (_Trace *TraceCaller) TracedataByNum(opts *bind.CallOpts, number string) (AgriTraOrgInfo, error) {
	var out []interface{}
	err := _Trace.contract.Call(opts, &out, "tracedataByNum", number)

	if err != nil {
		return *new(AgriTraOrgInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(AgriTraOrgInfo)).(*AgriTraOrgInfo)

	return out0, err

}

// TracedataByNum is a free data retrieval call binding the contract method 0x1bb2bfea.
//
// Solidity: function tracedataByNum(string number) view returns((string,string,string,string,string,string))
func (_Trace *TraceSession) TracedataByNum(number string) (AgriTraOrgInfo, error) {
	return _Trace.Contract.TracedataByNum(&_Trace.CallOpts, number)
}

// TracedataByNum is a free data retrieval call binding the contract method 0x1bb2bfea.
//
// Solidity: function tracedataByNum(string number) view returns((string,string,string,string,string,string))
func (_Trace *TraceCallerSession) TracedataByNum(number string) (AgriTraOrgInfo, error) {
	return _Trace.Contract.TracedataByNum(&_Trace.CallOpts, number)
}

// AddorupdateData is a paid mutator transaction binding the contract method 0x8a205085.
//
// Solidity: function addorupdateData(string number, string workamount, string persion, string workmethod, string worktime, string remarks) returns()
func (_Trace *TraceTransactor) AddorupdateData(opts *bind.TransactOpts, number string, workamount string, persion string, workmethod string, worktime string, remarks string) (*types.Transaction, error) {
	return _Trace.contract.Transact(opts, "addorupdateData", number, workamount, persion, workmethod, worktime, remarks)
}

// AddorupdateData is a paid mutator transaction binding the contract method 0x8a205085.
//
// Solidity: function addorupdateData(string number, string workamount, string persion, string workmethod, string worktime, string remarks) returns()
func (_Trace *TraceSession) AddorupdateData(number string, workamount string, persion string, workmethod string, worktime string, remarks string) (*types.Transaction, error) {
	return _Trace.Contract.AddorupdateData(&_Trace.TransactOpts, number, workamount, persion, workmethod, worktime, remarks)
}

// AddorupdateData is a paid mutator transaction binding the contract method 0x8a205085.
//
// Solidity: function addorupdateData(string number, string workamount, string persion, string workmethod, string worktime, string remarks) returns()
func (_Trace *TraceTransactorSession) AddorupdateData(number string, workamount string, persion string, workmethod string, worktime string, remarks string) (*types.Transaction, error) {
	return _Trace.Contract.AddorupdateData(&_Trace.TransactOpts, number, workamount, persion, workmethod, worktime, remarks)
}
