package service

import (
	"auth/internal/global"
	"context"

	pb "auth/api/version"
)

type VersionService struct {
	pb.UnimplementedVersionServer
}

func NewVersionService() *VersionService {
	return &VersionService{}
}

func (s *VersionService) Version(ctx context.Context, req *pb.VersionReq) (*pb.VersionReply, error) {
	return &pb.VersionReply{
		Name:    global.NAME,
		Version: global.VERSION,
	}, nil
}
