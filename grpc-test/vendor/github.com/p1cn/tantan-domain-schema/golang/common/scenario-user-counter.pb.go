// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common/scenario-user-counter.proto

package common

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ScenarioCategory int32

const (
	ScenarioCategory_SCENARIO_CATEGORY_FOOD  ScenarioCategory = 0
	ScenarioCategory_SCENARIO_CATEGORY_SPORT ScenarioCategory = 1
	ScenarioCategory_SCENARIO_CATEGORY_MOVIE ScenarioCategory = 2
	ScenarioCategory_SCENARIO_CATEGORY_MISC  ScenarioCategory = 3
)

var ScenarioCategory_name = map[int32]string{
	0: "SCENARIO_CATEGORY_FOOD",
	1: "SCENARIO_CATEGORY_SPORT",
	2: "SCENARIO_CATEGORY_MOVIE",
	3: "SCENARIO_CATEGORY_MISC",
}
var ScenarioCategory_value = map[string]int32{
	"SCENARIO_CATEGORY_FOOD":  0,
	"SCENARIO_CATEGORY_SPORT": 1,
	"SCENARIO_CATEGORY_MOVIE": 2,
	"SCENARIO_CATEGORY_MISC":  3,
}

func (x ScenarioCategory) String() string {
	return proto.EnumName(ScenarioCategory_name, int32(x))
}
func (ScenarioCategory) EnumDescriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

type ScenarioUserCounter struct {
	UserId                string           `protobuf:"bytes,1,opt,name=userId" json:"userId,omitempty"`
	Category              ScenarioCategory `protobuf:"varint,2,opt,name=category,enum=common.ScenarioCategory" json:"category,omitempty"`
	ScenarioIds           []string         `protobuf:"bytes,3,rep,name=scenarioIds" json:"scenarioIds,omitempty"`
	ScenarioExpiresTime   int64            `protobuf:"varint,4,opt,name=scenarioExpiresTime" json:"scenarioExpiresTime,omitempty"`
	ScenarioReceivedLikes int32            `protobuf:"varint,5,opt,name=scenarioReceivedLikes" json:"scenarioReceivedLikes,omitempty"`
}

func (m *ScenarioUserCounter) Reset()                    { *m = ScenarioUserCounter{} }
func (m *ScenarioUserCounter) String() string            { return proto.CompactTextString(m) }
func (*ScenarioUserCounter) ProtoMessage()               {}
func (*ScenarioUserCounter) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func (m *ScenarioUserCounter) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *ScenarioUserCounter) GetCategory() ScenarioCategory {
	if m != nil {
		return m.Category
	}
	return ScenarioCategory_SCENARIO_CATEGORY_FOOD
}

func (m *ScenarioUserCounter) GetScenarioIds() []string {
	if m != nil {
		return m.ScenarioIds
	}
	return nil
}

func (m *ScenarioUserCounter) GetScenarioExpiresTime() int64 {
	if m != nil {
		return m.ScenarioExpiresTime
	}
	return 0
}

func (m *ScenarioUserCounter) GetScenarioReceivedLikes() int32 {
	if m != nil {
		return m.ScenarioReceivedLikes
	}
	return 0
}

func init() {
	proto.RegisterType((*ScenarioUserCounter)(nil), "common.ScenarioUserCounter")
	proto.RegisterEnum("common.ScenarioCategory", ScenarioCategory_name, ScenarioCategory_value)
}

func init() { proto.RegisterFile("common/scenario-user-counter.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 310 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0x51, 0x4b, 0xf3, 0x30,
	0x14, 0x86, 0xbf, 0xac, 0xdf, 0x86, 0x8b, 0x20, 0x25, 0xc3, 0x59, 0xf4, 0xa6, 0xec, 0xaa, 0x08,
	0x5d, 0x75, 0xee, 0x0f, 0xcc, 0x5a, 0xa5, 0xe0, 0xac, 0xa4, 0x53, 0xd0, 0x9b, 0x91, 0xa5, 0x87,
	0x2e, 0x68, 0x93, 0x91, 0x74, 0xa2, 0xf7, 0xfe, 0x5d, 0xff, 0x83, 0xcc, 0xae, 0x22, 0xda, 0xcb,
	0x73, 0x9e, 0x27, 0x2f, 0x87, 0x37, 0x78, 0xc0, 0x55, 0x51, 0x28, 0x19, 0x18, 0x0e, 0x92, 0x69,
	0xa1, 0xfc, 0xb5, 0x01, 0xed, 0x73, 0xb5, 0x96, 0x25, 0xe8, 0xe1, 0x4a, 0xab, 0x52, 0x91, 0x4e,
	0xe5, 0x0c, 0x3e, 0x10, 0xee, 0xa5, 0x5b, 0xef, 0xce, 0x80, 0x0e, 0x2b, 0x8b, 0xf4, 0x71, 0x67,
	0xf3, 0x2a, 0xce, 0x1c, 0xe4, 0x22, 0xaf, 0x4b, 0xb7, 0x13, 0x19, 0xe3, 0x1d, 0xce, 0x4a, 0xc8,
	0x95, 0x7e, 0x73, 0x5a, 0x2e, 0xf2, 0xf6, 0x46, 0xce, 0xb0, 0x8a, 0x1a, 0xd6, 0x31, 0xe1, 0x96,
	0xd3, 0x6f, 0x93, 0xb8, 0x78, 0xb7, 0x3e, 0x26, 0xce, 0x8c, 0x63, 0xb9, 0x96, 0xd7, 0xa5, 0x3f,
	0x57, 0xe4, 0x04, 0xf7, 0xea, 0x31, 0x7a, 0x5d, 0x09, 0x0d, 0x66, 0x26, 0x0a, 0x70, 0xfe, 0xbb,
	0xc8, 0xb3, 0x68, 0x13, 0x22, 0x63, 0xbc, 0x5f, 0xaf, 0x29, 0x70, 0x10, 0x2f, 0x90, 0x5d, 0x8b,
	0x27, 0x30, 0x4e, 0xdb, 0x45, 0x5e, 0x9b, 0x36, 0xc3, 0xe3, 0x77, 0x84, 0xed, 0xdf, 0x87, 0x92,
	0x43, 0xdc, 0x4f, 0xc3, 0xe8, 0x66, 0x42, 0xe3, 0x64, 0x1e, 0x4e, 0x66, 0xd1, 0x55, 0x42, 0x1f,
	0xe6, 0x97, 0x49, 0x72, 0x61, 0xff, 0x23, 0x47, 0xf8, 0xe0, 0x2f, 0x4b, 0x6f, 0x13, 0x3a, 0xb3,
	0x51, 0x33, 0x9c, 0x26, 0xf7, 0x71, 0x64, 0xb7, 0x9a, 0x53, 0xa7, 0x71, 0x1a, 0xda, 0xd6, 0xf9,
	0xf8, 0x71, 0x94, 0x8b, 0x72, 0xb9, 0x5e, 0x6c, 0xca, 0x0b, 0x56, 0xa7, 0x5c, 0x06, 0x25, 0x93,
	0x25, 0x93, 0x7e, 0xa6, 0x0a, 0x26, 0xa4, 0x6f, 0xf8, 0x12, 0x0a, 0x16, 0xe4, 0xea, 0x99, 0xc9,
	0x3c, 0xa8, 0x1a, 0x5e, 0x74, 0xbe, 0xfe, 0xee, 0xec, 0x33, 0x00, 0x00, 0xff, 0xff, 0x53, 0xa7,
	0xc0, 0xc9, 0xe1, 0x01, 0x00, 0x00,
}