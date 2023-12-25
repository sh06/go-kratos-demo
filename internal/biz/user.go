package biz

import (
	"kratos-demo/internal/pkg/util"

	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type User struct {
	ID        uint64
	Username  string
	Password  string
	Token     string
	Email     string
	Bio       string
	Image     string
	CreatedAt string
}

type UserUsecase struct {
	repo UserRepo
	log  *log.Helper
}

type UserRepo interface {
	Save(ctx context.Context, g *RegisterReq) (uint64, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
}

func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}

type RegisterReq struct {
	Email    string
	Password string
	Username string
}

func (uc *UserUsecase) Register(ctx context.Context, registerReq *RegisterReq) error {
	registerReq.Password = util.HashPassword(registerReq.Password)

	if _, err := uc.repo.Save(ctx, registerReq); err != nil {
		return err
	}

	return nil
}

type LoginReq struct {
	Email    string
	Password string
}

func (uc *UserUsecase) Login(ctx context.Context, loginReq *LoginReq) (*User, error) {
	user, err := uc.repo.GetByEmail(ctx, loginReq.Email)

	if err != nil {
		return nil, err
	}

	if !util.VerifyPassword(user.Password, loginReq.Password) {
		return nil, err
	}

	// 生成token

	return user, nil
}
