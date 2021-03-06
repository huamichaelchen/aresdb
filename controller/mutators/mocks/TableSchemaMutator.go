//  Copyright (c) 2017-2018 Uber Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// Code generated by mockery v1.0.0
package mocks

import "github.com/uber/aresdb/metastore/common"
import "github.com/stretchr/testify/mock"

// TableSchemaMutator is an autogenerated mock type for the TableSchemaMutator type
type TableSchemaMutator struct {
	mock.Mock
}

// CreateTable provides a mock function with given fields: namespace, table, force
func (_m *TableSchemaMutator) CreateTable(namespace string, table *common.Table, force bool) error {
	ret := _m.Called(namespace, table, force)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, *common.Table, bool) error); ok {
		r0 = rf(namespace, table, force)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteTable provides a mock function with given fields: namespace, name
func (_m *TableSchemaMutator) DeleteTable(namespace string, name string) error {
	ret := _m.Called(namespace, name)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(namespace, name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetHash provides a mock function with given fields: namespace
func (_m *TableSchemaMutator) GetHash(namespace string) (string, error) {
	ret := _m.Called(namespace)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(namespace)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(namespace)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTable provides a mock function with given fields: namespace, name
func (_m *TableSchemaMutator) GetTable(namespace string, name string) (*common.Table, error) {
	ret := _m.Called(namespace, name)

	var r0 *common.Table
	if rf, ok := ret.Get(0).(func(string, string) *common.Table); ok {
		r0 = rf(namespace, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*common.Table)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(namespace, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListTables provides a mock function with given fields: namespace
func (_m *TableSchemaMutator) ListTables(namespace string) ([]string, error) {
	ret := _m.Called(namespace)

	var r0 []string
	if rf, ok := ret.Get(0).(func(string) []string); ok {
		r0 = rf(namespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(namespace)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateTable provides a mock function with given fields: namespace, table, force
func (_m *TableSchemaMutator) UpdateTable(namespace string, table common.Table, force bool) error {
	ret := _m.Called(namespace, table, force)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, common.Table, bool) error); ok {
		r0 = rf(namespace, table, force)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
