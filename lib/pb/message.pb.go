// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: lib/pb/message.proto

package message

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

type Result_Status int32

const (
	Result_Forbidden       Result_Status = 0
	Result_InvalidSelector Result_Status = 1
)

// Enum value maps for Result_Status.
var (
	Result_Status_name = map[int32]string{
		0: "Forbidden",
		1: "InvalidSelector",
	}
	Result_Status_value = map[string]int32{
		"Forbidden":       0,
		"InvalidSelector": 1,
	}
)

func (x Result_Status) Enum() *Result_Status {
	p := new(Result_Status)
	*p = x
	return p
}

func (x Result_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Result_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_lib_pb_message_proto_enumTypes[0].Descriptor()
}

func (Result_Status) Type() protoreflect.EnumType {
	return &file_lib_pb_message_proto_enumTypes[0]
}

func (x Result_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Result_Status.Descriptor instead.
func (Result_Status) EnumDescriptor() ([]byte, []int) {
	return file_lib_pb_message_proto_rawDescGZIP(), []int{1, 0}
}

type End_Status int32

const (
	End_OK     End_Status = 0
	End_Broken End_Status = 1
)

// Enum value maps for End_Status.
var (
	End_Status_name = map[int32]string{
		0: "OK",
		1: "Broken",
	}
	End_Status_value = map[string]int32{
		"OK":     0,
		"Broken": 1,
	}
)

func (x End_Status) Enum() *End_Status {
	p := new(End_Status)
	*p = x
	return p
}

func (x End_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (End_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_lib_pb_message_proto_enumTypes[1].Descriptor()
}

func (End_Status) Type() protoreflect.EnumType {
	return &file_lib_pb_message_proto_enumTypes[1]
}

func (x End_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use End_Status.Descriptor instead.
func (End_Status) EnumDescriptor() ([]byte, []int) {
	return file_lib_pb_message_proto_rawDescGZIP(), []int{5, 0}
}

// Message ...
type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// Types that are assignable to Content:
	//	*Message_Result
	//	*Message_Auth
	//	*Message_Connect
	//	*Message_Packet
	Content isMessage_Content `protobuf_oneof:"content"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lib_pb_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_lib_pb_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_lib_pb_message_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (m *Message) GetContent() isMessage_Content {
	if m != nil {
		return m.Content
	}
	return nil
}

func (x *Message) GetResult() *Result {
	if x, ok := x.GetContent().(*Message_Result); ok {
		return x.Result
	}
	return nil
}

func (x *Message) GetAuth() *Auth {
	if x, ok := x.GetContent().(*Message_Auth); ok {
		return x.Auth
	}
	return nil
}

func (x *Message) GetConnect() *Connect {
	if x, ok := x.GetContent().(*Message_Connect); ok {
		return x.Connect
	}
	return nil
}

func (x *Message) GetPacket() *Packet {
	if x, ok := x.GetContent().(*Message_Packet); ok {
		return x.Packet
	}
	return nil
}

type isMessage_Content interface {
	isMessage_Content()
}

type Message_Result struct {
	Result *Result `protobuf:"bytes,2,opt,name=result,proto3,oneof"`
}

type Message_Auth struct {
	Auth *Auth `protobuf:"bytes,3,opt,name=auth,proto3,oneof"`
}

type Message_Connect struct {
	Connect *Connect `protobuf:"bytes,4,opt,name=connect,proto3,oneof"`
}

type Message_Packet struct {
	Packet *Packet `protobuf:"bytes,5,opt,name=packet,proto3,oneof"`
}

func (*Message_Result) isMessage_Content() {}

func (*Message_Auth) isMessage_Content() {}

func (*Message_Connect) isMessage_Content() {}

func (*Message_Packet) isMessage_Content() {}

// Result ...
type Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *Result) Reset() {
	*x = Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lib_pb_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Result) ProtoMessage() {}

func (x *Result) ProtoReflect() protoreflect.Message {
	mi := &file_lib_pb_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Result.ProtoReflect.Descriptor instead.
func (*Result) Descriptor() ([]byte, []int) {
	return file_lib_pb_message_proto_rawDescGZIP(), []int{1}
}

func (x *Result) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

// Auth ...
type Auth struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token     string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Domain    string `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain,omitempty"`
	Subdomain string `protobuf:"bytes,3,opt,name=subdomain,proto3" json:"subdomain,omitempty"`
}

func (x *Auth) Reset() {
	*x = Auth{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lib_pb_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Auth) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Auth) ProtoMessage() {}

func (x *Auth) ProtoReflect() protoreflect.Message {
	mi := &file_lib_pb_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Auth.ProtoReflect.Descriptor instead.
func (*Auth) Descriptor() ([]byte, []int) {
	return file_lib_pb_message_proto_rawDescGZIP(), []int{2}
}

func (x *Auth) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *Auth) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *Auth) GetSubdomain() string {
	if x != nil {
		return x.Subdomain
	}
	return ""
}

// Connect ...
type Connect struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Connect) Reset() {
	*x = Connect{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lib_pb_message_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Connect) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Connect) ProtoMessage() {}

