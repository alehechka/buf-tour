// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0-devel
// 	protoc        (unknown)
// source: pet/v1/pet.proto

package petv1

import (
	v1alpha1 "github.com/alehechka/buf-tour/petstore/gen/proto/go/payment/v1alpha1"
	_ "github.com/ysugimoto/grpc-graphql-gateway/graphql"
	datetime "google.golang.org/genproto/googleapis/type/datetime"
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

// Pet represents a pet in the pet store.
type Pet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PetType   PetType            `protobuf:"varint,1,opt,name=pet_type,json=petType,proto3,enum=pet.v1.PetType" json:"pet_type,omitempty"`
	PetId     string             `protobuf:"bytes,2,opt,name=pet_id,json=petId,proto3" json:"pet_id,omitempty"`
	Name      string             `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	CreatedAt *datetime.DateTime `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *Pet) Reset() {
	*x = Pet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pet_v1_pet_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pet) ProtoMessage() {}

func (x *Pet) ProtoReflect() protoreflect.Message {
	mi := &file_pet_v1_pet_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pet.ProtoReflect.Descriptor instead.
func (*Pet) Descriptor() ([]byte, []int) {
	return file_pet_v1_pet_proto_rawDescGZIP(), []int{0}
}

func (x *Pet) GetPetType() PetType {
	if x != nil {
		return x.PetType
	}
	return PetType_PET_TYPE_UNSPECIFIED
}

func (x *Pet) GetPetId() string {
	if x != nil {
		return x.PetId
	}
	return ""
}

func (x *Pet) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Pet) GetCreatedAt() *datetime.DateTime {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type GetPetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PetId string `protobuf:"bytes,1,opt,name=pet_id,json=petId,proto3" json:"pet_id,omitempty"`
}

func (x *GetPetRequest) Reset() {
	*x = GetPetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pet_v1_pet_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPetRequest) ProtoMessage() {}

func (x *GetPetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pet_v1_pet_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPetRequest.ProtoReflect.Descriptor instead.
func (*GetPetRequest) Descriptor() ([]byte, []int) {
	return file_pet_v1_pet_proto_rawDescGZIP(), []int{1}
}

func (x *GetPetRequest) GetPetId() string {
	if x != nil {
		return x.PetId
	}
	return ""
}

type GetPetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pet *Pet `protobuf:"bytes,1,opt,name=pet,proto3" json:"pet,omitempty"`
}

func (x *GetPetResponse) Reset() {
	*x = GetPetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pet_v1_pet_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPetResponse) ProtoMessage() {}

func (x *GetPetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pet_v1_pet_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPetResponse.ProtoReflect.Descriptor instead.
func (*GetPetResponse) Descriptor() ([]byte, []int) {
	return file_pet_v1_pet_proto_rawDescGZIP(), []int{2}
}

func (x *GetPetResponse) GetPet() *Pet {
	if x != nil {
		return x.Pet
	}
	return nil
}

type PutPetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PetType PetType `protobuf:"varint,1,opt,name=pet_type,json=petType,proto3,enum=pet.v1.PetType" json:"pet_type,omitempty"`
	Name    string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *PutPetRequest) Reset() {
	*x = PutPetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pet_v1_pet_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutPetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutPetRequest) ProtoMessage() {}

