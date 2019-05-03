// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/ehazlett/stellar/api/services/cluster/v1/cluster.proto

package cluster

import (
	context "context"
	fmt "fmt"
	v1 "github.com/ehazlett/stellar/api/services/health/v1"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type InfoRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InfoRequest) Reset()         { *m = InfoRequest{} }
func (m *InfoRequest) String() string { return proto.CompactTextString(m) }
func (*InfoRequest) ProtoMessage()    {}
func (*InfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c077b095128b9733, []int{0}
}
func (m *InfoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InfoRequest.Unmarshal(m, b)
}
func (m *InfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InfoRequest.Marshal(b, m, deterministic)
}
func (m *InfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InfoRequest.Merge(m, src)
}
func (m *InfoRequest) XXX_Size() int {
	return xxx_messageInfo_InfoRequest.Size(m)
}
func (m *InfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_InfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_InfoRequest proto.InternalMessageInfo

type InfoResponse struct {
	ID                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InfoResponse) Reset()         { *m = InfoResponse{} }
func (m *InfoResponse) String() string { return proto.CompactTextString(m) }
func (*InfoResponse) ProtoMessage()    {}
func (*InfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c077b095128b9733, []int{1}
}
func (m *InfoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InfoResponse.Unmarshal(m, b)
}
func (m *InfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InfoResponse.Marshal(b, m, deterministic)
}
func (m *InfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InfoResponse.Merge(m, src)
}
func (m *InfoResponse) XXX_Size() int {
	return xxx_messageInfo_InfoResponse.Size(m)
}
func (m *InfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_InfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_InfoResponse proto.InternalMessageInfo

func (m *InfoResponse) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

type NodesRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodesRequest) Reset()         { *m = NodesRequest{} }
func (m *NodesRequest) String() string { return proto.CompactTextString(m) }
func (*NodesRequest) ProtoMessage()    {}
func (*NodesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c077b095128b9733, []int{2}
}
func (m *NodesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodesRequest.Unmarshal(m, b)
}
func (m *NodesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodesRequest.Marshal(b, m, deterministic)
}
func (m *NodesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodesRequest.Merge(m, src)
}
func (m *NodesRequest) XXX_Size() int {
	return xxx_messageInfo_NodesRequest.Size(m)
}
func (m *NodesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NodesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NodesRequest proto.InternalMessageInfo

type Node struct {
	ID                   string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Address              string            `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Labels               map[string]string `protobuf:"bytes,3,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Node) Reset()         { *m = Node{} }
func (m *Node) String() string { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()    {}
func (*Node) Descriptor() ([]byte, []int) {
	return fileDescriptor_c077b095128b9733, []int{3}
}
func (m *Node) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Node.Unmarshal(m, b)
}
func (m *Node) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Node.Marshal(b, m, deterministic)
}
func (m *Node) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Node.Merge(m, src)
}
func (m *Node) XXX_Size() int {
	return xxx_messageInfo_Node.Size(m)
}
func (m *Node) XXX_DiscardUnknown() {
	xxx_messageInfo_Node.DiscardUnknown(m)
}

var xxx_messageInfo_Node proto.InternalMessageInfo

func (m *Node) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Node) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Node) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

