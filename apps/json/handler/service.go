package handler

import (
	"context"
	l "github.com/elitecodegroovy/goutil/logger"

	us "github.com/micro/go-micro/v2/apps/json/model/user"
	s "github.com/micro/go-micro/v2/apps/json/proto/user"
)

var (
	log = l.GetLogger()
)

type Service struct{}

var (
	userService us.Service
)

// Init 初始化handler
func Init() {

	var err error
	userService, err = us.GetService()
	if err != nil {
		log.Fatal("[Init] 初始化Handler错误")
		return
	}
}

// QueryUserByName 通过参数中的名字返回用户
func (e *Service) QueryUserByName(ctx context.Context, req *s.Request, rsp *s.Response) error {
	user, err := userService.QueryUserByName(req.UserName)
	if err != nil {
		rsp.Error = &s.Error{
			Code:   500,
			Detail: err.Error(),
		}

		return nil
	}

	rsp.User = user
	return nil
}
