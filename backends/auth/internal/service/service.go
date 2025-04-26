package service

import (
	v1 "auth/internal/service/v1"
	ver "auth/internal/service/version"
	"github.com/google/wire"
)

// ProviderService is service providers.
var ProviderService = wire.NewSet(v1.NewAuthService, ver.NewVersionService)
