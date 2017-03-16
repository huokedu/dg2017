// Code generated by protoc-gen-go.
// source: articles.proto
// DO NOT EDIT!

/*
Package articles is a generated protocol buffer package.

It is generated from these files:
	articles.proto

It has these top-level messages:
	CreateRequest
	CreateResponse
	GetRequest
	GetResponse
	DeleteRequest
	DeleteResponse
	ListRequest
	ListResponse
	Article
*/
package articles

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CreateRequest struct {
	Article *Article `protobuf:"bytes,1,opt,name=article" json:"article,omitempty"`
}

func (m *CreateRequest) Reset()                    { *m = CreateRequest{} }
func (m *CreateRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()               {}
func (*CreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CreateRequest) GetArticle() *Article {
	if m != nil {
		return m.Article
	}
	return nil
}

type CreateResponse struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *CreateResponse) Reset()                    { *m = CreateResponse{} }
func (m *CreateResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()               {}
func (*CreateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CreateResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *GetRequest) Reset()                    { *m = GetRequest{} }
func (m *GetRequest) String() string            { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()               {}
func (*GetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *GetRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetResponse struct {
	Article *Article `protobuf:"bytes,1,opt,name=article" json:"article,omitempty"`
}

func (m *GetResponse) Reset()                    { *m = GetResponse{} }
func (m *GetResponse) String() string            { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()               {}
func (*GetResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *GetResponse) GetArticle() *Article {
	if m != nil {
		return m.Article
	}
	return nil
}

type DeleteRequest struct {
	Ids []string `protobuf:"bytes,1,rep,name=ids" json:"ids,omitempty"`
}

func (m *DeleteRequest) Reset()                    { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()               {}
func (*DeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *DeleteRequest) GetIds() []string {
	if m != nil {
		return m.Ids
	}
	return nil
}

type DeleteResponse struct {
}

func (m *DeleteResponse) Reset()                    { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()               {}
func (*DeleteResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type ListRequest struct {
	Ids []string `protobuf:"bytes,1,rep,name=ids" json:"ids,omitempty"`
}

func (m *ListRequest) Reset()                    { *m = ListRequest{} }
func (m *ListRequest) String() string            { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()               {}
func (*ListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ListRequest) GetIds() []string {
	if m != nil {
		return m.Ids
	}
	return nil
}

type ListResponse struct {
	Articles []*Article `protobuf:"bytes,1,rep,name=articles" json:"articles,omitempty"`
}

func (m *ListResponse) Reset()                    { *m = ListResponse{} }
func (m *ListResponse) String() string            { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()               {}
func (*ListResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *ListResponse) GetArticles() []*Article {
	if m != nil {
		return m.Articles
	}
	return nil
}

type Article struct {
	Title string `protobuf:"bytes,1,opt,name=title" json:"title,omitempty"`
	Body  string `protobuf:"bytes,2,opt,name=body" json:"body,omitempty"`
}

func (m *Article) Reset()                    { *m = Article{} }
func (m *Article) String() string            { return proto.CompactTextString(m) }
func (*Article) ProtoMessage()               {}
func (*Article) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *Article) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Article) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateRequest)(nil), "articles.CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "articles.CreateResponse")
	proto.RegisterType((*GetRequest)(nil), "articles.GetRequest")
	proto.RegisterType((*GetResponse)(nil), "articles.GetResponse")
	proto.RegisterType((*DeleteRequest)(nil), "articles.DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "articles.DeleteResponse")
	proto.RegisterType((*ListRequest)(nil), "articles.ListRequest")
	proto.RegisterType((*ListResponse)(nil), "articles.ListResponse")
	proto.RegisterType((*Article)(nil), "articles.Article")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Articles service

type ArticlesClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
}

type articlesClient struct {
	cc *grpc.ClientConn
}

func NewArticlesClient(cc *grpc.ClientConn) ArticlesClient {
	return &articlesClient{cc}
}

func (c *articlesClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := grpc.Invoke(ctx, "/articles.Articles/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articlesClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := grpc.Invoke(ctx, "/articles.Articles/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articlesClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := grpc.Invoke(ctx, "/articles.Articles/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articlesClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := grpc.Invoke(ctx, "/articles.Articles/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Articles service

type ArticlesServer interface {
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	Get(context.Context, *GetRequest) (*GetResponse, error)
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
	List(context.Context, *ListRequest) (*ListResponse, error)
}

func RegisterArticlesServer(s *grpc.Server, srv ArticlesServer) {
	s.RegisterService(&_Articles_serviceDesc, srv)
}

func _Articles_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticlesServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/articles.Articles/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticlesServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Articles_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticlesServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/articles.Articles/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticlesServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Articles_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticlesServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/articles.Articles/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticlesServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Articles_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticlesServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/articles.Articles/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticlesServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Articles_serviceDesc = grpc.ServiceDesc{
	ServiceName: "articles.Articles",
	HandlerType: (*ArticlesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Articles_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Articles_Get_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Articles_Delete_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Articles_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "articles.proto",
}

func init() { proto.RegisterFile("articles.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 300 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x4f, 0x4f, 0x83, 0x40,
	0x10, 0xc5, 0x4b, 0xa9, 0xb4, 0x3c, 0x2c, 0xa9, 0x93, 0xaa, 0x84, 0x98, 0x88, 0x9c, 0x9a, 0x18,
	0x7b, 0x68, 0x4d, 0x4c, 0x8c, 0x1c, 0x8c, 0x26, 0xbd, 0x78, 0xe2, 0x1b, 0xb4, 0xb2, 0x87, 0x4d,
	0x1a, 0xa9, 0xec, 0x7a, 0xf0, 0x93, 0x7b, 0x35, 0xdd, 0x3f, 0x40, 0x51, 0x0f, 0xde, 0x66, 0x67,
	0xde, 0xdb, 0x37, 0xfb, 0x03, 0x84, 0xeb, 0x4a, 0xf2, 0xd7, 0x2d, 0x13, 0xf3, 0x5d, 0x55, 0xca,
	0x92, 0x46, 0xf6, 0x9c, 0x3e, 0x60, 0xfc, 0x54, 0xb1, 0xb5, 0x64, 0x39, 0x7b, 0xff, 0x60, 0x42,
	0xd2, 0x35, 0x86, 0x66, 0x18, 0x39, 0x89, 0x33, 0x0b, 0x16, 0x27, 0xf3, 0xda, 0xfc, 0xa8, 0x8b,
	0xdc, 0x2a, 0xd2, 0x04, 0xa1, 0x75, 0x8b, 0x5d, 0xf9, 0x26, 0x18, 0x85, 0xe8, 0xf3, 0x42, 0x39,
	0xfd, 0xbc, 0xcf, 0x8b, 0xf4, 0x02, 0x58, 0x31, 0x69, 0x2f, 0xef, 0x4e, 0xef, 0x11, 0xa8, 0xa9,
	0x31, 0xff, 0x2b, 0xfb, 0x0a, 0xe3, 0x67, 0xb6, 0x65, 0xcd, 0xe6, 0x13, 0xb8, 0xbc, 0x10, 0x91,
	0x93, 0xb8, 0x33, 0x3f, 0xdf, 0x97, 0xe9, 0x04, 0xa1, 0x95, 0xe8, 0x84, 0xf4, 0x12, 0xc1, 0x0b,
	0x17, 0xf2, 0x6f, 0x4b, 0x86, 0x63, 0x2d, 0x30, 0x2b, 0xdd, 0xa0, 0x66, 0xa5, 0x64, 0xbf, 0xee,
	0xd4, 0xe0, 0x5c, 0x62, 0x68, 0x9a, 0x34, 0xc5, 0x91, 0xe4, 0xd2, 0x3c, 0xc5, 0xcf, 0xf5, 0x81,
	0x08, 0x83, 0x4d, 0x59, 0x7c, 0x46, 0x7d, 0xd5, 0x54, 0xf5, 0xe2, 0xcb, 0xc1, 0xc8, 0xb8, 0x04,
	0x65, 0xf0, 0x34, 0x52, 0x3a, 0x6f, 0x82, 0x0e, 0x3e, 0x51, 0x1c, 0xfd, 0x1c, 0x98, 0xe7, 0xf5,
	0xe8, 0x16, 0xee, 0x8a, 0x49, 0x9a, 0x36, 0x92, 0x06, 0x7f, 0x7c, 0xda, 0xe9, 0xd6, 0xae, 0x0c,
	0x9e, 0x06, 0xd5, 0x0e, 0x3d, 0xa0, 0xdb, 0x0e, 0xed, 0x30, 0xed, 0xd1, 0x1d, 0x06, 0x7b, 0x68,
	0xd4, 0xba, 0xbf, 0x45, 0x39, 0x3e, 0xeb, 0xb6, 0xad, 0x71, 0xe3, 0xa9, 0xdf, 0x71, 0xf9, 0x1d,
	0x00, 0x00, 0xff, 0xff, 0xba, 0x1f, 0xbc, 0xe5, 0xa0, 0x02, 0x00, 0x00,
}
