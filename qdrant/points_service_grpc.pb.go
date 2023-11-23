// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.22.2
// source: points_service.proto

package go_client

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

const (
	Points_Upsert_FullMethodName           = "/qdrant.Points/Upsert"
	Points_Delete_FullMethodName           = "/qdrant.Points/Delete"
	Points_Get_FullMethodName              = "/qdrant.Points/Get"
	Points_UpdateVectors_FullMethodName    = "/qdrant.Points/UpdateVectors"
	Points_DeleteVectors_FullMethodName    = "/qdrant.Points/DeleteVectors"
	Points_SetPayload_FullMethodName       = "/qdrant.Points/SetPayload"
	Points_OverwritePayload_FullMethodName = "/qdrant.Points/OverwritePayload"
	Points_DeletePayload_FullMethodName    = "/qdrant.Points/DeletePayload"
	Points_ClearPayload_FullMethodName     = "/qdrant.Points/ClearPayload"
	Points_CreateFieldIndex_FullMethodName = "/qdrant.Points/CreateFieldIndex"
	Points_DeleteFieldIndex_FullMethodName = "/qdrant.Points/DeleteFieldIndex"
	Points_Search_FullMethodName           = "/qdrant.Points/Search"
	Points_SearchBatch_FullMethodName      = "/qdrant.Points/SearchBatch"
	Points_SearchGroups_FullMethodName     = "/qdrant.Points/SearchGroups"
	Points_Scroll_FullMethodName           = "/qdrant.Points/Scroll"
	Points_Recommend_FullMethodName        = "/qdrant.Points/Recommend"
	Points_RecommendBatch_FullMethodName   = "/qdrant.Points/RecommendBatch"
	Points_RecommendGroups_FullMethodName  = "/qdrant.Points/RecommendGroups"
	Points_Count_FullMethodName            = "/qdrant.Points/Count"
	Points_UpdateBatch_FullMethodName      = "/qdrant.Points/UpdateBatch"
)

// PointsClient is the client API for Points service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PointsClient interface {
	//
	//Perform insert + updates on points. If a point with a given ID already exists - it will be overwritten.
	Upsert(ctx context.Context, in *UpsertPoints, opts ...grpc.CallOption) (*PointsOperationResponse, error)
	//
	//Delete points
	Delete(ctx context.Context, in *DeletePoints, opts ...grpc.CallOption) (*PointsOperationResponse, error)
	//
	//Retrieve points
	Get(ctx context.Context, in *GetPoints, opts ...grpc.CallOption) (*GetResponse, error)
	//
	//Update named vectors for point
	UpdateVectors(ctx context.Context, in *UpdatePointVectors, opts ...grpc.CallOption) (*PointsOperationResponse, error)
	//
	//Delete named vectors for points
	DeleteVectors(ctx context.Context, in *DeletePointVectors, opts ...grpc.CallOption) (*PointsOperationResponse, error)
	//
	//Set payload for points
	SetPayload(ctx context.Context, in *SetPayloadPoints, opts ...grpc.CallOption) (*PointsOperationResponse, error)
	//
	//Overwrite payload for points
	OverwritePayload(ctx context.Context, in *SetPayloadPoints, opts ...grpc.CallOption) (*PointsOperationResponse, error)
	//
	//Delete specified key payload for points
	DeletePayload(ctx context.Context, in *DeletePayloadPoints, opts ...grpc.CallOption) (*PointsOperationResponse, error)
	//
	//Remove all payload for specified points
	ClearPayload(ctx context.Context, in *ClearPayloadPoints, opts ...grpc.CallOption) (*PointsOperationResponse, error)
	//
	//Create index for field in collection
	CreateFieldIndex(ctx context.Context, in *CreateFieldIndexCollection, opts ...grpc.CallOption) (*PointsOperationResponse, error)
	//
	//Delete field index for collection
	DeleteFieldIndex(ctx context.Context, in *DeleteFieldIndexCollection, opts ...grpc.CallOption) (*PointsOperationResponse, error)
	//
	//Retrieve closest points based on vector similarity and given filtering conditions
	Search(ctx context.Context, in *SearchPoints, opts ...grpc.CallOption) (*SearchResponse, error)
	//
	//Retrieve closest points based on vector similarity and given filtering conditions
	SearchBatch(ctx context.Context, in *SearchBatchPoints, opts ...grpc.CallOption) (*SearchBatchResponse, error)
	//
	//Retrieve closest points based on vector similarity and given filtering conditions, grouped by a given field
	SearchGroups(ctx context.Context, in *SearchPointGroups, opts ...grpc.CallOption) (*SearchGroupsResponse, error)
	//
	//Iterate over all or filtered points
	Scroll(ctx context.Context, in *ScrollPoints, opts ...grpc.CallOption) (*ScrollResponse, error)
	//
	//Look for the points which are closer to stored positive examples and at the same time further to negative examples.
	Recommend(ctx context.Context, in *RecommendPoints, opts ...grpc.CallOption) (*RecommendResponse, error)
	//
	//Look for the points which are closer to stored positive examples and at the same time further to negative examples.
	RecommendBatch(ctx context.Context, in *RecommendBatchPoints, opts ...grpc.CallOption) (*RecommendBatchResponse, error)
	//
	//Look for the points which are closer to stored positive examples and at the same time further to negative examples, grouped by a given field
	RecommendGroups(ctx context.Context, in *RecommendPointGroups, opts ...grpc.CallOption) (*RecommendGroupsResponse, error)
	//
	//Count points in collection with given filtering conditions
	Count(ctx context.Context, in *CountPoints, opts ...grpc.CallOption) (*CountResponse, error)
	//
	//Perform multiple update operations in one request
	UpdateBatch(ctx context.Context, in *UpdateBatchPoints, opts ...grpc.CallOption) (*UpdateBatchResponse, error)
}

