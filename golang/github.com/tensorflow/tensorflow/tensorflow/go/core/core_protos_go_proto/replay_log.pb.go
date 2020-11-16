// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: tensorflow/core/protobuf/replay_log.proto

package core_protos_go_proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Records the creation of a new replay session.  We record the device listing
// here to capture the state of the cluster.
type NewReplaySession struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Devices       *ListDevicesResponse `protobuf:"bytes,1,opt,name=devices,proto3" json:"devices,omitempty"`
	SessionHandle string               `protobuf:"bytes,2,opt,name=session_handle,json=sessionHandle,proto3" json:"session_handle,omitempty"`
}

func (x *NewReplaySession) Reset() {
	*x = NewReplaySession{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_protobuf_replay_log_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewReplaySession) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewReplaySession) ProtoMessage() {}

func (x *NewReplaySession) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_protobuf_replay_log_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewReplaySession.ProtoReflect.Descriptor instead.
func (*NewReplaySession) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_protobuf_replay_log_proto_rawDescGZIP(), []int{0}
}

func (x *NewReplaySession) GetDevices() *ListDevicesResponse {
	if x != nil {
		return x.Devices
	}
	return nil
}

func (x *NewReplaySession) GetSessionHandle() string {
	if x != nil {
		return x.SessionHandle
	}
	return ""
}

type ReplayOp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartTimeUs float64 `protobuf:"fixed64,31,opt,name=start_time_us,json=startTimeUs,proto3" json:"start_time_us,omitempty"`
	EndTimeUs   float64 `protobuf:"fixed64,32,opt,name=end_time_us,json=endTimeUs,proto3" json:"end_time_us,omitempty"`
	// Types that are assignable to Op:
	//	*ReplayOp_CreateSession
	//	*ReplayOp_ExtendSession
	//	*ReplayOp_PartialRunSetup
	//	*ReplayOp_RunStep
	//	*ReplayOp_CloseSession
	//	*ReplayOp_ListDevices
	//	*ReplayOp_ResetRequest
	//	*ReplayOp_MakeCallable
	//	*ReplayOp_RunCallable
	//	*ReplayOp_ReleaseCallable
	//	*ReplayOp_NewReplaySession
	Op isReplayOp_Op `protobuf_oneof:"op"`
	// Types that are assignable to Response:
	//	*ReplayOp_CreateSessionResponse
	//	*ReplayOp_ExtendSessionResponse
	//	*ReplayOp_PartialRunSetupResponse
	//	*ReplayOp_RunStepResponse
	//	*ReplayOp_CloseSessionResponse
	//	*ReplayOp_ListDevicesResponse
	//	*ReplayOp_ResetRequestResponse
	//	*ReplayOp_MakeCallableResponse
	//	*ReplayOp_RunCallableResponse
	//	*ReplayOp_ReleaseCallableResponse
	Response isReplayOp_Response `protobuf_oneof:"response"`
}

func (x *ReplayOp) Reset() {
	*x = ReplayOp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_protobuf_replay_log_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReplayOp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplayOp) ProtoMessage() {}

func (x *ReplayOp) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_protobuf_replay_log_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplayOp.ProtoReflect.Descriptor instead.
func (*ReplayOp) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_protobuf_replay_log_proto_rawDescGZIP(), []int{1}
}

func (x *ReplayOp) GetStartTimeUs() float64 {
	if x != nil {
		return x.StartTimeUs
	}
	return 0
}

func (x *ReplayOp) GetEndTimeUs() float64 {
	if x != nil {
		return x.EndTimeUs
	}
	return 0
}

func (m *ReplayOp) GetOp() isReplayOp_Op {
	if m != nil {
		return m.Op
	}
	return nil
}

func (x *ReplayOp) GetCreateSession() *CreateSessionRequest {
	if x, ok := x.GetOp().(*ReplayOp_CreateSession); ok {
		return x.CreateSession
	}
	return nil
}

func (x *ReplayOp) GetExtendSession() *ExtendSessionRequest {
	if x, ok := x.GetOp().(*ReplayOp_ExtendSession); ok {
		return x.ExtendSession
	}
	return nil
}

func (x *ReplayOp) GetPartialRunSetup() *PartialRunSetupRequest {
	if x, ok := x.GetOp().(*ReplayOp_PartialRunSetup); ok {
		return x.PartialRunSetup
	}
	return nil
}

