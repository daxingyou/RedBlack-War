package internal

import "douniu/douniu_svr/log"

//RobotsCenter 机器人控制中心
type RobotsCenter struct {
	RobotsNumRoom int32            //每个房间放入机器数量
	mapRobotList  map[uint32]*Robot //机器人列表
}

//Robot 定义机器
type Robot struct {
	Player
	IsRun bool //机器人是否运行
}

//Init 初始机器人控制中心
func (rc *RobotsCenter) Init() {
	log.Debug("-------------- RobotsCenter Init~! ---------------")
	rc.mapRobotList = make(map[uint32]*Robot)
}

//CreateRobot 创建一个机器人
func (rc *RobotsCenter) CreateRobot() *Robot {
	robot := &Robot{}
	robot.Init()

	robot.IsRobot = true
	robot.Index = uint32(len(rc.mapRobotList))
	rc.mapRobotList[robot.Index] = robot
	return robot
}

//Running 机器人开始运行
func (ro *Robot) Running() {
	ro.IsRun = true

}

//Stop 机器人停止运行,一般在机器人下线调用
func (ro *Robot) Stop() {
	ro.IsRun = false
}
//Start 机器人开工~！
func (rc *RobotsCenter) Start(num int32) {
	gameHall.LoadHallRobots(num)
}

//Close 关闭机器人
func (rc *RobotsCenter) Close() {
	for _, r := range rc.mapRobotList {
		r.Stop()
		//r.ExitFromRoom()
	}

	rc.mapRobotList = make(map[uint32]*Robot)
}

var gRobotCenter RobotsCenter