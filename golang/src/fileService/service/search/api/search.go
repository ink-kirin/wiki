package main

import (
	"flag"
	"fmt"

	"fileService/service/search/api/internal/config"
	"fileService/service/search/api/internal/handler"
	"fileService/service/search/api/internal/svc"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/rest"
)

var configFile = flag.String("f", "etc/search-api.yaml", "the config file")

func main() {
	flag.Parse() // 解析命令行参数

	var c config.Config
	conf.MustLoad(*configFile, &c) // 加载配置

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf) // 加载服务
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
