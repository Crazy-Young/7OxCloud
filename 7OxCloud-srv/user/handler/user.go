package handler

import (
	"context"
	"fmt"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/palp1tate/7OxCloud/7OxCloud-srv/user/dao"
	"github.com/palp1tate/7OxCloud/7OxCloud-srv/user/model"
	"github.com/palp1tate/7OxCloud/7OxCloud-srv/user/proto"
	"github.com/palp1tate/go-crypto-guard/pbkdf2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserServer struct {
	proto.UnimplementedUserServiceServer
}

func (s *UserServer) CheckMobile(ctx context.Context, req *proto.CheckMobileRequest) (*proto.CheckMobileResponse, error) {
	_, err := dao.FindUserByMobile(req.Mobile)
	res := &proto.CheckMobileResponse{
		Exist: err == nil,
	}
	return res, nil
}

func (s *UserServer) Register(ctx context.Context, req *proto.RegisterRequest) (*proto.RegisterResponse, error) {
	var user model.User
	user, err := dao.FindUserByMobile(req.Mobile)
	if err == nil {
		return nil, status.Errorf(codes.AlreadyExists, "该手机号已经被注册")
	}
	password, _ := pwd.GenSHA512(req.Password, 16, 32, 50)
	user = model.User{
		Mobile:   req.Mobile,
		Password: password,
		Username: fmt.Sprintf("手机用户%s", req.Mobile[7:]),
	}
	err = dao.CreateUser(&user)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "注册失败")
	}
	res := &proto.RegisterResponse{
		Id: user.ID,
	}
	return res, nil
}

func (s *UserServer) LoginByPassword(ctx context.Context, req *proto.LoginByPasswordRequest) (*proto.LoginResponse, error) {
	user, err := dao.FindUserByMobile(req.Mobile)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "该手机号未注册")
	}
	if ok, _ := pwd.VerifySHA512(req.Password, user.Password); !ok {
		return nil, status.Errorf(codes.Unauthenticated, "密码错误")
	}
	res := &proto.LoginResponse{
		Id: user.ID,
	}
	return res, nil
}

func (s *UserServer) LoginBySMS(ctx context.Context, req *proto.LoginBySMSRequest) (*proto.LoginResponse, error) {
	user, _ := dao.FindUserByMobile(req.Mobile)
	res := &proto.LoginResponse{
		Id: user.ID,
	}
	return res, nil
}

func (s *UserServer) GetUser(ctx context.Context, req *proto.GetUserRequest) (*proto.GetUserResponse, error) {
	user, err := dao.FindUserById(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "该用户不存在")
	}
	res := &proto.GetUserResponse{
		User: UserModelToResponse(user, req.CurrentUserId),
	}
	return res, nil
}

func (s *UserServer) UpdateUser(ctx context.Context, req *proto.UpdateUserRequest) (*empty.Empty, error) {
	user, err := dao.FindUserById(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "该用户不存在")
	}
	user.Username = req.Username
	user.Age = int(req.Age)
	user.Location = req.Location
	user.Avatar = req.Avatar
	user.BackgroundImage = req.BackgroundImage
	user.Signature = req.Signature
	err = dao.UpdateUser(&user)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "个人信息更新失败")
	}
	return &empty.Empty{}, nil
}

func UserModelToResponse(user model.User, currentUserId int64) *proto.User {
	uid := user.ID
	followCount, _ := dao.GetFollowCount(uid)
	fanCount, _ := dao.GetFanCount(uid)
	isFollow, _ := dao.GetIsFollow(uid, currentUserId)
	likeCount, _ := dao.GetLikeCount(uid)
	likedCount, _ := dao.GetLikedCount(uid)
	workCount, _ := dao.GetWorkCount(uid)
	collectCount, _ := dao.GetCollectCount(uid)
	return &proto.User{
		Id:              user.ID,
		Age:             int64(user.Age),
		Username:        user.Username,
		Location:        user.Location,
		Avatar:          user.Avatar,
		BackgroundImage: user.BackgroundImage,
		Signature:       user.Signature,
		FollowCount:     followCount,
		FanCount:        fanCount,
		IsFollow:        isFollow,
		LikeCount:       likeCount,
		LikedCount:      likedCount,
		WorkCount:       workCount,
		CollectCount:    collectCount,
	}
}
