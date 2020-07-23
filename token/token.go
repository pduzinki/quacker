package token

import (
	"crypto/rand"
	"encoding/base64"
)

const RememberTokenLength = 32

func GenerateRememberToken() (string, error) {
	randomBytes := make([]byte, RememberTokenLength)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(randomBytes), nil
}

func GetTokenLength(base64string string) (int, error) {
	b, err := base64.URLEncoding.DecodeString(base64string)
	if err != nil {
		return -1, err
	}
	return len(b), nil
}
