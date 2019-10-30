// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/Recognition/face_recognition.proto

package Recognition

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type MatchParameters struct {
	UserId               int64    `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	FlowId               int64    `protobuf:"varint,2,opt,name=flowId,proto3" json:"flowId,omitempty"`
	LicenseId            int64    `protobuf:"varint,3,opt,name=licenseId,proto3" json:"licenseId,omitempty"`
	S3Reference          string   `protobuf:"bytes,4,opt,name=s3Reference,proto3" json:"s3Reference,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MatchParameters) Reset()         { *m = MatchParameters{} }
func (m *MatchParameters) String() string { return proto.CompactTextString(m) }
func (*MatchParameters) ProtoMessage()    {}
func (*MatchParameters) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d0473497cda6150, []int{0}
}

func (m *MatchParameters) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MatchParameters.Unmarshal(m, b)
}
func (m *MatchParameters) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MatchParameters.Marshal(b, m, deterministic)
}
func (m *MatchParameters) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MatchParameters.Merge(m, src)
}
func (m *MatchParameters) XXX_Size() int {
	return xxx_messageInfo_MatchParameters.Size(m)
}
func (m *MatchParameters) XXX_DiscardUnknown() {
	xxx_messageInfo_MatchParameters.DiscardUnknown(m)
}

var xxx_messageInfo_MatchParameters proto.InternalMessageInfo

func (m *MatchParameters) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *MatchParameters) GetFlowId() int64 {
	if m != nil {
		return m.FlowId
	}
	return 0
}

func (m *MatchParameters) GetLicenseId() int64 {
	if m != nil {
		return m.LicenseId
	}
	return 0
}

func (m *MatchParameters) GetS3Reference() string {
	if m != nil {
		return m.S3Reference
	}
	return ""
}

type GetMatchParameters struct {
	MatchId              int64    `protobuf:"varint,1,opt,name=matchId,proto3" json:"matchId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetMatchParameters) Reset()         { *m = GetMatchParameters{} }
func (m *GetMatchParameters) String() string { return proto.CompactTextString(m) }
func (*GetMatchParameters) ProtoMessage()    {}
func (*GetMatchParameters) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d0473497cda6150, []int{1}
}

func (m *GetMatchParameters) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMatchParameters.Unmarshal(m, b)
}
func (m *GetMatchParameters) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMatchParameters.Marshal(b, m, deterministic)
}
func (m *GetMatchParameters) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMatchParameters.Merge(m, src)
}
func (m *GetMatchParameters) XXX_Size() int {
	return xxx_messageInfo_GetMatchParameters.Size(m)
}
func (m *GetMatchParameters) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMatchParameters.DiscardUnknown(m)
}

var xxx_messageInfo_GetMatchParameters proto.InternalMessageInfo

func (m *GetMatchParameters) GetMatchId() int64 {
	if m != nil {
		return m.MatchId
	}
	return 0
}

type GetUserMatchesParameters struct {
	UserId               int64    `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserMatchesParameters) Reset()         { *m = GetUserMatchesParameters{} }
func (m *GetUserMatchesParameters) String() string { return proto.CompactTextString(m) }
func (*GetUserMatchesParameters) ProtoMessage()    {}
func (*GetUserMatchesParameters) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d0473497cda6150, []int{2}
}

func (m *GetUserMatchesParameters) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserMatchesParameters.Unmarshal(m, b)
}
func (m *GetUserMatchesParameters) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserMatchesParameters.Marshal(b, m, deterministic)
}
func (m *GetUserMatchesParameters) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserMatchesParameters.Merge(m, src)
}
func (m *GetUserMatchesParameters) XXX_Size() int {
	return xxx_messageInfo_GetUserMatchesParameters.Size(m)
}
func (m *GetUserMatchesParameters) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserMatchesParameters.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserMatchesParameters proto.InternalMessageInfo

