package auth

import (
	"context"
	"kratos-demo/internal/pkg/util"

	e "github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
)

func JwtAuth(secret string) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if tr, ok := transport.FromServerContext(ctx); ok {
				tokenString := tr.RequestHeader().Get("token")
				user, err := util.ParseJwt(tokenString, secret)
				if err != nil {
					return nil, e.New(401, "ERROR_AUTH_FAILED", "认证错误，请重新登录")
				}

				ctx = context.WithValue(ctx, "user_id", user.UserID)

				defer func() {
					// Do something on exiting
				}()
			}
			return handler(ctx, req)
		}
	}
}