func (x *ReplayOp) GetRunStep() *RunStepRequest {
	if x, ok := x.GetOp().(*ReplayOp_RunStep); ok {
		return x.RunStep
	}
	return nil
}

func (x *ReplayOp) GetCloseSession() *CloseSessionRequest {
	if x, ok := x.GetOp().(*ReplayOp_CloseSession); ok {
		return x.CloseSession
	}
	return nil
}

func (x *ReplayOp) GetListDevices() *ListDevicesRequest {
	if x, ok := x.GetOp().(*ReplayOp_ListDevices); ok {
		return x.ListDevices
	}
	return nil
}

func (x *ReplayOp) GetResetRequest() *ResetRequest {
	if x, ok := x.GetOp().(*ReplayOp_ResetRequest); ok {
		return x.ResetRequest
	}
	return nil
}

func (x *ReplayOp) GetMakeCallable() *MakeCallableRequest {
	if x, ok := x.GetOp().(*ReplayOp_MakeCallable); ok {
		return x.MakeCallable
	}
	return nil
}

func (x *ReplayOp) GetRunCallable() *RunCallableRequest {
	if x, ok := x.GetOp().(*ReplayOp_RunCallable); ok {
		return x.RunCallable
	}
	return nil
}

func (x *ReplayOp) GetReleaseCallable() *ReleaseCallableRequest {
	if x, ok := x.GetOp().(*ReplayOp_ReleaseCallable); ok {
		return x.ReleaseCallable
	}
	return nil
}

func (x *ReplayOp) GetNewReplaySession() *NewReplaySession {
	if x, ok := x.GetOp().(*ReplayOp_NewReplaySession); ok {
		return x.NewReplaySession
	}
	return nil
}

func (m *ReplayOp) GetResponse() isReplayOp_Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (x *ReplayOp) GetCreateSessionResponse() *CreateSessionResponse {
	if x, ok := x.GetResponse().(*ReplayOp_CreateSessionResponse); ok {
		return x.CreateSessionResponse
	}
	return nil
}

func (x *ReplayOp) GetExtendSessionResponse() *ExtendSessionResponse {
	if x, ok := x.GetResponse().(*ReplayOp_ExtendSessionResponse); ok {
		return x.ExtendSessionResponse
	}
	return nil
}

func (x *ReplayOp) GetPartialRunSetupResponse() *PartialRunSetupResponse {
	if x, ok := x.GetResponse().(*ReplayOp_PartialRunSetupResponse); ok {
		return x.PartialRunSetupResponse
	}
	return nil
}

func (x *ReplayOp) GetRunStepResponse() *RunStepResponse {
	if x, ok := x.GetResponse().(*ReplayOp_RunStepResponse); ok {
		return x.RunStepResponse
	}
	return nil
}

func (x *ReplayOp) GetCloseSessionResponse() *CloseSessionResponse {
	if x, ok := x.GetResponse().(*ReplayOp_CloseSessionResponse); ok {
		return x.CloseSessionResponse
	}
	return nil
}

func (x *ReplayOp) GetListDevicesResponse() *ListDevicesResponse {
	if x, ok := x.GetResponse().(*ReplayOp_ListDevicesResponse); ok {
		return x.ListDevicesResponse
	}
	return nil
}

func (x *ReplayOp) GetResetRequestResponse() *ResetResponse {
	if x, ok := x.GetResponse().(*ReplayOp_ResetRequestResponse); ok {
		return x.ResetRequestResponse
	}
	return nil
}

func (x *ReplayOp) GetMakeCallableResponse() *MakeCallableResponse {
	if x, ok := x.GetResponse().(*ReplayOp_MakeCallableResponse); ok {
		return x.MakeCallableResponse
	}
	return nil
}

func (x *ReplayOp) GetRunCallableResponse() *RunCallableResponse {
	if x, ok := x.GetResponse().(*ReplayOp_RunCallableResponse); ok {
		return x.RunCallableResponse
	}
	return nil
}

func (x *ReplayOp) GetReleaseCallableResponse() *ReleaseCallableResponse {
	if x, ok := x.GetResponse().(*ReplayOp_ReleaseCallableResponse); ok {
		return x.ReleaseCallableResponse
	}
	return nil
}

type isReplayOp_Op interface {
	isReplayOp_Op()
}

