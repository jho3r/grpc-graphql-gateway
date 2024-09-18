// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.11.4
// source: starwars/starwars.proto

package starwars

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/jho3r/grpc-graphql-gateway/graphql"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Type int32

const (
	Type_HUMAN Type = 0
	Type_DROID Type = 1
)

// Enum value maps for Type.
var (
	Type_name = map[int32]string{
		0: "HUMAN",
		1: "DROID",
	}
	Type_value = map[string]int32{
		"HUMAN": 0,
		"DROID": 1,
	}
)

func (x Type) Enum() *Type {
	p := new(Type)
	*p = x
	return p
}

func (x Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Type) Descriptor() protoreflect.EnumDescriptor {
	return file_starwars_starwars_proto_enumTypes[0].Descriptor()
}

func (Type) Type() protoreflect.EnumType {
	return &file_starwars_starwars_proto_enumTypes[0]
}

func (x Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Type.Descriptor instead.
func (Type) EnumDescriptor() ([]byte, []int) {
	return file_starwars_starwars_proto_rawDescGZIP(), []int{0}
}

type Episode int32

const (
	Episode__       Episode = 0
	Episode_NEWHOPE Episode = 1
	Episode_EMPIRE  Episode = 2
	Episode_JEDI    Episode = 3
)

// Enum value maps for Episode.
var (
	Episode_name = map[int32]string{
		0: "_",
		1: "NEWHOPE",
		2: "EMPIRE",
		3: "JEDI",
	}
	Episode_value = map[string]int32{
		"_":       0,
		"NEWHOPE": 1,
		"EMPIRE":  2,
		"JEDI":    3,
	}
)

func (x Episode) Enum() *Episode {
	p := new(Episode)
	*p = x
	return p
}

func (x Episode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Episode) Descriptor() protoreflect.EnumDescriptor {
	return file_starwars_starwars_proto_enumTypes[1].Descriptor()
}

func (Episode) Type() protoreflect.EnumType {
	return &file_starwars_starwars_proto_enumTypes[1]
}

func (x Episode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Episode.Descriptor instead.
func (Episode) EnumDescriptor() ([]byte, []int) {
	return file_starwars_starwars_proto_rawDescGZIP(), []int{1}
}

type Character struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              int64        `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name            string       `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Friends         []*Character `protobuf:"bytes,3,rep,name=friends,proto3" json:"friends,omitempty"`
	AppearsIn       []Episode    `protobuf:"varint,4,rep,packed,name=appears_in,json=appearsIn,proto3,enum=startwars.Episode" json:"appears_in,omitempty"`
	HomePlanet      string       `protobuf:"bytes,5,opt,name=home_planet,json=homePlanet,proto3" json:"home_planet,omitempty"`
	PrimaryFunction string       `protobuf:"bytes,6,opt,name=primary_function,json=primaryFunction,proto3" json:"primary_function,omitempty"`
	Type            Type         `protobuf:"varint,7,opt,name=type,proto3,enum=startwars.Type" json:"type,omitempty"`
}

func (x *Character) Reset() {
	*x = Character{}
	if protoimpl.UnsafeEnabled {
		mi := &file_starwars_starwars_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Character) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Character) ProtoMessage() {}

func (x *Character) ProtoReflect() protoreflect.Message {
	mi := &file_starwars_starwars_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Character.ProtoReflect.Descriptor instead.
func (*Character) Descriptor() ([]byte, []int) {
	return file_starwars_starwars_proto_rawDescGZIP(), []int{0}
}

func (x *Character) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Character) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Character) GetFriends() []*Character {
	if x != nil {
		return x.Friends
	}
	return nil
}

func (x *Character) GetAppearsIn() []Episode {
	if x != nil {
		return x.AppearsIn
	}
	return nil
}

func (x *Character) GetHomePlanet() string {
	if x != nil {
		return x.HomePlanet
	}
	return ""
}

func (x *Character) GetPrimaryFunction() string {
	if x != nil {
		return x.PrimaryFunction
	}
	return ""
}

func (x *Character) GetType() Type {
	if x != nil {
		return x.Type
	}
	return Type_HUMAN
}

type GetHeroRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// If omitted, returns the hero of the whope saga. If provided, returns the hero of that particular episode.
	Episode Episode `protobuf:"varint,1,opt,name=episode,proto3,enum=startwars.Episode" json:"episode,omitempty"`
}

