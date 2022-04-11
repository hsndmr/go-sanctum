package random

import (
	"crypto/rand"
	"math/big"
)

// GenerateString generates a random string of the specified length
// https://gist.github.com/dopey/c69559607800d2f2f90b1b1ed4e550fb
func GenerateString(size int) (string, error) {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-"
	ret := make([]byte, size)
	for i := 0; i < size; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return "", err
		}
		ret[i] = letters[num.Int64()]
	}

	return string(ret), nil
}