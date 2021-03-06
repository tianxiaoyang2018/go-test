// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user/user-counter.proto

package user

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type UserCounter struct {
	UserId              string  `protobuf:"bytes,1,opt,name=userId" json:"userId,omitempty"`
	Gender              string  `protobuf:"bytes,2,opt,name=gender" json:"gender,omitempty"`
	ReceivedLikes       int64   `protobuf:"varint,3,opt,name=receivedLikes" json:"receivedLikes,omitempty"`
	ReceivedDislikes    int64   `protobuf:"varint,4,opt,name=receivedDislikes" json:"receivedDislikes,omitempty"`
	ReceivedBlocks      int64   `protobuf:"varint,5,opt,name=receivedBlocks" json:"receivedBlocks,omitempty"`
	GivenLikes          int64   `protobuf:"varint,6,opt,name=givenLikes" json:"givenLikes,omitempty"`
	GivenDislikes       int64   `protobuf:"varint,7,opt,name=givenDislikes" json:"givenDislikes,omitempty"`
	GivenBlocks         int64   `protobuf:"varint,8,opt,name=givenBlocks" json:"givenBlocks,omitempty"`
	Matches             int64   `protobuf:"varint,9,opt,name=matches" json:"matches,omitempty"`
	LikeRating          int64   `protobuf:"varint,10,opt,name=likeRating" json:"likeRating,omitempty"`
	Popularity          float64 `protobuf:"fixed64,11,opt,name=popularity" json:"popularity,omitempty"`
	TotalCrushes        int64   `protobuf:"varint,12,opt,name=totalCrushes" json:"totalCrushes,omitempty"`
	GivenCrushes        int64   `protobuf:"varint,13,opt,name=givenCrushes" json:"givenCrushes,omitempty"`
	ReceivedCrushes     int64   `protobuf:"varint,14,opt,name=receivedCrushes" json:"receivedCrushes,omitempty"`
	Moments             int64   `protobuf:"varint,15,opt,name=moments" json:"moments,omitempty"`
	UnreadActivities    int64   `protobuf:"varint,16,opt,name=unreadActivities" json:"unreadActivities,omitempty"`
	GivenSuperLikes     int64   `protobuf:"varint,17,opt,name=givenSuperLikes" json:"givenSuperLikes,omitempty"`
	GivenUndos          int64   `protobuf:"varint,18,opt,name=givenUndos" json:"givenUndos,omitempty"`
	ReceivedSuperLikes  int64   `protobuf:"varint,19,opt,name=receivedSuperLikes" json:"receivedSuperLikes,omitempty"`
	LastResetTime       int64   `protobuf:"varint,20,opt,name=lastResetTime" json:"lastResetTime,omitempty"`
	LastShareTime       int64   `protobuf:"varint,21,opt,name=lastShareTime" json:"lastShareTime,omitempty"`
	LikesWithinLimit    int64   `protobuf:"varint,22,opt,name=likesWithinLimit" json:"likesWithinLimit,omitempty"`
	TotalLikesLimit     int64   `protobuf:"varint,23,opt,name=totalLikesLimit" json:"totalLikesLimit,omitempty"`
	RemainingLikes      int64   `protobuf:"varint,24,opt,name=remainingLikes" json:"remainingLikes,omitempty"`
	SuperLikeReward     int64   `protobuf:"varint,25,opt,name=superLikeReward" json:"superLikeReward,omitempty"`
	SuperLikeBuy        int64   `protobuf:"varint,26,opt,name=superLikeBuy" json:"superLikeBuy,omitempty"`
	SuperLikeRemaining  int64   `protobuf:"varint,27,opt,name=superLikeRemaining" json:"superLikeRemaining,omitempty"`
	SuperLikeCount      int64   `protobuf:"varint,28,opt,name=superLikeCount" json:"superLikeCount,omitempty"`
	SuperLikeQuota      int64   `protobuf:"varint,29,opt,name=superLikeQuota" json:"superLikeQuota,omitempty"`
	SuperLikeDailyQuota int64   `protobuf:"varint,30,opt,name=superLikeDailyQuota" json:"superLikeDailyQuota,omitempty"`
	UndoReward          int64   `protobuf:"varint,31,opt,name=undoReward" json:"undoReward,omitempty"`
	UndoBuy             int64   `protobuf:"varint,32,opt,name=undoBuy" json:"undoBuy,omitempty"`
	UndoRemaining       int64   `protobuf:"varint,33,opt,name=undoRemaining" json:"undoRemaining,omitempty"`
	UndoCount           int64   `protobuf:"varint,34,opt,name=undoCount" json:"undoCount,omitempty"`
	UndoQuota           int64   `protobuf:"varint,35,opt,name=undoQuota" json:"undoQuota,omitempty"`
}

