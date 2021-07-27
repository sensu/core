// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/sensu/sensu-go/api/core/v2/handler.proto

package v2

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	github_com_golang_protobuf_proto "github.com/golang/protobuf/proto"
	proto "github.com/golang/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
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

// A Handler is a handler specification.
type Handler struct {
	// Metadata contains the name, namespace, labels and annotations of the
	// handler
	ObjectMeta `protobuf:"bytes,1,opt,name=metadata,proto3,embedded=metadata" json:"metadata,omitempty"`
	// Type is the handler type, i.e. pipe.
	Type string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	// Mutator is the handler event data mutator.
	Mutator string `protobuf:"bytes,3,opt,name=mutator,proto3" json:"mutator,omitempty"`
	// Command is the command to be executed for a pipe handler.
	Command string `protobuf:"bytes,4,opt,name=command,proto3" json:"command,omitempty"`
	// Timeout is the handler timeout in seconds.
	Timeout uint32 `protobuf:"varint,5,opt,name=timeout,proto3" json:"timeout"`
	// Socket contains configuration for a TCP or UDP handler.
	Socket *HandlerSocket `protobuf:"bytes,6,opt,name=socket,proto3" json:"socket,omitempty"`
	// Handlers is a list of handlers for a handler set.
	Handlers []string `protobuf:"bytes,7,rep,name=handlers,proto3" json:"handlers"`
	// Filters is a list of filters name to evaluate before executing this
	// handler
	Filters []string `protobuf:"bytes,8,rep,name=filters,proto3" json:"filters"`
	// EnvVars is a list of environment variables to use with command execution
	EnvVars []string `protobuf:"bytes,9,rep,name=env_vars,json=envVars,proto3" json:"env_vars"`
	// RuntimeAssets are a list of assets required to execute a handler.
	RuntimeAssets []string `protobuf:"bytes,13,rep,name=runtime_assets,json=runtimeAssets,proto3" json:"runtime_assets"`
	// Secrets is the list of Sensu secrets to set for the handler's
	// execution environment.
	Secrets              []*Secret `protobuf:"bytes,14,rep,name=secrets,proto3" json:"secrets"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Handler) Reset()         { *m = Handler{} }
func (m *Handler) String() string { return proto.CompactTextString(m) }
func (*Handler) ProtoMessage()    {}
func (*Handler) Descriptor() ([]byte, []int) {
	return fileDescriptor_a415b3439792b693, []int{0}
}
func (m *Handler) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Handler) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Handler.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Handler) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Handler.Merge(m, src)
}
func (m *Handler) XXX_Size() int {
	return m.Size()
}
func (m *Handler) XXX_DiscardUnknown() {
	xxx_messageInfo_Handler.DiscardUnknown(m)
}

var xxx_messageInfo_Handler proto.InternalMessageInfo

// HandlerSocket contains configuration for a TCP or UDP handler.
type HandlerSocket struct {
	// Host is the socket peer address.
	Host string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	// Port is the socket peer port.
	Port                 uint32   `protobuf:"varint,2,opt,name=port,proto3" json:"port"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HandlerSocket) Reset()         { *m = HandlerSocket{} }
func (m *HandlerSocket) String() string { return proto.CompactTextString(m) }
func (*HandlerSocket) ProtoMessage()    {}
func (*HandlerSocket) Descriptor() ([]byte, []int) {
	return fileDescriptor_a415b3439792b693, []int{1}
}
func (m *HandlerSocket) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HandlerSocket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HandlerSocket.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HandlerSocket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HandlerSocket.Merge(m, src)
}
func (m *HandlerSocket) XXX_Size() int {
	return m.Size()
}
func (m *HandlerSocket) XXX_DiscardUnknown() {
	xxx_messageInfo_HandlerSocket.DiscardUnknown(m)
}

var xxx_messageInfo_HandlerSocket proto.InternalMessageInfo

