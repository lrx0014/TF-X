package service

import (
	"auth/internal/biz"
	"auth/internal/conf"
	"auth/internal/global"
	"context"

	pb "auth/api/auth/v1"
	_ "github.com/lrx0014/ScalableFlake/pkg/driver/redis"
)

type AuthService struct {
	pb.UnimplementedAuthServer
	appConf *conf.App

	uc *biz.UserUseCase
}

func NewAuthService(uc *biz.UserUseCase, appCond *conf.App) *AuthService {
	return &AuthService{
		appConf: appCond,
		uc:      uc,
	}
}

func (s *AuthService) CreateAccount(ctx context.Context, req *pb.CreateAccountReq) (resp *pb.CreateAccountResp, err error) {
	resp = &pb.CreateAccountResp{}
	// validate request params
	if int64(len(req.Username)) < s.appConf.ParamValidator.UsernameMin || int64(len(req.Username)) > s.appConf.ParamValidator.UsernameMax {
		err = global.ErrUsernameValidation
		return
	}
	if int64(len(req.Pwd)) < s.appConf.ParamValidator.PasswordMin || int64(len(req.Pwd)) > s.appConf.ParamValidator.PasswordMax {
		err = global.ErrPasswordValidation
		return
	}

	user := &biz.User{
		UID:      0, // generate in biz
		Username: req.Username,
		Pwd:      req.Pwd,
	}
	uid, err := s.uc.CreateUser(ctx, user)
	if err != nil {
		return
	}

	resp.Uid = int64(uid)
	return
}

func (s *AuthService) Login(ctx context.Context, req *pb.LoginReq) (resp *pb.LoginResp, err error) {
	resp = &pb.LoginResp{}
	accessToken, refreshToken, err := s.uc.Login(ctx, req.GetUsername(), req.GetPwd())
	if err != nil {
		return
	}

	resp.AccessToken = accessToken
	resp.RefreshToken = refreshToken
	return
}

func (s *AuthService) DeleteAccount(ctx context.Context, req *pb.DeleteAccountReq) (*pb.DeleteAccountResp, error) {
	return &pb.DeleteAccountResp{}, nil
}

func (s *AuthService) LockAccount(ctx context.Context, req *pb.LockAccountReq) (*pb.LockAccountResp, error) {
	return &pb.LockAccountResp{}, nil
}

func (s *AuthService) UnlockAccount(ctx context.Context, req *pb.UnlockAccountReq) (*pb.UnlockAccountResp, error) {
	return &pb.UnlockAccountResp{}, nil
}