func (m *UserCounter) Reset()                    { *m = UserCounter{} }
func (m *UserCounter) String() string            { return proto.CompactTextString(m) }
func (*UserCounter) ProtoMessage()               {}
func (*UserCounter) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *UserCounter) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *UserCounter) GetGender() string {
	if m != nil {
		return m.Gender
	}
	return ""
}

func (m *UserCounter) GetReceivedLikes() int64 {
	if m != nil {
		return m.ReceivedLikes
	}
	return 0
}

func (m *UserCounter) GetReceivedDislikes() int64 {
	if m != nil {
		return m.ReceivedDislikes
	}
	return 0
}

func (m *UserCounter) GetReceivedBlocks() int64 {
	if m != nil {
		return m.ReceivedBlocks
	}
	return 0
}

func (m *UserCounter) GetGivenLikes() int64 {
	if m != nil {
		return m.GivenLikes
	}
	return 0
}

func (m *UserCounter) GetGivenDislikes() int64 {
	if m != nil {
		return m.GivenDislikes
	}
	return 0
}

func (m *UserCounter) GetGivenBlocks() int64 {
	if m != nil {
		return m.GivenBlocks
	}
	return 0
}

func (m *UserCounter) GetMatches() int64 {
	if m != nil {
		return m.Matches
	}
	return 0
}

func (m *UserCounter) GetLikeRating() int64 {
	if m != nil {
		return m.LikeRating
	}
	return 0
}

func (m *UserCounter) GetPopularity() float64 {
	if m != nil {
		return m.Popularity
	}
	return 0
}

func (m *UserCounter) GetTotalCrushes() int64 {
	if m != nil {
		return m.TotalCrushes
	}
	return 0
}

func (m *UserCounter) GetGivenCrushes() int64 {
	if m != nil {
		return m.GivenCrushes
	}
	return 0
}

func (m *UserCounter) GetReceivedCrushes() int64 {
	if m != nil {
		return m.ReceivedCrushes
	}
	return 0
}

func (m *UserCounter) GetMoments() int64 {
	if m != nil {
		return m.Moments
	}
	return 0
}

func (m *UserCounter) GetUnreadActivities() int64 {
	if m != nil {
		return m.UnreadActivities
	}
	return 0
}

func (m *UserCounter) GetGivenSuperLikes() int64 {
	if m != nil {
		return m.GivenSuperLikes
	}
	return 0
}

func (m *UserCounter) GetGivenUndos() int64 {
	if m != nil {
		return m.GivenUndos
	}
	return 0
}

func (m *UserCounter) GetReceivedSuperLikes() int64 {
	if m != nil {
		return m.ReceivedSuperLikes
	}
	return 0
}

func (m *UserCounter) GetLastResetTime() int64 {
	if m != nil {
		return m.LastResetTime
	}
	return 0
}

func (m *UserCounter) GetLastShareTime() int64 {
	if m != nil {
		return m.LastShareTime
	}
	return 0
}

func (m *UserCounter) GetLikesWithinLimit() int64 {
	if m != nil {
		return m.LikesWithinLimit
	}
	return 0
}

func (m *UserCounter) GetTotalLikesLimit() int64 {
	if m != nil {
		return m.TotalLikesLimit
	}
	return 0
}

func (m *UserCounter) GetRemainingLikes() int64 {
	if m != nil {
		return m.RemainingLikes
	}
	return 0
}

func (m *UserCounter) GetSuperLikeReward() int64 {
	if m != nil {
		return m.SuperLikeReward
	}
	return 0
}

func (m *UserCounter) GetSuperLikeBuy() int64 {
	if m != nil {
		return m.SuperLikeBuy
	}
	return 0
}

func (m *UserCounter) GetSuperLikeRemaining() int64 {
	if m != nil {
		return m.SuperLikeRemaining
	}
	return 0
}

func (m *UserCounter) GetSuperLikeCount() int64 {
	if m != nil {
		return m.SuperLikeCount
	}
	return 0
}

func (m *UserCounter) GetSuperLikeQuota() int64 {
	if m != nil {
		return m.SuperLikeQuota
	}
	return 0
}

func (m *UserCounter) GetSuperLikeDailyQuota() int64 {
	if m != nil {
		return m.SuperLikeDailyQuota
	}
	return 0
}

