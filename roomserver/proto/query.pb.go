// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.12.4
// source: query.proto

package proto

import (
	proto "github.com/golang/protobuf/proto"
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

type ServerBannedFromRoomRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerName string `protobuf:"bytes,1,opt,name=ServerName,proto3" json:"ServerName,omitempty"`
	RoomID     string `protobuf:"bytes,2,opt,name=RoomID,proto3" json:"RoomID,omitempty"`
}

func (x *ServerBannedFromRoomRequest) Reset() {
	*x = ServerBannedFromRoomRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_query_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerBannedFromRoomRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerBannedFromRoomRequest) ProtoMessage() {}

func (x *ServerBannedFromRoomRequest) ProtoReflect() protoreflect.Message {
	mi := &file_query_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerBannedFromRoomRequest.ProtoReflect.Descriptor instead.
func (*ServerBannedFromRoomRequest) Descriptor() ([]byte, []int) {
	return file_query_proto_rawDescGZIP(), []int{0}
}

func (x *ServerBannedFromRoomRequest) GetServerName() string {
	if x != nil {
		return x.ServerName
	}
	return ""
}

func (x *ServerBannedFromRoomRequest) GetRoomID() string {
	if x != nil {
		return x.RoomID
	}
	return ""
}

type ServerBannedFromRoomResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Banned bool `protobuf:"varint,1,opt,name=Banned,proto3" json:"Banned,omitempty"`
}

func (x *ServerBannedFromRoomResponse) Reset() {
	*x = ServerBannedFromRoomResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_query_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerBannedFromRoomResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerBannedFromRoomResponse) ProtoMessage() {}

func (x *ServerBannedFromRoomResponse) ProtoReflect() protoreflect.Message {
	mi := &file_query_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerBannedFromRoomResponse.ProtoReflect.Descriptor instead.
func (*ServerBannedFromRoomResponse) Descriptor() ([]byte, []int) {
	return file_query_proto_rawDescGZIP(), []int{1}
}

func (x *ServerBannedFromRoomResponse) GetBanned() bool {
	if x != nil {
		return x.Banned
	}
	return false
}

type SharedUsersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID         string   `protobuf:"bytes,1,opt,name=UserID,json=user_id,omitempty,proto3" json:"UserID,omitempty"`
	ExcludeRoomIDs []string `protobuf:"bytes,2,rep,name=ExcludeRoomIDs,proto3" json:"ExcludeRoomIDs,omitempty"`
	IncludeRoomIDs []string `protobuf:"bytes,3,rep,name=IncludeRoomIDs,proto3" json:"IncludeRoomIDs,omitempty"`
}

func (x *SharedUsersRequest) Reset() {
	*x = SharedUsersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_query_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SharedUsersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SharedUsersRequest) ProtoMessage() {}

func (x *SharedUsersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_query_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SharedUsersRequest.ProtoReflect.Descriptor instead.
func (*SharedUsersRequest) Descriptor() ([]byte, []int) {
	return file_query_proto_rawDescGZIP(), []int{2}
}

func (x *SharedUsersRequest) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *SharedUsersRequest) GetExcludeRoomIDs() []string {
	if x != nil {
		return x.ExcludeRoomIDs
	}
	return nil
}

func (x *SharedUsersRequest) GetIncludeRoomIDs() []string {
	if x != nil {
		return x.IncludeRoomIDs
	}
	return nil
}

type SharedUsersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserIDsToCount map[string]int64 `protobuf:"bytes,1,rep,name=UserIDsToCount,proto3" json:"UserIDsToCount,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *SharedUsersResponse) Reset() {
	*x = SharedUsersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_query_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SharedUsersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SharedUsersResponse) ProtoMessage() {}

func (x *SharedUsersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_query_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SharedUsersResponse.ProtoReflect.Descriptor instead.
func (*SharedUsersResponse) Descriptor() ([]byte, []int) {
	return file_query_proto_rawDescGZIP(), []int{3}
}

func (x *SharedUsersResponse) GetUserIDsToCount() map[string]int64 {
	if x != nil {
		return x.UserIDsToCount
	}
	return nil
}

type RoomsForUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID         string `protobuf:"bytes,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	WantMembership string `protobuf:"bytes,2,opt,name=WantMembership,proto3" json:"WantMembership,omitempty"`
}

func (x *RoomsForUserRequest) Reset() {
	*x = RoomsForUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_query_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomsForUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomsForUserRequest) ProtoMessage() {}

