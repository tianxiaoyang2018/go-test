// Code generated by protoc-gen-go. DO NOT EDIT.
// source: chat/question.proto

package chat

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Question struct {
	Id          string    `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Text        string    `protobuf:"bytes,2,opt,name=text" json:"text,omitempty"`
	Answers     []*Answer `protobuf:"bytes,3,rep,name=answers" json:"answers,omitempty"`
	Category    string    `protobuf:"bytes,4,opt,name=category" json:"category,omitempty"`
	Usage       string    `protobuf:"bytes,5,opt,name=usage" json:"usage,omitempty"`
	Status      string    `protobuf:"bytes,6,opt,name=status" json:"status,omitempty"`
	Sensitivity int32     `protobuf:"varint,7,opt,name=sensitivity" json:"sensitivity,omitempty"`
	Priority    int32     `protobuf:"varint,8,opt,name=priority" json:"priority,omitempty"`
	Type        string    `protobuf:"bytes,9,opt,name=type" json:"type,omitempty"`
}

func (m *Question) Reset()                    { *m = Question{} }
func (m *Question) String() string            { return proto.CompactTextString(m) }
func (*Question) ProtoMessage()               {}
func (*Question) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{0} }

func (m *Question) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Question) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *Question) GetAnswers() []*Answer {
	if m != nil {
		return m.Answers
	}
	return nil
}

func (m *Question) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

func (m *Question) GetUsage() string {
	if m != nil {
		return m.Usage
	}
	return ""
}

func (m *Question) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Question) GetSensitivity() int32 {
	if m != nil {
		return m.Sensitivity
	}
	return 0
}

func (m *Question) GetPriority() int32 {
	if m != nil {
		return m.Priority
	}
	return 0
}

func (m *Question) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

type Answer struct {
	Id         string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	QuestionId string `protobuf:"bytes,2,opt,name=questionId" json:"questionId,omitempty"`
	Attitude   string `protobuf:"bytes,3,opt,name=attitude" json:"attitude,omitempty"`
	Priority   int32  `protobuf:"varint,4,opt,name=priority" json:"priority,omitempty"`
	Text       string `protobuf:"bytes,5,opt,name=text" json:"text,omitempty"`
	Type       string `protobuf:"bytes,6,opt,name=type" json:"type,omitempty"`
}

func (m *Answer) Reset()                    { *m = Answer{} }
func (m *Answer) String() string            { return proto.CompactTextString(m) }
func (*Answer) ProtoMessage()               {}
func (*Answer) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{1} }

func (m *Answer) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Answer) GetQuestionId() string {
	if m != nil {
		return m.QuestionId
	}
	return ""
}

func (m *Answer) GetAttitude() string {
	if m != nil {
		return m.Attitude
	}
	return ""
}

func (m *Answer) GetPriority() int32 {
	if m != nil {
		return m.Priority
	}
	return 0
}

func (m *Answer) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *Answer) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

type AnswerReply struct {
	AnswerId                        string `protobuf:"bytes,1,opt,name=answerId" json:"answerId,omitempty"`
	QuestionId                      string `protobuf:"bytes,2,opt,name=questionId" json:"questionId,omitempty"`
	PredefinedMessageTranslationKey string `protobuf:"bytes,3,opt,name=predefinedMessageTranslationKey" json:"predefinedMessageTranslationKey,omitempty"`
	Type                            string `protobuf:"bytes,4,opt,name=type" json:"type,omitempty"`
}

func (m *AnswerReply) Reset()                    { *m = AnswerReply{} }
func (m *AnswerReply) String() string            { return proto.CompactTextString(m) }
func (*AnswerReply) ProtoMessage()               {}
func (*AnswerReply) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{2} }

func (m *AnswerReply) GetAnswerId() string {
	if m != nil {
		return m.AnswerId
	}
	return ""
}

func (m *AnswerReply) GetQuestionId() string {
	if m != nil {
		return m.QuestionId
	}
	return ""
}

func (m *AnswerReply) GetPredefinedMessageTranslationKey() string {
	if m != nil {
		return m.PredefinedMessageTranslationKey
	}
	return ""
}

func (m *AnswerReply) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func init() {
	proto.RegisterType((*Question)(nil), "chat.Question")
	proto.RegisterType((*Answer)(nil), "chat.Answer")
	proto.RegisterType((*AnswerReply)(nil), "chat.AnswerReply")
}

func init() { proto.RegisterFile("chat/question.proto", fileDescriptor7) }

var fileDescriptor7 = []byte{
	// 361 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x4f, 0x4e, 0xe3, 0x30,
	0x18, 0xc5, 0x95, 0x36, 0x49, 0x5b, 0x77, 0x34, 0x0b, 0xcf, 0x68, 0x64, 0xcd, 0x02, 0xa2, 0x2e,
	0x50, 0x37, 0x4d, 0xa0, 0x9c, 0x00, 0x56, 0x54, 0x88, 0x05, 0x15, 0x2b, 0x76, 0x6e, 0xf2, 0x91,
	0x5a, 0x6a, 0xed, 0x60, 0x7f, 0x01, 0x72, 0x13, 0x6e, 0xc0, 0xfd, 0x38, 0x01, 0xb2, 0xf3, 0x87,
	0x20, 0x90, 0xba, 0xeb, 0x7b, 0xcf, 0x8d, 0x7f, 0xef, 0xc9, 0xe4, 0x4f, 0xba, 0xe5, 0x98, 0x3c,
	0x96, 0x60, 0x50, 0x28, 0x19, 0x17, 0x5a, 0xa1, 0xa2, 0xbe, 0x35, 0x67, 0xef, 0x1e, 0x19, 0xdf,
	0x36, 0x01, 0xfd, 0x4d, 0x06, 0x22, 0x63, 0x5e, 0xe4, 0xcd, 0x27, 0xeb, 0x81, 0xc8, 0x28, 0x25,
	0x3e, 0xc2, 0x0b, 0xb2, 0x81, 0x73, 0xdc, 0x6f, 0x7a, 0x42, 0x46, 0x5c, 0x9a, 0x67, 0xd0, 0x86,
	0x0d, 0xa3, 0xe1, 0x7c, 0xba, 0xfc, 0x15, 0xdb, 0x0f, 0xc5, 0x17, 0xce, 0x5c, 0xb7, 0x21, 0xfd,
	0x4f, 0xc6, 0x29, 0x47, 0xc8, 0x95, 0xae, 0x98, 0xef, 0xfe, 0xdf, 0x69, 0xfa, 0x97, 0x04, 0xa5,
	0xe1, 0x39, 0xb0, 0xc0, 0x05, 0xb5, 0xa0, 0xff, 0x48, 0x68, 0x90, 0x63, 0x69, 0x58, 0xe8, 0xec,
	0x46, 0xd1, 0x88, 0x4c, 0x0d, 0x48, 0x23, 0x50, 0x3c, 0x09, 0xac, 0xd8, 0x28, 0xf2, 0xe6, 0xc1,
	0xba, 0x6f, 0xd9, 0xbb, 0x0a, 0x2d, 0x94, 0xb6, 0xf1, 0xd8, 0xc5, 0x9d, 0x76, 0x1d, 0xaa, 0x02,
	0xd8, 0xa4, 0xe9, 0x50, 0x15, 0x30, 0x7b, 0xf5, 0x48, 0x58, 0xf3, 0x7e, 0xab, 0x7c, 0x44, 0x48,
	0xbb, 0xd3, 0x2a, 0x6b, 0x8a, 0xf7, 0x1c, 0x7b, 0x15, 0x47, 0x14, 0x58, 0x66, 0xc0, 0x86, 0x75,
	0xad, 0x56, 0x7f, 0xc1, 0xf0, 0x7f, 0xc0, 0xb0, 0x53, 0x06, 0xbd, 0x29, 0x5b, 0xb4, 0xb0, 0x87,
	0xf6, 0xe6, 0x91, 0x69, 0x33, 0x25, 0x14, 0x3b, 0x57, 0xad, 0x5e, 0x74, 0xd5, 0x52, 0x76, 0xfa,
	0x20, 0xeb, 0x15, 0x39, 0x2e, 0x34, 0x64, 0xf0, 0x20, 0x24, 0x64, 0x37, 0x60, 0xec, 0xca, 0x77,
	0x9a, 0x4b, 0xb3, 0xe3, 0xf6, 0xc4, 0x35, 0x54, 0x4d, 0x85, 0x43, 0xc7, 0x3a, 0x52, 0xff, 0x93,
	0xf4, 0x72, 0x79, 0x7f, 0x9a, 0x0b, 0xdc, 0x96, 0x9b, 0x38, 0x55, 0xfb, 0xa4, 0x38, 0x4b, 0x65,
	0x82, 0x5c, 0x22, 0x97, 0x8b, 0x4c, 0xed, 0xb9, 0x90, 0x0b, 0x93, 0x6e, 0x61, 0xcf, 0x93, 0x5c,
	0xed, 0xb8, 0xcc, 0x13, 0xfb, 0x48, 0x36, 0xa1, 0x7b, 0x7a, 0xe7, 0x1f, 0x01, 0x00, 0x00, 0xff,
	0xff, 0xf0, 0xbd, 0x42, 0xcd, 0x91, 0x02, 0x00, 0x00,
}
