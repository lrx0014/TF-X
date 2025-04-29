package biz

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	accessTokenSecret  = []byte("your-access-token-secret")
	refreshTokenSecret = []byte("your-refresh-token-secret")
)

const (
	TypeAccess      = "access"
	TypeRefresh     = "refresh"
	AccessTokenTTL  = time.Minute * 15
	RefreshTokenTTL = time.Hour * 24 * 7
)

type CustomClaims struct {
	UserID    int64  `json:"user_id"`
	Username  string `json:"username"`
	TokenType string `json:"type"` // "access" or "refresh"
	jwt.RegisteredClaims
}

func generateJWT(userID int64, username, tokenType string, ttl time.Duration, secret []byte) (string, error) {
	claims := CustomClaims{
		UserID:    userID,
		Username:  username,
		TokenType: tokenType,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ttl)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secret)
}
