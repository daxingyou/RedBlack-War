package internal

import (
	"fmt"
	"math/rand"
	"time"
)

func (r *Room) RoomInit() {

	r.RoomId = r.GetRoomNumber()
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
	return int32(len(r.RPotWinList))
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