func (m *HandlerSocket) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *HandlerSocket) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func init() {
	proto.RegisterType((*Handler)(nil), "sensu.core.v2.Handler")
	proto.RegisterType((*HandlerSocket)(nil), "sensu.core.v2.HandlerSocket")
}

func init() {
	proto.RegisterFile("github.com/sensu/sensu-go/api/core/v2/handler.proto", fileDescriptor_a415b3439792b693)
}

var fileDescriptor_a415b3439792b693 = []byte{
	// 501 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0xcd, 0x36, 0x21, 0x76, 0x36, 0xb8, 0x87, 0x95, 0x90, 0x96, 0xaa, 0xb2, 0xad, 0x4a, 0x08,
	0x1f, 0xc0, 0xa6, 0x0e, 0x17, 0x2a, 0x0e, 0xd4, 0x27, 0x2e, 0x08, 0x69, 0x2b, 0x38, 0x70, 0xa9,
	0x36, 0xce, 0x36, 0x09, 0xd4, 0xde, 0xc8, 0xbb, 0xb6, 0xd4, 0x3f, 0xe8, 0x27, 0x70, 0xec, 0xb1,
	0x9f, 0xc0, 0x27, 0xe4, 0xd8, 0x2f, 0xb0, 0x20, 0xdc, 0xfc, 0x05, 0x1c, 0x91, 0xc7, 0x76, 0xa0,
	0x39, 0xf5, 0xb2, 0x7a, 0xf3, 0xe6, 0xcd, 0x78, 0x9e, 0x67, 0xf0, 0x64, 0xbe, 0xd4, 0x8b, 0x7c,
	0xea, 0xc7, 0x32, 0x09, 0x94, 0x48, 0x55, 0xde, 0xbc, 0x2f, 0xe7, 0x32, 0xe0, 0xab, 0x65, 0x10,
	0xcb, 0x4c, 0x04, 0x45, 0x18, 0x2c, 0x78, 0x3a, 0xbb, 0x14, 0x99, 0xbf, 0xca, 0xa4, 0x96, 0xc4,
	0x02, 0x8d, 0x5f, 0x27, 0xfd, 0x22, 0x3c, 0x78, 0xfd, 0x5f, 0x8f, 0xb9, 0x9c, 0xcb, 0x00, 0x54,
	0xd3, 0xfc, 0xe2, 0x5d, 0x71, 0xec, 0x4f, 0xfc, 0x63, 0x20, 0x81, 0x03, 0xd4, 0x34, 0x39, 0x78,
	0xf5, 0xb0, 0x2f, 0x27, 0x42, 0xf3, 0xb6, 0x22, 0x7c, 0x58, 0x85, 0x12, 0x71, 0x26, 0x74, 0x53,
	0x73, 0x74, 0x3d, 0xc0, 0xc6, 0xfb, 0x66, 0x78, 0xf2, 0x09, 0x9b, 0x75, 0xb7, 0x19, 0xd7, 0x9c,
	0x22, 0x17, 0x79, 0xe3, 0xf0, 0xa9, 0x7f, 0xcf, 0x89, 0xff, 0x71, 0xfa, 0x55, 0xc4, 0xfa, 0x83,
	0xd0, 0x3c, 0xb2, 0xd7, 0xa5, 0xd3, 0xbb, 0x2b, 0x1d, 0x54, 0x95, 0x0e, 0xe9, 0xca, 0x5e, 0xc8,
	0x64, 0xa9, 0x45, 0xb2, 0xd2, 0x57, 0x6c, 0xdb, 0x8a, 0x10, 0x3c, 0xd0, 0x57, 0x2b, 0x41, 0xf7,
	0x5c, 0xe4, 0x8d, 0x18, 0x60, 0x42, 0xb1, 0x91, 0xe4, 0x9a, 0x6b, 0x99, 0xd1, 0x3e, 0xd0, 0x5d,
	0x58, 0x67, 0x62, 0x99, 0x24, 0x3c, 0x9d, 0xd1, 0x41, 0x93, 0x69, 0x43, 0xf2, 0x0c, 0x1b, 0x7a,
	0x99, 0x08, 0x99, 0x6b, 0xfa, 0xc8, 0x45, 0x9e, 0x15, 0x8d, 0xab, 0xd2, 0xe9, 0x28, 0xd6, 0x01,
	0x72, 0x82, 0x87, 0x4a, 0xc6, 0xdf, 0x84, 0xa6, 0x43, 0xf0, 0x70, 0xb8, 0xe3, 0xa1, 0x75, 0x7b,
	0x06, 0x9a, 0x68, 0xb0, 0x2e, 0x1d, 0xc4, 0xda, 0x0a, 0xe2, 0x61, 0xb3, 0xdd, 0xa4, 0xa2, 0x86,
	0xdb, 0xf7, 0x46, 0xd1, 0xe3, 0xaa, 0x74, 0xb6, 0x1c, 0xdb, 0xa2, 0x7a, 0x98, 0x8b, 0xe5, 0xa5,
	0xae, 0x85, 0x26, 0x08, 0x61, 0x98, 0x96, 0x62, 0x1d, 0x20, 0xcf, 0xb1, 0x29, 0xd2, 0xe2, 0xbc,
	0xe0, 0x99, 0xa2, 0xa3, 0x7f, 0x0d, 0x3b, 0x8e, 0x19, 0x22, 0x2d, 0x3e, 0xf3, 0x4c, 0x91, 0x37,
	0x78, 0x3f, 0xcb, 0xd3, 0xda, 0xc3, 0x39, 0x57, 0x4a, 0x68, 0x45, 0x2d, 0x90, 0x93, 0xaa, 0x74,
	0x76, 0x32, 0xcc, 0x6a, 0xe3, 0x53, 0x08, 0xc9, 0x5b, 0x6c, 0x34, 0x2b, 0x55, 0x74, 0xdf, 0xed,
	0x7b, 0xe3, 0xf0, 0xc9, 0x8e, 0xe3, 0x33, 0xc8, 0x36, 0x13, 0xb6, 0x4a, 0xd6, 0x81, 0x13, 0xf3,
	0xfa, 0xc6, 0xe9, 0xdd, 0xde, 0x38, 0xe8, 0xe8, 0x14, 0x5b, 0xf7, 0xfe, 0x4d, 0xbd, 0xb8, 0x85,
	0x54, 0x1a, 0x6e, 0x61, 0xc4, 0x00, 0x93, 0x43, 0x3c, 0x58, 0xc9, 0x4c, 0xc3, 0x32, 0xad, 0xc8,
	0xac, 0x4a, 0x07, 0x62, 0x06, 0x6f, 0xe4, 0xfe, 0xf9, 0x65, 0xa3, 0xdb, 0x8d, 0x8d, 0x7e, 0x6c,
	0x6c, 0xb4, 0xde, 0xd8, 0xe8, 0x6e, 0x63, 0xa3, 0x9f, 0x1b, 0x1b, 0x7d, 0xff, 0x6d, 0xf7, 0xbe,
	0xec, 0x15, 0xe1, 0x74, 0x08, 0x67, 0x37, 0xf9, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x72, 0x2a, 0x50,
	0xb5, 0x58, 0x03, 0x00, 0x00,
}

