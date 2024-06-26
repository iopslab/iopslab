// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.20.1
// source: entity/dependencies_service_v2_request.proto

package grpc

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Dependency struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *Dependency) Reset() {
	*x = Dependency{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entity_dependencies_service_v2_request_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Dependency) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Dependency) ProtoMessage() {}

func (x *Dependency) ProtoReflect() protoreflect.Message {
	mi := &file_entity_dependencies_service_v2_request_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Dependency.ProtoReflect.Descriptor instead.
func (*Dependency) Descriptor() ([]byte, []int) {
	return file_entity_dependencies_service_v2_request_proto_rawDescGZIP(), []int{0}
}

func (x *Dependency) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Dependency) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

type DependenciesServiceV2ConnectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeKey string `protobuf:"bytes,1,opt,name=node_key,json=nodeKey,proto3" json:"node_key,omitempty"`
}

func (x *DependenciesServiceV2ConnectRequest) Reset() {
	*x = DependenciesServiceV2ConnectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entity_dependencies_service_v2_request_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DependenciesServiceV2ConnectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DependenciesServiceV2ConnectRequest) ProtoMessage() {}

func (x *DependenciesServiceV2ConnectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_entity_dependencies_service_v2_request_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DependenciesServiceV2ConnectRequest.ProtoReflect.Descriptor instead.
func (*DependenciesServiceV2ConnectRequest) Descriptor() ([]byte, []int) {
	return file_entity_dependencies_service_v2_request_proto_rawDescGZIP(), []int{1}
}

func (x *DependenciesServiceV2ConnectRequest) GetNodeKey() string {
	if x != nil {
		return x.NodeKey
	}
	return ""
}

type DependenciesServiceV2SyncRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeKey      string        `protobuf:"bytes,1,opt,name=node_key,json=nodeKey,proto3" json:"node_key,omitempty"`
	Lang         string        `protobuf:"bytes,2,opt,name=lang,proto3" json:"lang,omitempty"`
	Dependencies []*Dependency `protobuf:"bytes,3,rep,name=dependencies,proto3" json:"dependencies,omitempty"`
}

func (x *DependenciesServiceV2SyncRequest) Reset() {
	*x = DependenciesServiceV2SyncRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entity_dependencies_service_v2_request_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DependenciesServiceV2SyncRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DependenciesServiceV2SyncRequest) ProtoMessage() {}

func (x *DependenciesServiceV2SyncRequest) ProtoReflect() protoreflect.Message {
	mi := &file_entity_dependencies_service_v2_request_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DependenciesServiceV2SyncRequest.ProtoReflect.Descriptor instead.
func (*DependenciesServiceV2SyncRequest) Descriptor() ([]byte, []int) {
	return file_entity_dependencies_service_v2_request_proto_rawDescGZIP(), []int{2}
}

func (x *DependenciesServiceV2SyncRequest) GetNodeKey() string {
	if x != nil {
		return x.NodeKey
	}
	return ""
}

func (x *DependenciesServiceV2SyncRequest) GetLang() string {
	if x != nil {
		return x.Lang
	}
	return ""
}

func (x *DependenciesServiceV2SyncRequest) GetDependencies() []*Dependency {
	if x != nil {
		return x.Dependencies
	}
	return nil
}

type DependenciesServiceV2InstallRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeKey      string        `protobuf:"bytes,1,opt,name=node_key,json=nodeKey,proto3" json:"node_key,omitempty"`
	Lang         string        `protobuf:"bytes,2,opt,name=lang,proto3" json:"lang,omitempty"`
	Dependencies []*Dependency `protobuf:"bytes,3,rep,name=dependencies,proto3" json:"dependencies,omitempty"`
	Proxy        string        `protobuf:"bytes,4,opt,name=proxy,proto3" json:"proxy,omitempty"`
}

func (x *DependenciesServiceV2InstallRequest) Reset() {
	*x = DependenciesServiceV2InstallRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entity_dependencies_service_v2_request_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DependenciesServiceV2InstallRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DependenciesServiceV2InstallRequest) ProtoMessage() {}

