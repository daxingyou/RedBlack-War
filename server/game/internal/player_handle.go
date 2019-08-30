package internal

import (
	"fmt"
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
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

//GetUserRoomInfo 用户重新登陆，获取房间信息
func (p *Player) GetUserRoomInfo() *Player {
	for _, v := range gameHall.roomList {
		if v != nil {
			for _, pl := range v.PlayerList {
				if pl != nil && pl.Id == p.Id {
					return pl
				}
			}
		}
	}
	return nil
}

//PlayerLoginAgain 用户重新登陆
func PlayerLoginAgain(p *Player, a gate.Agent) {
	fmt.Println("进来了呀~~~")
	p.room = userRoomMap[p.Id]
	for _, v := range p.room.PlayerList {
		if v.Id == p.Id {
			p = v
		}
	}
	p.IsOnline = true
	p.ConnAgent = a
	p.ConnAgent.SetUserData(p)

	//返回前端信息
	//fmt.Println("LoginAgain房间信息:", p.room)
	r := p.room.RspRoomData()
	enter := &pb_msg.EnterRoom_S2C{}
	enter.RoomData = r
	if p.room.GameStat == DownBet {
		enter.GameTime = DownBetTime - p.room.counter
		log.Debug("DownBetTime.GameTime ::: %v", enter.GameTime)
	} else {
		enter.GameTime = SettleTime - p.room.counter
		log.Debug("SettleTime.GameTime ::: %v", enter.GameTime)
	}
	p.SendMsg(enter)

	//更新房间列表
	p.room.UpdatePlayerList()
	maintainList := p.room.PackageRoomPlayerList()
	p.room.BroadCastExcept(maintainList, p)
	log.Debug("用户断线重连成功,返回客户端数据~ ")
}

//PlayerExitRoom 玩家退出房间
func (p *Player) PlayerReqExit() {
	if p.room != nil {
		p.room.ExitFromRoom(p)
	} else {
		log.Debug("Player Exit Room, But not found Player Room ~")
	}
}

//SetAction 设置玩家行动
func (p *Player) SetPlayerAction(m *pb_msg.PlayerAction_C2S) {
	//不是下注阶段不能进行下注
	if p.room.GameStat != DownBet {
		//返回前端信息
		msg := &pb_msg.MsgInfo_S2C{}
		msg.Msg = recodeText[RECODE_NOTDOWNBETSTATUS]
		p.SendMsg(msg)
		log.Debug("当前不是下注阶段,玩家不能行动~")
		return
	}

	//判断玩家金额是否足够下注的金额(这里其实金额不足玩家是不能在进行点击事件的。双重安全!)
	if p.Account < float64(m.DownBet) {
		msg := &pb_msg.MsgInfo_S2C{}
		msg.Error = recodeText[RECODE_NOTDOWNMONEY]
		p.SendMsg(msg)

		log.Debug("玩家金额不足,不能进行下注~")
		return
	}

	p.IsAction = m.IsAction
	//判断玩家是否行动,做相应处理
	if p.IsAction == true {
		//记录玩家在该房间总下注 和 房间注池的总金额
		if m.DownPot == pb_msg.PotType_RedPot {
			p.Account -= float64(m.DownBet)
			p.DownBetMoneys.RedDownBet += m.DownBet
			p.TotalAmountBet += m.DownBet
			p.room.PotMoneyCount.RedMoneyCount += m.DownBet
		}
		if m.DownPot == pb_msg.PotType_BlackPot {
			p.Account -= float64(m.DownBet)
			p.DownBetMoneys.BlackDownBet += m.DownBet
			p.TotalAmountBet += m.DownBet
			p.room.PotMoneyCount.BlackMoneyCount += m.DownBet
		}
		if m.DownPot == pb_msg.PotType_LuckPot {
			p.Account -= float64(m.DownBet)
			p.DownBetMoneys.LuckDownBet += m.DownBet
			p.TotalAmountBet += m.DownBet
			p.room.PotMoneyCount.LuckMoneyCount += m.DownBet
		}
		//记录续投下注的金额对应注池
		p.ContinueVot.DownBetMoneys.RedDownBet = p.DownBetMoneys.RedDownBet
		p.ContinueVot.DownBetMoneys.BlackDownBet = p.DownBetMoneys.BlackDownBet
		p.ContinueVot.DownBetMoneys.LuckDownBet = p.DownBetMoneys.LuckDownBet
		p.ContinueVot.TotalMoneyBet = p.ContinueVot.DownBetMoneys.RedDownBet + p.ContinueVot.DownBetMoneys.BlackDownBet + p.ContinueVot.DownBetMoneys.LuckDownBet
	}

	//返回前端玩家行动,更新玩家最新金额
	action := &pb_msg.PlayerAction_S2C{}
	action.Id = p.Id
	action.DownBet = m.DownBet
	action.DownPot = m.DownPot
	action.IsAction = m.IsAction
	action.Account = p.Account
	p.room.BroadCastMsg(action)
	//p.SendMsg(action)

	//广播玩家注池金额
	pot := &pb_msg.PotTotalMoney_S2C{}
	pot.PotMoneyCount = new(pb_msg.PotMoneyCount)
	pot.PotMoneyCount.RedMoneyCount = p.room.PotMoneyCount.RedMoneyCount
	pot.PotMoneyCount.BlackMoneyCount = p.room.PotMoneyCount.BlackMoneyCount
	pot.PotMoneyCount.LuckMoneyCount = p.room.PotMoneyCount.LuckMoneyCount
	p.room.BroadCastMsg(pot)

	//fmt.Println("玩家:", p.Id, "行动 红、黑、Luck下注: ", p.DownBetMoneys, "玩家总下注金额: ", p.TotalAmountBet)
	//fmt.Println("房间池红、黑、Luck总下注: ", p.room.PotMoneyCount, "续投总额:", p.ContinueVot.TotalMoneyBet)
}
