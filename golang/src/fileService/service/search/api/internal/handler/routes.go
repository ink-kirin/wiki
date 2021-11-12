// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"log"
	"net/http"

	"fileService/service/search/api/internal/svc"

	"github.com/tal-tech/go-zero/rest"
)

func RegisterHandlers(engine *rest.Server, serverCtx *svc.ServiceContext) {
	log.Printf("handler")
	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/search",
				Handler: SearchHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/search/batch",
				Handler: BatchSearchHandler(serverCtx),
			},
		},
	)
}