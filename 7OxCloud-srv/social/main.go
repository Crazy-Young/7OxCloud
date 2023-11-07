package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/palp1tate/7OxCloud/7OxCloud-srv/social/global"
	"github.com/palp1tate/7OxCloud/7OxCloud-srv/social/initialize"
	"github.com/palp1tate/7OxCloud/util"
	"go.uber.org/zap"
)

func main() {
	initialize.InitConfig()
	initialize.InitLogger()
	initialize.InitMySQL()

	host := global.ServerConfig.Service.Host
	port := flag.Int("p", 0, "端口号")
	flag.Parse()
	if *port == 0 {
		*port, _ = util.GetFreePort()
	}
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		zap.S().Panic("failed to listen:" + err.Error())
	}

	server := initialize.InitGRPC()
	zap.S().Info("host: ", host)
	zap.S().Info("port: ", *port)
	go func() {
		err = server.Serve(lis)
		if err != nil {
			zap.S().Panic("failed to serve:" + err.Error())
		}
	}()

	client := initialize.NewRegistryClient(global.ServerConfig.Consul.Host, global.ServerConfig.Consul.Port)
	serviceName := global.ServerConfig.Service.Name
	serviceTags := global.ServerConfig.Service.Tags
	serviceId := util.GenerateUUID()
	err = client.Register(host, *port, serviceName, serviceTags, serviceId)
	if err != nil {
		zap.S().Panic("服务注册失败:", err.Error())
	}

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	server.GracefulStop()

	if err = client.DeRegister(serviceId); err != nil {
		zap.S().Warnf("%s注销失败", serviceName)
	} else {
		zap.S().Infof("%s注销成功", serviceName)
	}
}
