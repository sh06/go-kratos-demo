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
	if err := s.uc.Register(ctx, &biz.RegisterReq{
		Email:    req.User.Email,
		Password: req.User.Password,
		Username: req.User.Username,
	}); err != nil {
		return nil, err
	}

	return &pb.OperateReply{
		Success: true,
	}, nil
}

func (s *UserService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginReply, error) {
	user, err := s.uc.Login(ctx, &biz.LoginReq{
		Email:    req.User.Email,
		Password: req.User.Password,
	})

	if err != nil {
		return nil, err
	}

	return &pb.LoginReply{
		User: &pb.LoginReply_User{
			Id:        int32(user.ID),
			Email:     user.Email,
			Token:     user.Token,
			Username:  user.Username,
			Bio:       user.Bio,
			Image:     user.Image,
			Following: 0,
			Followers: 0,
		},
	}, nil
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
