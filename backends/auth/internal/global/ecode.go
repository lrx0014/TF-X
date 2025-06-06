package global

import "errors"

var ErrInvalidDSN = errors.New("invalid DSN")
var ErrFailedConnect = errors.New("failed to connect")
var ErrDBInternalError = errors.New("database internal error")

var ErrUsernameValidation = errors.New("username validation error, too short or too long")
var ErrPasswordValidation = errors.New("password validation error, too short or too long")
var ErrUserNotFound = errors.New("user not found")
var ErrGenerateUIDFailed = errors.New("failed to generate uid")
var ErrHashFailed = errors.New("failed to hash password")
var ErrPasswordMismatch = errors.New("password mismatch")