func (x *PutPetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pet_v1_pet_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutPetRequest.ProtoReflect.Descriptor instead.
func (*PutPetRequest) Descriptor() ([]byte, []int) {
	return file_pet_v1_pet_proto_rawDescGZIP(), []int{3}
}

func (x *PutPetRequest) GetPetType() PetType {
	if x != nil {
		return x.PetType
	}
	return PetType_PET_TYPE_UNSPECIFIED
}

func (x *PutPetRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type PutPetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pet *Pet `protobuf:"bytes,1,opt,name=pet,proto3" json:"pet,omitempty"`
}

func (x *PutPetResponse) Reset() {
	*x = PutPetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pet_v1_pet_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutPetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutPetResponse) ProtoMessage() {}

func (x *PutPetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pet_v1_pet_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutPetResponse.ProtoReflect.Descriptor instead.
func (*PutPetResponse) Descriptor() ([]byte, []int) {
	return file_pet_v1_pet_proto_rawDescGZIP(), []int{4}
}

func (x *PutPetResponse) GetPet() *Pet {
	if x != nil {
		return x.Pet
	}
	return nil
}

type DeletePetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PetId string `protobuf:"bytes,1,opt,name=pet_id,json=petId,proto3" json:"pet_id,omitempty"`
}

func (x *DeletePetRequest) Reset() {
	*x = DeletePetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pet_v1_pet_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePetRequest) ProtoMessage() {}

func (x *DeletePetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pet_v1_pet_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePetRequest.ProtoReflect.Descriptor instead.
func (*DeletePetRequest) Descriptor() ([]byte, []int) {
	return file_pet_v1_pet_proto_rawDescGZIP(), []int{5}
}

func (x *DeletePetRequest) GetPetId() string {
	if x != nil {
		return x.PetId
	}
	return ""
}

type DeletePetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeletePetResponse) Reset() {
	*x = DeletePetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pet_v1_pet_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePetResponse) ProtoMessage() {}

func (x *DeletePetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pet_v1_pet_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePetResponse.ProtoReflect.Descriptor instead.
func (*DeletePetResponse) Descriptor() ([]byte, []int) {
	return file_pet_v1_pet_proto_rawDescGZIP(), []int{6}
}

type PurchasePetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PetId string          `protobuf:"bytes,1,opt,name=pet_id,json=petId,proto3" json:"pet_id,omitempty"`
	Order *v1alpha1.Order `protobuf:"bytes,2,opt,name=order,proto3" json:"order,omitempty"`
}

func (x *PurchasePetRequest) Reset() {
	*x = PurchasePetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pet_v1_pet_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PurchasePetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PurchasePetRequest) ProtoMessage() {}

func (x *PurchasePetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pet_v1_pet_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PurchasePetRequest.ProtoReflect.Descriptor instead.
func (*PurchasePetRequest) Descriptor() ([]byte, []int) {
	return file_pet_v1_pet_proto_rawDescGZIP(), []int{7}
}

func (x *PurchasePetRequest) GetPetId() string {
	if x != nil {
		return x.PetId
	}
	return ""
}

func (x *PurchasePetRequest) GetOrder() *v1alpha1.Order {
	if x != nil {
		return x.Order
	}
	return nil
}

type PurchasePetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PurchasePetResponse) Reset() {
	*x = PurchasePetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pet_v1_pet_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PurchasePetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PurchasePetResponse) ProtoMessage() {}

func (x *PurchasePetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pet_v1_pet_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PurchasePetResponse.ProtoReflect.Descriptor instead.
func (*PurchasePetResponse) Descriptor() ([]byte, []int) {
	return file_pet_v1_pet_proto_rawDescGZIP(), []int{8}
}

var File_pet_v1_pet_proto protoreflect.FileDescriptor

var file_pet_v1_pet_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x65, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x06, 0x70, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x1e, 0x70, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x70, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x64, 0x61, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x67, 0x72, 0x61, 0x70, 0x68, 0x71, 0x6c, 0x2f,
	0x67, 0x72, 0x61, 0x70, 0x68, 0x71, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x70,
	0x65, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x65, 0x74, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xa0, 0x01, 0x0a, 0x03, 0x50, 0x65, 0x74, 0x12, 0x31, 0x0a, 0x08, 0x70,
	0x65, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e,
	0x70, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x42, 0x05,
	0xba, 0x43, 0x02, 0x08, 0x01, 0x52, 0x07, 0x70, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x15,
	0x0a, 0x06, 0x70, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x70, 0x65, 0x74, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x05, 0xba, 0x43, 0x02, 0x08, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x34, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x2d, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x50, 0x65, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x06, 0x70, 0x65, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x05, 0xba, 0x43, 0x02, 0x08, 0x01, 0x52, 0x05,
	0x70, 0x65, 0x74, 0x49, 0x64, 0x22, 0x2f, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x50, 0x65, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x03, 0x70, 0x65, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65,
	0x74, 0x52, 0x03, 0x70, 0x65, 0x74, 0x22, 0x5d, 0x0a, 0x0d, 0x50, 0x75, 0x74, 0x50, 0x65, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x08, 0x70, 0x65, 0x74, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x70, 0x65, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x50, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x42, 0x05, 0xba, 0x43, 0x02, 0x08,
	0x01, 0x52, 0x07, 0x70, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x05, 0xba, 0x43, 0x02, 0x08, 0x01, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x2f, 0x0a, 0x0e, 0x50, 0x75, 0x74, 0x50, 0x65, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x03, 0x70, 0x65, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65,
	0x74, 0x52, 0x03, 0x70, 0x65, 0x74, 0x22, 0x30, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x50, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x06, 0x70, 0x65,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x05, 0xba, 0x43, 0x02, 0x08,
	0x01, 0x52, 0x05, 0x70, 0x65, 0x74, 0x49, 0x64, 0x22, 0x13, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x50, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x68, 0x0a,
	0x12, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x50, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x06, 0x70, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x05, 0xba, 0x43, 0x02, 0x08, 0x01, 0x52, 0x05, 0x70, 0x65, 0x74, 0x49,
	0x64, 0x12, 0x34, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x05, 0xba, 0x43, 0x02, 0x08, 0x01,
	0x52, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x22, 0x15, 0x0a, 0x13, 0x50, 0x75, 0x72, 0x63, 0x68,
	0x61, 0x73, 0x65, 0x50, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xe6,
	0x02, 0x0a, 0x0f, 0x50, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x44, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x50, 0x65, 0x74, 0x12, 0x15, 0x2e, 0x70,
	0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x50, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x0b, 0xba, 0x43, 0x08,
	0x12, 0x06, 0x67, 0x65, 0x74, 0x50, 0x65, 0x74, 0x12, 0x46, 0x0a, 0x06, 0x50, 0x75, 0x74, 0x50,
	0x65, 0x74, 0x12, 0x15, 0x2e, 0x70, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x75, 0x74, 0x50,
	0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x65, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x50, 0x75, 0x74, 0x50, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x0d, 0xba, 0x43, 0x0a, 0x08, 0x01, 0x12, 0x06, 0x70, 0x75, 0x74, 0x50, 0x65, 0x74,
	0x12, 0x52, 0x0a, 0x09, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x65, 0x74, 0x12, 0x18, 0x2e,
	0x70, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x65, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x65, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x10, 0xba, 0x43, 0x0d, 0x08, 0x01, 0x12, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x50, 0x65, 0x74, 0x12, 0x5a, 0x0a, 0x0b, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65,
	0x50, 0x65, 0x74, 0x12, 0x1a, 0x2e, 0x70, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x75, 0x72,
	0x63, 0x68, 0x61, 0x73, 0x65, 0x50, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1b, 0x2e, 0x70, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73,
	0x65, 0x50, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x12, 0xba, 0x43,
	0x0f, 0x08, 0x01, 0x12, 0x0b, 0x70, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x50, 0x65, 0x74,
	0x1a, 0x15, 0xba, 0x43, 0x12, 0x0a, 0x0e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x68, 0x6f, 0x73, 0x74,
	0x3a, 0x38, 0x30, 0x38, 0x30, 0x10, 0x01, 0x42, 0x91, 0x01, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x2e,
	0x70, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x42, 0x08, 0x50, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61,
	0x6c, 0x65, 0x68, 0x65, 0x63, 0x68, 0x6b, 0x61, 0x2f, 0x62, 0x75, 0x66, 0x2d, 0x74, 0x6f, 0x75,
	0x72, 0x2f, 0x70, 0x65, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x65, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x70,
	0x65, 0x74, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x50, 0x58, 0x58, 0xaa, 0x02, 0x06, 0x50, 0x65, 0x74,
	0x2e, 0x56, 0x31, 0xca, 0x02, 0x06, 0x50, 0x65, 0x74, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x12, 0x50,
	0x65, 0x74, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x07, 0x50, 0x65, 0x74, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_pet_v1_pet_proto_rawDescOnce sync.Once
	file_pet_v1_pet_proto_rawDescData = file_pet_v1_pet_proto_rawDesc
)

func file_pet_v1_pet_proto_rawDescGZIP() []byte {
	file_pet_v1_pet_proto_rawDescOnce.Do(func() {
		file_pet_v1_pet_proto_rawDescData = protoimpl.X.CompressGZIP(file_pet_v1_pet_proto_rawDescData)
	})
	return file_pet_v1_pet_proto_rawDescData
}

var file_pet_v1_pet_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_pet_v1_pet_proto_goTypes = []interface{}{
	(*Pet)(nil),                 // 0: pet.v1.Pet
	(*GetPetRequest)(nil),       // 1: pet.v1.GetPetRequest
	(*GetPetResponse)(nil),      // 2: pet.v1.GetPetResponse
	(*PutPetRequest)(nil),       // 3: pet.v1.PutPetRequest
	(*PutPetResponse)(nil),      // 4: pet.v1.PutPetResponse
	(*DeletePetRequest)(nil),    // 5: pet.v1.DeletePetRequest
	(*DeletePetResponse)(nil),   // 6: pet.v1.DeletePetResponse
	(*PurchasePetRequest)(nil),  // 7: pet.v1.PurchasePetRequest
	(*PurchasePetResponse)(nil), // 8: pet.v1.PurchasePetResponse
	(PetType)(0),                // 9: pet.v1.PetType
	(*datetime.DateTime)(nil),   // 10: google.type.DateTime
	(*v1alpha1.Order)(nil),      // 11: payment.v1alpha1.Order
}
var file_pet_v1_pet_proto_depIdxs = []int32{
	9,  // 0: pet.v1.Pet.pet_type:type_name -> pet.v1.PetType
	10, // 1: pet.v1.Pet.created_at:type_name -> google.type.DateTime
	0,  // 2: pet.v1.GetPetResponse.pet:type_name -> pet.v1.Pet
	9,  // 3: pet.v1.PutPetRequest.pet_type:type_name -> pet.v1.PetType
	0,  // 4: pet.v1.PutPetResponse.pet:type_name -> pet.v1.Pet
	11, // 5: pet.v1.PurchasePetRequest.order:type_name -> payment.v1alpha1.Order
	1,  // 6: pet.v1.PetStoreService.GetPet:input_type -> pet.v1.GetPetRequest
	3,  // 7: pet.v1.PetStoreService.PutPet:input_type -> pet.v1.PutPetRequest
	5,  // 8: pet.v1.PetStoreService.DeletePet:input_type -> pet.v1.DeletePetRequest
	7,  // 9: pet.v1.PetStoreService.PurchasePet:input_type -> pet.v1.PurchasePetRequest
	2,  // 10: pet.v1.PetStoreService.GetPet:output_type -> pet.v1.GetPetResponse
	4,  // 11: pet.v1.PetStoreService.PutPet:output_type -> pet.v1.PutPetResponse
	6,  // 12: pet.v1.PetStoreService.DeletePet:output_type -> pet.v1.DeletePetResponse
	8,  // 13: pet.v1.PetStoreService.PurchasePet:output_type -> pet.v1.PurchasePetResponse
	10, // [10:14] is the sub-list for method output_type
	6,  // [6:10] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_pet_v1_pet_proto_init() }
func file_pet_v1_pet_proto_init() {
	if File_pet_v1_pet_proto != nil {
		return
	}
	file_pet_v1_pettype_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_pet_v1_pet_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pet); i {
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
		file_pet_v1_pet_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPetRequest); i {
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
		file_pet_v1_pet_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPetResponse); i {
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
		file_pet_v1_pet_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutPetRequest); i {
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
		file_pet_v1_pet_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutPetResponse); i {
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
		file_pet_v1_pet_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePetRequest); i {
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
		file_pet_v1_pet_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePetResponse); i {
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
		file_pet_v1_pet_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PurchasePetRequest); i {
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
		file_pet_v1_pet_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PurchasePetResponse); i {
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
			RawDescriptor: file_pet_v1_pet_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pet_v1_pet_proto_goTypes,
		DependencyIndexes: file_pet_v1_pet_proto_depIdxs,
		MessageInfos:      file_pet_v1_pet_proto_msgTypes,
	}.Build()
	File_pet_v1_pet_proto = out.File
	file_pet_v1_pet_proto_rawDesc = nil
	file_pet_v1_pet_proto_goTypes = nil
	file_pet_v1_pet_proto_depIdxs = nil
}
