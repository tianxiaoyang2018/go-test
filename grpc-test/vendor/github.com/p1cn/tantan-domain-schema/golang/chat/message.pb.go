// Code generated by protoc-gen-go. DO NOT EDIT.
// source: chat/message.proto

package chat

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type MessageType int32

const (
	MessageType_TEXT_MESSAGE     MessageType = 0
	MessageType_VIDEO_MESSAGE    MessageType = 1
	MessageType_AUDIO_MESSAGE    MessageType = 2
	MessageType_PICTURE_MESSAGE  MessageType = 3
	MessageType_LOCATION_MESSAGE MessageType = 4
	MessageType_STICKER_MESSAGE  MessageType = 5
	MessageType_QUESTION_MESSAGE MessageType = 6
	MessageType_ANSWER_MESSAGE   MessageType = 7
	MessageType_COMMENT_MESSAGE  MessageType = 8
	MessageType_LIKE_MESSAGE     MessageType = 9
)

var MessageType_name = map[int32]string{
	0: "TEXT_MESSAGE",
	1: "VIDEO_MESSAGE",
	2: "AUDIO_MESSAGE",
	3: "PICTURE_MESSAGE",
	4: "LOCATION_MESSAGE",
	5: "STICKER_MESSAGE",
	6: "QUESTION_MESSAGE",
	7: "ANSWER_MESSAGE",
	8: "COMMENT_MESSAGE",
	9: "LIKE_MESSAGE",
}
var MessageType_value = map[string]int32{
	"TEXT_MESSAGE":     0,
	"VIDEO_MESSAGE":    1,
	"AUDIO_MESSAGE":    2,
	"PICTURE_MESSAGE":  3,
	"LOCATION_MESSAGE": 4,
	"STICKER_MESSAGE":  5,
	"QUESTION_MESSAGE": 6,
	"ANSWER_MESSAGE":   7,
	"COMMENT_MESSAGE":  8,
	"LIKE_MESSAGE":     9,
}

func (x MessageType) String() string {
	return proto.EnumName(MessageType_name, int32(x))
}
func (MessageType) EnumDescriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

type MessageStatus int32

const (
	MessageStatus_MESSAGE_STATUS_DEFAULT     MessageStatus = 0
	MessageStatus_MESSAGE_STATUS_DELETED     MessageStatus = 1
	MessageStatus_MESSAGE_STATUS_INACTIVATED MessageStatus = 2
)

var MessageStatus_name = map[int32]string{
	0: "MESSAGE_STATUS_DEFAULT",
	1: "MESSAGE_STATUS_DELETED",
	2: "MESSAGE_STATUS_INACTIVATED",
}
var MessageStatus_value = map[string]int32{
	"MESSAGE_STATUS_DEFAULT":     0,
	"MESSAGE_STATUS_DELETED":     1,
	"MESSAGE_STATUS_INACTIVATED": 2,
}

func (x MessageStatus) String() string {
	return proto.EnumName(MessageStatus_name, int32(x))
}
func (MessageStatus) EnumDescriptor() ([]byte, []int) { return fileDescriptor5, []int{1} }

