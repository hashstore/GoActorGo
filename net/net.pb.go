// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: net.proto

package net

import (
	proto "github.com/golang/protobuf/proto"
	base "github.com/hashstore/hashlogic/base"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type IP struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Octets []byte `protobuf:"bytes,1,opt,name=octets,proto3" json:"octets,omitempty"`
}

func (x *IP) Reset() {
	*x = IP{}
	if protoimpl.UnsafeEnabled {
		mi := &file_net_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IP) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IP) ProtoMessage() {}

func (x *IP) ProtoReflect() protoreflect.Message {
	mi := &file_net_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IP.ProtoReflect.Descriptor instead.
func (*IP) Descriptor() ([]byte, []int) {
	return file_net_proto_rawDescGZIP(), []int{0}
}

func (x *IP) GetOctets() []byte {
	if x != nil {
		return x.Octets
	}
	return nil
}

// TCP port and hostname hostname
type HostCoordinate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Port int32  `protobuf:"varint,1,opt,name=port,proto3" json:"port,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Ips  []*IP  `protobuf:"bytes,3,rep,name=ips,proto3" json:"ips,omitempty"`
}

func (x *HostCoordinate) Reset() {
	*x = HostCoordinate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_net_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HostCoordinate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HostCoordinate) ProtoMessage() {}

func (x *HostCoordinate) ProtoReflect() protoreflect.Message {
	mi := &file_net_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HostCoordinate.ProtoReflect.Descriptor instead.
func (*HostCoordinate) Descriptor() ([]byte, []int) {
	return file_net_proto_rawDescGZIP(), []int{1}
}

func (x *HostCoordinate) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *HostCoordinate) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *HostCoordinate) GetIps() []*IP {
	if x != nil {
		return x.Ips
	}
	return nil
}

// public information about host
type Host struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HostId  *base.Cake        `protobuf:"bytes,1,opt,name=host_id,json=hostId,proto3" json:"host_id,omitempty"` // Cake of host public key
	PubKey  []byte            `protobuf:"bytes,2,opt,name=pub_key,json=pubKey,proto3" json:"pub_key,omitempty"`
	EnvTags []string          `protobuf:"bytes,3,rep,name=env_tags,json=envTags,proto3" json:"env_tags,omitempty"`
	Coords  []*HostCoordinate `protobuf:"bytes,4,rep,name=coords,proto3" json:"coords,omitempty"`
}

func (x *Host) Reset() {
	*x = Host{}
	if protoimpl.UnsafeEnabled {
		mi := &file_net_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Host) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Host) ProtoMessage() {}

func (x *Host) ProtoReflect() protoreflect.Message {
	mi := &file_net_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Host.ProtoReflect.Descriptor instead.
func (*Host) Descriptor() ([]byte, []int) {
	return file_net_proto_rawDescGZIP(), []int{2}
}

func (x *Host) GetHostId() *base.Cake {
	if x != nil {
		return x.HostId
	}
	return nil
}

func (x *Host) GetPubKey() []byte {
	if x != nil {
		return x.PubKey
	}
	return nil
}

func (x *Host) GetEnvTags() []string {
	if x != nil {
		return x.EnvTags
	}
	return nil
}

func (x *Host) GetCoords() []*HostCoordinate {
	if x != nil {
		return x.Coords
	}
	return nil
}

// public information about actor
type ActorLocator struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ActorId  *base.Rake `protobuf:"bytes,1,opt,name=actor_id,json=actorId,proto3" json:"actor_id,omitempty"`
	HostId   *base.Rake `protobuf:"bytes,2,opt,name=host_id,json=hostId,proto3" json:"host_id,omitempty"`
	ConfigId *base.Cake `protobuf:"bytes,3,opt,name=config_id,json=configId,proto3" json:"config_id,omitempty"`
}

func (x *ActorLocator) Reset() {
	*x = ActorLocator{}
	if protoimpl.UnsafeEnabled {
		mi := &file_net_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActorLocator) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActorLocator) ProtoMessage() {}

func (x *ActorLocator) ProtoReflect() protoreflect.Message {
	mi := &file_net_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActorLocator.ProtoReflect.Descriptor instead.
func (*ActorLocator) Descriptor() ([]byte, []int) {
	return file_net_proto_rawDescGZIP(), []int{3}
}

func (x *ActorLocator) GetActorId() *base.Rake {
	if x != nil {
		return x.ActorId
	}
	return nil
}

func (x *ActorLocator) GetHostId() *base.Rake {
	if x != nil {
		return x.HostId
	}
	return nil
}

func (x *ActorLocator) GetConfigId() *base.Cake {
	if x != nil {
		return x.ConfigId
	}
	return nil
}

// message set between actors
type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ActorTo   *base.Rake     `protobuf:"bytes,1,opt,name=actor_to,json=actorTo,proto3" json:"actor_to,omitempty"`
	ActorFrom *base.Rake     `protobuf:"bytes,2,opt,name=actor_from,json=actorFrom,proto3" json:"actor_from,omitempty"`
	Sent      *base.Nanotime `protobuf:"bytes,3,opt,name=sent,proto3" json:"sent,omitempty"`
	Body      []byte         `protobuf:"bytes,4,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_net_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_net_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_net_proto_rawDescGZIP(), []int{4}
}

