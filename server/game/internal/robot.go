package internal

import (
	"fmt"
	"github.com/name5566/leaf/log"
	"math/rand"
	"time"
)

//机器人问题:
//1、机器人没钱怎么充值,不能再房间就直接充值,不然可以被其他用户看见
//2、机器人怎么下注，如果在桌面6个位置上，是否设置机器的下注速度和选择注池
//3、机器人选择注池的输赢,都要进行计算，只是不和盈余池牵扯，主要是前端做展示
//4、如果机器人金额如果小于50或不能参加游戏,则踢出房间删除机器人，在生成新的机器人加入该房间。

//Init 初始机器人控制中心
func (rc *RobotsCenter) Init() {
	log.Debug("-------------- RobotsCenter Init~! ---------------")
	rc.mapRobotList = make(map[uint32]*Player)
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
	money := rand.Intn(3000) + 2000
	r.Account = float64(money)

	r.Index = uint32(len(rc.mapRobotList))
	rc.mapRobotList[r.Index] = r
	return r
}

//RobotsDownBet 机器人进行下注
func (r *Room) RobotsDownBet() {

	// 线程下注
	go func() {
		for _, v := range r.PlayerList {
			fmt.Println("你好 我是机器人---", v.Id)
			if v != nil && v.IsRobot == true && r.GameStat == DownBet {
				//bet := RobotRandBet()
				fmt.Println("你好 我是机器人-------------------------------", v.Id)
			}
		}
		//select {
		//case t := <-RobotDownBetChan:
		//	if t == true {
		//		return
		//	}
		//}
	}()
}

//RandNumber 随机机器下注金额
func RobotRandBet() int32 {
	slice := []int32{1, 10, 50, 100}
	rand.Seed(int64(time.Now().UnixNano()))
	num := rand.Intn(4)
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
		"https://ss2.bdstatic.com/70cFvnSh_Q1YnxGkpoWK1HF6hhy/it/u=1785199001,3375299815&fm=26&gp=0.jpg",
		"https://ss2.bdstatic.com/70cFvnSh_Q1YnxGkpoWK1HF6hhy/it/u=3169113226,3838220660&fm=26&gp=0.jpg",
		"https://ss3.bdstatic.com/70cFv8Sh_Q1YnxGkpoWK1HF6hhy/it/u=1751914404,475980807&fm=26&gp=0.jpg",
		"https://ss0.bdstatic.com/70cFvHSh_Q1YnxGkpoWK1HF6hhy/it/u=3080153838,1785924790&fm=26&gp=0.jpg",
	}
	rand.Seed(int64(time.Now().UnixNano()))
	num := rand.Intn(4)

	return slice[num]
}
