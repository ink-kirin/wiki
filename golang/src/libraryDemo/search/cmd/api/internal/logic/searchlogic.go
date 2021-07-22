package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"libraryDemo/user/cmd/rpc/user"

	"libraryDemo/search/cmd/api/internal/svc"
	"libraryDemo/search/cmd/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type SearchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) SearchLogic {
	return SearchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchLogic) Search(req types.SearchReq) (*types.SearchReply, error) {
	userIdNumber := json.Number(fmt.Sprintf("%v", l.ctx.Value("userId")))
	logx.Infof("userId: %s", userIdNumber)
	userId, err := userIdNumber.Int64()
	if err != nil {
		return nil, err
	}

	// 使用user rpc
	_, err = l.svcCtx.UserRpc.GetUser(l.ctx, &user.IdReq{
		Id: userId,
	})
	if err != nil {
		return nil, err
	}

	return &types.SearchReply{
		Name:  req.Name,
		Count: 100,
	}, nil
}
