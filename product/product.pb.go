// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.14.0
// source: product/product.proto

package product

import (
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

type QueryProductListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Extra map[string]string `protobuf:"bytes,1,rep,name=extra,proto3" json:"extra,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *QueryProductListRequest) Reset() {
	*x = QueryProductListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_product_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryProductListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryProductListRequest) ProtoMessage() {}

func (x *QueryProductListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_product_product_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryProductListRequest.ProtoReflect.Descriptor instead.
func (*QueryProductListRequest) Descriptor() ([]byte, []int) {
	return file_product_product_proto_rawDescGZIP(), []int{0}
}

func (x *QueryProductListRequest) GetExtra() map[string]string {
	if x != nil {
		return x.Extra
	}
	return nil
}

type QueryProductListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data         *QueryProductListResponse_Data `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	RequestId    string                         `protobuf:"bytes,2,opt,name=RequestId,proto3" json:"RequestId,omitempty"`
	Code         string                         `protobuf:"bytes,3,opt,name=Code,proto3" json:"Code,omitempty"`
	Success      bool                           `protobuf:"varint,4,opt,name=Success,proto3" json:"Success,omitempty"`
	ErrorMessage string                         `protobuf:"bytes,5,opt,name=errorMessage,proto3" json:"errorMessage,omitempty"`
}

func (x *QueryProductListResponse) Reset() {
	*x = QueryProductListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_product_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryProductListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryProductListResponse) ProtoMessage() {}

func (x *QueryProductListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_product_product_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryProductListResponse.ProtoReflect.Descriptor instead.
func (*QueryProductListResponse) Descriptor() ([]byte, []int) {
	return file_product_product_proto_rawDescGZIP(), []int{1}
}

func (x *QueryProductListResponse) GetData() *QueryProductListResponse_Data {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *QueryProductListResponse) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *QueryProductListResponse) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *QueryProductListResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *QueryProductListResponse) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

