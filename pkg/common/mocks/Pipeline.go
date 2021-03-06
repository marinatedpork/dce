// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	codepipeline "github.com/aws/aws-sdk-go/service/codepipeline"

	mock "github.com/stretchr/testify/mock"
)

// Pipeline is an autogenerated mock type for the Pipeline type
type Pipeline struct {
	mock.Mock
}

// StartPipeline provides a mock function with given fields: _a0
func (_m *Pipeline) StartPipeline(_a0 *codepipeline.StartPipelineExecutionInput) (*codepipeline.StartPipelineExecutionOutput, error) {
	ret := _m.Called(_a0)

	var r0 *codepipeline.StartPipelineExecutionOutput
	if rf, ok := ret.Get(0).(func(*codepipeline.StartPipelineExecutionInput) *codepipeline.StartPipelineExecutionOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*codepipeline.StartPipelineExecutionOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*codepipeline.StartPipelineExecutionInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
