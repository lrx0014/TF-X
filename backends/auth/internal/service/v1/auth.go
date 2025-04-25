package service

import (
	"context"

	pb "auth/api/auth/v1"
)

type AuthService struct {
	pb.UnimplementedAuthServer
}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (s *AuthService) Signup(ctx context.Context, req *pb.SignupReq) (*pb.SignupResp, error) {
	return &pb.SignupResp{}, nil
}
func (s *AuthService) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginResp, error) {
	return &pb.LoginResp{}, nil
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
