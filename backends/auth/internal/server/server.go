package server

import (
	"github.com/google/wire"
)

// ProviderServer is server providers.
var ProviderServer = wire.NewSet(NewGRPCServer, NewHTTPServer)
