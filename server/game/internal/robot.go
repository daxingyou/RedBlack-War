package internal

import (
	"fmt"
	"github.com/name5566/leaf/log"
	"math/rand"
	pb_msg "server/msg/Protocal"
	"time"
)

//机器人问题:
//1、机器人没钱怎么充值,不能再房间就直接充值,不然可以被其他用户看见
//2、机器人怎么下注，如果在桌面6个位置上，是否设置机器的下注速度和选择注池
//3、机器人选择注池的输赢,都要进行计算，只是不和盈余池牵扯，主要是前端做展示
//4、如果机器人金额如果小于50或不能参加游戏,则踢出房间删除机器人，在生成新的机器人加入该房间。

//机器人下标
var RobotIndex uint32

//Init 初始机器人控制中心
func (rc *RobotsCenter) Init() {
	log.Debug("-------------- RobotsCenter Init~! ---------------")
	rc.mapRobotList = make(map[uint32]*Player)
	RobotIndex = 0
}

//CreateRobot 创建一个机器人
func (rc *RobotsCenter) CreateRobot() *Player {
	r := &Player{}
	r.Init()

	r.IsRobot = true
	//生成随机ID
	r.Id = RandomID()
	//生成随机头像IMG
	r.HeadImg = RandomIMG()
	//生成机器人金币随机数
	rand.Intn(int(time.Now().Unix()))
	money := rand.Intn(6000) + 1000
	r.Account = float64(money)

	r.Index = RobotIndex
	fmt.Println("robot Index :", r.Index)
	rc.mapRobotList[r.Index] = r
	RobotIndex++
	log.Debug("创建机器人~ : %v", r.Id)
	return r
}

//RobotsDownBet 机器人进行下注
func (r *Room) RobotsDownBet() {
	// 线程下注
	go func() {
		var count int32
		t := time.NewTicker(time.Second)
		for range t.C {
			for _, v := range r.PlayerList {
				time.Sleep(time.Millisecond * 200)
				if v != nil && v.IsRobot == true && r.GameStat == DownBet {
					fmt.Println("你好 我是机器人-------------------------------", v.Id)
					time.Sleep(time.Millisecond * 200) //这里时间不能大于500,不然执行时间会超出一秒

					bet1 := RobotRandBet()
					pot1 := RobotRandPot()
					v.IsAction = true

					if v.Account < float64(bet1) {
						log.Debug("机器人:%v 下注金额小于身上筹码,下注失败~", v.Id)
						continue
					}

					//记录玩家在该房间总下注 和 房间注池的总金额
					if pb_msg.PotType(pot1) == pb_msg.PotType_RedPot {
						v.Account -= float64(bet1)
						v.DownBetMoneys.RedDownBet += bet1
						v.TotalAmountBet += bet1
						r.PotMoneyCount.RedMoneyCount += bet1
					}
					if pb_msg.PotType(pot1) == pb_msg.PotType_BlackPot {
						v.Account -= float64(bet1)
						v.DownBetMoneys.BlackDownBet += bet1
						v.TotalAmountBet += bet1
						r.PotMoneyCount.BlackMoneyCount += bet1

					}
					if pb_msg.PotType(pot1) == pb_msg.PotType_LuckPot {
						v.Account -= float64(bet1)
						v.DownBetMoneys.LuckDownBet += bet1
						v.TotalAmountBet += bet1
						r.PotMoneyCount.LuckMoneyCount += bet1
					}
					//返回前端玩家行动,更新玩家最新金额
					action := &pb_msg.PlayerAction_S2C{}
					action.Id = v.Id
					action.DownBet = bet1
					action.DownPot = pb_msg.PotType(pot1)
					action.IsAction = v.IsAction
					action.Account = v.Account
					r.BroadCastMsg(action)

					//广播玩家注池金额
					pot := &pb_msg.PotTotalMoney_S2C{}
					pot.PotMoneyCount = new(pb_msg.PotMoneyCount)
					pot.PotMoneyCount.RedMoneyCount = r.PotMoneyCount.RedMoneyCount
					pot.PotMoneyCount.BlackMoneyCount = r.PotMoneyCount.BlackMoneyCount
					pot.PotMoneyCount.LuckMoneyCount = r.PotMoneyCount.LuckMoneyCount
					r.BroadCastMsg(pot)

					//fmt.Println("玩家:", v.Id, "行动 红、黑、Luck下注: ", v.DownBetMoneys, "玩家总下注金额: ", v.TotalAmountBet)
				}
			}
		}
		log.Debug("Robot clock : %v", count)
		count++
		if count == RobotDownTime {
			return
		}
	}()
}

//RandNumber 随机机器下注金额
func RobotRandBet() int32 {
	slice := []int32{1, 10, 50, 100}
	rand.Seed(int64(time.Now().UnixNano()))
	num := rand.Intn(4)
	return slice[num]
}

//RandNumber 随机机器下注金额
func RobotRandPot() int32 {
	slice := []int32{1, 2, 3}
	rand.Seed(int64(time.Now().UnixNano()))
	num := rand.Intn(3)
	return slice[num]
}

//Start 机器人开工~！
func (rc *RobotsCenter) Start() {
	rand.Seed(int64(time.Now().UnixNano()))
	num := rand.Intn(5) + 4
	gameHall.LoadHallRobots(num)
}

//生成随机机器人ID
func RandomID() string {
	RobotId := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(100000000))
	return RobotId
}

//生成随机机器人头像IMG
func RandomIMG() string {
	slice := []string{
		"http://img1.imgtn.bdimg.com/it/u=152883880,155381385&fm=26&gp=0.jpg",
		"https://ss0.bdstatic.com/70cFvHSh_Q1YnxGkpoWK1HF6hhy/it/u=3080153838,1785924790&fm=26&gp=0.jpg",
		"https://ss1.bdstatic.com/70cFuXSh_Q1YnxGkpoWK1HF6hhy/it/u=1088543337,3570028698&fm=26&gp=0.jpg",
		"https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRBxUZlM3mpvYiIcpJ6buM8v4facsI_uaTRaEpLp4iss-CJWtlwzA",
	}
	rand.Seed(int64(time.Now().UnixNano()))
	num := rand.Intn(4)

	return slice[num]
}
