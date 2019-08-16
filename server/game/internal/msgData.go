package internal

import (
	"encoding/json"
	"fmt"
)

const (
	RECODE_PLAYERDESTORY    = 1001
	RECODE_MONEYNOTFULL     = 1002
	RECODE_PLAYERBREAKLINE  = 1003
	RECODE_PEOPLENOTFULL    = 1004
	RECODE_SELLTENOTDOWNBET = 1005
)

var recodeText = map[int32]string{
	RECODE_PLAYERDESTORY:    "用户已在其他地方登录",
	RECODE_MONEYNOTFULL:     "玩家金额不足,设为观战",
	RECODE_PLAYERBREAKLINE:  "玩家已掉线,断开连接",
	RECODE_PEOPLENOTFULL:    "房间人数不够,不能开始游戏",
	RECODE_SELLTENOTDOWNBET: "当前结算阶段,不能进行操作",
}

func jsonData() {
	reCode, err := json.Marshal(recodeText)
	if err != nil {
		fmt.Println("json.Marshal err:", err)
		return
	}

	data := string(reCode)
	fmt.Println("S2C jsonData String ~", data)
}