func (x *DependenciesServiceV2InstallRequest) ProtoReflect() protoreflect.Message {
	mi := &file_entity_dependencies_service_v2_request_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DependenciesServiceV2InstallRequest.ProtoReflect.Descriptor instead.
func (*DependenciesServiceV2InstallRequest) Descriptor() ([]byte, []int) {
	return file_entity_dependencies_service_v2_request_proto_rawDescGZIP(), []int{3}
}

func (x *DependenciesServiceV2InstallRequest) GetNodeKey() string {
	if x != nil {
		return x.NodeKey
	}
	return ""
}

func (x *DependenciesServiceV2InstallRequest) GetLang() string {
	if x != nil {
		return x.Lang
	}
	return ""
}

func (x *DependenciesServiceV2InstallRequest) GetDependencies() []*Dependency {
	if x != nil {
		return x.Dependencies
	}
	return nil
}

func (x *DependenciesServiceV2InstallRequest) GetProxy() string {
	if x != nil {
		return x.Proxy
	}
	return ""
}

type DependenciesServiceV2UninstallRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeKey      string        `protobuf:"bytes,1,opt,name=node_key,json=nodeKey,proto3" json:"node_key,omitempty"`
	Lang         string        `protobuf:"bytes,2,opt,name=lang,proto3" json:"lang,omitempty"`
	Dependencies []*Dependency `protobuf:"bytes,3,rep,name=dependencies,proto3" json:"dependencies,omitempty"`
}

func (x *DependenciesServiceV2UninstallRequest) Reset() {
	*x = DependenciesServiceV2UninstallRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entity_dependencies_service_v2_request_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DependenciesServiceV2UninstallRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DependenciesServiceV2UninstallRequest) ProtoMessage() {}

func (x *DependenciesServiceV2UninstallRequest) ProtoReflect() protoreflect.Message {
	mi := &file_entity_dependencies_service_v2_request_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DependenciesServiceV2UninstallRequest.ProtoReflect.Descriptor instead.
func (*DependenciesServiceV2UninstallRequest) Descriptor() ([]byte, []int) {
	return file_entity_dependencies_service_v2_request_proto_rawDescGZIP(), []int{4}
}

func (x *DependenciesServiceV2UninstallRequest) GetNodeKey() string {
	if x != nil {
		return x.NodeKey
	}
	return ""
}

func (x *DependenciesServiceV2UninstallRequest) GetLang() string {
	if x != nil {
		return x.Lang
	}
	return ""
}

func (x *DependenciesServiceV2UninstallRequest) GetDependencies() []*Dependency {
	if x != nil {
		return x.Dependencies
	}
	return nil
}

var File_entity_dependencies_service_v2_request_proto protoreflect.FileDescriptor

var file_entity_dependencies_service_v2_request_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x64, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65,
	0x6e, 0x63, 0x69, 0x65, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x76, 0x32,
	0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04,
	0x67, 0x72, 0x70, 0x63, 0x22, 0x3a, 0x0a, 0x0a, 0x44, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e,
	0x63, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x22, 0x40, 0x0a, 0x23, 0x44, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x56, 0x32, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x5f,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6e, 0x6f, 0x64, 0x65, 0x4b,
	0x65, 0x79, 0x22, 0x87, 0x01, 0x0a, 0x20, 0x44, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63,
	0x69, 0x65, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x56, 0x32, 0x53, 0x79, 0x6e, 0x63,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x5f,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6e, 0x6f, 0x64, 0x65, 0x4b,
	0x65, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x61, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6c, 0x61, 0x6e, 0x67, 0x12, 0x34, 0x0a, 0x0c, 0x64, 0x65, 0x70, 0x65, 0x6e, 0x64,
	0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x44, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x0c,
	0x64, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x22, 0xa0, 0x01, 0x0a,
	0x23, 0x44, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x56, 0x32, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6e, 0x6f, 0x64, 0x65, 0x4b, 0x65, 0x79, 0x12,
	0x12, 0x0a, 0x04, 0x6c, 0x61, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c,
	0x61, 0x6e, 0x67, 0x12, 0x34, 0x0a, 0x0c, 0x64, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63,
	0x69, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x44, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x0c, 0x64, 0x65, 0x70,
	0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x6f,
	0x78, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x22,
	0x8c, 0x01, 0x0a, 0x25, 0x44, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x56, 0x32, 0x55, 0x6e, 0x69, 0x6e, 0x73, 0x74, 0x61,
	0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6e, 0x6f, 0x64,
	0x65, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6e, 0x6f, 0x64,
	0x65, 0x4b, 0x65, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x61, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6c, 0x61, 0x6e, 0x67, 0x12, 0x34, 0x0a, 0x0c, 0x64, 0x65, 0x70, 0x65,
	0x6e, 0x64, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x44, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x79,
	0x52, 0x0c, 0x64, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x42, 0x08,
	0x5a, 0x06, 0x2e, 0x3b, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_entity_dependencies_service_v2_request_proto_rawDescOnce sync.Once
	file_entity_dependencies_service_v2_request_proto_rawDescData = file_entity_dependencies_service_v2_request_proto_rawDesc
)

