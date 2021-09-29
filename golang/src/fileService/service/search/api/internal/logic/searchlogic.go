package logic

import (
	"context"
	"fileService/service/search/api/internal/svc"
	"fileService/service/search/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

// 搜索业务逻辑

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

func (l *SearchLogic) Search(req types.SearchReq) (*types.SearchRes, error) {
	annex, err := l.svcCtx.AnnexModel.FindOne(req.ID)
	if err != nil {
		return nil, err
	}
	return &types.SearchRes{
		ID:      annex.Id,
		UID:     annex.Uid,
		BaseUrl: annex.BaseUrl,
		Path:    annex.Path,
	}, nil
}
