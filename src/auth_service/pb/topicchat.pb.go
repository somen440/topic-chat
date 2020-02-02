// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/topicchat.proto

package topicchat

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7c4a8088c4bf324, []int{0}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type User struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	LoggedIn             bool     `protobuf:"varint,3,opt,name=loggedIn,proto3" json:"loggedIn,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7c4a8088c4bf324, []int{1}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetLoggedIn() bool {
	if m != nil {
		return m.LoggedIn
	}
	return false
}

type Topic struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Topic) Reset()         { *m = Topic{} }
func (m *Topic) String() string { return proto.CompactTextString(m) }
func (*Topic) ProtoMessage()    {}
func (*Topic) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7c4a8088c4bf324, []int{2}
}

func (m *Topic) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Topic.Unmarshal(m, b)
}
func (m *Topic) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Topic.Marshal(b, m, deterministic)
}
func (m *Topic) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Topic.Merge(m, src)
}
func (m *Topic) XXX_Size() int {
	return xxx_messageInfo_Topic.Size(m)
}
func (m *Topic) XXX_DiscardUnknown() {
	xxx_messageInfo_Topic.DiscardUnknown(m)
}

var xxx_messageInfo_Topic proto.InternalMessageInfo

func (m *Topic) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Topic) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type JoinRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JoinRequest) Reset()         { *m = JoinRequest{} }
func (m *JoinRequest) String() string { return proto.CompactTextString(m) }
func (*JoinRequest) ProtoMessage()    {}
func (*JoinRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7c4a8088c4bf324, []int{3}
}

func (m *JoinRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JoinRequest.Unmarshal(m, b)
}
func (m *JoinRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JoinRequest.Marshal(b, m, deterministic)
}
func (m *JoinRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JoinRequest.Merge(m, src)
}
func (m *JoinRequest) XXX_Size() int {
	return xxx_messageInfo_JoinRequest.Size(m)
}
func (m *JoinRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_JoinRequest.DiscardUnknown(m)
}

var xxx_messageInfo_JoinRequest proto.InternalMessageInfo

func (m *JoinRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type LoggedInRequest struct {
	UserId               string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoggedInRequest) Reset()         { *m = LoggedInRequest{} }
func (m *LoggedInRequest) String() string { return proto.CompactTextString(m) }
func (*LoggedInRequest) ProtoMessage()    {}
func (*LoggedInRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7c4a8088c4bf324, []int{4}
}

func (m *LoggedInRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoggedInRequest.Unmarshal(m, b)
}
func (m *LoggedInRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoggedInRequest.Marshal(b, m, deterministic)
}
func (m *LoggedInRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoggedInRequest.Merge(m, src)
}
func (m *LoggedInRequest) XXX_Size() int {
	return xxx_messageInfo_LoggedInRequest.Size(m)
}
func (m *LoggedInRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoggedInRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoggedInRequest proto.InternalMessageInfo

func (m *LoggedInRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type SignoutRequest struct {
	UserId               string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignoutRequest) Reset()         { *m = SignoutRequest{} }
func (m *SignoutRequest) String() string { return proto.CompactTextString(m) }
func (*SignoutRequest) ProtoMessage()    {}
func (*SignoutRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7c4a8088c4bf324, []int{5}
}

func (m *SignoutRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignoutRequest.Unmarshal(m, b)
}
func (m *SignoutRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignoutRequest.Marshal(b, m, deterministic)
}
func (m *SignoutRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignoutRequest.Merge(m, src)
}
func (m *SignoutRequest) XXX_Size() int {
	return xxx_messageInfo_SignoutRequest.Size(m)
}
func (m *SignoutRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SignoutRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SignoutRequest proto.InternalMessageInfo

func (m *SignoutRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type RoomJoinRequest struct {
	TopicId              string   `protobuf:"bytes,1,opt,name=topic_id,json=topicId,proto3" json:"topic_id,omitempty"`
	User                 *User    `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RoomJoinRequest) Reset()         { *m = RoomJoinRequest{} }
