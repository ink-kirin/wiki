package handler

import (
	"fileService/common/response"
	"net/http"

	"fileService/service/search/api/internal/logic"
	"fileService/service/search/api/internal/svc"
	"fileService/service/search/api/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func BatchSearchHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.BatchReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewBatchSearchLogic(r.Context(), ctx)
		resp, err := l.BatchSearch(req)
		response.Response(w, resp, err)
	}
}
