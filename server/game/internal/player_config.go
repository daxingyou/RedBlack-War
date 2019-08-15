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

const (
	GamesNumLimit = 72 //玩家获取房间游戏数据上限
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

	DownBetMoney  int32        //玩家本局下注金额
	DownPotType   int32        //玩家本局下注注池  红为 1,黑为 2,幸运为 3
	Status        PlayerStatus //玩家状态
	room          *Room        //玩家房间信息
	IsGodGambling bool         //玩家是否是赌神
	ContinueVot   float64      //续投，记录玩家上局的下注金额。

	WinTotalCount  int32          //玩家房间获胜Win总次数
	PotWinList     []*GameWinList //底池每局Win总列表
	CardTypeList   []int32        //卡牌类型列表
	ReadBlackList  []int32        //每局红黑Win总顺序列表  红为 1,黑为 2
	ReadWinCount   int32          //Win总列表红Win的局数
	BlackWinCount  int32          //Win总列表黑Win的局数
	LuckWinCount   int32          //Win总列表幸运的局数
	TotalAmountBet int32          //玩家总下注
	IsOnline       bool           //玩家是否在线
	//SeatNum        int32          //玩家座位号
}