type ReplayOp_CreateSession struct {
	CreateSession *CreateSessionRequest `protobuf:"bytes,1,opt,name=create_session,json=createSession,proto3,oneof"`
}

type ReplayOp_ExtendSession struct {
	ExtendSession *ExtendSessionRequest `protobuf:"bytes,2,opt,name=extend_session,json=extendSession,proto3,oneof"`
}

type ReplayOp_PartialRunSetup struct {
	PartialRunSetup *PartialRunSetupRequest `protobuf:"bytes,3,opt,name=partial_run_setup,json=partialRunSetup,proto3,oneof"`
}

type ReplayOp_RunStep struct {
	RunStep *RunStepRequest `protobuf:"bytes,4,opt,name=run_step,json=runStep,proto3,oneof"`
}

type ReplayOp_CloseSession struct {
	CloseSession *CloseSessionRequest `protobuf:"bytes,5,opt,name=close_session,json=closeSession,proto3,oneof"`
}

type ReplayOp_ListDevices struct {
	ListDevices *ListDevicesRequest `protobuf:"bytes,6,opt,name=list_devices,json=listDevices,proto3,oneof"`
}

type ReplayOp_ResetRequest struct {
	ResetRequest *ResetRequest `protobuf:"bytes,7,opt,name=reset_request,json=resetRequest,proto3,oneof"`
}

type ReplayOp_MakeCallable struct {
	MakeCallable *MakeCallableRequest `protobuf:"bytes,8,opt,name=make_callable,json=makeCallable,proto3,oneof"`
}

type ReplayOp_RunCallable struct {
	RunCallable *RunCallableRequest `protobuf:"bytes,9,opt,name=run_callable,json=runCallable,proto3,oneof"`
}

type ReplayOp_ReleaseCallable struct {
	ReleaseCallable *ReleaseCallableRequest `protobuf:"bytes,10,opt,name=release_callable,json=releaseCallable,proto3,oneof"`
}

type ReplayOp_NewReplaySession struct {
	NewReplaySession *NewReplaySession `protobuf:"bytes,11,opt,name=new_replay_session,json=newReplaySession,proto3,oneof"`
}

func (*ReplayOp_CreateSession) isReplayOp_Op() {}

func (*ReplayOp_ExtendSession) isReplayOp_Op() {}

func (*ReplayOp_PartialRunSetup) isReplayOp_Op() {}

func (*ReplayOp_RunStep) isReplayOp_Op() {}

func (*ReplayOp_CloseSession) isReplayOp_Op() {}

func (*ReplayOp_ListDevices) isReplayOp_Op() {}

func (*ReplayOp_ResetRequest) isReplayOp_Op() {}

func (*ReplayOp_MakeCallable) isReplayOp_Op() {}

func (*ReplayOp_RunCallable) isReplayOp_Op() {}

func (*ReplayOp_ReleaseCallable) isReplayOp_Op() {}

func (*ReplayOp_NewReplaySession) isReplayOp_Op() {}

type isReplayOp_Response interface {
	isReplayOp_Response()
}

type ReplayOp_CreateSessionResponse struct {
	CreateSessionResponse *CreateSessionResponse `protobuf:"bytes,21,opt,name=create_session_response,json=createSessionResponse,proto3,oneof"`
}

type ReplayOp_ExtendSessionResponse struct {
	ExtendSessionResponse *ExtendSessionResponse `protobuf:"bytes,22,opt,name=extend_session_response,json=extendSessionResponse,proto3,oneof"`
}

type ReplayOp_PartialRunSetupResponse struct {
	PartialRunSetupResponse *PartialRunSetupResponse `protobuf:"bytes,23,opt,name=partial_run_setup_response,json=partialRunSetupResponse,proto3,oneof"`
}

type ReplayOp_RunStepResponse struct {
	RunStepResponse *RunStepResponse `protobuf:"bytes,24,opt,name=run_step_response,json=runStepResponse,proto3,oneof"`
}

type ReplayOp_CloseSessionResponse struct {
	CloseSessionResponse *CloseSessionResponse `protobuf:"bytes,25,opt,name=close_session_response,json=closeSessionResponse,proto3,oneof"`
}

type ReplayOp_ListDevicesResponse struct {
	ListDevicesResponse *ListDevicesResponse `protobuf:"bytes,26,opt,name=list_devices_response,json=listDevicesResponse,proto3,oneof"`
}