type pointsClient struct {
	cc grpc.ClientConnInterface
}

func NewPointsClient(cc grpc.ClientConnInterface) PointsClient {
	return &pointsClient{cc}
}

func (c *pointsClient) Upsert(ctx context.Context, in *UpsertPoints, opts ...grpc.CallOption) (*PointsOperationResponse, error) {
	out := new(PointsOperationResponse)
	err := c.cc.Invoke(ctx, Points_Upsert_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pointsClient) Delete(ctx context.Context, in *DeletePoints, opts ...grpc.CallOption) (*PointsOperationResponse, error) {
	out := new(PointsOperationResponse)
	err := c.cc.Invoke(ctx, Points_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pointsClient) Get(ctx context.Context, in *GetPoints, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, Points_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pointsClient) UpdateVectors(ctx context.Context, in *UpdatePointVectors, opts ...grpc.CallOption) (*PointsOperationResponse, error) {
	out := new(PointsOperationResponse)
	err := c.cc.Invoke(ctx, Points_UpdateVectors_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pointsClient) DeleteVectors(ctx context.Context, in *DeletePointVectors, opts ...grpc.CallOption) (*PointsOperationResponse, error) {
	out := new(PointsOperationResponse)
	err := c.cc.Invoke(ctx, Points_DeleteVectors_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pointsClient) SetPayload(ctx context.Context, in *SetPayloadPoints, opts ...grpc.CallOption) (*PointsOperationResponse, error) {
	out := new(PointsOperationResponse)
	err := c.cc.Invoke(ctx, Points_SetPayload_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pointsClient) OverwritePayload(ctx context.Context, in *SetPayloadPoints, opts ...grpc.CallOption) (*PointsOperationResponse, error) {
	out := new(PointsOperationResponse)
	err := c.cc.Invoke(ctx, Points_OverwritePayload_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pointsClient) DeletePayload(ctx context.Context, in *DeletePayloadPoints, opts ...grpc.CallOption) (*PointsOperationResponse, error) {
	out := new(PointsOperationResponse)
	err := c.cc.Invoke(ctx, Points_DeletePayload_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pointsClient) ClearPayload(ctx context.Context, in *ClearPayloadPoints, opts ...grpc.CallOption) (*PointsOperationResponse, error) {
	out := new(PointsOperationResponse)
	err := c.cc.Invoke(ctx, Points_ClearPayload_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pointsClient) CreateFieldIndex(ctx context.Context, in *CreateFieldIndexCollection, opts ...grpc.CallOption) (*PointsOperationResponse, error) {
	out := new(PointsOperationResponse)
	err := c.cc.Invoke(ctx, Points_CreateFieldIndex_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pointsClient) DeleteFieldIndex(ctx context.Context, in *DeleteFieldIndexCollection, opts ...grpc.CallOption) (*PointsOperationResponse, error) {
	out := new(PointsOperationResponse)
	err := c.cc.Invoke(ctx, Points_DeleteFieldIndex_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pointsClient) Search(ctx context.Context, in *SearchPoints, opts ...grpc.CallOption) (*SearchResponse, error) {
	out := new(SearchResponse)
	err := c.cc.Invoke(ctx, Points_Search_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pointsClient) SearchBatch(ctx context.Context, in *SearchBatchPoints, opts ...grpc.CallOption) (*SearchBatchResponse, error) {
	out := new(SearchBatchResponse)
	err := c.cc.Invoke(ctx, Points_SearchBatch_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pointsClient) SearchGroups(ctx context.Context, in *SearchPointGroups, opts ...grpc.CallOption) (*SearchGroupsResponse, error) {
	out := new(SearchGroupsResponse)
	err := c.cc.Invoke(ctx, Points_SearchGroups_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pointsClient) Scroll(ctx context.Context, in *ScrollPoints, opts ...grpc.CallOption) (*ScrollResponse, error) {
	out := new(ScrollResponse)
	err := c.cc.Invoke(ctx, Points_Scroll_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pointsClient) Recommend(ctx context.Context, in *RecommendPoints, opts ...grpc.CallOption) (*RecommendResponse, error) {
	out := new(RecommendResponse)
	err := c.cc.Invoke(ctx, Points_Recommend_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pointsClient) RecommendBatch(ctx context.Context, in *RecommendBatchPoints, opts ...grpc.CallOption) (*RecommendBatchResponse, error) {
	out := new(RecommendBatchResponse)
	err := c.cc.Invoke(ctx, Points_RecommendBatch_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pointsClient) RecommendGroups(ctx context.Context, in *RecommendPointGroups, opts ...grpc.CallOption) (*RecommendGroupsResponse, error) {
	out := new(RecommendGroupsResponse)
	err := c.cc.Invoke(ctx, Points_RecommendGroups_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pointsClient) Count(ctx context.Context, in *CountPoints, opts ...grpc.CallOption) (*CountResponse, error) {
	out := new(CountResponse)
	err := c.cc.Invoke(ctx, Points_Count_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pointsClient) UpdateBatch(ctx context.Context, in *UpdateBatchPoints, opts ...grpc.CallOption) (*UpdateBatchResponse, error) {
	out := new(UpdateBatchResponse)
	err := c.cc.Invoke(ctx, Points_UpdateBatch_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PointsServer is the server API for Points service.
// All implementations must embed UnimplementedPointsServer
// for forward compatibility
type PointsServer interface {
	//
	//Perform insert + updates on points. If a point with a given ID already exists - it will be overwritten.
	Upsert(context.Context, *UpsertPoints) (*PointsOperationResponse, error)
	//
	//Delete points
	Delete(context.Context, *DeletePoints) (*PointsOperationResponse, error)
	//
	//Retrieve points
	Get(context.Context, *GetPoints) (*GetResponse, error)
	//
	//Update named vectors for point
	UpdateVectors(context.Context, *UpdatePointVectors) (*PointsOperationResponse, error)
	//
	//Delete named vectors for points
	DeleteVectors(context.Context, *DeletePointVectors) (*PointsOperationResponse, error)
	//
	//Set payload for points
	SetPayload(context.Context, *SetPayloadPoints) (*PointsOperationResponse, error)
	//
	//Overwrite payload for points
	OverwritePayload(context.Context, *SetPayloadPoints) (*PointsOperationResponse, error)
	//
	//Delete specified key payload for points
	DeletePayload(context.Context, *DeletePayloadPoints) (*PointsOperationResponse, error)
	//
	//Remove all payload for specified points
	ClearPayload(context.Context, *ClearPayloadPoints) (*PointsOperationResponse, error)
	//
	//Create index for field in collection
	CreateFieldIndex(context.Context, *CreateFieldIndexCollection) (*PointsOperationResponse, error)
	//
	//Delete field index for collection
	DeleteFieldIndex(context.Context, *DeleteFieldIndexCollection) (*PointsOperationResponse, error)
	//
	//Retrieve closest points based on vector similarity and given filtering conditions
	Search(context.Context, *SearchPoints) (*SearchResponse, error)
	//
	//Retrieve closest points based on vector similarity and given filtering conditions
	SearchBatch(context.Context, *SearchBatchPoints) (*SearchBatchResponse, error)
	//
	//Retrieve closest points based on vector similarity and given filtering conditions, grouped by a given field
	SearchGroups(context.Context, *SearchPointGroups) (*SearchGroupsResponse, error)
	//
	//Iterate over all or filtered points
	Scroll(context.Context, *ScrollPoints) (*ScrollResponse, error)
	//
	//Look for the points which are closer to stored positive examples and at the same time further to negative examples.
	Recommend(context.Context, *RecommendPoints) (*RecommendResponse, error)
	//
	//Look for the points which are closer to stored positive examples and at the same time further to negative examples.
	RecommendBatch(context.Context, *RecommendBatchPoints) (*RecommendBatchResponse, error)
	//
	//Look for the points which are closer to stored positive examples and at the same time further to negative examples, grouped by a given field
	RecommendGroups(context.Context, *RecommendPointGroups) (*RecommendGroupsResponse, error)
	//
	//Count points in collection with given filtering conditions
	Count(context.Context, *CountPoints) (*CountResponse, error)
	//
	//Perform multiple update operations in one request
	UpdateBatch(context.Context, *UpdateBatchPoints) (*UpdateBatchResponse, error)
	mustEmbedUnimplementedPointsServer()
}

// UnimplementedPointsServer must be embedded to have forward compatible implementations.
type UnimplementedPointsServer struct {
}

func (UnimplementedPointsServer) Upsert(context.Context, *UpsertPoints) (*PointsOperationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Upsert not implemented")
}
func (UnimplementedPointsServer) Delete(context.Context, *DeletePoints) (*PointsOperationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedPointsServer) Get(context.Context, *GetPoints) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedPointsServer) UpdateVectors(context.Context, *UpdatePointVectors) (*PointsOperationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateVectors not implemented")
}
func (UnimplementedPointsServer) DeleteVectors(context.Context, *DeletePointVectors) (*PointsOperationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteVectors not implemented")
}
func (UnimplementedPointsServer) SetPayload(context.Context, *SetPayloadPoints) (*PointsOperationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetPayload not implemented")
}
func (UnimplementedPointsServer) OverwritePayload(context.Context, *SetPayloadPoints) (*PointsOperationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OverwritePayload not implemented")
}
func (UnimplementedPointsServer) DeletePayload(context.Context, *DeletePayloadPoints) (*PointsOperationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePayload not implemented")
}
func (UnimplementedPointsServer) ClearPayload(context.Context, *ClearPayloadPoints) (*PointsOperationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClearPayload not implemented")
}
func (UnimplementedPointsServer) CreateFieldIndex(context.Context, *CreateFieldIndexCollection) (*PointsOperationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFieldIndex not implemented")
}
func (UnimplementedPointsServer) DeleteFieldIndex(context.Context, *DeleteFieldIndexCollection) (*PointsOperationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFieldIndex not implemented")
}
func (UnimplementedPointsServer) Search(context.Context, *SearchPoints) (*SearchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}
func (UnimplementedPointsServer) SearchBatch(context.Context, *SearchBatchPoints) (*SearchBatchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchBatch not implemented")
}
func (UnimplementedPointsServer) SearchGroups(context.Context, *SearchPointGroups) (*SearchGroupsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchGroups not implemented")
}
func (UnimplementedPointsServer) Scroll(context.Context, *ScrollPoints) (*ScrollResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Scroll not implemented")
}
func (UnimplementedPointsServer) Recommend(context.Context, *RecommendPoints) (*RecommendResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Recommend not implemented")
}
func (UnimplementedPointsServer) RecommendBatch(context.Context, *RecommendBatchPoints) (*RecommendBatchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RecommendBatch not implemented")
}
func (UnimplementedPointsServer) RecommendGroups(context.Context, *RecommendPointGroups) (*RecommendGroupsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RecommendGroups not implemented")
}
func (UnimplementedPointsServer) Count(context.Context, *CountPoints) (*CountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Count not implemented")
}
func (UnimplementedPointsServer) UpdateBatch(context.Context, *UpdateBatchPoints) (*UpdateBatchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBatch not implemented")
}
func (UnimplementedPointsServer) mustEmbedUnimplementedPointsServer() {}

// UnsafePointsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PointsServer will
// result in compilation errors.
type UnsafePointsServer interface {
	mustEmbedUnimplementedPointsServer()
}

func RegisterPointsServer(s grpc.ServiceRegistrar, srv PointsServer) {
	s.RegisterService(&Points_ServiceDesc, srv)
}

func _Points_Upsert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpsertPoints)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PointsServer).Upsert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Points_Upsert_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PointsServer).Upsert(ctx, req.(*UpsertPoints))
	}
	return interceptor(ctx, in, info, handler)
}

func _Points_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePoints)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PointsServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Points_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PointsServer).Delete(ctx, req.(*DeletePoints))
	}
	return interceptor(ctx, in, info, handler)
}

func _Points_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPoints)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PointsServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Points_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PointsServer).Get(ctx, req.(*GetPoints))
	}
	return interceptor(ctx, in, info, handler)
}

func _Points_UpdateVectors_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePointVectors)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PointsServer).UpdateVectors(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Points_UpdateVectors_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PointsServer).UpdateVectors(ctx, req.(*UpdatePointVectors))
	}
	return interceptor(ctx, in, info, handler)
}

func _Points_DeleteVectors_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePointVectors)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PointsServer).DeleteVectors(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Points_DeleteVectors_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PointsServer).DeleteVectors(ctx, req.(*DeletePointVectors))
	}
	return interceptor(ctx, in, info, handler)
}

