package initialize

import (
	"fmt"

	"github.com/palp1tate/7OxCloud/7OxCloud-srv/video/global"
	"github.com/palp1tate/7OxCloud/7OxCloud-srv/video/handler"
	"github.com/palp1tate/7OxCloud/proto/transport"
	"github.com/palp1tate/7OxCloud/proto/video"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

func InitGRPC() *grpc.Server {
	server := grpc.NewServer()
	videoProto.RegisterVideoServiceServer(server, &handler.VideoServer{})
	grpc_health_v1.RegisterHealthServer(server, health.NewServer())
	return server
}

func InitServiceConn() {
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", global.ServerConfig.Recommend.Host,
		global.ServerConfig.Recommend.Port),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic("failed to connect:" + err.Error())
	}
	global.TransportServiceClient = transportProto.NewTransportServiceClient(conn)
}
