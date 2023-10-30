package initialize

import (
	"github.com/palp1tate/7OxCloud/7OxCloud-srv/user/handler"
	"github.com/palp1tate/7OxCloud/7OxCloud-srv/user/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

func InitGRPC() *grpc.Server {
	server := grpc.NewServer()
	proto.RegisterUserServiceServer(server, &handler.UserServer{})
	grpc_health_v1.RegisterHealthServer(server, health.NewServer())
	return server
}
