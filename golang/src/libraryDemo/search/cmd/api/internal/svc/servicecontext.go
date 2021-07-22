package svc

import (
	"github.com/tal-tech/go-zero/rest"
	"github.com/tal-tech/go-zero/zrpc"
	"libraryDemo/search/cmd/api/internal/config"
	"libraryDemo/search/cmd/api/internal/middleware"
	"libraryDemo/user/cmd/rpc/user"
)

type ServiceContext struct {
	Config  config.Config
	Example rest.Middleware
	UserRpc user.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		Example: middleware.NewExampleMiddleware().Handle,
		UserRpc: user.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
