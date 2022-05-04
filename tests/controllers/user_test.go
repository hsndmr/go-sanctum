package controllertest

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hsndmr/go-sanctum/app"
	"github.com/hsndmr/go-sanctum/repositories"
	repositorytest "github.com/hsndmr/go-sanctum/tests/repositories"
	"github.com/stretchr/testify/assert"
)


func TestCreateUser(t *testing.T) {
	w := httptest.NewRecorder()

	body := struct {
    Name    string
    Email   string
    Password string
	}{
		Name:    "name",
		Email:   "email@email.com",
		Password: "secret",
	}
	
	req, _ := http.NewRequest("POST", "/api/v1/user", jsonReaderFactory(body))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestLoginUser(t *testing.T) {
	uf := &repositorytest.UserFactory{}
	u, _ := uf.Create()

	t.Run("it should login user", func(t *testing.T) {
		w := httptest.NewRecorder()
		body := struct {
			Email    string
			Password string
		}{
			Email:    u.Email,
			Password: "secret",
		}

		req, _ := http.NewRequest("POST", "/api/v1/auth/login", jsonReaderFactory(body))
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("it should not login user with wrong password", func(t *testing.T) {
		w := httptest.NewRecorder()
		body := struct {
			Email    string
			Password string
		}{
			Email:    u.Email,
			Password: "secret1",
		}

		req, _ := http.NewRequest("POST", "/api/v1/auth/login", jsonReaderFactory(body))
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

}

func TestGetUser(t *testing.T) {
	uf := &repositorytest.UserFactory{}
	u, _ := uf.Create()

	_, token,  _ := app.C.Repository.PersonalAccessToken.Create(&repositories.CreatePersonalAccessTokenDto{
		Name: "Personal Access Token",
		User: u,
	})
	
	t.Run("it should get user", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/v1/user", nil)
		req.Header.Set("Authorization", "Bearer "+token)
		router.ServeHTTP(w, req)
		
		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("it should throw authenticate error", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/v1/user", nil)
		req.Header.Set("Authorization", "Bearer token")
		router.ServeHTTP(w, req)
		
		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})
}