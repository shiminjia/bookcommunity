package utils

import (
	"fmt"
	"github.com/shiminjia/bookcommunity/config"
	"golang.org/x/crypto/scrypt"
)

func Encrypt(password string) (string, error) {
	//Encrypt password and change byte[] to string
	dk, err := scrypt.Key([]byte(password), []byte(config.SALT), 16384, 8, 1, 32)
	if err != nil {
		return "", err
	}
	dkStr := fmt.Sprintf("%x", dk)
	return dkStr, nil
}