func (x *Message) GetActorTo() *base.Rake {
	if x != nil {
		return x.ActorTo
	}
	return nil
}

func (x *Message) GetActorFrom() *base.Rake {
	if x != nil {
		return x.ActorFrom
	}
	return nil
}

func (x *Message) GetSent() *base.Nanotime {
	if x != nil {
		return x.Sent
	}
	return nil
}

func (x *Message) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

type ActorConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Match      *base.TagMatch `protobuf:"bytes,1,opt,name=match,proto3" json:"match,omitempty"`
	ConfigBody []byte         `protobuf:"bytes,2,opt,name=config_body,json=configBody,proto3" json:"config_body,omitempty"`
}

func (x *ActorConfig) Reset() {
	*x = ActorConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_net_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActorConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActorConfig) ProtoMessage() {}

func (x *ActorConfig) ProtoReflect() protoreflect.Message {
	mi := &file_net_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActorConfig.ProtoReflect.Descriptor instead.
func (*ActorConfig) Descriptor() ([]byte, []int) {
	return file_net_proto_rawDescGZIP(), []int{5}
}

func (x *ActorConfig) GetMatch() *base.TagMatch {
	if x != nil {
		return x.Match
	}
	return nil
}

func (x *ActorConfig) GetConfigBody() []byte {
	if x != nil {
		return x.ConfigBody
	}
	return nil
}

type Worker struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Locator *ActorLocator `protobuf:"bytes,1,opt,name=locator,proto3" json:"locator,omitempty"`
}

func (x *Worker) Reset() {
	*x = Worker{}
	if protoimpl.UnsafeEnabled {
		mi := &file_net_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Worker) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Worker) ProtoMessage() {}

func (x *Worker) ProtoReflect() protoreflect.Message {
	mi := &file_net_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Worker.ProtoReflect.Descriptor instead.
func (*Worker) Descriptor() ([]byte, []int) {
	return file_net_proto_rawDescGZIP(), []int{6}
}

func (x *Worker) GetLocator() *ActorLocator {
	if x != nil {
		return x.Locator
	}
	return nil
}

type Domain struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DomainId *base.Rake      `protobuf:"bytes,1,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	Hosts    []*Host         `protobuf:"bytes,2,rep,name=hosts,proto3" json:"hosts,omitempty"`
	Actors   []*ActorLocator `protobuf:"bytes,3,rep,name=actors,proto3" json:"actors,omitempty"`
	Configs  []*ActorConfig  `protobuf:"bytes,4,rep,name=configs,proto3" json:"configs,omitempty"`
}

func (x *Domain) Reset() {
	*x = Domain{}
	if protoimpl.UnsafeEnabled {
		mi := &file_net_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Domain) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Domain) ProtoMessage() {}

func (x *Domain) ProtoReflect() protoreflect.Message {
	mi := &file_net_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Domain.ProtoReflect.Descriptor instead.
func (*Domain) Descriptor() ([]byte, []int) {
	return file_net_proto_rawDescGZIP(), []int{7}
}

func (x *Domain) GetDomainId() *base.Rake {
	if x != nil {
		return x.DomainId
	}
	return nil
}

func (x *Domain) GetHosts() []*Host {
	if x != nil {
		return x.Hosts
	}
	return nil
}

func (x *Domain) GetActors() []*ActorLocator {
	if x != nil {
		return x.Actors
	}
	return nil
}

func (x *Domain) GetConfigs() []*ActorConfig {
	if x != nil {
		return x.Configs
	}
	return nil
}

// private host state
type HostSuperviser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Host        *Host     `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	PrivateKey  []byte    `protobuf:"bytes,2,opt,name=private_key,json=privateKey,proto3" json:"private_key,omitempty"`
	Workers     []*Worker `protobuf:"bytes,3,rep,name=workers,proto3" json:"workers,omitempty"`
	AddressBook []*Domain `protobuf:"bytes,4,rep,name=addressBook,proto3" json:"addressBook,omitempty"`
}

func (x *HostSuperviser) Reset() {
	*x = HostSuperviser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_net_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HostSuperviser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HostSuperviser) ProtoMessage() {}

