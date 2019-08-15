package internal

import (
	"github.com/name5566/leaf/log"
	pb_msg "server/msg/Protocal"
)

func (p *Player) Init() {
	p.ConnAgent = nil
	p.Index = 0

	p.HeadImg = "https://ss0.bdstatic.com/70cFvHSh_Q1YnxGkpoWK1HF6hhy/it/u=2873269578,797009742&fm=26&gp=0.jpg"
	p.Account = 4000

	p.CardTypeList = nil
	p.PotWinList = nil
	p.ContinueVot = 0
	p.IsGodGambling = false
	p.WinTotalCount = 0
	p.TotalAmountBet = 0

	p.room = nil
	p.IsOnline = true
}

// 用户缓存数据
var mapPlayerIndex int32
var mapGlobalPlayer map[int32]*Player
var mapUserIDPlayer map[string]*Player

// 初始化全局用户列表
func InitMapPlayer() {
	mapPlayerIndex = 0
	mapGlobalPlayer = make(map[int32]*Player)
	mapUserIDPlayer = make(map[string]*Player)
}

//CreatPlayer 创建用户信息
func CreatPlayer() *Player {
	p := &Player{}
	p.Init()
	mapGlobalPlayer[mapPlayerIndex] = p

	p.Index = mapPlayerIndex
	log.Debug("CreatePlayer index ~ : %v", p.Index)
	mapPlayerIndex++
	return p
}

//RegisterPlayer 注册用户信息
func RegisterPlayer(p *Player) {
	log.Debug("RegisterPlayer ~ : %v", p.Id)
	// 获取用户当前是否已经存在
	up, ok := mapUserIDPlayer[p.Id]

	// 如果有相同的ID，则断开和删除当前的用户链接，让新用户登录
	if ok {
		log.Debug("Have the same Player ID Login :%v", up.Id)

		errMsg := pb_msg.MsgInfo_S2C{}
		errMsg.Msg = recodeText[RECODE_PLAYERDESTORY]
		p.ConnAgent.WriteMsg(errMsg)
		log.Debug("用户已在其他地方登录~")

		up.ConnAgent.Destroy()
		up.ConnAgent.Close()
		DeletePlayer(up)
	}
	//将链接的Player数据赋值给map缓存
	mapUserIDPlayer[p.Id] = p
}

//DeletePlayer 删除用户信息
func DeletePlayer(p *Player) {
	// 删除mapGlobalPlayer用户索引
	delete(mapGlobalPlayer, p.Index)

	up, ok := mapUserIDPlayer[p.Id]
	if ok && up.Index == p.Index {
		// 删除mapUserIDPlayer用户索引
		delete(mapUserIDPlayer, p.Id)
		log.Debug("DeletePlayer SUCCESS ~ : %v", p.Id)
	} else {
		log.Debug("DeletePlayer come to nothing ~ : %v", p.Id)
	}
}
