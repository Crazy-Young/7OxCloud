package initialize

import (
	"github.com/palp1tate/7OxCloud/7OxCloud-srv/cron/global"
	"go.uber.org/zap"
)

func InitLogger() {
	logger, _ := zap.NewProduction()
	if global.Debug {
		logger, _ = zap.NewDevelopment()
	}
	zap.ReplaceGlobals(logger)
}
