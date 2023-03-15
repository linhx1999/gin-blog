package utils

import (
	"encoding/base64"
	"golang.org/x/crypto/scrypt"
	"log"
)

func ScryptPassword(password string) string {
	salt := make([]byte, 8, 8)
	salt = []byte{19, 99, 07, 31, 03, 07, 06, 12}

	dk, err := scrypt.Key([]byte(password), salt, 1<<15, 8, 1, 32)
	if err != nil {
		log.Fatal(err)
	}
	return base64.StdEncoding.EncodeToString(dk)
}
