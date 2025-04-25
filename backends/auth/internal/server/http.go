package server

import (
	v1 "auth/api/auth/v1"
	apiVer "auth/api/version"
	"auth/internal/conf"
	v1svc "auth/internal/service/v1"
	ver "auth/internal/service/version"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, auth *v1svc.AuthService, version *ver.VersionService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	apiVer.RegisterVersionHTTPServer(srv, version)
	v1.RegisterAuthHTTPServer(srv, auth)
	return srv
}
