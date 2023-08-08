// Code generated by mockery v1.0.0. DO NOT EDIT.

package vulnsrctest

import mock "github.com/stretchr/testify/mock"

// MockUpdater is an autogenerated mock type for the Updater type
type MockUpdater struct {
	mock.Mock
}

type UpdaterUpdateArgs struct {
	Dir         string
	DirAnything bool
}

type UpdaterUpdateReturns struct {
	Err error
}

type UpdaterUpdateExpectation struct {
	Args    UpdaterUpdateArgs
	Returns UpdaterUpdateReturns
}

func (_m *MockUpdater) ApplyUpdateExpectation(e UpdaterUpdateExpectation) {
	var args []interface{}
	if e.Args.DirAnything {
		args = append(args, mock.Anything)
	} else {
		args = append(args, e.Args.Dir)
	}
	_m.On("Update", args...).Return(e.Returns.Err)
}

func (_m *MockUpdater) ApplyUpdateExpectations(expectations []UpdaterUpdateExpectation) {
	for _, e := range expectations {
		_m.ApplyUpdateExpectation(e)
	}
}

// Update provides a mock function with given fields: dir
func (_m *MockUpdater) Update(dir string) error {
	ret := _m.Called(dir)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(dir)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}