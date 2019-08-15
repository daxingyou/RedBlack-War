// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb_msg.proto

package pb_msg

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

//游戏阶段
type GameStage int32

const (
	GameStage_XXX     GameStage = 0
	GameStage_DownBet GameStage = 1
	GameStage_Settle  GameStage = 2
)

var GameStage_name = map[int32]string{
	0: "XXX",
	1: "DownBet",
	2: "Settle",
}

var GameStage_value = map[string]int32{
	"XXX":     0,
	"DownBet": 1,
	"Settle":  2,
}

func (x GameStage) String() string {
	return proto.EnumName(GameStage_name, int32(x))
}

func (GameStage) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_bd13c7201b21be60, []int{0}
}

//卡牌类型
type CardsType int32

const (
	CardsType_XXX_Card CardsType = 0
	CardsType_Leaflet  CardsType = 1
	CardsType_Pair     CardsType = 2
	CardsType_Straight CardsType = 3
	CardsType_Golden   CardsType = 4
	CardsType_ShunJin  CardsType = 5
	CardsType_Leopard  CardsType = 6
)

var CardsType_name = map[int32]string{
	0: "XXX_Card",
	1: "Leaflet",
	2: "Pair",
	3: "Straight",
	4: "Golden",
	5: "ShunJin",
	6: "Leopard",
}

var CardsType_value = map[string]int32{
	"XXX_Card": 0,
	"Leaflet":  1,
	"Pair":     2,
	"Straight": 3,
	"Golden":   4,
	"ShunJin":  5,
	"Leopard":  6,
}

func (x CardsType) String() string {
	return proto.EnumName(CardsType_name, int32(x))
}

func (CardsType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_bd13c7201b21be60, []int{1}
}

type PotType int32

const (
	PotType_XXX_Pot  PotType = 0
	PotType_ReadPot  PotType = 1
	PotType_BlackPot PotType = 2
	PotType_LuckPot  PotType = 3
)

var PotType_name = map[int32]string{
	0: "XXX_Pot",
	1: "ReadPot",
	2: "BlackPot",
	3: "LuckPot",
}

var PotType_value = map[string]int32{
	"XXX_Pot":  0,
	"ReadPot":  1,
	"BlackPot": 2,
	"LuckPot":  3,
}

func (x PotType) String() string {
	return proto.EnumName(PotType_name, int32(x))
}

func (PotType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_bd13c7201b21be60, []int{2}
}

//--0
type Ping struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ping) Reset()         { *m = Ping{} }
func (m *Ping) String() string { return proto.CompactTextString(m) }
func (*Ping) ProtoMessage()    {}
func (*Ping) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd13c7201b21be60, []int{0}
}

func (m *Ping) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ping.Unmarshal(m, b)
}
func (m *Ping) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ping.Marshal(b, m, deterministic)
}
func (m *Ping) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ping.Merge(m, src)
}
func (m *Ping) XXX_Size() int {
	return xxx_messageInfo_Ping.Size(m)
}
func (m *Ping) XXX_DiscardUnknown() {
	xxx_messageInfo_Ping.DiscardUnknown(m)
}

var xxx_messageInfo_Ping proto.InternalMessageInfo

//--1
type Pong struct {
	ServerTime           int64    `protobuf:"varint,1,opt,name=serverTime,proto3" json:"serverTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Pong) Reset()         { *m = Pong{} }
func (m *Pong) String() string { return proto.CompactTextString(m) }
func (*Pong) ProtoMessage()    {}
func (*Pong) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd13c7201b21be60, []int{1}
}

func (m *Pong) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pong.Unmarshal(m, b)
}
func (m *Pong) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pong.Marshal(b, m, deterministic)
}
func (m *Pong) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pong.Merge(m, src)
}
func (m *Pong) XXX_Size() int {
	return xxx_messageInfo_Pong.Size(m)
}
func (m *Pong) XXX_DiscardUnknown() {
	xxx_messageInfo_Pong.DiscardUnknown(m)
}

var xxx_messageInfo_Pong proto.InternalMessageInfo

func (m *Pong) GetServerTime() int64 {
	if m != nil {
		return m.ServerTime
	}
	return 0
}

//--2
type MsgInfo_S2C struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Data                 string   `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MsgInfo_S2C) Reset()         { *m = MsgInfo_S2C{} }
func (m *MsgInfo_S2C) String() string { return proto.CompactTextString(m) }
func (*MsgInfo_S2C) ProtoMessage()    {}
func (*MsgInfo_S2C) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd13c7201b21be60, []int{2}
}

func (m *MsgInfo_S2C) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgInfo_S2C.Unmarshal(m, b)
}
func (m *MsgInfo_S2C) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgInfo_S2C.Marshal(b, m, deterministic)
}
func (m *MsgInfo_S2C) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgInfo_S2C.Merge(m, src)
}
func (m *MsgInfo_S2C) XXX_Size() int {
	return xxx_messageInfo_MsgInfo_S2C.Size(m)
}
func (m *MsgInfo_S2C) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgInfo_S2C.DiscardUnknown(m)
}

var xxx_messageInfo_MsgInfo_S2C proto.InternalMessageInfo

func (m *MsgInfo_S2C) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *MsgInfo_S2C) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *MsgInfo_S2C) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

//--3
type LoginInfo_C2S struct {
	Id                   string   `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	GameId               string   `protobuf:"bytes,2,opt,name=gameId,proto3" json:"gameId,omitempty"`
	ServerUrl            string   `protobuf:"bytes,3,opt,name=serverUrl,proto3" json:"serverUrl,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginInfo_C2S) Reset()         { *m = LoginInfo_C2S{} }
func (m *LoginInfo_C2S) String() string { return proto.CompactTextString(m) }
func (*LoginInfo_C2S) ProtoMessage()    {}
func (*LoginInfo_C2S) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd13c7201b21be60, []int{3}
}

