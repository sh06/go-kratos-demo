package service

import (
	"context"

	pb "kratos-demo/api/tag/v1"
)

type TagService struct {
	pb.UnimplementedTagServer
}

func NewTagService() *TagService {
	return &TagService{}
}

func (s *TagService) CreateTag(ctx context.Context, req *pb.CreateTagRequest) (*pb.CreateTagReply, error) {
	return &pb.CreateTagReply{}, nil
}
func (s *TagService) ListTag(ctx context.Context, req *pb.ListTagRequest) (*pb.ListTagReply, error) {
	return &pb.ListTagReply{}, nil
}
