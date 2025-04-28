package data

import (
	"auth/internal/biz"
	"context"
)

type userRepo struct {
	data *Data
}

func NewUserRepo(data *Data) biz.UserRepo {
	return &userRepo{
		data: data,
	}
}

func (u *userRepo) UserByID(ctx context.Context, uid int64) (user *biz.User, err error) {
	err = u.data.db.First(&user, "uid = ?", uid).Error
	return
}

func (u *userRepo) CreateUser(ctx context.Context, user *biz.User) (err error) {
	err = u.data.db.Create(&user).Error
	return
}
