// Copyright 2015-2018 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/aws/amazon-ecs-agent/agent/engine/dockerstate (interfaces: TaskEngineState)

package mock_dockerstate

import (
	container "github.com/aws/amazon-ecs-agent/agent/api/container"
	eni "github.com/aws/amazon-ecs-agent/agent/api/eni"
	task "github.com/aws/amazon-ecs-agent/agent/api/task"
	image "github.com/aws/amazon-ecs-agent/agent/engine/image"
	gomock "github.com/golang/mock/gomock"
)

// Mock of TaskEngineState interface
type MockTaskEngineState struct {
	ctrl     *gomock.Controller
	recorder *_MockTaskEngineStateRecorder
}

// Recorder for MockTaskEngineState (not exported)
type _MockTaskEngineStateRecorder struct {
	mock *MockTaskEngineState
}

func NewMockTaskEngineState(ctrl *gomock.Controller) *MockTaskEngineState {
	mock := &MockTaskEngineState{ctrl: ctrl}
	mock.recorder = &_MockTaskEngineStateRecorder{mock}
	return mock
}

func (_m *MockTaskEngineState) EXPECT() *_MockTaskEngineStateRecorder {
	return _m.recorder
}

func (_m *MockTaskEngineState) AddContainer(_param0 *container.DockerContainer, _param1 *task.Task) {
	_m.ctrl.Call(_m, "AddContainer", _param0, _param1)
}

func (_mr *_MockTaskEngineStateRecorder) AddContainer(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AddContainer", arg0, arg1)
}

func (_m *MockTaskEngineState) AddENIAttachment(_param0 *eni.ENIAttachment) {
	_m.ctrl.Call(_m, "AddENIAttachment", _param0)
}

func (_mr *_MockTaskEngineStateRecorder) AddENIAttachment(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AddENIAttachment", arg0)
}

func (_m *MockTaskEngineState) AddImageState(_param0 *image.ImageState) {
	_m.ctrl.Call(_m, "AddImageState", _param0)
}

func (_mr *_MockTaskEngineStateRecorder) AddImageState(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AddImageState", arg0)
}

func (_m *MockTaskEngineState) AddTask(_param0 *task.Task) {
	_m.ctrl.Call(_m, "AddTask", _param0)
}

func (_mr *_MockTaskEngineStateRecorder) AddTask(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AddTask", arg0)
}

func (_m *MockTaskEngineState) AddTaskIPAddress(_param0 string, _param1 string) {
	_m.ctrl.Call(_m, "AddTaskIPAddress", _param0, _param1)
}

func (_mr *_MockTaskEngineStateRecorder) AddTaskIPAddress(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AddTaskIPAddress", arg0, arg1)
}

func (_m *MockTaskEngineState) AllImageStates() []*image.ImageState {
	ret := _m.ctrl.Call(_m, "AllImageStates")
	ret0, _ := ret[0].([]*image.ImageState)
	return ret0
}

func (_mr *_MockTaskEngineStateRecorder) AllImageStates() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AllImageStates")
}

func (_m *MockTaskEngineState) AllTasks() []*task.Task {
	ret := _m.ctrl.Call(_m, "AllTasks")
	ret0, _ := ret[0].([]*task.Task)
	return ret0
}

func (_mr *_MockTaskEngineStateRecorder) AllTasks() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AllTasks")
}