type Message struct {
	Id          string                `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	UserId      string                `protobuf:"bytes,2,opt,name=userId" json:"userId,omitempty"`
	OtherUserId string                `protobuf:"bytes,3,opt,name=otherUserId" json:"otherUserId,omitempty"`
	Type        MessageType           `protobuf:"varint,4,opt,name=type,enum=chat.MessageType" json:"type,omitempty"`
	Recalled    bool                  `protobuf:"varint,5,opt,name=recalled" json:"recalled,omitempty"`
	SentFrom    string                `protobuf:"bytes,6,opt,name=sentFrom" json:"sentFrom,omitempty"`
	Status      MessageStatus         `protobuf:"varint,7,opt,name=status,enum=chat.MessageStatus" json:"status,omitempty"`
	CreatedTime int64                 `protobuf:"varint,8,opt,name=createdTime" json:"createdTime,omitempty"`
	UpdatedTime int64                 `protobuf:"varint,9,opt,name=updatedTime" json:"updatedTime,omitempty"`
	Text        *TextMessage          `protobuf:"bytes,10,opt,name=text" json:"text,omitempty"`
	Picture     *PictureMessage       `protobuf:"bytes,11,opt,name=picture" json:"picture,omitempty"`
	Video       *VideoMessage         `protobuf:"bytes,12,opt,name=video" json:"video,omitempty"`
	Audio       *AudioMessage         `protobuf:"bytes,13,opt,name=audio" json:"audio,omitempty"`
	Location    *LocationMessage      `protobuf:"bytes,14,opt,name=location" json:"location,omitempty"`
	Sticker     *StickerMessage       `protobuf:"bytes,15,opt,name=sticker" json:"sticker,omitempty"`
	Question    *QuestionMessage      `protobuf:"bytes,16,opt,name=question" json:"question,omitempty"`
	Answer      *AnswerMessage        `protobuf:"bytes,17,opt,name=answer" json:"answer,omitempty"`
	Comment     *MomentCommentMessage `protobuf:"bytes,18,opt,name=comment" json:"comment,omitempty"`
	MomentLike  *MomentLikeMessage    `protobuf:"bytes,19,opt,name=momentLike" json:"momentLike,omitempty"`
}

func (m *Message) Reset()                    { *m = Message{} }
func (m *Message) String() string            { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()               {}
func (*Message) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

func (m *Message) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Message) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *Message) GetOtherUserId() string {
	if m != nil {
		return m.OtherUserId
	}
	return ""
}

func (m *Message) GetType() MessageType {
	if m != nil {
		return m.Type
	}
	return MessageType_TEXT_MESSAGE
}

func (m *Message) GetRecalled() bool {
	if m != nil {
		return m.Recalled
	}
	return false
}

func (m *Message) GetSentFrom() string {
	if m != nil {
		return m.SentFrom
	}
	return ""
}

func (m *Message) GetStatus() MessageStatus {
	if m != nil {
		return m.Status
	}
	return MessageStatus_MESSAGE_STATUS_DEFAULT
}

func (m *Message) GetCreatedTime() int64 {
	if m != nil {
		return m.CreatedTime
	}
	return 0
}

func (m *Message) GetUpdatedTime() int64 {
	if m != nil {
		return m.UpdatedTime
	}
	return 0
}

func (m *Message) GetText() *TextMessage {
	if m != nil {
		return m.Text
	}
	return nil
}

func (m *Message) GetPicture() *PictureMessage {
	if m != nil {
		return m.Picture
	}
	return nil
}

func (m *Message) GetVideo() *VideoMessage {
	if m != nil {
		return m.Video
	}
	return nil
}

func (m *Message) GetAudio() *AudioMessage {
	if m != nil {
		return m.Audio
	}
	return nil
}

func (m *Message) GetLocation() *LocationMessage {
	if m != nil {
		return m.Location
	}
	return nil
}

func (m *Message) GetSticker() *StickerMessage {
	if m != nil {
		return m.Sticker
	}
	return nil
}

func (m *Message) GetQuestion() *QuestionMessage {
	if m != nil {
		return m.Question
	}
	return nil
}

func (m *Message) GetAnswer() *AnswerMessage {
	if m != nil {
		return m.Answer
	}
	return nil
}

func (m *Message) GetComment() *MomentCommentMessage {
	if m != nil {
		return m.Comment
	}
	return nil
}

func (m *Message) GetMomentLike() *MomentLikeMessage {
	if m != nil {
		return m.MomentLike
	}
	return nil
}

type LocationMessage struct {
	Name        string    `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Address     string    `protobuf:"bytes,2,opt,name=address" json:"address,omitempty"`
	Phone       string    `protobuf:"bytes,3,opt,name=phone" json:"phone,omitempty"`
	Coordinates []float64 `protobuf:"fixed64,4,rep,packed,name=coordinates" json:"coordinates,omitempty"`
}

func (m *LocationMessage) Reset()                    { *m = LocationMessage{} }
func (m *LocationMessage) String() string            { return proto.CompactTextString(m) }
func (*LocationMessage) ProtoMessage()               {}
func (*LocationMessage) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{1} }

func (m *LocationMessage) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *LocationMessage) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *LocationMessage) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *LocationMessage) GetCoordinates() []float64 {
	if m != nil {
		return m.Coordinates
	}
	return nil
}

type StickerMessage struct {
	StickerId string `protobuf:"bytes,1,opt,name=stickerId" json:"stickerId,omitempty"`
}

func (m *StickerMessage) Reset()                    { *m = StickerMessage{} }
func (m *StickerMessage) String() string            { return proto.CompactTextString(m) }
func (*StickerMessage) ProtoMessage()               {}
func (*StickerMessage) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{2} }

