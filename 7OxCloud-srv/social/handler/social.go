package handler

import (
	"context"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/palp1tate/7OxCloud/7OxCloud-srv/social/dao"
	"github.com/palp1tate/7OxCloud/7OxCloud-srv/social/global"
	"github.com/palp1tate/7OxCloud/model"
	"github.com/palp1tate/7OxCloud/proto/social"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type SocialServer struct {
	socialProto.UnimplementedSocialServiceServer
}

func (s *SocialServer) Follow(ctx context.Context, req *socialProto.FollowRequest) (*empty.Empty, error) {
	err := dao.Follow(req.UserId, req.FanId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "关注失败")
	}
	return &empty.Empty{}, nil
}

func (s *SocialServer) CancelFollow(ctx context.Context, req *socialProto.CancelFollowRequest) (*empty.Empty, error) {
	err := dao.CancelFollow(req.UserId, req.FanId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "取消关注失败")
	}
	return &empty.Empty{}, nil
}

func (s *SocialServer) GetFollowing(ctx context.Context, req *socialProto.GetFollowingRequest) (*socialProto.GetFollowingResponse, error) {
	followingList, count, err := dao.GetFollowing(req.CurrentUserId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "获取关注列表失败")
	}
	followings := make([]*socialProto.Following, len(followingList))
	for i, following := range followingList {
		followings[i] = UserModelToFollowingResponse(following, req.CurrentUserId)
	}
	return &socialProto.GetFollowingResponse{
		Followings: followings,
		Count:      count,
	}, nil
}

func (s *SocialServer) SearchFollowing(ctx context.Context, req *socialProto.SearchRequest) (*socialProto.SearchFollowingResponse, error) {
	followingList, err := dao.SearchFollowing(req.CurrentUserId, req.Keyword)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "搜索关注列表失败")
	}
	followings := make([]*socialProto.Following, len(followingList))
	for i, following := range followingList {
		followings[i] = UserModelToFollowingResponse(following, req.CurrentUserId)
	}
	return &socialProto.SearchFollowingResponse{
		Followings: followings,
	}, nil
}

func (s *SocialServer) GetFan(ctx context.Context, req *socialProto.GetFanRequest) (*socialProto.GetFanResponse, error) {
	fanList, count, err := dao.GetFan(req.CurrentUserId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "获取粉丝列表失败")
	}
	if len(fanList) == 0 {
		return &socialProto.GetFanResponse{}, nil
	}
	fans := make([]*socialProto.Fan, len(fanList))
	for i, fan := range fanList {
		fans[i] = UserModelToFanResponse(fan, req.CurrentUserId)
	}
	return &socialProto.GetFanResponse{
		Fans:  fans,
		Count: count,
	}, nil
}

func (s *SocialServer) SearchFan(ctx context.Context, req *socialProto.SearchRequest) (*socialProto.SearchFanResponse, error) {
	fanList, err := dao.SearchFan(req.CurrentUserId, req.Keyword)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "搜索粉丝列表失败")
	}
	fans := make([]*socialProto.Fan, len(fanList))
	for i, fan := range fanList {
		fans[i] = UserModelToFanResponse(fan, req.CurrentUserId)
	}
	return &socialProto.SearchFanResponse{
		Fans: fans,
	}, nil
}

func (s *SocialServer) FollowFeed(ctx context.Context, req *socialProto.FollowFeedRequest) (*socialProto.FeedResponse, error) {
	latestTime := time.Unix(req.LatestTime, 0)
	videoList, err := dao.GetFollowFeed(latestTime, req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "获取关注动态失败")
	}
	if len(videoList) == 0 {
		return &socialProto.FeedResponse{}, nil
	}
	videos := make([]*socialProto.Video, len(videoList))
	for i, video := range videoList {
		videos[i] = VideoModelToResponse(video, req.UserId)
	}
	return &socialProto.FeedResponse{
		Videos:   videos,
		NextTime: videoList[len(videoList)-1].CreatedAt.Unix(),
	}, nil
}

func (s *SocialServer) FriendFeed(ctx context.Context, req *socialProto.FriendFeedRequest) (*socialProto.FeedResponse, error) {
	latestTime := time.Unix(req.LatestTime, 0)
	videoList, err := dao.GetFriendFeed(latestTime, req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "获取好友动态失败")
	}
	if len(videoList) == 0 {
		return &socialProto.FeedResponse{}, nil
	}
	videos := make([]*socialProto.Video, len(videoList))
	for i, video := range videoList {
		videos[i] = VideoModelToResponse(video, req.UserId)
	}
	return &socialProto.FeedResponse{
		Videos:   videos,
		NextTime: videoList[len(videoList)-1].CreatedAt.Unix(),
	}, nil
}

func TopicModelToResponse(topic *model.Topic) *socialProto.Topic {
	return &socialProto.Topic{
		Id:          topic.ID,
		Description: topic.Description,
	}
}

func VideoModelToResponse(video *model.Video, currentUserId int64) *socialProto.Video {
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
	topics := make([]*socialProto.Topic, len(topicList))
	for i, topic := range topicList {
		topics[i] = TopicModelToResponse(topic)
	}
	if err := tx.Commit().Error; err != nil {
		return nil
	}
	return &socialProto.Video{
		Id:           vid,
		Description:  video.Description,
		PlayUrl:      video.PlayUrl,
		CoverUrl:     video.CoverUrl,
		LikeCount:    likeCount,
		CommentCount: commentCount,
		CollectCount: collectCount,
		IsLike:       isLike,
		IsCollect:    isCollect,
		Author:       UserModelToResponse(author),
		Topics:       topics,
		PublishTime:  video.CreatedAt.Unix(),
	}
}

func UserModelToResponse(user model.User) *socialProto.User {
	return &socialProto.User{
		Id:       user.ID,
		Username: user.Username,
		Avatar:   user.Avatar,
		Location: user.Location,
	}
}

func UserModelToFollowingResponse(user model.User, currentUserId int64) *socialProto.Following {
	isFan, _ := dao.GetIsFan(user.ID, currentUserId)
	return &socialProto.Following{
		Id:        user.ID,
		Username:  user.Username,
		Avatar:    user.Avatar,
		Signature: user.Signature,
		IsFan:     isFan,
	}
}

func UserModelToFanResponse(user model.User, currentUserId int64) *socialProto.Fan {
	isFollow, _ := dao.GetIsFollow(user.ID, currentUserId)
	return &socialProto.Fan{
		Id:        user.ID,
		Username:  user.Username,
		Avatar:    user.Avatar,
		Signature: user.Signature,
		IsFollow:  isFollow,
	}
}