type ReplayOp_ResetRequestResponse struct {
	ResetRequestResponse *ResetResponse `protobuf:"bytes,27,opt,name=reset_request_response,json=resetRequestResponse,proto3,oneof"`
}

type ReplayOp_MakeCallableResponse struct {
	MakeCallableResponse *MakeCallableResponse `protobuf:"bytes,28,opt,name=make_callable_response,json=makeCallableResponse,proto3,oneof"`
}

type ReplayOp_RunCallableResponse struct {
	RunCallableResponse *RunCallableResponse `protobuf:"bytes,29,opt,name=run_callable_response,json=runCallableResponse,proto3,oneof"`
}

type ReplayOp_ReleaseCallableResponse struct {
	ReleaseCallableResponse *ReleaseCallableResponse `protobuf:"bytes,30,opt,name=release_callable_response,json=releaseCallableResponse,proto3,oneof"`
}

func (*ReplayOp_CreateSessionResponse) isReplayOp_Response() {}

func (*ReplayOp_ExtendSessionResponse) isReplayOp_Response() {}

func (*ReplayOp_PartialRunSetupResponse) isReplayOp_Response() {}

func (*ReplayOp_RunStepResponse) isReplayOp_Response() {}

func (*ReplayOp_CloseSessionResponse) isReplayOp_Response() {}

func (*ReplayOp_ListDevicesResponse) isReplayOp_Response() {}

func (*ReplayOp_ResetRequestResponse) isReplayOp_Response() {}

func (*ReplayOp_MakeCallableResponse) isReplayOp_Response() {}

func (*ReplayOp_RunCallableResponse) isReplayOp_Response() {}

func (*ReplayOp_ReleaseCallableResponse) isReplayOp_Response() {}

var File_tensorflow_core_protobuf_replay_log_proto protoreflect.FileDescriptor

