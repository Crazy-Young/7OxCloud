package handler

import (
	"context"
	"time"

	"github.com/bwmarrin/snowflake"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/palp1tate/7OxCloud/7OxCloud-srv/video/dao"
	"github.com/palp1tate/7OxCloud/7OxCloud-srv/video/global"
	"github.com/palp1tate/7OxCloud/model"
	"github.com/palp1tate/7OxCloud/proto/video"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type VideoServer struct {
	videoProto.UnimplementedVideoServiceServer
}

func (s *VideoServer) Feed(ctx context.Context, req *videoProto.FeedRequest) (*videoProto.MiniFeedResponse, error) {
	latestTime := time.Unix(req.LatestTime, 0)
	videoList, err := dao.GetVideos(latestTime)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "获取视频流失败")
	}
	if len(videoList) == 0 {
		return &videoProto.MiniFeedResponse{}, nil
	}
	videos := make([]*videoProto.MiniVideo, len(videoList))
	for i, video := range videoList {
		videos[i] = MiniVideoModelToResponse(video, req.CurrentUserId)
	}
	res := &videoProto.MiniFeedResponse{
		Videos:   videos,
		NextTime: videoList[len(videoList)-1].CreatedAt.Unix(),
	}
	return res, nil
}

func (s *VideoServer) Publish(ctx context.Context, req *videoProto.PublishRequest) (*empty.Empty, error) {
	node, err := snowflake.NewNode(1)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "视频ID生成失败")
	}
	vid := node.Generate().Int64()
	video := model.Video{
		ID:          vid,
		Description: req.Description,
		PlayUrl:     req.PlayUrl,
		CoverUrl:    req.CoverUrl,
		AuthorId:    req.UserId,
	}
	err = dao.CreateVideo(&video, req.Topics, req.CategoryId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "发布视频失败")
	}
	return &empty.Empty{}, nil
}

func (s *VideoServer) PublishList(ctx context.Context, req *videoProto.PublishListRequest) (*videoProto.PublishListResponse, error) {
	latestTime := time.Unix(req.LatestTime, 0)
	videoList, err := dao.GetVideosByUserId(latestTime, req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "获取视频列表失败")
	}
	if len(videoList) == 0 {
		return &videoProto.PublishListResponse{}, nil
	}
	videos := make([]*videoProto.SmallVideo, len(videoList))
	for i, video := range videoList {
		videos[i] = SmallVideoModelToResponse(video)
	}
	res := &videoProto.PublishListResponse{
		Videos:   videos,
		NextTime: videoList[len(videoList)-1].CreatedAt.Unix(),
	}
	return res, nil
}

func (s *VideoServer) Video(ctx context.Context, req *videoProto.VideoRequest) (*videoProto.VideoResponse, error) {
	video, err := dao.GetVideo(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "视频不存在")
	}
	res := &videoProto.VideoResponse{
		Video: VideoModelToResponse(video, req.CurrentUserId),
	}
	return res, nil
}

func (s *VideoServer) FeedByTopic(ctx context.Context, req *videoProto.FeedByTopicRequest) (*videoProto.FeedResponse, error) {
	latestTime := time.Unix(req.LatestTime, 0)
	videoList, err := dao.GetVideosByTopicId(latestTime, req.TopicId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "获取视频流失败")
	}
	if len(videoList) == 0 {
		return &videoProto.FeedResponse{}, nil
	}
	videos := make([]*videoProto.Video, len(videoList))
	for i, video := range videoList {
		videos[i] = VideoModelToResponse(video, req.CurrentUserId)
	}
	res := &videoProto.FeedResponse{
		Videos:   videos,
		NextTime: videoList[len(videoList)-1].CreatedAt.Unix(),
	}
	return res, nil
}

func (s *VideoServer) CategoryList(ctx context.Context, req *empty.Empty) (*videoProto.CategoryListResponse, error) {
	categoryList, err := dao.GetCategories()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "获取分类列表失败")
	}
	categories := make([]*videoProto.Category, len(categoryList))
	for i, category := range categoryList {
		categories[i] = CategoryModelToResponse(category)
	}
	res := &videoProto.CategoryListResponse{
		Categories: categories,
	}
	return res, nil
}

func (s *VideoServer) TopicList(ctx context.Context, req *empty.Empty) (*videoProto.TopicListResponse, error) {
	topicList, err := dao.GetTopics()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "获取话题列表失败")
	}
	topics := make([]*videoProto.Topic, len(topicList))
	for i, topic := range topicList {
		topics[i] = TopicModelToResponse(topic)
	}
	res := &videoProto.TopicListResponse{
		Topics: topics,
	}
	return res, nil
}

func (s *VideoServer) FeedByCategory(ctx context.Context, req *videoProto.FeedByCategoryRequest) (*videoProto.MiniFeedResponse, error) {
	latestTime := time.Unix(req.LatestTime, 0)
	videoList, err := dao.GetVideosByCategoryId(latestTime, req.CategoryId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "获取视频流失败")
	}
	if len(videoList) == 0 {
		return &videoProto.MiniFeedResponse{}, nil
	}
	videos := make([]*videoProto.MiniVideo, len(videoList))
	for i, video := range videoList {
		videos[i] = MiniVideoModelToResponse(video, req.CurrentUserId)
	}
	res := &videoProto.MiniFeedResponse{
		Videos:   videos,
		NextTime: videoList[len(videoList)-1].CreatedAt.Unix(),
	}
	return res, nil
}

