package biz

import (
	"auth/internal/conf"
	"auth/internal/global"
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/lrx0014/ScalableFlake/pkg/snowflake"
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
	UserByUID(ctx context.Context, uid int64) (user *User, err error)
	UserByUsername(ctx context.Context, username string) (user *User, err error)
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

func (uc *UserUseCase) CreateUser(ctx context.Context, user *User) (uid uint64, err error) {
	// generate an UID
	uid, err = snowflake.GenerateUID("auth")
	if err != nil {
		err = global.ErrGenerateUIDFailed
		return
	}
	user.UID = int64(uid)

	// password hash
	hashedPwd, err := hashPassword(user.Pwd)
	if err != nil {
		log.Errorf("hash password failed: %v", err)
		err = global.ErrHashFailed
		return
	}
	user.Pwd = hashedPwd

	err = uc.repo.CreateUser(ctx, user)
	if err != nil {
		log.Errorf("[biz] CreateUser err: %+v", err.Error())
		err = global.ErrDBInternalError
		return
	}

	return
}

func (uc *UserUseCase) Login(ctx context.Context, username, pwd string) (accessToken, refreshToken string, err error) {
	user, err := uc.GetUserByUsername(ctx, username)
	if err != nil {
		return
	}

	matched := verifyPassword(pwd, user.Pwd)
	if !matched {
		log.Errorf("[biz] Login error: %s != %s", pwd, user.Pwd)
		err = global.ErrPasswordMismatch
		return
	}

	accessToken, err = generateJWT(user.UID, user.Username, TypeAccess, AccessTokenTTL, accessTokenSecret)
	if err != nil {
		log.Errorf("failed to generate access token: %v", err)
		return
	}

	refreshToken, err = generateJWT(user.UID, user.Username, TypeRefresh, RefreshTokenTTL, refreshTokenSecret)
	if err != nil {
		log.Errorf("failed to generate refresh token: %v", err)
		return
	}

	return
}

func (uc *UserUseCase) GetUserByUID(ctx context.Context, uid int64) (user *User, err error) {
	user, err = uc.repo.UserByUID(ctx, uid)
	if err != nil {
		log.Errorf("[biz] GetUserByID err: %+v", err.Error())
		err = global.ErrDBInternalError
		return
	}

	return
}

func (uc *UserUseCase) GetUserByUsername(ctx context.Context, username string) (user *User, err error) {
	user, err = uc.repo.UserByUsername(ctx, username)
	if err != nil {
		log.Errorf("[biz] GetUserByUsername err: %+v", err.Error())
		err = global.ErrDBInternalError
		return
	}

	return
}