func (m *LoginInfo_C2S) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginInfo_C2S.Unmarshal(m, b)
}
func (m *LoginInfo_C2S) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginInfo_C2S.Marshal(b, m, deterministic)
}
func (m *LoginInfo_C2S) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginInfo_C2S.Merge(m, src)
}
func (m *LoginInfo_C2S) XXX_Size() int {
	return xxx_messageInfo_LoginInfo_C2S.Size(m)
}
func (m *LoginInfo_C2S) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginInfo_C2S.DiscardUnknown(m)
}

var xxx_messageInfo_LoginInfo_C2S proto.InternalMessageInfo

func (m *LoginInfo_C2S) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *LoginInfo_C2S) GetGameId() string {
	if m != nil {
		return m.GameId
	}
	return ""
}

func (m *LoginInfo_C2S) GetServerUrl() string {
	if m != nil {
		return m.ServerUrl
	}
	return ""
}

type LoginData struct {
	Id                   string   `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	NickName             string   `protobuf:"bytes,2,opt,name=nickName,proto3" json:"nickName,omitempty"`
	HeadImg              string   `protobuf:"bytes,3,opt,name=headImg,proto3" json:"headImg,omitempty"`
	Account              float64  `protobuf:"fixed64,4,opt,name=account,proto3" json:"account,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginData) Reset()         { *m = LoginData{} }
func (m *LoginData) String() string { return proto.CompactTextString(m) }
func (*LoginData) ProtoMessage()    {}
func (*LoginData) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd13c7201b21be60, []int{4}
}

func (m *LoginData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginData.Unmarshal(m, b)
}
func (m *LoginData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginData.Marshal(b, m, deterministic)
}
func (m *LoginData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginData.Merge(m, src)
}
func (m *LoginData) XXX_Size() int {
	return xxx_messageInfo_LoginData.Size(m)
}
func (m *LoginData) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginData.DiscardUnknown(m)
}

var xxx_messageInfo_LoginData proto.InternalMessageInfo

func (m *LoginData) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *LoginData) GetNickName() string {
	if m != nil {
		return m.NickName
	}
	return ""
}

func (m *LoginData) GetHeadImg() string {
	if m != nil {
		return m.HeadImg
	}
	return ""
}

func (m *LoginData) GetAccount() float64 {
	if m != nil {
		return m.Account
	}
	return 0
}

//--4
type LoginInfo_S2C struct {
	LoginData            *LoginData `protobuf:"bytes,1,opt,name=loginData,proto3" json:"loginData,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *LoginInfo_S2C) Reset()         { *m = LoginInfo_S2C{} }
func (m *LoginInfo_S2C) String() string { return proto.CompactTextString(m) }
func (*LoginInfo_S2C) ProtoMessage()    {}
func (*LoginInfo_S2C) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd13c7201b21be60, []int{5}
}

func (m *LoginInfo_S2C) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginInfo_S2C.Unmarshal(m, b)
}
func (m *LoginInfo_S2C) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginInfo_S2C.Marshal(b, m, deterministic)
}
func (m *LoginInfo_S2C) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginInfo_S2C.Merge(m, src)
}
func (m *LoginInfo_S2C) XXX_Size() int {
	return xxx_messageInfo_LoginInfo_S2C.Size(m)
}
func (m *LoginInfo_S2C) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginInfo_S2C.DiscardUnknown(m)
}

var xxx_messageInfo_LoginInfo_S2C proto.InternalMessageInfo

func (m *LoginInfo_S2C) GetLoginData() *LoginData {
	if m != nil {
		return m.LoginData
	}
	return nil
}

//--5
type JoinRoom_C2S struct {
	RoomId               string   `protobuf:"bytes,1,opt,name=roomId,proto3" json:"roomId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JoinRoom_C2S) Reset()         { *m = JoinRoom_C2S{} }
func (m *JoinRoom_C2S) String() string { return proto.CompactTextString(m) }
func (*JoinRoom_C2S) ProtoMessage()    {}
func (*JoinRoom_C2S) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd13c7201b21be60, []int{6}
}

func (m *JoinRoom_C2S) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JoinRoom_C2S.Unmarshal(m, b)
}
func (m *JoinRoom_C2S) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JoinRoom_C2S.Marshal(b, m, deterministic)
}
func (m *JoinRoom_C2S) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JoinRoom_C2S.Merge(m, src)
}
func (m *JoinRoom_C2S) XXX_Size() int {
	return xxx_messageInfo_JoinRoom_C2S.Size(m)
}
func (m *JoinRoom_C2S) XXX_DiscardUnknown() {
	xxx_messageInfo_JoinRoom_C2S.DiscardUnknown(m)
}

var xxx_messageInfo_JoinRoom_C2S proto.InternalMessageInfo

func (m *JoinRoom_C2S) GetRoomId() string {
	if m != nil {
		return m.RoomId
	}
	return ""
}