func (m *UserCounter) GetUndoReward() int64 {
	if m != nil {
		return m.UndoReward
	}
	return 0
}

func (m *UserCounter) GetUndoBuy() int64 {
	if m != nil {
		return m.UndoBuy
	}
	return 0
}

func (m *UserCounter) GetUndoRemaining() int64 {
	if m != nil {
		return m.UndoRemaining
	}
	return 0
}

func (m *UserCounter) GetUndoCount() int64 {
	if m != nil {
		return m.UndoCount
	}
	return 0
}

func (m *UserCounter) GetUndoQuota() int64 {
	if m != nil {
		return m.UndoQuota
	}
	return 0
}

func init() {
	proto.RegisterType((*UserCounter)(nil), "user.UserCounter")
}

func init() { proto.RegisterFile("user/user-counter.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 583 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x94, 0x6f, 0x4f, 0xd4, 0x40,
	0x10, 0xc6, 0x73, 0x82, 0x20, 0xcb, 0x5f, 0x17, 0x85, 0x55, 0x11, 0x4f, 0x34, 0xe6, 0x62, 0xc2,
	0x1d, 0xea, 0x27, 0xf0, 0xe0, 0x8d, 0x09, 0x6f, 0x2c, 0x12, 0x13, 0xdf, 0x2d, 0xed, 0xa6, 0xb7,
	0xa1, 0xdd, 0xbd, 0xec, 0x6e, 0xcf, 0xdc, 0xc7, 0xf0, 0x1b, 0x9b, 0x99, 0xe9, 0xf6, 0xda, 0xc2,
	0x1b, 0x92, 0xf9, 0xcd, 0xc3, 0xcc, 0x3c, 0x3b, 0x73, 0x65, 0xc7, 0x95, 0x57, 0x6e, 0x02, 0x7f,
	0xce, 0x53, 0x5b, 0x99, 0xa0, 0xdc, 0x78, 0xee, 0x6c, 0xb0, 0x7c, 0x1d, 0xd8, 0xd9, 0x3f, 0xc6,
	0xb6, 0x6f, 0xbd, 0x72, 0x97, 0x94, 0xe3, 0x47, 0x6c, 0x03, 0xf8, 0x8f, 0x4c, 0x0c, 0x86, 0x83,
	0xd1, 0x56, 0x52, 0x47, 0xc0, 0x73, 0x65, 0x32, 0xe5, 0xc4, 0x13, 0xe2, 0x14, 0xf1, 0x8f, 0x6c,
	0xd7, 0xa9, 0x54, 0xe9, 0x85, 0xca, 0xae, 0xf5, 0xbd, 0xf2, 0x62, 0x6d, 0x38, 0x18, 0xad, 0x25,
	0x5d, 0xc8, 0x3f, 0xb3, 0x83, 0x08, 0xae, 0xb4, 0x2f, 0x50, 0xb8, 0x8e, 0xc2, 0x07, 0x9c, 0x7f,
	0x62, 0x7b, 0x91, 0x4d, 0x0b, 0x9b, 0xde, 0x7b, 0xf1, 0x14, 0x95, 0x3d, 0xca, 0x4f, 0x19, 0xcb,
	0xf5, 0x42, 0x19, 0x6a, 0xbb, 0x81, 0x9a, 0x16, 0x81, 0xc9, 0x30, 0x6a, 0x1a, 0x6e, 0xd2, 0x64,
	0x1d, 0xc8, 0x87, 0x6c, 0x1b, 0x41, 0xdd, 0xea, 0x19, 0x6a, 0xda, 0x88, 0x0b, 0xb6, 0x59, 0xca,
	0x90, 0xce, 0x94, 0x17, 0x5b, 0x98, 0x8d, 0x21, 0x4c, 0x00, 0x45, 0x12, 0x19, 0xb4, 0xc9, 0x05,
	0xa3, 0x09, 0x56, 0x04, 0xf2, 0x73, 0x3b, 0xaf, 0x0a, 0xe9, 0x74, 0x58, 0x8a, 0xed, 0xe1, 0x60,
	0x34, 0x48, 0x5a, 0x84, 0x9f, 0xb1, 0x9d, 0x60, 0x83, 0x2c, 0x2e, 0x5d, 0xe5, 0xa1, 0xfc, 0x0e,
	0x56, 0xe8, 0x30, 0xd0, 0xe0, 0x30, 0x51, 0xb3, 0x4b, 0x9a, 0x36, 0xe3, 0x23, 0xb6, 0x1f, 0xdf,
	0x26, 0xca, 0xf6, 0x50, 0xd6, 0xc7, 0xe8, 0xc5, 0x96, 0xca, 0x04, 0x2f, 0xf6, 0x6b, 0x2f, 0x14,
	0xc2, 0x86, 0x2a, 0xe3, 0x94, 0xcc, 0xbe, 0xa7, 0x41, 0x2f, 0x74, 0xd0, 0xca, 0x8b, 0x03, 0xda,
	0x50, 0x9f, 0x43, 0x3f, 0xec, 0x7f, 0x53, 0xcd, 0x95, 0xa3, 0xe7, 0x7f, 0x4e, 0xfd, 0x7a, 0xb8,
	0xd9, 0xd1, 0xad, 0xc9, 0xac, 0x17, 0xbc, 0xb5, 0x23, 0x24, 0x7c, 0xcc, 0x78, 0x1c, 0xb1, 0x55,
	0xec, 0x10, 0x75, 0x8f, 0x64, 0x60, 0xa7, 0x85, 0xf4, 0x21, 0x51, 0x5e, 0x85, 0x5f, 0xba, 0x54,
	0xe2, 0x05, 0xed, 0xb4, 0x03, 0xa3, 0xea, 0x66, 0x26, 0x9d, 0x42, 0xd5, 0xcb, 0x95, 0xaa, 0x81,
	0xe0, 0x18, 0x4f, 0xe0, 0xb7, 0x0e, 0x33, 0x6d, 0xae, 0x75, 0xa9, 0x83, 0x38, 0x22, 0xc7, 0x7d,
	0x0e, 0x8e, 0x71, 0x2b, 0x38, 0x05, 0x49, 0x8f, 0xc9, 0x71, 0x0f, 0xd3, 0xf5, 0x96, 0x52, 0x1b,
	0x6d, 0x72, 0x72, 0x23, 0xe2, 0xf5, 0xb6, 0x29, 0x54, 0xf4, 0xd1, 0x57, 0xa2, 0xfe, 0x4a, 0x97,
	0x89, 0x57, 0x54, 0xb1, 0x87, 0xe1, 0x02, 0x1a, 0x34, 0xad, 0x96, 0xe2, 0x35, 0x5d, 0x40, 0x9b,
	0xc1, 0x3b, 0xb6, 0xfe, 0xad, 0x6e, 0x24, 0xde, 0xd0, 0x3b, 0x3e, 0xcc, 0xc0, 0x94, 0x0d, 0xc5,
	0x5f, 0xbe, 0x38, 0xa1, 0x29, 0xbb, 0xb4, 0xa3, 0xfb, 0x59, 0xd9, 0x20, 0xc5, 0xdb, 0x9e, 0x0e,
	0x29, 0xbf, 0x60, 0x87, 0x0d, 0xb9, 0x92, 0xba, 0x58, 0x92, 0xf8, 0x14, 0xc5, 0x8f, 0xa5, 0xe0,
	0x32, 0x2a, 0x93, 0xd9, 0xda, 0xfa, 0x3b, 0xba, 0x8c, 0x15, 0x81, 0x4b, 0x85, 0x08, 0x0c, 0x0f,
	0xe9, 0x52, 0xeb, 0x10, 0xb6, 0x4b, 0xba, 0x68, 0xf3, 0x3d, 0x6d, 0xb7, 0x03, 0xf9, 0x09, 0xdb,
	0x02, 0x40, 0xe6, 0xce, 0x50, 0xb1, 0x02, 0x31, 0x4b, 0x53, 0x7e, 0x58, 0x65, 0x11, 0x4c, 0xbf,
	0xfe, 0xb9, 0xc8, 0x75, 0x98, 0x55, 0x77, 0xe3, 0xd4, 0x96, 0x93, 0xf9, 0x97, 0xd4, 0x4c, 0x82,
	0x34, 0x41, 0x9a, 0xf3, 0xcc, 0x42, 0xfd, 0x73, 0x9f, 0xce, 0x54, 0x29, 0x27, 0xb9, 0x2d, 0xa4,
	0xc9, 0xf1, 0xdb, 0x7a, 0xb7, 0x81, 0x1f, 0xd5, 0x6f, 0xff, 0x03, 0x00, 0x00, 0xff, 0xff, 0x9f,
	0xc6, 0x7b, 0x2a, 0x6f, 0x05, 0x00, 0x00,
}
