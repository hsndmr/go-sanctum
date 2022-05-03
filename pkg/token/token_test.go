package token

import (
	"testing"

	cryptoservice "github.com/hsndmr/go-sanctum/pkg/crypto"
)

// TestCreateToken tests the CreateToken function
func TestCreateToken(t *testing.T) {
	token := &Token{
		Crypto: &cryptoservice.Crypto{},
	}
	nt, err := token.Create()
	if err != nil {
		t.Errorf("CreateToken() error: %s", err)
	}
	
	if len(nt.Hash) != 64 {
		t.Errorf("CreateToken() error: invalid hash length")
	}

	if len(nt.GetPlainText("1")) != 42 {
		t.Errorf("CreateToken() error: invalid plain text length")
	}
}

func TestSplitToken(t *testing.T) {
	token := &Token{
		Crypto: &cryptoservice.Crypto{},
	}
	nt, _ := token.Create()

	id, hash, err := token.Split(nt.GetPlainText("1"))

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