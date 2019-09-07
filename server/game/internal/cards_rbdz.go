package internal

import (
	"github.com/name5566/leaf/log"
	pb_msg "server/msg/Protocal"
	"time"
)

// 红黑大战
// 游戏玩法：
// 游戏使用1副扑克牌，无大小王
// 红黑各派3张牌

//卡牌类型
type CardsType int32

const (
	Leaflet  CardsType = 1 //单张
	Pair     CardsType = 2 //对子
	Straight CardsType = 3 //顺子
	Golden   CardsType = 4 //金花
	Shunjin  CardsType = 5 //顺金
	Leopard  CardsType = 6 //豹子
)

const (
	RedWin   = 1 //红Win为 1
	BlackWin = 2 //黑Win为 2
)

// 0:红赢，1赔1，和 黑全输
// 1:黑赢，1赔1，和 红全输
const (
	WinLeopard  int32 = 10 //豹子10倍
	WinShunjin  int32 = 5  //顺金5倍
	WinGolden   int32 = 3  //金花3倍
	WinStraight int32 = 2  //顺子2倍
	WinBigPair  int32 = 1  //大对子(9-A)
	//WinRedBlack int32 = 1  //红黑赢倍数
)

type RBdzDealer struct {
	Poker  []byte //所有的牌
	Offset int    //牌的位置
}

var (
	dealer = NewGoldenFlowerDealer(true)
)

var (
	aCard []byte
	bCard []byte
)

func (this *RBdzDealer) Deal() ([]byte, []byte) {
	// 检查剩余牌数量
	offset := this.Offset
	if offset >= len(this.Poker)/2 {
		//获取牌值
		this.Poker = NewPoker(1, false, true)
		offset = 0
	}
	// 红黑各取3张牌
	a := this.Poker[offset : offset+3]
	b := this.Poker[offset+3 : offset+6]

	return a, b
}

