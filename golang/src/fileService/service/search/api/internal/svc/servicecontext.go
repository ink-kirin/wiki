package svc

import (
	"fileService/service/search/api/internal/config"
	"fileService/service/search/model"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config     config.Config
	AnnexModel model.XqdAlbumAnnex0Model // model interface
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:     c,
		AnnexModel: model.NewXqdAlbumAnnex0Model(conn, c.CacheRedis),
	}
}
