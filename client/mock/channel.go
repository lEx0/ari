// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/CyCoreSystems/ari (interfaces: Channel)

package mock

import (
	ari "github.com/CyCoreSystems/ari"
	gomock "github.com/golang/mock/gomock"
	time "time"
)

// Mock of Channel interface
type MockChannel struct {
	ctrl     *gomock.Controller
	recorder *_MockChannelRecorder
}

// Recorder for MockChannel (not exported)
type _MockChannelRecorder struct {
	mock *MockChannel
}

func NewMockChannel(ctrl *gomock.Controller) *MockChannel {
	mock := &MockChannel{ctrl: ctrl}
	mock.recorder = &_MockChannelRecorder{mock}
	return mock
}

func (_m *MockChannel) EXPECT() *_MockChannelRecorder {
	return _m.recorder
}

func (_m *MockChannel) Answer(_param0 string) error {
	ret := _m.ctrl.Call(_m, "Answer", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockChannelRecorder) Answer(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Answer", arg0)
}

func (_m *MockChannel) Busy(_param0 string) error {
	ret := _m.ctrl.Call(_m, "Busy", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockChannelRecorder) Busy(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Busy", arg0)
}

func (_m *MockChannel) Congestion(_param0 string) error {
	ret := _m.ctrl.Call(_m, "Congestion", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockChannelRecorder) Congestion(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Congestion", arg0)
}

func (_m *MockChannel) Continue(_param0 string, _param1 string, _param2 string, _param3 int) error {
	ret := _m.ctrl.Call(_m, "Continue", _param0, _param1, _param2, _param3)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockChannelRecorder) Continue(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Continue", arg0, arg1, arg2, arg3)
}

func (_m *MockChannel) Create(_param0 ari.ChannelCreateRequest) (ari.ChannelHandle, error) {
	ret := _m.ctrl.Call(_m, "Create", _param0)
	ret0, _ := ret[0].(ari.ChannelHandle)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockChannelRecorder) Create(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Create", arg0)
}

func (_m *MockChannel) Data(_param0 string) (*ari.ChannelData, error) {
	ret := _m.ctrl.Call(_m, "Data", _param0)
	ret0, _ := ret[0].(*ari.ChannelData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockChannelRecorder) Data(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Data", arg0)
}

func (_m *MockChannel) Dial(_param0 string, _param1 string, _param2 time.Duration) error {
	ret := _m.ctrl.Call(_m, "Dial", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockChannelRecorder) Dial(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Dial", arg0, arg1, arg2)
}

func (_m *MockChannel) Get(_param0 string) ari.ChannelHandle {
	ret := _m.ctrl.Call(_m, "Get", _param0)
	ret0, _ := ret[0].(ari.ChannelHandle)
	return ret0
}

func (_mr *_MockChannelRecorder) Get(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Get", arg0)
}

func (_m *MockChannel) Hangup(_param0 string, _param1 string) error {
	ret := _m.ctrl.Call(_m, "Hangup", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockChannelRecorder) Hangup(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Hangup", arg0, arg1)
}

func (_m *MockChannel) Hold(_param0 string) error {
	ret := _m.ctrl.Call(_m, "Hold", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockChannelRecorder) Hold(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Hold", arg0)
}

func (_m *MockChannel) List() ([]ari.ChannelHandle, error) {
	ret := _m.ctrl.Call(_m, "List")
	ret0, _ := ret[0].([]ari.ChannelHandle)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockChannelRecorder) List() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "List")
}

func (_m *MockChannel) MOH(_param0 string, _param1 string) error {
	ret := _m.ctrl.Call(_m, "MOH", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockChannelRecorder) MOH(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "MOH", arg0, arg1)
}

func (_m *MockChannel) Mute(_param0 string, _param1 ari.Direction) error {
	ret := _m.ctrl.Call(_m, "Mute", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockChannelRecorder) Mute(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Mute", arg0, arg1)
}

func (_m *MockChannel) Originate(_param0 ari.OriginateRequest) (ari.ChannelHandle, error) {
	ret := _m.ctrl.Call(_m, "Originate", _param0)
	ret0, _ := ret[0].(ari.ChannelHandle)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockChannelRecorder) Originate(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Originate", arg0)
}

func (_m *MockChannel) Play(_param0 string, _param1 string, _param2 string) (ari.PlaybackHandle, error) {
	ret := _m.ctrl.Call(_m, "Play", _param0, _param1, _param2)
	ret0, _ := ret[0].(ari.PlaybackHandle)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockChannelRecorder) Play(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Play", arg0, arg1, arg2)
}

func (_m *MockChannel) Record(_param0 string, _param1 string, _param2 *ari.RecordingOptions) (ari.LiveRecordingHandle, error) {
	ret := _m.ctrl.Call(_m, "Record", _param0, _param1, _param2)
	ret0, _ := ret[0].(ari.LiveRecordingHandle)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockChannelRecorder) Record(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Record", arg0, arg1, arg2)
}

func (_m *MockChannel) Ring(_param0 string) error {
	ret := _m.ctrl.Call(_m, "Ring", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockChannelRecorder) Ring(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Ring", arg0)
}

func (_m *MockChannel) SendDTMF(_param0 string, _param1 string, _param2 *ari.DTMFOptions) error {
	ret := _m.ctrl.Call(_m, "SendDTMF", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockChannelRecorder) SendDTMF(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SendDTMF", arg0, arg1, arg2)
}

func (_m *MockChannel) Silence(_param0 string) error {
	ret := _m.ctrl.Call(_m, "Silence", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockChannelRecorder) Silence(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Silence", arg0)
}

func (_m *MockChannel) Snoop(_param0 string, _param1 string, _param2 *ari.SnoopOptions) (ari.ChannelHandle, error) {
	ret := _m.ctrl.Call(_m, "Snoop", _param0, _param1, _param2)
	ret0, _ := ret[0].(ari.ChannelHandle)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockChannelRecorder) Snoop(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Snoop", arg0, arg1, arg2)
}

func (_m *MockChannel) StopHold(_param0 string) error {
	ret := _m.ctrl.Call(_m, "StopHold", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockChannelRecorder) StopHold(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "StopHold", arg0)
}

func (_m *MockChannel) StopMOH(_param0 string) error {
	ret := _m.ctrl.Call(_m, "StopMOH", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockChannelRecorder) StopMOH(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "StopMOH", arg0)
}

func (_m *MockChannel) StopRing(_param0 string) error {
	ret := _m.ctrl.Call(_m, "StopRing", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockChannelRecorder) StopRing(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "StopRing", arg0)
}

func (_m *MockChannel) StopSilence(_param0 string) error {
	ret := _m.ctrl.Call(_m, "StopSilence", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockChannelRecorder) StopSilence(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "StopSilence", arg0)
}

func (_m *MockChannel) Subscribe(_param0 string, _param1 ...string) ari.Subscription {
	_s := []interface{}{_param0}
	for _, _x := range _param1 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "Subscribe", _s...)
	ret0, _ := ret[0].(ari.Subscription)
	return ret0
}

func (_mr *_MockChannelRecorder) Subscribe(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0}, arg1...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Subscribe", _s...)
}

func (_m *MockChannel) Unmute(_param0 string, _param1 ari.Direction) error {
	ret := _m.ctrl.Call(_m, "Unmute", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockChannelRecorder) Unmute(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Unmute", arg0, arg1)
}

func (_m *MockChannel) Variables(_param0 string) ari.Variables {
	ret := _m.ctrl.Call(_m, "Variables", _param0)
	ret0, _ := ret[0].(ari.Variables)
	return ret0
}

func (_mr *_MockChannelRecorder) Variables(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Variables", arg0)
}
