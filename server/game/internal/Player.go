package internal

import (
	"github.com/name5566/leaf/log"
	pb_msg "server/msg/Protocal"
	"time"
)

func (p *Player) Init() {
	p.ConnAgent = nil
	p.uClientDelay = 0
	p.Index = 0

	p.HeadImg = "https://ss0.bdstatic.com/70cFvHSh_Q1YnxGkpoWK1HF6hhy/it/u=2873269578,797009742&fm=26&gp=0.jpg"
	p.Account = 4000

	p.DownBetMoneys = nil
	p.DownPotTypes = nil
	p.TotalAmountBet = 0
	p.IsAction = false
	p.IsGodGambling = false
	p.ContinueVot = nil
	p.ResultWinMoney = 0
	p.ResultLoseMoney = 0

	p.room = nil

	p.WinTotalCount = 0
	p.PotWinList = nil
	p.CardTypeList = nil
	p.RedBlackList = nil
	p.RedWinCount = 0
	p.BlackWinCount = 0
	p.LuckWinCount = 0
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

		errMsg := &pb_msg.MsgInfo_S2C{}
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

//onClientBreathe 客户端呼吸，长时间未执行该函数可能已经断网，将主动踢掉
func (p *Player) onClientBreathe() {
	p.uClientDelay = 0
}

//StartBreathe 开始呼吸
func (p *Player) StartBreathe() {
	ticker := time.NewTicker(time.Second * 3)
	go func() {
		for { //循环
			<-ticker.C
			p.uClientDelay++
			//已经超过9秒没有收到客户端心跳，踢掉好了
			if p.uClientDelay > 4 {
				p.IsOnline = false

				errMsg := &pb_msg.MsgInfo_S2C{}
				errMsg.Msg = recodeText[RECODE_BREATHSTOP]
				p.ConnAgent.WriteMsg(errMsg)

				log.Debug("用户长时间未响应心跳,停止心跳~ : %v", p.Id)
				p.ConnAgent.Destroy()
				p.ConnAgent.Close()
				return
			}
		}
	}()
}
