// Code generated by moq; DO NOT EDIT
// github.com/matryer/moq

package models

import (
	"sync"
)

var (
	lockIMyDataMockGetData sync.RWMutex
)

// IMyDataMock is a mock implementation of IMyData.
//
//     func TestSomethingThatUsesIMyData(t *testing.T) {
//
//         // make and configure a mocked IMyData
//         mockedIMyData := &IMyDataMock{
//             GetDataFunc: func(i int) (string, error) {
// 	               panic("TODO: mock out the GetData method")
//             },
//         }
//
//         // TODO: use mockedIMyData in code that requires IMyData
//         //       and then make assertions.
//
//     }
type IMyDataMock struct {
	// GetDataFunc mocks the GetData method.
	GetDataFunc func(i int) (string, error)

	// calls tracks calls to the methods.
	calls struct {
		// GetData holds details about calls to the GetData method.
		GetData []struct {
			// I is the i argument value.
			I int
		}
	}
}

// GetData calls GetDataFunc.
func (mock *IMyDataMock) GetData(i int) (string, error) {
	if mock.GetDataFunc == nil {
		panic("moq: IMyDataMock.GetDataFunc is nil but IMyData.GetData was just called")
	}
	callInfo := struct {
		I int
	}{
		I: i,
	}
	lockIMyDataMockGetData.Lock()
	mock.calls.GetData = append(mock.calls.GetData, callInfo)
	lockIMyDataMockGetData.Unlock()
	return mock.GetDataFunc(i)
}

// GetDataCalls gets all the calls that were made to GetData.
// Check the length with:
//     len(mockedIMyData.GetDataCalls())
func (mock *IMyDataMock) GetDataCalls() []struct {
	I int
} {
	var calls []struct {
		I int
	}
	lockIMyDataMockGetData.RLock()
	calls = mock.calls.GetData
	lockIMyDataMockGetData.RUnlock()
	return calls
}
