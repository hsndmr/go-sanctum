package token

import (
	"errors"
	"strings"

	cryptoservice "github.com/hsndmr/go-sanctum/pkg/crypto"
	"github.com/hsndmr/go-sanctum/pkg/random"
)

type Token struct {
	Crypto *cryptoservice.Crypto
}

// Create creates a new token
func (t *Token)  Create() (*NewToken, error) {
	plainText, err := random.GenerateString(40)
	if err != nil {
		return nil, err
	}

	return &NewToken{
		plain: plainText,
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
	plain string
	Hash string
}

// GetPlainText returns the plain text
// ModelID.PlainText
func (nt *NewToken) getPlainText(ID string) string {
	return ID+"|"+nt.plain
}


 