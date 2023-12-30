package server

import (
	"context"
	articleV1 "kratos-demo/api/article/v1"
	v1 "kratos-demo/api/helloworld/v1"
	tagV1 "kratos-demo/api/tag/v1"
	userV1 "kratos-demo/api/user/v1"
	"kratos-demo/internal/conf"
	"kratos-demo/internal/pkg/encode"
	"kratos-demo/internal/pkg/middleware/auth"
	"kratos-demo/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/transport/http"
)

func NewSkipListMatcher() selector.MatchFunc {
	SkipList := make(map[string]struct{})
	SkipList["/api.user.v1.User/Login"] = struct{}{}
	//SkipList["/api.user.v1.User/Register"] = struct{}{}
	return func(ctx context.Context, operation string) bool {
		if _, ok := SkipList[operation]; ok {
			return false
		}
		return true
	}
}

// NewHTTPServer new an HTTP server.
func NewHTTPServer(
	c *conf.Server,
	greeter *service.GreeterService,
	user *service.UserService,
	article *service.ArticleService,
	tag *service.TagService,
	logger log.Logger,
) *http.Server {
	var opts = []http.ServerOption{
		http.ResponseEncoder(encode.ResponseEncoder),
		http.ErrorEncoder(encode.ErrorEncoder),
		http.Middleware(
			recovery.Recovery(),
			selector.Server(auth.JwtAuth("secret")).Match(NewSkipListMatcher()).Build(),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	v1.RegisterGreeterHTTPServer(srv, greeter)
	userV1.RegisterUserHTTPServer(srv, user)
	articleV1.RegisterArticleHTTPServer(srv, article)
	tagV1.RegisterTagHTTPServer(srv, tag)
	return srv
}
