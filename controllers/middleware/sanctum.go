package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/hsndmr/go-sanctum/app"
)

func Sanctum() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := getBearerTokenFromHeader(c.GetHeader("Authorization"))

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized,gin.H{
				"message": "unauthorized",
			})
			return
		}

		user, err := app.C.Repository.PersonalAccessToken.Verify(token)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized,gin.H{
				"message": "unauthorized",
			})
			return
		}

		c.Set("user", user)

		c.Next()
	}
}

// getBearerTokenFromHeader returns the token from the header
func getBearerTokenFromHeader(header string) (string, error) {
	if header == "" {
		return "", errors.New("not found authorization header")
	}

	token := strings.Split(header, " ")
	if len(token) != 2 {
		return "", errors.New("incorrectly formatted authorization header")
	}

	if token[0] != "Bearer" {
		return "", errors.New("incorrectly formatted authorization header")
	}

	return token[1], nil
}