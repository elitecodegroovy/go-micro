package user

import (
	"github.com/micro/go-micro/v2/apps/json/plugins/db"
	proto "github.com/micro/go-micro/v2/apps/json/proto/user"
	"go.uber.org/zap"

	l "github.com/elitecodegroovy/goutil/logger"
)

var (
	log = l.GetLogger()
)

func (s *service) QueryUserByName(userName string) (ret *proto.User, err error) {
	queryString := `SELECT user_id, user_name, pwd FROM user WHERE user_name = ?`

	// 获取数据库
	o := db.GetDB()

	ret = &proto.User{}

	// 查询
	err = o.QueryRow(queryString, userName).Scan(&ret.Id, &ret.Name, &ret.Pwd)
	if err != nil {
		log.Fatal("[QueryUserByName] 查询数据失败，err：%s", zap.Any("error:", err))
		return
	}

	log.Info("req:"+userName, zap.Any("rsp: ", ret))

	return
}
