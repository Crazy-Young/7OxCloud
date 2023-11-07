package handler

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"

	"go.uber.org/zap"

	"github.com/go-redis/redis/v8"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/palp1tate/7OxCloud/7OxCloud-srv/interaction/consts"
	"github.com/palp1tate/7OxCloud/7OxCloud-srv/interaction/dao"
	"github.com/palp1tate/7OxCloud/7OxCloud-srv/interaction/global"
	"github.com/palp1tate/7OxCloud/7OxCloud-srv/interaction/mq"
	"github.com/palp1tate/7OxCloud/model"
	"github.com/palp1tate/7OxCloud/proto/interaction"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type InteractionServer struct {
	interactionProto.UnimplementedInteractionServiceServer
}

type MQMessage struct {
	UserId    int64 `json:"uid"`
	VideoId   int64 `json:"vid"`
	Timestamp int64 `json:"timestamp"`
}

func (s *InteractionServer) ViewVideo(ctx context.Context, req *interactionProto.ViewVideoRequest) (*empty.Empty, error) {
	message := &MQMessage{
		UserId:    req.UserId,
		VideoId:   req.VideoId,
		Timestamp: time.Now().Unix(),
	}
	body, _ := json.Marshal(&message)
	err := mq.SendMessage2MQ(body, consts.ViewVideoQueue)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "观看视频失败")
	}
	return &empty.Empty{}, nil
}

func (s *InteractionServer) LikeVideo(ctx context.Context, req *interactionProto.LikeVideoRequest) (*empty.Empty, error) {
	message := &MQMessage{
		UserId:    req.UserId,
		VideoId:   req.VideoId,
		Timestamp: time.Now().Unix(),
	}
	body, _ := json.Marshal(&message)
	err := mq.SendMessage2MQ(body, consts.LikeVideoQueue)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "点赞视频失败")
	}
	return &empty.Empty{}, nil
}

func (s *InteractionServer) CancelLikeVideo(ctx context.Context, req *interactionProto.CancelLikeVideoRequest) (*empty.Empty, error) {
	err := dao.CancelLikeVideo(req.UserId, req.VideoId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "取消点赞视频失败")
	}
	return &empty.Empty{}, nil
}

func (s *InteractionServer) CollectVideo(ctx context.Context, req *interactionProto.CollectVideoRequest) (*empty.Empty, error) {
	message := &MQMessage{
		UserId:    req.UserId,
		VideoId:   req.VideoId,
		Timestamp: time.Now().Unix(),
	}
	body, _ := json.Marshal(&message)
	err := mq.SendMessage2MQ(body, consts.CollectVideoQueue)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "收藏视频失败")
	}
	return &empty.Empty{}, nil
}

func (s *InteractionServer) CancelCollectVideo(ctx context.Context, req *interactionProto.CancelCollectVideoRequest) (*empty.Empty, error) {
	err := dao.CancelCollectVideo(req.UserId, req.VideoId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "取消收藏视频失败")
	}
	return &empty.Empty{}, nil
}

func (s *InteractionServer) GetLikeList(ctx context.Context, req *interactionProto.GetLikeListRequest) (*interactionProto.FeedResponse, error) {
	latestTime := time.Unix(req.LatestTime, 0)
	videoList, err := dao.GetUserLikedVideos(latestTime, req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "获取点赞列表失败")
	}
	if len(videoList) == 0 {
		return &interactionProto.FeedResponse{}, nil
	}
	videos := make([]*interactionProto.SmallVideo, len(videoList))
	for i, video := range videoList {
		videos[i] = SmallVideoModelToResponse(video)
	}
	res := &interactionProto.FeedResponse{
		Videos:   videos,
		NextTime: videoList[len(videoList)-1].CreatedAt.Unix(),
	}
	return res, nil
}

func (s *InteractionServer) GetCollectList(ctx context.Context, req *interactionProto.GetCollectListRequest) (*interactionProto.FeedResponse, error) {
	latestTime := time.Unix(req.LatestTime, 0)
	videoList, err := dao.GetUserCollectedVideos(latestTime, req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "获取收藏列表失败")
	}
	if len(videoList) == 0 {
		return &interactionProto.FeedResponse{}, nil
	}
	videos := make([]*interactionProto.SmallVideo, len(videoList))
	for i, video := range videoList {
		videos[i] = SmallVideoModelToResponse(video)
	}
	res := &interactionProto.FeedResponse{
		Videos:   videos,
		NextTime: videoList[len(videoList)-1].CreatedAt.Unix(),
	}
	return res, nil
}

