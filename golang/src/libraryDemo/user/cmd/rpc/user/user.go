// Code generated by goctl. DO NOT EDIT!
// Source: user.proto

//go:generate mockgen -destination ./user_mock.go -package user -source $GOFILE

package user

import (
	"context"

	"libraryDemo/user/cmd/rpc/userclient"

	"github.com/tal-tech/go-zero/zrpc"
)

type (
	IdReq         = userclient.IdReq
	UserInfoReply = userclient.UserInfoReply

	User interface {
		GetUser(ctx context.Context, in *IdReq) (*UserInfoReply, error)
	}

	defaultUser struct {
		cli zrpc.Client
	}
)

func NewUser(cli zrpc.Client) User {
	return &defaultUser{
		cli: cli,
	}
}

func (m *defaultUser) GetUser(ctx context.Context, in *IdReq) (*UserInfoReply, error) {
	client := userclient.NewUserClient(m.cli.Conn())
	return client.GetUser(ctx, in)
}
