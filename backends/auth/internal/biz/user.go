package biz

import (
	"context"
	"time"
)

type User struct {
	ID        int64     `gorm:"primaryKey;autoIncrement;column:id"`
	UID       int64     `gorm:"column:uid;not null;default:0"`
	Username  string    `gorm:"column:username;type:varchar(128);not null;default:''"`
	HashPwd   string    `gorm:"column:hash_pwd;type:varchar(512);not null;default:''"`
	IsDeleted int       `gorm:"column:is_deleted;default:0"`
	CTime     time.Time `gorm:"column:ctime;autoCreateTime"`
	MTime     time.Time `gorm:"column:mtime;autoUpdateTime"`
}

func (User) TableName() string {
	return "user"
}

type UserRepo interface {
	UserByID(ctx context.Context, uid int64) (*User, error)
}

type UserUseCase struct {
	repo UserRepo
}

func NewUserUseCase(repo UserRepo) *UserUseCase {
	return &UserUseCase{repo: repo}
}

func (uc *UserUseCase) GetUserByID(ctx context.Context, uid int64) (user *User, err error) {
	user, err = uc.repo.UserByID(ctx, uid)
	if err != nil {
		return
	}

	return
}
