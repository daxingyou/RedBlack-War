package internal

import (
	"github.com/name5566/leaf/log"
	pb_msg "server/msg/Protocal"
)

//JoinGameRoom 加入游戏房间
func (r *Room) JoinGameRoom(p *Player) {

	//寻找可用的座位号
	//p.SeatNum = r.FindUsableSeat()
	//r.PlayerList[p.SeatNum] = p

	//将用户添加到用户列表
	r.PlayerList = append(r.PlayerList, p)
	p.room = r

	userRoomMap = make(map[string]*Room)
	userRoomMap[p.Id] = r

	//进入房间玩家是否大于 50金币，否则处于观战状态
	p.PlayerMoneyHandler()

	//获取最新40局游戏数据(小于40局则全部显示出来)
	p.GetRoomCordData(r)

	//更新房间列表	todo 这里可以不需要发前端指令，因为加入房间要返回 roomData
	r.UpdatePlayerList()

	//判断房间人数是否小于两人，否则不能开始运行
	if r.PlayerLength() < 2 {
		//房间游戏不能开始,房间设为等待状态
		r.RoomStat = RoomStatusNone

		errMsg := &pb_msg.MsgInfo_S2C{}
		errMsg.Msg = recodeText[RECODE_PEOPLENOTFULL]
		p.ConnAgent.WriteMsg(errMsg)
		log.Debug("房间人数不够，不能开始游戏~")

		//返回前端房间信息
		msg := &pb_msg.JoinRoom_S2C{}
		roomData := p.RspRoomData()
		msg.RoomData = roomData
		p.ConnAgent.WriteMsg(msg)

		return
	}

	//只要不小于两人,就属于游戏状态
	p.Status = PlayGame

	//返回前端房间信息
	msg := &pb_msg.JoinRoom_S2C{}
	roomData := p.RspRoomData()
	msg.RoomData = roomData
	p.ConnAgent.WriteMsg(msg)

	if r.RoomStat != RoomStatusRun {
		// None和Over状态都直接开始运行游戏
		r.StartGameRun()
	} else {
		if r.GameStat == Settle { //这里给前端发送消息 做处理
			msg := &pb_msg.MsgInfo_S2C{}
			msg.Msg = recodeText[RECODE_SELLTENOTDOWNBET]
			p.ConnAgent.WriteMsg(msg)

			log.Debug("当前结算阶段,不能进行操作~")
		}
	}
}

//PlayerExitRoom 玩家退出房间
func (r *Room) PlayerReqExit(p *Player) {
	if p.room != nil {
		r.ExitFromRoom(p)
	} else {
		log.Debug("Player Exit Room, But not found Player Room ~")
	}
}

//ExitFromRoom 从房间退出处理
func (r *Room) ExitFromRoom(p *Player) {

	//从房间列表删除玩家信息,更新房间列表
	for k, v := range r.PlayerList {
		if v != nil && v == p {
			r.PlayerList = append(r.PlayerList[:k], r.PlayerList[k+1:]...)
		}
	}

	//更新房间列表
	r.UpdatePlayerList()
	maintainList := r.PackageRoomInfo()
	r.BroadCastExcept(maintainList, p)

	//广播其他玩家该玩家退出房间
	leave := &pb_msg.LeaveRoom_S2C{}
	leave.PlayerInfo.Id = p.Id
	leave.PlayerInfo.NickName = p.NickName
	leave.PlayerInfo.HeadImg = p.HeadImg
	leave.PlayerInfo.Account = p.Account

	r.BroadCastExcept(leave, p)
	log.Debug("Player Exit from the Room SUCCESS ~")
}
