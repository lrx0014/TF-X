package biz

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"golang.org/x/crypto/scrypt"
	"strings"
)

const (
	saltLen = 16
	N       = 1 << 15 // CPU/memory cost parameter
	r       = 8       // Block size parameter
	p       = 1       // Parallelization parameter
	keyLen  = 32      // Desired key length
)

func generateSalt() ([]byte, error) {
	salt := make([]byte, saltLen)
	_, err := rand.Read(salt)
	return salt, err
}

func hashPassword(password string) (hashed string, err error) {
	salt, err := generateSalt()
	if err != nil {
		return
	}

	dk, err := scrypt.Key([]byte(password), salt, N, r, p, keyLen)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s:%s",
		base64.StdEncoding.EncodeToString(salt),
		base64.StdEncoding.EncodeToString(dk)), nil
}

func verifyPassword(password, encoded string) bool {
	el := strings.Split(encoded, ":")

	saltEncoded := el[0]
	hashEncoded := el[1]

	salt, err := base64.StdEncoding.DecodeString(saltEncoded)
	if err != nil {
		log.Errorf("[biz] DecodeString error: %s", err)
		return false
	}

	dk, err := scrypt.Key([]byte(password), salt, N, r, p, keyLen)
	if err != nil {
		log.Errorf("[biz] scrypt error: %s", err)
		return false
	}

	return base64.StdEncoding.EncodeToString(dk) == hashEncoded
}