type QueryProductListResponse_Productinfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CreateAt    uint64            `protobuf:"varint,1,opt,name=CreateAt,proto3" json:"CreateAt,omitempty"`       //创建时间
	ProductId   string            `protobuf:"bytes,2,opt,name=ProductId,proto3" json:"ProductId,omitempty"`      //产品id
	ProductName string            `protobuf:"bytes,3,opt,name=ProductName,proto3" json:"ProductName,omitempty"`  //产品名称
	Description string            `protobuf:"bytes,4,opt,name=Description,proto3" json:"Description,omitempty"`  //描述
	NodeType    uint32            `protobuf:"varint,5,opt,name=NodeType,proto3" json:"NodeType,omitempty"`       //产品的节点类型 0设备，1网关
	DataFormat  uint32            `protobuf:"varint,6,opt,name=DataFormat,proto3" json:"DataFormat,omitempty"`   //上报数据的格式
	DeviceCount uint32            `protobuf:"varint,7,opt,name=DeviceCount,proto3" json:"DeviceCount,omitempty"` //设备数量
	Platform    string            `protobuf:"bytes,8,opt,name=Platform,proto3" json:"Platform,omitempty"`        //平台名称
	Extra       map[string]string `protobuf:"bytes,9,rep,name=extra,proto3" json:"extra,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *QueryProductListResponse_Productinfo) Reset() {
	*x = QueryProductListResponse_Productinfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_product_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryProductListResponse_Productinfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryProductListResponse_Productinfo) ProtoMessage() {}

func (x *QueryProductListResponse_Productinfo) ProtoReflect() protoreflect.Message {
	mi := &file_product_product_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryProductListResponse_Productinfo.ProtoReflect.Descriptor instead.
func (*QueryProductListResponse_Productinfo) Descriptor() ([]byte, []int) {
	return file_product_product_proto_rawDescGZIP(), []int{1, 0}
}

func (x *QueryProductListResponse_Productinfo) GetCreateAt() uint64 {
	if x != nil {
		return x.CreateAt
	}
	return 0
}

func (x *QueryProductListResponse_Productinfo) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *QueryProductListResponse_Productinfo) GetProductName() string {
	if x != nil {
		return x.ProductName
	}
	return ""
}

func (x *QueryProductListResponse_Productinfo) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *QueryProductListResponse_Productinfo) GetNodeType() uint32 {
	if x != nil {
		return x.NodeType
	}
	return 0
}

func (x *QueryProductListResponse_Productinfo) GetDataFormat() uint32 {
	if x != nil {
		return x.DataFormat
	}
	return 0
}

func (x *QueryProductListResponse_Productinfo) GetDeviceCount() uint32 {
	if x != nil {
		return x.DeviceCount
	}
	return 0
}

func (x *QueryProductListResponse_Productinfo) GetPlatform() string {
	if x != nil {
		return x.Platform
	}
	return ""
}

func (x *QueryProductListResponse_Productinfo) GetExtra() map[string]string {
	if x != nil {
		return x.Extra
	}
	return nil
}

type QueryProductListResponse_List struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductInfo []*QueryProductListResponse_Productinfo `protobuf:"bytes,1,rep,name=ProductInfo,proto3" json:"ProductInfo,omitempty"`
}

func (x *QueryProductListResponse_List) Reset() {
	*x = QueryProductListResponse_List{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_product_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryProductListResponse_List) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryProductListResponse_List) ProtoMessage() {}

func (x *QueryProductListResponse_List) ProtoReflect() protoreflect.Message {
	mi := &file_product_product_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryProductListResponse_List.ProtoReflect.Descriptor instead.
func (*QueryProductListResponse_List) Descriptor() ([]byte, []int) {
	return file_product_product_proto_rawDescGZIP(), []int{1, 1}
}

func (x *QueryProductListResponse_List) GetProductInfo() []*QueryProductListResponse_Productinfo {
	if x != nil {
		return x.ProductInfo
	}
	return nil
}

type QueryProductListResponse_Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total uint32                         `protobuf:"varint,4,opt,name=Total,proto3" json:"Total,omitempty"`
	List  *QueryProductListResponse_List `protobuf:"bytes,5,opt,name=List,proto3" json:"List,omitempty"`
}

func (x *QueryProductListResponse_Data) Reset() {
	*x = QueryProductListResponse_Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_product_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryProductListResponse_Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryProductListResponse_Data) ProtoMessage() {}

func (x *QueryProductListResponse_Data) ProtoReflect() protoreflect.Message {
	mi := &file_product_product_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryProductListResponse_Data.ProtoReflect.Descriptor instead.
func (*QueryProductListResponse_Data) Descriptor() ([]byte, []int) {
	return file_product_product_proto_rawDescGZIP(), []int{1, 2}
}

func (x *QueryProductListResponse_Data) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *QueryProductListResponse_Data) GetList() *QueryProductListResponse_List {
	if x != nil {
		return x.List
	}
	return nil
}

var File_product_product_proto protoreflect.FileDescriptor

var file_product_product_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x22, 0x96, 0x01, 0x0a, 0x17, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x41, 0x0a, 0x05,
	0x65, 0x78, 0x74, 0x72, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x45, 0x78,
	0x74, 0x72, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x1a,
	0x38, 0x0a, 0x0a, 0x45, 0x78, 0x74, 0x72, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x8b, 0x06, 0x0a, 0x18, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x1c, 0x0a, 0x09, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x43, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x22,
	0x0a, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x1a, 0x8f, 0x03, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6e,
	0x66, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x1c,
	0x0a, 0x09, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20,
	0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x1a, 0x0a, 0x08, 0x4e, 0x6f, 0x64, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x08, 0x4e, 0x6f, 0x64, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a, 0x0a,
	0x44, 0x61, 0x74, 0x61, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0a, 0x44, 0x61, 0x74, 0x61, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x20, 0x0a, 0x0b,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0b, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x4e, 0x0a, 0x05, 0x65, 0x78,
	0x74, 0x72, 0x61, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x45, 0x78, 0x74, 0x72, 0x61, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x1a, 0x38, 0x0a, 0x0a, 0x45, 0x78,
	0x74, 0x72, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x1a, 0x57, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x4f, 0x0a, 0x0b,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x2d, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6e, 0x66, 0x6f,
	0x52, 0x0b, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x58, 0x0a,
	0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x3a, 0x0a, 0x04, 0x4c,
	0x69, 0x73, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x32, 0x67, 0x0a, 0x0a, 0x52, 0x70, 0x63, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x59, 0x0a, 0x10, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77,
	0x69, 0x6e, 0x63, 0x2d, 0x6c, 0x69, 0x6e, 0x6b, 0x2f, 0x69, 0x6f, 0x74, 0x2d, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_product_product_proto_rawDescOnce sync.Once
	file_product_product_proto_rawDescData = file_product_product_proto_rawDesc
)

func file_product_product_proto_rawDescGZIP() []byte {
	file_product_product_proto_rawDescOnce.Do(func() {
		file_product_product_proto_rawDescData = protoimpl.X.CompressGZIP(file_product_product_proto_rawDescData)
	})
	return file_product_product_proto_rawDescData
}

var file_product_product_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_product_product_proto_goTypes = []interface{}{
	(*QueryProductListRequest)(nil),  // 0: product.QueryProductListRequest
	(*QueryProductListResponse)(nil), // 1: product.QueryProductListResponse
	nil,                              // 2: product.QueryProductListRequest.ExtraEntry
	(*QueryProductListResponse_Productinfo)(nil), // 3: product.QueryProductListResponse.Productinfo
	(*QueryProductListResponse_List)(nil),        // 4: product.QueryProductListResponse.List
	(*QueryProductListResponse_Data)(nil),        // 5: product.QueryProductListResponse.Data
	nil,                                          // 6: product.QueryProductListResponse.Productinfo.ExtraEntry
}
var file_product_product_proto_depIdxs = []int32{
	2, // 0: product.QueryProductListRequest.extra:type_name -> product.QueryProductListRequest.ExtraEntry
	5, // 1: product.QueryProductListResponse.data:type_name -> product.QueryProductListResponse.Data
	6, // 2: product.QueryProductListResponse.Productinfo.extra:type_name -> product.QueryProductListResponse.Productinfo.ExtraEntry
	3, // 3: product.QueryProductListResponse.List.ProductInfo:type_name -> product.QueryProductListResponse.Productinfo
	4, // 4: product.QueryProductListResponse.Data.List:type_name -> product.QueryProductListResponse.List
	0, // 5: product.RpcProduct.QueryProductList:input_type -> product.QueryProductListRequest
	1, // 6: product.RpcProduct.QueryProductList:output_type -> product.QueryProductListResponse
	6, // [6:7] is the sub-list for method output_type
	5, // [5:6] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_product_product_proto_init() }
func file_product_product_proto_init() {
	if File_product_product_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_product_product_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryProductListRequest); i {
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
		file_product_product_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryProductListResponse); i {
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
		file_product_product_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryProductListResponse_Productinfo); i {
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
		file_product_product_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryProductListResponse_List); i {
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
		file_product_product_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryProductListResponse_Data); i {
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
			RawDescriptor: file_product_product_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_product_product_proto_goTypes,
		DependencyIndexes: file_product_product_proto_depIdxs,
		MessageInfos:      file_product_product_proto_msgTypes,
	}.Build()
	File_product_product_proto = out.File
	file_product_product_proto_rawDesc = nil
	file_product_product_proto_goTypes = nil
	file_product_product_proto_depIdxs = nil
}
