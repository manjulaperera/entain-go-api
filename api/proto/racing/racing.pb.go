// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: racing/racing.proto

package racing

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Sort order/ direction of the given field
type OrderByField_Direction int32

const (
	// Ascending order
	OrderByField_ASC OrderByField_Direction = 0
	// Descending order
	OrderByField_DESC OrderByField_Direction = 1
)

// Enum value maps for OrderByField_Direction.
var (
	OrderByField_Direction_name = map[int32]string{
		0: "ASC",
		1: "DESC",
	}
	OrderByField_Direction_value = map[string]int32{
		"ASC":  0,
		"DESC": 1,
	}
)

func (x OrderByField_Direction) Enum() *OrderByField_Direction {
	p := new(OrderByField_Direction)
	*p = x
	return p
}

func (x OrderByField_Direction) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OrderByField_Direction) Descriptor() protoreflect.EnumDescriptor {
	return file_racing_racing_proto_enumTypes[0].Descriptor()
}

func (OrderByField_Direction) Type() protoreflect.EnumType {
	return &file_racing_racing_proto_enumTypes[0]
}

func (x OrderByField_Direction) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OrderByField_Direction.Descriptor instead.
func (OrderByField_Direction) EnumDescriptor() ([]byte, []int) {
	return file_racing_racing_proto_rawDescGZIP(), []int{5, 0}
}

// Request for ListRaces call.
type ListRacesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filter  *ListRacesRequestFilter  `protobuf:"bytes,1,opt,name=filter,proto3" json:"filter,omitempty"`
	OrderBy *ListRacesRequestOrderBy `protobuf:"bytes,2,opt,name=order_by,json=orderBy,proto3,oneof" json:"order_by,omitempty"`
}

func (x *ListRacesRequest) Reset() {
	*x = ListRacesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_racing_racing_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRacesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRacesRequest) ProtoMessage() {}

func (x *ListRacesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_racing_racing_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRacesRequest.ProtoReflect.Descriptor instead.
func (*ListRacesRequest) Descriptor() ([]byte, []int) {
	return file_racing_racing_proto_rawDescGZIP(), []int{0}
}

func (x *ListRacesRequest) GetFilter() *ListRacesRequestFilter {
	if x != nil {
		return x.Filter
	}
	return nil
}

func (x *ListRacesRequest) GetOrderBy() *ListRacesRequestOrderBy {
	if x != nil {
		return x.OrderBy
	}
	return nil
}

// Response to ListRaces call.
type ListRacesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Races []*Race `protobuf:"bytes,1,rep,name=races,proto3" json:"races,omitempty"`
}

func (x *ListRacesResponse) Reset() {
	*x = ListRacesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_racing_racing_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRacesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRacesResponse) ProtoMessage() {}

func (x *ListRacesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_racing_racing_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRacesResponse.ProtoReflect.Descriptor instead.
func (*ListRacesResponse) Descriptor() ([]byte, []int) {
	return file_racing_racing_proto_rawDescGZIP(), []int{1}
}

func (x *ListRacesResponse) GetRaces() []*Race {
	if x != nil {
		return x.Races
	}
	return nil
}

// Filter for listing races.
type ListRacesRequestFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MeetingIds []int64 `protobuf:"varint,1,rep,packed,name=meeting_ids,json=meetingIds,proto3" json:"meeting_ids,omitempty"`
	// Use this filter for filtering the race meets based on their visibility
	MeetingVisibility *bool `protobuf:"varint,2,opt,name=meeting_visibility,json=meetingVisibility,proto3,oneof" json:"meeting_visibility,omitempty"`
}

func (x *ListRacesRequestFilter) Reset() {
	*x = ListRacesRequestFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_racing_racing_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRacesRequestFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRacesRequestFilter) ProtoMessage() {}

func (x *ListRacesRequestFilter) ProtoReflect() protoreflect.Message {
	mi := &file_racing_racing_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRacesRequestFilter.ProtoReflect.Descriptor instead.
func (*ListRacesRequestFilter) Descriptor() ([]byte, []int) {
	return file_racing_racing_proto_rawDescGZIP(), []int{2}
}

func (x *ListRacesRequestFilter) GetMeetingIds() []int64 {
	if x != nil {
		return x.MeetingIds
	}
	return nil
}

func (x *ListRacesRequestFilter) GetMeetingVisibility() bool {
	if x != nil && x.MeetingVisibility != nil {
		return *x.MeetingVisibility
	}
	return false
}

// An array of order by fileds to be used to order the races list.
//The list will be ordered in the order the fields appear in this list
type ListRacesRequestOrderBy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderByFields []*OrderByField `protobuf:"bytes,1,rep,name=order_by_fields,json=orderByFields,proto3" json:"order_by_fields,omitempty"`
}

func (x *ListRacesRequestOrderBy) Reset() {
	*x = ListRacesRequestOrderBy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_racing_racing_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRacesRequestOrderBy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRacesRequestOrderBy) ProtoMessage() {}

