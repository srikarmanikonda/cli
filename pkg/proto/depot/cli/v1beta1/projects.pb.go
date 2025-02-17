// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: depot/cli/v1beta1/projects.proto

package cliv1beta1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ListProjectsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListProjectsRequest) Reset() {
	*x = ListProjectsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_depot_cli_v1beta1_projects_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListProjectsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListProjectsRequest) ProtoMessage() {}

func (x *ListProjectsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_depot_cli_v1beta1_projects_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListProjectsRequest.ProtoReflect.Descriptor instead.
func (*ListProjectsRequest) Descriptor() ([]byte, []int) {
	return file_depot_cli_v1beta1_projects_proto_rawDescGZIP(), []int{0}
}

type ListProjectsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Projects []*ListProjectsResponse_Project `protobuf:"bytes,1,rep,name=projects,proto3" json:"projects,omitempty"`
}

func (x *ListProjectsResponse) Reset() {
	*x = ListProjectsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_depot_cli_v1beta1_projects_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListProjectsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListProjectsResponse) ProtoMessage() {}

func (x *ListProjectsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_depot_cli_v1beta1_projects_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListProjectsResponse.ProtoReflect.Descriptor instead.
func (*ListProjectsResponse) Descriptor() ([]byte, []int) {
	return file_depot_cli_v1beta1_projects_proto_rawDescGZIP(), []int{1}
}

func (x *ListProjectsResponse) GetProjects() []*ListProjectsResponse_Project {
	if x != nil {
		return x.Projects
	}
	return nil
}

type ResetProjectCacheRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectId string `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
}

func (x *ResetProjectCacheRequest) Reset() {
	*x = ResetProjectCacheRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_depot_cli_v1beta1_projects_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResetProjectCacheRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResetProjectCacheRequest) ProtoMessage() {}

func (x *ResetProjectCacheRequest) ProtoReflect() protoreflect.Message {
	mi := &file_depot_cli_v1beta1_projects_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResetProjectCacheRequest.ProtoReflect.Descriptor instead.
func (*ResetProjectCacheRequest) Descriptor() ([]byte, []int) {
	return file_depot_cli_v1beta1_projects_proto_rawDescGZIP(), []int{2}
}

func (x *ResetProjectCacheRequest) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

type ResetProjectCacheResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	OrgName string `protobuf:"bytes,2,opt,name=org_name,json=orgName,proto3" json:"org_name,omitempty"`
}

func (x *ResetProjectCacheResponse) Reset() {
	*x = ResetProjectCacheResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_depot_cli_v1beta1_projects_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResetProjectCacheResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResetProjectCacheResponse) ProtoMessage() {}

func (x *ResetProjectCacheResponse) ProtoReflect() protoreflect.Message {
	mi := &file_depot_cli_v1beta1_projects_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResetProjectCacheResponse.ProtoReflect.Descriptor instead.
func (*ResetProjectCacheResponse) Descriptor() ([]byte, []int) {
	return file_depot_cli_v1beta1_projects_proto_rawDescGZIP(), []int{3}
}

func (x *ResetProjectCacheResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ResetProjectCacheResponse) GetOrgName() string {
	if x != nil {
		return x.OrgName
	}
	return ""
}

type ListProjectsResponse_Project struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	OrgId   string `protobuf:"bytes,3,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	OrgName string `protobuf:"bytes,4,opt,name=org_name,json=orgName,proto3" json:"org_name,omitempty"`
}

func (x *ListProjectsResponse_Project) Reset() {
	*x = ListProjectsResponse_Project{}
	if protoimpl.UnsafeEnabled {
		mi := &file_depot_cli_v1beta1_projects_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListProjectsResponse_Project) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListProjectsResponse_Project) ProtoMessage() {}

func (x *ListProjectsResponse_Project) ProtoReflect() protoreflect.Message {
	mi := &file_depot_cli_v1beta1_projects_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListProjectsResponse_Project.ProtoReflect.Descriptor instead.
func (*ListProjectsResponse_Project) Descriptor() ([]byte, []int) {
	return file_depot_cli_v1beta1_projects_proto_rawDescGZIP(), []int{1, 0}
}

func (x *ListProjectsResponse_Project) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ListProjectsResponse_Project) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ListProjectsResponse_Project) GetOrgId() string {
	if x != nil {
		return x.OrgId
	}
	return ""
}

func (x *ListProjectsResponse_Project) GetOrgName() string {
	if x != nil {
		return x.OrgName
	}
	return ""
}

var File_depot_cli_v1beta1_projects_proto protoreflect.FileDescriptor

