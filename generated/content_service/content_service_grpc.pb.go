// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: content_service.proto

package content_service

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ContentClient is the client API for Content service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ContentClient interface {
	CreateStories(ctx context.Context, in *CreateStoriesRequest, opts ...grpc.CallOption) (*CreateStoriesResponce, error)
	UpdateStories(ctx context.Context, in *UpdateStoriesRequest, opts ...grpc.CallOption) (*UpdateStoriesResponce, error)
	DeleteStories(ctx context.Context, in *DeleteStoriesRequest, opts ...grpc.CallOption) (*DeleteStoriesResponce, error)
	GetsStories(ctx context.Context, in *GetsStoriesRequest, opts ...grpc.CallOption) (*GetsStoriesResponce, error)
	GetStories(ctx context.Context, in *GetStoriesRequest, opts ...grpc.CallOption) (*GetStoriesResponce, error)
	CommentStories(ctx context.Context, in *CommentStoriesRequest, opts ...grpc.CallOption) (*CommentStoriesResponce, error)
	GetCommentStories(ctx context.Context, in *GetCommentStoriesRequest, opts ...grpc.CallOption) (*GetCommentStoriesResponce, error)
	LikeStories(ctx context.Context, in *LikeStoriesRequest, opts ...grpc.CallOption) (*LikeStoriesResponce, error)
	CreateItineraries(ctx context.Context, in *CreateItinerariesRequest, opts ...grpc.CallOption) (*CreateItinerariesResponce, error)
	UpdateItineraries(ctx context.Context, in *UpdateItinerariesRequest, opts ...grpc.CallOption) (*UpdateItinerariesResponce, error)
	DeleteItineraries(ctx context.Context, in *DeleteItinerariesRequest, opts ...grpc.CallOption) (*DeleteItinerariesResponce, error)
	GetsItineraries(ctx context.Context, in *GetsItinerariesRequest, opts ...grpc.CallOption) (*GetsItinerariesResponce, error)
	GetItineraries(ctx context.Context, in *GetItinerariesRequest, opts ...grpc.CallOption) (*GetItinerariesResponce, error)
	ItinerariesComment(ctx context.Context, in *ItinerariesCommentRequest, opts ...grpc.CallOption) (*ItinerariesCommentResponce, error)
	GetsDestinations(ctx context.Context, in *GetsDestinationsRequest, opts ...grpc.CallOption) (*GetsDestinationsResponce, error)
	GetDestinations(ctx context.Context, in *GetDestinationsRequest, opts ...grpc.CallOption) (*GetDestinationsResponce, error)
	Message(ctx context.Context, in *MessageRequest, opts ...grpc.CallOption) (*MessageResponce, error)
	GetMessage(ctx context.Context, in *GetMessageRequest, opts ...grpc.CallOption) (*GetMessageResponce, error)
	AddTravelTips(ctx context.Context, in *AddTravelTipsRequest, opts ...grpc.CallOption) (*AddTravelTipsResponce, error)
	GetTravelTips(ctx context.Context, in *GetTravelTipsRequest, opts ...grpc.CallOption) (*GetTravelTipsResponce, error)
	StatisticsUser(ctx context.Context, in *StatisticsUserRequest, opts ...grpc.CallOption) (*StatisticsUserResponce, error)
	GetTrendingDestinations(ctx context.Context, in *GetTrendingDestinationsRequest, opts ...grpc.CallOption) (*GetTrendingDestinationsResponce, error)
}

type contentClient struct {
	cc grpc.ClientConnInterface
}

func NewContentClient(cc grpc.ClientConnInterface) ContentClient {
	return &contentClient{cc}
}

