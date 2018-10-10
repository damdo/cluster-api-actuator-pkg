// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/firestore/admin/v1beta2/field.proto

package admin

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Represents a single field in the database.
//
// Fields are grouped by their "Collection Group", which represent all
// collections in the database with the same id.
type Field struct {
	// A field name of the form
	// `projects/{project_id}/databases/{database_id}/collectionGroups/{collection_id}/fields/{field_path}`
	//
	// A field path may be a simple field name, e.g. `address` or a path to fields
	// within map_value , e.g. `address.city`,
	// or a special field path. The only valid special field is `*`, which
	// represents any field.
	//
	// Field paths may be quoted using ` (backtick). The only character that needs
	// to be escaped within a quoted field path is the backtick character itself,
	// escaped using a backslash. Special characters in field paths that
	// must be quoted include: `*`, `.`,
	// ``` (backtick), `[`, `]`, as well as any ascii symbolic characters.
	//
	// Examples:
	// (Note: Comments here are written in markdown syntax, so there is an
	//  additional layer of backticks to represent a code block)
	// `\`address.city\`` represents a field named `address.city`, not the map key
	// `city` in the field `address`.
	// `\`*\`` represents a field named `*`, not any field.
	//
	// A special `Field` contains the default indexing settings for all fields.
	// This field's resource name is:
	// `projects/{project_id}/databases/{database_id}/collectionGroups/__default__/fields/*`
	// Indexes defined on this `Field` will be applied to all fields which do not
	// have their own `Field` index configuration.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The index configuration for this field. If unset, field indexing will
	// revert to the configuration defined by the `ancestor_field`. To
	// explicitly remove all indexes for this field, specify an index config
	// with an empty list of indexes.
	IndexConfig          *Field_IndexConfig `protobuf:"bytes,2,opt,name=index_config,json=indexConfig,proto3" json:"index_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Field) Reset()         { *m = Field{} }
func (m *Field) String() string { return proto.CompactTextString(m) }
func (*Field) ProtoMessage()    {}
func (*Field) Descriptor() ([]byte, []int) {
	return fileDescriptor_967ea3483ba729a5, []int{0}
}

func (m *Field) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Field.Unmarshal(m, b)
}
func (m *Field) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Field.Marshal(b, m, deterministic)
}
func (m *Field) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Field.Merge(m, src)
}
func (m *Field) XXX_Size() int {
	return xxx_messageInfo_Field.Size(m)
}
func (m *Field) XXX_DiscardUnknown() {
	xxx_messageInfo_Field.DiscardUnknown(m)
}

var xxx_messageInfo_Field proto.InternalMessageInfo

func (m *Field) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Field) GetIndexConfig() *Field_IndexConfig {
	if m != nil {
		return m.IndexConfig
	}
	return nil
}

// The index configuration for this field.
type Field_IndexConfig struct {
	// The indexes supported for this field.
	Indexes []*Index `protobuf:"bytes,1,rep,name=indexes,proto3" json:"indexes,omitempty"`
	// Output only.
	// When true, the `Field`'s index configuration is set from the
	// configuration specified by the `ancestor_field`.
	// When false, the `Field`'s index configuration is defined explicitly.
	UsesAncestorConfig bool `protobuf:"varint,2,opt,name=uses_ancestor_config,json=usesAncestorConfig,proto3" json:"uses_ancestor_config,omitempty"`
	// Output only.
	// Specifies the resource name of the `Field` from which this field's
	// index configuration is set (when `uses_ancestor_config` is true),
	// or from which it *would* be set if this field had no index configuration
	// (when `uses_ancestor_config` is false).
	AncestorField string `protobuf:"bytes,3,opt,name=ancestor_field,json=ancestorField,proto3" json:"ancestor_field,omitempty"`
	// Output only
	// When true, the `Field`'s index configuration is in the process of being
	// reverted. Once complete, the index config will transition to the same
	// state as the field specified by `ancestor_field`, at which point
	// `uses_ancestor_config` will be `true` and `reverting` will be `false`.
	Reverting            bool     `protobuf:"varint,4,opt,name=reverting,proto3" json:"reverting,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Field_IndexConfig) Reset()         { *m = Field_IndexConfig{} }
func (m *Field_IndexConfig) String() string { return proto.CompactTextString(m) }
func (*Field_IndexConfig) ProtoMessage()    {}
func (*Field_IndexConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_967ea3483ba729a5, []int{0, 0}
}

func (m *Field_IndexConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Field_IndexConfig.Unmarshal(m, b)
}
func (m *Field_IndexConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Field_IndexConfig.Marshal(b, m, deterministic)
}
func (m *Field_IndexConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Field_IndexConfig.Merge(m, src)
}
func (m *Field_IndexConfig) XXX_Size() int {
	return xxx_messageInfo_Field_IndexConfig.Size(m)
}
func (m *Field_IndexConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_Field_IndexConfig.DiscardUnknown(m)
}

var xxx_messageInfo_Field_IndexConfig proto.InternalMessageInfo

func (m *Field_IndexConfig) GetIndexes() []*Index {
	if m != nil {
		return m.Indexes
	}
	return nil
}

func (m *Field_IndexConfig) GetUsesAncestorConfig() bool {
	if m != nil {
		return m.UsesAncestorConfig
	}
	return false
}

func (m *Field_IndexConfig) GetAncestorField() string {
	if m != nil {
		return m.AncestorField
	}
	return ""
}

func (m *Field_IndexConfig) GetReverting() bool {
	if m != nil {
		return m.Reverting
	}
	return false
}

func init() {
	proto.RegisterType((*Field)(nil), "google.firestore.admin.v1beta2.Field")
	proto.RegisterType((*Field_IndexConfig)(nil), "google.firestore.admin.v1beta2.Field.IndexConfig")
}

func init() {
	proto.RegisterFile("google/firestore/admin/v1beta2/field.proto", fileDescriptor_967ea3483ba729a5)
}

var fileDescriptor_967ea3483ba729a5 = []byte{
	// 346 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xcf, 0x4a, 0xf3, 0x40,
	0x14, 0xc5, 0x99, 0xb4, 0xdf, 0xa7, 0x9d, 0xa8, 0x8b, 0xc1, 0x45, 0x28, 0x45, 0x4a, 0xb1, 0x50,
	0x5c, 0xcc, 0xd8, 0xba, 0x74, 0x21, 0x6d, 0xa0, 0xc5, 0x5d, 0x89, 0xe2, 0xc2, 0x4d, 0x99, 0x36,
	0xd3, 0x61, 0x20, 0x9d, 0x5b, 0x92, 0xb4, 0xf8, 0x3c, 0x82, 0x1b, 0xdf, 0xc0, 0x07, 0xf0, 0xbd,
	0x24, 0x77, 0xd2, 0x3f, 0x1b, 0xcd, 0x2e, 0x33, 0xe7, 0x77, 0xce, 0x3d, 0xb9, 0x43, 0x6f, 0x34,
	0x80, 0x4e, 0x94, 0x58, 0x9a, 0x54, 0x65, 0x39, 0xa4, 0x4a, 0xc8, 0x78, 0x65, 0xac, 0xd8, 0xf6,
	0xe7, 0x2a, 0x97, 0x03, 0xb1, 0x34, 0x2a, 0x89, 0xf9, 0x3a, 0x85, 0x1c, 0xd8, 0x95, 0x63, 0xf9,
	0x9e, 0xe5, 0xc8, 0xf2, 0x92, 0x6d, 0xb6, 0xca, 0x2c, 0xb9, 0x36, 0x42, 0x5a, 0x0b, 0xb9, 0xcc,
	0x0d, 0xd8, 0xcc, 0xb9, 0x9b, 0x55, 0x93, 0x8c, 0x8d, 0xd5, 0x9b, 0x63, 0x3b, 0x5f, 0x1e, 0xfd,
	0x37, 0x2e, 0x26, 0x33, 0x46, 0xeb, 0x56, 0xae, 0x54, 0x40, 0xda, 0xa4, 0xd7, 0x88, 0xf0, 0x9b,
	0x3d, 0xd3, 0x33, 0x84, 0x67, 0x0b, 0xb0, 0x4b, 0xa3, 0x03, 0xaf, 0x4d, 0x7a, 0xfe, 0xa0, 0xcf,
	0xff, 0xae, 0xc7, 0x31, 0x90, 0x3f, 0x16, 0xce, 0x10, 0x8d, 0x91, 0x6f, 0x0e, 0x87, 0xe6, 0x37,
	0xa1, 0xfe, 0x91, 0xc8, 0x1e, 0xe8, 0x09, 0xca, 0x2a, 0x0b, 0x48, 0xbb, 0xd6, 0xf3, 0x07, 0xdd,
	0xaa, 0x01, 0xe8, 0x8e, 0x76, 0x2e, 0x76, 0x4b, 0x2f, 0x37, 0x99, 0xca, 0x66, 0xd2, 0x2e, 0x10,
	0x3f, 0xae, 0x7b, 0x1a, 0xb1, 0x42, 0x1b, 0x96, 0x52, 0x39, 0xb2, 0x4b, 0x2f, 0xf6, 0x30, 0x2e,
	0x3e, 0xa8, 0xe1, 0x6f, 0x9f, 0xef, 0x6e, 0xdd, 0x4e, 0x5a, 0xb4, 0x91, 0xaa, 0xad, 0x4a, 0x73,
	0x63, 0x75, 0x50, 0xc7, 0xb4, 0xc3, 0xc5, 0xe8, 0x83, 0xd0, 0xce, 0x02, 0x56, 0x15, 0x65, 0x47,
	0x14, 0xb3, 0xa6, 0xc5, 0xba, 0xa7, 0xe4, 0x35, 0x2c, 0x69, 0x0d, 0x89, 0xb4, 0x9a, 0x43, 0xaa,
	0x85, 0x56, 0x16, 0x1f, 0x43, 0x38, 0x49, 0xae, 0x4d, 0xf6, 0xdb, 0xdb, 0xdd, 0xe3, 0xe9, 0xdd,
	0xab, 0x4f, 0xc2, 0xf1, 0xd3, 0xa7, 0x77, 0x3d, 0x71, 0x61, 0x61, 0x02, 0x9b, 0x98, 0x8f, 0xf7,
	0x05, 0x86, 0x58, 0xe0, 0xa5, 0x3f, 0x2a, 0x3c, 0xf3, 0xff, 0x98, 0x7e, 0xf7, 0x13, 0x00, 0x00,
	0xff, 0xff, 0x31, 0xd7, 0xaa, 0xc3, 0x82, 0x02, 0x00, 0x00,
}
