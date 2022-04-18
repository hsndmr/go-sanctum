package cryptoservice

import (
	"testing"
)

// TestGenerateSHA256 tests the generation of a SHA256 hash
func TestGenerateSHA256(t *testing.T) {
	c := Crypto{}
	h := c.SHA256("test")
	if h != "9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08" {
		t.Errorf("Error: %s", "SHA256 does not match")
	}
}

// TestHashPassword tests the hashing of a password
func TestHashPassword(t *testing.T) {
	c := Crypto{}
	h := c.HashPassword("1")

	if !c.CheckPasswordHash("1", h) {
		t.Errorf("Error: %s", "Password does not match")
	}
}