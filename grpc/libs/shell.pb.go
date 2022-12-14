// protoc --proto_path .\grpc\pb\ --go_out=. --go-grpc_out=. .\grpc\pb\shell.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: shell.proto

package libs

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

type Ping struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State bool `protobuf:"varint,1,opt,name=state,proto3" json:"state,omitempty"`
}

func (x *Ping) Reset() {
	*x = Ping{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shell_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ping) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ping) ProtoMessage() {}

func (x *Ping) ProtoReflect() protoreflect.Message {
	mi := &file_shell_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ping.ProtoReflect.Descriptor instead.
func (*Ping) Descriptor() ([]byte, []int) {
	return file_shell_proto_rawDescGZIP(), []int{0}
}

func (x *Ping) GetState() bool {
	if x != nil {
		return x.State
	}
	return false
}

type Pong struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State   bool   `protobuf:"varint,1,opt,name=state,proto3" json:"state,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Pong) Reset() {
	*x = Pong{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shell_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pong) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pong) ProtoMessage() {}

func (x *Pong) ProtoReflect() protoreflect.Message {
	mi := &file_shell_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pong.ProtoReflect.Descriptor instead.
func (*Pong) Descriptor() ([]byte, []int) {
	return file_shell_proto_rawDescGZIP(), []int{1}
}

func (x *Pong) GetState() bool {
	if x != nil {
		return x.State
	}
	return false
}

func (x *Pong) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ListState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State bool `protobuf:"varint,1,opt,name=state,proto3" json:"state,omitempty"`
}

func (x *ListState) Reset() {
	*x = ListState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shell_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListState) ProtoMessage() {}

func (x *ListState) ProtoReflect() protoreflect.Message {
	mi := &file_shell_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListState.ProtoReflect.Descriptor instead.
func (*ListState) Descriptor() ([]byte, []int) {
	return file_shell_proto_rawDescGZIP(), []int{2}
}

func (x *ListState) GetState() bool {
	if x != nil {
		return x.State
	}
	return false
}

type ListResults struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Scripts []string `protobuf:"bytes,1,rep,name=scripts,proto3" json:"scripts,omitempty"`
}

func (x *ListResults) Reset() {
	*x = ListResults{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shell_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListResults) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResults) ProtoMessage() {}

func (x *ListResults) ProtoReflect() protoreflect.Message {
	mi := &file_shell_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListResults.ProtoReflect.Descriptor instead.
func (*ListResults) Descriptor() ([]byte, []int) {
	return file_shell_proto_rawDescGZIP(), []int{3}
}

func (x *ListResults) GetScripts() []string {
	if x != nil {
		return x.Scripts
	}
	return nil
}

type RunMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ScriptName  string   `protobuf:"bytes,1,opt,name=script_name,json=scriptName,proto3" json:"script_name,omitempty"`
	ScriptParas []string `protobuf:"bytes,2,rep,name=script_paras,json=scriptParas,proto3" json:"script_paras,omitempty"`
}

func (x *RunMeta) Reset() {
	*x = RunMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shell_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunMeta) ProtoMessage() {}

func (x *RunMeta) ProtoReflect() protoreflect.Message {
	mi := &file_shell_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunMeta.ProtoReflect.Descriptor instead.
func (*RunMeta) Descriptor() ([]byte, []int) {
	return file_shell_proto_rawDescGZIP(), []int{4}
}

func (x *RunMeta) GetScriptName() string {
	if x != nil {
		return x.ScriptName
	}
	return ""
}

func (x *RunMeta) GetScriptParas() []string {
	if x != nil {
		return x.ScriptParas
	}
	return nil
}

type RunResults struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results string `protobuf:"bytes,1,opt,name=results,proto3" json:"results,omitempty"`
}

func (x *RunResults) Reset() {
	*x = RunResults{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shell_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunResults) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunResults) ProtoMessage() {}

func (x *RunResults) ProtoReflect() protoreflect.Message {
	mi := &file_shell_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunResults.ProtoReflect.Descriptor instead.
func (*RunResults) Descriptor() ([]byte, []int) {
	return file_shell_proto_rawDescGZIP(), []int{5}
}

func (x *RunResults) GetResults() string {
	if x != nil {
		return x.Results
	}
	return ""
}

type FileMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Md5  string `protobuf:"bytes,1,opt,name=md5,proto3" json:"md5,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Size int64  `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
	File []byte `protobuf:"bytes,4,opt,name=file,proto3" json:"file,omitempty"`
}

func (x *FileMeta) Reset() {
	*x = FileMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shell_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileMeta) ProtoMessage() {}

func (x *FileMeta) ProtoReflect() protoreflect.Message {
	mi := &file_shell_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileMeta.ProtoReflect.Descriptor instead.
func (*FileMeta) Descriptor() ([]byte, []int) {
	return file_shell_proto_rawDescGZIP(), []int{6}
}

func (x *FileMeta) GetMd5() string {
	if x != nil {
		return x.Md5
	}
	return ""
}

func (x *FileMeta) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FileMeta) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *FileMeta) GetFile() []byte {
	if x != nil {
		return x.File
	}
	return nil
}

type FileResults struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State bool `protobuf:"varint,1,opt,name=state,proto3" json:"state,omitempty"`
}

