package main

import (
	"context"

	"github.com/palp1tate/7OxCloud/7OxCloud-srv/cron/global"
	"github.com/palp1tate/7OxCloud/7OxCloud-srv/cron/handler"
	"github.com/palp1tate/7OxCloud/7OxCloud-srv/cron/initialize"
	"go.uber.org/zap"
)

func main() {
	initialize.InitConfig()
	initialize.InitLogger()
	initialize.InitRedis()
	initialize.InitServiceConn()

	defer global.RedisClient.Close()

	zap.S().Info("cron service started")
	go handler.TransportCSV(context.Background())
	go handler.ClearCache()
	select {}
}
