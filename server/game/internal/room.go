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
	p.room.PlayerCount++
	p.room = r

	//进入房间玩家是否大于 50金币，否则处于观战状态
	p.PlayerMoneyHandler()

	//获取最新40局游戏数据(小于40局则全部显示出来)
	p.GetRoomCordData(r)

	//更新房间列表
	//TODO 是否必要发送前端更新玩家列表
	r.UpdatePlayerList()

	//游戏开始
	r.GameStartRun()

}

//GameStart 游戏开始
func (r *Room) GameStartRun() {
	//判断房间人数是小于两人
	if r.PlayerLength() < 2 {
		return
	}
	log.Debug("Game Start Running ~")

	//下注阶段,玩家开始下注 15秒

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

	//p.room.PlayerList[p.SeatNum] = nil

	//从房间列表删除玩家信息,更新房间列表
	p.room.PlayerCount--
	for k, v := range r.PlayerList {
		if v != nil && v == p {
			r.PlayerList = append(r.PlayerList[:k], r.PlayerList[k+1:]...)
		}
	}

	r.UpdatePlayerList()
	//TODO 维护房间列表  ( 这里暂且有误
	//maintain := pb_msg.MaintainList_S2C{}
	//var ListSlice []*Player
	//maintain.PlayerList = append(maintain.PlayerList,r.PlayerList...)
	//r.BroadCastExcept(maintain.PlayerList, p)

	//广播其他玩家该玩家退出房间
	leave := pb_msg.LeaveRoom_S2C{}
	leave.LoginData.Id = p.Id
	leave.LoginData.NickName = p.NickName
	leave.LoginData.HeadImg = p.HeadImg
	leave.LoginData.Account = p.Account

	r.BroadCastExcept(leave, p)
	log.Debug("Player Exit from the Room SUCCESS ~")
}

//TODO 处理获取玩家行动
