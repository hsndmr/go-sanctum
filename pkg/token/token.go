package token

import (
	"errors"
	"strings"

	"github.com/hsndmr/go-sanctum/pkg/crypto"
	"github.com/hsndmr/go-sanctum/pkg/random"
)

// CreateToken creates a new token
func  CreateToken() (*NewToken, error) {
	plainText, err := random.GenerateString(40)
	if err != nil {
		return nil, err
	}

	return &NewToken{
		plain: plainText,
		Hash: crypto.SHA256(plainText),
	}, nil
}

// SplitToken returns the hash and ID
func SplitToken(token string) (string, string, error) {
	parts := strings.Split(token, "|")
	if(len(parts) != 2) {
		return "", "", errors.New("invalid token")
	}
	return parts[0], crypto.SHA256(parts[1]), nil
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


 