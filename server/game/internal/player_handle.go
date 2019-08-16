package internal

import (
	"github.com/name5566/leaf/gate"
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
	p.IsOnline = true
	p.room = userRoomMap[p.Id]
	p.ConnAgent = a
	p.ConnAgent.SetUserData(p)
	//返回前端信息
	r := p.RspRoomData()
	enter := &pb_msg.EnterRoom_S2C{}
	enter.RoomData = r
	p.ConnAgent.WriteMsg(enter)
}

//SetAction 设置玩家行动
func (p *Player) SetPlayerAction(m *pb_msg.PlayerAction_C2S) {
	p.DownBetMoneys.ReadDownBet = m.DownBetMoneys.ReadDownBet
	p.DownBetMoneys.BlackDownBet = m.DownBetMoneys.BlackDownBet
	p.DownBetMoneys.LuckDownBet = m.DownBetMoneys.LuckDownBet
	p.DownPotTypes.ReadDownPot = m.DownPotTypes.ReadDownPot
	p.DownPotTypes.BlackDownPot = m.DownPotTypes.BlackDownPot
	p.DownPotTypes.LuckDownPot = m.DownPotTypes.LuckDownPot
	p.IsAction = m.IsAction
}