type PlayerData struct {
	LoginData            *LoginData `protobuf:"bytes,1,opt,name=loginData,proto3" json:"loginData,omitempty"`
	ContinueVot          float64    `protobuf:"fixed64,2,opt,name=continueVot,proto3" json:"continueVot,omitempty"`
	IsGodGambling        bool       `protobuf:"varint,3,opt,name=IsGodGambling,proto3" json:"IsGodGambling,omitempty"`
	WinCount             int32      `protobuf:"varint,4,opt,name=winCount,proto3" json:"winCount,omitempty"`
	ResultWinMoney       float64    `protobuf:"fixed64,5,opt,name=resultWinMoney,proto3" json:"resultWinMoney,omitempty"`
	ResultLoseMoney      float64    `protobuf:"fixed64,6,opt,name=resultLoseMoney,proto3" json:"resultLoseMoney,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *PlayerData) Reset()         { *m = PlayerData{} }
func (m *PlayerData) String() string { return proto.CompactTextString(m) }
func (*PlayerData) ProtoMessage()    {}
func (*PlayerData) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd13c7201b21be60, []int{7}
}

func (m *PlayerData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlayerData.Unmarshal(m, b)
}
func (m *PlayerData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlayerData.Marshal(b, m, deterministic)
}
func (m *PlayerData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlayerData.Merge(m, src)
}
func (m *PlayerData) XXX_Size() int {
	return xxx_messageInfo_PlayerData.Size(m)
}
func (m *PlayerData) XXX_DiscardUnknown() {
	xxx_messageInfo_PlayerData.DiscardUnknown(m)
}

var xxx_messageInfo_PlayerData proto.InternalMessageInfo

func (m *PlayerData) GetLoginData() *LoginData {
	if m != nil {
		return m.LoginData
	}
	return nil
}

func (m *PlayerData) GetContinueVot() float64 {
	if m != nil {
		return m.ContinueVot
	}
	return 0
}

func (m *PlayerData) GetIsGodGambling() bool {
	if m != nil {
		return m.IsGodGambling
	}
	return false
}

func (m *PlayerData) GetWinCount() int32 {
	if m != nil {
		return m.WinCount
	}
	return 0
}

func (m *PlayerData) GetResultWinMoney() float64 {
	if m != nil {
		return m.ResultWinMoney
	}
	return 0
}

func (m *PlayerData) GetResultLoseMoney() float64 {
	if m != nil {
		return m.ResultLoseMoney
	}
	return 0
}

//房间注池金额总数量
type PotMoneyCount struct {
	ReadMoneyCount       int32    `protobuf:"varint,1,opt,name=ReadMoneyCount,proto3" json:"ReadMoneyCount,omitempty"`
	BlackMoneyCount      int32    `protobuf:"varint,2,opt,name=BlackMoneyCount,proto3" json:"BlackMoneyCount,omitempty"`
	LuckMoneyCount       int32    `protobuf:"varint,3,opt,name=LuckMoneyCount,proto3" json:"LuckMoneyCount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PotMoneyCount) Reset()         { *m = PotMoneyCount{} }
func (m *PotMoneyCount) String() string { return proto.CompactTextString(m) }
func (*PotMoneyCount) ProtoMessage()    {}
func (*PotMoneyCount) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd13c7201b21be60, []int{8}
}

func (m *PotMoneyCount) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PotMoneyCount.Unmarshal(m, b)
}
func (m *PotMoneyCount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PotMoneyCount.Marshal(b, m, deterministic)
}
func (m *PotMoneyCount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PotMoneyCount.Merge(m, src)
}
func (m *PotMoneyCount) XXX_Size() int {
	return xxx_messageInfo_PotMoneyCount.Size(m)
}
func (m *PotMoneyCount) XXX_DiscardUnknown() {
	xxx_messageInfo_PotMoneyCount.DiscardUnknown(m)
}

var xxx_messageInfo_PotMoneyCount proto.InternalMessageInfo

func (m *PotMoneyCount) GetReadMoneyCount() int32 {
	if m != nil {
		return m.ReadMoneyCount
	}
	return 0
}

func (m *PotMoneyCount) GetBlackMoneyCount() int32 {
	if m != nil {
		return m.BlackMoneyCount
	}
	return 0
}

func (m *PotMoneyCount) GetLuckMoneyCount() int32 {
	if m != nil {
		return m.LuckMoneyCount
	}
	return 0
}

//游戏输赢的数据
type GameWinList struct {
	ReadWin              int32     `protobuf:"varint,1,opt,name=ReadWin,proto3" json:"ReadWin,omitempty"`
	BlackWin             int32     `protobuf:"varint,2,opt,name=BlackWin,proto3" json:"BlackWin,omitempty"`
	LuckWin              int32     `protobuf:"varint,3,opt,name=LuckWin,proto3" json:"LuckWin,omitempty"`
	CardType             CardsType `protobuf:"varint,4,opt,name=cardType,proto3,enum=pb_msg.CardsType" json:"cardType,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GameWinList) Reset()         { *m = GameWinList{} }
func (m *GameWinList) String() string { return proto.CompactTextString(m) }
func (*GameWinList) ProtoMessage()    {}
func (*GameWinList) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd13c7201b21be60, []int{9}
}

func (m *GameWinList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GameWinList.Unmarshal(m, b)
}
func (m *GameWinList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GameWinList.Marshal(b, m, deterministic)
}
func (m *GameWinList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GameWinList.Merge(m, src)
}
func (m *GameWinList) XXX_Size() int {
	return xxx_messageInfo_GameWinList.Size(m)
}
func (m *GameWinList) XXX_DiscardUnknown() {
	xxx_messageInfo_GameWinList.DiscardUnknown(m)
}

var xxx_messageInfo_GameWinList proto.InternalMessageInfo

func (m *GameWinList) GetReadWin() int32 {
	if m != nil {
		return m.ReadWin
	}
	return 0
}

func (m *GameWinList) GetBlackWin() int32 {
	if m != nil {
		return m.BlackWin
	}
	return 0
}

func (m *GameWinList) GetLuckWin() int32 {
	if m != nil {
		return m.LuckWin
	}
	return 0
}

func (m *GameWinList) GetCardType() CardsType {
	if m != nil {
		return m.CardType
	}
	return CardsType_XXX_Card
}

type RoomData struct {
	RoomId               string         `protobuf:"bytes,1,opt,name=roomId,proto3" json:"roomId,omitempty"`
	PlayerList           []*PlayerData  `protobuf:"bytes,2,rep,name=playerList,proto3" json:"playerList,omitempty"`
	GodGableName         string         `protobuf:"bytes,3,opt,name=godGableName,proto3" json:"godGableName,omitempty"`
	GameStage            GameStage      `protobuf:"varint,4,opt,name=gameStage,proto3,enum=pb_msg.GameStage" json:"gameStage,omitempty"`
	PotMoneyCount        *PotMoneyCount `protobuf:"bytes,5,opt,name=potMoneyCount,proto3" json:"potMoneyCount,omitempty"`
	CardTypes            []int32        `protobuf:"varint,6,rep,packed,name=cardTypes,proto3" json:"cardTypes,omitempty"`
	RPotWinList          []*GameWinList `protobuf:"bytes,7,rep,name=rPotWinList,proto3" json:"rPotWinList,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *RoomData) Reset()         { *m = RoomData{} }
func (m *RoomData) String() string { return proto.CompactTextString(m) }
func (*RoomData) ProtoMessage()    {}
func (*RoomData) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd13c7201b21be60, []int{10}
}

