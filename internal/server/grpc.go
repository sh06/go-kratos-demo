package server

import (
	articleV1 "kratos-demo/api/article/v1"
	v1 "kratos-demo/api/helloworld/v1"
	tagV1 "kratos-demo/api/tag/v1"
	userV1 "kratos-demo/api/user/v1"
	"kratos-demo/internal/conf"
	"kratos-demo/internal/pkg/middleware/auth"
	"kratos-demo/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(
	c *conf.Server,
	greeter *service.GreeterService,
	user *service.UserService,
	article *service.ArticleService,
	tag *service.TagService,
	logger log.Logger,
) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
			auth.JwtAuth("secret"),
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	v1.RegisterGreeterServer(srv, greeter)
	userV1.RegisterUserServer(srv, user)
	articleV1.RegisterArticleServer(srv, article)
	tagV1.RegisterTagServer(srv, tag)
	return srv
}