func (x *RoomsForUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_query_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomsForUserRequest.ProtoReflect.Descriptor instead.
func (*RoomsForUserRequest) Descriptor() ([]byte, []int) {
	return file_query_proto_rawDescGZIP(), []int{4}
}

func (x *RoomsForUserRequest) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *RoomsForUserRequest) GetWantMembership() string {
	if x != nil {
		return x.WantMembership
	}
	return ""
}

type RoomsForUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomIDs []string `protobuf:"bytes,1,rep,name=RoomIDs,proto3" json:"RoomIDs,omitempty"`
}

func (x *RoomsForUserResponse) Reset() {
	*x = RoomsForUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_query_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomsForUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomsForUserResponse) ProtoMessage() {}

func (x *RoomsForUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_query_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomsForUserResponse.ProtoReflect.Descriptor instead.
func (*RoomsForUserResponse) Descriptor() ([]byte, []int) {
	return file_query_proto_rawDescGZIP(), []int{5}
}

func (x *RoomsForUserResponse) GetRoomIDs() []string {
	if x != nil {
		return x.RoomIDs
	}
	return nil
}

type PublishedRoomsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PublishedRoomsRequest) Reset() {
	*x = PublishedRoomsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_query_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishedRoomsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishedRoomsRequest) ProtoMessage() {}

func (x *PublishedRoomsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_query_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishedRoomsRequest.ProtoReflect.Descriptor instead.
func (*PublishedRoomsRequest) Descriptor() ([]byte, []int) {
	return file_query_proto_rawDescGZIP(), []int{6}
}

type PublishedRoomsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomIDs []string `protobuf:"bytes,1,rep,name=RoomIDs,proto3" json:"RoomIDs,omitempty"`
}

func (x *PublishedRoomsResponse) Reset() {
	*x = PublishedRoomsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_query_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishedRoomsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishedRoomsResponse) ProtoMessage() {}

func (x *PublishedRoomsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_query_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishedRoomsResponse.ProtoReflect.Descriptor instead.
func (*PublishedRoomsResponse) Descriptor() ([]byte, []int) {
	return file_query_proto_rawDescGZIP(), []int{7}
}

func (x *PublishedRoomsResponse) GetRoomIDs() []string {
	if x != nil {
		return x.RoomIDs
	}
	return nil
}

type RoomVersionForRoomRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomID string `protobuf:"bytes,1,opt,name=RoomID,proto3" json:"RoomID,omitempty"`
}

func (x *RoomVersionForRoomRequest) Reset() {
	*x = RoomVersionForRoomRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_query_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomVersionForRoomRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomVersionForRoomRequest) ProtoMessage() {}

func (x *RoomVersionForRoomRequest) ProtoReflect() protoreflect.Message {
	mi := &file_query_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomVersionForRoomRequest.ProtoReflect.Descriptor instead.
func (*RoomVersionForRoomRequest) Descriptor() ([]byte, []int) {
	return file_query_proto_rawDescGZIP(), []int{8}
}

func (x *RoomVersionForRoomRequest) GetRoomID() string {
	if x != nil {
		return x.RoomID
	}
	return ""
}

type RoomVersionForRoomResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomVersion RoomVersion `protobuf:"varint,1,opt,name=RoomVersion,proto3,enum=common.RoomVersion" json:"RoomVersion,omitempty"`
}

func (x *RoomVersionForRoomResponse) Reset() {
	*x = RoomVersionForRoomResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_query_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomVersionForRoomResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomVersionForRoomResponse) ProtoMessage() {}

func (x *RoomVersionForRoomResponse) ProtoReflect() protoreflect.Message {
	mi := &file_query_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomVersionForRoomResponse.ProtoReflect.Descriptor instead.
func (*RoomVersionForRoomResponse) Descriptor() ([]byte, []int) {
	return file_query_proto_rawDescGZIP(), []int{9}
}

func (x *RoomVersionForRoomResponse) GetRoomVersion() RoomVersion {
	if x != nil {
		return x.RoomVersion
	}
	return RoomVersion_UNKNOWN
}

type RoomVersionCapabilitiesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RoomVersionCapabilitiesRequest) Reset() {
	*x = RoomVersionCapabilitiesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_query_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomVersionCapabilitiesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomVersionCapabilitiesRequest) ProtoMessage() {}

func (x *RoomVersionCapabilitiesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_query_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomVersionCapabilitiesRequest.ProtoReflect.Descriptor instead.
func (*RoomVersionCapabilitiesRequest) Descriptor() ([]byte, []int) {
	return file_query_proto_rawDescGZIP(), []int{10}
}

type RoomVersionCapabilitiesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DefaultRoomVersion    RoomVersion       `protobuf:"varint,1,opt,name=DefaultRoomVersion,proto3,enum=common.RoomVersion" json:"DefaultRoomVersion,omitempty"`
	AvailableRoomVersions map[string]string `protobuf:"bytes,2,rep,name=AvailableRoomVersions,proto3" json:"AvailableRoomVersions,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // Must be manually handled!
}

func (x *RoomVersionCapabilitiesResponse) Reset() {
	*x = RoomVersionCapabilitiesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_query_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomVersionCapabilitiesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomVersionCapabilitiesResponse) ProtoMessage() {}

func (x *RoomVersionCapabilitiesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_query_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomVersionCapabilitiesResponse.ProtoReflect.Descriptor instead.
func (*RoomVersionCapabilitiesResponse) Descriptor() ([]byte, []int) {
	return file_query_proto_rawDescGZIP(), []int{11}
}

func (x *RoomVersionCapabilitiesResponse) GetDefaultRoomVersion() RoomVersion {
	if x != nil {
		return x.DefaultRoomVersion
	}
	return RoomVersion_UNKNOWN
}

func (x *RoomVersionCapabilitiesResponse) GetAvailableRoomVersions() map[string]string {
	if x != nil {
		return x.AvailableRoomVersions
	}
	return nil
}

var File_query_proto protoreflect.FileDescriptor

var file_query_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x1a, 0x11, 0x72, 0x6f, 0x6f, 0x6d, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x55, 0x0a, 0x1b, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x64, 0x46, 0x72, 0x6f, 0x6d, 0x52, 0x6f, 0x6f, 0x6d, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x6f, 0x6f, 0x6d, 0x49, 0x44,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x52, 0x6f, 0x6f, 0x6d, 0x49, 0x44, 0x22, 0x36,
	0x0a, 0x1c, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x64, 0x46, 0x72,
	0x6f, 0x6d, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06,
	0x42, 0x61, 0x6e, 0x6e, 0x65, 0x64, 0x22, 0x87, 0x01, 0x0a, 0x12, 0x53, 0x68, 0x61, 0x72, 0x65,
	0x64, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a,
	0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x2c, 0x6f, 0x6d, 0x69, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x12, 0x26, 0x0a, 0x0e, 0x45, 0x78, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x49,
	0x44, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0e, 0x45, 0x78, 0x63, 0x6c, 0x75, 0x64,
	0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x49, 0x44, 0x73, 0x12, 0x26, 0x0a, 0x0e, 0x49, 0x6e, 0x63, 0x6c,
	0x75, 0x64, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x49, 0x44, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x0e, 0x49, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x49, 0x44, 0x73,
	0x22, 0xb0, 0x01, 0x0a, 0x13, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x55, 0x73, 0x65, 0x72, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x56, 0x0a, 0x0e, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x44, 0x73, 0x54, 0x6f, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x2e, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x55,
	0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x44, 0x73, 0x54, 0x6f, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x0e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x73, 0x54, 0x6f, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x1a, 0x41, 0x0a, 0x13, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x73, 0x54, 0x6f, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x22, 0x55, 0x0a, 0x13, 0x52, 0x6f, 0x6f, 0x6d, 0x73, 0x46, 0x6f, 0x72, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x44, 0x12, 0x26, 0x0a, 0x0e, 0x57, 0x61, 0x6e, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x73, 0x68, 0x69, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x57, 0x61, 0x6e, 0x74,
	0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x22, 0x30, 0x0a, 0x14, 0x52, 0x6f,
	0x6f, 0x6d, 0x73, 0x46, 0x6f, 0x72, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x52, 0x6f, 0x6f, 0x6d, 0x49, 0x44, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x07, 0x52, 0x6f, 0x6f, 0x6d, 0x49, 0x44, 0x73, 0x22, 0x17, 0x0a, 0x15,
	0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x52, 0x6f, 0x6f, 0x6d, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x32, 0x0a, 0x16, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68,
	0x65, 0x64, 0x52, 0x6f, 0x6f, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x52, 0x6f, 0x6f, 0x6d, 0x49, 0x44, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x07, 0x52, 0x6f, 0x6f, 0x6d, 0x49, 0x44, 0x73, 0x22, 0x33, 0x0a, 0x19, 0x52, 0x6f, 0x6f,
	0x6d, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x46, 0x6f, 0x72, 0x52, 0x6f, 0x6f, 0x6d, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x6f, 0x6f, 0x6d, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x52, 0x6f, 0x6f, 0x6d, 0x49, 0x44, 0x22, 0x53,
	0x0a, 0x1a, 0x52, 0x6f, 0x6f, 0x6d, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x46, 0x6f, 0x72,
	0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x0b,
	0x52, 0x6f, 0x6f, 0x6d, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x13, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x52, 0x6f, 0x6f, 0x6d, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x22, 0x20, 0x0a, 0x1e, 0x52, 0x6f, 0x6f, 0x6d, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xa9, 0x02, 0x0a, 0x1f, 0x52, 0x6f, 0x6f, 0x6d, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x12, 0x44, 0x65, 0x66,
	0x61, 0x75, 0x6c, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52,
	0x6f, 0x6f, 0x6d, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x12, 0x44, 0x65, 0x66, 0x61,
	0x75, 0x6c, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x77,
	0x0a, 0x15, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x41, 0x2e,
	0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x52,
	0x6f, 0x6f, 0x6d, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x15, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x48, 0x0a, 0x1a, 0x41, 0x76, 0x61, 0x69, 0x6c,
	0x61, 0x62, 0x6c, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_query_proto_rawDescOnce sync.Once
	file_query_proto_rawDescData = file_query_proto_rawDesc
)

func file_query_proto_rawDescGZIP() []byte {
	file_query_proto_rawDescOnce.Do(func() {
		file_query_proto_rawDescData = protoimpl.X.CompressGZIP(file_query_proto_rawDescData)
	})
	return file_query_proto_rawDescData
}

var file_query_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_query_proto_goTypes = []interface{}{
	(*ServerBannedFromRoomRequest)(nil),     // 0: query.ServerBannedFromRoomRequest
	(*ServerBannedFromRoomResponse)(nil),    // 1: query.ServerBannedFromRoomResponse
	(*SharedUsersRequest)(nil),              // 2: query.SharedUsersRequest
	(*SharedUsersResponse)(nil),             // 3: query.SharedUsersResponse
	(*RoomsForUserRequest)(nil),             // 4: query.RoomsForUserRequest
	(*RoomsForUserResponse)(nil),            // 5: query.RoomsForUserResponse
	(*PublishedRoomsRequest)(nil),           // 6: query.PublishedRoomsRequest
	(*PublishedRoomsResponse)(nil),          // 7: query.PublishedRoomsResponse
	(*RoomVersionForRoomRequest)(nil),       // 8: query.RoomVersionForRoomRequest
	(*RoomVersionForRoomResponse)(nil),      // 9: query.RoomVersionForRoomResponse
	(*RoomVersionCapabilitiesRequest)(nil),  // 10: query.RoomVersionCapabilitiesRequest
	(*RoomVersionCapabilitiesResponse)(nil), // 11: query.RoomVersionCapabilitiesResponse
	nil,                                     // 12: query.SharedUsersResponse.UserIDsToCountEntry
	nil,                                     // 13: query.RoomVersionCapabilitiesResponse.AvailableRoomVersionsEntry
	(RoomVersion)(0),                        // 14: common.RoomVersion
}
var file_query_proto_depIdxs = []int32{
	12, // 0: query.SharedUsersResponse.UserIDsToCount:type_name -> query.SharedUsersResponse.UserIDsToCountEntry
	14, // 1: query.RoomVersionForRoomResponse.RoomVersion:type_name -> common.RoomVersion
	14, // 2: query.RoomVersionCapabilitiesResponse.DefaultRoomVersion:type_name -> common.RoomVersion
	13, // 3: query.RoomVersionCapabilitiesResponse.AvailableRoomVersions:type_name -> query.RoomVersionCapabilitiesResponse.AvailableRoomVersionsEntry
	4,  // [4:4] is the sub-list for method output_type
	4,  // [4:4] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_query_proto_init() }
func file_query_proto_init() {
	if File_query_proto != nil {
		return
	}
	file_roomversion_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_query_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerBannedFromRoomRequest); i {
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
		file_query_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerBannedFromRoomResponse); i {
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
		file_query_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SharedUsersRequest); i {
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
		file_query_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SharedUsersResponse); i {
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
		file_query_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoomsForUserRequest); i {
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
		file_query_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoomsForUserResponse); i {
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
		file_query_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublishedRoomsRequest); i {
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
		file_query_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublishedRoomsResponse); i {
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
		file_query_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoomVersionForRoomRequest); i {
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
		file_query_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoomVersionForRoomResponse); i {
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
		file_query_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoomVersionCapabilitiesRequest); i {
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
		file_query_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoomVersionCapabilitiesResponse); i {
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
			RawDescriptor: file_query_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_query_proto_goTypes,
		DependencyIndexes: file_query_proto_depIdxs,
		MessageInfos:      file_query_proto_msgTypes,
	}.Build()
	File_query_proto = out.File
	file_query_proto_rawDesc = nil
	file_query_proto_goTypes = nil
	file_query_proto_depIdxs = nil
}
