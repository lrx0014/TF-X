package biz

import "github.com/google/wire"

// ProviderBiz is biz providers.
var ProviderBiz = wire.NewSet(NewUserUseCase)
