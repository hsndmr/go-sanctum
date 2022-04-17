package repositories

import (
	"testing"
)

// TestCreateUser creates a user and asserts that the user is created.
func TestCreateUser(t *testing.T) {
	uf := &UserFactory{}
	u, _ := uf.Create()
	if u.Name == "" {
		t.Errorf("failed creating user")
	}
}