type NodesResponse struct {
	Nodes                []*Node  `protobuf:"bytes,1,rep,name=nodes,proto3" json:"nodes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodesResponse) Reset()         { *m = NodesResponse{} }
func (m *NodesResponse) String() string { return proto.CompactTextString(m) }
func (*NodesResponse) ProtoMessage()    {}
func (*NodesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c077b095128b9733, []int{4}
}
func (m *NodesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodesResponse.Unmarshal(m, b)
}
func (m *NodesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodesResponse.Marshal(b, m, deterministic)
}
func (m *NodesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodesResponse.Merge(m, src)
}
func (m *NodesResponse) XXX_Size() int {
	return xxx_messageInfo_NodesResponse.Size(m)
}
func (m *NodesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NodesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NodesResponse proto.InternalMessageInfo

func (m *NodesResponse) GetNodes() []*Node {
	if m != nil {
		return m.Nodes
	}
	return nil
}

type HealthRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HealthRequest) Reset()         { *m = HealthRequest{} }
func (m *HealthRequest) String() string { return proto.CompactTextString(m) }
func (*HealthRequest) ProtoMessage()    {}
func (*HealthRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c077b095128b9733, []int{5}
}
func (m *HealthRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthRequest.Unmarshal(m, b)
}
func (m *HealthRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthRequest.Marshal(b, m, deterministic)
}
func (m *HealthRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthRequest.Merge(m, src)
}
func (m *HealthRequest) XXX_Size() int {
	return xxx_messageInfo_HealthRequest.Size(m)
}
func (m *HealthRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HealthRequest proto.InternalMessageInfo

type NodeHealth struct {
	Node                 *Node          `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	Health               *v1.NodeHealth `protobuf:"bytes,2,opt,name=health,proto3" json:"health,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *NodeHealth) Reset()         { *m = NodeHealth{} }
func (m *NodeHealth) String() string { return proto.CompactTextString(m) }
func (*NodeHealth) ProtoMessage()    {}
func (*NodeHealth) Descriptor() ([]byte, []int) {
	return fileDescriptor_c077b095128b9733, []int{6}
}
func (m *NodeHealth) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeHealth.Unmarshal(m, b)
}
func (m *NodeHealth) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeHealth.Marshal(b, m, deterministic)
}
func (m *NodeHealth) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeHealth.Merge(m, src)
}
func (m *NodeHealth) XXX_Size() int {
	return xxx_messageInfo_NodeHealth.Size(m)
}
func (m *NodeHealth) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeHealth.DiscardUnknown(m)
}

var xxx_messageInfo_NodeHealth proto.InternalMessageInfo

func (m *NodeHealth) GetNode() *Node {
	if m != nil {
		return m.Node
	}
	return nil
}

func (m *NodeHealth) GetHealth() *v1.NodeHealth {
	if m != nil {
		return m.Health
	}
	return nil
}

type HealthResponse struct {
	Nodes                []*NodeHealth `protobuf:"bytes,1,rep,name=nodes,proto3" json:"nodes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *HealthResponse) Reset()         { *m = HealthResponse{} }
func (m *HealthResponse) String() string { return proto.CompactTextString(m) }
func (*HealthResponse) ProtoMessage()    {}
func (*HealthResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c077b095128b9733, []int{7}
}
func (m *HealthResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthResponse.Unmarshal(m, b)
}
func (m *HealthResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthResponse.Marshal(b, m, deterministic)
}
func (m *HealthResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthResponse.Merge(m, src)
}
func (m *HealthResponse) XXX_Size() int {
	return xxx_messageInfo_HealthResponse.Size(m)
}
func (m *HealthResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HealthResponse proto.InternalMessageInfo

func (m *HealthResponse) GetNodes() []*NodeHealth {
	if m != nil {
		return m.Nodes
	}
	return nil
}

func init() {
	proto.RegisterType((*InfoRequest)(nil), "stellar.services.cluster.v1.InfoRequest")
	proto.RegisterType((*InfoResponse)(nil), "stellar.services.cluster.v1.InfoResponse")
	proto.RegisterType((*NodesRequest)(nil), "stellar.services.cluster.v1.NodesRequest")
	proto.RegisterType((*Node)(nil), "stellar.services.cluster.v1.Node")
	proto.RegisterMapType((map[string]string)(nil), "stellar.services.cluster.v1.Node.LabelsEntry")
	proto.RegisterType((*NodesResponse)(nil), "stellar.services.cluster.v1.NodesResponse")
	proto.RegisterType((*HealthRequest)(nil), "stellar.services.cluster.v1.HealthRequest")
	proto.RegisterType((*NodeHealth)(nil), "stellar.services.cluster.v1.NodeHealth")
	proto.RegisterType((*HealthResponse)(nil), "stellar.services.cluster.v1.HealthResponse")
}

func init() {
	proto.RegisterFile("github.com/ehazlett/stellar/api/services/cluster/v1/cluster.proto", fileDescriptor_c077b095128b9733)
}

var fileDescriptor_c077b095128b9733 = []byte{
	// 485 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xdd, 0x8a, 0xd3, 0x40,
	0x14, 0x36, 0x69, 0x9b, 0xc5, 0xd3, 0xed, 0xaa, 0xe3, 0x22, 0x31, 0x0a, 0xae, 0x11, 0xd6, 0xae,
	0x62, 0xc2, 0x56, 0xc4, 0x9f, 0x65, 0x05, 0x57, 0x17, 0x76, 0x41, 0x14, 0x72, 0x25, 0x7a, 0x35,
	0x6d, 0x8e, 0x69, 0x70, 0xcc, 0xd4, 0xcc, 0xb4, 0x50, 0x2f, 0xf5, 0x11, 0x7c, 0x1b, 0xaf, 0x7d,
	0x03, 0xef, 0xbd, 0x10, 0x1f, 0x44, 0xe6, 0x0f, 0x5a, 0xc4, 0x36, 0xec, 0xdd, 0x37, 0x27, 0xdf,
	0x77, 0xce, 0x77, 0xe6, 0x9c, 0x09, 0x3c, 0x2b, 0x4a, 0x39, 0x9e, 0x0e, 0x93, 0x11, 0xff, 0x98,
	0xe2, 0x98, 0x7e, 0x66, 0x28, 0x65, 0x2a, 0x24, 0x32, 0x46, 0xeb, 0x94, 0x4e, 0xca, 0x54, 0x60,
	0x3d, 0x2b, 0x47, 0x28, 0xd2, 0x11, 0x9b, 0x0a, 0x89, 0x75, 0x3a, 0xdb, 0x77, 0x30, 0x99, 0xd4,
	0x5c, 0x72, 0x72, 0xcd, 0xd2, 0x13, 0x47, 0x4d, 0xdc, 0xf7, 0xd9, 0x7e, 0x74, 0xbd, 0xe0, 0xbc,
	0x60, 0xa8, 0x53, 0xd1, 0xaa, 0xe2, 0x92, 0xca, 0x92, 0x57, 0xc2, 0x48, 0xa3, 0xed, 0x82, 0x17,
	0x5c, 0xc3, 0x54, 0x21, 0x1b, 0xbd, 0xb5, 0x54, 0x77, 0x8c, 0x94, 0xc9, 0xb1, 0x2a, 0x6b, 0x90,
	0x21, 0xc5, 0x3d, 0xe8, 0x9e, 0x56, 0xef, 0x79, 0x86, 0x9f, 0xa6, 0x28, 0x64, 0xbc, 0x0b, 0x9b,
	0xe6, 0x28, 0x26, 0xbc, 0x12, 0x48, 0xae, 0x80, 0x5f, 0xe6, 0xa1, 0xb7, 0xe3, 0xf5, 0xcf, 0x1f,
	0x05, 0xbf, 0x7f, 0xdd, 0xf0, 0x4f, 0x5f, 0x64, 0x7e, 0x99, 0xc7, 0x5b, 0xb0, 0xf9, 0x8a, 0xe7,
	0x28, 0x9c, 0xee, 0xbb, 0x07, 0x6d, 0x15, 0xf8, 0x9f, 0x80, 0x84, 0xb0, 0x41, 0xf3, 0xbc, 0x46,
	0x21, 0x42, 0x5f, 0x7d, 0xcc, 0xdc, 0x91, 0x1c, 0x43, 0xc0, 0xe8, 0x10, 0x99, 0x08, 0x5b, 0x3b,
	0xad, 0x7e, 0x77, 0x70, 0x2f, 0x59, 0x71, 0x11, 0x89, 0x2a, 0x92, 0xbc, 0xd4, 0xfc, 0xe3, 0x4a,
	0xd6, 0xf3, 0xcc, 0x8a, 0xa3, 0xc7, 0xd0, 0x5d, 0x08, 0x93, 0x8b, 0xd0, 0xfa, 0x80, 0x73, 0x63,
	0x24, 0x53, 0x90, 0x6c, 0x43, 0x67, 0x46, 0xd9, 0x14, 0x6d, 0x7d, 0x73, 0x78, 0xe2, 0x3f, 0xf2,
	0xe2, 0x13, 0xe8, 0xd9, 0x66, 0x6c, 0xd7, 0x0f, 0xa1, 0x53, 0xa9, 0x40, 0xe8, 0x69, 0x47, 0x37,
	0xd7, 0x3a, 0xca, 0x0c, 0x3f, 0xbe, 0x00, 0xbd, 0x13, 0x7d, 0xbb, 0xee, 0x5e, 0xbe, 0x7a, 0x00,
	0x8a, 0x60, 0xa2, 0xe4, 0x01, 0xb4, 0x15, 0x51, 0xdb, 0x6a, 0x94, 0x57, 0xd3, 0xc9, 0x53, 0x08,
	0xcc, 0xd0, 0xb4, 0xf7, 0xee, 0x60, 0xf7, 0x5f, 0xa1, 0x1d, 0xaa, 0xd5, 0x59, 0x13, 0x56, 0x15,
	0xbf, 0x86, 0x2d, 0x67, 0xcb, 0x76, 0x78, 0xb8, 0xdc, 0xe1, 0xed, 0xb5, 0x4e, 0xac, 0xde, 0xa8,
	0x06, 0x3f, 0x7c, 0xd8, 0x78, 0x6e, 0x08, 0xe4, 0x1d, 0xb4, 0xd5, 0xca, 0x90, 0xfe, 0xca, 0x1c,
	0x0b, 0x4b, 0x16, 0xed, 0x35, 0x60, 0x5a, 0x9f, 0x73, 0xe8, 0xe8, 0xd1, 0x90, 0xbd, 0xb5, 0x0e,
	0xdd, 0x2e, 0x46, 0x77, 0x9a, 0x50, 0x4d, 0xfe, 0xf8, 0xea, 0x97, 0x9f, 0x7f, 0xbe, 0xf9, 0x97,
	0xc9, 0xa5, 0x85, 0xf7, 0x98, 0xea, 0x1e, 0x09, 0x85, 0xc0, 0x4e, 0x6d, 0x75, 0xc2, 0xa5, 0x81,
	0x47, 0x77, 0x1b, 0x71, 0x4d, 0xf5, 0xa3, 0xc3, 0xb7, 0x07, 0x67, 0xf8, 0x6f, 0x1c, 0x58, 0xf8,
	0xe6, 0xdc, 0x30, 0xd0, 0x8f, 0xf8, 0xfe, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc3, 0x15, 0x1f,
	0x3c, 0x7f, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ClusterClient is the client API for Cluster service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ClusterClient interface {
	Info(ctx context.Context, in *InfoRequest, opts ...grpc.CallOption) (*InfoResponse, error)
	Nodes(ctx context.Context, in *NodesRequest, opts ...grpc.CallOption) (*NodesResponse, error)
	Health(ctx context.Context, in *HealthRequest, opts ...grpc.CallOption) (*HealthResponse, error)
}

type clusterClient struct {
	cc *grpc.ClientConn
}

func NewClusterClient(cc *grpc.ClientConn) ClusterClient {
	return &clusterClient{cc}
}

func (c *clusterClient) Info(ctx context.Context, in *InfoRequest, opts ...grpc.CallOption) (*InfoResponse, error) {
	out := new(InfoResponse)
	err := c.cc.Invoke(ctx, "/stellar.services.cluster.v1.Cluster/Info", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterClient) Nodes(ctx context.Context, in *NodesRequest, opts ...grpc.CallOption) (*NodesResponse, error) {
	out := new(NodesResponse)
	err := c.cc.Invoke(ctx, "/stellar.services.cluster.v1.Cluster/Nodes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterClient) Health(ctx context.Context, in *HealthRequest, opts ...grpc.CallOption) (*HealthResponse, error) {
	out := new(HealthResponse)
	err := c.cc.Invoke(ctx, "/stellar.services.cluster.v1.Cluster/Health", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClusterServer is the server API for Cluster service.
type ClusterServer interface {
	Info(context.Context, *InfoRequest) (*InfoResponse, error)
	Nodes(context.Context, *NodesRequest) (*NodesResponse, error)
	Health(context.Context, *HealthRequest) (*HealthResponse, error)
}

func RegisterClusterServer(s *grpc.Server, srv ClusterServer) {
	s.RegisterService(&_Cluster_serviceDesc, srv)
}

func _Cluster_Info_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServer).Info(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stellar.services.cluster.v1.Cluster/Info",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServer).Info(ctx, req.(*InfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cluster_Nodes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServer).Nodes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stellar.services.cluster.v1.Cluster/Nodes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServer).Nodes(ctx, req.(*NodesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cluster_Health_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServer).Health(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stellar.services.cluster.v1.Cluster/Health",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServer).Health(ctx, req.(*HealthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Cluster_serviceDesc = grpc.ServiceDesc{
	ServiceName: "stellar.services.cluster.v1.Cluster",
	HandlerType: (*ClusterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Info",
			Handler:    _Cluster_Info_Handler,
		},
		{
			MethodName: "Nodes",
			Handler:    _Cluster_Nodes_Handler,
		},
		{
			MethodName: "Health",
			Handler:    _Cluster_Health_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/ehazlett/stellar/api/services/cluster/v1/cluster.proto",
}
