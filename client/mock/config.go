// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/CyCoreSystems/ari (interfaces: Config)

package mock

import (
	ari "github.com/CyCoreSystems/ari"
	gomock "github.com/golang/mock/gomock"
)

// Mock of Config interface
type MockConfig struct {
	ctrl     *gomock.Controller
	recorder *_MockConfigRecorder
}

// Recorder for MockConfig (not exported)
type _MockConfigRecorder struct {
	mock *MockConfig
}

func NewMockConfig(ctrl *gomock.Controller) *MockConfig {
	mock := &MockConfig{ctrl: ctrl}
	mock.recorder = &_MockConfigRecorder{mock}
	return mock
}

func (_m *MockConfig) EXPECT() *_MockConfigRecorder {
	return _m.recorder
}

func (_m *MockConfig) Data(_param0 string, _param1 string, _param2 string) (*ari.ConfigData, error) {
	ret := _m.ctrl.Call(_m, "Data", _param0, _param1, _param2)
	ret0, _ := ret[0].(*ari.ConfigData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockConfigRecorder) Data(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Data", arg0, arg1, arg2)
}

func (_m *MockConfig) Delete(_param0 string, _param1 string, _param2 string) error {
	ret := _m.ctrl.Call(_m, "Delete", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockConfigRecorder) Delete(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Delete", arg0, arg1, arg2)
}

func (_m *MockConfig) Get(_param0 string, _param1 string, _param2 string) ari.ConfigHandle {
	ret := _m.ctrl.Call(_m, "Get", _param0, _param1, _param2)
	ret0, _ := ret[0].(ari.ConfigHandle)
	return ret0
}

func (_mr *_MockConfigRecorder) Get(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Get", arg0, arg1, arg2)
}

func (_m *MockConfig) Update(_param0 string, _param1 string, _param2 string, _param3 []ari.ConfigTuple) error {
	ret := _m.ctrl.Call(_m, "Update", _param0, _param1, _param2, _param3)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockConfigRecorder) Update(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Update", arg0, arg1, arg2, arg3)
}
