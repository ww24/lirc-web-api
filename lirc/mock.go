// Automatically generated by MockGen. DO NOT EDIT!
// Source: lirc/interface.go

package lirc

import (
	gomock "github.com/golang/mock/gomock"
)

// Mock of ClientAPI interface
type MockClientAPI struct {
	ctrl     *gomock.Controller
	recorder *_MockClientAPIRecorder
}

// Recorder for MockClientAPI (not exported)
type _MockClientAPIRecorder struct {
	mock *MockClientAPI
}

func NewMockClientAPI(ctrl *gomock.Controller) *MockClientAPI {
	mock := &MockClientAPI{ctrl: ctrl}
	mock.recorder = &_MockClientAPIRecorder{mock}
	return mock
}

func (_m *MockClientAPI) EXPECT() *_MockClientAPIRecorder {
	return _m.recorder
}

func (_m *MockClientAPI) Version() (string, error) {
	ret := _m.ctrl.Call(_m, "Version")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockClientAPIRecorder) Version() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Version")
}

func (_m *MockClientAPI) List(remote string, code ...string) ([]string, error) {
	_s := []interface{}{remote}
	for _, _x := range code {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "List", _s...)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockClientAPIRecorder) List(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0}, arg1...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "List", _s...)
}

func (_m *MockClientAPI) SendOnce(remote string, code ...string) error {
	_s := []interface{}{remote}
	for _, _x := range code {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "SendOnce", _s...)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockClientAPIRecorder) SendOnce(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0}, arg1...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SendOnce", _s...)
}

func (_m *MockClientAPI) SendStart(remote string, code string) error {
	ret := _m.ctrl.Call(_m, "SendStart", remote, code)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockClientAPIRecorder) SendStart(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SendStart", arg0, arg1)
}

func (_m *MockClientAPI) SendStop(remote string, code string) error {
	ret := _m.ctrl.Call(_m, "SendStop", remote, code)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockClientAPIRecorder) SendStop(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SendStop", arg0, arg1)
}

func (_m *MockClientAPI) Close() error {
	ret := _m.ctrl.Call(_m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockClientAPIRecorder) Close() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Close")
}

func (_m *MockClientAPI) send(cmd string, remote string, code ...string) ([]Reply, error) {
	_s := []interface{}{cmd, remote}
	for _, _x := range code {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "send", _s...)
	ret0, _ := ret[0].([]Reply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockClientAPIRecorder) send(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "send", _s...)
}

func (_m *MockClientAPI) read() {
	_m.ctrl.Call(_m, "read")
}

func (_mr *_MockClientAPIRecorder) read() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "read")
}
