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
func (p *Player) PlayerLoginAgain(a gate.Agent) {
	p.IsOnline = true
	p.room = userRoomMap[p.Id]
	p.ConnAgent = a
	p.ConnAgent.SetUserData(p)
	//返回前端信息
	r := p.room.RspRoomData()
	enter := &pb_msg.EnterRoom_S2C{}
	enter.RoomData = r
	p.ConnAgent.WriteMsg(enter)
	log.Debug("用户断线重连成功,返回客户端数据~")
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
	//if p.room.GameStat != DownBet {
	//	//返回前端信息
	//	msg := &pb_msg.MsgInfo_S2C{}
	//	msg.Msg = recodeText[RECODE_NOTDOWNBETSTATUS]
	//	p.ConnAgent.WriteMsg(msg)
	//	log.Debug("当前不是下注阶段,玩家不能行动~")
	//	return
	//}

	p.DownBetMoneys = new(DownBetMoney)
	p.DownBetMoneys.RedDownBet = m.DownBetMoneys.RedDownBet
	p.DownBetMoneys.BlackDownBet = m.DownBetMoneys.BlackDownBet
	p.DownBetMoneys.LuckDownBet = m.DownBetMoneys.LuckDownBet
	p.DownPotTypes = new(DownPotType)
	p.DownPotTypes.RedDownPot = m.DownPotTypes.RedDownPot
	p.DownPotTypes.BlackDownPot = m.DownPotTypes.BlackDownPot
	p.DownPotTypes.LuckDownPot = m.DownPotTypes.LuckDownPot
	p.IsAction = m.IsAction

	fmt.Println("玩家下注数据: :: ", p.Id, p.DownBetMoneys, p.DownPotTypes)
	fmt.Println("玩家下注数据: :: ", p.room)
}
