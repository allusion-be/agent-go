// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.26.1
// source: node.proto

package v1

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

var (
	file_node_proto_rawDescOnce sync.Once
	file_node_proto_rawDescData = file_node_proto_rawDesc
)

var File_node_proto protoreflect.FileDescriptor

var file_node_proto_depIdxs = []int32{
	0, // 0: registry.node.v1.NodeRecord.xnet:type_name -> registry.node.v1.ConnectionEndpoint
	0, // 1: registry.node.v1.NodeRecord.http:type_name -> registry.node.v1.ConnectionEndpoint
	1, // 2: registry.node.v1.NodeRecord.public_ipv4_config:type_name -> registry.node.v1.IPv4InterfaceConfig
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

var file_node_proto_goTypes = []interface{}{
	(*ConnectionEndpoint)(nil),  // 0: registry.node.v1.ConnectionEndpoint
	(*IPv4InterfaceConfig)(nil), // 1: registry.node.v1.IPv4InterfaceConfig
	(*NodeRecord)(nil),          // 2: registry.node.v1.NodeRecord
}

var file_node_proto_msgTypes = make([]protoimpl.MessageInfo, 3)

var file_node_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x72, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x22, 0x47,
	0x0a, 0x12, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x70, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x70, 0x41, 0x64, 0x64, 0x72, 0x12, 0x12, 0x0a,
	0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x70, 0x6f, 0x72,
	0x74, 0x4a, 0x04, 0x08, 0x04, 0x10, 0x05, 0x22, 0x7b, 0x0a, 0x13, 0x49, 0x50, 0x76, 0x34, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x17,
	0x0a, 0x07, 0x69, 0x70, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x69, 0x70, 0x41, 0x64, 0x64, 0x72, 0x12, 0x26, 0x0a, 0x0f, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x5f, 0x69, 0x70, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x0d, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x49, 0x70, 0x41, 0x64, 0x64, 0x72, 0x12,
	0x23, 0x0a, 0x0d, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x5f, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x4c, 0x65,
	0x6e, 0x67, 0x74, 0x68, 0x22, 0xc7, 0x05, 0x0a, 0x0a, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x12, 0x38, 0x0a, 0x04, 0x78, 0x6e, 0x65, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x24, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x6e, 0x6f, 0x64,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x45,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x04, 0x78, 0x6e, 0x65, 0x74, 0x12, 0x38, 0x0a,
	0x04, 0x68, 0x74, 0x74, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x72, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x52, 0x04, 0x68, 0x74, 0x74, 0x70, 0x12, 0x28, 0x0a, 0x10, 0x6e, 0x6f, 0x64, 0x65, 0x5f,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x0e, 0x6e, 0x6f, 0x64, 0x65, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x49,
	0x64, 0x12, 0x1c, 0x0a, 0x07, 0x63, 0x68, 0x69, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x10, 0x20, 0x01,
	0x28, 0x0c, 0x48, 0x00, 0x52, 0x06, 0x63, 0x68, 0x69, 0x70, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12,
	0x2f, 0x0a, 0x11, 0x68, 0x6f, 0x73, 0x74, 0x6f, 0x73, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x5f, 0x69, 0x64, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0f, 0x68, 0x6f,
	0x73, 0x74, 0x6f, 0x73, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x88, 0x01, 0x01,
	0x12, 0x58, 0x0a, 0x12, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x69, 0x70, 0x76, 0x34, 0x5f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x12, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x72,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x49, 0x50, 0x76, 0x34, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x48, 0x02, 0x52, 0x10, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x49, 0x70, 0x76,
	0x34, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x64, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x06, 0x64, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x88, 0x01, 0x01, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x63, 0x68, 0x69, 0x70,
	0x5f, 0x69, 0x64, 0x42, 0x14, 0x0a, 0x12, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x6f, 0x73, 0x5f, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x42, 0x15, 0x0a, 0x13, 0x5f, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x5f, 0x69, 0x70, 0x76, 0x34, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x42, 0x09, 0x0a, 0x07, 0x5f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x4a, 0x04, 0x08, 0x01, 0x10,
	0x02, 0x4a, 0x04, 0x08, 0x02, 0x10, 0x03, 0x4a, 0x04, 0x08, 0x03, 0x10, 0x04, 0x4a, 0x04, 0x08,
	0x04, 0x10, 0x05, 0x4a, 0x04, 0x08, 0x07, 0x10, 0x08, 0x4a, 0x04, 0x08, 0x08, 0x10, 0x09, 0x4a,
	0x04, 0x08, 0x09, 0x10, 0x0a, 0x4a, 0x04, 0x08, 0x0a, 0x10, 0x0b, 0x4a, 0x04, 0x08, 0x0b, 0x10,
	0x0c, 0x4a, 0x04, 0x08, 0x0c, 0x10, 0x0d, 0x4a, 0x04, 0x08, 0x0d, 0x10, 0x0e, 0x4a, 0x04, 0x08,
	0x0e, 0x10, 0x0f, 0x52, 0x07, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x52, 0x0d, 0x67, 0x6f,
	0x73, 0x73, 0x69, 0x70, 0x5f, 0x61, 0x64, 0x76, 0x65, 0x72, 0x74, 0x52, 0x0e, 0x67, 0x6f, 0x73,
	0x73, 0x69, 0x70, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0f, 0x67, 0x6f, 0x73,
	0x73, 0x69, 0x70, 0x5f, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x52, 0x11, 0x64, 0x63,
	0x6f, 0x70, 0x5f, 0x70, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x52,
	0x12, 0x70, 0x32, 0x70, 0x5f, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x73, 0x52, 0x17, 0x70, 0x72, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x65, 0x75, 0x73, 0x5f,
	0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x5f, 0x68, 0x74, 0x74, 0x70, 0x52, 0x0a, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x5f, 0x61, 0x70, 0x69, 0x52, 0x0b, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74,
	0x65, 0x5f, 0x61, 0x70, 0x69, 0x52, 0x12, 0x70, 0x72, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x65, 0x75,
	0x73, 0x5f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x52, 0x19, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x70, 0x6f, 0x72, 0x74, 0x5f, 0x74, 0x6c, 0x73, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x65, 0x52, 0x08, 0x78, 0x6e, 0x65, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x42, 0x0a,
	0x5a, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

func file_node_proto_init() {
	if File_node_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_node_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectionEndpoint); i {
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
		file_node_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IPv4InterfaceConfig); i {
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
		file_node_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeRecord); i {
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
	file_node_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_node_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_node_proto_goTypes,
		DependencyIndexes: file_node_proto_depIdxs,
		MessageInfos:      file_node_proto_msgTypes,
	}.Build()
	File_node_proto = out.File
	file_node_proto_rawDesc = nil
	file_node_proto_goTypes = nil
	file_node_proto_depIdxs = nil
}