func (m *RoomJoinRequest) String() string { return proto.CompactTextString(m) }
func (*RoomJoinRequest) ProtoMessage()    {}
func (*RoomJoinRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7c4a8088c4bf324, []int{6}
}

func (m *RoomJoinRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoomJoinRequest.Unmarshal(m, b)
}
func (m *RoomJoinRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoomJoinRequest.Marshal(b, m, deterministic)
}
func (m *RoomJoinRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoomJoinRequest.Merge(m, src)
}
func (m *RoomJoinRequest) XXX_Size() int {
	return xxx_messageInfo_RoomJoinRequest.Size(m)
}
func (m *RoomJoinRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RoomJoinRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RoomJoinRequest proto.InternalMessageInfo

func (m *RoomJoinRequest) GetTopicId() string {
	if m != nil {
		return m.TopicId
	}
	return ""
}

func (m *RoomJoinRequest) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type ListTopicsResponse struct {
	Topics               []*Topic `protobuf:"bytes,1,rep,name=topics,proto3" json:"topics,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListTopicsResponse) Reset()         { *m = ListTopicsResponse{} }
func (m *ListTopicsResponse) String() string { return proto.CompactTextString(m) }
func (*ListTopicsResponse) ProtoMessage()    {}
func (*ListTopicsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7c4a8088c4bf324, []int{7}
}

func (m *ListTopicsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListTopicsResponse.Unmarshal(m, b)
}
func (m *ListTopicsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListTopicsResponse.Marshal(b, m, deterministic)
}
func (m *ListTopicsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListTopicsResponse.Merge(m, src)
}
func (m *ListTopicsResponse) XXX_Size() int {
	return xxx_messageInfo_ListTopicsResponse.Size(m)
}
func (m *ListTopicsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListTopicsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListTopicsResponse proto.InternalMessageInfo

func (m *ListTopicsResponse) GetTopics() []*Topic {
	if m != nil {
		return m.Topics
	}
	return nil
}

type GetTopicRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTopicRequest) Reset()         { *m = GetTopicRequest{} }
func (m *GetTopicRequest) String() string { return proto.CompactTextString(m) }
func (*GetTopicRequest) ProtoMessage()    {}
func (*GetTopicRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7c4a8088c4bf324, []int{8}
}

func (m *GetTopicRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTopicRequest.Unmarshal(m, b)
}
func (m *GetTopicRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTopicRequest.Marshal(b, m, deterministic)
}
func (m *GetTopicRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTopicRequest.Merge(m, src)
}
func (m *GetTopicRequest) XXX_Size() int {
	return xxx_messageInfo_GetTopicRequest.Size(m)
}
func (m *GetTopicRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTopicRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTopicRequest proto.InternalMessageInfo

func (m *GetTopicRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterType((*Empty)(nil), "topicchat.Empty")
	proto.RegisterType((*User)(nil), "topicchat.User")
	proto.RegisterType((*Topic)(nil), "topicchat.Topic")
	proto.RegisterType((*JoinRequest)(nil), "topicchat.JoinRequest")
	proto.RegisterType((*LoggedInRequest)(nil), "topicchat.LoggedInRequest")
	proto.RegisterType((*SignoutRequest)(nil), "topicchat.SignoutRequest")
	proto.RegisterType((*RoomJoinRequest)(nil), "topicchat.RoomJoinRequest")
	proto.RegisterType((*ListTopicsResponse)(nil), "topicchat.ListTopicsResponse")
	proto.RegisterType((*GetTopicRequest)(nil), "topicchat.GetTopicRequest")
}

func init() { proto.RegisterFile("pb/topicchat.proto", fileDescriptor_f7c4a8088c4bf324) }

var fileDescriptor_f7c4a8088c4bf324 = []byte{
	// 410 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0x4f, 0xcf, 0x93, 0x40,
	0x10, 0xc6, 0xa1, 0xa5, 0x85, 0x0e, 0x49, 0x31, 0x63, 0xa2, 0x94, 0xc4, 0xa4, 0x5d, 0x2f, 0xf8,
	0x27, 0x35, 0x62, 0x62, 0xd4, 0x83, 0xc6, 0x18, 0x35, 0x35, 0xbd, 0x48, 0xf5, 0x6c, 0x68, 0x59,
	0x29, 0x49, 0xcb, 0x22, 0xbb, 0x98, 0xf8, 0x35, 0xfc, 0x2e, 0x7e, 0x3f, 0xc3, 0x16, 0xe8, 0x96,
	0x92, 0xbc, 0xef, 0xad, 0xb3, 0xf3, 0xcc, 0x33, 0x33, 0xbf, 0x29, 0x80, 0xf9, 0xf6, 0x99, 0x60,
	0x79, 0xba, 0xdb, 0xed, 0x23, 0xb1, 0xcc, 0x0b, 0x26, 0x18, 0x4e, 0xda, 0x07, 0x62, 0xc2, 0xe8,
	0xe3, 0x31, 0x17, 0x7f, 0xc8, 0x27, 0x30, 0xbe, 0x73, 0x5a, 0xe0, 0x14, 0x06, 0x69, 0xec, 0xea,
	0x73, 0xdd, 0x9f, 0x84, 0x83, 0x34, 0x46, 0x04, 0x23, 0x8b, 0x8e, 0xd4, 0x1d, 0xc8, 0x17, 0xf9,
	0x1b, 0x3d, 0xb0, 0x0e, 0x2c, 0x49, 0x68, 0xbc, 0xca, 0xdc, 0xe1, 0x5c, 0xf7, 0xad, 0xb0, 0x8d,
	0xc9, 0x13, 0x18, 0x7d, 0xab, 0xdc, 0x6f, 0x63, 0x44, 0x16, 0x60, 0x7f, 0x61, 0x69, 0x16, 0xd2,
	0x5f, 0x25, 0xe5, 0xa2, 0x95, 0xe8, 0x8a, 0xe4, 0x31, 0x38, 0xeb, 0xda, 0xbb, 0x91, 0xdd, 0x07,
	0xb3, 0xe4, 0xb4, 0xf8, 0xd1, 0xda, 0x8f, 0xab, 0x70, 0x15, 0x93, 0x47, 0x30, 0xdd, 0xa4, 0x49,
	0xc6, 0x4a, 0x71, 0xa3, 0xf4, 0x2b, 0x38, 0x21, 0x63, 0x47, 0xb5, 0xfb, 0x0c, 0x2c, 0xc9, 0xe5,
	0x2c, 0x36, 0x65, 0xbc, 0x8a, 0xf1, 0x21, 0x18, 0x55, 0x9d, 0x9c, 0xdd, 0x0e, 0x9c, 0xe5, 0x19,
	0x68, 0xc5, 0x2c, 0x94, 0x49, 0xf2, 0x16, 0x70, 0x9d, 0x72, 0x21, 0xb7, 0xe7, 0x21, 0xe5, 0x39,
	0xcb, 0x38, 0x45, 0x1f, 0xc6, 0x52, 0xcd, 0x5d, 0x7d, 0x3e, 0xf4, 0xed, 0xe0, 0x8e, 0x52, 0x2c,
	0xa5, 0x61, 0x9d, 0x27, 0x0b, 0x70, 0x3e, 0xd3, 0x53, 0x79, 0x33, 0x52, 0x87, 0x61, 0xf0, 0x4f,
	0x07, 0xfb, 0x7d, 0x29, 0xf6, 0x1b, 0x5a, 0xfc, 0x4e, 0x77, 0x14, 0x9f, 0x83, 0x51, 0x6d, 0x80,
	0xf7, 0x14, 0x53, 0x65, 0x25, 0xaf, 0x3b, 0x29, 0xd1, 0xf0, 0x35, 0x58, 0x0d, 0x4f, 0xf4, 0x94,
	0x74, 0x07, 0x72, 0x5f, 0xe9, 0x2b, 0x30, 0x6b, 0xbc, 0x38, 0x53, 0xb2, 0x97, 0xc8, 0x3d, 0x75,
	0xc1, 0xd3, 0x5f, 0x4b, 0x0b, 0x38, 0xd8, 0x15, 0xed, 0x66, 0xec, 0x97, 0xf5, 0xd8, 0x6a, 0xff,
	0xce, 0x35, 0xfa, 0x6c, 0xf0, 0x29, 0x18, 0x6b, 0xfa, 0x53, 0xe0, 0x55, 0xae, 0xb7, 0xe9, 0x5f,
	0x1d, 0xee, 0x4a, 0x9a, 0x1f, 0x22, 0x11, 0x1d, 0x58, 0xd2, 0x74, 0x7f, 0x07, 0x70, 0xbe, 0x53,
	0x8f, 0xd7, 0x03, 0x95, 0xca, 0xd5, 0x41, 0x89, 0x86, 0x6f, 0xc0, 0x6a, 0x0e, 0x75, 0xb1, 0x42,
	0xe7, 0x7a, 0xde, 0xd5, 0xa9, 0x89, 0xb6, 0x1d, 0xcb, 0x2f, 0xf0, 0xc5, 0xff, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x1e, 0xb6, 0xac, 0x83, 0x97, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AuthServiceClient is the client API for AuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthServiceClient interface {
	Join(ctx context.Context, in *JoinRequest, opts ...grpc.CallOption) (*User, error)
	LoggedIn(ctx context.Context, in *LoggedInRequest, opts ...grpc.CallOption) (*User, error)
	Signout(ctx context.Context, in *SignoutRequest, opts ...grpc.CallOption) (*Empty, error)
}

type authServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthServiceClient(cc grpc.ClientConnInterface) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) Join(ctx context.Context, in *JoinRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/topicchat.AuthService/Join", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) LoggedIn(ctx context.Context, in *LoggedInRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/topicchat.AuthService/LoggedIn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) Signout(ctx context.Context, in *SignoutRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/topicchat.AuthService/Signout", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceServer is the server API for AuthService service.
type AuthServiceServer interface {
	Join(context.Context, *JoinRequest) (*User, error)
	LoggedIn(context.Context, *LoggedInRequest) (*User, error)
	Signout(context.Context, *SignoutRequest) (*Empty, error)
}

// UnimplementedAuthServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAuthServiceServer struct {
}

func (*UnimplementedAuthServiceServer) Join(ctx context.Context, req *JoinRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Join not implemented")
}
func (*UnimplementedAuthServiceServer) LoggedIn(ctx context.Context, req *LoggedInRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoggedIn not implemented")
}
func (*UnimplementedAuthServiceServer) Signout(ctx context.Context, req *SignoutRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Signout not implemented")
}

func RegisterAuthServiceServer(s *grpc.Server, srv AuthServiceServer) {
	s.RegisterService(&_AuthService_serviceDesc, srv)
}

func _AuthService_Join_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Join(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/topicchat.AuthService/Join",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Join(ctx, req.(*JoinRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_LoggedIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoggedInRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).LoggedIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/topicchat.AuthService/LoggedIn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).LoggedIn(ctx, req.(*LoggedInRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_Signout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignoutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Signout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/topicchat.AuthService/Signout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Signout(ctx, req.(*SignoutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "topicchat.AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Join",
			Handler:    _AuthService_Join_Handler,
		},
		{
			MethodName: "LoggedIn",
			Handler:    _AuthService_LoggedIn_Handler,
		},
		{
			MethodName: "Signout",
			Handler:    _AuthService_Signout_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/topicchat.proto",
}

// RoomServiceClient is the client API for RoomService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RoomServiceClient interface {
	Join(ctx context.Context, in *RoomJoinRequest, opts ...grpc.CallOption) (*Empty, error)
	Left(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
}

type roomServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRoomServiceClient(cc grpc.ClientConnInterface) RoomServiceClient {
	return &roomServiceClient{cc}
}

func (c *roomServiceClient) Join(ctx context.Context, in *RoomJoinRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/topicchat.RoomService/Join", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomServiceClient) Left(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/topicchat.RoomService/Left", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RoomServiceServer is the server API for RoomService service.
type RoomServiceServer interface {
	Join(context.Context, *RoomJoinRequest) (*Empty, error)
	Left(context.Context, *Empty) (*Empty, error)
}

// UnimplementedRoomServiceServer can be embedded to have forward compatible implementations.
type UnimplementedRoomServiceServer struct {
}

func (*UnimplementedRoomServiceServer) Join(ctx context.Context, req *RoomJoinRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Join not implemented")
}
func (*UnimplementedRoomServiceServer) Left(ctx context.Context, req *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Left not implemented")
}

func RegisterRoomServiceServer(s *grpc.Server, srv RoomServiceServer) {
	s.RegisterService(&_RoomService_serviceDesc, srv)
}

func _RoomService_Join_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoomJoinRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomServiceServer).Join(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/topicchat.RoomService/Join",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomServiceServer).Join(ctx, req.(*RoomJoinRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoomService_Left_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomServiceServer).Left(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/topicchat.RoomService/Left",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomServiceServer).Left(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _RoomService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "topicchat.RoomService",
	HandlerType: (*RoomServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Join",
			Handler:    _RoomService_Join_Handler,
		},
		{
			MethodName: "Left",
			Handler:    _RoomService_Left_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/topicchat.proto",
}

// TopicCatalogServiceClient is the client API for TopicCatalogService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TopicCatalogServiceClient interface {
	ListTopics(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ListTopicsResponse, error)
	GetTopic(ctx context.Context, in *GetTopicRequest, opts ...grpc.CallOption) (*Topic, error)
}

type topicCatalogServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTopicCatalogServiceClient(cc grpc.ClientConnInterface) TopicCatalogServiceClient {
	return &topicCatalogServiceClient{cc}
}

func (c *topicCatalogServiceClient) ListTopics(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ListTopicsResponse, error) {
	out := new(ListTopicsResponse)
	err := c.cc.Invoke(ctx, "/topicchat.TopicCatalogService/ListTopics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *topicCatalogServiceClient) GetTopic(ctx context.Context, in *GetTopicRequest, opts ...grpc.CallOption) (*Topic, error) {
	out := new(Topic)
	err := c.cc.Invoke(ctx, "/topicchat.TopicCatalogService/GetTopic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TopicCatalogServiceServer is the server API for TopicCatalogService service.
type TopicCatalogServiceServer interface {
	ListTopics(context.Context, *Empty) (*ListTopicsResponse, error)
	GetTopic(context.Context, *GetTopicRequest) (*Topic, error)
}

// UnimplementedTopicCatalogServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTopicCatalogServiceServer struct {
}

func (*UnimplementedTopicCatalogServiceServer) ListTopics(ctx context.Context, req *Empty) (*ListTopicsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTopics not implemented")
}
func (*UnimplementedTopicCatalogServiceServer) GetTopic(ctx context.Context, req *GetTopicRequest) (*Topic, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTopic not implemented")
}

func RegisterTopicCatalogServiceServer(s *grpc.Server, srv TopicCatalogServiceServer) {
	s.RegisterService(&_TopicCatalogService_serviceDesc, srv)
}

func _TopicCatalogService_ListTopics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TopicCatalogServiceServer).ListTopics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/topicchat.TopicCatalogService/ListTopics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TopicCatalogServiceServer).ListTopics(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _TopicCatalogService_GetTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTopicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TopicCatalogServiceServer).GetTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/topicchat.TopicCatalogService/GetTopic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TopicCatalogServiceServer).GetTopic(ctx, req.(*GetTopicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TopicCatalogService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "topicchat.TopicCatalogService",
	HandlerType: (*TopicCatalogServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListTopics",
			Handler:    _TopicCatalogService_ListTopics_Handler,
		},
		{
			MethodName: "GetTopic",
			Handler:    _TopicCatalogService_GetTopic_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/topicchat.proto",
}
