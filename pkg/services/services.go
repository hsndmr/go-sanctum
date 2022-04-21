package services

import (
	cryptoservice "github.com/hsndmr/go-sanctum/pkg/crypto"
	"github.com/hsndmr/go-sanctum/pkg/token"
)

type Service struct {
	Crypto *cryptoservice.Crypto
	Token token.TokenI
}