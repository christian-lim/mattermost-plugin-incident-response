// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mattermost/mattermost-plugin-incident-response/server/pluginkvstore (interfaces: UserAPI)

// Package mock_pluginkvstore is a generated GoMock package.
package mock_pluginkvstore

import (
	gomock "github.com/golang/mock/gomock"
	model "github.com/mattermost/mattermost-server/v5/model"
	reflect "reflect"
)

// MockUserAPI is a mock of UserAPI interface
type MockUserAPI struct {
	ctrl     *gomock.Controller
	recorder *MockUserAPIMockRecorder
}

// MockUserAPIMockRecorder is the mock recorder for MockUserAPI
type MockUserAPIMockRecorder struct {
	mock *MockUserAPI
}

// NewMockUserAPI creates a new mock instance
func NewMockUserAPI(ctrl *gomock.Controller) *MockUserAPI {
	mock := &MockUserAPI{ctrl: ctrl}
	mock.recorder = &MockUserAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUserAPI) EXPECT() *MockUserAPIMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockUserAPI) Get(arg0 string) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockUserAPIMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockUserAPI)(nil).Get), arg0)
}