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

	//进入房间玩家是否大于 50金币，否则处于观战状态
	p.PlayerMoneyHandler()

	//获取最新40局游戏数据(小于40局则全部显示出来)
	p.GetRoomCordData(r)

	//更新房间列表  TODO 是否必要发送前端更新玩家列表
	r.UpdatePlayerList()

	//房间游戏开始事件
	r.RoomEvent(p)
}

//RoomEvent 房间小于两人，则不能开始游戏
func (r *Room) RoomEvent(p *Player) {
	//判断房间人数是否小于两人，否则不能开始运行
	if r.PlayerLength() < 2 {
		//房间游戏不能开始,房间设为等待状态
		r.RoomStat = RoomStatusNone

		errMsg := pb_msg.ErrMsg_S2C{}
		errMsg.Msg = recodeText[RECODE_PEOPLENOTFULL]
		p.ConnAgent.WriteMsg(errMsg)
		log.Debug("房间人数不够，不能开始游戏~")

		//TODO 返回前端房间信息
		return
	}

	r.StartGameRun(p)
}

//GameStart 游戏开始运行
func (r *Room) StartGameRun(p *Player) {
	log.Debug("~~~~~~~~~~~~ Room Game Start Running ~~~~~~~~~~~~")

	//记录房间游戏总局数
	r.GameTotalCount++

	//开始游戏，两种情况：
	//1、玩家开始进入游戏开始，15秒倒计时下注
	//2、玩家中途加入游戏，截取当前下注倒计时时间
	if r.RoomStat == RoomStatusRun {
		if r.GameStat == DownBet {
			//TODO 截取下注当前倒计时时间

		} else {

		}

	} else { //TODO 首次开始None游戏和重新Over开始游戏,都重新设置房间表时间
		r.RoomStat = RoomStatusRun
		r.GameStat = DownBet
		//设置下注表时间
		r.SetClockTime()

	}

	//玩家开始下注
	r.PlayerAction()
}

//SetClockTime 设置表时间
func (r *Room) SetClockTime() {
	//房间倒计时
	defer r.clock.Stop()
	defer func() { r.counter = 0 }()

	for t := range r.clock.C {
		log.Debug("时间流动 : %v", t)
		r.counter++
		if r.counter == DownBetTime {
			//下注倒计时结束就做比牌结算
			r.CompareSettlement()
			break
		}
	}
}

//CompareSettlement 开始比牌结算
func (r *Room) CompareSettlement() {

}

//PlayerAction 玩家下注行动
func (r *Room) PlayerAction() {
	//遍历所有用户开始下注信息，观战用户也不能进行下注
	//TODO 存在一个问题，用户行动怎么对应 和 遍历获取玩家行动是否会阻塞
	for _, v := range r.PlayerList {
		if v != nil && v.Status != WatchGame {

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

	//p.room.PlayerList[p.SeatNum] = nil

	//从房间列表删除玩家信息,更新房间列表
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
