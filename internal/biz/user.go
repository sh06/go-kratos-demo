package biz

import (
	"context"
	"errors"

	pb "kratos-demo/api/user/v1"
	"kratos-demo/internal/pkg/util"

	e "github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/spf13/cast"
	"gorm.io/gorm"
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
	GetByUsername(ctx context.Context, username string) (*User, error)
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
	if user, _ := uc.repo.GetByEmail(ctx, registerReq.Email); user != nil {
		return pb.ErrorErrorUserEmailExist("邮箱已注册")
	}

	if user, _ := uc.repo.GetByUsername(ctx, registerReq.Username); user != nil {
		return pb.ErrorErrorUserUsernameExist("用户名已注册")
	}

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
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, pb.ErrorErrorUserNotFound("用户不存在")
		}
		return nil, err
	}

	if !util.VerifyPassword(user.Password, loginReq.Password) {
		return nil, pb.ErrorErrorUserPasswordError("用户密码错误")
	}

	token, err := util.GenerateJWT(cast.ToString(user.ID), "secret")
	if err != nil {
		return nil, e.New(500, "ERROR_CREATE_TOKEN_FAILED", "系统错误，请稍后再试")
	}

	user.Token = token
	return user, nil
}
