package global

import "github.com/go-kratos/kratos/v2/errors"

// db error (reserved code from 100 - 199)

var ErrInvalidDSN = errors.New(100, "invalid DSN", "invalid DSN")
var ErrFailedConnect = errors.New(101, "failed to connect", "failed to connect")

var ErrUserNotFound = errors.New(404, "user not found", "user not found")