func (_m *MockTaskEngineState) ContainerByID(_param0 string) (*container.DockerContainer, bool) {
	ret := _m.ctrl.Call(_m, "ContainerByID", _param0)
	ret0, _ := ret[0].(*container.DockerContainer)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

func (_mr *_MockTaskEngineStateRecorder) ContainerByID(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ContainerByID", arg0)
}

func (_m *MockTaskEngineState) ContainerByV3EndpointID(_param0 string) (*container.DockerContainer, bool) {
	ret := _m.ctrl.Call(_m, "ContainerByV3EndpointID", _param0)
	ret0, _ := ret[0].(*container.DockerContainer)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

func (_mr *_MockTaskEngineStateRecorder) ContainerByV3EndpointID(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ContainerByV3EndpointID", arg0)
}

func (_m *MockTaskEngineState) ContainerMapByArn(_param0 string) (map[string]*container.DockerContainer, bool) {
	ret := _m.ctrl.Call(_m, "ContainerMapByArn", _param0)
	ret0, _ := ret[0].(map[string]*container.DockerContainer)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

func (_mr *_MockTaskEngineStateRecorder) ContainerMapByArn(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ContainerMapByArn", arg0)
}

func (_m *MockTaskEngineState) ENIByMac(_param0 string) (*eni.ENIAttachment, bool) {
	ret := _m.ctrl.Call(_m, "ENIByMac", _param0)
	ret0, _ := ret[0].(*eni.ENIAttachment)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

func (_mr *_MockTaskEngineStateRecorder) ENIByMac(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ENIByMac", arg0)
}

func (_m *MockTaskEngineState) GetAllContainerIDs() []string {
	ret := _m.ctrl.Call(_m, "GetAllContainerIDs")
	ret0, _ := ret[0].([]string)
	return ret0
}

func (_mr *_MockTaskEngineStateRecorder) GetAllContainerIDs() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetAllContainerIDs")
}

func (_m *MockTaskEngineState) GetTaskByIPAddress(_param0 string) (string, bool) {
	ret := _m.ctrl.Call(_m, "GetTaskByIPAddress", _param0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

func (_mr *_MockTaskEngineStateRecorder) GetTaskByIPAddress(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetTaskByIPAddress", arg0)
}

func (_m *MockTaskEngineState) MarshalJSON() ([]byte, error) {
	ret := _m.ctrl.Call(_m, "MarshalJSON")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockTaskEngineStateRecorder) MarshalJSON() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "MarshalJSON")
}

func (_m *MockTaskEngineState) RemoveENIAttachment(_param0 string) {
	_m.ctrl.Call(_m, "RemoveENIAttachment", _param0)
}

func (_mr *_MockTaskEngineStateRecorder) RemoveENIAttachment(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RemoveENIAttachment", arg0)
}

func (_m *MockTaskEngineState) RemoveImageState(_param0 *image.ImageState) {
	_m.ctrl.Call(_m, "RemoveImageState", _param0)
}

func (_mr *_MockTaskEngineStateRecorder) RemoveImageState(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RemoveImageState", arg0)
}

func (_m *MockTaskEngineState) RemoveTask(_param0 *task.Task) {
	_m.ctrl.Call(_m, "RemoveTask", _param0)
}

func (_mr *_MockTaskEngineStateRecorder) RemoveTask(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RemoveTask", arg0)
}

func (_m *MockTaskEngineState) Reset() {
	_m.ctrl.Call(_m, "Reset")
}

func (_mr *_MockTaskEngineStateRecorder) Reset() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Reset")
}

func (_m *MockTaskEngineState) TaskByArn(_param0 string) (*task.Task, bool) {
	ret := _m.ctrl.Call(_m, "TaskByArn", _param0)
	ret0, _ := ret[0].(*task.Task)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

func (_mr *_MockTaskEngineStateRecorder) TaskByArn(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "TaskByArn", arg0)
}

func (_m *MockTaskEngineState) TaskByID(_param0 string) (*task.Task, bool) {
	ret := _m.ctrl.Call(_m, "TaskByID", _param0)
	ret0, _ := ret[0].(*task.Task)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

func (_mr *_MockTaskEngineStateRecorder) TaskByID(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "TaskByID", arg0)
}

func (_m *MockTaskEngineState) TaskByShortID(_param0 string) ([]*task.Task, bool) {
	ret := _m.ctrl.Call(_m, "TaskByShortID", _param0)
	ret0, _ := ret[0].([]*task.Task)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

func (_mr *_MockTaskEngineStateRecorder) TaskByShortID(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "TaskByShortID", arg0)
}

func (_m *MockTaskEngineState) TaskByV3EndpointID(_param0 string) (*task.Task, bool) {
	ret := _m.ctrl.Call(_m, "TaskByV3EndpointID", _param0)
	ret0, _ := ret[0].(*task.Task)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

func (_mr *_MockTaskEngineStateRecorder) TaskByV3EndpointID(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "TaskByV3EndpointID", arg0)
}

func (_m *MockTaskEngineState) UnmarshalJSON(_param0 []byte) error {
	ret := _m.ctrl.Call(_m, "UnmarshalJSON", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockTaskEngineStateRecorder) UnmarshalJSON(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "UnmarshalJSON", arg0)
}
