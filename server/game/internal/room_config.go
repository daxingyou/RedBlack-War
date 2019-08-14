package internal

import "server/game/card"

type RoomStatus int32

const (
	RoomStatusNone RoomStatus = 1 // 房间等待状态
	RoomStatusRun  RoomStatus = 2 // 房间运行状态
	RoomStatusOver RoomStatus = 3 // 房间结束状态
)

//房间状态 (分为下注阶段、比牌结算阶段)
//主要是针对于玩家中途加入房间，如果是下注阶段，玩家可直接进行下注。比牌结算阶段，玩家则视为观战
type GameStatus int32

const (
	DownBet GameStatus = 1 //下注阶段
	Settle  GameStatus = 2 //比牌结算阶段
)

const (
	RoomCordCount  = 40 //玩家进入房间获取房间的战绩数量。
	RoomLimitMoney = 50 //房间限定金额50,否则处于观战状态
)

type GameWinList struct {
	ReadWin   int32          //红Win为 1
	BlackWin  int32          //黑Win为 1
	LuckWin   int32          //幸运luck为 1
	CardTypes card.CardsType //比牌类型  1 单张,2 对子,3 顺子,4 金花,5 顺金,6 豹子
}

//房间注池数量
type RoomPoolCount struct {
	ReadMoneyCount  int32 //红池金额数量
	BlackMoneyCount int32 //黑池金额数量
	LuckMoneyCount  int32 //Luck金额数量
}

type Room struct {
	RoomId      string    //房间号
	PlayerList  []*Player //玩家列表
	PlayerCount int32     //房间当前人数

	RoomStat      RoomStatus //房间状态
	GameStat      GameStatus //游戏状态
	GodGambleName string     //赌神id

	PoolMoneyCount *RoomPoolCount
	CardTypeList   []int32        //卡牌类型的总集合 1 单张,2 对子,3 顺子,4 金花,5 顺金,6 豹子
	RPotWinList    []*GameWinList //红黑Win、Luck、比牌类型的总集合
	TotalCount     int32          //房间游戏的总局数
}
