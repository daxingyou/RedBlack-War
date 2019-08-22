package internal

import (
	"fmt"
	"github.com/name5566/leaf/log"
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

// 结算,欧赔方式计算,赔率放大100倍
const Radix = 100
const lostRadix = 0

// 0:红赢，1赔1，和 黑全输
// 1:黑赢，1赔1，和 红全输
const (
	WinLeopard       = 10*Radix + Radix //三同10倍
	WinStraightFlush = 5*Radix + Radix  //顺金5倍
	WinFlush         = 3*Radix + Radix  //金花3倍
	WinStraight      = 2*Radix + Radix  //顺子2倍
	WinBigPair       = 1*Radix + Radix  //大对子(9-A)
)

type RBdzDealer struct {
	Poker  []byte //所有的牌
	Offset int    //牌的位置
}

var (
	dealer = NewGoldenFlowerDealer(true)
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
func RBdzPk() {
	rb := &RBdzDealer{}
	a, b := rb.Deal()

	ha := Hex(a)
	log.Debug("花牌 数据Red~ : %v", ha)
	hb := Hex(b)
	log.Debug("花牌 数据Black~ : %v", hb)

	//红黑池牌型赋值
	r := &Room{}
	r.Cards = new(CardData)
	r.Cards.ReadCard = HexInt(a)
	r.Cards.BlackCard = HexInt(b)

	//字符串牌型
	note := PokerArrayString(a) + " | " + PokerArrayString(b)
	log.Debug("花牌 牌型~ : %v", note)

	// 可下注的选项数量(0:红赢,1:黑赢,2:幸运一击)
	ag := dealer.GetGroup(a)
	bg := dealer.GetGroup(b)

	gw := &GameWinList{}

	//获取Pot池Win
	if ag.Weight > bg.Weight { //redWin
		log.Debug("Red Win ~")
		gw.RedWin = 1

		if ag.IsThreeKind() {
			r.Cards.LuckType = CardsType(Leopard)
			gw.LuckWin = 1
			gw.CardTypes = Leopard
			r.CardTypeList = append(r.CardTypeList, int32(Leopard))
		}
		if ag.IsStraightFlush() {
			r.Cards.LuckType = CardsType(Shunjin)
			gw.LuckWin = 1
			gw.CardTypes = Shunjin
			r.CardTypeList = append(r.CardTypeList, int32(Shunjin))
		}
		if ag.IsFlush() {
			r.Cards.LuckType = CardsType(Golden)
			gw.LuckWin = 1
			gw.CardTypes = Golden
			r.CardTypeList = append(r.CardTypeList, int32(Golden))
		}
		if ag.IsStraight() {
			r.Cards.LuckType = CardsType(Straight)
			gw.LuckWin = 1
			gw.CardTypes = Straight
			r.CardTypeList = append(r.CardTypeList, int32(Straight))
		}
		if (ag.Key.Pair() >> 8) >= 9 {
			r.Cards.LuckType = CardsType(Pair)
			gw.LuckWin = 1
			gw.CardTypes = Pair
			r.CardTypeList = append(r.CardTypeList, int32(Pair))
		} else if ag.IsPair() {
			gw.CardTypes = Pair
			r.CardTypeList = append(r.CardTypeList, int32(Pair))
		}
		if ag.IsZilch() {
			gw.CardTypes = Leaflet
			r.CardTypeList = append(r.CardTypeList, int32(Leaflet))
		}
	} else if ag.Weight < bg.Weight { //blackWin
		log.Debug("Black Win ~")
		gw.BlackWin = 1

		if bg.IsThreeKind() {
			r.Cards.LuckType = CardsType(Leopard)
			gw.LuckWin = 1
			gw.CardTypes = Leopard
			r.CardTypeList = append(r.CardTypeList, int32(Leopard))
		}
		if bg.IsStraightFlush() {
			r.Cards.LuckType = CardsType(Shunjin)
			gw.LuckWin = 1
			gw.CardTypes = Shunjin
			r.CardTypeList = append(r.CardTypeList, int32(Shunjin))
		}
		if bg.IsFlush() {
			r.Cards.LuckType = CardsType(Golden)
			gw.LuckWin = 1
			gw.CardTypes = Golden
			r.CardTypeList = append(r.CardTypeList, int32(Golden))
		}
		if bg.IsStraight() {
			r.Cards.LuckType = CardsType(Straight)
			gw.LuckWin = 1
			gw.CardTypes = Straight
			r.CardTypeList = append(r.CardTypeList, int32(Straight))
		}
		if (bg.Key.Pair() >> 8) >= 9 {
			r.Cards.LuckType = CardsType(Pair)
			gw.LuckWin = 1
			gw.CardTypes = Pair
			r.CardTypeList = append(r.CardTypeList, int32(Pair))
		} else if bg.IsPair() {
			gw.CardTypes = Pair
			r.CardTypeList = append(r.CardTypeList, int32(Pair))
		}
		if bg.IsZilch() {
			gw.CardTypes = Leaflet
			r.CardTypeList = append(r.CardTypeList, int32(Leaflet))
		}
	}

	//追加每局红黑Win、Luck、比牌类型的总集合
	r.RPotWinList = append(r.RPotWinList, gw)

	//获取牌型处理
	if ag.IsThreeKind() {
		r.Cards.RedType = CardsType(Leopard)
		log.Debug("Red 三同10倍")
	}
	if bg.IsThreeKind() {
		r.Cards.BlackType = CardsType(Leopard)
		log.Debug("Black 三同10倍")
	}
	if ag.IsStraightFlush() {
		r.Cards.RedType = CardsType(Shunjin)
		log.Debug("Red 顺金5倍")
	}
	if bg.IsStraightFlush() {
		r.Cards.BlackType = CardsType(Shunjin)
		log.Debug("Black 顺金5倍")
	}
	if ag.IsFlush() {
		r.Cards.RedType = CardsType(Golden)
		log.Debug("Red 金花3倍")
	}
	if bg.IsFlush() {
		r.Cards.BlackType = CardsType(Golden)
		log.Debug("Black 金花3倍")
	}
	if ag.IsStraight() {
		r.Cards.RedType = CardsType(Straight)
		log.Debug("Red 顺子2倍")
	}
	if bg.IsStraight() {
		r.Cards.BlackType = CardsType(Straight)
		log.Debug("Black 顺子2倍")
	}
	if (ag.Key.Pair() >> 8) >= 9 {
		r.Cards.RedType = CardsType(Pair)
		log.Debug("Red 大对子(9-A)")
	} else if ag.IsPair() {
		r.Cards.RedType = CardsType(Pair)
		log.Debug("Red 小对子(2-8)")
	}
	if (bg.Key.Pair() >> 8) >= 9 {
		r.Cards.BlackType = CardsType(Pair)
		log.Debug("Black 大对子(9-A)")
	} else if bg.IsPair() {
		r.Cards.BlackType = CardsType(Pair)
		log.Debug("Black 小对子(2-8)")
	}
	if ag.IsZilch() {
		r.Cards.RedType = CardsType(Leaflet)
		log.Debug("Red 单张")
	}
	if bg.IsZilch() {
		r.Cards.BlackType = CardsType(Leaflet)
		log.Debug("Black 单张")
	}

	fmt.Println("Cards Data :", r.Cards)

}
