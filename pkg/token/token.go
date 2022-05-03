package token

import (
	"errors"
	"strings"

	cryptoservice "github.com/hsndmr/go-sanctum/pkg/crypto"
	"github.com/hsndmr/go-sanctum/pkg/random"
)

type TokenI interface {
	Create() (*NewToken, error)
	Split(token string) (string, string, error)
}

type Token struct {
	Crypto cryptoservice.CryptoI
}

// Create creates a new token
func (t *Token)  Create() (*NewToken, error) {
	plainText, err := random.GenerateString(40)
	if err != nil {
		return nil, err
	}

	return &NewToken{
		PlainText: plainText,
		Hash: t.Crypto.SHA256(plainText),
	}, nil
}

// Split returns the hash and ID
func (t *Token) Split(token string) (string, string, error) {
	parts := strings.Split(token, "|")
	if(len(parts) != 2) {
		return "", "", errors.New("invalid token")
	}
	return parts[0], t.Crypto.SHA256(parts[1]), nil
}

// New Token
type NewToken struct {
	PlainText string
	Hash string
}

// GetPlainText returns the plain text
// ModelID.PlainText
func (nt *NewToken) GetPlainText(ID string) string {
	return ID+"|"+nt.PlainText
}


 