func (m *GetUserMatchesParameters) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type ClearUserReferenceImageParameters struct {
	UserId               int64    `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClearUserReferenceImageParameters) Reset()         { *m = ClearUserReferenceImageParameters{} }
func (m *ClearUserReferenceImageParameters) String() string { return proto.CompactTextString(m) }
func (*ClearUserReferenceImageParameters) ProtoMessage()    {}
func (*ClearUserReferenceImageParameters) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d0473497cda6150, []int{3}
}

func (m *ClearUserReferenceImageParameters) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClearUserReferenceImageParameters.Unmarshal(m, b)
}
func (m *ClearUserReferenceImageParameters) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClearUserReferenceImageParameters.Marshal(b, m, deterministic)
}
func (m *ClearUserReferenceImageParameters) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClearUserReferenceImageParameters.Merge(m, src)
}
func (m *ClearUserReferenceImageParameters) XXX_Size() int {
	return xxx_messageInfo_ClearUserReferenceImageParameters.Size(m)
}
func (m *ClearUserReferenceImageParameters) XXX_DiscardUnknown() {
	xxx_messageInfo_ClearUserReferenceImageParameters.DiscardUnknown(m)
}

var xxx_messageInfo_ClearUserReferenceImageParameters proto.InternalMessageInfo

func (m *ClearUserReferenceImageParameters) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type SetUserReferenceImageParameters struct {
	UserId               int64    `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	MatchId              int64    `protobuf:"varint,2,opt,name=matchId,proto3" json:"matchId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetUserReferenceImageParameters) Reset()         { *m = SetUserReferenceImageParameters{} }
func (m *SetUserReferenceImageParameters) String() string { return proto.CompactTextString(m) }
func (*SetUserReferenceImageParameters) ProtoMessage()    {}
func (*SetUserReferenceImageParameters) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d0473497cda6150, []int{4}
}

func (m *SetUserReferenceImageParameters) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetUserReferenceImageParameters.Unmarshal(m, b)
}
func (m *SetUserReferenceImageParameters) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetUserReferenceImageParameters.Marshal(b, m, deterministic)
}
func (m *SetUserReferenceImageParameters) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetUserReferenceImageParameters.Merge(m, src)
}
func (m *SetUserReferenceImageParameters) XXX_Size() int {
	return xxx_messageInfo_SetUserReferenceImageParameters.Size(m)
}
func (m *SetUserReferenceImageParameters) XXX_DiscardUnknown() {
	xxx_messageInfo_SetUserReferenceImageParameters.DiscardUnknown(m)
}

var xxx_messageInfo_SetUserReferenceImageParameters proto.InternalMessageInfo

func (m *SetUserReferenceImageParameters) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *SetUserReferenceImageParameters) GetMatchId() int64 {
	if m != nil {
		return m.MatchId
	}
	return 0
}

type Match struct {
	MatchId              int64    `protobuf:"varint,1,opt,name=matchId,proto3" json:"matchId,omitempty"`
	UserId               int64    `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
	FlowId               int64    `protobuf:"varint,3,opt,name=flowId,proto3" json:"flowId,omitempty"`
	Confidence           float64  `protobuf:"fixed64,4,opt,name=confidence,proto3" json:"confidence,omitempty"`
	Similarity           float64  `protobuf:"fixed64,5,opt,name=similarity,proto3" json:"similarity,omitempty"`
	S3Reference          string   `protobuf:"bytes,6,opt,name=s3Reference,proto3" json:"s3Reference,omitempty"`
	IsReferenceImage     bool     `protobuf:"varint,7,opt,name=isReferenceImage,proto3" json:"isReferenceImage,omitempty"`
	CreatedAt            int64    `protobuf:"varint,8,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Match) Reset()         { *m = Match{} }
func (m *Match) String() string { return proto.CompactTextString(m) }
func (*Match) ProtoMessage()    {}
func (*Match) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d0473497cda6150, []int{5}
}

func (m *Match) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Match.Unmarshal(m, b)
}
func (m *Match) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Match.Marshal(b, m, deterministic)
}
func (m *Match) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Match.Merge(m, src)
}
func (m *Match) XXX_Size() int {
	return xxx_messageInfo_Match.Size(m)
}
func (m *Match) XXX_DiscardUnknown() {
	xxx_messageInfo_Match.DiscardUnknown(m)
}

var xxx_messageInfo_Match proto.InternalMessageInfo

func (m *Match) GetMatchId() int64 {
	if m != nil {
		return m.MatchId
	}
	return 0
}

func (m *Match) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *Match) GetFlowId() int64 {
	if m != nil {
		return m.FlowId
	}
	return 0
}

func (m *Match) GetConfidence() float64 {
	if m != nil {
		return m.Confidence
	}
	return 0
}

func (m *Match) GetSimilarity() float64 {
	if m != nil {
		return m.Similarity
	}
	return 0
}

func (m *Match) GetS3Reference() string {
	if m != nil {
		return m.S3Reference
	}
	return ""
}

func (m *Match) GetIsReferenceImage() bool {
	if m != nil {
		return m.IsReferenceImage
	}
	return false
}

func (m *Match) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

type MatchError struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Code                 int64    `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MatchError) Reset()         { *m = MatchError{} }
func (m *MatchError) String() string { return proto.CompactTextString(m) }
func (*MatchError) ProtoMessage()    {}
func (*MatchError) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d0473497cda6150, []int{6}
}