func (m *StickerMessage) GetStickerId() string {
	if m != nil {
		return m.StickerId
	}
	return ""
}

type QuestionMessage struct {
	QuestionId string `protobuf:"bytes,1,opt,name=questionId" json:"questionId,omitempty"`
}

func (m *QuestionMessage) Reset()                    { *m = QuestionMessage{} }
func (m *QuestionMessage) String() string            { return proto.CompactTextString(m) }
func (*QuestionMessage) ProtoMessage()               {}
func (*QuestionMessage) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{3} }

func (m *QuestionMessage) GetQuestionId() string {
	if m != nil {
		return m.QuestionId
	}
	return ""
}

type AnswerMessage struct {
	ReferenceId string   `protobuf:"bytes,1,opt,name=referenceId" json:"referenceId,omitempty"`
	StickerId   string   `protobuf:"bytes,2,opt,name=stickerId" json:"stickerId,omitempty"`
	Value       string   `protobuf:"bytes,3,opt,name=value" json:"value,omitempty"`
	Media       []*Media `protobuf:"bytes,4,rep,name=media" json:"media,omitempty"`
}

func (m *AnswerMessage) Reset()                    { *m = AnswerMessage{} }
func (m *AnswerMessage) String() string            { return proto.CompactTextString(m) }
func (*AnswerMessage) ProtoMessage()               {}
func (*AnswerMessage) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{4} }

func (m *AnswerMessage) GetReferenceId() string {
	if m != nil {
		return m.ReferenceId
	}
	return ""
}

func (m *AnswerMessage) GetStickerId() string {
	if m != nil {
		return m.StickerId
	}
	return ""
}

func (m *AnswerMessage) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *AnswerMessage) GetMedia() []*Media {
	if m != nil {
		return m.Media
	}
	return nil
}

type MomentCommentMessage struct {
	MomentId string `protobuf:"bytes,1,opt,name=momentId" json:"momentId,omitempty"`
	Value    string `protobuf:"bytes,3,opt,name=value" json:"value,omitempty"`
}

func (m *MomentCommentMessage) Reset()                    { *m = MomentCommentMessage{} }
func (m *MomentCommentMessage) String() string            { return proto.CompactTextString(m) }
func (*MomentCommentMessage) ProtoMessage()               {}
func (*MomentCommentMessage) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{5} }

func (m *MomentCommentMessage) GetMomentId() string {
	if m != nil {
		return m.MomentId
	}
	return ""
}

func (m *MomentCommentMessage) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type MomentLikeMessage struct {
	MomentId string `protobuf:"bytes,1,opt,name=momentId" json:"momentId,omitempty"`
}

func (m *MomentLikeMessage) Reset()                    { *m = MomentLikeMessage{} }
func (m *MomentLikeMessage) String() string            { return proto.CompactTextString(m) }
func (*MomentLikeMessage) ProtoMessage()               {}
func (*MomentLikeMessage) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{6} }

func (m *MomentLikeMessage) GetMomentId() string {
	if m != nil {
		return m.MomentId
	}
	return ""
}

