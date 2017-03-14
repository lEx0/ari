// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/CyCoreSystems/ari (interfaces: DeviceState)

package mock

import (
	ari "github.com/CyCoreSystems/ari"
	gomock "github.com/golang/mock/gomock"
)

// Mock of DeviceState interface
type MockDeviceState struct {
	ctrl     *gomock.Controller
	recorder *_MockDeviceStateRecorder
}

// Recorder for MockDeviceState (not exported)
type _MockDeviceStateRecorder struct {
	mock *MockDeviceState
}

func NewMockDeviceState(ctrl *gomock.Controller) *MockDeviceState {
	mock := &MockDeviceState{ctrl: ctrl}
	mock.recorder = &_MockDeviceStateRecorder{mock}
	return mock
}

func (_m *MockDeviceState) EXPECT() *_MockDeviceStateRecorder {
	return _m.recorder
}

func (_m *MockDeviceState) Data(_param0 string) (*ari.DeviceStateData, error) {
	ret := _m.ctrl.Call(_m, "Data", _param0)
	ret0, _ := ret[0].(*ari.DeviceStateData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockDeviceStateRecorder) Data(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Data", arg0)
}

func (_m *MockDeviceState) Delete(_param0 string) error {
	ret := _m.ctrl.Call(_m, "Delete", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockDeviceStateRecorder) Delete(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Delete", arg0)
}

func (_m *MockDeviceState) Get(_param0 string) ari.DeviceStateHandle {
	ret := _m.ctrl.Call(_m, "Get", _param0)
	ret0, _ := ret[0].(ari.DeviceStateHandle)
	return ret0
}

func (_mr *_MockDeviceStateRecorder) Get(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Get", arg0)
}

func (_m *MockDeviceState) List() ([]ari.DeviceStateHandle, error) {
	ret := _m.ctrl.Call(_m, "List")
	ret0, _ := ret[0].([]ari.DeviceStateHandle)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockDeviceStateRecorder) List() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "List")
}

func (_m *MockDeviceState) Update(_param0 string, _param1 string) error {
	ret := _m.ctrl.Call(_m, "Update", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockDeviceStateRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Update", arg0, arg1)
}
