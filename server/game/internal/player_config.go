package internal

import (
	"github.com/name5566/leaf/gate"
)

//玩家状态 (玩家金额小于50视为观战，玩家中途加入结算阶段，视为观战)
//(观战状态针对于房间内部观战玩家禁止下注。)
type PlayerStatus int32

const (
	PlayGame  PlayerStatus = 1 //游戏状态
	WatchGame PlayerStatus = 2 //观战状态
)

//定义全局用户房间
var userRoomMap map[string]*Room

type Player struct {
	// 玩家代理链接
	ConnAgent gate.Agent
	Index     int32

	Id       string
	NickName string
	HeadImg  string
	Account  float64 // 玩家金额

	Status         PlayerStatus
	room           *Room
	IsGodGambling  bool           //玩家是否是赌神
	ContinueVot    float64        //续投，记录玩家上局的下注金额。
	WinCount       int32          //玩家当前房间Win的局数
	CardTypeList   []int32        //卡牌类型列表
	PotWinList     []*GameWinList //底池每局Win总列表
	ReadBlackList  []int32        //每局红黑Win总顺序列表  红为 1,黑为 2
	ReadWinCount   int32          //Win总列表红Win的局数
	BlackWinCount  int32          //Win总列表黑Win的局数
	LuckWinCount   int32          //Win总列表幸运的局数
	TotalAmountBet int32          //玩家总下注
	IsOnline       bool           //玩家是否在线
	//SeatNum        int32          //玩家座位号
}
