package service

import (
	"github.com/google/wire"

	version "oss-mng/internal/service/version"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(version.NewVersionService)
