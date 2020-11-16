// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: tensorflow/core/protobuf/tensor_bundle.proto

package core_protos_go_proto

import (
	tensor_shape_go_proto "github.com/tensorflow/tensorflow/tensorflow/go/core/framework/tensor_shape_go_proto"
	tensor_slice_go_proto "github.com/tensorflow/tensorflow/tensorflow/go/core/framework/tensor_slice_go_proto"
	types_go_proto "github.com/tensorflow/tensorflow/tensorflow/go/core/framework/types_go_proto"
	versions_go_proto "github.com/tensorflow/tensorflow/tensorflow/go/core/framework/versions_go_proto"
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

// An enum indicating the endianness of the platform that produced this
// bundle.  A bundle can only be read by a platform with matching endianness.
// Defaults to LITTLE, as most modern platforms are little-endian.
//
// Affects the binary tensor data bytes only, not the metadata in protobufs.
type BundleHeaderProto_Endianness int32

const (
	BundleHeaderProto_LITTLE BundleHeaderProto_Endianness = 0
	BundleHeaderProto_BIG    BundleHeaderProto_Endianness = 1
)

// Enum value maps for BundleHeaderProto_Endianness.
var (
	BundleHeaderProto_Endianness_name = map[int32]string{
		0: "LITTLE",
		1: "BIG",
	}
	BundleHeaderProto_Endianness_value = map[string]int32{
		"LITTLE": 0,
		"BIG":    1,
	}
)

func (x BundleHeaderProto_Endianness) Enum() *BundleHeaderProto_Endianness {
	p := new(BundleHeaderProto_Endianness)
	*p = x
	return p
}

func (x BundleHeaderProto_Endianness) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BundleHeaderProto_Endianness) Descriptor() protoreflect.EnumDescriptor {
	return file_tensorflow_core_protobuf_tensor_bundle_proto_enumTypes[0].Descriptor()
}

func (BundleHeaderProto_Endianness) Type() protoreflect.EnumType {
	return &file_tensorflow_core_protobuf_tensor_bundle_proto_enumTypes[0]
}

func (x BundleHeaderProto_Endianness) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BundleHeaderProto_Endianness.Descriptor instead.
func (BundleHeaderProto_Endianness) EnumDescriptor() ([]byte, []int) {
	return file_tensorflow_core_protobuf_tensor_bundle_proto_rawDescGZIP(), []int{0, 0}
}

