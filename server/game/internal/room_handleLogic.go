package internal

import (
	"fmt"
	"github.com/name5566/leaf/log"
	pb_msg "server/msg/Protocal"
	"time"
)

//BroadCastExcept 向当前玩家之外的玩家广播
func (r *Room) BroadCastExcept(msg interface{}, p *Player) {
	for _, v := range r.PlayerList {
		if v != nil && v.Id != p.Id {
			v.SendMsg(msg)
		}
	}
}

//BroadCastMsg 进行广播消息
func (r *Room) BroadCastMsg(msg interface{}) {
	for _, v := range r.PlayerList {
		if v != nil {
			v.SendMsg(msg)
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
	//首先
	//临时切片
	var playerSlice []*Player
	//1、赌神
	for _, v := range r.PlayerList {
		if v != nil && v.Id == r.GodGambleName {
			playerSlice = append(playerSlice, v)
		}
	}
	//2、玩家下注总金额
	var p1 []*Player //所有下注过的用户
	var p2 []*Player //所有下注金额为0的用户
	for _, v := range r.PlayerList {
		if v != nil && v.TotalAmountBet != 0 {
			p1 = append(p1, v)
		} else {
			p2 = append(p2, v)
		}
	}
	//根据玩家总下注进行排序
	for i := 0; i < len(p1); i++ {
		for j := 1; j < len(p1)-i; j++ {
			if p1[j].TotalAmountBet > p1[j-1].TotalAmountBet {
				//交换
				p1[j], p1[j-1] = p1[j-1], p1[j]
			}
		}
	}
	//将用户总下注金额顺序追加到临时切片
	playerSlice = append(playerSlice, p1...)
	//3、玩家金额,总下注为0,按用户金额排序
	for i := 0; i < len(p2); i++ {
		for j := 1; j < len(p2)-i; j++ {
			if p2[j].Account > p2[j-1].Account {
				//交换
				p2[j], p2[j-1] = p2[j-1], p2[j]
			}
		}
	}
	//将用户余额排序追加到临时切片
	playerSlice = append(playerSlice, p2...)

	//将房间列表置为空,将更新的数据追加到房间列表
	r.PlayerList = nil
	r.PlayerList = append(r.PlayerList, playerSlice...)
}

//GetGodGableId 获取赌神ID
func (r *Room) GetGodGableId() {
	var GodSlice []*Player
	GodSlice = append(GodSlice, r.PlayerList...)

	var WinCount []*Player
	for _, v := range GodSlice {
		if v != nil && v.WinTotalCount != 0 {
			WinCount = append(WinCount, v)
		}
	}
	if len(WinCount) == 0 {
		log.Debug("---------- 没有获取到赌神 ~")
		return
	}

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

//GatherRCardType 房间所有卡牌类型集合  ( 这里可以直接每局游戏摊牌 追加牌型类型 (这里可以不需要这个函数)
func (r *Room) GatherRCardType() {
	for _, v := range r.RPotWinList {
		if v != nil {
			// 这里存在一个问题,卡牌类型是房间的，不是用户的，用户只是截取 40局类型
			r.CardTypeList = append(r.CardTypeList, int32(v.CardTypes))
		}
	}
}

//UpdateGamesNum 更新玩家局数
func (r *Room) UpdateGamesNum() {
	for _, v := range r.PlayerList {
		//玩家局数达到72局，就清空一次玩家房间数据
		if v != nil && v.GetPotWinCount() == GamesNumLimit {
			v.RedWinCount = 0
			v.BlackWinCount = 0
			v.LuckWinCount = 0

			v.PotWinList = nil
			v.RedBlackList = nil

			//游戏结束玩家金额不足设为观战
			v.PlayerMoneyHandler()
		}
	}
}

//PackageRoomInfo 封装房间信息
func (r *Room) PackageRoomPlayerList() *pb_msg.MaintainList_S2C {
	msg := &pb_msg.MaintainList_S2C{}

	for _, v := range r.PlayerList {
		if v != nil {
			data := &pb_msg.PlayerData{}
			data.PlayerInfo = new(pb_msg.PlayerInfo)
			data.PlayerInfo.Id = v.Id
			data.PlayerInfo.NickName = v.NickName
			data.PlayerInfo.HeadImg = v.HeadImg
			data.PlayerInfo.Account = v.Account
			data.DownBetMoneys = new(pb_msg.DownBetMoney)
			data.DownBetMoneys.RedDownBet = v.DownBetMoneys.RedDownBet
			data.DownBetMoneys.BlackDownBet = v.DownBetMoneys.BlackDownBet
			data.DownBetMoneys.LuckDownBet = v.DownBetMoneys.LuckDownBet
			data.TotalAmountBet = v.TotalAmountBet
			data.Status = pb_msg.PlayerStatus(v.Status)
			data.ContinueVot = new(pb_msg.ContinueBet)
			data.ContinueVot.DownBetMoneys = new(pb_msg.DownBetMoney)
			data.ContinueVot.DownBetMoneys.RedDownBet = v.ContinueVot.DownBetMoneys.RedDownBet
			data.ContinueVot.DownBetMoneys.BlackDownBet = v.ContinueVot.DownBetMoneys.BlackDownBet
			data.ContinueVot.DownBetMoneys.LuckDownBet = v.ContinueVot.DownBetMoneys.LuckDownBet
			data.ContinueVot.TotalMoneyBet = v.ContinueVot.TotalMoneyBet
			data.ResultMoney = v.ResultMoney
			data.WinTotalCount = v.WinTotalCount
			data.CardTypeList = v.CardTypeList
			for _, val := range v.PotWinList {
				pot := &pb_msg.PotWinList{}
				pot.RedWin = val.RedWin
				pot.BlackWin = val.BlackWin
				pot.LuckWin = val.LuckWin
				pot.CardType = pb_msg.CardsType(val.CardTypes)
				data.PotWinList = append(data.PotWinList, pot)
			}
			data.RedBlackList = v.RedBlackList
			data.RedWinCount = v.RedWinCount
			data.BlackWinCount = v.BlackWinCount
			data.LuckWinCount = v.LuckWinCount
			data.IsOnline = v.IsOnline
			msg.PlayerList = append(msg.PlayerList, data)
		}
	}
	return msg
}

//GameStart 游戏开始运行
func (r *Room) StartGameRun() {
	//重新开始也要判断房间是否小于两人
	if r.PlayerLength() < 2 {
		//房间游戏不能开始,房间设为等待状态
		r.RoomStat = RoomStatusNone
		msg := &pb_msg.MsgInfo_S2C{}
		msg.Msg = recodeText[RECODE_PEOPLENOTFULL]
		r.BroadCastMsg(msg)

		log.Debug("房间人数不够，不能重新开始游戏~")
		return
	}

	log.Debug("~~~~~~~~~~~~ Room Game Start Running ~~~~~~~~~~~~")
	//返回下注阶段倒计时
	msg := &pb_msg.DownBetTime_S2C{}
	msg.StartTime = DownBetTime
	r.BroadCastMsg(msg)

	log.Debug("~~~~~~~~ 下注阶段 Start : %v", time.Now().Format("2006.01.02 15:04:05")+" ~~~~~~~~")

	//记录房间游戏总局数
	r.GameTotalCount++
	r.RoomStat = RoomStatusRun
	r.GameStat = DownBet

	//下注阶段定时任务
	r.DownBetTimerTask()

	//机器人进行下注 todo
	//r.RobotsDownBet()

	//开始发牌,这里开始计算牌型盈余池。如果亏损就换牌
	//RBdzPk()

	//结算阶段定时任务
	r.SettlerTimerTask()

}

//TimerTask 下注阶段定时器任务
func (r *Room) DownBetTimerTask() {
	//go func() {
	//	//下注阶段定时器
	//	timer := time.NewTicker(time.Second * DownBetTime)
	//	select {
	//	case <-timer.C:
	//		DownBetChannel <- true
	//		return
	//	}
	//}()

	go func() {
		for range r.clock.C {
			r.counter++
			log.Debug("clock : %v ", r.counter)
			if r.counter == DownBetTime {
				r.counter = 0
				DownBetChannel <- true
				return
			}
		}
	}()
}

//TimerTask 结算阶段定时器任务
func (r *Room) SettlerTimerTask() {
	go func() {
		select {
		case t := <-DownBetChannel:
			if t == true {
				//开始比牌结算任务
				r.CompareSettlement()

				//开始新一轮游戏,重复调用StartGameRun函数
				r.StartGameRun()
				return
			}
		}
	}()
}

//PlayerAction 玩家游戏结算
func (r *Room) GameCheckout() {
	//遍历所有用户开始下注信息，观战用户也不能进行下注
	for _, v := range r.PlayerList {
		if v != nil && v.Status != WatchGame {
			//获取玩家下注处理
			v.ActionHandler()
		}
	}
}

//CompareSettlement 开始比牌结算
func (r *Room) CompareSettlement() {

	//返回结算阶段倒计时
	msg := &pb_msg.SettlerTime_S2C{}
	msg.StartTime = SettleTime
	r.BroadCastMsg(msg)

	log.Debug("~~~~~~~~ 结算阶段 Start : %v", time.Now().Format("2006.01.02 15:04:05")+" ~~~~~~~~")

	var count int32
	t := time.NewTicker(time.Second)

	//开始发牌,这里开始计算牌型盈余池。如果亏损就换牌
	RBdzPk()

	//玩家游戏结算  todo
	r.GameCheckout()

	r.GameStat = Settle

	// 摊牌,要在摊牌之前发牌,做盈余池计算,可以进行换牌

	// 比牌
	// Who Win?
	// 注池结算

	//计时数又重置为0,开始新的下注阶段时间倒计时
	r.RoomStat = RoomStatusOver

	//处理玩家局数 和 玩家金额
	r.UpdateGamesNum()

	//清空玩家数据,开始下一句游戏
	r.CleanPlayerData()

	//更新房间赌神ID
	r.GetGodGableId()

	//更新房间列表
	r.UpdatePlayerList()
	maintainList := r.PackageRoomPlayerList()
	r.BroadCastMsg(maintainList)

	//踢出房间断线玩家
	r.KickOutPlayer()

	//todo 这里会发送前端房间数据，前端做处理

	//测试，打印数据
	r.PrintPlayerList()

	for range t.C {
		count++
		log.Debug("clock : %v ", count)
		if count == SettleTime {
			count = 0
			return
		}
	}
}

//KickOutPlayer 踢出房间断线玩家
func (r *Room) KickOutPlayer() {
	for _, v := range r.PlayerList {
		if v != nil && v.IsOnline == false {
			v.PlayerReqExit()
			log.Debug("踢出房间断线玩家 : %v", v.Id)
		}
	}
}

//CleanPlayerData 清空玩家数据,开始下一句游戏
func (r *Room) CleanPlayerData() {
	fmt.Println("进来了~~~")
	for _, v := range r.PlayerList {
		if v != nil {
			v.DownBetMoneys = new(DownBetMoney)
			v.IsAction = false
		}
	}
	for _, v := range r.PlayerList {
		if v != nil && v.IsRobot == true {
			if v.Account < RoomLimitMoney {
				//退出一个机器人就在创建一个机器人
				log.Debug("删除机器人！~~~~~~~~~~~~~~~~~~~~~~~~~")
				v.PlayerReqExit()

				robot := gRobotCenter.CreateRobot()
				r.JoinGameRoom(robot)
			}
		}
	}
}

//看数据用,为了打印房间玩家列表
func (r *Room) PrintPlayerList() {
	for _, v := range r.PlayerList {
		if v != nil {
			fmt.Println("玩家ID ：", v.Id, "金额 :", v.Account, "下注总金额 :", v.TotalAmountBet)
			//fmt.Println("玩家:", v.Id, "行动 红、黑、Luck下注: ", v.DownBetMoneys, "玩家总下注金额: ", v.TotalAmountBet)
			//fmt.Println("房间池红、黑、Luck总下注: ", v.room.PotMoneyCount, "续投总额:", v.ContinueVot.TotalMoneyBet)
		}
	}
}
