// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v5.26.1
// source: tweet.proto

package pb

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

type CreateTweetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *CreateTweetRequest) Reset() {
	*x = CreateTweetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tweet_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTweetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTweetRequest) ProtoMessage() {}

func (x *CreateTweetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tweet_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTweetRequest.ProtoReflect.Descriptor instead.
func (*CreateTweetRequest) Descriptor() ([]byte, []int) {
	return file_tweet_proto_rawDescGZIP(), []int{0}
}

func (x *CreateTweetRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type AuthorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID       string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=Username,proto3" json:"Username,omitempty"`
}

func (x *AuthorResponse) Reset() {
	*x = AuthorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tweet_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorResponse) ProtoMessage() {}

func (x *AuthorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tweet_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorResponse.ProtoReflect.Descriptor instead.
func (*AuthorResponse) Descriptor() ([]byte, []int) {
	return file_tweet_proto_rawDescGZIP(), []int{1}
}

func (x *AuthorResponse) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *AuthorResponse) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type TweetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string          `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Text      string          `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	CreatedAt string          `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Author    *AuthorResponse `protobuf:"bytes,4,opt,name=author,proto3" json:"author,omitempty"`
}

func (x *TweetResponse) Reset() {
	*x = TweetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tweet_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TweetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TweetResponse) ProtoMessage() {}

func (x *TweetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tweet_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TweetResponse.ProtoReflect.Descriptor instead.
func (*TweetResponse) Descriptor() ([]byte, []int) {
	return file_tweet_proto_rawDescGZIP(), []int{2}
}

func (x *TweetResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *TweetResponse) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *TweetResponse) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *TweetResponse) GetAuthor() *AuthorResponse {
	if x != nil {
		return x.Author
	}
	return nil
}

var File_tweet_proto protoreflect.FileDescriptor

var file_tweet_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x74, 0x77, 0x65, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x28, 0x0a,
	0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x77, 0x65, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x22, 0x3c, 0x0a, 0x0e, 0x41, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x55, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x7b, 0x0a, 0x0d, 0x54, 0x77, 0x65, 0x65, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x27, 0x0a, 0x06, 0x61, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x41, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x32, 0x3d, 0x0a, 0x0c, 0x54, 0x77, 0x65, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x2d, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x13, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x77, 0x65, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x0e, 0x2e, 0x54, 0x77, 0x65, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x08, 0x5a, 0x06, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_tweet_proto_rawDescOnce sync.Once
	file_tweet_proto_rawDescData = file_tweet_proto_rawDesc
)

func file_tweet_proto_rawDescGZIP() []byte {
	file_tweet_proto_rawDescOnce.Do(func() {
		file_tweet_proto_rawDescData = protoimpl.X.CompressGZIP(file_tweet_proto_rawDescData)
	})
	return file_tweet_proto_rawDescData
}

var file_tweet_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_tweet_proto_goTypes = []interface{}{
	(*CreateTweetRequest)(nil), // 0: CreateTweetRequest
	(*AuthorResponse)(nil),     // 1: AuthorResponse
	(*TweetResponse)(nil),      // 2: TweetResponse
}
var file_tweet_proto_depIdxs = []int32{
	1, // 0: TweetResponse.author:type_name -> AuthorResponse
	0, // 1: TweetService.Create:input_type -> CreateTweetRequest
	2, // 2: TweetService.Create:output_type -> TweetResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_tweet_proto_init() }
func file_tweet_proto_init() {
	if File_tweet_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tweet_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTweetRequest); i {
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
		file_tweet_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthorResponse); i {
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
		file_tweet_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TweetResponse); i {
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
			RawDescriptor: file_tweet_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tweet_proto_goTypes,
		DependencyIndexes: file_tweet_proto_depIdxs,
		MessageInfos:      file_tweet_proto_msgTypes,
	}.Build()
	File_tweet_proto = out.File
	file_tweet_proto_rawDesc = nil
	file_tweet_proto_goTypes = nil
	file_tweet_proto_depIdxs = nil
}