func (m *MatchError) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MatchError.Unmarshal(m, b)
}
func (m *MatchError) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MatchError.Marshal(b, m, deterministic)
}
func (m *MatchError) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MatchError.Merge(m, src)
}
func (m *MatchError) XXX_Size() int {
	return xxx_messageInfo_MatchError.Size(m)
}
func (m *MatchError) XXX_DiscardUnknown() {
	xxx_messageInfo_MatchError.DiscardUnknown(m)
}

var xxx_messageInfo_MatchError proto.InternalMessageInfo

func (m *MatchError) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *MatchError) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

type MatchResponse struct {
	MatchId              int64       `protobuf:"varint,1,opt,name=matchId,proto3" json:"matchId,omitempty"`
	Confidence           float64     `protobuf:"fixed64,2,opt,name=confidence,proto3" json:"confidence,omitempty"`
	Similarity           float64     `protobuf:"fixed64,3,opt,name=similarity,proto3" json:"similarity,omitempty"`
	IsNewReferenceImage  bool        `protobuf:"varint,4,opt,name=isNewReferenceImage,proto3" json:"isNewReferenceImage,omitempty"`
	Error                *MatchError `protobuf:"bytes,5,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *MatchResponse) Reset()         { *m = MatchResponse{} }
func (m *MatchResponse) String() string { return proto.CompactTextString(m) }
func (*MatchResponse) ProtoMessage()    {}
func (*MatchResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d0473497cda6150, []int{7}
}

func (m *MatchResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MatchResponse.Unmarshal(m, b)
}
func (m *MatchResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MatchResponse.Marshal(b, m, deterministic)
}
func (m *MatchResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MatchResponse.Merge(m, src)
}
func (m *MatchResponse) XXX_Size() int {
	return xxx_messageInfo_MatchResponse.Size(m)
}
func (m *MatchResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MatchResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MatchResponse proto.InternalMessageInfo

func (m *MatchResponse) GetMatchId() int64 {
	if m != nil {
		return m.MatchId
	}
	return 0
}

func (m *MatchResponse) GetConfidence() float64 {
	if m != nil {
		return m.Confidence
	}
	return 0
}

func (m *MatchResponse) GetSimilarity() float64 {
	if m != nil {
		return m.Similarity
	}
	return 0
}

func (m *MatchResponse) GetIsNewReferenceImage() bool {
	if m != nil {
		return m.IsNewReferenceImage
	}
	return false
}

func (m *MatchResponse) GetError() *MatchError {
	if m != nil {
		return m.Error
	}
	return nil
}

type GetMatchResponse struct {
	Match                *Match   `protobuf:"bytes,1,opt,name=match,proto3" json:"match,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetMatchResponse) Reset()         { *m = GetMatchResponse{} }
func (m *GetMatchResponse) String() string { return proto.CompactTextString(m) }
func (*GetMatchResponse) ProtoMessage()    {}
func (*GetMatchResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d0473497cda6150, []int{8}
}

func (m *GetMatchResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMatchResponse.Unmarshal(m, b)
}
func (m *GetMatchResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMatchResponse.Marshal(b, m, deterministic)
}
func (m *GetMatchResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMatchResponse.Merge(m, src)
}
func (m *GetMatchResponse) XXX_Size() int {
	return xxx_messageInfo_GetMatchResponse.Size(m)
}
func (m *GetMatchResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMatchResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetMatchResponse proto.InternalMessageInfo

func (m *GetMatchResponse) GetMatch() *Match {
	if m != nil {
		return m.Match
	}
	return nil
}

type GetUserMatchesResponse struct {
	Matches              []*Match `protobuf:"bytes,1,rep,name=matches,proto3" json:"matches,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserMatchesResponse) Reset()         { *m = GetUserMatchesResponse{} }
func (m *GetUserMatchesResponse) String() string { return proto.CompactTextString(m) }
func (*GetUserMatchesResponse) ProtoMessage()    {}
func (*GetUserMatchesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d0473497cda6150, []int{9}
}

func (m *GetUserMatchesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserMatchesResponse.Unmarshal(m, b)
}
func (m *GetUserMatchesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserMatchesResponse.Marshal(b, m, deterministic)
}
func (m *GetUserMatchesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserMatchesResponse.Merge(m, src)
}
func (m *GetUserMatchesResponse) XXX_Size() int {
	return xxx_messageInfo_GetUserMatchesResponse.Size(m)
}
func (m *GetUserMatchesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserMatchesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserMatchesResponse proto.InternalMessageInfo

func (m *GetUserMatchesResponse) GetMatches() []*Match {
	if m != nil {
		return m.Matches
	}
	return nil
}

type ClearUserReferenceImageResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClearUserReferenceImageResponse) Reset()         { *m = ClearUserReferenceImageResponse{} }
func (m *ClearUserReferenceImageResponse) String() string { return proto.CompactTextString(m) }
func (*ClearUserReferenceImageResponse) ProtoMessage()    {}
func (*ClearUserReferenceImageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d0473497cda6150, []int{10}
}

func (m *ClearUserReferenceImageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClearUserReferenceImageResponse.Unmarshal(m, b)
}
func (m *ClearUserReferenceImageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClearUserReferenceImageResponse.Marshal(b, m, deterministic)
}
func (m *ClearUserReferenceImageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClearUserReferenceImageResponse.Merge(m, src)
}
func (m *ClearUserReferenceImageResponse) XXX_Size() int {
	return xxx_messageInfo_ClearUserReferenceImageResponse.Size(m)
}
func (m *ClearUserReferenceImageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ClearUserReferenceImageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ClearUserReferenceImageResponse proto.InternalMessageInfo

type SetUserReferenceImageResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetUserReferenceImageResponse) Reset()         { *m = SetUserReferenceImageResponse{} }
func (m *SetUserReferenceImageResponse) String() string { return proto.CompactTextString(m) }
func (*SetUserReferenceImageResponse) ProtoMessage()    {}
func (*SetUserReferenceImageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d0473497cda6150, []int{11}
}

func (m *SetUserReferenceImageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetUserReferenceImageResponse.Unmarshal(m, b)
}
func (m *SetUserReferenceImageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetUserReferenceImageResponse.Marshal(b, m, deterministic)
}
func (m *SetUserReferenceImageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetUserReferenceImageResponse.Merge(m, src)
}
func (m *SetUserReferenceImageResponse) XXX_Size() int {
	return xxx_messageInfo_SetUserReferenceImageResponse.Size(m)
}
func (m *SetUserReferenceImageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SetUserReferenceImageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SetUserReferenceImageResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MatchParameters)(nil), "Proto.Face.Recognition.MatchParameters")
	proto.RegisterType((*GetMatchParameters)(nil), "Proto.Face.Recognition.GetMatchParameters")
	proto.RegisterType((*GetUserMatchesParameters)(nil), "Proto.Face.Recognition.GetUserMatchesParameters")
	proto.RegisterType((*ClearUserReferenceImageParameters)(nil), "Proto.Face.Recognition.ClearUserReferenceImageParameters")
	proto.RegisterType((*SetUserReferenceImageParameters)(nil), "Proto.Face.Recognition.SetUserReferenceImageParameters")
	proto.RegisterType((*Match)(nil), "Proto.Face.Recognition.Match")
	proto.RegisterType((*MatchError)(nil), "Proto.Face.Recognition.MatchError")
	proto.RegisterType((*MatchResponse)(nil), "Proto.Face.Recognition.MatchResponse")
	proto.RegisterType((*GetMatchResponse)(nil), "Proto.Face.Recognition.GetMatchResponse")
	proto.RegisterType((*GetUserMatchesResponse)(nil), "Proto.Face.Recognition.GetUserMatchesResponse")
	proto.RegisterType((*ClearUserReferenceImageResponse)(nil), "Proto.Face.Recognition.ClearUserReferenceImageResponse")
	proto.RegisterType((*SetUserReferenceImageResponse)(nil), "Proto.Face.Recognition.SetUserReferenceImageResponse")
}

func init() {
	proto.RegisterFile("proto/Recognition/face_recognition.proto", fileDescriptor_1d0473497cda6150)
}

var fileDescriptor_1d0473497cda6150 = []byte{
	// 579 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x55, 0x4d, 0x6f, 0xd3, 0x4c,
	0x10, 0x7e, 0x37, 0x5f, 0x4d, 0x27, 0x7a, 0xa1, 0x5a, 0x44, 0xb0, 0x22, 0x4a, 0xd2, 0x95, 0x10,
	0x51, 0x0f, 0x6e, 0x95, 0x08, 0x95, 0x8f, 0x13, 0x42, 0x10, 0xe5, 0x00, 0x2a, 0x8e, 0x38, 0x70,
	0x42, 0xee, 0x7a, 0x12, 0x56, 0x4a, 0xec, 0x6a, 0x77, 0xa1, 0xe2, 0xc8, 0x91, 0x13, 0xbf, 0x8e,
	0x5f, 0xc1, 0x6f, 0x40, 0x42, 0x5e, 0xc7, 0x5f, 0x89, 0x1d, 0x37, 0xa7, 0x64, 0x9f, 0x9d, 0x67,
	0x66, 0x9e, 0x67, 0x77, 0xd6, 0x30, 0xbc, 0x96, 0x81, 0x0e, 0xce, 0x1c, 0xe4, 0xc1, 0xc2, 0x17,
	0x5a, 0x04, 0xfe, 0xd9, 0xdc, 0xe5, 0xf8, 0x59, 0xa6, 0x80, 0x6d, 0x42, 0x68, 0xf7, 0x32, 0xfc,
	0xb1, 0xdf, 0xba, 0x1c, 0xed, 0x4c, 0x38, 0xfb, 0x41, 0xe0, 0xee, 0x3b, 0x57, 0xf3, 0x2f, 0x97,
	0xae, 0x74, 0x57, 0xa8, 0x51, 0x2a, 0xda, 0x85, 0xd6, 0x57, 0x85, 0x72, 0xea, 0x59, 0x64, 0x40,
	0x86, 0x75, 0x67, 0xbd, 0x0a, 0xf1, 0xf9, 0x32, 0xb8, 0x99, 0x7a, 0x56, 0x2d, 0xc2, 0xa3, 0x15,
	0x7d, 0x08, 0x87, 0x4b, 0xc1, 0xd1, 0x57, 0x38, 0xf5, 0xac, 0xba, 0xd9, 0x4a, 0x01, 0x3a, 0x80,
	0x8e, 0x1a, 0x3b, 0x38, 0x47, 0x89, 0x3e, 0x47, 0xab, 0x31, 0x20, 0xc3, 0x43, 0x27, 0x0b, 0x31,
	0x1b, 0xe8, 0x04, 0xf5, 0x66, 0x17, 0x16, 0x1c, 0xac, 0x42, 0x28, 0x69, 0x23, 0x5e, 0xb2, 0x11,
	0x58, 0x13, 0xd4, 0x1f, 0x15, 0x4a, 0xc3, 0x41, 0x55, 0xdd, 0x3b, 0x7b, 0x09, 0x27, 0xaf, 0x97,
	0xe8, 0xca, 0x90, 0x95, 0x54, 0x9e, 0xae, 0xdc, 0x05, 0xde, 0x82, 0x3c, 0x83, 0xfe, 0x2c, 0x2a,
	0xb8, 0x2f, 0x35, 0xab, 0xa2, 0x96, 0x57, 0xf1, 0x97, 0x40, 0xd3, 0xf4, 0x5f, 0xae, 0x34, 0x93,
	0xb5, 0x56, 0x72, 0x12, 0xf5, 0xdc, 0x49, 0x3c, 0x02, 0xe0, 0x81, 0x3f, 0x17, 0x5e, 0x62, 0x35,
	0x71, 0x32, 0x48, 0xb8, 0xaf, 0xc4, 0x4a, 0x2c, 0x5d, 0x29, 0xf4, 0x77, 0xab, 0x19, 0xed, 0xa7,
	0xc8, 0xe6, 0x59, 0xb5, 0xb6, 0xce, 0x8a, 0x9e, 0xc2, 0x91, 0x50, 0x79, 0x17, 0xac, 0x83, 0x01,
	0x19, 0xb6, 0x9d, 0x2d, 0x3c, 0xbc, 0x17, 0x5c, 0xa2, 0xab, 0xd1, 0x7b, 0xa5, 0xad, 0x76, 0x74,
	0x2f, 0x12, 0x80, 0xbd, 0x00, 0x30, 0xf2, 0xdf, 0x48, 0x19, 0x48, 0xe3, 0x01, 0x2a, 0x15, 0xa6,
	0x23, 0xa6, 0x6a, 0xbc, 0xa4, 0x14, 0x1a, 0x3c, 0xf0, 0x70, 0xed, 0x80, 0xf9, 0xcf, 0x7e, 0x13,
	0xf8, 0xdf, 0x90, 0x1d, 0x54, 0xd7, 0x81, 0xaf, 0x70, 0x87, 0x87, 0x79, 0x4f, 0x6a, 0x15, 0x9e,
	0xd4, 0xb7, 0x3c, 0x39, 0x87, 0x7b, 0x42, 0xbd, 0xc7, 0x9b, 0x0d, 0xd1, 0x0d, 0x23, 0xba, 0x68,
	0x8b, 0x3e, 0x83, 0x26, 0x86, 0xa2, 0x8c, 0xc1, 0x9d, 0x11, 0xb3, 0x8b, 0x67, 0xcf, 0x4e, 0xe5,
	0x3b, 0x11, 0x81, 0x4d, 0xe0, 0x28, 0x9e, 0x84, 0x44, 0xd9, 0x18, 0x9a, 0x46, 0x8a, 0xd1, 0xd5,
	0x19, 0x1d, 0xef, 0xcc, 0xe6, 0x44, 0xb1, 0xec, 0x03, 0x74, 0xf3, 0x23, 0x92, 0xa4, 0xbb, 0x58,
	0x1b, 0x85, 0xca, 0x22, 0x83, 0x7a, 0x75, 0xc2, 0x38, 0x9a, 0x9d, 0x40, 0xbf, 0x64, 0x82, 0xe2,
	0xdc, 0xac, 0x0f, 0xc7, 0x85, 0x73, 0x12, 0x07, 0x8c, 0xfe, 0x34, 0xa0, 0x1b, 0xd6, 0xc9, 0x94,
	0x99, 0xa1, 0xfc, 0x26, 0x38, 0xd2, 0x4f, 0x6b, 0x99, 0xf4, 0xc9, 0xce, 0x7e, 0xd2, 0x91, 0xeb,
	0x3d, 0xde, 0xdd, 0x78, 0xdc, 0xd4, 0x7f, 0xf4, 0x0a, 0xda, 0x8b, 0xb5, 0xab, 0xf4, 0xb4, 0x8c,
	0xb4, 0xfd, 0x02, 0xf5, 0x86, 0x55, 0xb1, 0x99, 0x1a, 0x1a, 0xee, 0x2c, 0x72, 0x86, 0xd3, 0xf3,
	0x1d, 0xec, 0xc2, 0xb7, 0xab, 0x67, 0xdf, 0x8e, 0x91, 0xa9, 0xfa, 0x8b, 0xc0, 0x03, 0x5e, 0x7c,
	0x28, 0xf4, 0x79, 0x59, 0xb6, 0xca, 0x77, 0xb0, 0x77, 0xb1, 0x27, 0x35, 0xd3, 0xd1, 0x4f, 0x02,
	0xf7, 0x55, 0xd1, 0x1d, 0xa0, 0xa5, 0x49, 0x2b, 0x9e, 0xd6, 0xde, 0xd3, 0xbd, 0x88, 0x69, 0x2f,
	0x57, 0x2d, 0xf3, 0xe9, 0x1b, 0xff, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x75, 0xd6, 0xf5, 0x17, 0x26,
	0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FaceRecognitionServiceClient is the client API for FaceRecognitionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FaceRecognitionServiceClient interface {
	Match(ctx context.Context, in *MatchParameters, opts ...grpc.CallOption) (*MatchResponse, error)
	GetMatch(ctx context.Context, in *GetMatchParameters, opts ...grpc.CallOption) (*GetMatchResponse, error)
	GetUserMatches(ctx context.Context, in *GetUserMatchesParameters, opts ...grpc.CallOption) (*GetUserMatchesResponse, error)
	ClearUserReferenceImage(ctx context.Context, in *ClearUserReferenceImageParameters, opts ...grpc.CallOption) (*ClearUserReferenceImageResponse, error)
	SetUserReferenceImage(ctx context.Context, in *SetUserReferenceImageParameters, opts ...grpc.CallOption) (*SetUserReferenceImageResponse, error)
}

type faceRecognitionServiceClient struct {
	cc *grpc.ClientConn
}

func NewFaceRecognitionServiceClient(cc *grpc.ClientConn) FaceRecognitionServiceClient {
	return &faceRecognitionServiceClient{cc}
}

func (c *faceRecognitionServiceClient) Match(ctx context.Context, in *MatchParameters, opts ...grpc.CallOption) (*MatchResponse, error) {
	out := new(MatchResponse)
	err := c.cc.Invoke(ctx, "/Proto.Face.Recognition.FaceRecognitionService/match", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *faceRecognitionServiceClient) GetMatch(ctx context.Context, in *GetMatchParameters, opts ...grpc.CallOption) (*GetMatchResponse, error) {
	out := new(GetMatchResponse)
	err := c.cc.Invoke(ctx, "/Proto.Face.Recognition.FaceRecognitionService/getMatch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *faceRecognitionServiceClient) GetUserMatches(ctx context.Context, in *GetUserMatchesParameters, opts ...grpc.CallOption) (*GetUserMatchesResponse, error) {
	out := new(GetUserMatchesResponse)
	err := c.cc.Invoke(ctx, "/Proto.Face.Recognition.FaceRecognitionService/getUserMatches", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *faceRecognitionServiceClient) ClearUserReferenceImage(ctx context.Context, in *ClearUserReferenceImageParameters, opts ...grpc.CallOption) (*ClearUserReferenceImageResponse, error) {
	out := new(ClearUserReferenceImageResponse)
	err := c.cc.Invoke(ctx, "/Proto.Face.Recognition.FaceRecognitionService/clearUserReferenceImage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *faceRecognitionServiceClient) SetUserReferenceImage(ctx context.Context, in *SetUserReferenceImageParameters, opts ...grpc.CallOption) (*SetUserReferenceImageResponse, error) {
	out := new(SetUserReferenceImageResponse)
	err := c.cc.Invoke(ctx, "/Proto.Face.Recognition.FaceRecognitionService/setUserReferenceImage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FaceRecognitionServiceServer is the server API for FaceRecognitionService service.
type FaceRecognitionServiceServer interface {
	Match(context.Context, *MatchParameters) (*MatchResponse, error)
	GetMatch(context.Context, *GetMatchParameters) (*GetMatchResponse, error)
	GetUserMatches(context.Context, *GetUserMatchesParameters) (*GetUserMatchesResponse, error)
	ClearUserReferenceImage(context.Context, *ClearUserReferenceImageParameters) (*ClearUserReferenceImageResponse, error)
	SetUserReferenceImage(context.Context, *SetUserReferenceImageParameters) (*SetUserReferenceImageResponse, error)
}

// UnimplementedFaceRecognitionServiceServer can be embedded to have forward compatible implementations.
type UnimplementedFaceRecognitionServiceServer struct {
}

func (*UnimplementedFaceRecognitionServiceServer) Match(ctx context.Context, req *MatchParameters) (*MatchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Match not implemented")
}
func (*UnimplementedFaceRecognitionServiceServer) GetMatch(ctx context.Context, req *GetMatchParameters) (*GetMatchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMatch not implemented")
}
func (*UnimplementedFaceRecognitionServiceServer) GetUserMatches(ctx context.Context, req *GetUserMatchesParameters) (*GetUserMatchesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserMatches not implemented")
}
func (*UnimplementedFaceRecognitionServiceServer) ClearUserReferenceImage(ctx context.Context, req *ClearUserReferenceImageParameters) (*ClearUserReferenceImageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClearUserReferenceImage not implemented")
}
func (*UnimplementedFaceRecognitionServiceServer) SetUserReferenceImage(ctx context.Context, req *SetUserReferenceImageParameters) (*SetUserReferenceImageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetUserReferenceImage not implemented")
}

func RegisterFaceRecognitionServiceServer(s *grpc.Server, srv FaceRecognitionServiceServer) {
	s.RegisterService(&_FaceRecognitionService_serviceDesc, srv)
}

func _FaceRecognitionService_Match_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MatchParameters)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FaceRecognitionServiceServer).Match(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Proto.Face.Recognition.FaceRecognitionService/Match",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FaceRecognitionServiceServer).Match(ctx, req.(*MatchParameters))
	}
	return interceptor(ctx, in, info, handler)
}

func _FaceRecognitionService_GetMatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMatchParameters)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FaceRecognitionServiceServer).GetMatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Proto.Face.Recognition.FaceRecognitionService/GetMatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FaceRecognitionServiceServer).GetMatch(ctx, req.(*GetMatchParameters))
	}
	return interceptor(ctx, in, info, handler)
}

func _FaceRecognitionService_GetUserMatches_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserMatchesParameters)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FaceRecognitionServiceServer).GetUserMatches(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Proto.Face.Recognition.FaceRecognitionService/GetUserMatches",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FaceRecognitionServiceServer).GetUserMatches(ctx, req.(*GetUserMatchesParameters))
	}
	return interceptor(ctx, in, info, handler)
}

func _FaceRecognitionService_ClearUserReferenceImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClearUserReferenceImageParameters)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FaceRecognitionServiceServer).ClearUserReferenceImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Proto.Face.Recognition.FaceRecognitionService/ClearUserReferenceImage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FaceRecognitionServiceServer).ClearUserReferenceImage(ctx, req.(*ClearUserReferenceImageParameters))
	}
	return interceptor(ctx, in, info, handler)
}

func _FaceRecognitionService_SetUserReferenceImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetUserReferenceImageParameters)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FaceRecognitionServiceServer).SetUserReferenceImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Proto.Face.Recognition.FaceRecognitionService/SetUserReferenceImage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FaceRecognitionServiceServer).SetUserReferenceImage(ctx, req.(*SetUserReferenceImageParameters))
	}
	return interceptor(ctx, in, info, handler)
}

var _FaceRecognitionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Proto.Face.Recognition.FaceRecognitionService",
	HandlerType: (*FaceRecognitionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "match",
			Handler:    _FaceRecognitionService_Match_Handler,
		},
		{
			MethodName: "getMatch",
			Handler:    _FaceRecognitionService_GetMatch_Handler,
		},
		{
			MethodName: "getUserMatches",
			Handler:    _FaceRecognitionService_GetUserMatches_Handler,
		},
		{
			MethodName: "clearUserReferenceImage",
			Handler:    _FaceRecognitionService_ClearUserReferenceImage_Handler,
		},
		{
			MethodName: "setUserReferenceImage",
			Handler:    _FaceRecognitionService_SetUserReferenceImage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/Recognition/face_recognition.proto",
}