var file_tensorflow_core_protobuf_replay_log_proto_rawDesc = []byte{
	0x0a, 0x29, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x72, 0x65, 0x70, 0x6c, 0x61,
	0x79, 0x5f, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x74, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x1a, 0x25, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66,
	0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x74,
	0x0a, 0x10, 0x4e, 0x65, 0x77, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x79, 0x53, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x39, 0x0a, 0x07, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x52, 0x07, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x25, 0x0a,
	0x0e, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x48, 0x61,
	0x6e, 0x64, 0x6c, 0x65, 0x22, 0xfc, 0x0d, 0x0a, 0x08, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x79, 0x4f,
	0x70, 0x12, 0x22, 0x0a, 0x0d, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f,
	0x75, 0x73, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54,
	0x69, 0x6d, 0x65, 0x55, 0x73, 0x12, 0x1e, 0x0a, 0x0b, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x5f, 0x75, 0x73, 0x18, 0x20, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x65, 0x6e, 0x64, 0x54,
	0x69, 0x6d, 0x65, 0x55, 0x73, 0x12, 0x49, 0x0a, 0x0e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f,
	0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e,
	0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48,
	0x00, 0x52, 0x0d, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x49, 0x0a, 0x0e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x5f, 0x73, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x53, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x0d, 0x65, 0x78,
	0x74, 0x65, 0x6e, 0x64, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x50, 0x0a, 0x11, 0x70,
	0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x72, 0x75, 0x6e, 0x5f, 0x73, 0x65, 0x74, 0x75, 0x70,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66,
	0x6c, 0x6f, 0x77, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x52, 0x75, 0x6e, 0x53, 0x65,
	0x74, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x0f, 0x70, 0x61,
	0x72, 0x74, 0x69, 0x61, 0x6c, 0x52, 0x75, 0x6e, 0x53, 0x65, 0x74, 0x75, 0x70, 0x12, 0x37, 0x0a,
	0x08, 0x72, 0x75, 0x6e, 0x5f, 0x73, 0x74, 0x65, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x52, 0x75, 0x6e,
	0x53, 0x74, 0x65, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x07, 0x72,
	0x75, 0x6e, 0x53, 0x74, 0x65, 0x70, 0x12, 0x46, 0x0a, 0x0d, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x5f,
	0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e,
	0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x43, 0x6c, 0x6f, 0x73, 0x65,
	0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00,
	0x52, 0x0c, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x43,
	0x0a, 0x0c, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f,
	0x77, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x0b, 0x6c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x12, 0x3f, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x65, 0x74, 0x5f, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x74, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x65, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x46, 0x0a, 0x0d, 0x6d, 0x61, 0x6b, 0x65, 0x5f, 0x63, 0x61, 0x6c,
	0x6c, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x74, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x4d, 0x61, 0x6b, 0x65, 0x43, 0x61, 0x6c,
	0x6c, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x0c,
	0x6d, 0x61, 0x6b, 0x65, 0x43, 0x61, 0x6c, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x43, 0x0a, 0x0c,
	0x72, 0x75, 0x6e, 0x5f, 0x63, 0x61, 0x6c, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e,
	0x52, 0x75, 0x6e, 0x43, 0x61, 0x6c, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x48, 0x00, 0x52, 0x0b, 0x72, 0x75, 0x6e, 0x43, 0x61, 0x6c, 0x6c, 0x61, 0x62, 0x6c,
	0x65, 0x12, 0x4f, 0x0a, 0x10, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x5f, 0x63, 0x61, 0x6c,
	0x6c, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x74, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65,
	0x43, 0x61, 0x6c, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48,
	0x00, 0x52, 0x0f, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x43, 0x61, 0x6c, 0x6c, 0x61, 0x62,
	0x6c, 0x65, 0x12, 0x4c, 0x0a, 0x12, 0x6e, 0x65, 0x77, 0x5f, 0x72, 0x65, 0x70, 0x6c, 0x61, 0x79,
	0x5f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x4e, 0x65, 0x77, 0x52,
	0x65, 0x70, 0x6c, 0x61, 0x79, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x10,
	0x6e, 0x65, 0x77, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x79, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x5b, 0x0a, 0x17, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x73, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x15, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x21, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x48, 0x01, 0x52, 0x15, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5b, 0x0a,
	0x17, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x5f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x16, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21,
	0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x45, 0x78, 0x74, 0x65,
	0x6e, 0x64, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x48, 0x01, 0x52, 0x15, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x53, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x62, 0x0a, 0x1a, 0x70, 0x61,
	0x72, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x72, 0x75, 0x6e, 0x5f, 0x73, 0x65, 0x74, 0x75, 0x70, 0x5f,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x17, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23,
	0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x50, 0x61, 0x72, 0x74,
	0x69, 0x61, 0x6c, 0x52, 0x75, 0x6e, 0x53, 0x65, 0x74, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x48, 0x01, 0x52, 0x17, 0x70, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x52, 0x75,
	0x6e, 0x53, 0x65, 0x74, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49,
	0x0a, 0x11, 0x72, 0x75, 0x6e, 0x5f, 0x73, 0x74, 0x65, 0x70, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x18, 0x18, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x74, 0x65, 0x6e, 0x73,
	0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x52, 0x75, 0x6e, 0x53, 0x74, 0x65, 0x70, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x01, 0x52, 0x0f, 0x72, 0x75, 0x6e, 0x53, 0x74, 0x65,
	0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x58, 0x0a, 0x16, 0x63, 0x6c, 0x6f,
	0x73, 0x65, 0x5f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x18, 0x19, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x74, 0x65, 0x6e, 0x73,
	0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x53, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x01, 0x52, 0x14, 0x63,
	0x6c, 0x6f, 0x73, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x55, 0x0a, 0x15, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x1a, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x48, 0x01, 0x52, 0x13, 0x6c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x51, 0x0a, 0x16, 0x72, 0x65,
	0x73, 0x65, 0x74, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x18, 0x1b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x74, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x01, 0x52, 0x14, 0x72, 0x65, 0x73, 0x65, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x58, 0x0a,
	0x16, 0x6d, 0x61, 0x6b, 0x65, 0x5f, 0x63, 0x61, 0x6c, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x1c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e,
	0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x4d, 0x61, 0x6b, 0x65, 0x43,
	0x61, 0x6c, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48,
	0x01, 0x52, 0x14, 0x6d, 0x61, 0x6b, 0x65, 0x43, 0x61, 0x6c, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x55, 0x0a, 0x15, 0x72, 0x75, 0x6e, 0x5f, 0x63,
	0x61, 0x6c, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x18, 0x1d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66,
	0x6c, 0x6f, 0x77, 0x2e, 0x52, 0x75, 0x6e, 0x43, 0x61, 0x6c, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x01, 0x52, 0x13, 0x72, 0x75, 0x6e, 0x43, 0x61,
	0x6c, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x61,
	0x0a, 0x19, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x5f, 0x63, 0x61, 0x6c, 0x6c, 0x61, 0x62,
	0x6c, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x1e, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x23, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x52,
	0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x43, 0x61, 0x6c, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x01, 0x52, 0x17, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73,
	0x65, 0x43, 0x61, 0x6c, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x04, 0x0a, 0x02, 0x6f, 0x70, 0x42, 0x0a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x4d, 0x5a, 0x48, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c,
	0x6f, 0x77, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x5f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x5f, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xf8,
	0x01, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tensorflow_core_protobuf_replay_log_proto_rawDescOnce sync.Once
	file_tensorflow_core_protobuf_replay_log_proto_rawDescData = file_tensorflow_core_protobuf_replay_log_proto_rawDesc
)

