package initialize

import (
	"fmt"
	"sync"

	_ "github.com/mbobakov/grpc-consul-resolver"
	"github.com/palp1tate/7OxCloud/7OxCloud-api/global"
	"github.com/palp1tate/7OxCloud/proto/user"
	"github.com/palp1tate/7OxCloud/proto/video"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func InitServiceConn() {
	consul := global.ServerConfig.Consul
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		userConn, err := grpc.Dial(
			fmt.Sprintf("consul://%s:%d/%s?wait=14s",
				consul.Host, consul.Port, global.ServerConfig.Service.User),
			grpc.WithTransportCredentials(insecure.NewCredentials()),
			grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
		)
		if err != nil {
			zap.S().Fatal("连接用户服务失败")
		}
		userServiceClient := userProto.NewUserServiceClient(userConn)
		global.UserServiceClient = userServiceClient
	}()

	go func() {
		defer wg.Done()
		videoConn, err := grpc.Dial(
			fmt.Sprintf("consul://%s:%d/%s?wait=14s",
				consul.Host, consul.Port, global.ServerConfig.Service.Video),
			grpc.WithTransportCredentials(insecure.NewCredentials()),
			grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
		)
		if err != nil {
			zap.S().Fatal("连接视频服务失败")
		}
		videoServiceClient := videoProto.NewVideoServiceClient(videoConn)
		global.VideoServiceClient = videoServiceClient
	}()

	wg.Wait()
}