func (m *RoomData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoomData.Unmarshal(m, b)
}
func (m *RoomData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoomData.Marshal(b, m, deterministic)
}
func (m *RoomData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoomData.Merge(m, src)
}
func (m *RoomData) XXX_Size() int {
	return xxx_messageInfo_RoomData.Size(m)
}
func (m *RoomData) XXX_DiscardUnknown() {
	xxx_messageInfo_RoomData.DiscardUnknown(m)
}

var xxx_messageInfo_RoomData proto.InternalMessageInfo

func (m *RoomData) GetRoomId() string {
	if m != nil {
		return m.RoomId
	}
	return ""
}

func (m *RoomData) GetPlayerList() []*PlayerData {
	if m != nil {
		return m.PlayerList
	}
	return nil
}

func (m *RoomData) GetGodGableName() string {
	if m != nil {
		return m.GodGableName
	}
	return ""
}

func (m *RoomData) GetGameStage() GameStage {
	if m != nil {
		return m.GameStage
	}
	return GameStage_XXX
}

func (m *RoomData) GetPotMoneyCount() *PotMoneyCount {
	if m != nil {
		return m.PotMoneyCount
	}
	return nil
}

func (m *RoomData) GetCardTypes() []int32 {
	if m != nil {
		return m.CardTypes
	}
	return nil
}

func (m *RoomData) GetRPotWinList() []*GameWinList {
	if m != nil {
		return m.RPotWinList
	}
	return nil
}

//--6
type JoinRoom_S2C struct {
	RoomData             *RoomData `protobuf:"bytes,1,opt,name=roomData,proto3" json:"roomData,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *JoinRoom_S2C) Reset()         { *m = JoinRoom_S2C{} }
func (m *JoinRoom_S2C) String() string { return proto.CompactTextString(m) }
func (*JoinRoom_S2C) ProtoMessage()    {}
func (*JoinRoom_S2C) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd13c7201b21be60, []int{11}
}

func (m *JoinRoom_S2C) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JoinRoom_S2C.Unmarshal(m, b)
}
func (m *JoinRoom_S2C) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JoinRoom_S2C.Marshal(b, m, deterministic)
}
func (m *JoinRoom_S2C) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JoinRoom_S2C.Merge(m, src)
}
func (m *JoinRoom_S2C) XXX_Size() int {
	return xxx_messageInfo_JoinRoom_S2C.Size(m)
}
func (m *JoinRoom_S2C) XXX_DiscardUnknown() {
	xxx_messageInfo_JoinRoom_S2C.DiscardUnknown(m)
}

var xxx_messageInfo_JoinRoom_S2C proto.InternalMessageInfo

func (m *JoinRoom_S2C) GetRoomData() *RoomData {
	if m != nil {
		return m.RoomData
	}
	return nil
}

//--7
type LeaveRoom_C2S struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LeaveRoom_C2S) Reset()         { *m = LeaveRoom_C2S{} }
func (m *LeaveRoom_C2S) String() string { return proto.CompactTextString(m) }
func (*LeaveRoom_C2S) ProtoMessage()    {}
func (*LeaveRoom_C2S) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd13c7201b21be60, []int{12}
}

func (m *LeaveRoom_C2S) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LeaveRoom_C2S.Unmarshal(m, b)
}
func (m *LeaveRoom_C2S) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LeaveRoom_C2S.Marshal(b, m, deterministic)
}
func (m *LeaveRoom_C2S) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LeaveRoom_C2S.Merge(m, src)
}
func (m *LeaveRoom_C2S) XXX_Size() int {
	return xxx_messageInfo_LeaveRoom_C2S.Size(m)
}
func (m *LeaveRoom_C2S) XXX_DiscardUnknown() {
	xxx_messageInfo_LeaveRoom_C2S.DiscardUnknown(m)
}

var xxx_messageInfo_LeaveRoom_C2S proto.InternalMessageInfo

//--8
type LeaveRoom_S2C struct {
	LoginData            *LoginData `protobuf:"bytes,1,opt,name=loginData,proto3" json:"loginData,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *LeaveRoom_S2C) Reset()         { *m = LeaveRoom_S2C{} }
func (m *LeaveRoom_S2C) String() string { return proto.CompactTextString(m) }
func (*LeaveRoom_S2C) ProtoMessage()    {}
func (*LeaveRoom_S2C) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd13c7201b21be60, []int{13}
}