var file_depot_cli_v1beta1_projects_proto_rawDesc = []byte{
	0x0a, 0x20, 0x64, 0x65, 0x70, 0x6f, 0x74, 0x2f, 0x63, 0x6c, 0x69, 0x2f, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x11, 0x64, 0x65, 0x70, 0x6f, 0x74, 0x2e, 0x63, 0x6c, 0x69, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x15, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xc4, 0x01,
	0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4b, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x64, 0x65, 0x70, 0x6f, 0x74,
	0x2e, 0x63, 0x6c, 0x69, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x73, 0x1a, 0x5f, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x6f, 0x72, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x6f, 0x72, 0x67, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x67,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x67,
	0x4e, 0x61, 0x6d, 0x65, 0x22, 0x39, 0x0a, 0x18, 0x52, 0x65, 0x73, 0x65, 0x74, 0x50, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x43, 0x61, 0x63, 0x68, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x22,
	0x4a, 0x0a, 0x19, 0x52, 0x65, 0x73, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x43,
	0x61, 0x63, 0x68, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x67, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0x32, 0xe2, 0x01, 0x0a, 0x0f,
	0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x5f, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x12,
	0x26, 0x2e, 0x64, 0x65, 0x70, 0x6f, 0x74, 0x2e, 0x63, 0x6c, 0x69, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x64, 0x65, 0x70, 0x6f, 0x74, 0x2e,
	0x63, 0x6c, 0x69, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x6e, 0x0a, 0x11, 0x52, 0x65, 0x73, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x43, 0x61, 0x63, 0x68, 0x65, 0x12, 0x2b, 0x2e, 0x64, 0x65, 0x70, 0x6f, 0x74, 0x2e, 0x63, 0x6c,
	0x69, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x74, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x43, 0x61, 0x63, 0x68, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x64, 0x65, 0x70, 0x6f, 0x74, 0x2e, 0x63, 0x6c, 0x69, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x43, 0x61, 0x63, 0x68, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0xc9, 0x01, 0x0a, 0x15, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x65, 0x70, 0x6f, 0x74, 0x2e, 0x63,
	0x6c, 0x69, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x42, 0x0d, 0x50, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3b, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x65, 0x70, 0x6f, 0x74, 0x2f, 0x63, 0x6c,
	0x69, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x65, 0x70, 0x6f,
	0x74, 0x2f, 0x63, 0x6c, 0x69, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x3b, 0x63, 0x6c,
	0x69, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x44, 0x43, 0x58, 0xaa, 0x02,
	0x11, 0x44, 0x65, 0x70, 0x6f, 0x74, 0x2e, 0x43, 0x6c, 0x69, 0x2e, 0x56, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0xca, 0x02, 0x11, 0x44, 0x65, 0x70, 0x6f, 0x74, 0x5c, 0x43, 0x6c, 0x69, 0x5c, 0x56,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xe2, 0x02, 0x1d, 0x44, 0x65, 0x70, 0x6f, 0x74, 0x5c, 0x43,
	0x6c, 0x69, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x13, 0x44, 0x65, 0x70, 0x6f, 0x74, 0x3a, 0x3a,
	0x43, 0x6c, 0x69, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_depot_cli_v1beta1_projects_proto_rawDescOnce sync.Once
	file_depot_cli_v1beta1_projects_proto_rawDescData = file_depot_cli_v1beta1_projects_proto_rawDesc
)

func file_depot_cli_v1beta1_projects_proto_rawDescGZIP() []byte {
	file_depot_cli_v1beta1_projects_proto_rawDescOnce.Do(func() {
		file_depot_cli_v1beta1_projects_proto_rawDescData = protoimpl.X.CompressGZIP(file_depot_cli_v1beta1_projects_proto_rawDescData)
	})
	return file_depot_cli_v1beta1_projects_proto_rawDescData
}

var file_depot_cli_v1beta1_projects_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_depot_cli_v1beta1_projects_proto_goTypes = []interface{}{
	(*ListProjectsRequest)(nil),          // 0: depot.cli.v1beta1.ListProjectsRequest
	(*ListProjectsResponse)(nil),         // 1: depot.cli.v1beta1.ListProjectsResponse
	(*ResetProjectCacheRequest)(nil),     // 2: depot.cli.v1beta1.ResetProjectCacheRequest
	(*ResetProjectCacheResponse)(nil),    // 3: depot.cli.v1beta1.ResetProjectCacheResponse
	(*ListProjectsResponse_Project)(nil), // 4: depot.cli.v1beta1.ListProjectsResponse.Project
}
var file_depot_cli_v1beta1_projects_proto_depIdxs = []int32{
	4, // 0: depot.cli.v1beta1.ListProjectsResponse.projects:type_name -> depot.cli.v1beta1.ListProjectsResponse.Project
	0, // 1: depot.cli.v1beta1.ProjectsService.ListProjects:input_type -> depot.cli.v1beta1.ListProjectsRequest
	2, // 2: depot.cli.v1beta1.ProjectsService.ResetProjectCache:input_type -> depot.cli.v1beta1.ResetProjectCacheRequest
	1, // 3: depot.cli.v1beta1.ProjectsService.ListProjects:output_type -> depot.cli.v1beta1.ListProjectsResponse
	3, // 4: depot.cli.v1beta1.ProjectsService.ResetProjectCache:output_type -> depot.cli.v1beta1.ResetProjectCacheResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_depot_cli_v1beta1_projects_proto_init() }
func file_depot_cli_v1beta1_projects_proto_init() {
	if File_depot_cli_v1beta1_projects_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_depot_cli_v1beta1_projects_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListProjectsRequest); i {
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
		file_depot_cli_v1beta1_projects_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListProjectsResponse); i {
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
		file_depot_cli_v1beta1_projects_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResetProjectCacheRequest); i {
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
		file_depot_cli_v1beta1_projects_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResetProjectCacheResponse); i {
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
		file_depot_cli_v1beta1_projects_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListProjectsResponse_Project); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_depot_cli_v1beta1_projects_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_depot_cli_v1beta1_projects_proto_goTypes,
		DependencyIndexes: file_depot_cli_v1beta1_projects_proto_depIdxs,
		MessageInfos:      file_depot_cli_v1beta1_projects_proto_msgTypes,
	}.Build()
	File_depot_cli_v1beta1_projects_proto = out.File
	file_depot_cli_v1beta1_projects_proto_rawDesc = nil
	file_depot_cli_v1beta1_projects_proto_goTypes = nil
	file_depot_cli_v1beta1_projects_proto_depIdxs = nil
}
