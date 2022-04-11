package token

import "testing"

// TestCreateToken tests the CreateToken function
func TestCreateToken(t *testing.T) {
	nt, err := CreateToken()
	if err != nil {
		t.Errorf("CreateToken() error: %s", err)
	}
	
	if len(nt.Hash) != 64 {
		t.Errorf("CreateToken() error: invalid hash length")
	}

	if len(nt.getPlainText("1")) != 42 {
		t.Errorf("CreateToken() error: invalid plain text length")
	}
}

func TestSplitToken(t *testing.T) {
	nt, _ := CreateToken()

	id, hash, err := SplitToken(nt.getPlainText("1"))

	if err != nil {
		t.Errorf("SplitToken() error: %s", err)
	}

	if id != "1" {
		t.Errorf("SplitToken() error: invalid ID")
	}

	if hash != nt.Hash {
		t.Errorf("SplitToken() error: invalid hash")
	}
}