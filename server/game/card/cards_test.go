package card

import (
	"testing"
)

func Test_Dealer(t *testing.T) {
	//d := &Cards{}
	//
	//d.Shuffle()
	//
	//var ca Cards
	//ca = Cards{d.Take(), d.Take(), d.Take()}
	//fmt.Println("牌型类型1 :", ca)
	//fmt.Println("牌型类型2 :", ca.Hex())
	//fmt.Println("牌型类型3 :", ca.HexInt())
	//
	//fmt.Println("牌型剩余Take数量 ~ :", len(*d))

	//this := &RBdzDealer{}
	//
	//// 检查剩余牌数量
	//offset := this.Offset
	//if offset >= len(this.Poker)/2 {
	//	this.Poker = NewPoker(1, false, true)
	//	offset = 0
	//}
	//aaa := Hex(this.Poker)
	//fmt.Println("12:", aaa)
	//// 红黑各取3张牌
	//a := this.Poker[offset : offset+3]
	//b := this.Poker[offset+3 : offset+6]
	//
	//note := PokerArrayString(a) + "|" + PokerArrayString(b)
	//fmt.Println("note:::", note)
	//
	//hexa := HexInt(a)
	//str1 := fmt.Sprintf("%#v", a)
	//fmt.Println("offfff1 ::", str1)
	//fmt.Println("111:", hexa)
	//hexb := HexInt(b)
	//str2 := fmt.Sprintf("%#v", b)
	//fmt.Println("offfff2 ::", str2)
	//fmt.Println("222:", hexb)
	//
	RBdzPk()
}
