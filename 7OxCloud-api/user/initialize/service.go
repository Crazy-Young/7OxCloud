package initialize

import (
	"fmt"

	_ "github.com/mbobakov/grpc-consul-resolver"
	"github.com/palp1tate/7OxCloud/7OxCloud-api/user/global"
	"github.com/palp1tate/7OxCloud/7OxCloud-api/user/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func InitServiceConn() {
	consul := global.ServerConfig.Consul
	userConn, err := grpc.Dial(
		fmt.Sprintf("consul://%s:%d/%s?wait=14s", consul.Host, consul.Port, global.ServerConfig.Service.Name),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)
	if err != nil {
		zap.S().Fatal("连接用户服务失败")
	}
	userServiceClient := proto.NewUserServiceClient(userConn)
	global.UserServiceClient = userServiceClient
}