// Special header that is associated with a bundle.
//
// TODO(zongheng,zhifengc): maybe in the future, we can add information about
// which binary produced this checkpoint, timestamp, etc. Sometime, these can be
// valuable debugging information. And if needed, these can be used as defensive
// information ensuring reader (binary version) of the checkpoint and the writer
// (binary version) must match within certain range, etc.
type BundleHeaderProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Number of data files in the bundle.
	NumShards  int32                        `protobuf:"varint,1,opt,name=num_shards,json=numShards,proto3" json:"num_shards,omitempty"`
	Endianness BundleHeaderProto_Endianness `protobuf:"varint,2,opt,name=endianness,proto3,enum=tensorflow.BundleHeaderProto_Endianness" json:"endianness,omitempty"`
	// Versioning of the tensor bundle format.
	Version *versions_go_proto.VersionDef `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *BundleHeaderProto) Reset() {
	*x = BundleHeaderProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_protobuf_tensor_bundle_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BundleHeaderProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BundleHeaderProto) ProtoMessage() {}

func (x *BundleHeaderProto) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_protobuf_tensor_bundle_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BundleHeaderProto.ProtoReflect.Descriptor instead.
func (*BundleHeaderProto) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_protobuf_tensor_bundle_proto_rawDescGZIP(), []int{0}
}

func (x *BundleHeaderProto) GetNumShards() int32 {
	if x != nil {
		return x.NumShards
	}
	return 0
}

func (x *BundleHeaderProto) GetEndianness() BundleHeaderProto_Endianness {
	if x != nil {
		return x.Endianness
	}
	return BundleHeaderProto_LITTLE
}

func (x *BundleHeaderProto) GetVersion() *versions_go_proto.VersionDef {
	if x != nil {
		return x.Version
	}
	return nil
}

// Describes the metadata related to a checkpointed tensor.
type BundleEntryProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The tensor dtype and shape.
	Dtype types_go_proto.DataType                 `protobuf:"varint,1,opt,name=dtype,proto3,enum=tensorflow.DataType" json:"dtype,omitempty"`
	Shape *tensor_shape_go_proto.TensorShapeProto `protobuf:"bytes,2,opt,name=shape,proto3" json:"shape,omitempty"`
	// The binary content of the tensor lies in:
	//   File "shard_id": bytes [offset, offset + size).
	ShardId int32 `protobuf:"varint,3,opt,name=shard_id,json=shardId,proto3" json:"shard_id,omitempty"`
	Offset  int64 `protobuf:"varint,4,opt,name=offset,proto3" json:"offset,omitempty"`
	Size    int64 `protobuf:"varint,5,opt,name=size,proto3" json:"size,omitempty"`
	// The CRC32C checksum of the tensor bytes.
	Crc32C uint32 `protobuf:"fixed32,6,opt,name=crc32c,proto3" json:"crc32c,omitempty"`
	// Iff present, this entry represents a partitioned tensor.  The previous
	// fields are interpreted as follows:
	//
	//   "dtype", "shape": describe the full tensor.
	//   "shard_id", "offset", "size", "crc32c": all IGNORED.
	//      These information for each slice can be looked up in their own
	//      BundleEntryProto, keyed by each "slice_name".
	Slices []*tensor_slice_go_proto.TensorSliceProto `protobuf:"bytes,7,rep,name=slices,proto3" json:"slices,omitempty"`
}

func (x *BundleEntryProto) Reset() {
	*x = BundleEntryProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_protobuf_tensor_bundle_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BundleEntryProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BundleEntryProto) ProtoMessage() {}

func (x *BundleEntryProto) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_protobuf_tensor_bundle_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BundleEntryProto.ProtoReflect.Descriptor instead.
func (*BundleEntryProto) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_protobuf_tensor_bundle_proto_rawDescGZIP(), []int{1}
}

func (x *BundleEntryProto) GetDtype() types_go_proto.DataType {
	if x != nil {
		return x.Dtype
	}
	return types_go_proto.DataType_DT_INVALID
}

func (x *BundleEntryProto) GetShape() *tensor_shape_go_proto.TensorShapeProto {
	if x != nil {
		return x.Shape
	}
	return nil
}

func (x *BundleEntryProto) GetShardId() int32 {
	if x != nil {
		return x.ShardId
	}
	return 0
}

func (x *BundleEntryProto) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *BundleEntryProto) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *BundleEntryProto) GetCrc32C() uint32 {
	if x != nil {
		return x.Crc32C
	}
	return 0
}

func (x *BundleEntryProto) GetSlices() []*tensor_slice_go_proto.TensorSliceProto {
	if x != nil {
		return x.Slices
	}
	return nil
}

var File_tensorflow_core_protobuf_tensor_bundle_proto protoreflect.FileDescriptor

var file_tensorflow_core_protobuf_tensor_bundle_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x5f, 0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a,
	0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x1a, 0x2c, 0x74, 0x65, 0x6e, 0x73,
	0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x66, 0x72, 0x61, 0x6d,
	0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x5f, 0x73, 0x68, 0x61,
	0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77,
	0x6f, 0x72, 0x6b, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x5f, 0x73, 0x6c, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x25, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c,
	0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72,
	0x6b, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x74,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x66,
	0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd1, 0x01, 0x0a, 0x11, 0x42, 0x75, 0x6e, 0x64,
	0x6c, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x0a,
	0x0a, 0x6e, 0x75, 0x6d, 0x5f, 0x73, 0x68, 0x61, 0x72, 0x64, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x09, 0x6e, 0x75, 0x6d, 0x53, 0x68, 0x61, 0x72, 0x64, 0x73, 0x12, 0x48, 0x0a, 0x0a,
	0x65, 0x6e, 0x64, 0x69, 0x61, 0x6e, 0x6e, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x28, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x42, 0x75,
	0x6e, 0x64, 0x6c, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x45, 0x6e, 0x64, 0x69, 0x61, 0x6e, 0x6e, 0x65, 0x73, 0x73, 0x52, 0x0a, 0x65, 0x6e, 0x64, 0x69,
	0x61, 0x6e, 0x6e, 0x65, 0x73, 0x73, 0x12, 0x30, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x66, 0x52,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x21, 0x0a, 0x0a, 0x45, 0x6e, 0x64, 0x69,
	0x61, 0x6e, 0x6e, 0x65, 0x73, 0x73, 0x12, 0x0a, 0x0a, 0x06, 0x4c, 0x49, 0x54, 0x54, 0x4c, 0x45,
	0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x42, 0x49, 0x47, 0x10, 0x01, 0x22, 0x87, 0x02, 0x0a, 0x10,
	0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x2a, 0x0a, 0x05, 0x64, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x14, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x44, 0x61, 0x74,
	0x61, 0x54, 0x79, 0x70, 0x65, 0x52, 0x05, 0x64, 0x74, 0x79, 0x70, 0x65, 0x12, 0x32, 0x0a, 0x05,
	0x73, 0x68, 0x61, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x74, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x53,
	0x68, 0x61, 0x70, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x05, 0x73, 0x68, 0x61, 0x70, 0x65,
	0x12, 0x19, 0x0a, 0x08, 0x73, 0x68, 0x61, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x07, 0x73, 0x68, 0x61, 0x72, 0x64, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6f, 0x66, 0x66,
	0x73, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x72, 0x63, 0x33, 0x32,
	0x63, 0x18, 0x06, 0x20, 0x01, 0x28, 0x07, 0x52, 0x06, 0x63, 0x72, 0x63, 0x33, 0x32, 0x63, 0x12,
	0x34, 0x0a, 0x06, 0x73, 0x6c, 0x69, 0x63, 0x65, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x54, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x06, 0x73,
	0x6c, 0x69, 0x63, 0x65, 0x73, 0x42, 0x78, 0x0a, 0x13, 0x6f, 0x72, 0x67, 0x2e, 0x74, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x75, 0x74, 0x69, 0x6c, 0x42, 0x12, 0x54, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x50, 0x01, 0x5a, 0x48, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f,
	0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x5f, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xf8, 0x01, 0x01, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tensorflow_core_protobuf_tensor_bundle_proto_rawDescOnce sync.Once
	file_tensorflow_core_protobuf_tensor_bundle_proto_rawDescData = file_tensorflow_core_protobuf_tensor_bundle_proto_rawDesc
)

func file_tensorflow_core_protobuf_tensor_bundle_proto_rawDescGZIP() []byte {
	file_tensorflow_core_protobuf_tensor_bundle_proto_rawDescOnce.Do(func() {
		file_tensorflow_core_protobuf_tensor_bundle_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_core_protobuf_tensor_bundle_proto_rawDescData)
	})
	return file_tensorflow_core_protobuf_tensor_bundle_proto_rawDescData
}

var file_tensorflow_core_protobuf_tensor_bundle_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_tensorflow_core_protobuf_tensor_bundle_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_tensorflow_core_protobuf_tensor_bundle_proto_goTypes = []interface{}{
	(BundleHeaderProto_Endianness)(0),              // 0: tensorflow.BundleHeaderProto.Endianness
	(*BundleHeaderProto)(nil),                      // 1: tensorflow.BundleHeaderProto
	(*BundleEntryProto)(nil),                       // 2: tensorflow.BundleEntryProto
	(*versions_go_proto.VersionDef)(nil),           // 3: tensorflow.VersionDef
	(types_go_proto.DataType)(0),                   // 4: tensorflow.DataType
	(*tensor_shape_go_proto.TensorShapeProto)(nil), // 5: tensorflow.TensorShapeProto
	(*tensor_slice_go_proto.TensorSliceProto)(nil), // 6: tensorflow.TensorSliceProto
}
var file_tensorflow_core_protobuf_tensor_bundle_proto_depIdxs = []int32{
	0, // 0: tensorflow.BundleHeaderProto.endianness:type_name -> tensorflow.BundleHeaderProto.Endianness
	3, // 1: tensorflow.BundleHeaderProto.version:type_name -> tensorflow.VersionDef
	4, // 2: tensorflow.BundleEntryProto.dtype:type_name -> tensorflow.DataType
	5, // 3: tensorflow.BundleEntryProto.shape:type_name -> tensorflow.TensorShapeProto
	6, // 4: tensorflow.BundleEntryProto.slices:type_name -> tensorflow.TensorSliceProto
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_tensorflow_core_protobuf_tensor_bundle_proto_init() }
func file_tensorflow_core_protobuf_tensor_bundle_proto_init() {
	if File_tensorflow_core_protobuf_tensor_bundle_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_core_protobuf_tensor_bundle_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BundleHeaderProto); i {
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
		file_tensorflow_core_protobuf_tensor_bundle_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BundleEntryProto); i {
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
			RawDescriptor: file_tensorflow_core_protobuf_tensor_bundle_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_core_protobuf_tensor_bundle_proto_goTypes,
		DependencyIndexes: file_tensorflow_core_protobuf_tensor_bundle_proto_depIdxs,
		EnumInfos:         file_tensorflow_core_protobuf_tensor_bundle_proto_enumTypes,
		MessageInfos:      file_tensorflow_core_protobuf_tensor_bundle_proto_msgTypes,
	}.Build()
	File_tensorflow_core_protobuf_tensor_bundle_proto = out.File
	file_tensorflow_core_protobuf_tensor_bundle_proto_rawDesc = nil
	file_tensorflow_core_protobuf_tensor_bundle_proto_goTypes = nil
	file_tensorflow_core_protobuf_tensor_bundle_proto_depIdxs = nil
}