func (x *Connect) ProtoReflect() protoreflect.Message {
	mi := &file_lib_pb_message_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Connect.ProtoReflect.Descriptor instead.
func (*Connect) Descriptor() ([]byte, []int) {
	return file_lib_pb_message_proto_rawDescGZIP(), []int{3}
}

// Packet ...
type Packet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Packet) Reset() {
	*x = Packet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lib_pb_message_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Packet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Packet) ProtoMessage() {}

func (x *Packet) ProtoReflect() protoreflect.Message {
	mi := &file_lib_pb_message_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Packet.ProtoReflect.Descriptor instead.
func (*Packet) Descriptor() ([]byte, []int) {
	return file_lib_pb_message_proto_rawDescGZIP(), []int{4}
}

func (x *Packet) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

// End connection
type End struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *End) Reset() {
	*x = End{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lib_pb_message_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *End) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*End) ProtoMessage() {}

func (x *End) ProtoReflect() protoreflect.Message {
	mi := &file_lib_pb_message_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use End.ProtoReflect.Descriptor instead.
func (*End) Descriptor() ([]byte, []int) {
	return file_lib_pb_message_proto_rawDescGZIP(), []int{5}
}

func (x *End) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

var File_lib_pb_message_proto protoreflect.FileDescriptor

var file_lib_pb_message_proto_rawDesc = []byte{
	0x0a, 0x14, 0x6c, 0x69, 0x62, 0x2f, 0x70, 0x62, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xad, 0x01, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x21, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x07, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x48, 0x00, 0x52, 0x06, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1b, 0x0a, 0x04, 0x61, 0x75, 0x74, 0x68, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x48, 0x00, 0x52, 0x04, 0x61, 0x75,
	0x74, 0x68, 0x12, 0x24, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x48, 0x00, 0x52,
	0x07, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x12, 0x21, 0x0a, 0x06, 0x70, 0x61, 0x63, 0x6b,
	0x65, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x65,
	0x74, 0x48, 0x00, 0x52, 0x06, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x42, 0x09, 0x0a, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x58, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0x2c, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0d, 0x0a, 0x09,
	0x46, 0x6f, 0x72, 0x62, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x49,
	0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x10, 0x01,
	0x22, 0x52, 0x0a, 0x04, 0x41, 0x75, 0x74, 0x68, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x16,
	0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x75, 0x62, 0x64, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x75, 0x62, 0x64, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x22, 0x09, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x22,
	0x1c, 0x0a, 0x06, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x45, 0x0a,
	0x03, 0x45, 0x6e, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x1c, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4b, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x42, 0x72, 0x6f, 0x6b,
	0x65, 0x6e, 0x10, 0x01, 0x42, 0x12, 0x5a, 0x10, 0x2e, 0x2f, 0x6c, 0x69, 0x62, 0x2f, 0x70, 0x62,
	0x3b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_lib_pb_message_proto_rawDescOnce sync.Once
	file_lib_pb_message_proto_rawDescData = file_lib_pb_message_proto_rawDesc
)

func file_lib_pb_message_proto_rawDescGZIP() []byte {
	file_lib_pb_message_proto_rawDescOnce.Do(func() {
		file_lib_pb_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_lib_pb_message_proto_rawDescData)
	})
	return file_lib_pb_message_proto_rawDescData
}

var file_lib_pb_message_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_lib_pb_message_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_lib_pb_message_proto_goTypes = []interface{}{
	(Result_Status)(0), // 0: Result.Status
	(End_Status)(0),    // 1: End.Status
	(*Message)(nil),    // 2: Message
	(*Result)(nil),     // 3: Result
	(*Auth)(nil),       // 4: Auth
	(*Connect)(nil),    // 5: Connect
	(*Packet)(nil),     // 6: Packet
	(*End)(nil),        // 7: End
}
var file_lib_pb_message_proto_depIdxs = []int32{
	3, // 0: Message.result:type_name -> Result
	4, // 1: Message.auth:type_name -> Auth
	5, // 2: Message.connect:type_name -> Connect
	6, // 3: Message.packet:type_name -> Packet
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_lib_pb_message_proto_init() }
func file_lib_pb_message_proto_init() {
	if File_lib_pb_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_lib_pb_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
		file_lib_pb_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Result); i {
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
		file_lib_pb_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Auth); i {
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
		file_lib_pb_message_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Connect); i {
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
		file_lib_pb_message_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Packet); i {
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
		file_lib_pb_message_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*End); i {
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
	file_lib_pb_message_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Message_Result)(nil),
		(*Message_Auth)(nil),
		(*Message_Connect)(nil),
		(*Message_Packet)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_lib_pb_message_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_lib_pb_message_proto_goTypes,
		DependencyIndexes: file_lib_pb_message_proto_depIdxs,
		EnumInfos:         file_lib_pb_message_proto_enumTypes,
		MessageInfos:      file_lib_pb_message_proto_msgTypes,
	}.Build()
	File_lib_pb_message_proto = out.File
	file_lib_pb_message_proto_rawDesc = nil
	file_lib_pb_message_proto_goTypes = nil
	file_lib_pb_message_proto_depIdxs = nil
}
