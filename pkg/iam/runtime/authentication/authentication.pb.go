// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.24.3
// source: authentication/authentication.proto

package authentication

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

type AuthenticateSubjectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// credential is the literal credential for a subject (such as a bearer token) passed to the
	// application with no transformations applied.
	Credential string `protobuf:"bytes,1,opt,name=credential,proto3" json:"credential,omitempty"`
}

func (x *AuthenticateSubjectRequest) Reset() {
	*x = AuthenticateSubjectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authentication_authentication_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthenticateSubjectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthenticateSubjectRequest) ProtoMessage() {}

func (x *AuthenticateSubjectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_authentication_authentication_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthenticateSubjectRequest.ProtoReflect.Descriptor instead.
func (*AuthenticateSubjectRequest) Descriptor() ([]byte, []int) {
	return file_authentication_authentication_proto_rawDescGZIP(), []int{0}
}

func (x *AuthenticateSubjectRequest) GetCredential() string {
	if x != nil {
		return x.Credential
	}
	return ""
}

type AuthenticateSubjectResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// subject_claims is a map of claims about the subject (such as ID and scopes).
	SubjectClaims map[string]string `protobuf:"bytes,1,rep,name=subject_claims,json=subjectClaims,proto3" json:"subject_claims,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *AuthenticateSubjectResponse) Reset() {
	*x = AuthenticateSubjectResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authentication_authentication_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthenticateSubjectResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthenticateSubjectResponse) ProtoMessage() {}

func (x *AuthenticateSubjectResponse) ProtoReflect() protoreflect.Message {
	mi := &file_authentication_authentication_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthenticateSubjectResponse.ProtoReflect.Descriptor instead.
func (*AuthenticateSubjectResponse) Descriptor() ([]byte, []int) {
	return file_authentication_authentication_proto_rawDescGZIP(), []int{1}
}

func (x *AuthenticateSubjectResponse) GetSubjectClaims() map[string]string {
	if x != nil {
		return x.SubjectClaims
	}
	return nil
}

var File_authentication_authentication_proto protoreflect.FileDescriptor

var file_authentication_authentication_proto_rawDesc = []byte{
	0x0a, 0x23, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x69,
	0x61, 0x6d, 0x2e, 0x76, 0x31, 0x22, 0x3c, 0x0a, 0x1a, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74,
	0x69, 0x63, 0x61, 0x74, 0x65, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x61, 0x6c, 0x22, 0xc6, 0x01, 0x0a, 0x1b, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69,
	0x63, 0x61, 0x74, 0x65, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x65, 0x0a, 0x0e, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x63,
	0x6c, 0x61, 0x69, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3e, 0x2e, 0x72, 0x75,
	0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74,
	0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x43, 0x6c, 0x61, 0x69, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0d, 0x73, 0x75, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x73, 0x1a, 0x40, 0x0a, 0x12, 0x53, 0x75,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x32, 0x82, 0x01, 0x0a,
	0x0e, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x70, 0x0a, 0x13, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x53,
	0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x2a, 0x2e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65,
	0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69,
	0x63, 0x61, 0x74, 0x65, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x69, 0x61, 0x6d,
	0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65,
	0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x41, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6d, 0x65, 0x74, 0x61, 0x6c, 0x2d, 0x74, 0x6f, 0x6f, 0x6c, 0x62, 0x6f, 0x78, 0x2f, 0x69, 0x61,
	0x6d, 0x2d, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x72, 0x75,
	0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_authentication_authentication_proto_rawDescOnce sync.Once
	file_authentication_authentication_proto_rawDescData = file_authentication_authentication_proto_rawDesc
)

func file_authentication_authentication_proto_rawDescGZIP() []byte {
	file_authentication_authentication_proto_rawDescOnce.Do(func() {
		file_authentication_authentication_proto_rawDescData = protoimpl.X.CompressGZIP(file_authentication_authentication_proto_rawDescData)
	})
	return file_authentication_authentication_proto_rawDescData
}

var file_authentication_authentication_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_authentication_authentication_proto_goTypes = []interface{}{
	(*AuthenticateSubjectRequest)(nil),  // 0: runtime.iam.v1.AuthenticateSubjectRequest
	(*AuthenticateSubjectResponse)(nil), // 1: runtime.iam.v1.AuthenticateSubjectResponse
	nil,                                 // 2: runtime.iam.v1.AuthenticateSubjectResponse.SubjectClaimsEntry
}
var file_authentication_authentication_proto_depIdxs = []int32{
	2, // 0: runtime.iam.v1.AuthenticateSubjectResponse.subject_claims:type_name -> runtime.iam.v1.AuthenticateSubjectResponse.SubjectClaimsEntry
	0, // 1: runtime.iam.v1.Authentication.AuthenticateSubject:input_type -> runtime.iam.v1.AuthenticateSubjectRequest
	1, // 2: runtime.iam.v1.Authentication.AuthenticateSubject:output_type -> runtime.iam.v1.AuthenticateSubjectResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_authentication_authentication_proto_init() }
func file_authentication_authentication_proto_init() {
	if File_authentication_authentication_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_authentication_authentication_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthenticateSubjectRequest); i {
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
		file_authentication_authentication_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthenticateSubjectResponse); i {
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
			RawDescriptor: file_authentication_authentication_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_authentication_authentication_proto_goTypes,
		DependencyIndexes: file_authentication_authentication_proto_depIdxs,
		MessageInfos:      file_authentication_authentication_proto_msgTypes,
	}.Build()
	File_authentication_authentication_proto = out.File
	file_authentication_authentication_proto_rawDesc = nil
	file_authentication_authentication_proto_goTypes = nil
	file_authentication_authentication_proto_depIdxs = nil
}