func (x *ListRacesRequestOrderBy) ProtoReflect() protoreflect.Message {
	mi := &file_racing_racing_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRacesRequestOrderBy.ProtoReflect.Descriptor instead.
func (*ListRacesRequestOrderBy) Descriptor() ([]byte, []int) {
	return file_racing_racing_proto_rawDescGZIP(), []int{3}
}

func (x *ListRacesRequestOrderBy) GetOrderByFields() []*OrderByField {
	if x != nil {
		return x.OrderByFields
	}
	return nil
}

// A race resource.
type Race struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID represents a unique identifier for the race.
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// MeetingID represents a unique identifier for the races meeting.
	MeetingId int64 `protobuf:"varint,2,opt,name=meeting_id,json=meetingId,proto3" json:"meeting_id,omitempty"`
	// Name is the official name given to the race.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// Number represents the number of the race.
	Number int64 `protobuf:"varint,4,opt,name=number,proto3" json:"number,omitempty"`
	// Visible represents whether or not the race is visible.
	Visible bool `protobuf:"varint,5,opt,name=visible,proto3" json:"visible,omitempty"`
	// AdvertisedStartTime is the time the race is advertised to run.
	AdvertisedStartTime *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=advertised_start_time,json=advertisedStartTime,proto3" json:"advertised_start_time,omitempty"`
}

func (x *Race) Reset() {
	*x = Race{}
	if protoimpl.UnsafeEnabled {
		mi := &file_racing_racing_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Race) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Race) ProtoMessage() {}

func (x *Race) ProtoReflect() protoreflect.Message {
	mi := &file_racing_racing_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Race.ProtoReflect.Descriptor instead.
func (*Race) Descriptor() ([]byte, []int) {
	return file_racing_racing_proto_rawDescGZIP(), []int{4}
}

func (x *Race) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Race) GetMeetingId() int64 {
	if x != nil {
		return x.MeetingId
	}
	return 0
}

func (x *Race) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Race) GetNumber() int64 {
	if x != nil {
		return x.Number
	}
	return 0
}

func (x *Race) GetVisible() bool {
	if x != nil {
		return x.Visible
	}
	return false
}

func (x *Race) GetAdvertisedStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.AdvertisedStartTime
	}
	return nil
}

