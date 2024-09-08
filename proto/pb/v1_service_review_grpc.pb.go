// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ReviewServiceClient is the client API for ReviewService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReviewServiceClient interface {
	CreateSnippetTopic(ctx context.Context, in *CreateSnippetTopicRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	SubmitNewSnippetVersion(ctx context.Context, in *SubmitNewSnippetVersionRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	CommentOnSnippetTopic(ctx context.Context, in *CommentOnSnippetTopicRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	FetchSnippetTopics(ctx context.Context, in *FetchSnippetTopicsRequest, opts ...grpc.CallOption) (*FetchSnippetTopicsReply, error)
	FetchSnippets(ctx context.Context, in *FetchSnippetsRequest, opts ...grpc.CallOption) (*FetchSnippetsReply, error)
	FetchComments(ctx context.Context, in *FetchCommentsRequest, opts ...grpc.CallOption) (*FetchCommentsReply, error)
}

type reviewServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewReviewServiceClient(cc grpc.ClientConnInterface) ReviewServiceClient {
	return &reviewServiceClient{cc}
}

func (c *reviewServiceClient) CreateSnippetTopic(ctx context.Context, in *CreateSnippetTopicRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/proto.ReviewService/CreateSnippetTopic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reviewServiceClient) SubmitNewSnippetVersion(ctx context.Context, in *SubmitNewSnippetVersionRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/proto.ReviewService/SubmitNewSnippetVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reviewServiceClient) CommentOnSnippetTopic(ctx context.Context, in *CommentOnSnippetTopicRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/proto.ReviewService/CommentOnSnippetTopic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reviewServiceClient) FetchSnippetTopics(ctx context.Context, in *FetchSnippetTopicsRequest, opts ...grpc.CallOption) (*FetchSnippetTopicsReply, error) {
	out := new(FetchSnippetTopicsReply)
	err := c.cc.Invoke(ctx, "/proto.ReviewService/FetchSnippetTopics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reviewServiceClient) FetchSnippets(ctx context.Context, in *FetchSnippetsRequest, opts ...grpc.CallOption) (*FetchSnippetsReply, error) {
	out := new(FetchSnippetsReply)
	err := c.cc.Invoke(ctx, "/proto.ReviewService/FetchSnippets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reviewServiceClient) FetchComments(ctx context.Context, in *FetchCommentsRequest, opts ...grpc.CallOption) (*FetchCommentsReply, error) {
	out := new(FetchCommentsReply)
	err := c.cc.Invoke(ctx, "/proto.ReviewService/FetchComments", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReviewServiceServer is the server API for ReviewService service.
// All implementations must embed UnimplementedReviewServiceServer
// for forward compatibility
type ReviewServiceServer interface {
	CreateSnippetTopic(context.Context, *CreateSnippetTopicRequest) (*emptypb.Empty, error)
	SubmitNewSnippetVersion(context.Context, *SubmitNewSnippetVersionRequest) (*emptypb.Empty, error)
	CommentOnSnippetTopic(context.Context, *CommentOnSnippetTopicRequest) (*emptypb.Empty, error)
	FetchSnippetTopics(context.Context, *FetchSnippetTopicsRequest) (*FetchSnippetTopicsReply, error)
	FetchSnippets(context.Context, *FetchSnippetsRequest) (*FetchSnippetsReply, error)
	FetchComments(context.Context, *FetchCommentsRequest) (*FetchCommentsReply, error)
	mustEmbedUnimplementedReviewServiceServer()
}

// UnimplementedReviewServiceServer must be embedded to have forward compatible implementations.
type UnimplementedReviewServiceServer struct {
}

func (UnimplementedReviewServiceServer) CreateSnippetTopic(context.Context, *CreateSnippetTopicRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSnippetTopic not implemented")
}
func (UnimplementedReviewServiceServer) SubmitNewSnippetVersion(context.Context, *SubmitNewSnippetVersionRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitNewSnippetVersion not implemented")
}
func (UnimplementedReviewServiceServer) CommentOnSnippetTopic(context.Context, *CommentOnSnippetTopicRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CommentOnSnippetTopic not implemented")
}
func (UnimplementedReviewServiceServer) FetchSnippetTopics(context.Context, *FetchSnippetTopicsRequest) (*FetchSnippetTopicsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchSnippetTopics not implemented")
}
func (UnimplementedReviewServiceServer) FetchSnippets(context.Context, *FetchSnippetsRequest) (*FetchSnippetsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchSnippets not implemented")
}
func (UnimplementedReviewServiceServer) FetchComments(context.Context, *FetchCommentsRequest) (*FetchCommentsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchComments not implemented")
}
func (UnimplementedReviewServiceServer) mustEmbedUnimplementedReviewServiceServer() {}

// UnsafeReviewServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReviewServiceServer will
// result in compilation errors.
type UnsafeReviewServiceServer interface {
	mustEmbedUnimplementedReviewServiceServer()
}

func RegisterReviewServiceServer(s grpc.ServiceRegistrar, srv ReviewServiceServer) {
	s.RegisterService(&ReviewService_ServiceDesc, srv)
}

func _ReviewService_CreateSnippetTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSnippetTopicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewServiceServer).CreateSnippetTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ReviewService/CreateSnippetTopic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewServiceServer).CreateSnippetTopic(ctx, req.(*CreateSnippetTopicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReviewService_SubmitNewSnippetVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitNewSnippetVersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewServiceServer).SubmitNewSnippetVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ReviewService/SubmitNewSnippetVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewServiceServer).SubmitNewSnippetVersion(ctx, req.(*SubmitNewSnippetVersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReviewService_CommentOnSnippetTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommentOnSnippetTopicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewServiceServer).CommentOnSnippetTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ReviewService/CommentOnSnippetTopic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewServiceServer).CommentOnSnippetTopic(ctx, req.(*CommentOnSnippetTopicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReviewService_FetchSnippetTopics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchSnippetTopicsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewServiceServer).FetchSnippetTopics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ReviewService/FetchSnippetTopics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewServiceServer).FetchSnippetTopics(ctx, req.(*FetchSnippetTopicsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReviewService_FetchSnippets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchSnippetsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewServiceServer).FetchSnippets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ReviewService/FetchSnippets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewServiceServer).FetchSnippets(ctx, req.(*FetchSnippetsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReviewService_FetchComments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchCommentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewServiceServer).FetchComments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ReviewService/FetchComments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewServiceServer).FetchComments(ctx, req.(*FetchCommentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ReviewService_ServiceDesc is the grpc.ServiceDesc for ReviewService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ReviewService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.ReviewService",
	HandlerType: (*ReviewServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateSnippetTopic",
			Handler:    _ReviewService_CreateSnippetTopic_Handler,
		},
		{
			MethodName: "SubmitNewSnippetVersion",
			Handler:    _ReviewService_SubmitNewSnippetVersion_Handler,
		},
		{
			MethodName: "CommentOnSnippetTopic",
			Handler:    _ReviewService_CommentOnSnippetTopic_Handler,
		},
		{
			MethodName: "FetchSnippetTopics",
			Handler:    _ReviewService_FetchSnippetTopics_Handler,
		},
		{
			MethodName: "FetchSnippets",
			Handler:    _ReviewService_FetchSnippets_Handler,
		},
		{
			MethodName: "FetchComments",
			Handler:    _ReviewService_FetchComments_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1_service_review.proto",
}
