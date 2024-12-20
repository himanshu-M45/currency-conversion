// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.2
// source: proto/currency.proto

package proto

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

type ConvertRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SenderCurrencyType   string  `protobuf:"bytes,1,opt,name=senderCurrencyType,proto3" json:"senderCurrencyType,omitempty"`
	ReceiverCurrencyType string  `protobuf:"bytes,2,opt,name=receiverCurrencyType,proto3" json:"receiverCurrencyType,omitempty"`
	Amount               float64 `protobuf:"fixed64,3,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *ConvertRequest) Reset() {
	*x = ConvertRequest{}
	mi := &file_proto_currency_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConvertRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConvertRequest) ProtoMessage() {}

func (x *ConvertRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_currency_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConvertRequest.ProtoReflect.Descriptor instead.
func (*ConvertRequest) Descriptor() ([]byte, []int) {
	return file_proto_currency_proto_rawDescGZIP(), []int{0}
}

func (x *ConvertRequest) GetSenderCurrencyType() string {
	if x != nil {
		return x.SenderCurrencyType
	}
	return ""
}

func (x *ConvertRequest) GetReceiverCurrencyType() string {
	if x != nil {
		return x.ReceiverCurrencyType
	}
	return ""
}

func (x *ConvertRequest) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type ConvertResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConvertedAmount float64 `protobuf:"fixed64,1,opt,name=convertedAmount,proto3" json:"convertedAmount,omitempty"`
}

func (x *ConvertResponse) Reset() {
	*x = ConvertResponse{}
	mi := &file_proto_currency_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConvertResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConvertResponse) ProtoMessage() {}

func (x *ConvertResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_currency_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConvertResponse.ProtoReflect.Descriptor instead.
func (*ConvertResponse) Descriptor() ([]byte, []int) {
	return file_proto_currency_proto_rawDescGZIP(), []int{1}
}

func (x *ConvertResponse) GetConvertedAmount() float64 {
	if x != nil {
		return x.ConvertedAmount
	}
	return 0
}

var File_proto_currency_proto protoreflect.FileDescriptor

var file_proto_currency_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79,
	0x22, 0x8c, 0x01, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x12, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x43, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x63, 0x79, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x12, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x32, 0x0a, 0x14, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x43,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x14, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x43, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22,
	0x3b, 0x0a, 0x0f, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x65, 0x64, 0x41,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0f, 0x63, 0x6f, 0x6e,
	0x76, 0x65, 0x72, 0x74, 0x65, 0x64, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x32, 0x54, 0x0a, 0x12,
	0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x3e, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x12, 0x18, 0x2e,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x79, 0x2e, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x11, 0x5a, 0x0f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x63, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_currency_proto_rawDescOnce sync.Once
	file_proto_currency_proto_rawDescData = file_proto_currency_proto_rawDesc
)

func file_proto_currency_proto_rawDescGZIP() []byte {
	file_proto_currency_proto_rawDescOnce.Do(func() {
		file_proto_currency_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_currency_proto_rawDescData)
	})
	return file_proto_currency_proto_rawDescData
}

var file_proto_currency_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_currency_proto_goTypes = []any{
	(*ConvertRequest)(nil),  // 0: currency.ConvertRequest
	(*ConvertResponse)(nil), // 1: currency.ConvertResponse
}
var file_proto_currency_proto_depIdxs = []int32{
	0, // 0: currency.CurrencyConversion.Convert:input_type -> currency.ConvertRequest
	1, // 1: currency.CurrencyConversion.Convert:output_type -> currency.ConvertResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_currency_proto_init() }
func file_proto_currency_proto_init() {
	if File_proto_currency_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_currency_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_currency_proto_goTypes,
		DependencyIndexes: file_proto_currency_proto_depIdxs,
		MessageInfos:      file_proto_currency_proto_msgTypes,
	}.Build()
	File_proto_currency_proto = out.File
	file_proto_currency_proto_rawDesc = nil
	file_proto_currency_proto_goTypes = nil
	file_proto_currency_proto_depIdxs = nil
}
