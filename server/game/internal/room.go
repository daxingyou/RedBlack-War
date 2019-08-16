package internal

import (
	"fmt"
	"math/rand"
	"time"
)

func (r *Room) RoomInit() {

	r.RoomId = r.GetRoomNumber()
	r.PlayerList = nil

	r.GodGambleName = ""
	r.RoomStat = RoomStatusNone

	r.PotMoneyCount = nil
	r.CardTypeList = nil
	r.RPotWinList = nil
	r.GameTotalCount = 0
}

func (r *Room) GetRoomNumber() string {
	roomNumber := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
	return roomNumber
}


