package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"zero_demo/api_user/internal/config"
	"zero_demo/rpc_user/userclient"
)

type ServiceContext struct {
	Config config.Config
	User   userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		User:   userclient.NewUser(zrpc.MustNewClient(c.UserRpcConf)),
	}
}