func file_node_proto_rawDescGZIP() []byte {
	file_node_proto_rawDescOnce.Do(func() {
		file_node_proto_rawDescData = protoimpl.X.CompressGZIP(file_node_proto_rawDescData)
	})
	return file_node_proto_rawDescData
}

// A connection endpoint.
type ConnectionEndpoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The IP address. Senders SHOULD use dotted-quad notation for IPv4 addresses
	// and RFC5952 representation for IPv6 addresses (which means that IPv6
	// addresses are *not* enclosed in `[` and `]`, as they are not written
	// with the port in the same field).
	//
	// Clients MUST be prepared to accept IPv6 addresses in the forms shown in
	// RFC4291.
	IpAddr string `protobuf:"bytes,1,opt,name=ip_addr,json=ipAddr,proto3" json:"ip_addr,omitempty"`
	Port   uint32 `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
}

// Deprecated: Use ConnectionEndpoint.ProtoReflect.Descriptor instead.
func (*ConnectionEndpoint) Descriptor() ([]byte, []int) {
	return file_node_proto_rawDescGZIP(), []int{0}
}

func (x *ConnectionEndpoint) GetIpAddr() string {
	if x != nil {
		return x.IpAddr
	}
	return ""
}

func (x *ConnectionEndpoint) GetPort() uint32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (*ConnectionEndpoint) ProtoMessage() {}

func (x *ConnectionEndpoint) ProtoReflect() protoreflect.Message {
	mi := &file_node_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *ConnectionEndpoint) Reset() {
	*x = ConnectionEndpoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectionEndpoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

type IPv4InterfaceConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IpAddr        string   `protobuf:"bytes,1,opt,name=ip_addr,json=ipAddr,proto3" json:"ip_addr,omitempty"`
	GatewayIpAddr []string `protobuf:"bytes,2,rep,name=gateway_ip_addr,json=gatewayIpAddr,proto3" json:"gateway_ip_addr,omitempty"`
	PrefixLength  uint32   `protobuf:"varint,3,opt,name=prefix_length,json=prefixLength,proto3" json:"prefix_length,omitempty"`
}

// Deprecated: Use IPv4InterfaceConfig.ProtoReflect.Descriptor instead.
func (*IPv4InterfaceConfig) Descriptor() ([]byte, []int) {
	return file_node_proto_rawDescGZIP(), []int{1}
}

func (x *IPv4InterfaceConfig) GetGatewayIpAddr() []string {
	if x != nil {
		return x.GatewayIpAddr
	}
	return nil
}

func (x *IPv4InterfaceConfig) GetIpAddr() string {
	if x != nil {
		return x.IpAddr
	}
	return ""
}

func (x *IPv4InterfaceConfig) GetPrefixLength() uint32 {
	if x != nil {
		return x.PrefixLength
	}
	return 0
}

func (*IPv4InterfaceConfig) ProtoMessage() {}

func (x *IPv4InterfaceConfig) ProtoReflect() protoreflect.Message {
	mi := &file_node_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *IPv4InterfaceConfig) Reset() {
	*x = IPv4InterfaceConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IPv4InterfaceConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

// A node: one machine running a replica instance.
type NodeRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The endpoint where this node receives xnet messages.
	Xnet *ConnectionEndpoint `protobuf:"bytes,5,opt,name=xnet,proto3" json:"xnet,omitempty"`
	// The endpoint where this node receives http requests.
	Http *ConnectionEndpoint `protobuf:"bytes,6,opt,name=http,proto3" json:"http,omitempty"`
	// The id of the node operator that added this node.
	NodeOperatorId []byte `protobuf:"bytes,15,opt,name=node_operator_id,json=nodeOperatorId,proto3" json:"node_operator_id,omitempty"`
	// The SEV-SNP chip_identifier for this node.
	ChipId []byte `protobuf:"bytes,16,opt,name=chip_id,json=chipId,proto3,oneof" json:"chip_id,omitempty"`
	// ID of the HostOS version to run.
	HostosVersionId *string `protobuf:"bytes,17,opt,name=hostos_version_id,json=hostosVersionId,proto3,oneof" json:"hostos_version_id,omitempty"`
	// IPv4 interface configuration
	PublicIpv4Config *IPv4InterfaceConfig `protobuf:"bytes,18,opt,name=public_ipv4_config,json=publicIpv4Config,proto3,oneof" json:"public_ipv4_config,omitempty"`
	// Domain name, which resolves into Node's IPv4 and IPv6.
	// If a Node is to be converted into the ApiBoundaryNode, the domain field should be set.
	Domain *string `protobuf:"bytes,19,opt,name=domain,proto3,oneof" json:"domain,omitempty"`
}

// Deprecated: Use NodeRecord.ProtoReflect.Descriptor instead.
func (*NodeRecord) Descriptor() ([]byte, []int) {
	return file_node_proto_rawDescGZIP(), []int{2}
}

func (x *NodeRecord) GetChipId() []byte {
	if x != nil {
		return x.ChipId
	}
	return nil
}

func (x *NodeRecord) GetDomain() string {
	if x != nil && x.Domain != nil {
		return *x.Domain
	}
	return ""
}

func (x *NodeRecord) GetHostosVersionId() string {
	if x != nil && x.HostosVersionId != nil {
		return *x.HostosVersionId
	}
	return ""
}

func (x *NodeRecord) GetHttp() *ConnectionEndpoint {
	if x != nil {
		return x.Http
	}
	return nil
}

func (x *NodeRecord) GetNodeOperatorId() []byte {
	if x != nil {
		return x.NodeOperatorId
	}
	return nil
}

func (x *NodeRecord) GetPublicIpv4Config() *IPv4InterfaceConfig {
	if x != nil {
		return x.PublicIpv4Config
	}
	return nil
}

func (x *NodeRecord) GetXnet() *ConnectionEndpoint {
	if x != nil {
		return x.Xnet
	}
	return nil
}

func (*NodeRecord) ProtoMessage() {}
func (x *NodeRecord) ProtoReflect() protoreflect.Message {
	mi := &file_node_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (x *NodeRecord) Reset() {
	*x = NodeRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func init() { file_node_proto_init() }
