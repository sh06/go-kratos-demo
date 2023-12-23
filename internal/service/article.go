package service

import (
	"context"

	pb "kratos-demo/api/article/v1"
)

type ArticleService struct {
	pb.UnimplementedArticleServer
}

func NewArticleService() *ArticleService {
	return &ArticleService{}
}

func (s *ArticleService) CreateArticle(ctx context.Context, req *pb.CreateArticleRequest) (*pb.SingleArticleReply, error) {
	return &pb.SingleArticleReply{}, nil
}
func (s *ArticleService) UpdateArticle(ctx context.Context, req *pb.UpdateArticleRequest) (*pb.SingleArticleReply, error) {
	return &pb.SingleArticleReply{}, nil
}
func (s *ArticleService) DeleteArticle(ctx context.Context, req *pb.DeleteArticleRequest) (*pb.OperateReply, error) {
	return &pb.OperateReply{}, nil
}
func (s *ArticleService) GetArticle(ctx context.Context, req *pb.GetArticleRequest) (*pb.SingleArticleReply, error) {
	return &pb.SingleArticleReply{}, nil
}
func (s *ArticleService) ListArticle(ctx context.Context, req *pb.ListArticleRequest) (*pb.MultipleArticlesReply, error) {
	return &pb.MultipleArticlesReply{}, nil
}
