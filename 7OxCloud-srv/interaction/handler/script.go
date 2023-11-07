package handler

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"os"

	"github.com/palp1tate/7OxCloud/7OxCloud-srv/interaction/consts"
	"github.com/palp1tate/7OxCloud/7OxCloud-srv/interaction/mq"
	"go.uber.org/zap"
)

type SyncVideo struct{}

func SyncViewVideo(ctx context.Context) {
	Sync := new(SyncVideo)
	err := Sync.ViewVideoSync(ctx, consts.ViewVideoQueue)
	if err != nil {
		zap.S().Warnf("failed to sync view video: %s", err.Error())
		return
	}
}

func SyncLikeVideo(ctx context.Context) {
	Sync := new(SyncVideo)
	err := Sync.LikeVideoSync(ctx, consts.LikeVideoQueue)
	if err != nil {
		zap.S().Warnf("failed to sync like video: %s", err.Error())
		return
	}
}

func SyncCollectVideo(ctx context.Context) {
	Sync := new(SyncVideo)
	err := Sync.CollectVideoSync(ctx, consts.CollectVideoQueue)
	if err != nil {
		zap.S().Warnf("failed to sync collect video: %s", err.Error())
		return
	}
}

func (s *SyncVideo) ViewVideoSync(ctx context.Context, queueName string) (err error) {
	msg, err := mq.ConsumeMessage(ctx, queueName)
	if err != nil {
		zap.S().Warnf("failed to consume message: %s", err.Error())
		return err
	}
	var forever chan struct{}
	go func() {
		for d := range msg {
			var message *MQMessage
			err = json.Unmarshal(d.Body, &message)
			if err != nil {
				zap.S().Warnf("failed to unmarshal message: %s", err.Error())
				return
			}
			err = RecordViewVideo(ctx, message)
			if err != nil {
				zap.S().Warnf("failed to record view video: %s", err.Error())
				return
			}
			d.Ack(false)
		}
	}()
	<-forever
	return nil
}

func (s *SyncVideo) LikeVideoSync(ctx context.Context, queueName string) (err error) {
	msg, err := mq.ConsumeMessage(ctx, queueName)
	if err != nil {
		zap.S().Warnf("failed to consume message: %s", err.Error())
		return err
	}
	var forever chan struct{}
	go func() {
		for d := range msg {
			var message *MQMessage
			err = json.Unmarshal(d.Body, &message)
			if err != nil {
				zap.S().Warnf("failed to unmarshal message: %s", err.Error())
				return
			}
			err = RecordLikeVideo(ctx, message)
			if err != nil {
				zap.S().Warnf("failed to record like video: %s", err.Error())
				return
			}
			d.Ack(false)
		}
	}()
	<-forever
	return nil
}

func (s *SyncVideo) CollectVideoSync(ctx context.Context, queueName string) (err error) {
	msg, err := mq.ConsumeMessage(ctx, queueName)
	if err != nil {
		zap.S().Warnf("failed to consume message: %s", err.Error())
		return err
	}
	var forever chan struct{}
	go func() {
		for d := range msg {
			var message *MQMessage
			err = json.Unmarshal(d.Body, &message)
			if err != nil {
				zap.S().Warnf("failed to unmarshal message: %s", err.Error())
				return
			}
			err = RecordCollectVideo(ctx, message)
			if err != nil {
				zap.S().Warnf("failed to record collect video: %s", err.Error())
				return
			}
			d.Ack(false)
		}
	}()
	<-forever
	return nil
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
