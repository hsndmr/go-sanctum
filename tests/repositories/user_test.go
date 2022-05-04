package repositorytest

import (
	"testing"

	"github.com/hsndmr/go-sanctum/app"
)

// TestCreateUser creates a user and asserts that the user is created.
func TestCreateUser(t *testing.T) {
	uf := &UserFactory{}
	u, _ := uf.Create()
	if u.Name == "" {
		t.Errorf("failed creating user")
	}
}

// TestFindByEmail finds a user by email and asserts that the user is found.
func TestFindUserByEmail(t *testing.T) {
	uf := &UserFactory{}
	u, _ := uf.Create()
	ur := app.C.Repository.User
	u2, _ := ur.FindByEmail(u.Email)
	if u2.Name != u.Name {
		t.Errorf("failed finding user by email")
	}
}

func TestFindUserByID(t *testing.T) {
	uf := &UserFactory{}
	u, _ := uf.Create()
	ur := app.C.Repository.User
	u2, _ := ur.FindByID(u.ID)
	if u2.ID != u.ID {
		t.Errorf("failed finding user by id")
	}
}