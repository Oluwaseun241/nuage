package repositories

import (
	"nuage/internal/entity"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

// Creates a new user with unique email
func TestCreateUser(t *testing.T) {
	repo := &InMemoryUserRepository{
		users: []*entity.User{
			{
				ID:       uuid.New(),
				Email:    "existing@example.com",
				Password: "password",
				FullName: "Existing User",
			},
		},
	}

	email := "new@example.com"
	password := "password"
	fullName := "New User"

	user, err := repo.CreateUser(email, password, fullName)

	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, email, user.Email)
	assert.Equal(t, password, user.Password)
	assert.Equal(t, fullName, user.FullName)
}
