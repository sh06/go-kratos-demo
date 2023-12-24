package service

import (
	"context"

	pb "kratos-demo/api/user/v1"
	"kratos-demo/internal/biz"
)

type UserService struct {
	pb.UnimplementedUserServer

	uc *biz.UserUsecase
}

func NewUserService(uc *biz.UserUsecase) *UserService {
	return &UserService{
		uc: uc,
	}
}

func (s *UserService) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.OperateReply, error) {
	error := s.uc.Register(ctx, &biz.RegisterReq{
		Email:    req.User.Email,
		Password: req.User.Password,
		Username: req.User.Username,
	})

	ret := true
	if error != nil {
		ret = false
	}

	return &pb.OperateReply{
		Success: ret,
	}, nil
}
func (s *UserService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginReply, error) {
	return &pb.LoginReply{}, nil
}
func (s *UserService) GetCurrentUser(ctx context.Context, req *pb.GetCurrentUserRequest) (*pb.UserReply, error) {
	return &pb.UserReply{}, nil
}
func (s *UserService) GetUserById(ctx context.Context, req *pb.GetUserByIdRequest) (*pb.UserReply, error) {
	return &pb.UserReply{}, nil
}
func (s *UserService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UserReply, error) {
	return &pb.UserReply{}, nil
}
func (s *UserService) FollowUser(ctx context.Context, req *pb.FollowUserRequest) (*pb.UserReply, error) {
	return &pb.UserReply{}, nil
}
func (s *UserService) UnfollowUser(ctx context.Context, req *pb.UnfollowUserRequest) (*pb.UserReply, error) {
	return &pb.UserReply{}, nil
}
