// Code generated by protoc-gen-go. DO NOT EDIT.
// source: paginate.proto

/*
Package paginate is a generated protocol buffer package.

It is generated from these files:
	paginate.proto

It has these top-level messages:
	PageOptions
	PageResult
*/
package paginate

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// PaginateMode : MaxId or PageNumber, default PageNumber
// PaginateMode ： 0 == PageNumber , 1 == MaxId
// ------------------------------------------------------------
// Action : 1 == NextPage , 0 == PreviousPage
// default Action  NextPage
// default MaxId 0
// default first page : select * from table order by id desc limit 0,PageSize
// ------------------------------------------------------------
// OrderBy default id desc
// 自定义 OrderBy : usernaem desc,age desc,id asc
// 数组使用 strings.Join([]string, ",")
// 映射使用 helper.ImplodeMapInt32String(map[int32]string, ",")
// 映射使用 helper.ImplodeMapIntString(map[int]string, ",")
type PageOptions struct {
	PaginateMode int32  `protobuf:"varint,1,opt,name=PaginateMode" json:"PaginateMode,omitempty"`
	PageSize     int32  `protobuf:"varint,2,opt,name=PageSize" json:"PageSize,omitempty"`
	MaxId        int32  `protobuf:"varint,3,opt,name=MaxId" json:"MaxId,omitempty"`
	Action       int32  `protobuf:"varint,4,opt,name=Action" json:"Action,omitempty"`
	PageNumber   int32  `protobuf:"varint,5,opt,name=PageNumber" json:"PageNumber,omitempty"`
	OrderBy      string `protobuf:"bytes,6,opt,name=OrderBy" json:"OrderBy,omitempty"`
}

func (m *PageOptions) Reset()                    { *m = PageOptions{} }
func (m *PageOptions) String() string            { return proto.CompactTextString(m) }
func (*PageOptions) ProtoMessage()               {}
func (*PageOptions) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *PageOptions) GetPaginateMode() int32 {
	if m != nil {
		return m.PaginateMode
	}
	return 0
}

func (m *PageOptions) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *PageOptions) GetMaxId() int32 {
	if m != nil {
		return m.MaxId
	}
	return 0
}

func (m *PageOptions) GetAction() int32 {
	if m != nil {
		return m.Action
	}
	return 0
}

func (m *PageOptions) GetPageNumber() int32 {
	if m != nil {
		return m.PageNumber
	}
	return 0
}

func (m *PageOptions) GetOrderBy() string {
	if m != nil {
		return m.OrderBy
	}
	return ""
}

// total TotalRecords entries , total TotalPages pages
type PageResult struct {
	PaginateMode string `protobuf:"bytes,1,opt,name=PaginateMode" json:"PaginateMode,omitempty"`
	TotalRecords int32  `protobuf:"varint,2,opt,name=TotalRecords" json:"TotalRecords,omitempty"`
	TotalPages   int32  `protobuf:"varint,3,opt,name=TotalPages" json:"TotalPages,omitempty"`
	PageSize     int32  `protobuf:"varint,4,opt,name=PageSize" json:"PageSize,omitempty"`
	MaxId        int32  `protobuf:"varint,5,opt,name=MaxId" json:"MaxId,omitempty"`
	MinId        int32  `protobuf:"varint,6,opt,name=MinId" json:"MinId,omitempty"`
	PageNumber   int32  `protobuf:"varint,7,opt,name=PageNumber" json:"PageNumber,omitempty"`
	Action       string `protobuf:"bytes,8,opt,name=Action" json:"Action,omitempty"`
}

func (m *PageResult) Reset()                    { *m = PageResult{} }
func (m *PageResult) String() string            { return proto.CompactTextString(m) }
func (*PageResult) ProtoMessage()               {}
func (*PageResult) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *PageResult) GetPaginateMode() string {
	if m != nil {
		return m.PaginateMode
	}
	return ""
}

