package internal

import (
	"github.com/name5566/leaf/gate"
	"server/game/card"
	pb_msg "server/msg/Protocal"
	"time"
)

//HeartBeatHandle 用户心跳处理
func HeartBeatHandle(a gate.Agent) {

	timer := time.Now().UnixNano() / 1e6

	pong := &pb_msg.Pong{
		ServerTime: timer,
	}
	a.WriteMsg(pong)
}

//PlayerLoginAgain 用户重新登陆
func (p *Player) PlayerLoginAgain(a gate.Agent) {

}

//SetAction 设置玩家行动
func (p *Player) SetPlayerAction() {

}

//PlayerMoneyHandler 玩家进入房间金额处理
func (p *Player) PlayerMoneyHandler() {
	if p.Account < RoomLimitMoney {
		//玩家观战状态不能进行投注
		p.Status = WatchGame
		return
	}
}

//GetPotWinCount 获取玩家在房间内的总局数
func (p *Player) GetPotWinCount() int32 {
	return int32(len(p.PotWinList))
}

//GetPlayerTableData 获取房间战绩数据
func (p *Player) GetRoomCordData(r *Room) {
	//最新40局游戏数据、红黑Win顺序列表、每局Win牌局类型、红黑Luck的总数量

	roomGCount := r.RoomGameCount()
	//判断房间数据是否大于40局
	if roomGCount > RoomCordCount {
		//大于40局则截取最新40局数据
		num := roomGCount - RoomCordCount
		p.PotWinList = append(p.PotWinList, r.RPotWinList[num:]...)
		//TODO 这里 r.CardTypeList 每局游戏摊完牌追加 牌局类型
		p.CardTypeList = append(p.CardTypeList, r.CardTypeList[num:]...)
		for _, v := range p.PotWinList {
			if v.ReadWin == 1 {
				p.ReadWinCount++
				p.ReadBlackList = append(p.ReadBlackList, card.ReadWin)
			}
			if v.BlackWin == 1 {
				p.BlackWinCount++
				p.ReadBlackList = append(p.ReadBlackList, card.BlackWin)
			}
			if v.LuckWin == 1 {
				p.LuckWinCount++
			}
		}
	} else {
		//小于40局则截取全部房间数据
		p.PotWinList = append(p.PotWinList, r.RPotWinList...)
		p.CardTypeList = append(p.CardTypeList, r.CardTypeList...)
		for _, v := range p.PotWinList {
			if v.ReadWin == 1 {
				p.ReadWinCount++
			}
			if v.BlackWin == 1 {
				p.BlackWinCount++
			}
			if v.LuckWin == 1 {
				p.LuckWinCount++
			}
		}
	}
}
