package internal

import (
	"C"
	"github.com/name5566/leaf/log"
	"gopkg.in/mgo.v2"
	"server/conf"
	"time"
)

var (
	session *mgo.Session
)

const (
	dbName = "HONGHEIDAZHAN-Game"
	userDB = "room_data"
)

// 连接数据库集合的函数 传入集合 默认连接IM数据库
func initMongoDB() {
	// 此处连接正式线上数据库  下面是模拟的直接连接
	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    []string{conf.Server.MongoDBAddr},
		Timeout:  60 * time.Second,
		Database: conf.Server.MongoDBAuth,
		Username: conf.Server.MongoDBUser,
		Password: conf.Server.MongoDBPwd,
	}

	var err error
	session, err = mgo.DialWithInfo(mongoDBDialInfo)
	if err != nil {
		log.Fatal("Connect DataBase 数据库连接失败: %v ", err)
	}
	log.Debug("Connect DataBase 数据库连接成功~")

	//打开数据库
	session.SetMode(mgo.Monotonic, true)

}

func connect(dbName, cName string) (*mgo.Session, *mgo.Collection) {
	s := session.Copy()
	c := s.DB(dbName).C(cName)
	return s, c
}

// 插入房间数据
func InsertRoomData(r *Room) {
	s, c := connect(dbName, userDB)
	defer s.Close()

	err := c.Insert(r)
	if err != nil {
		log.Error("<----- 数据库插入房间数据失败 ~ ----->")
		return
	}
	log.Debug("<----- 数据库插入房间数据成功 ~ ----->")
}
