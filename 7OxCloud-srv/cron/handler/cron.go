package handler

import (
	"context"
	"encoding/csv"
	"fmt"
	"os"
	"time"

	"github.com/palp1tate/7OxCloud/7OxCloud-srv/cron/global"
	"github.com/palp1tate/7OxCloud/7OxCloud-srv/interaction/consts"
	"github.com/palp1tate/7OxCloud/proto/transport"
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
)

func TransportCSV(ctx context.Context) {
	k := 1
	c := cron.New()
	_, err := c.AddFunc("@every 20m", func() {
		filePath := consts.CSVFilePath
		data, err := os.ReadFile(filePath)
		if err != nil {
			zap.S().Warnf("failed to read csv file: %s", err.Error())
			return
		}
		_, err = global.TransportServiceClient.SendCSVFile(ctx, &transportProto.SendRequest{Csv: data})
		if err != nil {
			zap.S().Warnf("failed to transport csv file: %s", err.Error())
			return
		}
		_ = os.Remove(consts.CSVFilePath)
		InitCSVFile()
		zap.S().Infof("第%d次上传csv文件成功", k)
		k++
	})
	if err != nil {
		zap.S().Panicf("failed to add cron job: %s", err.Error())
	}
	c.Start()
	defer c.Stop()
	select {}
}

func InitCSVFile() {
	fileName := consts.CSVFilePath
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		file, err := os.Create(fileName)
		if err != nil {
			zap.S().Panic("创建文件失败: " + err.Error())
		}
		defer file.Close()
		writer := csv.NewWriter(file)
		defer writer.Flush()
		err = writer.Write([]string{"uid", "vid", "isLike", "isCollect", "timestamp"})
		if err != nil {
			zap.S().Panic("写入文件失败: " + err.Error())
		}
	}
	return
}

func ClearCache() {
	e := global.ServerConfig.Redis.Expiration
	k := 1
	c := cron.New()
	_, err := c.AddFunc(fmt.Sprintf("0 0 */%d * *", e),
		func() {
			zap.S().Info("clear cache.......")
			ctx := context.Background()
			keys, err := global.RedisClient.Keys(ctx, "history:*").Result()
			if err != nil {
				zap.S().Warnf("failed to get keys: %s", err.Error())
			}
			pipe := global.RedisClient.Pipeline()
			minScore := "-inf"
			maxScore := fmt.Sprintf("%f", float64(time.Now().Add(-time.Duration(e)*time.Hour).Unix()))
			for _, key := range keys {
				_, err := pipe.ZRemRangeByScore(ctx, key, minScore, maxScore).Result()
				if err != nil {
					zap.S().Warnf("failed to remove keys: %s", err.Error())
				}
			}
			_, err = pipe.Exec(ctx)
			if err != nil {
				zap.S().Warnf("failed to exec: %s", err.Error())
			}
			zap.S().Infof("第%d次清理缓存成功", k)
			k++
		})
	if err != nil {
		zap.S().Panicf("failed to add cron job: %s", err.Error())
	}
	c.Start()
	defer c.Stop()
	select {}
}