// A field with it's sort order to be used as a sort/ order by field in the ListRacesRequest
type OrderByField struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The field to be used for sorting/ ordering
	Field     string                 `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
	Direction OrderByField_Direction `protobuf:"varint,2,opt,name=direction,proto3,enum=racing.OrderByField_Direction" json:"direction,omitempty"`
}

func (x *OrderByField) Reset() {
	*x = OrderByField{}
	if protoimpl.UnsafeEnabled {
		mi := &file_racing_racing_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderByField) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderByField) ProtoMessage() {}

func (x *OrderByField) ProtoReflect() protoreflect.Message {
	mi := &file_racing_racing_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderByField.ProtoReflect.Descriptor instead.
func (*OrderByField) Descriptor() ([]byte, []int) {
	return file_racing_racing_proto_rawDescGZIP(), []int{5}
}

func (x *OrderByField) GetField() string {
	if x != nil {
		return x.Field
	}
	return ""
}

func (x *OrderByField) GetDirection() OrderByField_Direction {
	if x != nil {
		return x.Direction
	}
	return OrderByField_ASC
}

var File_racing_racing_proto protoreflect.FileDescriptor

var file_racing_racing_proto_rawDesc = []byte{
	0x0a, 0x13, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x2f, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x98, 0x01, 0x0a,
	0x10, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x61, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x36, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1e, 0x2e, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x61, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x3f, 0x0a, 0x08, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x5f, 0x62, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x72, 0x61,
	0x63, 0x69, 0x6e, 0x67, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x61, 0x63, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x48, 0x00, 0x52, 0x07,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x88, 0x01, 0x01, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x5f, 0x62, 0x79, 0x22, 0x37, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x61, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x05,
	0x72, 0x61, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x72, 0x61,
	0x63, 0x69, 0x6e, 0x67, 0x2e, 0x52, 0x61, 0x63, 0x65, 0x52, 0x05, 0x72, 0x61, 0x63, 0x65, 0x73,
	0x22, 0x84, 0x01, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x61, 0x63, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x6d,
	0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x03,
	0x52, 0x0a, 0x6d, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x73, 0x12, 0x32, 0x0a, 0x12,
	0x6d, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x76, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x11, 0x6d, 0x65, 0x65, 0x74,
	0x69, 0x6e, 0x67, 0x56, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x88, 0x01, 0x01,
	0x42, 0x15, 0x0a, 0x13, 0x5f, 0x6d, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x76, 0x69, 0x73,
	0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x22, 0x57, 0x0a, 0x17, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x61, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x42, 0x79, 0x12, 0x3c, 0x0a, 0x0f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x62, 0x79, 0x5f, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x72, 0x61,
	0x63, 0x69, 0x6e, 0x67, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x52, 0x0d, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73,
	0x22, 0xcb, 0x01, 0x0a, 0x04, 0x52, 0x61, 0x63, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x65, 0x65,
	0x74, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x6d,
	0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x69, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x76, 0x69, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x12, 0x4e,
	0x0a, 0x15, 0x61, 0x64, 0x76, 0x65, 0x72, 0x74, 0x69, 0x73, 0x65, 0x64, 0x5f, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x13, 0x61, 0x64, 0x76, 0x65, 0x72,
	0x74, 0x69, 0x73, 0x65, 0x64, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x82,
	0x01, 0x0a, 0x0c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x3c, 0x0a, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x72, 0x61, 0x63, 0x69, 0x6e,
	0x67, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x2e, 0x44,
	0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0x1e, 0x0a, 0x09, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x07, 0x0a, 0x03, 0x41, 0x53, 0x43, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x45, 0x53,
	0x43, 0x10, 0x01, 0x32, 0x65, 0x0a, 0x06, 0x52, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x12, 0x5b, 0x0a,
	0x09, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x61, 0x63, 0x65, 0x73, 0x12, 0x18, 0x2e, 0x72, 0x61, 0x63,
	0x69, 0x6e, 0x67, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x61, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x61, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x22, 0x0e, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x69, 0x73,
	0x74, 0x2d, 0x72, 0x61, 0x63, 0x65, 0x73, 0x3a, 0x01, 0x2a, 0x42, 0x09, 0x5a, 0x07, 0x2f, 0x72,
	0x61, 0x63, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_racing_racing_proto_rawDescOnce sync.Once
	file_racing_racing_proto_rawDescData = file_racing_racing_proto_rawDesc
)

func file_racing_racing_proto_rawDescGZIP() []byte {
	file_racing_racing_proto_rawDescOnce.Do(func() {
		file_racing_racing_proto_rawDescData = protoimpl.X.CompressGZIP(file_racing_racing_proto_rawDescData)
	})
	return file_racing_racing_proto_rawDescData
}

var file_racing_racing_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_racing_racing_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_racing_racing_proto_goTypes = []interface{}{
	(OrderByField_Direction)(0),     // 0: racing.OrderByField.Direction
	(*ListRacesRequest)(nil),        // 1: racing.ListRacesRequest
	(*ListRacesResponse)(nil),       // 2: racing.ListRacesResponse
	(*ListRacesRequestFilter)(nil),  // 3: racing.ListRacesRequestFilter
	(*ListRacesRequestOrderBy)(nil), // 4: racing.ListRacesRequestOrderBy
	(*Race)(nil),                    // 5: racing.Race
	(*OrderByField)(nil),            // 6: racing.OrderByField
	(*timestamppb.Timestamp)(nil),   // 7: google.protobuf.Timestamp
}
var file_racing_racing_proto_depIdxs = []int32{
	3, // 0: racing.ListRacesRequest.filter:type_name -> racing.ListRacesRequestFilter
	4, // 1: racing.ListRacesRequest.order_by:type_name -> racing.ListRacesRequestOrderBy
	5, // 2: racing.ListRacesResponse.races:type_name -> racing.Race
	6, // 3: racing.ListRacesRequestOrderBy.order_by_fields:type_name -> racing.OrderByField
	7, // 4: racing.Race.advertised_start_time:type_name -> google.protobuf.Timestamp
	0, // 5: racing.OrderByField.direction:type_name -> racing.OrderByField.Direction
	1, // 6: racing.Racing.ListRaces:input_type -> racing.ListRacesRequest
	2, // 7: racing.Racing.ListRaces:output_type -> racing.ListRacesResponse
	7, // [7:8] is the sub-list for method output_type
	6, // [6:7] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_racing_racing_proto_init() }
func file_racing_racing_proto_init() {
	if File_racing_racing_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_racing_racing_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRacesRequest); i {
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
		file_racing_racing_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRacesResponse); i {
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
		file_racing_racing_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRacesRequestFilter); i {
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
		file_racing_racing_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRacesRequestOrderBy); i {
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
		file_racing_racing_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Race); i {
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
		file_racing_racing_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderByField); i {
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
	file_racing_racing_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_racing_racing_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_racing_racing_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_racing_racing_proto_goTypes,
		DependencyIndexes: file_racing_racing_proto_depIdxs,
		EnumInfos:         file_racing_racing_proto_enumTypes,
		MessageInfos:      file_racing_racing_proto_msgTypes,
	}.Build()
	File_racing_racing_proto = out.File
	file_racing_racing_proto_rawDesc = nil
	file_racing_racing_proto_goTypes = nil
	file_racing_racing_proto_depIdxs = nil
}
