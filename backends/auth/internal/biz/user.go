package biz

import (
	"auth/internal/conf"
	"auth/internal/global"
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

type User struct {
	ID        int64     `gorm:"primaryKey;autoIncrement;column:id"`
	UID       int64     `gorm:"column:uid;not null;default:0"`
	Username  string    `gorm:"column:username;type:varchar(128);not null;default:''"`
	Pwd       string    `gorm:"column:pwd;type:varchar(512);not null;default:''"`
	IsDeleted int       `gorm:"column:is_deleted;default:0"`
	CTime     time.Time `gorm:"column:ctime;autoCreateTime"`
	MTime     time.Time `gorm:"column:mtime;autoUpdateTime"`
}

func (User) TableName() string {
	return "user"
}

type UserRepo interface {
	UserByID(ctx context.Context, uid int64) (user *User, err error)
	CreateUser(ctx context.Context, user *User) (err error)
}

type UserUseCase struct {
	appConf *conf.App
	repo    UserRepo
}

func NewUserUseCase(repo UserRepo, appConf *conf.App) *UserUseCase {
	return &UserUseCase{
		appConf: appConf,
		repo:    repo,
	}
}

func (uc *UserUseCase) CreateUser(ctx context.Context, user *User) (err error) {
	err = uc.repo.CreateUser(ctx, user)
	if err != nil {
		log.Errorf("[biz] CreateUser err: %+v", err.Error())
		err = global.ErrDBInternalError
		return
	}

	return
}

func (uc *UserUseCase) GetUserByID(ctx context.Context, uid int64) (user *User, err error) {
	user, err = uc.repo.UserByID(ctx, uid)
	if err != nil {
		log.Errorf("[biz] GetUserByID err: %+v", err.Error())
		err = global.ErrDBInternalError
		return
	}

	return
}