func (m *LeaveRoom_S2C) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LeaveRoom_S2C.Unmarshal(m, b)
}
func (m *LeaveRoom_S2C) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LeaveRoom_S2C.Marshal(b, m, deterministic)
}
func (m *LeaveRoom_S2C) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LeaveRoom_S2C.Merge(m, src)
}
func (m *LeaveRoom_S2C) XXX_Size() int {
	return xxx_messageInfo_LeaveRoom_S2C.Size(m)
}
func (m *LeaveRoom_S2C) XXX_DiscardUnknown() {
	xxx_messageInfo_LeaveRoom_S2C.DiscardUnknown(m)
}

var xxx_messageInfo_LeaveRoom_S2C proto.InternalMessageInfo

func (m *LeaveRoom_S2C) GetLoginData() *LoginData {
	if m != nil {
		return m.LoginData
	}
	return nil
}

//--9  玩家行动
type PlayerAction_C2S struct {
	Bet                  float64  `protobuf:"fixed64,1,opt,name=bet,proto3" json:"bet,omitempty"`
	Pot                  PotType  `protobuf:"varint,2,opt,name=pot,proto3,enum=pb_msg.PotType" json:"pot,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PlayerAction_C2S) Reset()         { *m = PlayerAction_C2S{} }
func (m *PlayerAction_C2S) String() string { return proto.CompactTextString(m) }
func (*PlayerAction_C2S) ProtoMessage()    {}
func (*PlayerAction_C2S) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd13c7201b21be60, []int{14}
}

func (m *PlayerAction_C2S) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlayerAction_C2S.Unmarshal(m, b)
}
func (m *PlayerAction_C2S) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlayerAction_C2S.Marshal(b, m, deterministic)
}
func (m *PlayerAction_C2S) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlayerAction_C2S.Merge(m, src)
}
func (m *PlayerAction_C2S) XXX_Size() int {
	return xxx_messageInfo_PlayerAction_C2S.Size(m)
}
func (m *PlayerAction_C2S) XXX_DiscardUnknown() {
	xxx_messageInfo_PlayerAction_C2S.DiscardUnknown(m)
}

var xxx_messageInfo_PlayerAction_C2S proto.InternalMessageInfo

func (m *PlayerAction_C2S) GetBet() float64 {
	if m != nil {
		return m.Bet
	}
	return 0
}

func (m *PlayerAction_C2S) GetPot() PotType {
	if m != nil {
		return m.Pot
	}
	return PotType_XXX_Pot
}

//--10
type PlayerAction_S2C struct {
	RoomData             *RoomData `protobuf:"bytes,1,opt,name=roomData,proto3" json:"roomData,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *PlayerAction_S2C) Reset()         { *m = PlayerAction_S2C{} }
func (m *PlayerAction_S2C) String() string { return proto.CompactTextString(m) }
func (*PlayerAction_S2C) ProtoMessage()    {}
func (*PlayerAction_S2C) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd13c7201b21be60, []int{15}
}

func (m *PlayerAction_S2C) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlayerAction_S2C.Unmarshal(m, b)
}
func (m *PlayerAction_S2C) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlayerAction_S2C.Marshal(b, m, deterministic)
}
func (m *PlayerAction_S2C) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlayerAction_S2C.Merge(m, src)
}
func (m *PlayerAction_S2C) XXX_Size() int {
	return xxx_messageInfo_PlayerAction_S2C.Size(m)
}
func (m *PlayerAction_S2C) XXX_DiscardUnknown() {
	xxx_messageInfo_PlayerAction_S2C.DiscardUnknown(m)
}

var xxx_messageInfo_PlayerAction_S2C proto.InternalMessageInfo

func (m *PlayerAction_S2C) GetRoomData() *RoomData {
	if m != nil {
		return m.RoomData
	}
	return nil
}

//--11  更新房间列表
type MaintainList_S2C struct {
	PlayerList           []*PlayerData `protobuf:"bytes,1,rep,name=playerList,proto3" json:"playerList,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *MaintainList_S2C) Reset()         { *m = MaintainList_S2C{} }
func (m *MaintainList_S2C) String() string { return proto.CompactTextString(m) }
func (*MaintainList_S2C) ProtoMessage()    {}
func (*MaintainList_S2C) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd13c7201b21be60, []int{16}
}

func (m *MaintainList_S2C) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MaintainList_S2C.Unmarshal(m, b)
}
func (m *MaintainList_S2C) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MaintainList_S2C.Marshal(b, m, deterministic)
}
func (m *MaintainList_S2C) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MaintainList_S2C.Merge(m, src)
}
func (m *MaintainList_S2C) XXX_Size() int {
	return xxx_messageInfo_MaintainList_S2C.Size(m)
}
func (m *MaintainList_S2C) XXX_DiscardUnknown() {
	xxx_messageInfo_MaintainList_S2C.DiscardUnknown(m)
}

var xxx_messageInfo_MaintainList_S2C proto.InternalMessageInfo

func (m *MaintainList_S2C) GetPlayerList() []*PlayerData {
	if m != nil {
		return m.PlayerList
	}
	return nil
}

//--12  开牌结果
type OpenCardResult_S2C struct {
	//1、卡牌数据
	CardData []int32 `protobuf:"varint,1,rep,packed,name=cardData,proto3" json:"cardData,omitempty"`
	//2、比牌类型
	CardType CardsType `protobuf:"varint,2,opt,name=cardType,proto3,enum=pb_msg.CardsType" json:"cardType,omitempty"`
	//3、注池Win类型
	PotType              PotType  `protobuf:"varint,3,opt,name=potType,proto3,enum=pb_msg.PotType" json:"potType,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OpenCardResult_S2C) Reset()         { *m = OpenCardResult_S2C{} }
func (m *OpenCardResult_S2C) String() string { return proto.CompactTextString(m) }
func (*OpenCardResult_S2C) ProtoMessage()    {}
func (*OpenCardResult_S2C) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd13c7201b21be60, []int{17}
}

