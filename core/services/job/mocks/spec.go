// Code generated by mockery v2.3.0. DO NOT EDIT.

package mocks

import (
	job "github.com/smartcontractkit/chainlink/core/services/job"
	mock "github.com/stretchr/testify/mock"

	pipeline "github.com/smartcontractkit/chainlink/core/services/pipeline"
)

// Spec is an autogenerated mock type for the Spec type
type Spec struct {
	mock.Mock
}

// JobID provides a mock function with given fields:
func (_m *Spec) JobID() int32 {
	ret := _m.Called()

	var r0 int32
	if rf, ok := ret.Get(0).(func() int32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int32)
	}

	return r0
}

// JobType provides a mock function with given fields:
func (_m *Spec) JobType() job.Type {
	ret := _m.Called()

	var r0 job.Type
	if rf, ok := ret.Get(0).(func() job.Type); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(job.Type)
	}

	return r0
}

// TableName provides a mock function with given fields:
func (_m *Spec) TableName() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// TaskDAG provides a mock function with given fields:
func (_m *Spec) TaskDAG() pipeline.TaskDAG {
	ret := _m.Called()

	var r0 pipeline.TaskDAG
	if rf, ok := ret.Get(0).(func() pipeline.TaskDAG); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(pipeline.TaskDAG)
	}

	return r0
}