func file_tensorflow_core_protobuf_replay_log_proto_rawDescGZIP() []byte {
	file_tensorflow_core_protobuf_replay_log_proto_rawDescOnce.Do(func() {
		file_tensorflow_core_protobuf_replay_log_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_core_protobuf_replay_log_proto_rawDescData)
	})
	return file_tensorflow_core_protobuf_replay_log_proto_rawDescData
}

var file_tensorflow_core_protobuf_replay_log_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_tensorflow_core_protobuf_replay_log_proto_goTypes = []interface{}{
	(*NewReplaySession)(nil),        // 0: tensorflow.NewReplaySession
	(*ReplayOp)(nil),                // 1: tensorflow.ReplayOp
	(*ListDevicesResponse)(nil),     // 2: tensorflow.ListDevicesResponse
	(*CreateSessionRequest)(nil),    // 3: tensorflow.CreateSessionRequest
	(*ExtendSessionRequest)(nil),    // 4: tensorflow.ExtendSessionRequest
	(*PartialRunSetupRequest)(nil),  // 5: tensorflow.PartialRunSetupRequest
	(*RunStepRequest)(nil),          // 6: tensorflow.RunStepRequest
	(*CloseSessionRequest)(nil),     // 7: tensorflow.CloseSessionRequest
	(*ListDevicesRequest)(nil),      // 8: tensorflow.ListDevicesRequest
	(*ResetRequest)(nil),            // 9: tensorflow.ResetRequest
	(*MakeCallableRequest)(nil),     // 10: tensorflow.MakeCallableRequest
	(*RunCallableRequest)(nil),      // 11: tensorflow.RunCallableRequest
	(*ReleaseCallableRequest)(nil),  // 12: tensorflow.ReleaseCallableRequest
	(*CreateSessionResponse)(nil),   // 13: tensorflow.CreateSessionResponse
	(*ExtendSessionResponse)(nil),   // 14: tensorflow.ExtendSessionResponse
	(*PartialRunSetupResponse)(nil), // 15: tensorflow.PartialRunSetupResponse
	(*RunStepResponse)(nil),         // 16: tensorflow.RunStepResponse
	(*CloseSessionResponse)(nil),    // 17: tensorflow.CloseSessionResponse
	(*ResetResponse)(nil),           // 18: tensorflow.ResetResponse
	(*MakeCallableResponse)(nil),    // 19: tensorflow.MakeCallableResponse
	(*RunCallableResponse)(nil),     // 20: tensorflow.RunCallableResponse
	(*ReleaseCallableResponse)(nil), // 21: tensorflow.ReleaseCallableResponse
}
var file_tensorflow_core_protobuf_replay_log_proto_depIdxs = []int32{
	2,  // 0: tensorflow.NewReplaySession.devices:type_name -> tensorflow.ListDevicesResponse
	3,  // 1: tensorflow.ReplayOp.create_session:type_name -> tensorflow.CreateSessionRequest
	4,  // 2: tensorflow.ReplayOp.extend_session:type_name -> tensorflow.ExtendSessionRequest
	5,  // 3: tensorflow.ReplayOp.partial_run_setup:type_name -> tensorflow.PartialRunSetupRequest
	6,  // 4: tensorflow.ReplayOp.run_step:type_name -> tensorflow.RunStepRequest
	7,  // 5: tensorflow.ReplayOp.close_session:type_name -> tensorflow.CloseSessionRequest
	8,  // 6: tensorflow.ReplayOp.list_devices:type_name -> tensorflow.ListDevicesRequest
	9,  // 7: tensorflow.ReplayOp.reset_request:type_name -> tensorflow.ResetRequest
	10, // 8: tensorflow.ReplayOp.make_callable:type_name -> tensorflow.MakeCallableRequest
	11, // 9: tensorflow.ReplayOp.run_callable:type_name -> tensorflow.RunCallableRequest
	12, // 10: tensorflow.ReplayOp.release_callable:type_name -> tensorflow.ReleaseCallableRequest
	0,  // 11: tensorflow.ReplayOp.new_replay_session:type_name -> tensorflow.NewReplaySession
	13, // 12: tensorflow.ReplayOp.create_session_response:type_name -> tensorflow.CreateSessionResponse
	14, // 13: tensorflow.ReplayOp.extend_session_response:type_name -> tensorflow.ExtendSessionResponse
	15, // 14: tensorflow.ReplayOp.partial_run_setup_response:type_name -> tensorflow.PartialRunSetupResponse
	16, // 15: tensorflow.ReplayOp.run_step_response:type_name -> tensorflow.RunStepResponse
	17, // 16: tensorflow.ReplayOp.close_session_response:type_name -> tensorflow.CloseSessionResponse
	2,  // 17: tensorflow.ReplayOp.list_devices_response:type_name -> tensorflow.ListDevicesResponse
	18, // 18: tensorflow.ReplayOp.reset_request_response:type_name -> tensorflow.ResetResponse
	19, // 19: tensorflow.ReplayOp.make_callable_response:type_name -> tensorflow.MakeCallableResponse
	20, // 20: tensorflow.ReplayOp.run_callable_response:type_name -> tensorflow.RunCallableResponse
	21, // 21: tensorflow.ReplayOp.release_callable_response:type_name -> tensorflow.ReleaseCallableResponse
	22, // [22:22] is the sub-list for method output_type
	22, // [22:22] is the sub-list for method input_type
	22, // [22:22] is the sub-list for extension type_name
	22, // [22:22] is the sub-list for extension extendee
	0,  // [0:22] is the sub-list for field type_name
}

