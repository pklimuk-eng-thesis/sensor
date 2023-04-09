// Code generated by mockery v2.23.2. DO NOT EDIT.

package service

import (
	domain "github.com/pklimuk-eng-thesis/sensor/pkg/domain"
	mock "github.com/stretchr/testify/mock"
)

// MockSensorService is an autogenerated mock type for the SensorService type
type MockSensorService struct {
	mock.Mock
}

type MockSensorService_Expecter struct {
	mock *mock.Mock
}

func (_m *MockSensorService) EXPECT() *MockSensorService_Expecter {
	return &MockSensorService_Expecter{mock: &_m.Mock}
}

// GetInfo provides a mock function with given fields:
func (_m *MockSensorService) GetInfo() *domain.Sensor {
	ret := _m.Called()

	var r0 *domain.Sensor
	if rf, ok := ret.Get(0).(func() *domain.Sensor); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Sensor)
		}
	}

	return r0
}

// MockSensorService_GetInfo_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetInfo'
type MockSensorService_GetInfo_Call struct {
	*mock.Call
}

// GetInfo is a helper method to define mock.On call
func (_e *MockSensorService_Expecter) GetInfo() *MockSensorService_GetInfo_Call {
	return &MockSensorService_GetInfo_Call{Call: _e.mock.On("GetInfo")}
}

func (_c *MockSensorService_GetInfo_Call) Run(run func()) *MockSensorService_GetInfo_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockSensorService_GetInfo_Call) Return(_a0 *domain.Sensor) *MockSensorService_GetInfo_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSensorService_GetInfo_Call) RunAndReturn(run func() *domain.Sensor) *MockSensorService_GetInfo_Call {
	_c.Call.Return(run)
	return _c
}

// ToggleDetected provides a mock function with given fields:
func (_m *MockSensorService) ToggleDetected() (*domain.Sensor, error) {
	ret := _m.Called()

	var r0 *domain.Sensor
	var r1 error
	if rf, ok := ret.Get(0).(func() (*domain.Sensor, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() *domain.Sensor); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Sensor)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockSensorService_ToggleDetected_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ToggleDetected'
type MockSensorService_ToggleDetected_Call struct {
	*mock.Call
}

// ToggleDetected is a helper method to define mock.On call
func (_e *MockSensorService_Expecter) ToggleDetected() *MockSensorService_ToggleDetected_Call {
	return &MockSensorService_ToggleDetected_Call{Call: _e.mock.On("ToggleDetected")}
}

func (_c *MockSensorService_ToggleDetected_Call) Run(run func()) *MockSensorService_ToggleDetected_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockSensorService_ToggleDetected_Call) Return(_a0 *domain.Sensor, _a1 error) *MockSensorService_ToggleDetected_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockSensorService_ToggleDetected_Call) RunAndReturn(run func() (*domain.Sensor, error)) *MockSensorService_ToggleDetected_Call {
	_c.Call.Return(run)
	return _c
}

// ToggleSensorEnabled provides a mock function with given fields:
func (_m *MockSensorService) ToggleSensorEnabled() *domain.Sensor {
	ret := _m.Called()

	var r0 *domain.Sensor
	if rf, ok := ret.Get(0).(func() *domain.Sensor); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Sensor)
		}
	}

	return r0
}

// MockSensorService_ToggleSensorEnabled_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ToggleSensorEnabled'
type MockSensorService_ToggleSensorEnabled_Call struct {
	*mock.Call
}

// ToggleSensorEnabled is a helper method to define mock.On call
func (_e *MockSensorService_Expecter) ToggleSensorEnabled() *MockSensorService_ToggleSensorEnabled_Call {
	return &MockSensorService_ToggleSensorEnabled_Call{Call: _e.mock.On("ToggleSensorEnabled")}
}

func (_c *MockSensorService_ToggleSensorEnabled_Call) Run(run func()) *MockSensorService_ToggleSensorEnabled_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockSensorService_ToggleSensorEnabled_Call) Return(_a0 *domain.Sensor) *MockSensorService_ToggleSensorEnabled_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSensorService_ToggleSensorEnabled_Call) RunAndReturn(run func() *domain.Sensor) *MockSensorService_ToggleSensorEnabled_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewMockSensorService interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockSensorService creates a new instance of MockSensorService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockSensorService(t mockConstructorTestingTNewMockSensorService) *MockSensorService {
	mock := &MockSensorService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
