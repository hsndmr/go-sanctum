package random

import (
	"testing"
)

// TestGenerateRandomString is a test for Random.Make function
func TestGenerateRandomString(t *testing.T) {
	s, err := GenerateString(40)
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	if len(s) != 40 {
		t.Errorf("Expected length 40, got %d", len(s))
	}
}