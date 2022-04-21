package cryptoservice

import (
	"crypto/sha256"
	"encoding/hex"

	"golang.org/x/crypto/bcrypt"
)
type CryptoI interface {
	SHA256(plainText string) string
	HashPassword(password string) string
	CheckPasswordHash(password, hash string) bool
}

type Crypto struct {}

// SHA256 Generates the SHA hash of a string
func (c *Crypto) SHA256(plainText string) string {
	hash := sha256.Sum256([]byte(plainText))
	return hex.EncodeToString(hash[:])
}

// GeneratePasswordHash generates a hash of a password
func (c *Crypto) HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return ""
	}
	return string(bytes)
}

// CheckPasswordHash checks if a password is correct
func (c *Crypto) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}