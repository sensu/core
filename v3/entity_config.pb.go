// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/sensu/core/v3/entity_config.proto

package v3

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	v2 "github.com/sensu/core/v2"
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

// EntityConfig represents entity configuration.
type EntityConfig struct {
	// Metadata contains the name, namespace, labels and annotations of the
	// entity.
	Metadata *v2.ObjectMeta `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata"`
	// EntityClass represents the class of the entity. It can be "agent",
	// "proxy", or "backend".
	EntityClass string `protobuf:"bytes,2,opt,name=entity_class,json=entityClass,proto3" json:"entity_class"`
	// User is the username the entity is connecting as, if the entity is an
	// agent entity.
	User string `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
	// Subscriptions are a weak relationship between entities and checks. The
	// scheduler uses subscriptions to make entities to checks when scheduling.
	Subscriptions []string `protobuf:"bytes,4,rep,name=subscriptions,proto3" json:"subscriptions"`
	// Deregister, if true, will result in the entity being deleted when the
	// entity is an agent, and the agent disconnects its session.
	Deregister bool `protobuf:"varint,5,opt,name=deregister,proto3" json:"deregister"`
	// Deregistration contains configuration for Sensu entity de-registration.
	Deregistration v2.Deregistration `protobuf:"bytes,6,opt,name=deregistration,proto3" json:"deregistration"`
	// KeepaliveHandlers contains a list of handlers to use for the entity's
	// keepalive events.
	KeepaliveHandlers []string `protobuf:"bytes,7,rep,name=keepalive_handlers,json=keepaliveHandlers,proto3" json:"keepalive_handlers,omitempty"`
	// Redact contains the fields to redact on the entity, if the entity is an
	// agent entity.
	Redact []string `protobuf:"bytes,8,rep,name=redact,proto3" json:"redact,omitempty"`
	// Keepalive contains configuration for Sensu entity keepalives.
	Keepalive            EntityKeepalive `protobuf:"bytes,9,opt,name=keepalive,proto3" json:"keepalive"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *EntityConfig) Reset()         { *m = EntityConfig{} }
func (m *EntityConfig) String() string { return proto.CompactTextString(m) }
func (*EntityConfig) ProtoMessage()    {}
func (*EntityConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_80e9ddc662fdc76f, []int{0}
}
func (m *EntityConfig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EntityConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EntityConfig.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EntityConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EntityConfig.Merge(m, src)
}
func (m *EntityConfig) XXX_Size() int {
	return m.Size()
}
func (m *EntityConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_EntityConfig.DiscardUnknown(m)
}

var xxx_messageInfo_EntityConfig proto.InternalMessageInfo

func (m *EntityConfig) GetMetadata() *v2.ObjectMeta {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *EntityConfig) GetEntityClass() string {
	if m != nil {
		return m.EntityClass
	}
	return ""
}

func (m *EntityConfig) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *EntityConfig) GetSubscriptions() []string {
	if m != nil {
		return m.Subscriptions
	}
	return nil
}

func (m *EntityConfig) GetDeregister() bool {
	if m != nil {
		return m.Deregister
	}
	return false
}

func (m *EntityConfig) GetDeregistration() v2.Deregistration {
	if m != nil {
		return m.Deregistration
	}
	return v2.Deregistration{}
}

func (m *EntityConfig) GetKeepaliveHandlers() []string {
	if m != nil {
		return m.KeepaliveHandlers
	}
	return nil
}

func (m *EntityConfig) GetRedact() []string {
	if m != nil {
		return m.Redact
	}
	return nil
}

func (m *EntityConfig) GetKeepalive() EntityKeepalive {
	if m != nil {
		return m.Keepalive
	}
	return EntityKeepalive{}
}

