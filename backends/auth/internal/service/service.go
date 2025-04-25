package service

import (
	v1 "auth/internal/service/v1"
	ver "auth/internal/service/version"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(v1.NewAuthService, ver.NewVersionService)
