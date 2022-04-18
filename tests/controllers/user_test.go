package controllertest

import (
	"net/http"
	"net/http/httptest"
	"testing"

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
	assert.Equal(t, 201, w.Code)
}