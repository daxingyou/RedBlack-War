package internal

import (
	"github.com/name5566/leaf/log"
	"server/game/card"
	pb_msg "server/msg/Protocal"
)

//PlayerMoneyHandler 玩家进入房间金额处理
func (p *Player) PlayerMoneyHandler() {
	if p.Account < RoomLimitMoney {
		//玩家观战状态不能进行投注
		p.Status = WatchGame

		errMsg := &pb_msg.MsgInfo_S2C{}
		errMsg.Msg = recodeText[RECODE_MONEYNOTFULL]
		p.ConnAgent.WriteMsg(errMsg)

		log.Debug("玩家金额不足,设为观战~")
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

//RspRoomData 返回房间信息
func (p *Player) RspRoomData() *pb_msg.RoomData {
	r := &pb_msg.RoomData{}
	r.RoomId = p.room.RoomId

	for _, v := range p.room.PlayerList {
		if v != nil {
			data := &pb_msg.PlayerData{}
			data.PlayerInfo = new(pb_msg.PlayerInfo)
			data.PlayerInfo.Id = v.Id
			data.PlayerInfo.NickName = v.NickName
			data.PlayerInfo.HeadImg = v.HeadImg
			data.PlayerInfo.Account = v.Account
			data.Status = pb_msg.PlayerStatus(v.Status)
			data.IsGodGambling = v.IsGodGambling
			data.ContinueVot = new(pb_msg.ContinueBet)
			data.ContinueVot.DownBetMoneys = new(pb_msg.DownBetMoney)
			data.ContinueVot.DownBetMoneys.ReadDownBet = v.ContinueVot.DownBetMoneys.ReadDownBet
			data.ContinueVot.DownBetMoneys.BlackDownBet = v.ContinueVot.DownBetMoneys.BlackDownBet
			data.ContinueVot.DownBetMoneys.LuckDownBet = v.ContinueVot.DownBetMoneys.LuckDownBet
			data.ContinueVot.TotalMoneyBet = v.ContinueVot.TotalMoneyBet
			data.ResultWinMoney = v.ResultWinMoney
			data.ResultLoseMoney = v.ResultLoseMoney
			data.WinTotalCount = v.WinTotalCount
			data.CardTypeList = v.CardTypeList
			for _, val := range v.PotWinList {
				pot := &pb_msg.PotWinList{}
				pot.ReadWin = val.ReadWin
				pot.BlackWin = val.BlackWin
				pot.LuckWin = val.LuckWin
				pot.CardType = pb_msg.CardsType(val.CardTypes)
				data.PotWinList = append(data.PotWinList, pot)
			}
			data.ReadBlackList = v.ReadBlackList
			data.ReadWinCount = v.ReadWinCount
			data.BlackWinCount = v.BlackWinCount
			data.LuckWinCount = v.LuckWinCount
			data.TotalAmountBet = v.TotalAmountBet
			data.IsOnline = v.IsOnline
			r.PlayerList = append(r.PlayerList, data)
		}
	}
	r.GodGableName = p.room.GodGambleName
	r.GameStage = pb_msg.GameStage(p.room.GameStat)
	r.PotMoneyCount = new(pb_msg.PotMoneyCount)
	r.PotMoneyCount.ReadMoneyCount = p.room.PotMoneyCount.ReadMoneyCount
	r.PotMoneyCount.BlackMoneyCount = p.room.PotMoneyCount.BlackMoneyCount
	r.PotMoneyCount.LuckMoneyCount = p.room.PotMoneyCount.LuckMoneyCount
	r.CardTypeList = p.room.CardTypeList
	for _, value := range p.room.RPotWinList {
		pot := &pb_msg.PotWinList{}
		value.ReadWin = pot.ReadWin
		value.BlackWin = pot.BlackWin
		value.LuckWin = pot.LuckWin
		value.CardTypes = card.CardsType(pot.CardType)
		r.RPotWinList = append(r.RPotWinList, pot)
	}
	return r
}

//PlayerActionDownBet 玩家行动下注
func (p *Player) ActionHandler() {
	//判断玩家是否行动,做相应处理
	if p.IsAction == true {
		//记录玩家在该房间总下注 和 房间注池的总金额
		if p.DownPotTypes.ReadDownPot == true {
			p.TotalAmountBet += p.DownBetMoneys.ReadDownBet
			p.room.PotMoneyCount.ReadMoneyCount += p.DownBetMoneys.ReadDownBet
		}
		if p.DownPotTypes.ReadDownPot == true {
			p.TotalAmountBet += p.DownBetMoneys.BlackDownBet
			p.room.PotMoneyCount.BlackMoneyCount += p.DownBetMoneys.BlackDownBet

		}
		if p.DownPotTypes.ReadDownPot == true {
			p.TotalAmountBet += p.DownBetMoneys.LuckDownBet
			p.room.PotMoneyCount.LuckMoneyCount += p.DownBetMoneys.LuckDownBet
		}

		//记录续投下注的金额对应注池
		p.ContinueVot.DownBetMoneys.ReadDownBet = p.DownBetMoneys.ReadDownBet
		p.ContinueVot.DownBetMoneys.BlackDownBet = p.DownBetMoneys.BlackDownBet
		p.ContinueVot.DownBetMoneys.LuckDownBet = p.DownBetMoneys.LuckDownBet
		p.ContinueVot.TotalMoneyBet = p.DownBetMoneys.ReadDownBet + p.DownBetMoneys.BlackDownBet + p.DownBetMoneys.LuckDownBet

	}
}
