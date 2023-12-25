package data

import (
	"context"
	"kratos-demo/internal/biz"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

// user 表模型定义
type User struct {
	ID        uint64 `gorm:"column:id"`
	Username  string `gorm:"column:username"`
	Password  string `gorm:"column:password"`
	Email     string `gorm:"column:email"`
	Bio       string `gorm:"column:bio"`
	Image     string `gorm:"column:image"`
	CreatedAt string `gorm:"column:created_at"`
	UpdatedAt string `gorm:"column:updated_at"`
}

func (u *User) TableName() string {
	return "users"
}

type userRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *userRepo) GetByEmail(ctx context.Context, email string) (*biz.User, error) {
	var user User
	result := r.data.db.Where("email = ?", email).First(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return &biz.User{
		ID:        user.ID,
		Username:  user.Username,
		Password:  user.Password,
		Email:     user.Email,
		Bio:       user.Bio,
		Image:     user.Image,
		CreatedAt: user.CreatedAt,
	}, nil
}

func (r *userRepo) Save(ctx context.Context, ur *biz.RegisterReq) (uint64, error) {
	currentTime := time.Now().Format("2006-01-02 15:04:05")

	user := User{
		Username:  ur.Username,
		Password:  ur.Password,
		Email:     ur.Email,
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
	}

	result := r.data.db.Create(&user)

	if result.Error != nil {
		return 0, result.Error
	}

	return user.ID, nil
}
