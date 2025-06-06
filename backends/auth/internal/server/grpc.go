package server

import (
	v1 "auth/api/auth/v1"
	apiVer "auth/api/version"
	"auth/internal/conf"
	v1svc "auth/internal/service/v1"
	verSvc "auth/internal/service/version"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Server, auth *v1svc.AuthService, version *verSvc.VersionService, logger log.Logger) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	apiVer.RegisterVersionServer(srv, version)
	v1.RegisterAuthServer(srv, auth)
	return srv
}
