package internal

import "github.com/name5566/leaf/log"

const (
	taxRate    float64 = 0.06 //税率   //todo 正式是 0.06
	SurplusTax float64 = 0.2  //指定盈余池的百分随机数
)

//盈余池
var SurplusPool float64 = 0

//记录进入大厅玩家的数量,为了统计 盈余池 * 6
var AllPlayerCount []string

//返回记录的玩家总数量
func RecordPlayerCount() int32 {
	log.Debug("游戏玩过总人数数量: %v", int32(len(AllPlayerCount)))
	return int32(len(AllPlayerCount))
}