func (s *InteractionServer) Comment(ctx context.Context, req *interactionProto.CommentRequest) (*empty.Empty, error) {
	comment := model.Comment{
		UserId:  req.UserId,
		VideoId: req.VideoId,
		Content: req.Content,
	}
	if req.ParentId != 0 {
		comment.ParentId = &req.ParentId
	}
	err := dao.CreateComment(&comment)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "评论失败")
	}
	return &empty.Empty{}, nil
}

func (s *InteractionServer) DeleteComment(ctx context.Context, req *interactionProto.DeleteCommentRequest) (*empty.Empty, error) {
	err := dao.DeleteComment(req.UserId, req.CommentId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "删除评论失败")
	}
	return &empty.Empty{}, nil
}

func (s *InteractionServer) GetCommentList(ctx context.Context, req *interactionProto.GetCommentListRequest) (*interactionProto.GetCommentListResponse, error) {
	comments, err := dao.GetCommentList(req.VideoId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "获取评论列表失败")
	}
	res := &interactionProto.GetCommentListResponse{
		Comments: comments,
	}
	return res, nil
}

func SmallVideoModelToResponse(video *model.Video) *interactionProto.SmallVideo {
	vid := video.ID
	tx := global.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	likeCount, _ := dao.GetVideoLikeCount(tx, vid)
	topicList, _ := dao.GetVideoTopics(tx, vid)
	if err := tx.Commit().Error; err != nil {
		return nil
	}
	topics := make([]*interactionProto.Topic, len(topicList))
	for i, topic := range topicList {
		topics[i] = TopicModelToResponse(topic)
	}
	return &interactionProto.SmallVideo{
		Id:          vid,
		CoverUrl:    video.CoverUrl,
		PlayUrl:     video.PlayUrl,
		LikeCount:   likeCount,
		Topics:      topics,
		Description: video.Description,
	}
}

func TopicModelToResponse(topic *model.Topic) *interactionProto.Topic {
	return &interactionProto.Topic{
		Id:          topic.ID,
		Description: topic.Description,
	}
}

func VideoLogToCSV(uid int64, vid int64, isLike int64, isCollect int64, timestamp int64) (err error) {
	InitCSVFile()
	fileName := consts.CSVFilePath
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0o666)
	if err != nil {
		zap.S().Warnf("打开文件失败: %s", err.Error())
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()
	err = writer.Write([]string{
		strconv.FormatInt(uid, 10),
		strconv.FormatInt(vid, 10),
		strconv.FormatInt(isLike, 10),
		strconv.FormatInt(isCollect, 10),
		strconv.FormatInt(timestamp, 10),
	})
	if err != nil {
		zap.S().Warnf("写入文件失败: %s", err.Error())
	}
	return
}

func RecordViewVideo(ctx context.Context, message *MQMessage) (err error) {
	err = global.RedisClient.ZAdd(ctx, fmt.Sprintf("history:%d", message.UserId),
		&redis.Z{Score: float64(message.Timestamp), Member: message.VideoId}).Err()
	if err != nil {
		zap.S().Warnf("记录观看历史失败: %s", err.Error())
	}
	err = VideoLogToCSV(message.UserId, message.VideoId, 0, 0, message.Timestamp)
	return
}

func RecordLikeVideo(ctx context.Context, message *MQMessage) (err error) {
	err = dao.LikeVideo(message.UserId, message.VideoId)
	if err != nil {
		zap.S().Warnf("点赞视频失败: %s", err.Error())
	}
	err = VideoLogToCSV(message.UserId, message.VideoId, 1, 0, message.Timestamp)
	return
}

func RecordCollectVideo(ctx context.Context, message *MQMessage) (err error) {
	err = dao.CollectVideo(message.UserId, message.VideoId)
	if err != nil {
		zap.S().Warnf("收藏视频失败: %s", err.Error())
	}
	err = VideoLogToCSV(message.UserId, message.VideoId, 0, 1, message.Timestamp)
	return
}
