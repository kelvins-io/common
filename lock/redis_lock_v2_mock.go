// Code generated by MockGen. DO NOT EDIT.
// Source: redis_lock_v2.go

// Package lock is a generated GoMock package.
package lock

import (
	reflect "reflect"

	redis "github.com/garyburd/redigo/redis"
	gomock "github.com/golang/mock/gomock"
)

// MockRedisLockIface is a mock of RedisLockIface interface
type MockRedisLockIface struct {
	ctrl     *gomock.Controller
	recorder *MockRedisLockIfaceMockRecorder
}

// MockRedisLockIfaceMockRecorder is the mock recorder for MockRedisLockIface
type MockRedisLockIfaceMockRecorder struct {
	mock *MockRedisLockIface
}

// NewMockRedisLockIface creates a new mock instance
func NewMockRedisLockIface(ctrl *gomock.Controller) *MockRedisLockIface {
	mock := &MockRedisLockIface{ctrl: ctrl}
	mock.recorder = &MockRedisLockIfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRedisLockIface) EXPECT() *MockRedisLockIfaceMockRecorder {
	return m.recorder
}

// Set mocks base method
func (m *MockRedisLockIface) Set(arg0 *redis.Pool, arg1 string, arg2 uint32) (bool, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", arg0, arg1, arg2)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Set indicates an expected call of Set
func (mr *MockRedisLockIfaceMockRecorder) Set(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockRedisLockIface)(nil).Set), arg0, arg1, arg2)
}

// Release mocks base method
func (m *MockRedisLockIface) Release(arg0 *redis.Pool, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Release", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Release indicates an expected call of Release
func (mr *MockRedisLockIfaceMockRecorder) Release(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Release", reflect.TypeOf((*MockRedisLockIface)(nil).Release), arg0, arg1, arg2)
}