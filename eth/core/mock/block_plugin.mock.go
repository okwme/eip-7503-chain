// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mock

import (
	"context"
	"github.com/berachain/polaris/eth/core"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"sync"
)

// Ensure, that BlockPluginMock does implement core.BlockPlugin.
// If this is not the case, regenerate this file with moq.
var _ core.BlockPlugin = &BlockPluginMock{}

// BlockPluginMock is a mock implementation of core.BlockPlugin.
//
//	func TestSomethingThatUsesBlockPlugin(t *testing.T) {
//
//		// make and configure a mocked core.BlockPlugin
//		mockedBlockPlugin := &BlockPluginMock{
//			GetHeaderByHashFunc: func(hash common.Hash) (*types.Header, error) {
//				panic("mock out the GetHeaderByHash method")
//			},
//			GetHeaderByNumberFunc: func(v uint64) (*types.Header, error) {
//				panic("mock out the GetHeaderByNumber method")
//			},
//			PrepareFunc: func(contextMoqParam context.Context)  {
//				panic("mock out the Prepare method")
//			},
//			StoreHeaderFunc: func(header *types.Header) error {
//				panic("mock out the StoreHeader method")
//			},
//		}
//
//		// use mockedBlockPlugin in code that requires core.BlockPlugin
//		// and then make assertions.
//
//	}
type BlockPluginMock struct {
	// GetHeaderByHashFunc mocks the GetHeaderByHash method.
	GetHeaderByHashFunc func(hash common.Hash) (*types.Header, error)

	// GetHeaderByNumberFunc mocks the GetHeaderByNumber method.
	GetHeaderByNumberFunc func(v uint64) (*types.Header, error)

	// PrepareFunc mocks the Prepare method.
	PrepareFunc func(contextMoqParam context.Context)

	// StoreHeaderFunc mocks the StoreHeader method.
	StoreHeaderFunc func(header *types.Header) error

	// calls tracks calls to the methods.
	calls struct {
		// GetHeaderByHash holds details about calls to the GetHeaderByHash method.
		GetHeaderByHash []struct {
			// Hash is the hash argument value.
			Hash common.Hash
		}
		// GetHeaderByNumber holds details about calls to the GetHeaderByNumber method.
		GetHeaderByNumber []struct {
			// V is the v argument value.
			V uint64
		}
		// Prepare holds details about calls to the Prepare method.
		Prepare []struct {
			// ContextMoqParam is the contextMoqParam argument value.
			ContextMoqParam context.Context
		}
		// StoreHeader holds details about calls to the StoreHeader method.
		StoreHeader []struct {
			// Header is the header argument value.
			Header *types.Header
		}
	}
	lockGetHeaderByHash   sync.RWMutex
	lockGetHeaderByNumber sync.RWMutex
	lockPrepare           sync.RWMutex
	lockStoreHeader       sync.RWMutex
}

// GetHeaderByHash calls GetHeaderByHashFunc.
func (mock *BlockPluginMock) GetHeaderByHash(hash common.Hash) (*types.Header, error) {
	if mock.GetHeaderByHashFunc == nil {
		panic("BlockPluginMock.GetHeaderByHashFunc: method is nil but BlockPlugin.GetHeaderByHash was just called")
	}
	callInfo := struct {
		Hash common.Hash
	}{
		Hash: hash,
	}
	mock.lockGetHeaderByHash.Lock()
	mock.calls.GetHeaderByHash = append(mock.calls.GetHeaderByHash, callInfo)
	mock.lockGetHeaderByHash.Unlock()
	return mock.GetHeaderByHashFunc(hash)
}

// GetHeaderByHashCalls gets all the calls that were made to GetHeaderByHash.
// Check the length with:
//
//	len(mockedBlockPlugin.GetHeaderByHashCalls())
func (mock *BlockPluginMock) GetHeaderByHashCalls() []struct {
	Hash common.Hash
} {
	var calls []struct {
		Hash common.Hash
	}
	mock.lockGetHeaderByHash.RLock()
	calls = mock.calls.GetHeaderByHash
	mock.lockGetHeaderByHash.RUnlock()
	return calls
}

// GetHeaderByNumber calls GetHeaderByNumberFunc.
func (mock *BlockPluginMock) GetHeaderByNumber(v uint64) (*types.Header, error) {
	if mock.GetHeaderByNumberFunc == nil {
		panic("BlockPluginMock.GetHeaderByNumberFunc: method is nil but BlockPlugin.GetHeaderByNumber was just called")
	}
	callInfo := struct {
		V uint64
	}{
		V: v,
	}
	mock.lockGetHeaderByNumber.Lock()
	mock.calls.GetHeaderByNumber = append(mock.calls.GetHeaderByNumber, callInfo)
	mock.lockGetHeaderByNumber.Unlock()
	return mock.GetHeaderByNumberFunc(v)
}

// GetHeaderByNumberCalls gets all the calls that were made to GetHeaderByNumber.
// Check the length with:
//
//	len(mockedBlockPlugin.GetHeaderByNumberCalls())
func (mock *BlockPluginMock) GetHeaderByNumberCalls() []struct {
	V uint64
} {
	var calls []struct {
		V uint64
	}
	mock.lockGetHeaderByNumber.RLock()
	calls = mock.calls.GetHeaderByNumber
	mock.lockGetHeaderByNumber.RUnlock()
	return calls
}

// Prepare calls PrepareFunc.
func (mock *BlockPluginMock) Prepare(contextMoqParam context.Context) {
	if mock.PrepareFunc == nil {
		panic("BlockPluginMock.PrepareFunc: method is nil but BlockPlugin.Prepare was just called")
	}
	callInfo := struct {
		ContextMoqParam context.Context
	}{
		ContextMoqParam: contextMoqParam,
	}
	mock.lockPrepare.Lock()
	mock.calls.Prepare = append(mock.calls.Prepare, callInfo)
	mock.lockPrepare.Unlock()
	mock.PrepareFunc(contextMoqParam)
}

// PrepareCalls gets all the calls that were made to Prepare.
// Check the length with:
//
//	len(mockedBlockPlugin.PrepareCalls())
func (mock *BlockPluginMock) PrepareCalls() []struct {
	ContextMoqParam context.Context
} {
	var calls []struct {
		ContextMoqParam context.Context
	}
	mock.lockPrepare.RLock()
	calls = mock.calls.Prepare
	mock.lockPrepare.RUnlock()
	return calls
}

// StoreHeader calls StoreHeaderFunc.
func (mock *BlockPluginMock) StoreHeader(header *types.Header) error {
	if mock.StoreHeaderFunc == nil {
		panic("BlockPluginMock.StoreHeaderFunc: method is nil but BlockPlugin.StoreHeader was just called")
	}
	callInfo := struct {
		Header *types.Header
	}{
		Header: header,
	}
	mock.lockStoreHeader.Lock()
	mock.calls.StoreHeader = append(mock.calls.StoreHeader, callInfo)
	mock.lockStoreHeader.Unlock()
	return mock.StoreHeaderFunc(header)
}

// StoreHeaderCalls gets all the calls that were made to StoreHeader.
// Check the length with:
//
//	len(mockedBlockPlugin.StoreHeaderCalls())
func (mock *BlockPluginMock) StoreHeaderCalls() []struct {
	Header *types.Header
} {
	var calls []struct {
		Header *types.Header
	}
	mock.lockStoreHeader.RLock()
	calls = mock.calls.StoreHeader
	mock.lockStoreHeader.RUnlock()
	return calls
}
