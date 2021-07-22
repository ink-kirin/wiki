package logic

import (
	"context"

	"libraryDemo/user/cmd/rpc/internal/svc"
	"libraryDemo/user/cmd/rpc/userclient"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserLogic) GetUser(in *userclient.IdReq) (*userclient.UserInfoReply, error) {
	one, err := l.svcCtx.UserModel.FindOne(in.Id)
	if err != nil {
		return nil, err
	}

	return &userclient.UserInfoReply{
		Id:     one.Id,
		Name:   one.Name,
		Number: one.Number,
		Gender: one.Gender,
	}, nil
}
