package crypto

import (
	"crypto/sha256"
	"encoding/hex"
)

// SHA256 Generates the SHA hash of a string
func SHA256(plainText string) string {
	hash := sha256.Sum256([]byte(plainText))
	return hex.EncodeToString(hash[:])
}