package internal

import "time"

//PlayerAction 玩家下注行动
func (r *Room) PlayerAction() {
	//下注阶段定时器
	timer1 := time.NewTicker(time.Second * DownBetTime)

	go func() {
		//遍历所有用户开始下注信息，观战用户也不能进行下注
		for _, v := range r.PlayerList {
			if v != nil && v.Status != WatchGame {
				//获取玩家下注处理
				v.PlayerActionDownBet()
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
	//结算阶段定时器
	timer2 := time.NewTicker(time.Second * DownBetTime)

	r.GameStat = Settle

	// 摊牌
	// 比牌
	// Who Win?
	// 注池结算

	//计时数又重置为0,开始新的下注阶段时间倒计时
	r.RoomStat = RoomStatusOver

	//处理玩家局数
	r.UpdateGamesNum()

	//更新房间列表
	r.UpdatePlayerList()

	select {
	case <-timer2.C:
		r.CompareSettlement()
	}

	//TODO 游戏重新开始
	r.StartGameRun()
}
