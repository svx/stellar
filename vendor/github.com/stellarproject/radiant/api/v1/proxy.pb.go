// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/stellarproject/radiant/api/v1/proxy.proto

package api // import "github.com/stellarproject/radiant/api/v1"

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import types "github.com/gogo/protobuf/types"

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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Policy int32

const (
	Policy_RANDOM      Policy = 0
	Policy_LEAST_CONN  Policy = 1
	Policy_ROUND_ROBIN Policy = 2
	Policy_FIRST       Policy = 3
	Policy_IP_HASH     Policy = 4
	Policy_URI_HASH    Policy = 5
	Policy_HEADER      Policy = 6
)

var Policy_name = map[int32]string{
	0: "RANDOM",
	1: "LEAST_CONN",
	2: "ROUND_ROBIN",
	3: "FIRST",
	4: "IP_HASH",
	5: "URI_HASH",
	6: "HEADER",
}
var Policy_value = map[string]int32{
	"RANDOM":      0,
	"LEAST_CONN":  1,
	"ROUND_ROBIN": 2,
	"FIRST":       3,
	"IP_HASH":     4,
	"URI_HASH":    5,
	"HEADER":      6,
}

func (x Policy) String() string {
	return proto.EnumName(Policy_name, int32(x))
}
func (Policy) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_proxy_4c08720d3c16443d, []int{0}
}

