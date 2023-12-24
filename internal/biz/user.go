package biz

import (
	"context"
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
)

type UserUsecase struct {
	log *log.Helper
}

func NewUserUsecase(logger log.Logger) *UserUsecase {
	return &UserUsecase{log: log.NewHelper(logger)}
}

type RegisterReq struct {
	Email    string
	Password string
	Username string
}

func (uc *UserUsecase) Register(ctx context.Context, registerReq *RegisterReq) error {
	fmt.Println(registerReq)
	return nil
}
