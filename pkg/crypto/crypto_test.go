package crypto

import (
	"testing"
)

// TestGenerateSHA256 is a test for SHA256 function
func TestGenerateSHA256(t *testing.T) {
	h := SHA256("test")
	if h != "9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08" {
		t.Errorf("Error: %s", "SHA256 does not match")
	}
}