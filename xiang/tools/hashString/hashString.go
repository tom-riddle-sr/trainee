package hashString

import (
	"crypto/sha512"
	"encoding/hex"
)

func HashString(s string) string {
	hash := sha512.Sum512([]byte(s))
	return hex.EncodeToString(hash[:])
}
