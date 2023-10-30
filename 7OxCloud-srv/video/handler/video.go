package handler

import (
	"context"

	"github.com/palp1tate/7OxCloud/7OxCloud-srv/video/global"

	"github.com/bwmarrin/snowflake"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/palp1tate/7OxCloud/7OxCloud-srv/video/dao"
	"github.com/palp1tate/7OxCloud/7OxCloud-srv/video/model"
	"github.com/palp1tate/7OxCloud/7OxCloud-srv/video/proto"
	"github.com/palp1tate/7OxCloud/7OxCloud-srv/video/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type VideoServer struct {
	proto.UnimplementedVideoServiceServer
}

func (s *VideoServer) Feed(ctx context.Context, req *proto.FeedRequest) (*proto.FeedResponse, error) {
	latestTime := util.ConvertTimestamp2Time(req.LatestTime)
	videoList, err := dao.GetVideos(latestTime)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "获取视频流失败")
	}
	videos := make([]*proto.Video, len(videoList))
	for i, video := range videoList {
		videos[i] = VideoModelToResponse(video, req.CurrentUserId)
	}
	res := &proto.FeedResponse{
		Videos:   videos,
		NextTime: util.ConvertTime2Timestamp(videoList[len(videoList)-1].CreatedAt),
	}
	return res, nil
}

func (s *VideoServer) Publish(ctx context.Context, req *proto.PublishRequest) (*empty.Empty, error) {
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

func (s *VideoServer) PublishList(ctx context.Context, req *proto.PublishListRequest) (*proto.PublishListResponse, error) {
	videoList, err := dao.GetVideosByUserId(req.Uid)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "获取视频列表失败")
	}
	videos := make([]*proto.MiniVideo, len(videoList))
	for i, video := range videoList {
		videos[i] = MiniVideoModelToResponse(video)
	}
	res := &proto.PublishListResponse{
		Videos: videos,
	}
	return res, nil
}

func (s *VideoServer) Video(ctx context.Context, req *proto.VideoRequest) (*proto.VideoResponse, error) {
	video, err := dao.GetVideo(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "视频不存在")
	}
	res := &proto.VideoResponse{
		Video: VideoModelToResponse(video, req.CurrentUserId),
	}
	return res, nil
}

func (s *VideoServer) FeedByTopic(ctx context.Context, req *proto.FeedByTopicRequest) (*proto.FeedResponse, error) {
	latestTime := util.ConvertTimestamp2Time(req.LatestTime)
	videoList, err := dao.GetVideosByTopicId(latestTime, req.TopicId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "获取视频流失败")
	}
	videos := make([]*proto.Video, len(videoList))
	for i, video := range videoList {
		videos[i] = VideoModelToResponse(video, req.CurrentUserId)
	}
	res := &proto.FeedResponse{
		Videos:   videos,
		NextTime: util.ConvertTime2Timestamp(videoList[len(videoList)-1].CreatedAt),
	}
	return res, nil
}

func (s *VideoServer) CategoryList(ctx context.Context, req *empty.Empty) (*proto.CategoryListResponse, error) {
	categoryList, err := dao.GetCategories()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "获取分类列表失败")
	}
	categories := make([]*proto.Category, len(categoryList))
	for i, category := range categoryList {
		categories[i] = CategoryModelToResponse(category)
	}
	res := &proto.CategoryListResponse{
		Categories: categories,
	}
	return res, nil
}

func (s *VideoServer) TopicList(ctx context.Context, req *empty.Empty) (*proto.TopicListResponse, error) {
	topicList, err := dao.GetTopics()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "获取话题列表失败")
	}
	topics := make([]*proto.Topic, len(topicList))
	for i, topic := range topicList {
		topics[i] = TopicModelToResponse(topic)
	}
	res := &proto.TopicListResponse{
		Topics: topics,
	}
	return res, nil
}

func (s *VideoServer) FeedByCategory(ctx context.Context, req *proto.FeedByCategoryRequest) (*proto.FeedResponse, error) {
	latestTime := util.ConvertTimestamp2Time(req.LatestTime)
	videoList, err := dao.GetVideosByCategoryId(latestTime, req.CategoryId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "获取视频流失败")
	}
	videos := make([]*proto.Video, len(videoList))
	for i, video := range videoList {
		videos[i] = VideoModelToResponse(video, req.CurrentUserId)
	}
	res := &proto.FeedResponse{
		Videos:   videos,
		NextTime: util.ConvertTime2Timestamp(videoList[len(videoList)-1].CreatedAt),
	}
	return res, nil
}

func (s *VideoServer) FeedBySearch(ctx context.Context, req *proto.FeedBySearchRequest) (*proto.FeedResponse, error) {
	latestTime := util.ConvertTimestamp2Time(req.LatestTime)
	videoList, err := dao.GetVideosBySearch(latestTime, req.Keyword)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "获取视频流失败")
	}
	videos := make([]*proto.Video, len(videoList))
	for i, video := range videoList {
		videos[i] = VideoModelToResponse(video, req.CurrentUserId)
	}
	res := &proto.FeedResponse{
		Videos:   videos,
		NextTime: util.ConvertTime2Timestamp(videoList[len(videoList)-1].CreatedAt),
	}
	return res, nil
}

func VideoModelToResponse(video *model.Video, currentUserId int64) *proto.Video {
	vid := video.ID
	likeCount, _ := dao.GetVideoLikeCount(vid)
	tx := global.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	commentCount, _ := dao.GetVideoCommentCount(tx, vid)
	collectCount, _ := dao.GetVideoCollectCount(tx, vid)
	isLike, _ := dao.GetIsLike(tx, vid, currentUserId)
	author, _ := dao.FindAuthor(tx, vid)
	topicList, _ := dao.GetVideoTopics(tx, vid)
	if err := tx.Commit().Error; err != nil {
		return nil
	}
	topics := make([]*proto.Topic, len(topicList))
	for i, topic := range topicList {
		topics[i] = TopicModelToResponse(topic)
	}
	return &proto.Video{
		Id:           vid,
		Description:  video.Description,
		PlayUrl:      video.PlayUrl,
		CoverUrl:     video.CoverUrl,
		LikeCount:    likeCount,
		CommentCount: commentCount,
		CollectCount: collectCount,
		IsLike:       isLike,
		Author:       UserModelToResponse(author, currentUserId),
		Topics:       topics,
		PublishTime:  util.ConvertTime2Timestamp(video.CreatedAt),
	}
}

func UserModelToResponse(user model.User, currentUserId int64) *proto.User {
	isFollow, _ := dao.GetIsFollow(user.ID, currentUserId)
	return &proto.User{
		Id:       user.ID,
		Username: user.Username,
		Avatar:   user.Avatar,
		Location: user.Location,
		IsFollow: isFollow,
	}
}

func TopicModelToResponse(topic *model.Topic) *proto.Topic {
	return &proto.Topic{
		Id:          topic.ID,
		Description: topic.Description,
	}
}

func MiniVideoModelToResponse(video *model.Video) *proto.MiniVideo {
	likeCount, _ := dao.GetVideoLikeCount(video.ID)
	return &proto.MiniVideo{
		Id:        video.ID,
		CoverUrl:  video.CoverUrl,
		LikeCount: likeCount,
	}
}

func CategoryModelToResponse(category *model.Category) *proto.Category {
	return &proto.Category{
		Id:   category.ID,
		Name: category.Name,
	}
}
