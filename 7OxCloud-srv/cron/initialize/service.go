package initialize

import (
	"fmt"

	"github.com/palp1tate/7OxCloud/7OxCloud-srv/cron/global"
	"github.com/palp1tate/7OxCloud/proto/transport"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func InitServiceConn() {
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", global.ServerConfig.Recommend.Host,
		global.ServerConfig.Recommend.Port),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		zap.S().Panic("failed to connect recommend service:" + err.Error())
	}
	global.TransportServiceClient = transportProto.NewTransportServiceClient(conn)
}