func _Points_SetPayload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetPayloadPoints)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PointsServer).SetPayload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Points_SetPayload_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PointsServer).SetPayload(ctx, req.(*SetPayloadPoints))
	}
	return interceptor(ctx, in, info, handler)
}

func _Points_OverwritePayload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetPayloadPoints)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PointsServer).OverwritePayload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Points_OverwritePayload_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PointsServer).OverwritePayload(ctx, req.(*SetPayloadPoints))
	}
	return interceptor(ctx, in, info, handler)
}

func _Points_DeletePayload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePayloadPoints)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PointsServer).DeletePayload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Points_DeletePayload_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PointsServer).DeletePayload(ctx, req.(*DeletePayloadPoints))
	}
	return interceptor(ctx, in, info, handler)
}

func _Points_ClearPayload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClearPayloadPoints)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PointsServer).ClearPayload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Points_ClearPayload_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PointsServer).ClearPayload(ctx, req.(*ClearPayloadPoints))
	}
	return interceptor(ctx, in, info, handler)
}

func _Points_CreateFieldIndex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFieldIndexCollection)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PointsServer).CreateFieldIndex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Points_CreateFieldIndex_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PointsServer).CreateFieldIndex(ctx, req.(*CreateFieldIndexCollection))
	}
	return interceptor(ctx, in, info, handler)
}