func (s *VideoServer) FeedBySearch(ctx context.Context, req *videoProto.FeedBySearchRequest) (*videoProto.FeedResponse, error) {
	latestTime := time.Unix(req.LatestTime, 0)
	videoList, err := dao.GetVideosBySearch(latestTime, req.Keyword)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "获取视频流失败")
	}
	if len(videoList) == 0 {
		return &videoProto.FeedResponse{}, nil
	}
	videos := make([]*videoProto.Video, len(videoList))
	for i, video := range videoList {
		videos[i] = VideoModelToResponse(video, req.CurrentUserId)
	}
	res := &videoProto.FeedResponse{
		Videos:   videos,
		NextTime: videoList[len(videoList)-1].CreatedAt.Unix(),
	}
	return res, nil
}

func (s *VideoServer) HotFeed(ctx context.Context, req *videoProto.FeedRequest) (*videoProto.MiniFeedResponse, error) {
	latestTime := time.Unix(req.LatestTime, 0)
	videoList, err := dao.GetHotVideos(latestTime)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "获取热点视频流失败")
	}
	if len(videoList) == 0 {
		return &videoProto.MiniFeedResponse{}, nil
	}
	videos := make([]*videoProto.MiniVideo, len(videoList))
	for i, video := range videoList {
		videos[i] = MiniVideoModelToResponse(video, req.CurrentUserId)
	}
	res := &videoProto.MiniFeedResponse{
		Videos:   videos,
		NextTime: videoList[len(videoList)-1].CreatedAt.Unix(),
	}
	return res, nil
}

func (s *VideoServer) DeleteVideo(ctx context.Context, req *videoProto.VideoRequest) (*empty.Empty, error) {
	err := dao.DeleteVideo(req.Id, req.CurrentUserId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "删除视频失败")
	}
	return &empty.Empty{}, nil
}

func VideoModelToResponse(video *model.Video, currentUserId int64) *videoProto.Video {
	vid := video.ID
	tx := global.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	likeCount, _ := dao.GetVideoLikeCount(tx, vid)
	commentCount, _ := dao.GetVideoCommentCount(tx, vid)
	collectCount, _ := dao.GetVideoCollectCount(tx, vid)
	isLike, _ := dao.GetIsLike(tx, vid, currentUserId)
	isCollect, _ := dao.GetIsCollect(tx, vid, currentUserId)
	author, _ := dao.FindAuthor(tx, vid)
	topicList, _ := dao.GetVideoTopics(tx, vid)
	topics := make([]*videoProto.Topic, len(topicList))
	for i, topic := range topicList {
		topics[i] = TopicModelToResponse(topic)
	}
	user := UserModelToResponse(tx, author, currentUserId)
	if err := tx.Commit().Error; err != nil {
		return nil
	}
	return &videoProto.Video{
		Id:           vid,
		Description:  video.Description,
		PlayUrl:      video.PlayUrl,
		CoverUrl:     video.CoverUrl,
		LikeCount:    likeCount,
		CommentCount: commentCount,
		CollectCount: collectCount,
		IsLike:       isLike,
		IsCollect:    isCollect,
		Author:       user,
		Topics:       topics,
		PublishTime:  video.CreatedAt.Unix(),
	}
}

func UserModelToResponse(tx *gorm.DB, user model.User, currentUserId int64) *videoProto.User {
	isFollow, _ := dao.GetIsFollow(tx, user.ID, currentUserId)
	return &videoProto.User{
		Id:       user.ID,
		Username: user.Username,
		Avatar:   user.Avatar,
		Location: user.Location,
		IsFollow: isFollow,
	}
}

func TopicModelToResponse(topic *model.Topic) *videoProto.Topic {
	return &videoProto.Topic{
		Id:          topic.ID,
		Description: topic.Description,
	}
}

func MiniVideoModelToResponse(video *model.Video, currentUserId int64) *videoProto.MiniVideo {
	vid := video.ID
	tx := global.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	likeCount, _ := dao.GetVideoLikeCount(tx, vid)
	topicList, _ := dao.GetVideoTopics(tx, vid)
	author, _ := dao.FindAuthor(tx, vid)
	miniUser := MiniUserModelToResponse(tx, author, currentUserId)
	if err := tx.Commit().Error; err != nil {
		return nil
	}
	topics := make([]*videoProto.Topic, len(topicList))
	for i, topic := range topicList {
		topics[i] = TopicModelToResponse(topic)
	}
	return &videoProto.MiniVideo{
		Id:          vid,
		CoverUrl:    video.CoverUrl,
		PlayUrl:     video.PlayUrl,
		LikeCount:   likeCount,
		Topics:      topics,
		Description: video.Description,
		Author:      miniUser,
		PublishTime: video.CreatedAt.Unix(),
	}
}

func SmallVideoModelToResponse(video *model.Video) *videoProto.SmallVideo {
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
	topics := make([]*videoProto.Topic, len(topicList))
	for i, topic := range topicList {
		topics[i] = TopicModelToResponse(topic)
	}
	return &videoProto.SmallVideo{
		Id:          vid,
		Description: video.Description,
		PlayUrl:     video.PlayUrl,
		CoverUrl:    video.CoverUrl,
		LikeCount:   likeCount,
		Topics:      topics,
	}
}

func MiniUserModelToResponse(tx *gorm.DB, user model.User, currentUserId int64) *videoProto.MiniUser {
	isFollow, _ := dao.GetIsFollow(tx, user.ID, currentUserId)
	return &videoProto.MiniUser{
		Id:       user.ID,
		Username: user.Username,
		IsFollow: isFollow,
	}
}

func CategoryModelToResponse(category *model.Category) *videoProto.Category {
	return &videoProto.Category{
		Id:   category.ID,
		Name: category.Name,
	}
}