func (this *Handler) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Handler)
	if !ok {
		that2, ok := that.(Handler)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.ObjectMeta.Equal(&that1.ObjectMeta) {
		return false
	}
	if this.Type != that1.Type {
		return false
	}
	if this.Mutator != that1.Mutator {
		return false
	}
	if this.Command != that1.Command {
		return false
	}
	if this.Timeout != that1.Timeout {
		return false
	}
	if !this.Socket.Equal(that1.Socket) {
		return false
	}
	if len(this.Handlers) != len(that1.Handlers) {
		return false
	}
	for i := range this.Handlers {
		if this.Handlers[i] != that1.Handlers[i] {
			return false
		}
	}
	if len(this.Filters) != len(that1.Filters) {
		return false
	}
	for i := range this.Filters {
		if this.Filters[i] != that1.Filters[i] {
			return false
		}
	}
	if len(this.EnvVars) != len(that1.EnvVars) {
		return false
	}
	for i := range this.EnvVars {
		if this.EnvVars[i] != that1.EnvVars[i] {
			return false
		}
	}
	if len(this.RuntimeAssets) != len(that1.RuntimeAssets) {
		return false
	}
	for i := range this.RuntimeAssets {
		if this.RuntimeAssets[i] != that1.RuntimeAssets[i] {
			return false
		}
	}
	if len(this.Secrets) != len(that1.Secrets) {
		return false
	}
	for i := range this.Secrets {
		if !this.Secrets[i].Equal(that1.Secrets[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *HandlerSocket) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*HandlerSocket)
	if !ok {
		that2, ok := that.(HandlerSocket)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Host != that1.Host {
		return false
	}
	if this.Port != that1.Port {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}

type HandlerFace interface {
	Proto() github_com_golang_protobuf_proto.Message
	GetObjectMeta() ObjectMeta
	GetType() string
	GetMutator() string
	GetCommand() string
	GetTimeout() uint32
	GetSocket() *HandlerSocket
	GetHandlers() []string
	GetFilters() []string
	GetEnvVars() []string
	GetRuntimeAssets() []string
	GetSecrets() []*Secret
}

func (this *Handler) Proto() github_com_golang_protobuf_proto.Message {
	return this
}

func (this *Handler) TestProto() github_com_golang_protobuf_proto.Message {
	return NewHandlerFromFace(this)
}

func (this *Handler) GetObjectMeta() ObjectMeta {
	return this.ObjectMeta
}

func (this *Handler) GetType() string {
	return this.Type
}

func (this *Handler) GetMutator() string {
	return this.Mutator
}

func (this *Handler) GetCommand() string {
	return this.Command
}

func (this *Handler) GetTimeout() uint32 {
	return this.Timeout
}

func (this *Handler) GetSocket() *HandlerSocket {
	return this.Socket
}

func (this *Handler) GetHandlers() []string {
	return this.Handlers
}

func (this *Handler) GetFilters() []string {
	return this.Filters
}

func (this *Handler) GetEnvVars() []string {
	return this.EnvVars
}

func (this *Handler) GetRuntimeAssets() []string {
	return this.RuntimeAssets
}

func (this *Handler) GetSecrets() []*Secret {
	return this.Secrets
}

func NewHandlerFromFace(that HandlerFace) *Handler {
	this := &Handler{}
	this.ObjectMeta = that.GetObjectMeta()
	this.Type = that.GetType()
	this.Mutator = that.GetMutator()
	this.Command = that.GetCommand()
	this.Timeout = that.GetTimeout()
	this.Socket = that.GetSocket()
	this.Handlers = that.GetHandlers()
	this.Filters = that.GetFilters()
	this.EnvVars = that.GetEnvVars()
	this.RuntimeAssets = that.GetRuntimeAssets()
	this.Secrets = that.GetSecrets()
	return this
}

func (m *Handler) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Handler) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Handler) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Secrets) > 0 {
		for iNdEx := len(m.Secrets) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Secrets[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintHandler(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x72
		}
	}
	if len(m.RuntimeAssets) > 0 {
		for iNdEx := len(m.RuntimeAssets) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.RuntimeAssets[iNdEx])
			copy(dAtA[i:], m.RuntimeAssets[iNdEx])
			i = encodeVarintHandler(dAtA, i, uint64(len(m.RuntimeAssets[iNdEx])))
			i--
			dAtA[i] = 0x6a
		}
	}
	if len(m.EnvVars) > 0 {
		for iNdEx := len(m.EnvVars) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.EnvVars[iNdEx])
			copy(dAtA[i:], m.EnvVars[iNdEx])
			i = encodeVarintHandler(dAtA, i, uint64(len(m.EnvVars[iNdEx])))
			i--
			dAtA[i] = 0x4a
		}
	}
	if len(m.Filters) > 0 {
		for iNdEx := len(m.Filters) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Filters[iNdEx])
			copy(dAtA[i:], m.Filters[iNdEx])
			i = encodeVarintHandler(dAtA, i, uint64(len(m.Filters[iNdEx])))
			i--
			dAtA[i] = 0x42
		}
	}
	if len(m.Handlers) > 0 {
		for iNdEx := len(m.Handlers) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Handlers[iNdEx])
			copy(dAtA[i:], m.Handlers[iNdEx])
			i = encodeVarintHandler(dAtA, i, uint64(len(m.Handlers[iNdEx])))
			i--
			dAtA[i] = 0x3a
		}
	}
	if m.Socket != nil {
		{
			size, err := m.Socket.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintHandler(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x32
	}
	if m.Timeout != 0 {
		i = encodeVarintHandler(dAtA, i, uint64(m.Timeout))
		i--
		dAtA[i] = 0x28
	}
	if len(m.Command) > 0 {
		i -= len(m.Command)
		copy(dAtA[i:], m.Command)
		i = encodeVarintHandler(dAtA, i, uint64(len(m.Command)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Mutator) > 0 {
		i -= len(m.Mutator)
		copy(dAtA[i:], m.Mutator)
		i = encodeVarintHandler(dAtA, i, uint64(len(m.Mutator)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Type) > 0 {
		i -= len(m.Type)
		copy(dAtA[i:], m.Type)
		i = encodeVarintHandler(dAtA, i, uint64(len(m.Type)))
		i--
		dAtA[i] = 0x12
	}
	{
		size, err := m.ObjectMeta.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintHandler(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *HandlerSocket) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HandlerSocket) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HandlerSocket) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Port != 0 {
		i = encodeVarintHandler(dAtA, i, uint64(m.Port))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Host) > 0 {
		i -= len(m.Host)
		copy(dAtA[i:], m.Host)
		i = encodeVarintHandler(dAtA, i, uint64(len(m.Host)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintHandler(dAtA []byte, offset int, v uint64) int {
	offset -= sovHandler(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func NewPopulatedHandler(r randyHandler, easy bool) *Handler {
	this := &Handler{}
	v1 := NewPopulatedObjectMeta(r, easy)
	this.ObjectMeta = *v1
	this.Type = string(randStringHandler(r))
	this.Mutator = string(randStringHandler(r))
	this.Command = string(randStringHandler(r))
	this.Timeout = uint32(r.Uint32())
	if r.Intn(5) != 0 {
		this.Socket = NewPopulatedHandlerSocket(r, easy)
	}
	v2 := r.Intn(10)
	this.Handlers = make([]string, v2)
	for i := 0; i < v2; i++ {
		this.Handlers[i] = string(randStringHandler(r))
	}
	v3 := r.Intn(10)
	this.Filters = make([]string, v3)
	for i := 0; i < v3; i++ {
		this.Filters[i] = string(randStringHandler(r))
	}
	v4 := r.Intn(10)
	this.EnvVars = make([]string, v4)
	for i := 0; i < v4; i++ {
		this.EnvVars[i] = string(randStringHandler(r))
	}
	v5 := r.Intn(10)
	this.RuntimeAssets = make([]string, v5)
	for i := 0; i < v5; i++ {
		this.RuntimeAssets[i] = string(randStringHandler(r))
	}
	if r.Intn(5) != 0 {
		v6 := r.Intn(5)
		this.Secrets = make([]*Secret, v6)
		for i := 0; i < v6; i++ {
			this.Secrets[i] = NewPopulatedSecret(r, easy)
		}
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedHandler(r, 15)
	}
	return this
}

func NewPopulatedHandlerSocket(r randyHandler, easy bool) *HandlerSocket {
	this := &HandlerSocket{}
	this.Host = string(randStringHandler(r))
	this.Port = uint32(r.Uint32())
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedHandler(r, 3)
	}
	return this
}

type randyHandler interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneHandler(r randyHandler) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringHandler(r randyHandler) string {
	v7 := r.Intn(100)
	tmps := make([]rune, v7)
	for i := 0; i < v7; i++ {
		tmps[i] = randUTF8RuneHandler(r)
	}
	return string(tmps)
}
func randUnrecognizedHandler(r randyHandler, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldHandler(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldHandler(dAtA []byte, r randyHandler, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateHandler(dAtA, uint64(key))
		v8 := r.Int63()
		if r.Intn(2) == 0 {
			v8 *= -1
		}
		dAtA = encodeVarintPopulateHandler(dAtA, uint64(v8))
	case 1:
		dAtA = encodeVarintPopulateHandler(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateHandler(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateHandler(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateHandler(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateHandler(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *Handler) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ObjectMeta.Size()
	n += 1 + l + sovHandler(uint64(l))
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovHandler(uint64(l))
	}
	l = len(m.Mutator)
	if l > 0 {
		n += 1 + l + sovHandler(uint64(l))
	}
	l = len(m.Command)
	if l > 0 {
		n += 1 + l + sovHandler(uint64(l))
	}
	if m.Timeout != 0 {
		n += 1 + sovHandler(uint64(m.Timeout))
	}
	if m.Socket != nil {
		l = m.Socket.Size()
		n += 1 + l + sovHandler(uint64(l))
	}
	if len(m.Handlers) > 0 {
		for _, s := range m.Handlers {
			l = len(s)
			n += 1 + l + sovHandler(uint64(l))
		}
	}
	if len(m.Filters) > 0 {
		for _, s := range m.Filters {
			l = len(s)
			n += 1 + l + sovHandler(uint64(l))
		}
	}
	if len(m.EnvVars) > 0 {
		for _, s := range m.EnvVars {
			l = len(s)
			n += 1 + l + sovHandler(uint64(l))
		}
	}
	if len(m.RuntimeAssets) > 0 {
		for _, s := range m.RuntimeAssets {
			l = len(s)
			n += 1 + l + sovHandler(uint64(l))
		}
	}
	if len(m.Secrets) > 0 {
		for _, e := range m.Secrets {
			l = e.Size()
			n += 1 + l + sovHandler(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *HandlerSocket) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Host)
	if l > 0 {
		n += 1 + l + sovHandler(uint64(l))
	}
	if m.Port != 0 {
		n += 1 + sovHandler(uint64(m.Port))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovHandler(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozHandler(x uint64) (n int) {
	return sovHandler(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Handler) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHandler
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Handler: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Handler: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ObjectMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHandler
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthHandler
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthHandler
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ObjectMeta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHandler
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthHandler
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHandler
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Mutator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHandler
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthHandler
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHandler
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Mutator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Command", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHandler
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthHandler
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHandler
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Command = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timeout", wireType)
			}
			m.Timeout = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHandler
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Timeout |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Socket", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHandler
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthHandler
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthHandler
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Socket == nil {
				m.Socket = &HandlerSocket{}
			}
			if err := m.Socket.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Handlers", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHandler
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthHandler
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHandler
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Handlers = append(m.Handlers, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Filters", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHandler
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthHandler
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHandler
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Filters = append(m.Filters, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EnvVars", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHandler
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthHandler
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHandler
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EnvVars = append(m.EnvVars, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RuntimeAssets", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHandler
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthHandler
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHandler
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RuntimeAssets = append(m.RuntimeAssets, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 14:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Secrets", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHandler
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthHandler
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthHandler
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Secrets = append(m.Secrets, &Secret{})
			if err := m.Secrets[len(m.Secrets)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHandler(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthHandler
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *HandlerSocket) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHandler
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: HandlerSocket: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HandlerSocket: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Host", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHandler
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthHandler
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHandler
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Host = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Port", wireType)
			}
			m.Port = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHandler
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Port |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipHandler(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthHandler
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipHandler(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowHandler
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowHandler
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowHandler
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthHandler
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupHandler
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthHandler
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthHandler        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowHandler          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupHandler = fmt.Errorf("proto: unexpected end of group")
)