func (x *GetHeroRequest) Reset() {
	*x = GetHeroRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_starwars_starwars_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHeroRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHeroRequest) ProtoMessage() {}

func (x *GetHeroRequest) ProtoReflect() protoreflect.Message {
	mi := &file_starwars_starwars_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHeroRequest.ProtoReflect.Descriptor instead.
func (*GetHeroRequest) Descriptor() ([]byte, []int) {
	return file_starwars_starwars_proto_rawDescGZIP(), []int{1}
}

func (x *GetHeroRequest) GetEpisode() Episode {
	if x != nil {
		return x.Episode
	}
	return Episode__
}

type GetHumanRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id of the human
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetHumanRequest) Reset() {
	*x = GetHumanRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_starwars_starwars_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHumanRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHumanRequest) ProtoMessage() {}

func (x *GetHumanRequest) ProtoReflect() protoreflect.Message {
	mi := &file_starwars_starwars_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHumanRequest.ProtoReflect.Descriptor instead.
func (*GetHumanRequest) Descriptor() ([]byte, []int) {
	return file_starwars_starwars_proto_rawDescGZIP(), []int{2}
}

func (x *GetHumanRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetDroidRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id of the droid
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetDroidRequest) Reset() {
	*x = GetDroidRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_starwars_starwars_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDroidRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDroidRequest) ProtoMessage() {}

func (x *GetDroidRequest) ProtoReflect() protoreflect.Message {
	mi := &file_starwars_starwars_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDroidRequest.ProtoReflect.Descriptor instead.
func (*GetDroidRequest) Descriptor() ([]byte, []int) {
	return file_starwars_starwars_proto_rawDescGZIP(), []int{3}
}

func (x *GetDroidRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ListEmptyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListEmptyRequest) Reset() {
	*x = ListEmptyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_starwars_starwars_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListEmptyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListEmptyRequest) ProtoMessage() {}

func (x *ListEmptyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_starwars_starwars_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListEmptyRequest.ProtoReflect.Descriptor instead.
func (*ListEmptyRequest) Descriptor() ([]byte, []int) {
	return file_starwars_starwars_proto_rawDescGZIP(), []int{4}
}

type ListHumansResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Humans []*Character `protobuf:"bytes,1,rep,name=humans,proto3" json:"humans,omitempty"`
}

func (x *ListHumansResponse) Reset() {
	*x = ListHumansResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_starwars_starwars_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListHumansResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListHumansResponse) ProtoMessage() {}

func (x *ListHumansResponse) ProtoReflect() protoreflect.Message {
	mi := &file_starwars_starwars_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListHumansResponse.ProtoReflect.Descriptor instead.
func (*ListHumansResponse) Descriptor() ([]byte, []int) {
	return file_starwars_starwars_proto_rawDescGZIP(), []int{5}
}

func (x *ListHumansResponse) GetHumans() []*Character {
	if x != nil {
		return x.Humans
	}
	return nil
}

type ListDroidsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Droids []*Character `protobuf:"bytes,1,rep,name=droids,proto3" json:"droids,omitempty"`
}

func (x *ListDroidsResponse) Reset() {
	*x = ListDroidsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_starwars_starwars_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListDroidsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDroidsResponse) ProtoMessage() {}

func (x *ListDroidsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_starwars_starwars_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDroidsResponse.ProtoReflect.Descriptor instead.
func (*ListDroidsResponse) Descriptor() ([]byte, []int) {
	return file_starwars_starwars_proto_rawDescGZIP(), []int{6}
}

func (x *ListDroidsResponse) GetDroids() []*Character {
	if x != nil {
		return x.Droids
	}
	return nil
}

var File_starwars_starwars_proto protoreflect.FileDescriptor

var file_starwars_starwars_proto_rawDesc = []byte{
	0x0a, 0x17, 0x73, 0x74, 0x61, 0x72, 0x77, 0x61, 0x72, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x72, 0x77,
	0x61, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x77, 0x61, 0x72, 0x73, 0x1a, 0x15, 0x67, 0x72, 0x61, 0x70, 0x68, 0x71, 0x6c, 0x2f, 0x67, 0x72,
	0x61, 0x70, 0x68, 0x71, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x83, 0x02, 0x0a, 0x09,
	0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2e, 0x0a,
	0x07, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x73, 0x74, 0x61, 0x72, 0x74, 0x77, 0x61, 0x72, 0x73, 0x2e, 0x43, 0x68, 0x61, 0x72, 0x61,
	0x63, 0x74, 0x65, 0x72, 0x52, 0x07, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x12, 0x31, 0x0a,
	0x0a, 0x61, 0x70, 0x70, 0x65, 0x61, 0x72, 0x73, 0x5f, 0x69, 0x6e, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x0e, 0x32, 0x12, 0x2e, 0x73, 0x74, 0x61, 0x72, 0x74, 0x77, 0x61, 0x72, 0x73, 0x2e, 0x45, 0x70,
	0x69, 0x73, 0x6f, 0x64, 0x65, 0x52, 0x09, 0x61, 0x70, 0x70, 0x65, 0x61, 0x72, 0x73, 0x49, 0x6e,
	0x12, 0x1f, 0x0a, 0x0b, 0x68, 0x6f, 0x6d, 0x65, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x68, 0x6f, 0x6d, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x65,
	0x74, 0x12, 0x29, 0x0a, 0x10, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x5f, 0x66, 0x75, 0x6e,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x70, 0x72, 0x69,
	0x6d, 0x61, 0x72, 0x79, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x77, 0x61, 0x72, 0x73, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x22, 0x3e, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x48, 0x65, 0x72, 0x6f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x07, 0x65, 0x70, 0x69, 0x73, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x73, 0x74, 0x61, 0x72, 0x74, 0x77, 0x61, 0x72, 0x73,
	0x2e, 0x45, 0x70, 0x69, 0x73, 0x6f, 0x64, 0x65, 0x52, 0x07, 0x65, 0x70, 0x69, 0x73, 0x6f, 0x64,
	0x65, 0x22, 0x28, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x48, 0x75, 0x6d, 0x61, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x15, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x42, 0x05, 0xba, 0x43, 0x02, 0x08, 0x01, 0x52, 0x02, 0x69, 0x64, 0x22, 0x28, 0x0a, 0x0f, 0x47,
	0x65, 0x74, 0x44, 0x72, 0x6f, 0x69, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x15,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x05, 0xba, 0x43, 0x02, 0x08,
	0x01, 0x52, 0x02, 0x69, 0x64, 0x22, 0x12, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x42, 0x0a, 0x12, 0x4c, 0x69, 0x73,
	0x74, 0x48, 0x75, 0x6d, 0x61, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x2c, 0x0a, 0x06, 0x68, 0x75, 0x6d, 0x61, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x73, 0x74, 0x61, 0x72, 0x74, 0x77, 0x61, 0x72, 0x73, 0x2e, 0x43, 0x68, 0x61, 0x72,
	0x61, 0x63, 0x74, 0x65, 0x72, 0x52, 0x06, 0x68, 0x75, 0x6d, 0x61, 0x6e, 0x73, 0x22, 0x42, 0x0a,
	0x12, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x72, 0x6f, 0x69, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x06, 0x64, 0x72, 0x6f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x74, 0x61, 0x72, 0x74, 0x77, 0x61, 0x72, 0x73, 0x2e,
	0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x52, 0x06, 0x64, 0x72, 0x6f, 0x69, 0x64,
	0x73, 0x2a, 0x1c, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x48, 0x55, 0x4d,
	0x41, 0x4e, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x44, 0x52, 0x4f, 0x49, 0x44, 0x10, 0x01, 0x2a,
	0x33, 0x0a, 0x07, 0x45, 0x70, 0x69, 0x73, 0x6f, 0x64, 0x65, 0x12, 0x05, 0x0a, 0x01, 0x5f, 0x10,
	0x00, 0x12, 0x0b, 0x0a, 0x07, 0x4e, 0x45, 0x57, 0x48, 0x4f, 0x50, 0x45, 0x10, 0x01, 0x12, 0x0a,
	0x0a, 0x06, 0x45, 0x4d, 0x50, 0x49, 0x52, 0x45, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x4a, 0x45,
	0x44, 0x49, 0x10, 0x03, 0x32, 0xc2, 0x03, 0x0a, 0x10, 0x53, 0x74, 0x61, 0x72, 0x74, 0x77, 0x61,
	0x72, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x45, 0x0a, 0x07, 0x47, 0x65, 0x74,
	0x48, 0x65, 0x72, 0x6f, 0x12, 0x19, 0x2e, 0x73, 0x74, 0x61, 0x72, 0x74, 0x77, 0x61, 0x72, 0x73,
	0x2e, 0x47, 0x65, 0x74, 0x48, 0x65, 0x72, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x14, 0x2e, 0x73, 0x74, 0x61, 0x72, 0x74, 0x77, 0x61, 0x72, 0x73, 0x2e, 0x43, 0x68, 0x61, 0x72,
	0x61, 0x63, 0x74, 0x65, 0x72, 0x22, 0x09, 0xba, 0x43, 0x06, 0x12, 0x04, 0x68, 0x65, 0x72, 0x6f,
	0x12, 0x48, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x48, 0x75, 0x6d, 0x61, 0x6e, 0x12, 0x1a, 0x2e, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x77, 0x61, 0x72, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x48, 0x75, 0x6d, 0x61,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x77, 0x61, 0x72, 0x73, 0x2e, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x22, 0x0a,
	0xba, 0x43, 0x07, 0x12, 0x05, 0x68, 0x75, 0x6d, 0x61, 0x6e, 0x12, 0x48, 0x0a, 0x08, 0x47, 0x65,
	0x74, 0x44, 0x72, 0x6f, 0x69, 0x64, 0x12, 0x1a, 0x2e, 0x73, 0x74, 0x61, 0x72, 0x74, 0x77, 0x61,
	0x72, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x72, 0x6f, 0x69, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x14, 0x2e, 0x73, 0x74, 0x61, 0x72, 0x74, 0x77, 0x61, 0x72, 0x73, 0x2e, 0x43,
	0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x22, 0x0a, 0xba, 0x43, 0x07, 0x12, 0x05, 0x64,
	0x72, 0x6f, 0x69, 0x64, 0x12, 0x5f, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x75, 0x6d, 0x61,
	0x6e, 0x73, 0x12, 0x1b, 0x2e, 0x73, 0x74, 0x61, 0x72, 0x74, 0x77, 0x61, 0x72, 0x73, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1d, 0x2e, 0x73, 0x74, 0x61, 0x72, 0x74, 0x77, 0x61, 0x72, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x48, 0x75, 0x6d, 0x61, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x15,
	0xba, 0x43, 0x12, 0x12, 0x06, 0x68, 0x75, 0x6d, 0x61, 0x6e, 0x73, 0x22, 0x08, 0x12, 0x06, 0x68,
	0x75, 0x6d, 0x61, 0x6e, 0x73, 0x12, 0x5f, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x72, 0x6f,
	0x69, 0x64, 0x73, 0x12, 0x1b, 0x2e, 0x73, 0x74, 0x61, 0x72, 0x74, 0x77, 0x61, 0x72, 0x73, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1d, 0x2e, 0x73, 0x74, 0x61, 0x72, 0x74, 0x77, 0x61, 0x72, 0x73, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x44, 0x72, 0x6f, 0x69, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x15, 0xba, 0x43, 0x12, 0x12, 0x06, 0x64, 0x72, 0x6f, 0x69, 0x64, 0x73, 0x22, 0x08, 0x12, 0x06,
	0x64, 0x72, 0x6f, 0x69, 0x64, 0x73, 0x1a, 0x11, 0xba, 0x43, 0x0e, 0x0a, 0x0a, 0x67, 0x72, 0x70,
	0x63, 0x3a, 0x35, 0x30, 0x30, 0x35, 0x31, 0x10, 0x01, 0x42, 0x4a, 0x5a, 0x48, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x73, 0x75, 0x67, 0x69, 0x6d, 0x6f, 0x74,
	0x6f, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x67, 0x72, 0x61, 0x70, 0x68, 0x71, 0x6c, 0x2d, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x73,
	0x74, 0x61, 0x72, 0x77, 0x61, 0x72, 0x73, 0x2f, 0x73, 0x70, 0x65, 0x63, 0x2f, 0x73, 0x74, 0x61,
	0x72, 0x77, 0x61, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_starwars_starwars_proto_rawDescOnce sync.Once
	file_starwars_starwars_proto_rawDescData = file_starwars_starwars_proto_rawDesc
)

func file_starwars_starwars_proto_rawDescGZIP() []byte {
	file_starwars_starwars_proto_rawDescOnce.Do(func() {
		file_starwars_starwars_proto_rawDescData = protoimpl.X.CompressGZIP(file_starwars_starwars_proto_rawDescData)
	})
	return file_starwars_starwars_proto_rawDescData
}

var file_starwars_starwars_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_starwars_starwars_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_starwars_starwars_proto_goTypes = []interface{}{
	(Type)(0),                  // 0: startwars.Type
	(Episode)(0),               // 1: startwars.Episode
	(*Character)(nil),          // 2: startwars.Character
	(*GetHeroRequest)(nil),     // 3: startwars.GetHeroRequest
	(*GetHumanRequest)(nil),    // 4: startwars.GetHumanRequest
	(*GetDroidRequest)(nil),    // 5: startwars.GetDroidRequest
	(*ListEmptyRequest)(nil),   // 6: startwars.ListEmptyRequest
	(*ListHumansResponse)(nil), // 7: startwars.ListHumansResponse
	(*ListDroidsResponse)(nil), // 8: startwars.ListDroidsResponse
}
var file_starwars_starwars_proto_depIdxs = []int32{
	2,  // 0: startwars.Character.friends:type_name -> startwars.Character
	1,  // 1: startwars.Character.appears_in:type_name -> startwars.Episode
	0,  // 2: startwars.Character.type:type_name -> startwars.Type
	1,  // 3: startwars.GetHeroRequest.episode:type_name -> startwars.Episode
	2,  // 4: startwars.ListHumansResponse.humans:type_name -> startwars.Character
	2,  // 5: startwars.ListDroidsResponse.droids:type_name -> startwars.Character
	3,  // 6: startwars.StartwarsService.GetHero:input_type -> startwars.GetHeroRequest
	4,  // 7: startwars.StartwarsService.GetHuman:input_type -> startwars.GetHumanRequest
	5,  // 8: startwars.StartwarsService.GetDroid:input_type -> startwars.GetDroidRequest
	6,  // 9: startwars.StartwarsService.ListHumans:input_type -> startwars.ListEmptyRequest
	6,  // 10: startwars.StartwarsService.ListDroids:input_type -> startwars.ListEmptyRequest
	2,  // 11: startwars.StartwarsService.GetHero:output_type -> startwars.Character
	2,  // 12: startwars.StartwarsService.GetHuman:output_type -> startwars.Character
	2,  // 13: startwars.StartwarsService.GetDroid:output_type -> startwars.Character
	7,  // 14: startwars.StartwarsService.ListHumans:output_type -> startwars.ListHumansResponse
	8,  // 15: startwars.StartwarsService.ListDroids:output_type -> startwars.ListDroidsResponse
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_starwars_starwars_proto_init() }
func file_starwars_starwars_proto_init() {
	if File_starwars_starwars_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_starwars_starwars_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Character); i {
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
		file_starwars_starwars_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetHeroRequest); i {
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
		file_starwars_starwars_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetHumanRequest); i {
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
		file_starwars_starwars_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDroidRequest); i {
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
		file_starwars_starwars_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListEmptyRequest); i {
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
		file_starwars_starwars_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListHumansResponse); i {
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
		file_starwars_starwars_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListDroidsResponse); i {
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
			RawDescriptor: file_starwars_starwars_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_starwars_starwars_proto_goTypes,
		DependencyIndexes: file_starwars_starwars_proto_depIdxs,
		EnumInfos:         file_starwars_starwars_proto_enumTypes,
		MessageInfos:      file_starwars_starwars_proto_msgTypes,
	}.Build()
	File_starwars_starwars_proto = out.File
	file_starwars_starwars_proto_rawDesc = nil
	file_starwars_starwars_proto_goTypes = nil
	file_starwars_starwars_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// StartwarsServiceClient is the client API for StartwarsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StartwarsServiceClient interface {
	GetHero(ctx context.Context, in *GetHeroRequest, opts ...grpc.CallOption) (*Character, error)
	GetHuman(ctx context.Context, in *GetHumanRequest, opts ...grpc.CallOption) (*Character, error)
	GetDroid(ctx context.Context, in *GetDroidRequest, opts ...grpc.CallOption) (*Character, error)
	ListHumans(ctx context.Context, in *ListEmptyRequest, opts ...grpc.CallOption) (*ListHumansResponse, error)
	ListDroids(ctx context.Context, in *ListEmptyRequest, opts ...grpc.CallOption) (*ListDroidsResponse, error)
}

type startwarsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStartwarsServiceClient(cc grpc.ClientConnInterface) StartwarsServiceClient {
	return &startwarsServiceClient{cc}
}

func (c *startwarsServiceClient) GetHero(ctx context.Context, in *GetHeroRequest, opts ...grpc.CallOption) (*Character, error) {
	out := new(Character)
	err := c.cc.Invoke(ctx, "/startwars.StartwarsService/GetHero", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *startwarsServiceClient) GetHuman(ctx context.Context, in *GetHumanRequest, opts ...grpc.CallOption) (*Character, error) {
	out := new(Character)
	err := c.cc.Invoke(ctx, "/startwars.StartwarsService/GetHuman", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *startwarsServiceClient) GetDroid(ctx context.Context, in *GetDroidRequest, opts ...grpc.CallOption) (*Character, error) {
	out := new(Character)
	err := c.cc.Invoke(ctx, "/startwars.StartwarsService/GetDroid", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *startwarsServiceClient) ListHumans(ctx context.Context, in *ListEmptyRequest, opts ...grpc.CallOption) (*ListHumansResponse, error) {
	out := new(ListHumansResponse)
	err := c.cc.Invoke(ctx, "/startwars.StartwarsService/ListHumans", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *startwarsServiceClient) ListDroids(ctx context.Context, in *ListEmptyRequest, opts ...grpc.CallOption) (*ListDroidsResponse, error) {
	out := new(ListDroidsResponse)
	err := c.cc.Invoke(ctx, "/startwars.StartwarsService/ListDroids", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StartwarsServiceServer is the server API for StartwarsService service.
type StartwarsServiceServer interface {
	GetHero(context.Context, *GetHeroRequest) (*Character, error)
	GetHuman(context.Context, *GetHumanRequest) (*Character, error)
	GetDroid(context.Context, *GetDroidRequest) (*Character, error)
	ListHumans(context.Context, *ListEmptyRequest) (*ListHumansResponse, error)
	ListDroids(context.Context, *ListEmptyRequest) (*ListDroidsResponse, error)
}

// UnimplementedStartwarsServiceServer can be embedded to have forward compatible implementations.
type UnimplementedStartwarsServiceServer struct {
}

func (*UnimplementedStartwarsServiceServer) GetHero(context.Context, *GetHeroRequest) (*Character, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHero not implemented")
}
func (*UnimplementedStartwarsServiceServer) GetHuman(context.Context, *GetHumanRequest) (*Character, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHuman not implemented")
}
func (*UnimplementedStartwarsServiceServer) GetDroid(context.Context, *GetDroidRequest) (*Character, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDroid not implemented")
}
func (*UnimplementedStartwarsServiceServer) ListHumans(context.Context, *ListEmptyRequest) (*ListHumansResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListHumans not implemented")
}
func (*UnimplementedStartwarsServiceServer) ListDroids(context.Context, *ListEmptyRequest) (*ListDroidsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDroids not implemented")
}

func RegisterStartwarsServiceServer(s *grpc.Server, srv StartwarsServiceServer) {
	s.RegisterService(&_StartwarsService_serviceDesc, srv)
}

func _StartwarsService_GetHero_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHeroRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StartwarsServiceServer).GetHero(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/startwars.StartwarsService/GetHero",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StartwarsServiceServer).GetHero(ctx, req.(*GetHeroRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StartwarsService_GetHuman_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHumanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StartwarsServiceServer).GetHuman(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/startwars.StartwarsService/GetHuman",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StartwarsServiceServer).GetHuman(ctx, req.(*GetHumanRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StartwarsService_GetDroid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDroidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StartwarsServiceServer).GetDroid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/startwars.StartwarsService/GetDroid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StartwarsServiceServer).GetDroid(ctx, req.(*GetDroidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StartwarsService_ListHumans_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListEmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StartwarsServiceServer).ListHumans(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/startwars.StartwarsService/ListHumans",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StartwarsServiceServer).ListHumans(ctx, req.(*ListEmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StartwarsService_ListDroids_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListEmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StartwarsServiceServer).ListDroids(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/startwars.StartwarsService/ListDroids",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StartwarsServiceServer).ListDroids(ctx, req.(*ListEmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _StartwarsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "startwars.StartwarsService",
	HandlerType: (*StartwarsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetHero",
			Handler:    _StartwarsService_GetHero_Handler,
		},
		{
			MethodName: "GetHuman",
			Handler:    _StartwarsService_GetHuman_Handler,
		},
		{
			MethodName: "GetDroid",
			Handler:    _StartwarsService_GetDroid_Handler,
		},
		{
			MethodName: "ListHumans",
			Handler:    _StartwarsService_ListHumans_Handler,
		},
		{
			MethodName: "ListDroids",
			Handler:    _StartwarsService_ListDroids_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "starwars/starwars.proto",
}