type HealthCheck struct {
	HealthCheck          string          `protobuf:"bytes,1,opt,name=health_check,json=healthCheck,proto3" json:"health_check,omitempty"`
	HealthCheckInterval  *types.Duration `protobuf:"bytes,2,opt,name=health_check_interval,json=healthCheckInterval" json:"health_check_interval,omitempty"`
	HealthCheckTimeout   *types.Duration `protobuf:"bytes,3,opt,name=health_check_timeout,json=healthCheckTimeout" json:"health_check_timeout,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *HealthCheck) Reset()         { *m = HealthCheck{} }
func (m *HealthCheck) String() string { return proto.CompactTextString(m) }
func (*HealthCheck) ProtoMessage()    {}
func (*HealthCheck) Descriptor() ([]byte, []int) {
	return fileDescriptor_proxy_4c08720d3c16443d, []int{0}
}
func (m *HealthCheck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheck.Unmarshal(m, b)
}
func (m *HealthCheck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheck.Marshal(b, m, deterministic)
}
func (dst *HealthCheck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheck.Merge(dst, src)
}
func (m *HealthCheck) XXX_Size() int {
	return xxx_messageInfo_HealthCheck.Size(m)
}
func (m *HealthCheck) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheck.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheck proto.InternalMessageInfo

func (m *HealthCheck) GetHealthCheck() string {
	if m != nil {
		return m.HealthCheck
	}
	return ""
}

func (m *HealthCheck) GetHealthCheckInterval() *types.Duration {
	if m != nil {
		return m.HealthCheckInterval
	}
	return nil
}

func (m *HealthCheck) GetHealthCheckTimeout() *types.Duration {
	if m != nil {
		return m.HealthCheckTimeout
	}
	return nil
}

type Server struct {
	Host                 string            `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	Path                 string            `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	TLS                  bool              `protobuf:"varint,3,opt,name=tls,proto3" json:"tls,omitempty"`
	Policy               Policy            `protobuf:"varint,4,opt,name=policy,proto3,enum=io.stellarproject.radiant.api.v1.Policy" json:"policy,omitempty"`
	Timeouts             *types.Duration   `protobuf:"bytes,5,opt,name=timeouts" json:"timeouts,omitempty"`
	Upstreams            []string          `protobuf:"bytes,6,rep,name=upstreams" json:"upstreams,omitempty"`
	HealthCheck          *HealthCheck      `protobuf:"bytes,7,opt,name=health_check,json=healthCheck" json:"health_check,omitempty"`
	InsecureSkipVerify   bool              `protobuf:"varint,8,opt,name=insecure_skip_verify,json=insecureSkipVerify,proto3" json:"insecure_skip_verify,omitempty"`
	Preset               string            `protobuf:"bytes,9,opt,name=preset,proto3" json:"preset,omitempty"`
	ProxyUpstreamHeaders map[string]string `protobuf:"bytes,10,rep,name=proxy_upstream_headers,json=proxyUpstreamHeaders" json:"proxy_upstream_headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Limits               string            `protobuf:"bytes,11,opt,name=limits,proto3" json:"limits,omitempty"`
	ProxyTryDuration     *types.Duration   `protobuf:"bytes,12,opt,name=proxy_try_duration,json=proxyTryDuration" json:"proxy_try_duration,omitempty"`
	ProxyFailTimeout     *types.Duration   `protobuf:"bytes,13,opt,name=proxy_fail_timeout,json=proxyFailTimeout" json:"proxy_fail_timeout,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Server) Reset()         { *m = Server{} }
func (m *Server) String() string { return proto.CompactTextString(m) }
func (*Server) ProtoMessage()    {}
func (*Server) Descriptor() ([]byte, []int) {
	return fileDescriptor_proxy_4c08720d3c16443d, []int{1}
}
func (m *Server) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Server.Unmarshal(m, b)
}
func (m *Server) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Server.Marshal(b, m, deterministic)
}
func (dst *Server) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Server.Merge(dst, src)
}
func (m *Server) XXX_Size() int {
	return xxx_messageInfo_Server.Size(m)
}
func (m *Server) XXX_DiscardUnknown() {
	xxx_messageInfo_Server.DiscardUnknown(m)
}

var xxx_messageInfo_Server proto.InternalMessageInfo

func (m *Server) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *Server) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *Server) GetTLS() bool {
	if m != nil {
		return m.TLS
	}
	return false
}

func (m *Server) GetPolicy() Policy {
	if m != nil {
		return m.Policy
	}
	return Policy_RANDOM
}

func (m *Server) GetTimeouts() *types.Duration {
	if m != nil {
		return m.Timeouts
	}
	return nil
}

func (m *Server) GetUpstreams() []string {
	if m != nil {
		return m.Upstreams
	}
	return nil
}

func (m *Server) GetHealthCheck() *HealthCheck {
	if m != nil {
		return m.HealthCheck
	}
	return nil
}

func (m *Server) GetInsecureSkipVerify() bool {
	if m != nil {
		return m.InsecureSkipVerify
	}
	return false
}

func (m *Server) GetPreset() string {
	if m != nil {
		return m.Preset
	}
	return ""
}

func (m *Server) GetProxyUpstreamHeaders() map[string]string {
	if m != nil {
		return m.ProxyUpstreamHeaders
	}
	return nil
}

func (m *Server) GetLimits() string {
	if m != nil {
		return m.Limits
	}
	return ""
}

func (m *Server) GetProxyTryDuration() *types.Duration {
	if m != nil {
		return m.ProxyTryDuration
	}
	return nil
}

func (m *Server) GetProxyFailTimeout() *types.Duration {
	if m != nil {
		return m.ProxyFailTimeout
	}
	return nil
}

type AddServerRequest struct {
	Server               *Server  `protobuf:"bytes,1,opt,name=server" json:"server,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddServerRequest) Reset()         { *m = AddServerRequest{} }
func (m *AddServerRequest) String() string { return proto.CompactTextString(m) }
func (*AddServerRequest) ProtoMessage()    {}
func (*AddServerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_proxy_4c08720d3c16443d, []int{2}
}
func (m *AddServerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddServerRequest.Unmarshal(m, b)
}
func (m *AddServerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddServerRequest.Marshal(b, m, deterministic)
}
func (dst *AddServerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddServerRequest.Merge(dst, src)
}
func (m *AddServerRequest) XXX_Size() int {
	return xxx_messageInfo_AddServerRequest.Size(m)
}
func (m *AddServerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddServerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddServerRequest proto.InternalMessageInfo

func (m *AddServerRequest) GetServer() *Server {
	if m != nil {
		return m.Server
	}
	return nil
}

type RemoveServerRequest struct {
	Host                 string   `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveServerRequest) Reset()         { *m = RemoveServerRequest{} }
func (m *RemoveServerRequest) String() string { return proto.CompactTextString(m) }
func (*RemoveServerRequest) ProtoMessage()    {}
func (*RemoveServerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_proxy_4c08720d3c16443d, []int{3}
}
func (m *RemoveServerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveServerRequest.Unmarshal(m, b)
}
func (m *RemoveServerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveServerRequest.Marshal(b, m, deterministic)
}
func (dst *RemoveServerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveServerRequest.Merge(dst, src)
}
func (m *RemoveServerRequest) XXX_Size() int {
	return xxx_messageInfo_RemoveServerRequest.Size(m)
}
func (m *RemoveServerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveServerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveServerRequest proto.InternalMessageInfo

func (m *RemoveServerRequest) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

type ReloadRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReloadRequest) Reset()         { *m = ReloadRequest{} }
func (m *ReloadRequest) String() string { return proto.CompactTextString(m) }
func (*ReloadRequest) ProtoMessage()    {}
func (*ReloadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_proxy_4c08720d3c16443d, []int{4}
}
func (m *ReloadRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReloadRequest.Unmarshal(m, b)
}
func (m *ReloadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReloadRequest.Marshal(b, m, deterministic)
}
func (dst *ReloadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReloadRequest.Merge(dst, src)
}
func (m *ReloadRequest) XXX_Size() int {
	return xxx_messageInfo_ReloadRequest.Size(m)
}
func (m *ReloadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReloadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReloadRequest proto.InternalMessageInfo

type ServersRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServersRequest) Reset()         { *m = ServersRequest{} }
func (m *ServersRequest) String() string { return proto.CompactTextString(m) }
func (*ServersRequest) ProtoMessage()    {}
func (*ServersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_proxy_4c08720d3c16443d, []int{5}
}
func (m *ServersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServersRequest.Unmarshal(m, b)
}
func (m *ServersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServersRequest.Marshal(b, m, deterministic)
}
func (dst *ServersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServersRequest.Merge(dst, src)
}
func (m *ServersRequest) XXX_Size() int {
	return xxx_messageInfo_ServersRequest.Size(m)
}
func (m *ServersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ServersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ServersRequest proto.InternalMessageInfo

type ServersResponse struct {
	Servers              []*Server `protobuf:"bytes,1,rep,name=servers" json:"servers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ServersResponse) Reset()         { *m = ServersResponse{} }
func (m *ServersResponse) String() string { return proto.CompactTextString(m) }
func (*ServersResponse) ProtoMessage()    {}
func (*ServersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_proxy_4c08720d3c16443d, []int{6}
}
func (m *ServersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServersResponse.Unmarshal(m, b)
}
func (m *ServersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServersResponse.Marshal(b, m, deterministic)
}
func (dst *ServersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServersResponse.Merge(dst, src)
}
func (m *ServersResponse) XXX_Size() int {
	return xxx_messageInfo_ServersResponse.Size(m)
}
func (m *ServersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ServersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ServersResponse proto.InternalMessageInfo

func (m *ServersResponse) GetServers() []*Server {
	if m != nil {
		return m.Servers
	}
	return nil
}

type ConfigRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConfigRequest) Reset()         { *m = ConfigRequest{} }
func (m *ConfigRequest) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest) ProtoMessage()    {}
func (*ConfigRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_proxy_4c08720d3c16443d, []int{7}
}
func (m *ConfigRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigRequest.Unmarshal(m, b)
}
func (m *ConfigRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigRequest.Marshal(b, m, deterministic)
}
func (dst *ConfigRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigRequest.Merge(dst, src)
}
func (m *ConfigRequest) XXX_Size() int {
	return xxx_messageInfo_ConfigRequest.Size(m)
}
func (m *ConfigRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigRequest proto.InternalMessageInfo

type ConfigResponse struct {
	Data                 []byte   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConfigResponse) Reset()         { *m = ConfigResponse{} }
func (m *ConfigResponse) String() string { return proto.CompactTextString(m) }
func (*ConfigResponse) ProtoMessage()    {}
func (*ConfigResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_proxy_4c08720d3c16443d, []int{8}
}
func (m *ConfigResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigResponse.Unmarshal(m, b)
}
func (m *ConfigResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigResponse.Marshal(b, m, deterministic)
}
func (dst *ConfigResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigResponse.Merge(dst, src)
}
func (m *ConfigResponse) XXX_Size() int {
	return xxx_messageInfo_ConfigResponse.Size(m)
}
func (m *ConfigResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigResponse proto.InternalMessageInfo

func (m *ConfigResponse) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*HealthCheck)(nil), "io.stellarproject.radiant.api.v1.HealthCheck")
	proto.RegisterType((*Server)(nil), "io.stellarproject.radiant.api.v1.Server")
	proto.RegisterMapType((map[string]string)(nil), "io.stellarproject.radiant.api.v1.Server.ProxyUpstreamHeadersEntry")
	proto.RegisterType((*AddServerRequest)(nil), "io.stellarproject.radiant.api.v1.AddServerRequest")
	proto.RegisterType((*RemoveServerRequest)(nil), "io.stellarproject.radiant.api.v1.RemoveServerRequest")
	proto.RegisterType((*ReloadRequest)(nil), "io.stellarproject.radiant.api.v1.ReloadRequest")
	proto.RegisterType((*ServersRequest)(nil), "io.stellarproject.radiant.api.v1.ServersRequest")
	proto.RegisterType((*ServersResponse)(nil), "io.stellarproject.radiant.api.v1.ServersResponse")
	proto.RegisterType((*ConfigRequest)(nil), "io.stellarproject.radiant.api.v1.ConfigRequest")
	proto.RegisterType((*ConfigResponse)(nil), "io.stellarproject.radiant.api.v1.ConfigResponse")
	proto.RegisterEnum("io.stellarproject.radiant.api.v1.Policy", Policy_name, Policy_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ProxyClient is the client API for Proxy service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProxyClient interface {
	AddServer(ctx context.Context, in *AddServerRequest, opts ...grpc.CallOption) (*types.Empty, error)
	RemoveServer(ctx context.Context, in *RemoveServerRequest, opts ...grpc.CallOption) (*types.Empty, error)
	Servers(ctx context.Context, in *ServersRequest, opts ...grpc.CallOption) (*ServersResponse, error)
	Reload(ctx context.Context, in *ReloadRequest, opts ...grpc.CallOption) (*types.Empty, error)
	Config(ctx context.Context, in *ConfigRequest, opts ...grpc.CallOption) (*ConfigResponse, error)
}

type proxyClient struct {
	cc *grpc.ClientConn
}

func NewProxyClient(cc *grpc.ClientConn) ProxyClient {
	return &proxyClient{cc}
}

func (c *proxyClient) AddServer(ctx context.Context, in *AddServerRequest, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/io.stellarproject.radiant.api.v1.Proxy/AddServer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *proxyClient) RemoveServer(ctx context.Context, in *RemoveServerRequest, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/io.stellarproject.radiant.api.v1.Proxy/RemoveServer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *proxyClient) Servers(ctx context.Context, in *ServersRequest, opts ...grpc.CallOption) (*ServersResponse, error) {
	out := new(ServersResponse)
	err := c.cc.Invoke(ctx, "/io.stellarproject.radiant.api.v1.Proxy/Servers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *proxyClient) Reload(ctx context.Context, in *ReloadRequest, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/io.stellarproject.radiant.api.v1.Proxy/Reload", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *proxyClient) Config(ctx context.Context, in *ConfigRequest, opts ...grpc.CallOption) (*ConfigResponse, error) {
	out := new(ConfigResponse)
	err := c.cc.Invoke(ctx, "/io.stellarproject.radiant.api.v1.Proxy/Config", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProxyServer is the server API for Proxy service.
type ProxyServer interface {
	AddServer(context.Context, *AddServerRequest) (*types.Empty, error)
	RemoveServer(context.Context, *RemoveServerRequest) (*types.Empty, error)
	Servers(context.Context, *ServersRequest) (*ServersResponse, error)
	Reload(context.Context, *ReloadRequest) (*types.Empty, error)
	Config(context.Context, *ConfigRequest) (*ConfigResponse, error)
}

func RegisterProxyServer(s *grpc.Server, srv ProxyServer) {
	s.RegisterService(&_Proxy_serviceDesc, srv)
}

func _Proxy_AddServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddServerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProxyServer).AddServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/io.stellarproject.radiant.api.v1.Proxy/AddServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProxyServer).AddServer(ctx, req.(*AddServerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Proxy_RemoveServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveServerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProxyServer).RemoveServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/io.stellarproject.radiant.api.v1.Proxy/RemoveServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProxyServer).RemoveServer(ctx, req.(*RemoveServerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Proxy_Servers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProxyServer).Servers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/io.stellarproject.radiant.api.v1.Proxy/Servers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProxyServer).Servers(ctx, req.(*ServersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Proxy_Reload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReloadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProxyServer).Reload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/io.stellarproject.radiant.api.v1.Proxy/Reload",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProxyServer).Reload(ctx, req.(*ReloadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Proxy_Config_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProxyServer).Config(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/io.stellarproject.radiant.api.v1.Proxy/Config",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProxyServer).Config(ctx, req.(*ConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Proxy_serviceDesc = grpc.ServiceDesc{
	ServiceName: "io.stellarproject.radiant.api.v1.Proxy",
	HandlerType: (*ProxyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddServer",
			Handler:    _Proxy_AddServer_Handler,
		},
		{
			MethodName: "RemoveServer",
			Handler:    _Proxy_RemoveServer_Handler,
		},
		{
			MethodName: "Servers",
			Handler:    _Proxy_Servers_Handler,
		},
		{
			MethodName: "Reload",
			Handler:    _Proxy_Reload_Handler,
		},
		{
			MethodName: "Config",
			Handler:    _Proxy_Config_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/stellarproject/radiant/api/v1/proxy.proto",
}

func init() {
	proto.RegisterFile("github.com/stellarproject/radiant/api/v1/proxy.proto", fileDescriptor_proxy_4c08720d3c16443d)
}

var fileDescriptor_proxy_4c08720d3c16443d = []byte{
	// 827 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x55, 0x51, 0x6f, 0x2a, 0x45,
	0x14, 0x96, 0x52, 0x96, 0x72, 0xe0, 0xb6, 0x9b, 0xb9, 0xd8, 0x6c, 0xd1, 0x28, 0x12, 0x1f, 0xd0,
	0xe8, 0x6e, 0x8b, 0xde, 0xc4, 0xe8, 0x8b, 0xd0, 0x72, 0x2f, 0xc4, 0x7b, 0x69, 0x1d, 0xa8, 0x26,
	0x26, 0x66, 0x33, 0x85, 0x29, 0x8c, 0x2c, 0xcc, 0x3a, 0x33, 0x90, 0xcb, 0x7f, 0xf4, 0x37, 0xf8,
	0xe0, 0xab, 0x0f, 0xfe, 0x05, 0xb3, 0x33, 0xb3, 0x5c, 0xa8, 0x6d, 0xc0, 0xb7, 0x39, 0x67, 0xce,
	0xf9, 0xce, 0x77, 0xe6, 0xfb, 0x16, 0xe0, 0xeb, 0x31, 0x53, 0x93, 0xc5, 0x9d, 0x3f, 0xe4, 0xb3,
	0x40, 0x2a, 0x1a, 0x45, 0x44, 0xc4, 0x82, 0xff, 0x46, 0x87, 0x2a, 0x10, 0x64, 0xc4, 0xc8, 0x5c,
	0x05, 0x24, 0x66, 0xc1, 0xf2, 0x22, 0x88, 0x05, 0x7f, 0xbb, 0xf2, 0x63, 0xc1, 0x15, 0x47, 0x55,
	0xc6, 0xfd, 0xed, 0x6a, 0xdf, 0x56, 0xfb, 0x24, 0x66, 0xfe, 0xf2, 0xa2, 0x52, 0x1e, 0xf3, 0x31,
	0xd7, 0xc5, 0x41, 0x72, 0x32, 0x7d, 0x95, 0x0f, 0xc6, 0x9c, 0x8f, 0x23, 0x1a, 0xe8, 0xe8, 0x6e,
	0x71, 0x1f, 0xd0, 0x59, 0xac, 0x2c, 0x68, 0xe5, 0xa3, 0x87, 0x97, 0xa3, 0x85, 0x20, 0x8a, 0xf1,
	0xb9, 0xb9, 0xaf, 0xfd, 0x91, 0x81, 0x62, 0x87, 0x92, 0x48, 0x4d, 0x2e, 0x27, 0x74, 0x38, 0x45,
	0x9f, 0x40, 0x69, 0xa2, 0xc3, 0x70, 0x98, 0xc4, 0x5e, 0xa6, 0x9a, 0xa9, 0x17, 0x70, 0x71, 0xb2,
	0x51, 0xf2, 0x06, 0xde, 0xdf, 0x2c, 0x09, 0xd9, 0x5c, 0x51, 0xb1, 0x24, 0x91, 0x77, 0x50, 0xcd,
	0xd4, 0x8b, 0x8d, 0x33, 0xdf, 0x8c, 0xf4, 0xd3, 0x91, 0xfe, 0x95, 0x1d, 0x89, 0x9f, 0x6f, 0xc0,
	0x74, 0x6d, 0x17, 0xfa, 0x01, 0xca, 0x5b, 0x70, 0x8a, 0xcd, 0x28, 0x5f, 0x28, 0x2f, 0xbb, 0x0b,
	0x0d, 0x6d, 0xa0, 0x0d, 0x4c, 0x53, 0xed, 0xef, 0x1c, 0x38, 0x7d, 0x2a, 0x96, 0x54, 0x20, 0x04,
	0x87, 0x13, 0x2e, 0x95, 0xdd, 0x40, 0x9f, 0x93, 0x5c, 0x4c, 0xd4, 0x44, 0x33, 0x2d, 0x60, 0x7d,
	0x46, 0x67, 0x90, 0x55, 0x91, 0xd4, 0xe3, 0x8e, 0x5a, 0xf9, 0xbf, 0xfe, 0xfc, 0x38, 0x3b, 0x78,
	0xdd, 0xc7, 0x49, 0x0e, 0x7d, 0x0f, 0x4e, 0xcc, 0x23, 0x36, 0x5c, 0x79, 0x87, 0xd5, 0x4c, 0xfd,
	0xb8, 0x51, 0xf7, 0x77, 0x49, 0xe4, 0xdf, 0xe8, 0x7a, 0x6c, 0xfb, 0xd0, 0x0b, 0x38, 0xb2, 0xfb,
	0x48, 0x2f, 0xb7, 0x6b, 0xa1, 0x75, 0x29, 0xfa, 0x10, 0x0a, 0x8b, 0x58, 0x2a, 0x41, 0xc9, 0x4c,
	0x7a, 0x4e, 0x35, 0x5b, 0x2f, 0xe0, 0x77, 0x09, 0x74, 0xf3, 0x40, 0xa3, 0xbc, 0x06, 0xfe, 0x72,
	0x37, 0xb9, 0x0d, 0xa1, 0xb7, 0x25, 0x3d, 0x87, 0x32, 0x9b, 0x4b, 0x3a, 0x5c, 0x08, 0x1a, 0xca,
	0x29, 0x8b, 0xc3, 0x25, 0x15, 0xec, 0x7e, 0xe5, 0x1d, 0x25, 0x8f, 0x82, 0x51, 0x7a, 0xd7, 0x9f,
	0xb2, 0xf8, 0x27, 0x7d, 0x83, 0x4e, 0xc1, 0x89, 0x05, 0x95, 0x54, 0x79, 0x05, 0xfd, 0x96, 0x36,
	0x42, 0x6f, 0xe1, 0x54, 0x7b, 0x3a, 0x4c, 0xe9, 0x86, 0x13, 0x4a, 0x46, 0x54, 0x48, 0x0f, 0xaa,
	0xd9, 0x7a, 0xb1, 0xd1, 0xda, 0xcd, 0xd2, 0xe8, 0xe7, 0xdf, 0x24, 0x30, 0xb7, 0x16, 0xa5, 0x63,
	0x40, 0xda, 0x73, 0x25, 0x56, 0xb8, 0x1c, 0x3f, 0x72, 0x95, 0x30, 0x8a, 0xd8, 0x8c, 0x29, 0xe9,
	0x15, 0x0d, 0x23, 0x13, 0xa1, 0x57, 0x80, 0x0c, 0x23, 0x25, 0x56, 0x61, 0xea, 0x7e, 0xaf, 0xb4,
	0x4b, 0x0c, 0x57, 0x37, 0x0d, 0xc4, 0x2a, 0xcd, 0xbc, 0x03, 0xba, 0x27, 0x2c, 0x5a, 0xdb, 0xf4,
	0xd9, 0x7e, 0x40, 0x2f, 0x09, 0x8b, 0xac, 0x49, 0x2b, 0xaf, 0xe0, 0xec, 0xc9, 0xe5, 0x90, 0x0b,
	0xd9, 0x29, 0x5d, 0x59, 0xd7, 0x26, 0x47, 0x54, 0x86, 0xdc, 0x92, 0x44, 0x0b, 0x6a, 0x5d, 0x6b,
	0x82, 0x6f, 0x0f, 0xbe, 0xc9, 0xd4, 0x06, 0xe0, 0x36, 0x47, 0x23, 0xf3, 0x5e, 0x98, 0xfe, 0xbe,
	0xa0, 0x52, 0x25, 0x9e, 0x95, 0x3a, 0xa1, 0x21, 0x8a, 0xfb, 0x78, 0xd6, 0x02, 0xd8, 0xbe, 0xda,
	0x67, 0xf0, 0x1c, 0xd3, 0x19, 0x5f, 0xd2, 0x6d, 0xe0, 0x47, 0xbe, 0xa7, 0xda, 0x09, 0x3c, 0xc3,
	0x34, 0xe2, 0x64, 0x64, 0x8b, 0x6a, 0x2e, 0x1c, 0x9b, 0x2e, 0x99, 0x66, 0x6e, 0xe1, 0x64, 0x9d,
	0x91, 0x31, 0x9f, 0x4b, 0x8a, 0x5a, 0x90, 0x37, 0xa3, 0xa4, 0x97, 0xd1, 0xa6, 0xd8, 0x9f, 0x63,
	0xda, 0x98, 0x4c, 0xbe, 0xe4, 0xf3, 0x7b, 0x36, 0x4e, 0xe7, 0x7c, 0x0a, 0xc7, 0x69, 0xc2, 0x8e,
	0x41, 0x70, 0x38, 0x22, 0x8a, 0x68, 0xc2, 0x25, 0xac, 0xcf, 0x9f, 0x8f, 0xc1, 0x31, 0x5f, 0x28,
	0x02, 0x70, 0x70, 0xb3, 0x77, 0x75, 0xfd, 0xc6, 0x7d, 0x0f, 0x1d, 0x03, 0xbc, 0x6e, 0x37, 0xfb,
	0x83, 0xf0, 0xf2, 0xba, 0xd7, 0x73, 0x33, 0xe8, 0x04, 0x8a, 0xf8, 0xfa, 0xb6, 0x77, 0x15, 0xe2,
	0xeb, 0x56, 0xb7, 0xe7, 0x1e, 0xa0, 0x02, 0xe4, 0x5e, 0x76, 0x71, 0x7f, 0xe0, 0x66, 0x51, 0x11,
	0xf2, 0xdd, 0x9b, 0xb0, 0xd3, 0xec, 0x77, 0xdc, 0x43, 0x54, 0x82, 0xa3, 0x5b, 0xdc, 0x35, 0x51,
	0x2e, 0x81, 0xec, 0xb4, 0x9b, 0x57, 0x6d, 0xec, 0x3a, 0x8d, 0x7f, 0xb2, 0x90, 0xd3, 0x22, 0xa3,
	0x9f, 0xa1, 0xb0, 0x16, 0x09, 0x35, 0x76, 0x6f, 0xfa, 0x50, 0xd1, 0xca, 0xe9, 0x7f, 0xbc, 0xd5,
	0x4e, 0x7e, 0xe0, 0xd1, 0xaf, 0x50, 0xda, 0xd4, 0x09, 0xbd, 0xd8, 0x8d, 0xfd, 0x88, 0xae, 0x4f,
	0xc2, 0xcf, 0x21, 0x6f, 0x85, 0x43, 0xe7, 0xfb, 0xea, 0x93, 0xaa, 0x5e, 0xb9, 0xf8, 0x1f, 0x1d,
	0x56, 0xae, 0x1f, 0xc1, 0x31, 0x5e, 0x42, 0xc1, 0x3e, 0x8b, 0x6c, 0xb8, 0xee, 0xc9, 0x15, 0xa6,
	0xe0, 0x18, 0x4f, 0xec, 0x03, 0xb9, 0x65, 0xa7, 0xca, 0xf9, 0xfe, 0x0d, 0x86, 0x7f, 0xcb, 0xff,
	0xe5, 0x8b, 0x7d, 0xff, 0xf6, 0xbf, 0x23, 0x31, 0xbb, 0x73, 0x34, 0xd9, 0xaf, 0xfe, 0x0d, 0x00,
	0x00, 0xff, 0xff, 0x81, 0x91, 0xf4, 0xf3, 0x2d, 0x08, 0x00, 0x00,
}