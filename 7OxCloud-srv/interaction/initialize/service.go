package initialize

import (
	"github.com/palp1tate/7OxCloud/7OxCloud-srv/interaction/handler"
	"github.com/palp1tate/7OxCloud/proto/interaction"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

func InitGRPC() *grpc.Server {
	server := grpc.NewServer()
	interactionProto.RegisterInteractionServiceServer(server, &handler.InteractionServer{})
	grpc_health_v1.RegisterHealthServer(server, health.NewServer())
	return server
}