func file_entity_dependencies_service_v2_request_proto_rawDescGZIP() []byte {
	file_entity_dependencies_service_v2_request_proto_rawDescOnce.Do(func() {
		file_entity_dependencies_service_v2_request_proto_rawDescData = protoimpl.X.CompressGZIP(file_entity_dependencies_service_v2_request_proto_rawDescData)
	})
	return file_entity_dependencies_service_v2_request_proto_rawDescData
}

var file_entity_dependencies_service_v2_request_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_entity_dependencies_service_v2_request_proto_goTypes = []interface{}{
	(*Dependency)(nil),                            // 0: grpc.Dependency
	(*DependenciesServiceV2ConnectRequest)(nil),   // 1: grpc.DependenciesServiceV2ConnectRequest
	(*DependenciesServiceV2SyncRequest)(nil),      // 2: grpc.DependenciesServiceV2SyncRequest
	(*DependenciesServiceV2InstallRequest)(nil),   // 3: grpc.DependenciesServiceV2InstallRequest
	(*DependenciesServiceV2UninstallRequest)(nil), // 4: grpc.DependenciesServiceV2UninstallRequest
}
var file_entity_dependencies_service_v2_request_proto_depIdxs = []int32{
	0, // 0: grpc.DependenciesServiceV2SyncRequest.dependencies:type_name -> grpc.Dependency
	0, // 1: grpc.DependenciesServiceV2InstallRequest.dependencies:type_name -> grpc.Dependency
	0, // 2: grpc.DependenciesServiceV2UninstallRequest.dependencies:type_name -> grpc.Dependency
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_entity_dependencies_service_v2_request_proto_init() }
func file_entity_dependencies_service_v2_request_proto_init() {
	if File_entity_dependencies_service_v2_request_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_entity_dependencies_service_v2_request_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Dependency); i {
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
		file_entity_dependencies_service_v2_request_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DependenciesServiceV2ConnectRequest); i {
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
		file_entity_dependencies_service_v2_request_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DependenciesServiceV2SyncRequest); i {
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
		file_entity_dependencies_service_v2_request_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DependenciesServiceV2InstallRequest); i {
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
		file_entity_dependencies_service_v2_request_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DependenciesServiceV2UninstallRequest); i {
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
			RawDescriptor: file_entity_dependencies_service_v2_request_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_entity_dependencies_service_v2_request_proto_goTypes,
		DependencyIndexes: file_entity_dependencies_service_v2_request_proto_depIdxs,
		MessageInfos:      file_entity_dependencies_service_v2_request_proto_msgTypes,
	}.Build()
	File_entity_dependencies_service_v2_request_proto = out.File
	file_entity_dependencies_service_v2_request_proto_rawDesc = nil
	file_entity_dependencies_service_v2_request_proto_goTypes = nil
	file_entity_dependencies_service_v2_request_proto_depIdxs = nil
}
