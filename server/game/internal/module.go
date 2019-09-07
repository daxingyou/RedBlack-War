package internal

import (
	"github.com/name5566/leaf/module"
	"server/base"
)

var (
	skeleton = base.NewSkeleton()
	ChanRPC  = skeleton.ChanRPCServer

	c4c = &Conn4Center{}
)

type Module struct {
	*module.Skeleton
}

func (m *Module) OnInit() {
	m.Skeleton = skeleton

	gameHall.Init()
	InitMapPlayer()

	gRobotCenter.Init()
	gRobotCenter.Start()

	//中心服初始化,主动请求Token
	c4c.Init()
	c4c.ReqCenterToken()
}

func (m *Module) OnDestroy() {
	c4c.onDestroy()
}
