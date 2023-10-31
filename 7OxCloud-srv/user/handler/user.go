package handler

import (
	"context"
	"fmt"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/palp1tate/7OxCloud/7OxCloud-srv/user/dao"
	"github.com/palp1tate/7OxCloud/7OxCloud-srv/user/global"
	"github.com/palp1tate/7OxCloud/model"
	"github.com/palp1tate/7OxCloud/proto/user"
	"github.com/palp1tate/go-crypto-guard/pbkdf2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserServer struct {
	userProto.UnimplementedUserServiceServer
}

func (s *UserServer) CheckMobile(ctx context.Context, req *userProto.CheckMobileRequest) (*userProto.CheckMobileResponse, error) {
	_, err := dao.FindUserByMobile(req.Mobile)
	res := &userProto.CheckMobileResponse{
		Exist: err == nil,
	}
	return res, nil
}

func (s *UserServer) Register(ctx context.Context, req *userProto.RegisterRequest) (*userProto.RegisterResponse, error) {
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
	res := &userProto.RegisterResponse{
		Id: user.ID,
	}
	return res, nil
}

func (s *UserServer) LoginByPassword(ctx context.Context, req *userProto.LoginByPasswordRequest) (*userProto.LoginResponse, error) {
	user, err := dao.FindUserByMobile(req.Mobile)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "该手机号未注册")
	}
	if ok, _ := pwd.VerifySHA512(req.Password, user.Password); !ok {
		return nil, status.Errorf(codes.Unauthenticated, "密码错误")
	}
	res := &userProto.LoginResponse{
		Id: user.ID,
	}
	return res, nil
}

func (s *UserServer) LoginBySMS(ctx context.Context, req *userProto.LoginBySMSRequest) (*userProto.LoginResponse, error) {
	user, _ := dao.FindUserByMobile(req.Mobile)
	res := &userProto.LoginResponse{
		Id: user.ID,
	}
	return res, nil
}

func (s *UserServer) GetUser(ctx context.Context, req *userProto.GetUserRequest) (*userProto.GetUserResponse, error) {
	user, err := dao.FindUserById(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "该用户不存在")
	}
	res := &userProto.GetUserResponse{
		User: UserModelToResponse(user, req.CurrentUserId),
	}
	return res, nil
}

func (s *UserServer) UpdateUser(ctx context.Context, req *userProto.UpdateUserRequest) (*empty.Empty, error) {
	user, err := dao.FindUserById(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "该用户不存在")
	}
	user.Username = req.Username
	user.Age = int(req.Age)
	user.Location = req.Location
	user.Avatar = req.Avatar
	user.Gender = int(req.Gender)
	user.Signature = req.Signature
	err = dao.UpdateUser(&user)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "个人信息更新失败")
	}
	return &empty.Empty{}, nil
}

func (s *UserServer) ResetPassword(ctx context.Context, req *userProto.ResetPasswordRequest) (*empty.Empty, error) {
	user, err := dao.FindUserByMobile(req.Mobile)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "该手机号未注册")
	}
	password, _ := pwd.GenSHA512(req.Password, 16, 32, 50)
	err = dao.UpdateUserPassword(&user, password)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "修改密码失败")
	}
	return &empty.Empty{}, nil
}

func UserModelToResponse(user model.User, currentUserId int64) *userProto.User {
	tx := global.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	uid := user.ID
	followCount, _ := dao.GetFollowCount(tx, uid)
	fanCount, _ := dao.GetFanCount(tx, uid)
	isFollow, _ := dao.GetIsFollow(tx, uid, currentUserId)
	likeCount, _ := dao.GetLikeCount(tx, uid)
	likedCount, _ := dao.GetLikedCount(tx, uid)
	workCount, _ := dao.GetWorkCount(tx, uid)
	collectCount, _ := dao.GetCollectCount(tx, uid)
	if err := tx.Commit().Error; err != nil {
		return nil
	}
	return &userProto.User{
		Id:           user.ID,
		Age:          int64(user.Age),
		Username:     user.Username,
		Location:     user.Location,
		Avatar:       user.Avatar,
		Gender:       int64(user.Gender),
		Signature:    user.Signature,
		FollowCount:  followCount,
		FanCount:     fanCount,
		IsFollow:     isFollow,
		LikeCount:    likeCount,
		LikedCount:   likedCount,
		WorkCount:    workCount,
		CollectCount: collectCount,
	}
}