type TextMessage struct {
	Value string `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
}

func (m *TextMessage) Reset()                    { *m = TextMessage{} }
func (m *TextMessage) String() string            { return proto.CompactTextString(m) }
func (*TextMessage) ProtoMessage()               {}
func (*TextMessage) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{7} }

func (m *TextMessage) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type VideoMessage struct {
	Media []*Media `protobuf:"bytes,11,rep,name=media" json:"media,omitempty"`
}

func (m *VideoMessage) Reset()                    { *m = VideoMessage{} }
func (m *VideoMessage) String() string            { return proto.CompactTextString(m) }
func (*VideoMessage) ProtoMessage()               {}
func (*VideoMessage) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{8} }

func (m *VideoMessage) GetMedia() []*Media {
	if m != nil {
		return m.Media
	}
	return nil
}

type AudioMessage struct {
	Media []*Media `protobuf:"bytes,11,rep,name=media" json:"media,omitempty"`
}

func (m *AudioMessage) Reset()                    { *m = AudioMessage{} }
func (m *AudioMessage) String() string            { return proto.CompactTextString(m) }
func (*AudioMessage) ProtoMessage()               {}
func (*AudioMessage) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{9} }

func (m *AudioMessage) GetMedia() []*Media {
	if m != nil {
		return m.Media
	}
	return nil
}

type PictureMessage struct {
	Media []*Media `protobuf:"bytes,11,rep,name=media" json:"media,omitempty"`
}

func (m *PictureMessage) Reset()                    { *m = PictureMessage{} }
func (m *PictureMessage) String() string            { return proto.CompactTextString(m) }
func (*PictureMessage) ProtoMessage()               {}
func (*PictureMessage) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{10} }

func (m *PictureMessage) GetMedia() []*Media {
	if m != nil {
		return m.Media
	}
	return nil
}

func init() {
	proto.RegisterType((*Message)(nil), "chat.Message")
	proto.RegisterType((*LocationMessage)(nil), "chat.LocationMessage")
	proto.RegisterType((*StickerMessage)(nil), "chat.StickerMessage")
	proto.RegisterType((*QuestionMessage)(nil), "chat.QuestionMessage")
	proto.RegisterType((*AnswerMessage)(nil), "chat.AnswerMessage")
	proto.RegisterType((*MomentCommentMessage)(nil), "chat.MomentCommentMessage")
	proto.RegisterType((*MomentLikeMessage)(nil), "chat.MomentLikeMessage")
	proto.RegisterType((*TextMessage)(nil), "chat.TextMessage")
	proto.RegisterType((*VideoMessage)(nil), "chat.VideoMessage")
	proto.RegisterType((*AudioMessage)(nil), "chat.AudioMessage")
	proto.RegisterType((*PictureMessage)(nil), "chat.PictureMessage")
	proto.RegisterEnum("chat.MessageType", MessageType_name, MessageType_value)
	proto.RegisterEnum("chat.MessageStatus", MessageStatus_name, MessageStatus_value)
}

func init() { proto.RegisterFile("chat/message.proto", fileDescriptor5) }

var fileDescriptor5 = []byte{
	// 830 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x95, 0xdf, 0x6f, 0xdb, 0x36,
	0x10, 0xc7, 0x2b, 0xc7, 0xbf, 0x72, 0x4e, 0x1c, 0x87, 0xc9, 0x3a, 0xc2, 0x18, 0x0a, 0x4f, 0xc3,
	0x00, 0xa3, 0x43, 0xed, 0x39, 0x1d, 0xb0, 0x67, 0xcd, 0x51, 0x37, 0xa1, 0x4e, 0xd2, 0x4a, 0x72,
	0x36, 0xec, 0x25, 0x60, 0x25, 0xce, 0x16, 0x6a, 0x89, 0x9e, 0x44, 0xf5, 0xc7, 0xdb, 0x1e, 0xf7,
	0x17, 0xee, 0xef, 0x19, 0x48, 0x51, 0x34, 0x9d, 0x65, 0x45, 0xdf, 0xc2, 0xef, 0xf7, 0x73, 0xba,
	0x3b, 0x5e, 0x8e, 0x06, 0x14, 0xad, 0x09, 0x9f, 0xa6, 0xb4, 0x28, 0xc8, 0x8a, 0x4e, 0xb6, 0x39,
	0xe3, 0x0c, 0x35, 0x85, 0x36, 0x1c, 0x28, 0x27, 0x4e, 0x48, 0xa5, 0xdb, 0x7f, 0xb5, 0xa1, 0x73,
	0x55, 0x91, 0xa8, 0x0f, 0x8d, 0x24, 0xc6, 0xd6, 0xc8, 0x1a, 0x1f, 0xfa, 0x8d, 0x24, 0x46, 0x8f,
	0xa1, 0x5d, 0x16, 0x34, 0xf7, 0x62, 0xdc, 0x90, 0x9a, 0x3a, 0xa1, 0x11, 0xf4, 0x18, 0x5f, 0xd3,
	0x7c, 0x59, 0x99, 0x07, 0xd2, 0x34, 0x25, 0xf4, 0x2d, 0x34, 0xf9, 0xc7, 0x2d, 0xc5, 0xcd, 0x91,
	0x35, 0xee, 0x5f, 0x9c, 0x4e, 0x44, 0xda, 0x89, 0x4a, 0x13, 0x7e, 0xdc, 0x52, 0x5f, 0xda, 0x68,
	0x08, 0xdd, 0x9c, 0x46, 0x64, 0xb3, 0xa1, 0x31, 0x6e, 0x8d, 0xac, 0x71, 0xd7, 0xd7, 0x67, 0xe1,
	0x15, 0x34, 0xe3, 0x2f, 0x72, 0x96, 0xe2, 0xb6, 0xcc, 0xa0, 0xcf, 0xe8, 0x3b, 0x68, 0x17, 0x9c,
	0xf0, 0xb2, 0xc0, 0x1d, 0x99, 0xe0, 0x6c, 0x2f, 0x41, 0x20, 0x2d, 0x5f, 0x21, 0xa2, 0xda, 0x28,
	0xa7, 0x84, 0xd3, 0x38, 0x4c, 0x52, 0x8a, 0xbb, 0x23, 0x6b, 0x7c, 0xe0, 0x9b, 0x92, 0x20, 0xca,
	0x6d, 0xac, 0x89, 0xc3, 0x8a, 0x30, 0x24, 0xd9, 0x0f, 0xfd, 0xc0, 0x31, 0x8c, 0xac, 0x71, 0xaf,
	0xee, 0x27, 0xa4, 0x1f, 0xb8, 0x4a, 0xe9, 0x4b, 0x1b, 0x4d, 0xa0, 0xb3, 0x4d, 0x22, 0x5e, 0xe6,
	0x14, 0xf7, 0x24, 0x79, 0x5e, 0x91, 0xaf, 0x2a, 0xb1, 0x86, 0x6b, 0x08, 0x8d, 0xa1, 0xf5, 0x2e,
	0x89, 0x29, 0xc3, 0x47, 0x92, 0x46, 0x15, 0x7d, 0x2b, 0xa4, 0x9a, 0xad, 0x00, 0x41, 0x92, 0x32,
	0x4e, 0x18, 0x3e, 0x36, 0x49, 0x47, 0x48, 0x9a, 0x94, 0x00, 0x9a, 0x41, 0x77, 0xc3, 0x22, 0xc2,
	0x13, 0x96, 0xe1, 0xbe, 0x84, 0xbf, 0xa8, 0xe0, 0x85, 0x52, 0x6b, 0x5e, 0x63, 0xa2, 0xec, 0x82,
	0x27, 0xd1, 0x5b, 0x9a, 0xe3, 0x13, 0xb3, 0xec, 0xa0, 0x12, 0x75, 0xd9, 0x0a, 0x12, 0x29, 0xfe,
	0x2c, 0x69, 0x21, 0x53, 0x0c, 0xcc, 0x14, 0xaf, 0x95, 0xaa, 0x53, 0xd4, 0x98, 0x98, 0x18, 0xc9,
	0x8a, 0xf7, 0x34, 0xc7, 0xa7, 0x32, 0x40, 0x4d, 0xcc, 0x91, 0x5a, 0x8d, 0x2b, 0x04, 0xfd, 0x00,
	0x9d, 0x88, 0xa5, 0x29, 0xcd, 0x38, 0x46, 0x92, 0x1e, 0xaa, 0xf9, 0x32, 0xa1, 0xcd, 0x2b, 0x4b,
	0x57, 0xa5, 0x50, 0xf4, 0x23, 0x40, 0x2a, 0x81, 0x45, 0xf2, 0x96, 0xe2, 0x33, 0x19, 0xf8, 0xa5,
	0x19, 0x28, 0xf4, 0x3a, 0xca, 0x40, 0xed, 0xf7, 0x70, 0x72, 0xef, 0x6e, 0x10, 0x82, 0x66, 0x46,
	0x52, 0xaa, 0x76, 0x41, 0xfe, 0x8d, 0x30, 0x74, 0x48, 0x1c, 0xe7, 0xb4, 0x28, 0xd4, 0x3a, 0xd4,
	0x47, 0x74, 0x0e, 0xad, 0xed, 0x9a, 0x65, 0x54, 0x6d, 0x42, 0x75, 0x90, 0xff, 0x77, 0x8c, 0xe5,
	0x71, 0x92, 0x11, 0x4e, 0x0b, 0xdc, 0x1c, 0x1d, 0x8c, 0x2d, 0xdf, 0x94, 0xec, 0x09, 0xf4, 0xf7,
	0xaf, 0x18, 0x7d, 0x05, 0x87, 0xea, 0x92, 0xbd, 0x7a, 0x11, 0x77, 0x82, 0x3d, 0x83, 0x93, 0x7b,
	0x37, 0x8c, 0x9e, 0x00, 0xd4, 0x77, 0xac, 0x23, 0x0c, 0xc5, 0xfe, 0xdb, 0x82, 0xe3, 0xbd, 0x4b,
	0x16, 0x65, 0xe5, 0xf4, 0x0f, 0x9a, 0xd3, 0x2c, 0xa2, 0x3a, 0xc4, 0x94, 0xf6, 0x8b, 0x68, 0xdc,
	0x2b, 0x42, 0x34, 0xfb, 0x8e, 0x6c, 0x4a, 0xdd, 0xac, 0x3c, 0xa0, 0xaf, 0xa1, 0x25, 0x5f, 0x15,
	0xd9, 0x66, 0xef, 0xa2, 0x57, 0x2f, 0x64, 0x9c, 0x10, 0xbf, 0x72, 0xec, 0x5f, 0xe0, 0xfc, 0xa1,
	0x01, 0x8a, 0x45, 0xaf, 0x86, 0xa1, 0xab, 0xd1, 0xe7, 0x87, 0x93, 0xd9, 0x53, 0x38, 0xfd, 0xcf,
	0x44, 0x3f, 0xf5, 0x19, 0xfb, 0x1b, 0xe8, 0x19, 0xcb, 0xba, 0xfb, 0xaa, 0x65, 0x7e, 0x75, 0x06,
	0x47, 0xe6, 0xe6, 0xed, 0x5a, 0xea, 0xfd, 0x6f, 0x4b, 0x33, 0x38, 0x32, 0x57, 0xf0, 0x73, 0x42,
	0x9e, 0x43, 0x7f, 0xff, 0x35, 0xf8, 0x8c, 0xa0, 0xa7, 0xff, 0x58, 0xd0, 0x33, 0x5e, 0x4f, 0x34,
	0x80, 0xa3, 0xd0, 0xfd, 0x2d, 0xbc, 0xbb, 0x72, 0x83, 0xc0, 0xf9, 0xd9, 0x1d, 0x3c, 0x42, 0xa7,
	0x70, 0x7c, 0xeb, 0x5d, 0xba, 0x37, 0x5a, 0xb2, 0x84, 0xe4, 0x2c, 0x2f, 0xbd, 0x9d, 0xd4, 0x40,
	0x67, 0x70, 0xf2, 0xca, 0x9b, 0x87, 0x4b, 0xdf, 0xd5, 0xe2, 0x01, 0x3a, 0x87, 0xc1, 0xe2, 0x66,
	0xee, 0x84, 0xde, 0xcd, 0xb5, 0x56, 0x9b, 0x02, 0x0d, 0x42, 0x6f, 0xfe, 0xd2, 0xf5, 0xb5, 0xd8,
	0x12, 0xe8, 0xeb, 0xa5, 0x1b, 0xec, 0xa1, 0x6d, 0x84, 0xa0, 0xef, 0x5c, 0x07, 0xbf, 0x1a, 0x64,
	0x47, 0x84, 0xcf, 0x6f, 0xae, 0xae, 0xdc, 0xeb, 0x5d, 0x91, 0x5d, 0x51, 0xf6, 0xc2, 0x7b, 0xb9,
	0xcb, 0x7d, 0xf8, 0x74, 0x05, 0xc7, 0x7b, 0x8f, 0x36, 0x1a, 0xc2, 0x63, 0xe5, 0xde, 0x05, 0xa1,
	0x13, 0x2e, 0x83, 0xbb, 0x4b, 0xf7, 0x85, 0xb3, 0x5c, 0x84, 0x83, 0x47, 0x0f, 0x7a, 0x0b, 0x37,
	0x74, 0x2f, 0x07, 0x16, 0x7a, 0x02, 0xc3, 0x7b, 0x9e, 0x77, 0xed, 0xcc, 0x43, 0xef, 0xd6, 0x11,
	0x7e, 0xe3, 0xa7, 0x8b, 0xdf, 0xbf, 0x5f, 0x25, 0x7c, 0x5d, 0xbe, 0x99, 0x44, 0x2c, 0x9d, 0x6e,
	0x67, 0x51, 0x36, 0xe5, 0x24, 0xe3, 0x24, 0x7b, 0x16, 0xb3, 0x94, 0x24, 0xd9, 0xb3, 0x22, 0x5a,
	0xd3, 0x94, 0x4c, 0x57, 0x6c, 0x43, 0xb2, 0xd5, 0x54, 0x8c, 0xe0, 0x4d, 0x5b, 0xfe, 0x42, 0x3e,
	0xff, 0x37, 0x00, 0x00, 0xff, 0xff, 0xc1, 0xcc, 0xeb, 0x82, 0x4f, 0x07, 0x00, 0x00,
}