func (x *HostSuperviser) ProtoReflect() protoreflect.Message {
	mi := &file_net_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HostSuperviser.ProtoReflect.Descriptor instead.
func (*HostSuperviser) Descriptor() ([]byte, []int) {
	return file_net_proto_rawDescGZIP(), []int{8}
}

func (x *HostSuperviser) GetHost() *Host {
	if x != nil {
		return x.Host
	}
	return nil
}

func (x *HostSuperviser) GetPrivateKey() []byte {
	if x != nil {
		return x.PrivateKey
	}
	return nil
}

func (x *HostSuperviser) GetWorkers() []*Worker {
	if x != nil {
		return x.Workers
	}
	return nil
}

func (x *HostSuperviser) GetAddressBook() []*Domain {
	if x != nil {
		return x.AddressBook
	}
	return nil
}

var File_net_proto protoreflect.FileDescriptor

var file_net_proto_rawDesc = []byte{
	0x0a, 0x09, 0x6e, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x6e, 0x65, 0x74,
	0x1a, 0x0a, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1c, 0x0a, 0x02,
	0x49, 0x50, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x63, 0x74, 0x65, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x06, 0x6f, 0x63, 0x74, 0x65, 0x74, 0x73, 0x22, 0x53, 0x0a, 0x0e, 0x48, 0x6f,
	0x73, 0x74, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x70, 0x6f, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x03, 0x69, 0x70, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x07, 0x2e, 0x6e, 0x65, 0x74, 0x2e, 0x49, 0x50, 0x52, 0x03, 0x69, 0x70, 0x73, 0x22,
	0x8c, 0x01, 0x0a, 0x04, 0x48, 0x6f, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x07, 0x68, 0x6f, 0x73, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x62, 0x61, 0x73, 0x65,
	0x2e, 0x43, 0x61, 0x6b, 0x65, 0x52, 0x06, 0x68, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x12, 0x17, 0x0a,
	0x07, 0x70, 0x75, 0x62, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06,
	0x70, 0x75, 0x62, 0x4b, 0x65, 0x79, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x76, 0x5f, 0x74, 0x61,
	0x67, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x76, 0x54, 0x61, 0x67,
	0x73, 0x12, 0x2b, 0x0a, 0x06, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x6e, 0x65, 0x74, 0x2e, 0x48, 0x6f, 0x73, 0x74, 0x43, 0x6f, 0x6f, 0x72,
	0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x52, 0x06, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x73, 0x22, 0x83,
	0x01, 0x0a, 0x0c, 0x41, 0x63, 0x74, 0x6f, 0x72, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x12,
	0x25, 0x0a, 0x08, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0a, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x52, 0x61, 0x6b, 0x65, 0x52, 0x07, 0x61,
	0x63, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x07, 0x68, 0x6f, 0x73, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x52,
	0x61, 0x6b, 0x65, 0x52, 0x06, 0x68, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x09, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a,
	0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x43, 0x61, 0x6b, 0x65, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x49, 0x64, 0x22, 0x93, 0x01, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x25, 0x0a, 0x08, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x74, 0x6f, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x52, 0x61, 0x6b, 0x65, 0x52, 0x07,
	0x61, 0x63, 0x74, 0x6f, 0x72, 0x54, 0x6f, 0x12, 0x29, 0x0a, 0x0a, 0x61, 0x63, 0x74, 0x6f, 0x72,
	0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x62, 0x61,
	0x73, 0x65, 0x2e, 0x52, 0x61, 0x6b, 0x65, 0x52, 0x09, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x46, 0x72,
	0x6f, 0x6d, 0x12, 0x22, 0x0a, 0x04, 0x73, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0e, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x4e, 0x61, 0x6e, 0x6f, 0x74, 0x69, 0x6d, 0x65,
	0x52, 0x04, 0x73, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x54, 0x0a, 0x0b, 0x41, 0x63,
	0x74, 0x6f, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x24, 0x0a, 0x05, 0x6d, 0x61, 0x74,
	0x63, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e,
	0x54, 0x61, 0x67, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x05, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x12,
	0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x6f, 0x64, 0x79,
	0x22, 0x35, 0x0a, 0x06, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x12, 0x2b, 0x0a, 0x07, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6e, 0x65,
	0x74, 0x2e, 0x41, 0x63, 0x74, 0x6f, 0x72, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x07,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x22, 0xa9, 0x01, 0x0a, 0x06, 0x44, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x12, 0x27, 0x0a, 0x09, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x52, 0x61, 0x6b,
	0x65, 0x52, 0x08, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x05, 0x68,
	0x6f, 0x73, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x6e, 0x65, 0x74,
	0x2e, 0x48, 0x6f, 0x73, 0x74, 0x52, 0x05, 0x68, 0x6f, 0x73, 0x74, 0x73, 0x12, 0x29, 0x0a, 0x06,
	0x61, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6e,
	0x65, 0x74, 0x2e, 0x41, 0x63, 0x74, 0x6f, 0x72, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x52,
	0x06, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x12, 0x2a, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6e, 0x65, 0x74, 0x2e, 0x41,
	0x63, 0x74, 0x6f, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x73, 0x22, 0xa6, 0x01, 0x0a, 0x0e, 0x48, 0x6f, 0x73, 0x74, 0x53, 0x75, 0x70, 0x65,
	0x72, 0x76, 0x69, 0x73, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x6e, 0x65, 0x74, 0x2e, 0x48, 0x6f, 0x73, 0x74, 0x52,
	0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65,
	0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x70, 0x72, 0x69, 0x76,
	0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x25, 0x0a, 0x07, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x6e, 0x65, 0x74, 0x2e, 0x57, 0x6f,
	0x72, 0x6b, 0x65, 0x72, 0x52, 0x07, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x73, 0x12, 0x2d, 0x0a,
	0x0b, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x42, 0x6f, 0x6f, 0x6b, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x6e, 0x65, 0x74, 0x2e, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x52,
	0x0b, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x42, 0x6f, 0x6f, 0x6b, 0x42, 0x28, 0x5a, 0x26,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x61, 0x73, 0x68, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x2f, 0x68, 0x61, 0x73, 0x68, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x2f, 0x6e,
	0x65, 0x74, 0x3b, 0x6e, 0x65, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_net_proto_rawDescOnce sync.Once
	file_net_proto_rawDescData = file_net_proto_rawDesc
)

func file_net_proto_rawDescGZIP() []byte {
	file_net_proto_rawDescOnce.Do(func() {
		file_net_proto_rawDescData = protoimpl.X.CompressGZIP(file_net_proto_rawDescData)
	})
	return file_net_proto_rawDescData
}

var file_net_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_net_proto_goTypes = []interface{}{
	(*IP)(nil),             // 0: net.IP
	(*HostCoordinate)(nil), // 1: net.HostCoordinate
	(*Host)(nil),           // 2: net.Host
	(*ActorLocator)(nil),   // 3: net.ActorLocator
	(*Message)(nil),        // 4: net.Message
	(*ActorConfig)(nil),    // 5: net.ActorConfig
	(*Worker)(nil),         // 6: net.Worker
	(*Domain)(nil),         // 7: net.Domain
	(*HostSuperviser)(nil), // 8: net.HostSuperviser
	(*base.Cake)(nil),      // 9: base.Cake
	(*base.Rake)(nil),      // 10: base.Rake
	(*base.Nanotime)(nil),  // 11: base.Nanotime
	(*base.TagMatch)(nil),  // 12: base.TagMatch
}
var file_net_proto_depIdxs = []int32{
	0,  // 0: net.HostCoordinate.ips:type_name -> net.IP
	9,  // 1: net.Host.host_id:type_name -> base.Cake
	1,  // 2: net.Host.coords:type_name -> net.HostCoordinate
	10, // 3: net.ActorLocator.actor_id:type_name -> base.Rake
	10, // 4: net.ActorLocator.host_id:type_name -> base.Rake
	9,  // 5: net.ActorLocator.config_id:type_name -> base.Cake
	10, // 6: net.Message.actor_to:type_name -> base.Rake
	10, // 7: net.Message.actor_from:type_name -> base.Rake
	11, // 8: net.Message.sent:type_name -> base.Nanotime
	12, // 9: net.ActorConfig.match:type_name -> base.TagMatch
	3,  // 10: net.Worker.locator:type_name -> net.ActorLocator
	10, // 11: net.Domain.domain_id:type_name -> base.Rake
	2,  // 12: net.Domain.hosts:type_name -> net.Host
	3,  // 13: net.Domain.actors:type_name -> net.ActorLocator
	5,  // 14: net.Domain.configs:type_name -> net.ActorConfig
	2,  // 15: net.HostSuperviser.host:type_name -> net.Host
	6,  // 16: net.HostSuperviser.workers:type_name -> net.Worker
	7,  // 17: net.HostSuperviser.addressBook:type_name -> net.Domain
	18, // [18:18] is the sub-list for method output_type
	18, // [18:18] is the sub-list for method input_type
	18, // [18:18] is the sub-list for extension type_name
	18, // [18:18] is the sub-list for extension extendee
	0,  // [0:18] is the sub-list for field type_name
}

func init() { file_net_proto_init() }
func file_net_proto_init() {
	if File_net_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_net_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IP); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_net_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HostCoordinate); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_net_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Host); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_net_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActorLocator); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_net_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_net_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActorConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_net_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Worker); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_net_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Domain); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_net_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HostSuperviser); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_net_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_net_proto_goTypes,
		DependencyIndexes: file_net_proto_depIdxs,
		MessageInfos:      file_net_proto_msgTypes,
	}.Build()
	File_net_proto = out.File
	file_net_proto_rawDesc = nil
	file_net_proto_goTypes = nil
	file_net_proto_depIdxs = nil
}