func _Points_DeleteFieldIndex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteFieldIndexCollection)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PointsServer).DeleteFieldIndex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Points_DeleteFieldIndex_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PointsServer).DeleteFieldIndex(ctx, req.(*DeleteFieldIndexCollection))
	}
	return interceptor(ctx, in, info, handler)
}

func _Points_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchPoints)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PointsServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Points_Search_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PointsServer).Search(ctx, req.(*SearchPoints))
	}
	return interceptor(ctx, in, info, handler)
}

func _Points_SearchBatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchBatchPoints)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PointsServer).SearchBatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Points_SearchBatch_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PointsServer).SearchBatch(ctx, req.(*SearchBatchPoints))
	}
	return interceptor(ctx, in, info, handler)
}

func _Points_SearchGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchPointGroups)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PointsServer).SearchGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Points_SearchGroups_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PointsServer).SearchGroups(ctx, req.(*SearchPointGroups))
	}
	return interceptor(ctx, in, info, handler)
}

func _Points_Scroll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScrollPoints)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PointsServer).Scroll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Points_Scroll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PointsServer).Scroll(ctx, req.(*ScrollPoints))
	}
	return interceptor(ctx, in, info, handler)
}

func _Points_Recommend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecommendPoints)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PointsServer).Recommend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Points_Recommend_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PointsServer).Recommend(ctx, req.(*RecommendPoints))
	}
	return interceptor(ctx, in, info, handler)
}

