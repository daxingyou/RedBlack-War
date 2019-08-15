package internal

import (
	"fmt"
	"math/rand"
	"time"
)

func (r *Room) RoomInit() {

	r.RoomId = r.GetRoomNumber()
	r.PlayerList = nil

	r.RoomStat = RoomStatusNone
	//r.clock = time.NewTicker(time.Second)

	r.GodGambleName = ""
	r.CardTypeList = nil
	r.RPotWinList = nil
	r.GameTotalCount = 0
}

func (r *Room) GetRoomNumber() string {
	roomNumber := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
	return roomNumber
}

//BroadCastExcept 向当前玩家之外的玩家广播
func (r *Room) BroadCastExcept(msg interface{}, p *Player) {
	for _, v := range r.PlayerList {
		if v != nil && v != p {
			v.ConnAgent.WriteMsg(msg)
		}
	}
}

//BroadCastMsg 进行广播消息
func (r *Room) BroadCastMsg(msg interface{}) {
	for _, v := range r.PlayerList {
		if v != nil {
			v.ConnAgent.WriteMsg(msg)
		}
	}
}

//PlayerLen 房间当前人数
func (r *Room) PlayerLength() int32 {
	var num int32
	for _, v := range r.PlayerList {
		if v != nil {
			num++
		}
	}
	return num
}

//RoomGameCount 获取房间游戏总数量
func (r *Room) RoomGameCount() int32 {
	return r.GameTotalCount
}

//FindUsableSeat 寻找可用座位
func (r *Room) FindUsableSeat() int32 {
	for k, v := range r.PlayerList {
		if v == nil {
			return int32(k)
		}
	}
	panic("The Room logic was Wrong, don't find able seat, panic err!")
}

//PlayerListSort 玩家列表排序(进入房间、退出房间、重新开始)
func (r *Room) UpdatePlayerList() {

	//先将玩家信息列表置为空
	var PlayerSort []*Player
	var playerSlice []*Player

	for _, v := range r.PlayerList {
		if v != nil && v.TotalAmountBet != 0 {
			playerSlice = append(playerSlice, v)
		}
	}

	var ps []*Player
	for _, v := range playerSlice {
		if v != nil && v.Id == r.GodGambleName {
			PlayerSort = append(PlayerSort, v)
		} else {
			ps = append(ps, v)
		}
	}
	for i := 0; i < len(ps); i++ {
		for j := 1; j < len(ps)-i; j++ {
			if ps[j].TotalAmountBet > ps[j-1].TotalAmountBet {
				//交换
				ps[j], ps[j-1] = ps[j-1], ps[j]
			}

		}
	}

	var ps2 []*Player
	for _, v := range r.PlayerList {
		if v != nil && v.TotalAmountBet == 0 {
			ps2 = append(ps2, v)
		}
	}
	for i := 0; i < len(ps2); i++ {
		for j := 1; j < len(ps2)-i; j++ {
			if ps2[j].TotalAmountBet > ps2[j-1].TotalAmountBet {
				//交换
				ps2[j], ps2[j-1] = ps2[j-1], ps2[j]
			}

		}
	}

	for _, v := range ps {
		if v != nil {
			PlayerSort = append(PlayerSort, v)
		}
	}
	for _, v := range ps2 {
		if v != nil {
			PlayerSort = append(PlayerSort, v)
		}
	}

	//将房间列表置为空,将更新的数据追加到房间列表
	r.PlayerList = nil
	r.PlayerList = append(r.PlayerList, PlayerSort...)
}

//GetGodGableId 获取赌神ID  TODO 每局游戏结束更新赌神
func (r *Room) GetGodGableId() {
	var GodSlice []*Player
	GodSlice = append(GodSlice, r.PlayerList...)

	for i := 0; i < len(GodSlice); i++ {
		for j := 1; j < len(GodSlice)-i; j++ {
			if GodSlice[j].WinTotalCount > GodSlice[j-1].WinTotalCount {
				//交换
				GodSlice[j], GodSlice[j-1] = GodSlice[j-1], GodSlice[j]
			}
		}
	}
	r.GodGambleName = GodSlice[0].Id
}

//GatherRCardType 房间所有卡牌类型集合  ( 这里可以直接每局游戏摊牌 追加牌型类型
func (r *Room) GatherRCardType() {
	for _, v := range r.RPotWinList {
		if v != nil {
			//TODO 这里存在一个问题,卡牌类型是房间的，不是用户的，用户只是截取 40局类型
			r.CardTypeList = append(r.CardTypeList, int32(v.CardTypes))
		}
	}
}

//DisposeGamesNum 处理玩家局数
func (r *Room) UpdateGamesNum() {
	for _, v := range r.PlayerList {
		//玩家局数达到72局，就清空一次玩家房间数据
		if v != nil && v.GetPotWinCount() == GamesNumLimit {
			v.ReadWinCount = 0
			v.BlackWinCount = 0
			v.LuckWinCount = 0
			v.ReadBlackList = nil

			//游戏结束玩家金额不足设为观战
			v.PlayerMoneyHandler()
		}
	}
}