func (c *contentClient) CreateStories(ctx context.Context, in *CreateStoriesRequest, opts ...grpc.CallOption) (*CreateStoriesResponce, error) {
	out := new(CreateStoriesResponce)
	err := c.cc.Invoke(ctx, "/content_service.Content/CreateStories", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentClient) UpdateStories(ctx context.Context, in *UpdateStoriesRequest, opts ...grpc.CallOption) (*UpdateStoriesResponce, error) {
	out := new(UpdateStoriesResponce)
	err := c.cc.Invoke(ctx, "/content_service.Content/UpdateStories", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentClient) DeleteStories(ctx context.Context, in *DeleteStoriesRequest, opts ...grpc.CallOption) (*DeleteStoriesResponce, error) {
	out := new(DeleteStoriesResponce)
	err := c.cc.Invoke(ctx, "/content_service.Content/DeleteStories", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentClient) GetsStories(ctx context.Context, in *GetsStoriesRequest, opts ...grpc.CallOption) (*GetsStoriesResponce, error) {
	out := new(GetsStoriesResponce)
	err := c.cc.Invoke(ctx, "/content_service.Content/GetsStories", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentClient) GetStories(ctx context.Context, in *GetStoriesRequest, opts ...grpc.CallOption) (*GetStoriesResponce, error) {
	out := new(GetStoriesResponce)
	err := c.cc.Invoke(ctx, "/content_service.Content/GetStories", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentClient) CommentStories(ctx context.Context, in *CommentStoriesRequest, opts ...grpc.CallOption) (*CommentStoriesResponce, error) {
	out := new(CommentStoriesResponce)
	err := c.cc.Invoke(ctx, "/content_service.Content/CommentStories", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentClient) GetCommentStories(ctx context.Context, in *GetCommentStoriesRequest, opts ...grpc.CallOption) (*GetCommentStoriesResponce, error) {
	out := new(GetCommentStoriesResponce)
	err := c.cc.Invoke(ctx, "/content_service.Content/GetCommentStories", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentClient) LikeStories(ctx context.Context, in *LikeStoriesRequest, opts ...grpc.CallOption) (*LikeStoriesResponce, error) {
	out := new(LikeStoriesResponce)
	err := c.cc.Invoke(ctx, "/content_service.Content/LikeStories", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentClient) CreateItineraries(ctx context.Context, in *CreateItinerariesRequest, opts ...grpc.CallOption) (*CreateItinerariesResponce, error) {
	out := new(CreateItinerariesResponce)
	err := c.cc.Invoke(ctx, "/content_service.Content/CreateItineraries", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentClient) UpdateItineraries(ctx context.Context, in *UpdateItinerariesRequest, opts ...grpc.CallOption) (*UpdateItinerariesResponce, error) {
	out := new(UpdateItinerariesResponce)
	err := c.cc.Invoke(ctx, "/content_service.Content/UpdateItineraries", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentClient) DeleteItineraries(ctx context.Context, in *DeleteItinerariesRequest, opts ...grpc.CallOption) (*DeleteItinerariesResponce, error) {
	out := new(DeleteItinerariesResponce)
	err := c.cc.Invoke(ctx, "/content_service.Content/DeleteItineraries", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentClient) GetsItineraries(ctx context.Context, in *GetsItinerariesRequest, opts ...grpc.CallOption) (*GetsItinerariesResponce, error) {
	out := new(GetsItinerariesResponce)
	err := c.cc.Invoke(ctx, "/content_service.Content/GetsItineraries", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentClient) GetItineraries(ctx context.Context, in *GetItinerariesRequest, opts ...grpc.CallOption) (*GetItinerariesResponce, error) {
	out := new(GetItinerariesResponce)
	err := c.cc.Invoke(ctx, "/content_service.Content/GetItineraries", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentClient) ItinerariesComment(ctx context.Context, in *ItinerariesCommentRequest, opts ...grpc.CallOption) (*ItinerariesCommentResponce, error) {
	out := new(ItinerariesCommentResponce)
	err := c.cc.Invoke(ctx, "/content_service.Content/ItinerariesComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentClient) GetsDestinations(ctx context.Context, in *GetsDestinationsRequest, opts ...grpc.CallOption) (*GetsDestinationsResponce, error) {
	out := new(GetsDestinationsResponce)
	err := c.cc.Invoke(ctx, "/content_service.Content/GetsDestinations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentClient) GetDestinations(ctx context.Context, in *GetDestinationsRequest, opts ...grpc.CallOption) (*GetDestinationsResponce, error) {
	out := new(GetDestinationsResponce)
	err := c.cc.Invoke(ctx, "/content_service.Content/GetDestinations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentClient) Message(ctx context.Context, in *MessageRequest, opts ...grpc.CallOption) (*MessageResponce, error) {
	out := new(MessageResponce)
	err := c.cc.Invoke(ctx, "/content_service.Content/Message", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentClient) GetMessage(ctx context.Context, in *GetMessageRequest, opts ...grpc.CallOption) (*GetMessageResponce, error) {
	out := new(GetMessageResponce)
	err := c.cc.Invoke(ctx, "/content_service.Content/GetMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentClient) AddTravelTips(ctx context.Context, in *AddTravelTipsRequest, opts ...grpc.CallOption) (*AddTravelTipsResponce, error) {
	out := new(AddTravelTipsResponce)
	err := c.cc.Invoke(ctx, "/content_service.Content/AddTravelTips", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentClient) GetTravelTips(ctx context.Context, in *GetTravelTipsRequest, opts ...grpc.CallOption) (*GetTravelTipsResponce, error) {
	out := new(GetTravelTipsResponce)
	err := c.cc.Invoke(ctx, "/content_service.Content/GetTravelTips", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentClient) StatisticsUser(ctx context.Context, in *StatisticsUserRequest, opts ...grpc.CallOption) (*StatisticsUserResponce, error) {
	out := new(StatisticsUserResponce)
	err := c.cc.Invoke(ctx, "/content_service.Content/StatisticsUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentClient) GetTrendingDestinations(ctx context.Context, in *GetTrendingDestinationsRequest, opts ...grpc.CallOption) (*GetTrendingDestinationsResponce, error) {
	out := new(GetTrendingDestinationsResponce)
	err := c.cc.Invoke(ctx, "/content_service.Content/GetTrendingDestinations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ContentServer is the server API for Content service.
// All implementations must embed UnimplementedContentServer
// for forward compatibility
type ContentServer interface {
	CreateStories(context.Context, *CreateStoriesRequest) (*CreateStoriesResponce, error)
	UpdateStories(context.Context, *UpdateStoriesRequest) (*UpdateStoriesResponce, error)
	DeleteStories(context.Context, *DeleteStoriesRequest) (*DeleteStoriesResponce, error)
	GetsStories(context.Context, *GetsStoriesRequest) (*GetsStoriesResponce, error)
	GetStories(context.Context, *GetStoriesRequest) (*GetStoriesResponce, error)
	CommentStories(context.Context, *CommentStoriesRequest) (*CommentStoriesResponce, error)
	GetCommentStories(context.Context, *GetCommentStoriesRequest) (*GetCommentStoriesResponce, error)
	LikeStories(context.Context, *LikeStoriesRequest) (*LikeStoriesResponce, error)
	CreateItineraries(context.Context, *CreateItinerariesRequest) (*CreateItinerariesResponce, error)
	UpdateItineraries(context.Context, *UpdateItinerariesRequest) (*UpdateItinerariesResponce, error)
	DeleteItineraries(context.Context, *DeleteItinerariesRequest) (*DeleteItinerariesResponce, error)
	GetsItineraries(context.Context, *GetsItinerariesRequest) (*GetsItinerariesResponce, error)
	GetItineraries(context.Context, *GetItinerariesRequest) (*GetItinerariesResponce, error)
	ItinerariesComment(context.Context, *ItinerariesCommentRequest) (*ItinerariesCommentResponce, error)
	GetsDestinations(context.Context, *GetsDestinationsRequest) (*GetsDestinationsResponce, error)
	GetDestinations(context.Context, *GetDestinationsRequest) (*GetDestinationsResponce, error)
	Message(context.Context, *MessageRequest) (*MessageResponce, error)
	GetMessage(context.Context, *GetMessageRequest) (*GetMessageResponce, error)
	AddTravelTips(context.Context, *AddTravelTipsRequest) (*AddTravelTipsResponce, error)
	GetTravelTips(context.Context, *GetTravelTipsRequest) (*GetTravelTipsResponce, error)
	StatisticsUser(context.Context, *StatisticsUserRequest) (*StatisticsUserResponce, error)
	GetTrendingDestinations(context.Context, *GetTrendingDestinationsRequest) (*GetTrendingDestinationsResponce, error)
	mustEmbedUnimplementedContentServer()
}

// UnimplementedContentServer must be embedded to have forward compatible implementations.
type UnimplementedContentServer struct {
}

func (UnimplementedContentServer) CreateStories(context.Context, *CreateStoriesRequest) (*CreateStoriesResponce, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateStories not implemented")
}
func (UnimplementedContentServer) UpdateStories(context.Context, *UpdateStoriesRequest) (*UpdateStoriesResponce, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateStories not implemented")
}
func (UnimplementedContentServer) DeleteStories(context.Context, *DeleteStoriesRequest) (*DeleteStoriesResponce, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteStories not implemented")
}
func (UnimplementedContentServer) GetsStories(context.Context, *GetsStoriesRequest) (*GetsStoriesResponce, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetsStories not implemented")
}
func (UnimplementedContentServer) GetStories(context.Context, *GetStoriesRequest) (*GetStoriesResponce, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStories not implemented")
}
func (UnimplementedContentServer) CommentStories(context.Context, *CommentStoriesRequest) (*CommentStoriesResponce, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CommentStories not implemented")
}
func (UnimplementedContentServer) GetCommentStories(context.Context, *GetCommentStoriesRequest) (*GetCommentStoriesResponce, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCommentStories not implemented")
}
func (UnimplementedContentServer) LikeStories(context.Context, *LikeStoriesRequest) (*LikeStoriesResponce, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LikeStories not implemented")
}
func (UnimplementedContentServer) CreateItineraries(context.Context, *CreateItinerariesRequest) (*CreateItinerariesResponce, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateItineraries not implemented")
}
func (UnimplementedContentServer) UpdateItineraries(context.Context, *UpdateItinerariesRequest) (*UpdateItinerariesResponce, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateItineraries not implemented")
}
func (UnimplementedContentServer) DeleteItineraries(context.Context, *DeleteItinerariesRequest) (*DeleteItinerariesResponce, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteItineraries not implemented")
}
func (UnimplementedContentServer) GetsItineraries(context.Context, *GetsItinerariesRequest) (*GetsItinerariesResponce, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetsItineraries not implemented")
}
func (UnimplementedContentServer) GetItineraries(context.Context, *GetItinerariesRequest) (*GetItinerariesResponce, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetItineraries not implemented")
}
func (UnimplementedContentServer) ItinerariesComment(context.Context, *ItinerariesCommentRequest) (*ItinerariesCommentResponce, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ItinerariesComment not implemented")
}
func (UnimplementedContentServer) GetsDestinations(context.Context, *GetsDestinationsRequest) (*GetsDestinationsResponce, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetsDestinations not implemented")
}
func (UnimplementedContentServer) GetDestinations(context.Context, *GetDestinationsRequest) (*GetDestinationsResponce, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDestinations not implemented")
}
func (UnimplementedContentServer) Message(context.Context, *MessageRequest) (*MessageResponce, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Message not implemented")
}
func (UnimplementedContentServer) GetMessage(context.Context, *GetMessageRequest) (*GetMessageResponce, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMessage not implemented")
}
func (UnimplementedContentServer) AddTravelTips(context.Context, *AddTravelTipsRequest) (*AddTravelTipsResponce, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTravelTips not implemented")
}
func (UnimplementedContentServer) GetTravelTips(context.Context, *GetTravelTipsRequest) (*GetTravelTipsResponce, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTravelTips not implemented")
}
func (UnimplementedContentServer) StatisticsUser(context.Context, *StatisticsUserRequest) (*StatisticsUserResponce, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StatisticsUser not implemented")
}
func (UnimplementedContentServer) GetTrendingDestinations(context.Context, *GetTrendingDestinationsRequest) (*GetTrendingDestinationsResponce, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTrendingDestinations not implemented")
}
func (UnimplementedContentServer) mustEmbedUnimplementedContentServer() {}

// UnsafeContentServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ContentServer will
// result in compilation errors.
type UnsafeContentServer interface {
	mustEmbedUnimplementedContentServer()
}

func RegisterContentServer(s grpc.ServiceRegistrar, srv ContentServer) {
	s.RegisterService(&Content_ServiceDesc, srv)
}

func _Content_CreateStories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateStoriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServer).CreateStories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content_service.Content/CreateStories",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServer).CreateStories(ctx, req.(*CreateStoriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Content_UpdateStories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateStoriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServer).UpdateStories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content_service.Content/UpdateStories",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServer).UpdateStories(ctx, req.(*UpdateStoriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Content_DeleteStories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteStoriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServer).DeleteStories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content_service.Content/DeleteStories",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServer).DeleteStories(ctx, req.(*DeleteStoriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Content_GetsStories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetsStoriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServer).GetsStories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content_service.Content/GetsStories",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServer).GetsStories(ctx, req.(*GetsStoriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Content_GetStories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStoriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServer).GetStories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content_service.Content/GetStories",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServer).GetStories(ctx, req.(*GetStoriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Content_CommentStories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommentStoriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServer).CommentStories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content_service.Content/CommentStories",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServer).CommentStories(ctx, req.(*CommentStoriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Content_GetCommentStories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCommentStoriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServer).GetCommentStories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content_service.Content/GetCommentStories",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServer).GetCommentStories(ctx, req.(*GetCommentStoriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Content_LikeStories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LikeStoriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServer).LikeStories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content_service.Content/LikeStories",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServer).LikeStories(ctx, req.(*LikeStoriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Content_CreateItineraries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateItinerariesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServer).CreateItineraries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content_service.Content/CreateItineraries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServer).CreateItineraries(ctx, req.(*CreateItinerariesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Content_UpdateItineraries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateItinerariesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServer).UpdateItineraries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content_service.Content/UpdateItineraries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServer).UpdateItineraries(ctx, req.(*UpdateItinerariesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Content_DeleteItineraries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteItinerariesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServer).DeleteItineraries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content_service.Content/DeleteItineraries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServer).DeleteItineraries(ctx, req.(*DeleteItinerariesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Content_GetsItineraries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetsItinerariesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServer).GetsItineraries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content_service.Content/GetsItineraries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServer).GetsItineraries(ctx, req.(*GetsItinerariesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Content_GetItineraries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetItinerariesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServer).GetItineraries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content_service.Content/GetItineraries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServer).GetItineraries(ctx, req.(*GetItinerariesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Content_ItinerariesComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ItinerariesCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServer).ItinerariesComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content_service.Content/ItinerariesComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServer).ItinerariesComment(ctx, req.(*ItinerariesCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Content_GetsDestinations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetsDestinationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServer).GetsDestinations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content_service.Content/GetsDestinations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServer).GetsDestinations(ctx, req.(*GetsDestinationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Content_GetDestinations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDestinationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServer).GetDestinations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content_service.Content/GetDestinations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServer).GetDestinations(ctx, req.(*GetDestinationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Content_Message_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServer).Message(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content_service.Content/Message",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServer).Message(ctx, req.(*MessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Content_GetMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServer).GetMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content_service.Content/GetMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServer).GetMessage(ctx, req.(*GetMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Content_AddTravelTips_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddTravelTipsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServer).AddTravelTips(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content_service.Content/AddTravelTips",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServer).AddTravelTips(ctx, req.(*AddTravelTipsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Content_GetTravelTips_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTravelTipsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServer).GetTravelTips(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content_service.Content/GetTravelTips",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServer).GetTravelTips(ctx, req.(*GetTravelTipsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Content_StatisticsUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatisticsUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServer).StatisticsUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content_service.Content/StatisticsUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServer).StatisticsUser(ctx, req.(*StatisticsUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Content_GetTrendingDestinations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTrendingDestinationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServer).GetTrendingDestinations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content_service.Content/GetTrendingDestinations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServer).GetTrendingDestinations(ctx, req.(*GetTrendingDestinationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Content_ServiceDesc is the grpc.ServiceDesc for Content service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Content_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "content_service.Content",
	HandlerType: (*ContentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateStories",
			Handler:    _Content_CreateStories_Handler,
		},
		{
			MethodName: "UpdateStories",
			Handler:    _Content_UpdateStories_Handler,
		},
		{
			MethodName: "DeleteStories",
			Handler:    _Content_DeleteStories_Handler,
		},
		{
			MethodName: "GetsStories",
			Handler:    _Content_GetsStories_Handler,
		},
		{
			MethodName: "GetStories",
			Handler:    _Content_GetStories_Handler,
		},
		{
			MethodName: "CommentStories",
			Handler:    _Content_CommentStories_Handler,
		},
		{
			MethodName: "GetCommentStories",
			Handler:    _Content_GetCommentStories_Handler,
		},
		{
			MethodName: "LikeStories",
			Handler:    _Content_LikeStories_Handler,
		},
		{
			MethodName: "CreateItineraries",
			Handler:    _Content_CreateItineraries_Handler,
		},
		{
			MethodName: "UpdateItineraries",
			Handler:    _Content_UpdateItineraries_Handler,
		},
		{
			MethodName: "DeleteItineraries",
			Handler:    _Content_DeleteItineraries_Handler,
		},
		{
			MethodName: "GetsItineraries",
			Handler:    _Content_GetsItineraries_Handler,
		},
		{
			MethodName: "GetItineraries",
			Handler:    _Content_GetItineraries_Handler,
		},
		{
			MethodName: "ItinerariesComment",
			Handler:    _Content_ItinerariesComment_Handler,
		},
		{
			MethodName: "GetsDestinations",
			Handler:    _Content_GetsDestinations_Handler,
		},
		{
			MethodName: "GetDestinations",
			Handler:    _Content_GetDestinations_Handler,
		},
		{
			MethodName: "Message",
			Handler:    _Content_Message_Handler,
		},
		{
			MethodName: "GetMessage",
			Handler:    _Content_GetMessage_Handler,
		},
		{
			MethodName: "AddTravelTips",
			Handler:    _Content_AddTravelTips_Handler,
		},
		{
			MethodName: "GetTravelTips",
			Handler:    _Content_GetTravelTips_Handler,
		},
		{
			MethodName: "StatisticsUser",
			Handler:    _Content_StatisticsUser_Handler,
		},
		{
			MethodName: "GetTrendingDestinations",
			Handler:    _Content_GetTrendingDestinations_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "content_service.proto",
}