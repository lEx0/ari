package arimocks

import mock "github.com/stretchr/testify/mock"

// MessageRawer is an autogenerated mock type for the MessageRawer type
type MessageRawer struct {
	mock.Mock
}

// GetRaw provides a mock function with given fields:
func (_m *MessageRawer) GetRaw() []byte {
	ret := _m.Called()

	var r0 []byte
	if rf, ok := ret.Get(0).(func() []byte); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	return r0
}

// SetRaw provides a mock function with given fields: _a0
func (_m *MessageRawer) SetRaw(_a0 []byte) {
	_m.Called(_a0)
}
