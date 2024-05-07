package tools

import (
	"crypto/sha512"
	"encoding/hex"
)

func Sha512(value string) string {
	hashValue := sha512.Sum512([]byte(value))
	return hex.EncodeToString(hashValue[:])
}
