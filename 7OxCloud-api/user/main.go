package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/palp1tate/7OxCloud/7OxCloud-api/user/global"
	"github.com/palp1tate/7OxCloud/7OxCloud-api/user/initialize"
	"github.com/palp1tate/7OxCloud/7OxCloud-api/user/util"
	"go.uber.org/zap"
)

func main() {
	initialize.InitConfig()
	initialize.InitLogger()
	initialize.InitRedis()
	Router := initialize.Routers()
	if err := initialize.InitTrans("zh"); err != nil {
		zap.S().Warn("初始化翻译器失败:", err.Error())
		panic(err)
	}
	initialize.InitValidator()
	initialize.InitServiceConn()

	host := global.ServerConfig.Api.Host
	port := flag.Int("p", 0, "端口号")
	flag.Parse()
	if *port == 0 {
		*port, _ = util.GetFreePort()
	}
	zap.S().Info("host: ", host)
	zap.S().Info("port: ", *port)
	go func() {
		if err := Router.Run(fmt.Sprintf(":%d", *port)); err != nil {
			zap.S().Panic("启动失败:", err.Error())
		}
	}()

	client := initialize.NewRegistryClient(global.ServerConfig.Consul.Host, global.ServerConfig.Consul.Port)
	apiName := global.ServerConfig.Api.Name
	apiTags := global.ServerConfig.Api.Tags
	apiId := util.GenerateUUID()
	err := client.Register(host, *port, apiName, apiTags, apiId)
	if err != nil {
		zap.S().Panic("服务注册失败:", err.Error())
	}

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	if err := client.DeRegister(apiId); err != nil {
		zap.S().Warnf("%s注销失败", apiName)
	} else {
		zap.S().Infof("%s注销成功", apiName)
	}
}
