// Protobuf message that are used in the local_store.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.3
// 	protoc        v5.29.2
// source: local.proto

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

// Enum value maps for MutationType.
var (
	MutationType_name = map[int32]string{
		0: "INVALID_STATE",
		1: "SET",
		2: "UNSET",
	}
	MutationType_value = map[string]int32{
		"INVALID_STATE": 0,
		"SET":           1,
		"UNSET":         2,
	}
)

var (
	file_local_proto_rawDescOnce sync.Once
	file_local_proto_rawDescData = file_local_proto_rawDesc
)

var File_local_proto protoreflect.FileDescriptor

var file_local_proto_depIdxs = []int32{
	2, // 0: ic_registry_common.pb.local_store.v1.ChangelogEntry.key_mutations:type_name -> ic_registry_common.pb.local_store.v1.KeyMutation
	0, // 1: ic_registry_common.pb.local_store.v1.KeyMutation.mutation_type:type_name -> ic_registry_common.pb.local_store.v1.MutationType
	1, // 2: ic_registry_common.pb.local_store.v1.Delta.changelog:type_name -> ic_registry_common.pb.local_store.v1.ChangelogEntry
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

var file_local_proto_enumTypes = make([]protoimpl.EnumInfo, 1)

var file_local_proto_goTypes = []any{
	(MutationType)(0),      // 0: ic_registry_common.pb.local_store.v1.MutationType
	(*ChangelogEntry)(nil), // 1: ic_registry_common.pb.local_store.v1.ChangelogEntry
	(*KeyMutation)(nil),    // 2: ic_registry_common.pb.local_store.v1.KeyMutation
	(*CertifiedTime)(nil),  // 3: ic_registry_common.pb.local_store.v1.CertifiedTime
	(*Delta)(nil),          // 4: ic_registry_common.pb.local_store.v1.Delta
}

var file_local_proto_msgTypes = make([]protoimpl.MessageInfo, 4)

var file_local_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x24, 0x69,
	0x63, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x70, 0x62, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x2e, 0x76, 0x31, 0x22, 0x68, 0x0a, 0x0e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x6c, 0x6f, 0x67,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x56, 0x0a, 0x0d, 0x6b, 0x65, 0x79, 0x5f, 0x6d, 0x75, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x69,
	0x63, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x70, 0x62, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x4b, 0x65, 0x79, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x0c, 0x6b, 0x65, 0x79, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x8e, 0x01,
	0x0a, 0x0b, 0x4b, 0x65, 0x79, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x57, 0x0a, 0x0d, 0x6d, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x32, 0x2e, 0x69,
	0x63, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x70, 0x62, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x0c, 0x6d, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x22, 0x39,
	0x0a, 0x0d, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x28, 0x0a, 0x10, 0x75, 0x6e, 0x69, 0x78, 0x5f, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x5f, 0x6e, 0x61,
	0x6e, 0x6f, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0e, 0x75, 0x6e, 0x69, 0x78, 0x45,
	0x70, 0x6f, 0x63, 0x68, 0x4e, 0x61, 0x6e, 0x6f, 0x73, 0x22, 0x86, 0x01, 0x0a, 0x05, 0x44, 0x65,
	0x6c, 0x74, 0x61, 0x12, 0x29, 0x0a, 0x10, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x5f,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0f, 0x72,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x52,
	0x0a, 0x09, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x6c, 0x6f, 0x67, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x34, 0x2e, 0x69, 0x63, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x5f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x62, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x6c,
	0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x6c,
	0x6f, 0x67, 0x2a, 0x35, 0x0a, 0x0c, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x11, 0x0a, 0x0d, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x53, 0x54,
	0x41, 0x54, 0x45, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x53, 0x45, 0x54, 0x10, 0x01, 0x12, 0x09,
	0x0a, 0x05, 0x55, 0x4e, 0x53, 0x45, 0x54, 0x10, 0x02, 0x42, 0x0a, 0x5a, 0x08, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

func file_local_proto_rawDescGZIP() []byte {
	file_local_proto_rawDescOnce.Do(func() {
		file_local_proto_rawDescData = protoimpl.X.CompressGZIP(file_local_proto_rawDescData)
	})
	return file_local_proto_rawDescData
}

// The time when the last certified update was successfully received.
type CertifiedTime struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Number of nano seconds since UNIX EPOCH
	UnixEpochNanos uint64 `protobuf:"varint,1,opt,name=unix_epoch_nanos,json=unixEpochNanos,proto3" json:"unix_epoch_nanos,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

// Deprecated: Use CertifiedTime.ProtoReflect.Descriptor instead.
func (*CertifiedTime) Descriptor() ([]byte, []int) {
	return file_local_proto_rawDescGZIP(), []int{2}
}

func (x *CertifiedTime) GetUnixEpochNanos() uint64 {
	if x != nil {
		return x.UnixEpochNanos
	}
	return 0
}

func (*CertifiedTime) ProtoMessage() {}

func (x *CertifiedTime) ProtoReflect() protoreflect.Message {
	mi := &file_local_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *CertifiedTime) Reset() {
	*x = CertifiedTime{}
	mi := &file_local_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CertifiedTime) String() string {
	return protoimpl.X.MessageStringOf(x)
}

// Set of all mutations that, when applied to the registry at version v,
// produce the registry at version v+1
type ChangelogEntry struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The default, an empty list, is _invalid_ here.
	KeyMutations  []*KeyMutation `protobuf:"bytes,1,rep,name=key_mutations,json=keyMutations,proto3" json:"key_mutations,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

// Deprecated: Use ChangelogEntry.ProtoReflect.Descriptor instead.
func (*ChangelogEntry) Descriptor() ([]byte, []int) {
	return file_local_proto_rawDescGZIP(), []int{0}
}

func (x *ChangelogEntry) GetKeyMutations() []*KeyMutation {
	if x != nil {
		return x.KeyMutations
	}
	return nil
}

func (*ChangelogEntry) ProtoMessage() {}

func (x *ChangelogEntry) ProtoReflect() protoreflect.Message {
	mi := &file_local_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *ChangelogEntry) Reset() {
	*x = ChangelogEntry{}
	mi := &file_local_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ChangelogEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

// A changelog that is applicable at a specific registry version.
type Delta struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	RegistryVersion uint64                 `protobuf:"varint,1,opt,name=registry_version,json=registryVersion,proto3" json:"registry_version,omitempty"`
	Changelog       []*ChangelogEntry      `protobuf:"bytes,2,rep,name=changelog,proto3" json:"changelog,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

// Deprecated: Use Delta.ProtoReflect.Descriptor instead.
func (*Delta) Descriptor() ([]byte, []int) {
	return file_local_proto_rawDescGZIP(), []int{3}
}

func (x *Delta) GetChangelog() []*ChangelogEntry {
	if x != nil {
		return x.Changelog
	}
	return nil
}

func (x *Delta) GetRegistryVersion() uint64 {
	if x != nil {
		return x.RegistryVersion
	}
	return 0
}

func (*Delta) ProtoMessage() {}

func (x *Delta) ProtoReflect() protoreflect.Message {
	mi := &file_local_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *Delta) Reset() {
	*x = Delta{}
	mi := &file_local_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Delta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

// A mutation of a single key.
type KeyMutation struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Key.
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// Protobuf encoded value.
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	// If this is `UNSET`, `value` must assume the default value.
	MutationType  MutationType `protobuf:"varint,3,opt,name=mutation_type,json=mutationType,proto3,enum=ic_registry_common.pb.local_store.v1.MutationType" json:"mutation_type,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

// Deprecated: Use KeyMutation.ProtoReflect.Descriptor instead.
func (*KeyMutation) Descriptor() ([]byte, []int) {
	return file_local_proto_rawDescGZIP(), []int{1}
}

func (x *KeyMutation) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *KeyMutation) GetMutationType() MutationType {
	if x != nil {
		return x.MutationType
	}
	return MutationType_INVALID_STATE
}

func (x *KeyMutation) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

func (*KeyMutation) ProtoMessage() {}

func (x *KeyMutation) ProtoReflect() protoreflect.Message {
	mi := &file_local_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *KeyMutation) Reset() {
	*x = KeyMutation{}
	mi := &file_local_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *KeyMutation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

type MutationType int32

const (
	// Illegal state.
	MutationType_INVALID_STATE MutationType = 0
	// The value was SET in this delta.
	MutationType_SET MutationType = 1
	// The value was UNSET in this delta.
	MutationType_UNSET MutationType = 2
)

func (MutationType) Descriptor() protoreflect.EnumDescriptor {
	return file_local_proto_enumTypes[0].Descriptor()
}

func (x MutationType) Enum() *MutationType {
	p := new(MutationType)
	*p = x
	return p
}

// Deprecated: Use MutationType.Descriptor instead.
func (MutationType) EnumDescriptor() ([]byte, []int) {
	return file_local_proto_rawDescGZIP(), []int{0}
}
func (x MutationType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}
func (x MutationType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}
func (MutationType) Type() protoreflect.EnumType {
	return &file_local_proto_enumTypes[0]
}

func init() { file_local_proto_init() }
func file_local_proto_init() {
	if File_local_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_local_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_local_proto_goTypes,
		DependencyIndexes: file_local_proto_depIdxs,
		EnumInfos:         file_local_proto_enumTypes,
		MessageInfos:      file_local_proto_msgTypes,
	}.Build()
	File_local_proto = out.File
	file_local_proto_rawDesc = nil
	file_local_proto_goTypes = nil
	file_local_proto_depIdxs = nil
}
