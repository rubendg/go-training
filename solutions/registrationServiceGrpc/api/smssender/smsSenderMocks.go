// Code generated by MockGen. DO NOT EDIT.
// Source: smsSender.go

// Package smssender is a generated GoMock package.
package smssender

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockSmsSender is a mock of SmsSender interface.
type MockSmsSender struct {
	ctrl     *gomock.Controller
	recorder *MockSmsSenderMockRecorder
}

// MockSmsSenderMockRecorder is the mock recorder for MockSmsSender.
type MockSmsSenderMockRecorder struct {
	mock *MockSmsSender
}

// NewMockSmsSender creates a new mock instance.
func NewMockSmsSender(ctrl *gomock.Controller) *MockSmsSender {
	mock := &MockSmsSender{ctrl: ctrl}
	mock.recorder = &MockSmsSenderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSmsSender) EXPECT() *MockSmsSenderMockRecorder {
	return m.recorder
}

// SendSms mocks base method.
func (m *MockSmsSender) SendSms(phoneNumber, smsContent string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendSms", phoneNumber, smsContent)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendSms indicates an expected call of SendSms.
func (mr *MockSmsSenderMockRecorder) SendSms(phoneNumber, smsContent interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendSms", reflect.TypeOf((*MockSmsSender)(nil).SendSms), phoneNumber, smsContent)
}