//获取牌型并比牌
func (r *Room) RBdzPk(a []byte, b []byte) {
	ha := Hex(a)
	log.Debug("花牌 数据Red~ : %v", ha)
	hb := Hex(b)
	log.Debug("花牌 数据Black~ : %v", hb)

	//红黑池牌型赋值
	r.Cards.ReadCard = HexInt(a)
	r.Cards.BlackCard = HexInt(b)

	//字符串牌型
	note := PokerArrayString(a) + " | " + PokerArrayString(b)
	log.Debug("花牌 牌型~ : %v", note)

	// 可下注的选项数量(0:红赢,1:黑赢,2:幸运一击)
	ag := dealer.GetGroup(a)
	bg := dealer.GetGroup(b)

	var hallCard int32
	var hallRBWin int32

	//获取牌型处理
	if ag.IsThreeKind() {
		r.Cards.RedType = CardsType(Leopard)
		hallCard = int32(Leopard)
		log.Debug("Red 三同10倍")
	}
	if bg.IsThreeKind() {
		r.Cards.BlackType = CardsType(Leopard)
		hallCard = int32(Leopard)
		log.Debug("Black 三同10倍")
	}
	if ag.IsStraightFlush() {
		r.Cards.RedType = CardsType(Shunjin)
		hallCard = int32(Shunjin)
		log.Debug("Red 顺金5倍")
	}
	if bg.IsStraightFlush() {
		r.Cards.BlackType = CardsType(Shunjin)
		hallCard = int32(Shunjin)
		log.Debug("Black 顺金5倍")
	}
	if ag.IsFlush() {
		r.Cards.RedType = CardsType(Golden)
		hallCard = int32(Golden)
		log.Debug("Red 金花3倍")
	}
	if bg.IsFlush() {
		r.Cards.BlackType = CardsType(Golden)
		hallCard = int32(Golden)
		log.Debug("Black 金花3倍")
	}
	if ag.IsStraight() {
		r.Cards.RedType = CardsType(Straight)
		hallCard = int32(Straight)
		log.Debug("Red 顺子2倍")
	}
	if bg.IsStraight() {
		r.Cards.BlackType = CardsType(Straight)
		hallCard = int32(Straight)
		log.Debug("Black 顺子2倍")
	}
	if (ag.Key.Pair() >> 8) >= 9 {
		r.Cards.RedType = CardsType(Pair)
		hallCard = int32(Pair)
		log.Debug("Red 大对子(9-A)")
	} else if ag.IsPair() {
		r.Cards.RedType = CardsType(Pair)
		hallCard = int32(Pair)
		log.Debug("Red 小对子(2-8)")
	}
	if (bg.Key.Pair() >> 8) >= 9 {
		r.Cards.BlackType = CardsType(Pair)
		hallCard = int32(Pair)
		log.Debug("Black 大对子(9-A)")
	} else if bg.IsPair() {
		r.Cards.BlackType = CardsType(Pair)
		hallCard = int32(Pair)
		log.Debug("Black 小对子(2-8)")
	}
	if ag.IsZilch() {
		r.Cards.RedType = CardsType(Leaflet)
		hallCard = int32(Leaflet)
		log.Debug("Red 单张")
	}
	if bg.IsZilch() {
		r.Cards.BlackType = CardsType(Leaflet)
		hallCard = int32(Leaflet)
		log.Debug("Black 单张")
	}

	log.Debug("Cards Data :%v", r.Cards)

	log.Debug("<-------- 更新盈余池金额为Pre: %v --------->", SurplusPool)

	gw := &GameWinList{}

	res := &pb_msg.OpenCardResult_S2C{}
	res.PotWinTypes = new(pb_msg.DownPotType)
	res.RedCard = r.Cards.ReadCard
	res.BlackCard = r.Cards.BlackCard

	//获取Pot池Win
	if ag.Weight > bg.Weight { //redWin
		log.Debug("Red Win ~")
		gw.RedWin = 1
		hallRBWin = int32(RedWin)
		res.PotWinTypes.RedDownPot = true

		if ag.IsThreeKind() {
			r.Cards.LuckType = CardsType(Leopard)
			gw.LuckWin = 1
			gw.CardTypes = Leopard
			r.CardTypeList = append(r.CardTypeList, int32(Leopard))
			res.PotWinTypes.LuckDownPot = true
			res.RedType = pb_msg.CardsType(Leopard)
		}
		if ag.IsStraightFlush() {
			r.Cards.LuckType = CardsType(Shunjin)
			gw.LuckWin = 1
			gw.CardTypes = Shunjin
			r.CardTypeList = append(r.CardTypeList, int32(Shunjin))
			res.PotWinTypes.LuckDownPot = true
			res.RedType = pb_msg.CardsType(Shunjin)
		}
		if ag.IsFlush() {
			r.Cards.LuckType = CardsType(Golden)
			gw.LuckWin = 1
			gw.CardTypes = Golden
			r.CardTypeList = append(r.CardTypeList, int32(Golden))
			res.PotWinTypes.LuckDownPot = true
			res.RedType = pb_msg.CardsType(Golden)
		}
		if ag.IsStraight() {
			r.Cards.LuckType = CardsType(Straight)
			gw.LuckWin = 1
			gw.CardTypes = Straight
			r.CardTypeList = append(r.CardTypeList, int32(Straight))
			res.PotWinTypes.LuckDownPot = true
			res.RedType = pb_msg.CardsType(Straight)
		}
		if (ag.Key.Pair() >> 8) >= 9 {
			r.Cards.LuckType = CardsType(Pair)
			gw.LuckWin = 1
			gw.CardTypes = Pair
			r.CardTypeList = append(r.CardTypeList, int32(Pair))
			res.PotWinTypes.LuckDownPot = true
			res.RedType = pb_msg.CardsType(Pair)
		} else if ag.IsPair() {
			gw.CardTypes = Pair
			r.CardTypeList = append(r.CardTypeList, int32(Pair))
			res.RedType = pb_msg.CardsType(Pair)
		}
		if ag.IsZilch() {
			gw.CardTypes = Leaflet
			r.CardTypeList = append(r.CardTypeList, int32(Leaflet))
			res.RedType = pb_msg.CardsType(Leaflet)
		}

		for _, v := range r.PlayerList {
			var taxMoney float64
			var totalWinMoney float64
			var totalLoseMoney float64

			if v != nil && v.IsAction == true {
				if v.IsRobot == false {
					totalWinMoney += float64(v.DownBetMoneys.RedDownBet)
					taxMoney += float64(v.DownBetMoneys.RedDownBet)

					totalLoseMoney += float64(v.DownBetMoneys.RedDownBet)
					totalLoseMoney += float64(v.DownBetMoneys.BlackDownBet)
					totalLoseMoney += float64(v.DownBetMoneys.LuckDownBet)

					if gw.LuckWin == 1 {
						if gw.CardTypes == Leopard {
							totalWinMoney += float64(v.DownBetMoneys.LuckDownBet)
							taxMoney += float64(v.DownBetMoneys.LuckDownBet * WinLeopard)
						}
						if gw.CardTypes == Shunjin {
							totalWinMoney += float64(v.DownBetMoneys.LuckDownBet)
							taxMoney += float64(v.DownBetMoneys.LuckDownBet * WinShunjin)
						}
						if gw.CardTypes == Golden {
							totalWinMoney += float64(v.DownBetMoneys.LuckDownBet)
							taxMoney += float64(v.DownBetMoneys.LuckDownBet * WinGolden)
						}
						if gw.CardTypes == Straight {
							totalWinMoney += float64(v.DownBetMoneys.LuckDownBet)
							taxMoney += float64(v.DownBetMoneys.LuckDownBet * WinStraight)
						}
						if gw.CardTypes == Pair {
							totalWinMoney += float64(v.DownBetMoneys.LuckDownBet)
							taxMoney += float64(v.DownBetMoneys.LuckDownBet * WinBigPair)
						}
					}
					tax := taxMoney * taxRate
					v.ResultMoney = totalWinMoney + taxMoney - tax
					v.Account += v.ResultMoney
					v.ResultMoney -= totalLoseMoney
					if v.ResultMoney > 0 {
						//将玩家的税收金额添加到盈余池
						SurplusPool += tax
						SurplusPool -= v.ResultMoney //盈余池结算减去玩家Win金额
						v.WinTotalCount++

						v.WinResultMoney = v.ResultMoney
						log.Debug("玩家金额: %v, 进来了Win: %v", v.Account, v.WinResultMoney)

						timeStr := time.Now().Format("2006-01-02_15:04:05")
						nowTime := time.Now().Unix()
						reason := "ResultWinScore"

						//同时同步赢分和输分
						c4c.UserSyncWinScore(v, nowTime, timeStr, reason)
					} else if v.ResultMoney < 0 {
						v.LoseResultMoney = v.ResultMoney
						log.Debug("玩家金额: %v, 进来了Lose: %v", v.Account, v.LoseResultMoney)

						timeStr := time.Now().Format("2006-01-02_15:04:05")
						nowTime := time.Now().Unix()
						reason := "ResultLoseScore"

						//同时同步赢分和输分
						c4c.UserSyncLoseScore(v, nowTime, timeStr, reason)

						//将玩家输的金额添加到盈余池
						SurplusPool -= v.ResultMoney //这个Res是负数 负负得正
					}
					log.Debug("<<===== 玩家下注: %v, 结算: %v =====>>", v.DownBetMoneys, v.ResultMoney)

				} else {
					totalWinMoney += float64(v.DownBetMoneys.RedDownBet)
					taxMoney += float64(v.DownBetMoneys.RedDownBet)

					totalLoseMoney += float64(v.DownBetMoneys.RedDownBet)
					totalLoseMoney += float64(v.DownBetMoneys.BlackDownBet)
					totalLoseMoney += float64(v.DownBetMoneys.LuckDownBet)
					if gw.LuckWin == 1 {
						if gw.CardTypes == Leopard {
							totalWinMoney += float64(v.DownBetMoneys.LuckDownBet)
							taxMoney += float64(v.DownBetMoneys.LuckDownBet * WinLeopard)
						}
						if gw.CardTypes == Shunjin {
							totalWinMoney += float64(v.DownBetMoneys.LuckDownBet)
							taxMoney += float64(v.DownBetMoneys.LuckDownBet * WinShunjin)
						}
						if gw.CardTypes == Golden {
							totalWinMoney += float64(v.DownBetMoneys.LuckDownBet)
							taxMoney += float64(v.DownBetMoneys.LuckDownBet * WinGolden)
						}
						if gw.CardTypes == Straight {
							totalWinMoney += float64(v.DownBetMoneys.LuckDownBet)
							taxMoney += float64(v.DownBetMoneys.LuckDownBet * WinStraight)
						}
						if gw.CardTypes == Pair {
							totalWinMoney += float64(v.DownBetMoneys.LuckDownBet)
							taxMoney += float64(v.DownBetMoneys.LuckDownBet * WinBigPair)
						}
					}
					tax := taxMoney * taxRate
					v.ResultMoney = totalWinMoney + taxMoney - tax
					v.Account += v.ResultMoney
					v.ResultMoney -= totalLoseMoney
					if v.ResultMoney > 0 {
						v.WinTotalCount++
					}
					log.Debug("<----- 机器人下注: %v, 结算: %v ----->", v.DownBetMoneys, v.ResultMoney)
				}
			}
		}
	} else if ag.Weight < bg.Weight { //blackWin
		log.Debug("Black Win ~")
		gw.BlackWin = 1
		hallRBWin = int32(BlackWin)
		res.PotWinTypes.BlackDownPot = true

		if bg.IsThreeKind() {
			r.Cards.LuckType = CardsType(Leopard)
			gw.LuckWin = 1
			gw.CardTypes = Leopard
			r.CardTypeList = append(r.CardTypeList, int32(Leopard))
			res.PotWinTypes.LuckDownPot = true
			res.BlackType = pb_msg.CardsType(Leopard)
		}
		if bg.IsStraightFlush() {
			r.Cards.LuckType = CardsType(Shunjin)
			gw.LuckWin = 1
			gw.CardTypes = Shunjin
			r.CardTypeList = append(r.CardTypeList, int32(Shunjin))
			res.PotWinTypes.LuckDownPot = true
			res.BlackType = pb_msg.CardsType(Shunjin)
		}
		if bg.IsFlush() {
			r.Cards.LuckType = CardsType(Golden)
			gw.LuckWin = 1
			gw.CardTypes = Golden
			r.CardTypeList = append(r.CardTypeList, int32(Golden))
			res.PotWinTypes.LuckDownPot = true
			res.BlackType = pb_msg.CardsType(Golden)
		}
		if bg.IsStraight() {
			r.Cards.LuckType = CardsType(Straight)
			gw.LuckWin = 1
			gw.CardTypes = Straight
			r.CardTypeList = append(r.CardTypeList, int32(Straight))
			res.PotWinTypes.LuckDownPot = true
			res.BlackType = pb_msg.CardsType(Straight)
		}
		if (bg.Key.Pair() >> 8) >= 9 {
			r.Cards.LuckType = CardsType(Pair)
			gw.LuckWin = 1
			gw.CardTypes = Pair
			r.CardTypeList = append(r.CardTypeList, int32(Pair))
			res.PotWinTypes.LuckDownPot = true
			res.BlackType = pb_msg.CardsType(Pair)
		} else if bg.IsPair() {
			gw.CardTypes = Pair
			r.CardTypeList = append(r.CardTypeList, int32(Pair))
			res.BlackType = pb_msg.CardsType(Pair)
		}
		if bg.IsZilch() {
			gw.CardTypes = Leaflet
			r.CardTypeList = append(r.CardTypeList, int32(Leaflet))
			res.BlackType = pb_msg.CardsType(Leaflet)
		}

		for _, v := range r.PlayerList {
			var taxMoney float64
			var totalWinMoney float64
			var totalLoseMoney float64

			if v != nil && v.IsAction == true {
				if v.IsRobot == false {
					totalWinMoney += float64(v.DownBetMoneys.BlackDownBet)
					taxMoney += float64(v.DownBetMoneys.BlackDownBet)

					totalLoseMoney += float64(v.DownBetMoneys.RedDownBet)
					totalLoseMoney += float64(v.DownBetMoneys.BlackDownBet)
					totalLoseMoney += float64(v.DownBetMoneys.LuckDownBet)
					if gw.LuckWin == 1 {
						if gw.CardTypes == Leopard {
							totalWinMoney += float64(v.DownBetMoneys.LuckDownBet)
							taxMoney += float64(v.DownBetMoneys.LuckDownBet * WinLeopard)
						}
						if gw.CardTypes == Shunjin {
							totalWinMoney += float64(v.DownBetMoneys.LuckDownBet)
							taxMoney += float64(v.DownBetMoneys.LuckDownBet * WinShunjin)
						}
						if gw.CardTypes == Golden {
							totalWinMoney += float64(v.DownBetMoneys.LuckDownBet)
							taxMoney += float64(v.DownBetMoneys.LuckDownBet * WinGolden)
						}
						if gw.CardTypes == Straight {
							totalWinMoney += float64(v.DownBetMoneys.LuckDownBet)
							taxMoney += float64(v.DownBetMoneys.LuckDownBet * WinStraight)
						}
						if gw.CardTypes == Pair {
							totalWinMoney += float64(v.DownBetMoneys.LuckDownBet)
							taxMoney += float64(v.DownBetMoneys.LuckDownBet * WinBigPair)
						}
					}
					tax := taxMoney * taxRate
					v.ResultMoney = totalWinMoney + taxMoney - tax
					v.Account += v.ResultMoney
					v.ResultMoney -= totalLoseMoney
					if v.ResultMoney > 0 {
						SurplusPool += tax
						SurplusPool -= v.ResultMoney //盈余池结算减去玩家Win金额
						v.WinTotalCount++

						v.WinResultMoney = v.ResultMoney
						log.Debug("玩家金额: %v, 进来了Win: %v", v.Account, v.WinResultMoney)

						timeStr := time.Now().Format("2006-01-02_15:04:05")
						nowTime := time.Now().Unix()
						reason := "ResultWinScore"

						//同时同步赢分和输分
						c4c.UserSyncWinScore(v, nowTime, timeStr, reason)
					} else if v.ResultMoney < 0 {
						v.LoseResultMoney = v.ResultMoney
						log.Debug("玩家金额: %v, 进来了Lose: %v", v.Account, v.LoseResultMoney)

						timeStr := time.Now().Format("2006-01-02_15:04:05")
						nowTime := time.Now().Unix()
						reason := "ResultLoseScore"

						//同时同步赢分和输分
						c4c.UserSyncLoseScore(v, nowTime, timeStr, reason)

						//将玩家输的金额添加到盈余池
						SurplusPool -= v.ResultMoney
					}
					log.Debug("<<===== 玩家下注: %v, 结算: %v =====>>", v.DownBetMoneys, v.ResultMoney)
				} else {
					totalWinMoney += float64(v.DownBetMoneys.BlackDownBet)
					taxMoney += float64(v.DownBetMoneys.BlackDownBet)

					totalLoseMoney += float64(v.DownBetMoneys.RedDownBet)
					totalLoseMoney += float64(v.DownBetMoneys.BlackDownBet)
					totalLoseMoney += float64(v.DownBetMoneys.LuckDownBet)
					if gw.LuckWin == 1 {
						if gw.CardTypes == Leopard {
							totalWinMoney += float64(v.DownBetMoneys.LuckDownBet)
							taxMoney += float64(v.DownBetMoneys.LuckDownBet * WinLeopard)
						}
						if gw.CardTypes == Shunjin {
							totalWinMoney += float64(v.DownBetMoneys.LuckDownBet)
							taxMoney += float64(v.DownBetMoneys.LuckDownBet * WinShunjin)
						}
						if gw.CardTypes == Golden {
							totalWinMoney += float64(v.DownBetMoneys.LuckDownBet)
							taxMoney += float64(v.DownBetMoneys.LuckDownBet * WinGolden)
						}
						if gw.CardTypes == Straight {
							totalWinMoney += float64(v.DownBetMoneys.LuckDownBet)
							taxMoney += float64(v.DownBetMoneys.LuckDownBet * WinStraight)
						}
						if gw.CardTypes == Pair {
							totalWinMoney += float64(v.DownBetMoneys.LuckDownBet)
							taxMoney += float64(v.DownBetMoneys.LuckDownBet * WinBigPair)
						}
					}
					tax := taxMoney * taxRate
					v.ResultMoney = totalWinMoney + taxMoney - tax
					v.Account += v.ResultMoney
					v.ResultMoney -= totalLoseMoney
					if v.ResultMoney > 0 {
						v.WinTotalCount++
					}
					log.Debug("<----- 机器人下注: %v, 结算: %v ----->", v.DownBetMoneys, v.ResultMoney)
				}
			}
		}
	}
	//广播开牌结果
	r.BroadCastMsg(res)

	//大厅用户添加列表数据
	hallData := &pb_msg.GameHallData_S2C{}
	for _, v := range mapUserIDPlayer {
		if v != nil && v.GameState == InGameHall {
			for _, data := range v.HallRoomData {
				hd := &pb_msg.HallData{}
				hd.RoomId = data.Rid
				if data.Rid == r.RoomId {
					// 判断该房间大厅数据列表是否已大于指定数据
					if len(data.HallCardTypeList) == 48 {
						log.Debug("<---------- 清空大厅列表数据~ ---------->")
						data.HallCardTypeList = nil
						data.HallRedBlackList = nil
					}
					data.HallCardTypeList = append(data.HallCardTypeList, hallCard)
					data.HallRedBlackList = append(data.HallRedBlackList, hallRBWin)
					hd.CardTypeList = data.HallCardTypeList
					hd.RedBlackList = data.HallRedBlackList
				} else {
					hd.CardTypeList = data.HallCardTypeList
					hd.RedBlackList = data.HallRedBlackList
				}
				hallData.HallData = append(hallData.HallData, hd)
			}
			v.SendMsg(hallData)
		}
	}

	//追加每局红黑Win、Luck、比牌类型的总集合
	r.RPotWinList = append(r.RPotWinList, gw)
	log.Debug("当前房间数据长度为: %v ~", len(r.RPotWinList))

	if len(r.RPotWinList) > 72 {
		r.RPotWinList = r.RPotWinList[1:]
	}
	if len(r.CardTypeList) > 72 {
		r.CardTypeList = r.CardTypeList[1:]
	}
	log.Debug("<-------- 更新盈余池金额为Last: %v --------->", SurplusPool)
}