// EntityKeepalive contains configuration for Sensu entity keepalives.
//sensu:nogen
type EntityKeepalive struct {
	Pipelines            []*v2.ResourceReference `protobuf:"bytes,1,rep,name=pipelines,proto3" json:"pipelines"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *EntityKeepalive) Reset()         { *m = EntityKeepalive{} }
func (m *EntityKeepalive) String() string { return proto.CompactTextString(m) }
func (*EntityKeepalive) ProtoMessage()    {}
func (*EntityKeepalive) Descriptor() ([]byte, []int) {
	return fileDescriptor_80e9ddc662fdc76f, []int{1}
}
func (m *EntityKeepalive) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EntityKeepalive) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EntityKeepalive.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EntityKeepalive) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EntityKeepalive.Merge(m, src)
}
func (m *EntityKeepalive) XXX_Size() int {
	return m.Size()
}
func (m *EntityKeepalive) XXX_DiscardUnknown() {
	xxx_messageInfo_EntityKeepalive.DiscardUnknown(m)
}

var xxx_messageInfo_EntityKeepalive proto.InternalMessageInfo

func (m *EntityKeepalive) GetPipelines() []*v2.ResourceReference {
	if m != nil {
		return m.Pipelines
	}
	return nil
}

func init() {
	proto.RegisterType((*EntityConfig)(nil), "sensu.core.v3.EntityConfig")
	proto.RegisterType((*EntityKeepalive)(nil), "sensu.core.v3.EntityKeepalive")
}

func init() {
	proto.RegisterFile("github.com/sensu/core/v3/entity_config.proto", fileDescriptor_80e9ddc662fdc76f)
}

var fileDescriptor_80e9ddc662fdc76f = []byte{
	// 484 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xcf, 0x6e, 0xd3, 0x4e,
	0x10, 0xc7, 0xbb, 0x4d, 0x7e, 0xf9, 0x25, 0x9b, 0xa6, 0xb4, 0x7b, 0x40, 0x4b, 0x25, 0x6c, 0xab,
	0x08, 0x29, 0x07, 0x58, 0xab, 0x31, 0x12, 0x57, 0xe4, 0x82, 0x40, 0xaa, 0x2a, 0xa4, 0x3d, 0x72,
	0x09, 0x8e, 0x33, 0x71, 0x0d, 0xa9, 0x6d, 0xed, 0xae, 0x2d, 0xf1, 0x26, 0x3c, 0x02, 0x8f, 0xc0,
	0x23, 0xe4, 0xc8, 0x13, 0x58, 0xd4, 0xdc, 0xfc, 0x04, 0x1c, 0x91, 0xd7, 0xce, 0x1f, 0x47, 0xca,
	0x6d, 0xfe, 0x7c, 0xbe, 0xb3, 0x33, 0xb3, 0x83, 0x5f, 0x04, 0xa1, 0xba, 0x4b, 0x67, 0xcc, 0x8f,
	0xef, 0x6d, 0x09, 0x91, 0x4c, 0x6d, 0x3f, 0x16, 0x60, 0x67, 0x8e, 0x0d, 0x91, 0x0a, 0xd5, 0xb7,
	0xa9, 0x1f, 0x47, 0x8b, 0x30, 0x60, 0x89, 0x88, 0x55, 0x4c, 0x46, 0x1a, 0x61, 0x15, 0xc2, 0x32,
	0xe7, 0xe2, 0xd5, 0x8e, 0x38, 0x88, 0x83, 0xd8, 0xd6, 0xd4, 0x2c, 0x5d, 0xbc, 0xc9, 0xae, 0x98,
	0xc3, 0x26, 0x3a, 0xa8, 0x63, 0xda, 0xaa, 0x8b, 0x5c, 0x3c, 0x3b, 0xf0, 0xe4, 0xc4, 0xbe, 0x07,
	0xe5, 0x35, 0xd0, 0xf3, 0x83, 0x50, 0xdd, 0x57, 0x83, 0x5d, 0x1d, 0xc4, 0x04, 0xc8, 0x38, 0x15,
	0x3e, 0x4c, 0x05, 0x2c, 0x40, 0x40, 0xe4, 0x43, 0x2d, 0xb9, 0x7c, 0xe8, 0xe0, 0x93, 0x77, 0xba,
	0xc6, 0xb5, 0x1e, 0x8d, 0xbc, 0xc7, 0xfd, 0xea, 0xe1, 0xb9, 0xa7, 0x3c, 0x8a, 0x2c, 0x34, 0x1e,
	0x4e, 0x9e, 0xb0, 0xdd, 0x39, 0x27, 0xec, 0xe3, 0xec, 0x0b, 0xf8, 0xea, 0x16, 0x94, 0xe7, 0x9e,
	0xad, 0x72, 0x13, 0x95, 0xb9, 0xb9, 0x91, 0xf0, 0x8d, 0x45, 0x1c, 0x7c, 0xb2, 0x5e, 0xda, 0xd2,
	0x93, 0x92, 0x1e, 0x5b, 0x68, 0x3c, 0x70, 0xcf, 0xca, 0xdc, 0x6c, 0xc5, 0xf9, 0xb0, 0xf6, 0xae,
	0x2b, 0x87, 0x10, 0xdc, 0x4d, 0x25, 0x08, 0xda, 0xa9, 0x60, 0xae, 0x6d, 0xf2, 0x1a, 0x8f, 0x64,
	0x3a, 0x93, 0xbe, 0x08, 0x13, 0x15, 0xc6, 0x91, 0xa4, 0x5d, 0xab, 0x33, 0x1e, 0xb8, 0xe7, 0x65,
	0x6e, 0xb6, 0x13, 0xbc, 0xed, 0x12, 0x86, 0xf1, 0x1c, 0x04, 0x04, 0xa1, 0x54, 0x20, 0xe8, 0x7f,
	0x16, 0x1a, 0xf7, 0xdd, 0xd3, 0x32, 0x37, 0x77, 0xa2, 0x7c, 0xc7, 0x26, 0x37, 0xf8, 0x74, 0xed,
	0x09, 0xaf, 0x2a, 0x41, 0x7b, 0x7a, 0x01, 0x4f, 0xf7, 0x16, 0xf0, 0xb6, 0x05, 0xb9, 0xdd, 0x55,
	0x6e, 0x1e, 0xf1, 0x3d, 0x29, 0x79, 0x89, 0xc9, 0x57, 0x80, 0xc4, 0x5b, 0x86, 0x19, 0x4c, 0xef,
	0xbc, 0x68, 0xbe, 0x04, 0x21, 0xe9, 0xff, 0x55, 0xeb, 0xfc, 0x7c, 0x93, 0xf9, 0xd0, 0x24, 0xc8,
	0x63, 0xdc, 0x13, 0x30, 0xf7, 0x7c, 0x45, 0xfb, 0x1a, 0x69, 0x3c, 0xe2, 0xe2, 0xc1, 0x06, 0xa6,
	0x03, 0xdd, 0x8e, 0xd1, 0x6a, 0xc7, 0x61, 0xf5, 0xf7, 0xdd, 0xac, 0xa9, 0xa6, 0x9f, 0xad, 0xec,
	0xf2, 0x33, 0x7e, 0xb4, 0xc7, 0x90, 0x5b, 0x3c, 0x48, 0xc2, 0x04, 0x96, 0x61, 0x04, 0x92, 0x22,
	0xab, 0x33, 0x1e, 0x4e, 0xac, 0xbd, 0x29, 0x79, 0x73, 0x32, 0x7c, 0x7d, 0x31, 0xee, 0xa8, 0xcc,
	0xcd, 0xad, 0x8c, 0x6f, 0x4d, 0xd7, 0xfa, 0xfb, 0x60, 0xa0, 0x1f, 0x85, 0x81, 0x7e, 0x16, 0x06,
	0x5a, 0x15, 0x06, 0xfa, 0x55, 0x18, 0xe8, 0x77, 0x61, 0xa0, 0xef, 0x7f, 0x8c, 0xa3, 0x4f, 0xc7,
	0x99, 0x33, 0xeb, 0xe9, 0x73, 0x73, 0xfe, 0x05, 0x00, 0x00, 0xff, 0xff, 0xaa, 0xf6, 0xf9, 0xa0,
	0x62, 0x03, 0x00, 0x00,
}

func (this *EntityConfig) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*EntityConfig)
	if !ok {
		that2, ok := that.(EntityConfig)
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
	if !this.Metadata.Equal(that1.Metadata) {
		return false
	}
	if this.EntityClass != that1.EntityClass {
		return false
	}
	if this.User != that1.User {
		return false
	}
	if len(this.Subscriptions) != len(that1.Subscriptions) {
		return false
	}
	for i := range this.Subscriptions {
		if this.Subscriptions[i] != that1.Subscriptions[i] {
			return false
		}
	}
	if this.Deregister != that1.Deregister {
		return false
	}
	if !this.Deregistration.Equal(&that1.Deregistration) {
		return false
	}
	if len(this.KeepaliveHandlers) != len(that1.KeepaliveHandlers) {
		return false
	}
	for i := range this.KeepaliveHandlers {
		if this.KeepaliveHandlers[i] != that1.KeepaliveHandlers[i] {
			return false
		}
	}
	if len(this.Redact) != len(that1.Redact) {
		return false
	}
	for i := range this.Redact {
		if this.Redact[i] != that1.Redact[i] {
			return false
		}
	}
	if !this.Keepalive.Equal(&that1.Keepalive) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *EntityKeepalive) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*EntityKeepalive)
	if !ok {
		that2, ok := that.(EntityKeepalive)
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
	if len(this.Pipelines) != len(that1.Pipelines) {
		return false
	}
	for i := range this.Pipelines {
		if !this.Pipelines[i].Equal(that1.Pipelines[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (m *EntityConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EntityConfig) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EntityConfig) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	{
		size, err := m.Keepalive.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintEntityConfig(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x4a
	if len(m.Redact) > 0 {
		for iNdEx := len(m.Redact) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Redact[iNdEx])
			copy(dAtA[i:], m.Redact[iNdEx])
			i = encodeVarintEntityConfig(dAtA, i, uint64(len(m.Redact[iNdEx])))
			i--
			dAtA[i] = 0x42
		}
	}
	if len(m.KeepaliveHandlers) > 0 {
		for iNdEx := len(m.KeepaliveHandlers) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.KeepaliveHandlers[iNdEx])
			copy(dAtA[i:], m.KeepaliveHandlers[iNdEx])
			i = encodeVarintEntityConfig(dAtA, i, uint64(len(m.KeepaliveHandlers[iNdEx])))
			i--
			dAtA[i] = 0x3a
		}
	}
	{
		size, err := m.Deregistration.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintEntityConfig(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	if m.Deregister {
		i--
		if m.Deregister {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x28
	}
	if len(m.Subscriptions) > 0 {
		for iNdEx := len(m.Subscriptions) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Subscriptions[iNdEx])
			copy(dAtA[i:], m.Subscriptions[iNdEx])
			i = encodeVarintEntityConfig(dAtA, i, uint64(len(m.Subscriptions[iNdEx])))
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.User) > 0 {
		i -= len(m.User)
		copy(dAtA[i:], m.User)
		i = encodeVarintEntityConfig(dAtA, i, uint64(len(m.User)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.EntityClass) > 0 {
		i -= len(m.EntityClass)
		copy(dAtA[i:], m.EntityClass)
		i = encodeVarintEntityConfig(dAtA, i, uint64(len(m.EntityClass)))
		i--
		dAtA[i] = 0x12
	}
	if m.Metadata != nil {
		{
			size, err := m.Metadata.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintEntityConfig(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EntityKeepalive) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EntityKeepalive) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EntityKeepalive) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Pipelines) > 0 {
		for iNdEx := len(m.Pipelines) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Pipelines[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintEntityConfig(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintEntityConfig(dAtA []byte, offset int, v uint64) int {
	offset -= sovEntityConfig(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func NewPopulatedEntityConfig(r randyEntityConfig, easy bool) *EntityConfig {
	this := &EntityConfig{}
	if r.Intn(5) != 0 {
		this.Metadata = v2.NewPopulatedObjectMeta(r, easy)
	}
	this.EntityClass = string(randStringEntityConfig(r))
	this.User = string(randStringEntityConfig(r))
	v1 := r.Intn(10)
	this.Subscriptions = make([]string, v1)
	for i := 0; i < v1; i++ {
		this.Subscriptions[i] = string(randStringEntityConfig(r))
	}
	this.Deregister = bool(bool(r.Intn(2) == 0))
	v2 := v2.NewPopulatedDeregistration(r, easy)
	this.Deregistration = *v2
	v3 := r.Intn(10)
	this.KeepaliveHandlers = make([]string, v3)
	for i := 0; i < v3; i++ {
		this.KeepaliveHandlers[i] = string(randStringEntityConfig(r))
	}
	v4 := r.Intn(10)
	this.Redact = make([]string, v4)
	for i := 0; i < v4; i++ {
		this.Redact[i] = string(randStringEntityConfig(r))
	}
	v5 := NewPopulatedEntityKeepalive(r, easy)
	this.Keepalive = *v5
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedEntityConfig(r, 10)
	}
	return this
}

func NewPopulatedEntityKeepalive(r randyEntityConfig, easy bool) *EntityKeepalive {
	this := &EntityKeepalive{}
	if r.Intn(5) != 0 {
		v6 := r.Intn(5)
		this.Pipelines = make([]*v2.ResourceReference, v6)
		for i := 0; i < v6; i++ {
			this.Pipelines[i] = v2.NewPopulatedResourceReference(r, easy)
		}
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedEntityConfig(r, 2)
	}
	return this
}

type randyEntityConfig interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneEntityConfig(r randyEntityConfig) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringEntityConfig(r randyEntityConfig) string {
	v7 := r.Intn(100)
	tmps := make([]rune, v7)
	for i := 0; i < v7; i++ {
		tmps[i] = randUTF8RuneEntityConfig(r)
	}
	return string(tmps)
}
func randUnrecognizedEntityConfig(r randyEntityConfig, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldEntityConfig(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldEntityConfig(dAtA []byte, r randyEntityConfig, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateEntityConfig(dAtA, uint64(key))
		v8 := r.Int63()
		if r.Intn(2) == 0 {
			v8 *= -1
		}
		dAtA = encodeVarintPopulateEntityConfig(dAtA, uint64(v8))
	case 1:
		dAtA = encodeVarintPopulateEntityConfig(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateEntityConfig(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateEntityConfig(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateEntityConfig(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateEntityConfig(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *EntityConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Metadata != nil {
		l = m.Metadata.Size()
		n += 1 + l + sovEntityConfig(uint64(l))
	}
	l = len(m.EntityClass)
	if l > 0 {
		n += 1 + l + sovEntityConfig(uint64(l))
	}
	l = len(m.User)
	if l > 0 {
		n += 1 + l + sovEntityConfig(uint64(l))
	}
	if len(m.Subscriptions) > 0 {
		for _, s := range m.Subscriptions {
			l = len(s)
			n += 1 + l + sovEntityConfig(uint64(l))
		}
	}
	if m.Deregister {
		n += 2
	}
	l = m.Deregistration.Size()
	n += 1 + l + sovEntityConfig(uint64(l))
	if len(m.KeepaliveHandlers) > 0 {
		for _, s := range m.KeepaliveHandlers {
			l = len(s)
			n += 1 + l + sovEntityConfig(uint64(l))
		}
	}
	if len(m.Redact) > 0 {
		for _, s := range m.Redact {
			l = len(s)
			n += 1 + l + sovEntityConfig(uint64(l))
		}
	}
	l = m.Keepalive.Size()
	n += 1 + l + sovEntityConfig(uint64(l))
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *EntityKeepalive) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Pipelines) > 0 {
		for _, e := range m.Pipelines {
			l = e.Size()
			n += 1 + l + sovEntityConfig(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovEntityConfig(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEntityConfig(x uint64) (n int) {
	return sovEntityConfig(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EntityConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEntityConfig
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
			return fmt.Errorf("proto: EntityConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EntityConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEntityConfig
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
				return ErrInvalidLengthEntityConfig
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEntityConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Metadata == nil {
				m.Metadata = &v2.ObjectMeta{}
			}
			if err := m.Metadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EntityClass", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEntityConfig
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
				return ErrInvalidLengthEntityConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEntityConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EntityClass = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field User", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEntityConfig
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
				return ErrInvalidLengthEntityConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEntityConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.User = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Subscriptions", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEntityConfig
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
				return ErrInvalidLengthEntityConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEntityConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Subscriptions = append(m.Subscriptions, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Deregister", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEntityConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Deregister = bool(v != 0)
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Deregistration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEntityConfig
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
				return ErrInvalidLengthEntityConfig
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEntityConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Deregistration.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field KeepaliveHandlers", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEntityConfig
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
				return ErrInvalidLengthEntityConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEntityConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.KeepaliveHandlers = append(m.KeepaliveHandlers, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Redact", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEntityConfig
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
				return ErrInvalidLengthEntityConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEntityConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Redact = append(m.Redact, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Keepalive", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEntityConfig
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
				return ErrInvalidLengthEntityConfig
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEntityConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Keepalive.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEntityConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEntityConfig
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
func (m *EntityKeepalive) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEntityConfig
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
			return fmt.Errorf("proto: EntityKeepalive: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EntityKeepalive: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pipelines", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEntityConfig
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
				return ErrInvalidLengthEntityConfig
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEntityConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Pipelines = append(m.Pipelines, &v2.ResourceReference{})
			if err := m.Pipelines[len(m.Pipelines)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEntityConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEntityConfig
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
func skipEntityConfig(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEntityConfig
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
					return 0, ErrIntOverflowEntityConfig
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
					return 0, ErrIntOverflowEntityConfig
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
				return 0, ErrInvalidLengthEntityConfig
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEntityConfig
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEntityConfig
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEntityConfig        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEntityConfig          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEntityConfig = fmt.Errorf("proto: unexpected end of group")
)