func (x *FileResults) Reset() {
	*x = FileResults{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shell_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileResults) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileResults) ProtoMessage() {}

func (x *FileResults) ProtoReflect() protoreflect.Message {
	mi := &file_shell_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileResults.ProtoReflect.Descriptor instead.
func (*FileResults) Descriptor() ([]byte, []int) {
	return file_shell_proto_rawDescGZIP(), []int{7}
}

func (x *FileResults) GetState() bool {
	if x != nil {
		return x.State
	}
	return false
}

var File_shell_proto protoreflect.FileDescriptor

var file_shell_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x73, 0x68, 0x65, 0x6c, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70,
	0x62, 0x22, 0x1c, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22,
	0x36, 0x0a, 0x04, 0x50, 0x6f, 0x6e, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x21, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0x27, 0x0a, 0x0b, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x73, 0x22, 0x4d, 0x0a, 0x07, 0x52, 0x75, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x1f,
	0x0a, 0x0b, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x21, 0x0a, 0x0c, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x50, 0x61, 0x72,
	0x61, 0x73, 0x22, 0x26, 0x0a, 0x0a, 0x52, 0x75, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73,
	0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x22, 0x58, 0x0a, 0x08, 0x46, 0x69,
	0x6c, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x64, 0x35, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x64, 0x35, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04,
	0x66, 0x69, 0x6c, 0x65, 0x22, 0x23, 0x0a, 0x0b, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x32, 0xbf, 0x01, 0x0a, 0x0c, 0x53, 0x68,
	0x65, 0x6c, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x23, 0x0a, 0x0b, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x50,
	0x69, 0x6e, 0x67, 0x1a, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x6f, 0x6e, 0x67, 0x22, 0x00, 0x12,
	0x2f, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x73, 0x12, 0x0d,
	0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x1a, 0x0f, 0x2e,
	0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x22, 0x00,
	0x12, 0x2a, 0x0a, 0x09, 0x52, 0x75, 0x6e, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x12, 0x0b, 0x2e,
	0x70, 0x62, 0x2e, 0x52, 0x75, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x1a, 0x0e, 0x2e, 0x70, 0x62, 0x2e,
	0x52, 0x75, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x22, 0x00, 0x12, 0x2d, 0x0a, 0x0a,
	0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x0c, 0x2e, 0x70, 0x62, 0x2e,
	0x46, 0x69, 0x6c, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x1a, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69,
	0x6c, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x22, 0x00, 0x42, 0x11, 0x5a, 0x0f, 0x67,
	0x72, 0x70, 0x63, 0x2f, 0x6c, 0x69, 0x62, 0x73, 0x2f, 0x3b, 0x6c, 0x69, 0x62, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_shell_proto_rawDescOnce sync.Once
	file_shell_proto_rawDescData = file_shell_proto_rawDesc
)

func file_shell_proto_rawDescGZIP() []byte {
	file_shell_proto_rawDescOnce.Do(func() {
		file_shell_proto_rawDescData = protoimpl.X.CompressGZIP(file_shell_proto_rawDescData)
	})
	return file_shell_proto_rawDescData
}

var file_shell_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_shell_proto_goTypes = []interface{}{
	(*Ping)(nil),        // 0: pb.Ping
	(*Pong)(nil),        // 1: pb.Pong
	(*ListState)(nil),   // 2: pb.ListState
	(*ListResults)(nil), // 3: pb.ListResults
	(*RunMeta)(nil),     // 4: pb.RunMeta
	(*RunResults)(nil),  // 5: pb.RunResults
	(*FileMeta)(nil),    // 6: pb.FileMeta
	(*FileResults)(nil), // 7: pb.FileResults
}
var file_shell_proto_depIdxs = []int32{
	0, // 0: pb.ShellService.ServerCheck:input_type -> pb.Ping
	2, // 1: pb.ShellService.ListScripts:input_type -> pb.ListState
	4, // 2: pb.ShellService.RunScript:input_type -> pb.RunMeta
	6, // 3: pb.ShellService.UploadFile:input_type -> pb.FileMeta
	1, // 4: pb.ShellService.ServerCheck:output_type -> pb.Pong
	3, // 5: pb.ShellService.ListScripts:output_type -> pb.ListResults
	5, // 6: pb.ShellService.RunScript:output_type -> pb.RunResults
	7, // 7: pb.ShellService.UploadFile:output_type -> pb.FileResults
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_shell_proto_init() }
func file_shell_proto_init() {
	if File_shell_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_shell_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ping); i {
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
		file_shell_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pong); i {
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
		file_shell_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListState); i {
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
		file_shell_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListResults); i {
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
		file_shell_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunMeta); i {
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
		file_shell_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunResults); i {
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
		file_shell_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileMeta); i {
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
		file_shell_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileResults); i {
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
			RawDescriptor: file_shell_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_shell_proto_goTypes,
		DependencyIndexes: file_shell_proto_depIdxs,
		MessageInfos:      file_shell_proto_msgTypes,
	}.Build()
	File_shell_proto = out.File
	file_shell_proto_rawDesc = nil
	file_shell_proto_goTypes = nil
	file_shell_proto_depIdxs = nil
}
