// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/dmsclient/dmsclient.go

// Package mock_dmsclient is a generated GoMock package.
package mock_dmsclient

import (
	gomock "github.com/golang/mock/gomock"
	dmsclient "github.com/openshift/deadmanssnitch-operator/pkg/dmsclient"
	reflect "reflect"
)

// MockClient is a mock of Client interface
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// ListAll mocks base method
func (m *MockClient) ListAll() ([]dmsclient.Snitch, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAll")
	ret0, _ := ret[0].([]dmsclient.Snitch)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAll indicates an expected call of ListAll
func (mr *MockClientMockRecorder) ListAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAll", reflect.TypeOf((*MockClient)(nil).ListAll))
}

// List mocks base method
func (m *MockClient) List(snitchToken string) (dmsclient.Snitch, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", snitchToken)
	ret0, _ := ret[0].(dmsclient.Snitch)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockClientMockRecorder) List(snitchToken interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockClient)(nil).List), snitchToken)
}

// Create mocks base method
func (m *MockClient) Create(newSnitch dmsclient.Snitch) (dmsclient.Snitch, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", newSnitch)
	ret0, _ := ret[0].(dmsclient.Snitch)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockClientMockRecorder) Create(newSnitch interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockClient)(nil).Create), newSnitch)
}

// Delete mocks base method
func (m *MockClient) Delete(snitchToken string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", snitchToken)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete
func (mr *MockClientMockRecorder) Delete(snitchToken interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockClient)(nil).Delete), snitchToken)
}

// FindSnitchesByName mocks base method
func (m *MockClient) FindSnitchesByName(snitchName string) ([]dmsclient.Snitch, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindSnitchesByName", snitchName)
	ret0, _ := ret[0].([]dmsclient.Snitch)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindSnitchesByName indicates an expected call of FindSnitchesByName
func (mr *MockClientMockRecorder) FindSnitchesByName(snitchName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindSnitchesByName", reflect.TypeOf((*MockClient)(nil).FindSnitchesByName), snitchName)
}

// Update mocks base method
func (m *MockClient) Update(updateSnitch dmsclient.Snitch) (dmsclient.Snitch, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", updateSnitch)
	ret0, _ := ret[0].(dmsclient.Snitch)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update
func (mr *MockClientMockRecorder) Update(updateSnitch interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockClient)(nil).Update), updateSnitch)
}

// CheckIn mocks base method
func (m *MockClient) CheckIn(s dmsclient.Snitch) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckIn", s)
	ret0, _ := ret[0].(error)
	return ret0
}

// CheckIn indicates an expected call of CheckIn
func (mr *MockClientMockRecorder) CheckIn(s interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckIn", reflect.TypeOf((*MockClient)(nil).CheckIn), s)
}
