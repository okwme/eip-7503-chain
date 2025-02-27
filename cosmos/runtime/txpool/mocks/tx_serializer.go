// Code generated by mockery v2.36.1. DO NOT EDIT.

package mocks

import (
	cosmos_sdktypes "github.com/cosmos/cosmos-sdk/types"
	mock "github.com/stretchr/testify/mock"

	types "github.com/ethereum/go-ethereum/core/types"
)

// TxSerializer is an autogenerated mock type for the TxSerializer type
type TxSerializer struct {
	mock.Mock
}

type TxSerializer_Expecter struct {
	mock *mock.Mock
}

func (_m *TxSerializer) EXPECT() *TxSerializer_Expecter {
	return &TxSerializer_Expecter{mock: &_m.Mock}
}

// ToSdkTx provides a mock function with given fields: signedTx, gasLimit
func (_m *TxSerializer) ToSdkTx(signedTx *types.Transaction, gasLimit uint64) (cosmos_sdktypes.Tx, error) {
	ret := _m.Called(signedTx, gasLimit)

	var r0 cosmos_sdktypes.Tx
	var r1 error
	if rf, ok := ret.Get(0).(func(*types.Transaction, uint64) (cosmos_sdktypes.Tx, error)); ok {
		return rf(signedTx, gasLimit)
	}
	if rf, ok := ret.Get(0).(func(*types.Transaction, uint64) cosmos_sdktypes.Tx); ok {
		r0 = rf(signedTx, gasLimit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(cosmos_sdktypes.Tx)
		}
	}

	if rf, ok := ret.Get(1).(func(*types.Transaction, uint64) error); ok {
		r1 = rf(signedTx, gasLimit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TxSerializer_ToSdkTx_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ToSdkTx'
type TxSerializer_ToSdkTx_Call struct {
	*mock.Call
}

// ToSdkTx is a helper method to define mock.On call
//   - signedTx *types.Transaction
//   - gasLimit uint64
func (_e *TxSerializer_Expecter) ToSdkTx(signedTx interface{}, gasLimit interface{}) *TxSerializer_ToSdkTx_Call {
	return &TxSerializer_ToSdkTx_Call{Call: _e.mock.On("ToSdkTx", signedTx, gasLimit)}
}

func (_c *TxSerializer_ToSdkTx_Call) Run(run func(signedTx *types.Transaction, gasLimit uint64)) *TxSerializer_ToSdkTx_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*types.Transaction), args[1].(uint64))
	})
	return _c
}

func (_c *TxSerializer_ToSdkTx_Call) Return(_a0 cosmos_sdktypes.Tx, _a1 error) *TxSerializer_ToSdkTx_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *TxSerializer_ToSdkTx_Call) RunAndReturn(run func(*types.Transaction, uint64) (cosmos_sdktypes.Tx, error)) *TxSerializer_ToSdkTx_Call {
	_c.Call.Return(run)
	return _c
}

// ToSdkTxBytes provides a mock function with given fields: signedTx, gasLimit
func (_m *TxSerializer) ToSdkTxBytes(signedTx *types.Transaction, gasLimit uint64) ([]byte, error) {
	ret := _m.Called(signedTx, gasLimit)

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(*types.Transaction, uint64) ([]byte, error)); ok {
		return rf(signedTx, gasLimit)
	}
	if rf, ok := ret.Get(0).(func(*types.Transaction, uint64) []byte); ok {
		r0 = rf(signedTx, gasLimit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(*types.Transaction, uint64) error); ok {
		r1 = rf(signedTx, gasLimit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TxSerializer_ToSdkTxBytes_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ToSdkTxBytes'
type TxSerializer_ToSdkTxBytes_Call struct {
	*mock.Call
}

// ToSdkTxBytes is a helper method to define mock.On call
//   - signedTx *types.Transaction
//   - gasLimit uint64
func (_e *TxSerializer_Expecter) ToSdkTxBytes(signedTx interface{}, gasLimit interface{}) *TxSerializer_ToSdkTxBytes_Call {
	return &TxSerializer_ToSdkTxBytes_Call{Call: _e.mock.On("ToSdkTxBytes", signedTx, gasLimit)}
}

func (_c *TxSerializer_ToSdkTxBytes_Call) Run(run func(signedTx *types.Transaction, gasLimit uint64)) *TxSerializer_ToSdkTxBytes_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*types.Transaction), args[1].(uint64))
	})
	return _c
}

func (_c *TxSerializer_ToSdkTxBytes_Call) Return(_a0 []byte, _a1 error) *TxSerializer_ToSdkTxBytes_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *TxSerializer_ToSdkTxBytes_Call) RunAndReturn(run func(*types.Transaction, uint64) ([]byte, error)) *TxSerializer_ToSdkTxBytes_Call {
	_c.Call.Return(run)
	return _c
}

// NewTxSerializer creates a new instance of TxSerializer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTxSerializer(t interface {
	mock.TestingT
	Cleanup(func())
}) *TxSerializer {
	mock := &TxSerializer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
