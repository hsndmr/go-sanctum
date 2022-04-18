package controllertest

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestCreateUser(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/v1/user", interface{
		"name": "test",
	})
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}