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
	ChooseSnippetTopicSolution(ctx context.Context, in *ChooseSnippetTopicSolutionRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ModifySnippetTopic(ctx context.Context, in *ModifySnippetTopicRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ModifySnippet(ctx context.Context, in *ModifySnippetRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ChooseSnippetTopicBestAnswer(ctx context.Context, in *ChooseSnippetTopicBestAnswerRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	VoteOnComment(ctx context.Context, in *VoteOnCommentRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	RemoveVoteFromComment(ctx context.Context, in *RemoveVoteFromCommentRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	RemoveCommentFromSnippetTopic(ctx context.Context, in *RemoveCommentFromSnippetTopicRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteSnippetTopic(ctx context.Context, in *DeleteSnippetTopicRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
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

func (c *reviewServiceClient) ChooseSnippetTopicSolution(ctx context.Context, in *ChooseSnippetTopicSolutionRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/proto.ReviewService/ChooseSnippetTopicSolution", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reviewServiceClient) ModifySnippetTopic(ctx context.Context, in *ModifySnippetTopicRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/proto.ReviewService/ModifySnippetTopic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reviewServiceClient) ModifySnippet(ctx context.Context, in *ModifySnippetRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/proto.ReviewService/ModifySnippet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reviewServiceClient) ChooseSnippetTopicBestAnswer(ctx context.Context, in *ChooseSnippetTopicBestAnswerRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/proto.ReviewService/ChooseSnippetTopicBestAnswer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reviewServiceClient) VoteOnComment(ctx context.Context, in *VoteOnCommentRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/proto.ReviewService/VoteOnComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reviewServiceClient) RemoveVoteFromComment(ctx context.Context, in *RemoveVoteFromCommentRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/proto.ReviewService/RemoveVoteFromComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reviewServiceClient) RemoveCommentFromSnippetTopic(ctx context.Context, in *RemoveCommentFromSnippetTopicRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/proto.ReviewService/RemoveCommentFromSnippetTopic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reviewServiceClient) DeleteSnippetTopic(ctx context.Context, in *DeleteSnippetTopicRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/proto.ReviewService/DeleteSnippetTopic", in, out, opts...)
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
	ChooseSnippetTopicSolution(context.Context, *ChooseSnippetTopicSolutionRequest) (*emptypb.Empty, error)
	ModifySnippetTopic(context.Context, *ModifySnippetTopicRequest) (*emptypb.Empty, error)
	ModifySnippet(context.Context, *ModifySnippetRequest) (*emptypb.Empty, error)
	ChooseSnippetTopicBestAnswer(context.Context, *ChooseSnippetTopicBestAnswerRequest) (*emptypb.Empty, error)
	VoteOnComment(context.Context, *VoteOnCommentRequest) (*emptypb.Empty, error)
	RemoveVoteFromComment(context.Context, *RemoveVoteFromCommentRequest) (*emptypb.Empty, error)
	RemoveCommentFromSnippetTopic(context.Context, *RemoveCommentFromSnippetTopicRequest) (*emptypb.Empty, error)
	DeleteSnippetTopic(context.Context, *DeleteSnippetTopicRequest) (*emptypb.Empty, error)
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
func (UnimplementedReviewServiceServer) ChooseSnippetTopicSolution(context.Context, *ChooseSnippetTopicSolutionRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChooseSnippetTopicSolution not implemented")
}
func (UnimplementedReviewServiceServer) ModifySnippetTopic(context.Context, *ModifySnippetTopicRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ModifySnippetTopic not implemented")
}
func (UnimplementedReviewServiceServer) ModifySnippet(context.Context, *ModifySnippetRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ModifySnippet not implemented")
}
func (UnimplementedReviewServiceServer) ChooseSnippetTopicBestAnswer(context.Context, *ChooseSnippetTopicBestAnswerRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChooseSnippetTopicBestAnswer not implemented")
}
func (UnimplementedReviewServiceServer) VoteOnComment(context.Context, *VoteOnCommentRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VoteOnComment not implemented")
}
func (UnimplementedReviewServiceServer) RemoveVoteFromComment(context.Context, *RemoveVoteFromCommentRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveVoteFromComment not implemented")
}
func (UnimplementedReviewServiceServer) RemoveCommentFromSnippetTopic(context.Context, *RemoveCommentFromSnippetTopicRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveCommentFromSnippetTopic not implemented")
}
func (UnimplementedReviewServiceServer) DeleteSnippetTopic(context.Context, *DeleteSnippetTopicRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSnippetTopic not implemented")
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

func _ReviewService_ChooseSnippetTopicSolution_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChooseSnippetTopicSolutionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewServiceServer).ChooseSnippetTopicSolution(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ReviewService/ChooseSnippetTopicSolution",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewServiceServer).ChooseSnippetTopicSolution(ctx, req.(*ChooseSnippetTopicSolutionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReviewService_ModifySnippetTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModifySnippetTopicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewServiceServer).ModifySnippetTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ReviewService/ModifySnippetTopic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewServiceServer).ModifySnippetTopic(ctx, req.(*ModifySnippetTopicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReviewService_ModifySnippet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModifySnippetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewServiceServer).ModifySnippet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ReviewService/ModifySnippet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewServiceServer).ModifySnippet(ctx, req.(*ModifySnippetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReviewService_ChooseSnippetTopicBestAnswer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChooseSnippetTopicBestAnswerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewServiceServer).ChooseSnippetTopicBestAnswer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ReviewService/ChooseSnippetTopicBestAnswer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewServiceServer).ChooseSnippetTopicBestAnswer(ctx, req.(*ChooseSnippetTopicBestAnswerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReviewService_VoteOnComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VoteOnCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewServiceServer).VoteOnComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ReviewService/VoteOnComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewServiceServer).VoteOnComment(ctx, req.(*VoteOnCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReviewService_RemoveVoteFromComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveVoteFromCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewServiceServer).RemoveVoteFromComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ReviewService/RemoveVoteFromComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewServiceServer).RemoveVoteFromComment(ctx, req.(*RemoveVoteFromCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReviewService_RemoveCommentFromSnippetTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveCommentFromSnippetTopicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewServiceServer).RemoveCommentFromSnippetTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ReviewService/RemoveCommentFromSnippetTopic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewServiceServer).RemoveCommentFromSnippetTopic(ctx, req.(*RemoveCommentFromSnippetTopicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReviewService_DeleteSnippetTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSnippetTopicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewServiceServer).DeleteSnippetTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ReviewService/DeleteSnippetTopic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewServiceServer).DeleteSnippetTopic(ctx, req.(*DeleteSnippetTopicRequest))
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
			MethodName: "ChooseSnippetTopicSolution",
			Handler:    _ReviewService_ChooseSnippetTopicSolution_Handler,
		},
		{
			MethodName: "ModifySnippetTopic",
			Handler:    _ReviewService_ModifySnippetTopic_Handler,
		},
		{
			MethodName: "ModifySnippet",
			Handler:    _ReviewService_ModifySnippet_Handler,
		},
		{
			MethodName: "ChooseSnippetTopicBestAnswer",
			Handler:    _ReviewService_ChooseSnippetTopicBestAnswer_Handler,
		},
		{
			MethodName: "VoteOnComment",
			Handler:    _ReviewService_VoteOnComment_Handler,
		},
		{
			MethodName: "RemoveVoteFromComment",
			Handler:    _ReviewService_RemoveVoteFromComment_Handler,
		},
		{
			MethodName: "RemoveCommentFromSnippetTopic",
			Handler:    _ReviewService_RemoveCommentFromSnippetTopic_Handler,
		},
		{
			MethodName: "DeleteSnippetTopic",
			Handler:    _ReviewService_DeleteSnippetTopic_Handler,
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
