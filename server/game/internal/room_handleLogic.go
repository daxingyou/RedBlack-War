package internal

import (
	"github.com/name5566/leaf/log"
	pb_msg "server/msg/Protocal"
	"time"
)

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

//GetGodGableId 获取赌神ID
func (r *Room) GetGodGableId() {
	var GodSlice []*Player
	GodSlice = append(GodSlice, r.PlayerList...)

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
			//TODO 这里存在一个问题,卡牌类型是房间的，不是用户的，用户只是截取 40局类型
			r.CardTypeList = append(r.CardTypeList, int32(v.CardTypes))
		}
	}
}

//DisposeGamesNum 处理玩家局数
func (r *Room) UpdateGamesNum() {
	for _, v := range r.PlayerList {
		//玩家局数达到72局，就清空一次玩家房间数据
		if v != nil && v.GetPotWinCount() == GamesNumLimit {
			v.ReadWinCount = 0
			v.BlackWinCount = 0
			v.LuckWinCount = 0
			v.ReadBlackList = nil

			//游戏结束玩家金额不足设为观战
			v.PlayerMoneyHandler()
		}
	}
}

//PackageRoomInfo 封装房间信息
func (r *Room) PackageRoomInfo() *pb_msg.MaintainList_S2C {
	msg := &pb_msg.MaintainList_S2C{}

	for _, v := range r.PlayerList {
		if v != nil {
			data := &pb_msg.PlayerData{}
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
	log.Debug("下注阶段开始倒计时~ : %v", time.Now())

	//记录房间游戏总局数
	r.GameTotalCount++
	r.RoomStat = RoomStatusRun
	r.GameStat = DownBet

	//玩家开始下注
	r.PlayerAction()
}

//PlayerAction 玩家下注行动
func (r *Room) PlayerAction() {
	//下注阶段定时器
	timer1 := time.NewTicker(time.Second * DownBetTime)

	go func() {
		//遍历所有用户开始下注信息，观战用户也不能进行下注
		for _, v := range r.PlayerList {
			if v != nil && v.Status != WatchGame {
				//获取玩家下注处理
				v.ActionHandler()
			}
		}
	}()

	select {
	case <-timer1.C:
		r.CompareSettlement()
	}
}

//CompareSettlement 开始比牌结算
func (r *Room) CompareSettlement() {
	//返回结算阶段倒计时
	msg := &pb_msg.DownBetTime_S2C{}
	msg.StartTime = SettleTime
	r.BroadCastMsg(msg)
	log.Debug("结算阶段开始倒计时~ : %v", time.Now())

	//结算阶段定时器
	timer2 := time.NewTicker(time.Second * SettleTime)

	r.GameStat = Settle

	// 摊牌
	// 比牌
	// Who Win?
	// 注池结算

	//计时数又重置为0,开始新的下注阶段时间倒计时
	r.RoomStat = RoomStatusOver

	//处理玩家局数
	r.UpdateGamesNum()

	//更新房间赌神ID
	r.GetGodGableId()

	//更新房间列表
	r.UpdatePlayerList()
	maintainList := r.PackageRoomInfo()
	r.BroadCastMsg(maintainList)

	select {
	case <-timer2.C:
		//开始新一轮游戏,重复调用StartGameRun函数
		r.StartGameRun()
	}
}