func (m *PageResult) GetTotalRecords() int32 {
	if m != nil {
		return m.TotalRecords
	}
	return 0
}

func (m *PageResult) GetTotalPages() int32 {
	if m != nil {
		return m.TotalPages
	}
	return 0
}

func (m *PageResult) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *PageResult) GetMaxId() int32 {
	if m != nil {
		return m.MaxId
	}
	return 0
}

func (m *PageResult) GetMinId() int32 {
	if m != nil {
		return m.MinId
	}
	return 0
}

func (m *PageResult) GetPageNumber() int32 {
	if m != nil {
		return m.PageNumber
	}
	return 0
}

func (m *PageResult) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func init() {
	proto.RegisterType((*PageOptions)(nil), "paginate.PageOptions")
	proto.RegisterType((*PageResult)(nil), "paginate.PageResult")
}

func init() { proto.RegisterFile("paginate.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 249 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2b, 0x48, 0x4c, 0xcf,
	0xcc, 0x4b, 0x2c, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x80, 0xf1, 0x95, 0xb6,
	0x32, 0x72, 0x71, 0x07, 0x24, 0xa6, 0xa7, 0xfa, 0x17, 0x94, 0x64, 0xe6, 0xe7, 0x15, 0x0b, 0x29,
	0x71, 0xf1, 0x04, 0x40, 0xe5, 0x7c, 0xf3, 0x53, 0x52, 0x25, 0x18, 0x15, 0x18, 0x35, 0x58, 0x83,
	0x50, 0xc4, 0x84, 0xa4, 0xb8, 0x38, 0x40, 0x5a, 0x82, 0x33, 0xab, 0x52, 0x25, 0x98, 0xc0, 0xf2,
	0x70, 0xbe, 0x90, 0x08, 0x17, 0xab, 0x6f, 0x62, 0x85, 0x67, 0x8a, 0x04, 0x33, 0x58, 0x02, 0xc2,
	0x11, 0x12, 0xe3, 0x62, 0x73, 0x4c, 0x06, 0x59, 0x20, 0xc1, 0x02, 0x16, 0x86, 0xf2, 0x84, 0xe4,
	0xb8, 0xb8, 0x40, 0x3a, 0xfd, 0x4a, 0x73, 0x93, 0x52, 0x8b, 0x24, 0x58, 0xc1, 0x72, 0x48, 0x22,
	0x42, 0x12, 0x5c, 0xec, 0xfe, 0x45, 0x29, 0xa9, 0x45, 0x4e, 0x95, 0x12, 0x6c, 0x0a, 0x8c, 0x1a,
	0x9c, 0x41, 0x30, 0xae, 0xd2, 0x17, 0x46, 0x88, 0xd6, 0xa0, 0xd4, 0xe2, 0xd2, 0x9c, 0x12, 0xac,
	0xce, 0xe6, 0x44, 0x73, 0xb6, 0x12, 0x17, 0x4f, 0x48, 0x7e, 0x49, 0x62, 0x4e, 0x50, 0x6a, 0x72,
	0x7e, 0x51, 0x4a, 0x31, 0xd4, 0xe9, 0x28, 0x62, 0x20, 0x07, 0x81, 0xf9, 0x20, 0xa3, 0x8b, 0xa1,
	0x7e, 0x40, 0x12, 0x41, 0xf1, 0x3a, 0x0b, 0x2e, 0xaf, 0xb3, 0x22, 0x7b, 0x1d, 0x24, 0x9a, 0x99,
	0xe7, 0x99, 0x02, 0xf6, 0x00, 0x48, 0x14, 0xc4, 0x41, 0xf3, 0x38, 0x3b, 0x86, 0xc7, 0x11, 0x01,
	0xc6, 0x01, 0xf6, 0x09, 0x94, 0x97, 0xc4, 0x06, 0x8e, 0x3f, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x37, 0x07, 0x03, 0x58, 0xd1, 0x01, 0x00, 0x00,
}
