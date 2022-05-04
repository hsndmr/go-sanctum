package middleware

import "testing"


func TestGetBearerTokenFromHeader(t *testing.T) {
	t.Run("it should return the bearer token from the header", func(t *testing.T) {
		token, _ := getBearerTokenFromHeader("Bearer token")
		if token != "token" {
			t.Errorf("getBearerTokenFromHeader() error: invalid token")
		}
	})
	t.Run("it should return an empty string if the header is not valid", func(t *testing.T) {
		token, _ := getBearerTokenFromHeader("Bearer")
		if token != "" {
			t.Errorf("token should be empty")
		}
	})
}