func init() { file_tensorflow_core_protobuf_replay_log_proto_init() }
func file_tensorflow_core_protobuf_replay_log_proto_init() {
	if File_tensorflow_core_protobuf_replay_log_proto != nil {
		return
	}
	file_tensorflow_core_protobuf_master_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_core_protobuf_replay_log_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewReplaySession); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tensorflow_core_protobuf_replay_log_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReplayOp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_tensorflow_core_protobuf_replay_log_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*ReplayOp_CreateSession)(nil),
		(*ReplayOp_ExtendSession)(nil),
		(*ReplayOp_PartialRunSetup)(nil),
		(*ReplayOp_RunStep)(nil),
		(*ReplayOp_CloseSession)(nil),
		(*ReplayOp_ListDevices)(nil),
		(*ReplayOp_ResetRequest)(nil),
		(*ReplayOp_MakeCallable)(nil),
		(*ReplayOp_RunCallable)(nil),
		(*ReplayOp_ReleaseCallable)(nil),
		(*ReplayOp_NewReplaySession)(nil),
		(*ReplayOp_CreateSessionResponse)(nil),
		(*ReplayOp_ExtendSessionResponse)(nil),
		(*ReplayOp_PartialRunSetupResponse)(nil),
		(*ReplayOp_RunStepResponse)(nil),
		(*ReplayOp_CloseSessionResponse)(nil),
		(*ReplayOp_ListDevicesResponse)(nil),
		(*ReplayOp_ResetRequestResponse)(nil),
		(*ReplayOp_MakeCallableResponse)(nil),
		(*ReplayOp_RunCallableResponse)(nil),
		(*ReplayOp_ReleaseCallableResponse)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tensorflow_core_protobuf_replay_log_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_core_protobuf_replay_log_proto_goTypes,
		DependencyIndexes: file_tensorflow_core_protobuf_replay_log_proto_depIdxs,
		MessageInfos:      file_tensorflow_core_protobuf_replay_log_proto_msgTypes,
	}.Build()
	File_tensorflow_core_protobuf_replay_log_proto = out.File
	file_tensorflow_core_protobuf_replay_log_proto_rawDesc = nil
	file_tensorflow_core_protobuf_replay_log_proto_goTypes = nil
	file_tensorflow_core_protobuf_replay_log_proto_depIdxs = nil
}
