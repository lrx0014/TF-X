package service

import (
	"auth/internal/biz"
	"auth/internal/global"
	"context"

	pb "auth/api/auth/v1"
)

type AuthService struct {
	pb.UnimplementedAuthServer

	uc *biz.UserUseCase
}

func NewAuthService(uc *biz.UserUseCase) *AuthService {
	return &AuthService{
		uc: uc,
	}
}

func (s *AuthService) Signup(ctx context.Context, req *pb.SignupReq) (*pb.SignupResp, error) {
	return &pb.SignupResp{}, nil
}
func (s *AuthService) Login(ctx context.Context, req *pb.LoginReq) (resp *pb.LoginResp, err error) {
	resp = &pb.LoginResp{}
	user, err := s.uc.GetUserByID(ctx, req.Uid)
	if err != nil {
		return
	}

	if user == nil || user.ID == 0 {
		err = global.ErrUserNotFound
		return
	}

	// TODO
	resp.AccessToken = user.Username
	resp.RefreshToken = user.Username

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
