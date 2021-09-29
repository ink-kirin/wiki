package logic

import (
	"context"
	"fileService/service/search/api/internal/svc"
	"fileService/service/search/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

// 批量搜索业务逻辑

type BatchSearchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBatchSearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) BatchSearchLogic {
	return BatchSearchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BatchSearchLogic) BatchSearch(req types.BatchReq) ([]types.SearchRes, error) {
	var batch []types.SearchRes
	var res []interface{}
	res = append(res, []interface{}{"uid", "=", 11033})
	rep := make(map[string]interface{})
	rep["orderBy"] = "id asc"
	sql, param := l.svcCtx.AnnexModel.Collation(res, rep)
	annex, err := l.svcCtx.AnnexModel.FindAll(sql, param)
	if err != nil {
		return nil, err
	}
	for _, v := range annex {
		val := types.SearchRes{
			ID:      v.Id,
			UID:     v.Uid,
			BaseUrl: v.BaseUrl,
			Path:    v.Path,
		}
		batch = append(batch, val)
	}
	return batch, nil
}
