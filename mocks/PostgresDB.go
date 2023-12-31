// Code generated by mockery v2.32.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	sql "database/sql"
)

// PostgresDB is an autogenerated mock type for the PostgresDB type
type PostgresDB struct {
	mock.Mock
}

// DeleteTodo provides a mock function with given fields: _a0
func (_m *PostgresDB) DeleteTodo(_a0 string) (int64, error) {
	ret := _m.Called(_a0)

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (int64, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(string) int64); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertTodo provides a mock function with given fields: _a0, _a1
func (_m *PostgresDB) InsertTodo(_a0 string, _a1 string) (int64, error) {
	ret := _m.Called(_a0, _a1)

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (int64, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(string, string) int64); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SearchTodoById provides a mock function with given fields: _a0
func (_m *PostgresDB) SearchTodoById(_a0 int64) *sql.Row {
	ret := _m.Called(_a0)

	var r0 *sql.Row
	if rf, ok := ret.Get(0).(func(int64) *sql.Row); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sql.Row)
		}
	}

	return r0
}

// SearchTodoByTitle provides a mock function with given fields: _a0
func (_m *PostgresDB) SearchTodoByTitle(_a0 string) (*sql.Rows, error) {
	ret := _m.Called(_a0)

	var r0 *sql.Rows
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*sql.Rows, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(string) *sql.Rows); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sql.Rows)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewPostgresDB creates a new instance of PostgresDB. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPostgresDB(t interface {
	mock.TestingT
	Cleanup(func())
}) *PostgresDB {
	mock := &PostgresDB{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