func (m *OpenCardResult_S2C) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OpenCardResult_S2C.Unmarshal(m, b)
}
func (m *OpenCardResult_S2C) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OpenCardResult_S2C.Marshal(b, m, deterministic)
}
func (m *OpenCardResult_S2C) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OpenCardResult_S2C.Merge(m, src)
}
func (m *OpenCardResult_S2C) XXX_Size() int {
	return xxx_messageInfo_OpenCardResult_S2C.Size(m)
}
func (m *OpenCardResult_S2C) XXX_DiscardUnknown() {
	xxx_messageInfo_OpenCardResult_S2C.DiscardUnknown(m)
}

var xxx_messageInfo_OpenCardResult_S2C proto.InternalMessageInfo

func (m *OpenCardResult_S2C) GetCardData() []int32 {
	if m != nil {
		return m.CardData
	}
	return nil
}

func (m *OpenCardResult_S2C) GetCardType() CardsType {
	if m != nil {
		return m.CardType
	}
	return CardsType_XXX_Card
}

func (m *OpenCardResult_S2C) GetPotType() PotType {
	if m != nil {
		return m.PotType
	}
	return PotType_XXX_Pot
}

//--13  房间结算数据
type RoomSettleData_S2C struct {
	RoomData             *RoomData `protobuf:"bytes,1,opt,name=roomData,proto3" json:"roomData,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *RoomSettleData_S2C) Reset()         { *m = RoomSettleData_S2C{} }
func (m *RoomSettleData_S2C) String() string { return proto.CompactTextString(m) }
func (*RoomSettleData_S2C) ProtoMessage()    {}
func (*RoomSettleData_S2C) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd13c7201b21be60, []int{18}
}

func (m *RoomSettleData_S2C) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoomSettleData_S2C.Unmarshal(m, b)
}
func (m *RoomSettleData_S2C) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoomSettleData_S2C.Marshal(b, m, deterministic)
}
func (m *RoomSettleData_S2C) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoomSettleData_S2C.Merge(m, src)
}
func (m *RoomSettleData_S2C) XXX_Size() int {
	return xxx_messageInfo_RoomSettleData_S2C.Size(m)
}
func (m *RoomSettleData_S2C) XXX_DiscardUnknown() {
	xxx_messageInfo_RoomSettleData_S2C.DiscardUnknown(m)
}

var xxx_messageInfo_RoomSettleData_S2C proto.InternalMessageInfo

func (m *RoomSettleData_S2C) GetRoomData() *RoomData {
	if m != nil {
		return m.RoomData
	}
	return nil
}

func init() {
	proto.RegisterEnum("pb_msg.GameStage", GameStage_name, GameStage_value)
	proto.RegisterEnum("pb_msg.CardsType", CardsType_name, CardsType_value)
	proto.RegisterEnum("pb_msg.PotType", PotType_name, PotType_value)
	proto.RegisterType((*Ping)(nil), "pb_msg.Ping")
	proto.RegisterType((*Pong)(nil), "pb_msg.Pong")
	proto.RegisterType((*MsgInfo_S2C)(nil), "pb_msg.MsgInfo_S2C")
	proto.RegisterType((*LoginInfo_C2S)(nil), "pb_msg.LoginInfo_C2S")
	proto.RegisterType((*LoginData)(nil), "pb_msg.LoginData")
	proto.RegisterType((*LoginInfo_S2C)(nil), "pb_msg.LoginInfo_S2C")
	proto.RegisterType((*JoinRoom_C2S)(nil), "pb_msg.JoinRoom_C2S")
	proto.RegisterType((*PlayerData)(nil), "pb_msg.PlayerData")
	proto.RegisterType((*PotMoneyCount)(nil), "pb_msg.PotMoneyCount")
	proto.RegisterType((*GameWinList)(nil), "pb_msg.GameWinList")
	proto.RegisterType((*RoomData)(nil), "pb_msg.RoomData")
	proto.RegisterType((*JoinRoom_S2C)(nil), "pb_msg.JoinRoom_S2C")
	proto.RegisterType((*LeaveRoom_C2S)(nil), "pb_msg.LeaveRoom_C2S")
	proto.RegisterType((*LeaveRoom_S2C)(nil), "pb_msg.LeaveRoom_S2C")
	proto.RegisterType((*PlayerAction_C2S)(nil), "pb_msg.PlayerAction_C2S")
	proto.RegisterType((*PlayerAction_S2C)(nil), "pb_msg.PlayerAction_S2C")
	proto.RegisterType((*MaintainList_S2C)(nil), "pb_msg.MaintainList_S2C")
	proto.RegisterType((*OpenCardResult_S2C)(nil), "pb_msg.OpenCardResult_S2C")
	proto.RegisterType((*RoomSettleData_S2C)(nil), "pb_msg.RoomSettleData_S2C")
}

func init() { proto.RegisterFile("pb_msg.proto", fileDescriptor_bd13c7201b21be60) }

var fileDescriptor_bd13c7201b21be60 = []byte{
	// 882 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x56, 0x6d, 0x8f, 0xdb, 0x44,
	0x10, 0xc6, 0x76, 0xe2, 0x24, 0x93, 0x7b, 0x31, 0x0b, 0x54, 0x56, 0x85, 0x50, 0xb0, 0xd0, 0x29,
	0x9c, 0x68, 0x91, 0x82, 0xf8, 0x04, 0x42, 0xe5, 0xae, 0x22, 0x4a, 0x95, 0x83, 0x68, 0xd3, 0x72,
	0xf9, 0x76, 0xda, 0xc4, 0x5b, 0xdf, 0xea, 0xec, 0x5d, 0xcb, 0xde, 0xb4, 0xba, 0x9f, 0x00, 0xe2,
	0x2f, 0xf1, 0xab, 0xf8, 0x03, 0x68, 0x76, 0x6d, 0xc7, 0x49, 0x41, 0x70, 0xfd, 0xe6, 0xe7, 0x99,
	0xd9, 0x67, 0x5e, 0x76, 0x66, 0x65, 0x38, 0xca, 0xd7, 0x37, 0x59, 0x99, 0x3c, 0xcd, 0x0b, 0xa5,
	0x15, 0xf1, 0x2d, 0x8a, 0x7c, 0xe8, 0x2c, 0x84, 0x4c, 0xa2, 0x33, 0xe8, 0x2c, 0x94, 0x4c, 0xc8,
	0x67, 0x00, 0x25, 0x2f, 0xde, 0xf0, 0xe2, 0xa5, 0xc8, 0x78, 0xe8, 0x8c, 0x9c, 0xb1, 0x47, 0x5b,
	0x4c, 0x34, 0x83, 0xe1, 0x55, 0x99, 0xcc, 0xe4, 0x6b, 0x75, 0xb3, 0x9c, 0x5c, 0x92, 0x00, 0xbc,
	0xac, 0x4c, 0x8c, 0xdf, 0x80, 0xe2, 0x27, 0xf9, 0x18, 0xba, 0xbc, 0x28, 0x54, 0x11, 0xba, 0x86,
	0xb3, 0x80, 0x10, 0xe8, 0xc4, 0x4c, 0xb3, 0xd0, 0x33, 0xa4, 0xf9, 0x8e, 0x5e, 0xc1, 0xf1, 0x5c,
	0x25, 0x42, 0x1a, 0xb1, 0xcb, 0xc9, 0x92, 0x9c, 0x80, 0x3b, 0x8b, 0x2b, 0x2d, 0x77, 0x16, 0x93,
	0x47, 0xe0, 0x27, 0x2c, 0xe3, 0xb3, 0xb8, 0xd2, 0xaa, 0x10, 0xf9, 0x14, 0x06, 0x36, 0xa3, 0x57,
	0x45, 0x5a, 0x29, 0xee, 0x88, 0xe8, 0x0e, 0x06, 0x46, 0xf6, 0x39, 0xd3, 0xec, 0x1d, 0xc9, 0xc7,
	0xd0, 0x97, 0x62, 0x73, 0xf7, 0x33, 0xcb, 0x78, 0x25, 0xda, 0x60, 0x12, 0x42, 0xef, 0x96, 0xb3,
	0x78, 0x96, 0x25, 0x95, 0x68, 0x0d, 0xd1, 0xc2, 0x36, 0x1b, 0xb5, 0x95, 0x3a, 0xec, 0x8c, 0x9c,
	0xb1, 0x43, 0x6b, 0x18, 0x3d, 0x6b, 0xd7, 0x80, 0x0d, 0xf9, 0x1a, 0x06, 0x69, 0x1d, 0xdd, 0xc4,
	0x1d, 0x4e, 0x3e, 0x7c, 0x5a, 0x75, 0xbe, 0x49, 0x8b, 0xee, 0x7c, 0xa2, 0x33, 0x38, 0x7a, 0xa1,
	0x84, 0xa4, 0x4a, 0x65, 0xa6, 0x09, 0x8f, 0xc0, 0x2f, 0x94, 0xca, 0x9a, 0xac, 0x2b, 0x14, 0xfd,
	0xe5, 0x00, 0x2c, 0x52, 0x76, 0xcf, 0x0b, 0x53, 0xd8, 0x43, 0xe3, 0x90, 0x11, 0x0c, 0x37, 0x4a,
	0x6a, 0x21, 0xb7, 0xfc, 0x57, 0xa5, 0x4d, 0xf1, 0x0e, 0x6d, 0x53, 0xe4, 0x0b, 0x38, 0x9e, 0x95,
	0x53, 0x15, 0x4f, 0x59, 0xb6, 0x4e, 0x85, 0xb4, 0x5d, 0xe8, 0xd3, 0x7d, 0x12, 0x3b, 0xf8, 0x56,
	0xc8, 0xcb, 0xa6, 0x19, 0x5d, 0xda, 0x60, 0x72, 0x06, 0x27, 0x05, 0x2f, 0xb7, 0xa9, 0xbe, 0x16,
	0xf2, 0x4a, 0x49, 0x7e, 0x1f, 0x76, 0x4d, 0x98, 0x03, 0x96, 0x8c, 0xe1, 0xd4, 0x32, 0x73, 0x55,
	0x72, 0xeb, 0xe8, 0x1b, 0xc7, 0x43, 0x3a, 0xfa, 0xcd, 0x81, 0xe3, 0x85, 0xd2, 0x06, 0x34, 0x31,
	0x28, 0x67, 0xf1, 0x8e, 0x31, 0xd5, 0x77, 0xe9, 0x01, 0x8b, 0x31, 0x2e, 0x52, 0xb6, 0xb9, 0x6b,
	0x39, 0xba, 0xc6, 0xf1, 0x90, 0x46, 0xc5, 0xf9, 0x76, 0xcf, 0xd1, 0xb3, 0x8a, 0xfb, 0x6c, 0xf4,
	0x87, 0x03, 0xc3, 0x29, 0xcb, 0xf8, 0xb5, 0x90, 0x73, 0x51, 0x6a, 0x9c, 0x0a, 0x8c, 0x79, 0x2d,
	0x64, 0x95, 0x42, 0x0d, 0xb1, 0x47, 0x26, 0x08, 0x9a, 0x6c, 0xd0, 0x06, 0xe3, 0x29, 0xd4, 0x45,
	0x93, 0x0d, 0x53, 0x43, 0xf2, 0x04, 0xfa, 0x1b, 0x56, 0xc4, 0x2f, 0xef, 0x73, 0x6e, 0x3a, 0x7b,
	0xb2, 0xbb, 0xd1, 0x4b, 0x56, 0xc4, 0x25, 0x1a, 0x68, 0xe3, 0x12, 0xfd, 0xe9, 0x42, 0x1f, 0xa7,
	0xc6, 0xdc, 0xee, 0xbf, 0x4c, 0x0d, 0x99, 0x00, 0xe4, 0x66, 0x68, 0x30, 0xe3, 0xd0, 0x1d, 0x79,
	0xe3, 0xe1, 0x84, 0xd4, 0xaa, 0xbb, 0x71, 0xa2, 0x2d, 0x2f, 0x12, 0xc1, 0x51, 0x82, 0x17, 0xbe,
	0x4e, 0xb9, 0xd9, 0x13, 0xbb, 0x0c, 0x7b, 0x1c, 0x8e, 0x1f, 0x2e, 0xe3, 0x52, 0xb3, 0xe4, 0x9d,
	0x64, 0xa7, 0xb5, 0x81, 0xee, 0x7c, 0xc8, 0x77, 0x70, 0x9c, 0xb7, 0xef, 0xd1, 0x4c, 0xc6, 0x70,
	0xf2, 0x49, 0x93, 0x4b, 0xdb, 0x48, 0xf7, 0x7d, 0x71, 0xe1, 0xeb, 0xb2, 0xcb, 0xd0, 0x1f, 0x79,
	0xe3, 0x2e, 0xdd, 0x11, 0xe4, 0x5b, 0x18, 0x16, 0x0b, 0xa5, 0xab, 0x6b, 0x09, 0x7b, 0xa6, 0xc8,
	0x8f, 0xda, 0xd9, 0x54, 0x26, 0xda, 0xf6, 0x8b, 0xbe, 0x6f, 0x2d, 0x1e, 0x6e, 0xee, 0x57, 0xd0,
	0x2f, 0xaa, 0x76, 0x56, 0x0b, 0x15, 0xd4, 0x1a, 0x75, 0x9b, 0x69, 0xe3, 0x11, 0x9d, 0xc2, 0xf1,
	0x9c, 0xb3, 0x37, 0xbc, 0xde, 0x5b, 0xf3, 0x12, 0x34, 0xc4, 0x7b, 0xbd, 0x04, 0x53, 0x08, 0xec,
	0x8d, 0xfc, 0xb8, 0xd1, 0x42, 0x49, 0xf3, 0x1a, 0x04, 0xe0, 0xad, 0xb9, 0x1d, 0x71, 0x87, 0xe2,
	0x27, 0xf9, 0x1c, 0xbc, 0xbc, 0xda, 0xdf, 0x93, 0xc9, 0x69, 0xab, 0x7d, 0x66, 0x3c, 0xd0, 0x16,
	0x3d, 0x3b, 0x10, 0x7a, 0x78, 0x75, 0x3f, 0x41, 0x70, 0xc5, 0x84, 0xd4, 0xcc, 0xf6, 0xca, 0x28,
	0xec, 0x8f, 0x92, 0xf3, 0x7f, 0x46, 0x29, 0xfa, 0xdd, 0x01, 0xf2, 0x4b, 0xce, 0x25, 0xce, 0x2f,
	0x35, 0xab, 0x6d, 0xa4, 0x1e, 0xdb, 0x49, 0xaf, 0x92, 0xc1, 0xeb, 0x6c, 0xf0, 0xde, 0x16, 0xb8,
	0xff, 0xb9, 0x05, 0xe4, 0x4b, 0xe8, 0xe5, 0xb6, 0x76, 0x33, 0xa7, 0xff, 0xd0, 0x92, 0xda, 0x1e,
	0x5d, 0x00, 0xc1, 0x52, 0x97, 0x5c, 0xeb, 0x94, 0x63, 0xac, 0x87, 0x37, 0xe6, 0xfc, 0x09, 0x0c,
	0x9a, 0xf1, 0x26, 0x3d, 0xf0, 0x56, 0xab, 0x55, 0xf0, 0x01, 0x19, 0x42, 0xef, 0xb9, 0x7a, 0x2b,
	0x2f, 0xb8, 0x0e, 0x1c, 0x02, 0xe0, 0xdb, 0x10, 0x81, 0x7b, 0x1e, 0xc3, 0xa0, 0x49, 0x9a, 0x1c,
	0x41, 0x7f, 0xb5, 0x5a, 0xdd, 0x20, 0x61, 0xcf, 0xcc, 0x39, 0x7b, 0x9d, 0x9a, 0x33, 0x7d, 0xe8,
	0x2c, 0x98, 0x28, 0x02, 0x17, 0x9d, 0x96, 0xba, 0x60, 0x22, 0xb9, 0xd5, 0x81, 0x87, 0x5a, 0x53,
	0x95, 0xc6, 0x5c, 0x06, 0x1d, 0x3c, 0xb0, 0xbc, 0xdd, 0xca, 0x17, 0x42, 0x06, 0x5d, 0x7b, 0x5a,
	0xe5, 0x28, 0xe5, 0x9f, 0xff, 0x00, 0xbd, 0xaa, 0x58, 0xe4, 0x31, 0xc6, 0x42, 0x69, 0x1b, 0x02,
	0x5f, 0x24, 0x04, 0x0e, 0x0a, 0x9b, 0x37, 0x08, 0x91, 0x6b, 0xce, 0x6f, 0x2d, 0xf0, 0xd6, 0xbe,
	0xf9, 0x25, 0xf8, 0xe6, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0xce, 0x68, 0xdc, 0x37, 0x22, 0x08,
	0x00, 0x00,
}
