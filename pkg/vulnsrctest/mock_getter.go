// Code generated by mockery v1.0.0. DO NOT EDIT.

package vulnsrctest

import (
	types "github.com/khulnasoft-labs/vul-db/pkg/types"
	mock "github.com/stretchr/testify/mock"
)

// MockGetter is an autogenerated mock type for the Getter type
type MockGetter struct {
	mock.Mock
}

type GetterGetArgs struct {
	_a0         string
	_a0Anything bool
	_a1         string
	_a1Anything bool
}

type GetterGetReturns struct {
	_a0 []types.Advisory
	_a1 error
}

type GetterGetExpectation struct {
	Args    GetterGetArgs
	Returns GetterGetReturns
}

func (_m *MockGetter) ApplyGetExpectation(e GetterGetExpectation) {
	var args []interface{}
	if e.Args._a0Anything {
		args = append(args, mock.Anything)
	} else {
		args = append(args, e.Args._a0)
	}
	if e.Args._a1Anything {
		args = append(args, mock.Anything)
	} else {
		args = append(args, e.Args._a1)
	}
	_m.On("Get", args...).Return(e.Returns._a0, e.Returns._a1)
}

func (_m *MockGetter) ApplyGetExpectations(expectations []GetterGetExpectation) {
	for _, e := range expectations {
		_m.ApplyGetExpectation(e)
	}
}

// Get provides a mock function with given fields: _a0, _a1
func (_m *MockGetter) Get(_a0 string, _a1 string) ([]types.Advisory, error) {
	ret := _m.Called(_a0, _a1)

	var r0 []types.Advisory
	if rf, ok := ret.Get(0).(func(string, string) []types.Advisory); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.Advisory)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