func _Points_RecommendBatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecommendBatchPoints)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PointsServer).RecommendBatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Points_RecommendBatch_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PointsServer).RecommendBatch(ctx, req.(*RecommendBatchPoints))
	}
	return interceptor(ctx, in, info, handler)
}

func _Points_RecommendGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecommendPointGroups)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PointsServer).RecommendGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Points_RecommendGroups_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PointsServer).RecommendGroups(ctx, req.(*RecommendPointGroups))
	}
	return interceptor(ctx, in, info, handler)
}

func _Points_Count_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountPoints)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PointsServer).Count(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Points_Count_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PointsServer).Count(ctx, req.(*CountPoints))
	}
	return interceptor(ctx, in, info, handler)
}

func _Points_UpdateBatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBatchPoints)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PointsServer).UpdateBatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Points_UpdateBatch_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PointsServer).UpdateBatch(ctx, req.(*UpdateBatchPoints))
	}
	return interceptor(ctx, in, info, handler)
}

// Points_ServiceDesc is the grpc.ServiceDesc for Points service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Points_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "qdrant.Points",
	HandlerType: (*PointsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Upsert",
			Handler:    _Points_Upsert_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Points_Delete_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Points_Get_Handler,
		},
		{
			MethodName: "UpdateVectors",
			Handler:    _Points_UpdateVectors_Handler,
		},
		{
			MethodName: "DeleteVectors",
			Handler:    _Points_DeleteVectors_Handler,
		},
		{
			MethodName: "SetPayload",
			Handler:    _Points_SetPayload_Handler,
		},
		{
			MethodName: "OverwritePayload",
			Handler:    _Points_OverwritePayload_Handler,
		},
		{
			MethodName: "DeletePayload",
			Handler:    _Points_DeletePayload_Handler,
		},
		{
			MethodName: "ClearPayload",
			Handler:    _Points_ClearPayload_Handler,
		},
		{
			MethodName: "CreateFieldIndex",
			Handler:    _Points_CreateFieldIndex_Handler,
		},
		{
			MethodName: "DeleteFieldIndex",
			Handler:    _Points_DeleteFieldIndex_Handler,
		},
		{
			MethodName: "Search",
			Handler:    _Points_Search_Handler,
		},
		{
			MethodName: "SearchBatch",
			Handler:    _Points_SearchBatch_Handler,
		},
		{
			MethodName: "SearchGroups",
			Handler:    _Points_SearchGroups_Handler,
		},
		{
			MethodName: "Scroll",
			Handler:    _Points_Scroll_Handler,
		},
		{
			MethodName: "Recommend",
			Handler:    _Points_Recommend_Handler,
		},
		{
			MethodName: "RecommendBatch",
			Handler:    _Points_RecommendBatch_Handler,
		},
		{
			MethodName: "RecommendGroups",
			Handler:    _Points_RecommendGroups_Handler,
		},
		{
			MethodName: "Count",
			Handler:    _Points_Count_Handler,
		},
		{
			MethodName: "UpdateBatch",
			Handler:    _Points_UpdateBatch_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "points_service.proto",
}
