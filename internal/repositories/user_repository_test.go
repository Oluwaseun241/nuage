package repositories

import (
	"nuage/internal/entities"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

// Creates a new user with unique email
func TestCreateUser(t *testing.T) {
	repo := &InMemoryUserRepository{
		users: []*entities.User{
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

func TestUpdateUser(t *testing.T) {
	repo := &InMemoryUserRepository{}
	user := &entities.User{
		ID:       uuid.New(),
		Email:    "test@example.com",
		Password: "password",
		FullName: "Test User",
	}

	repo.users = append(repo.users, user)

	newFullName := "Demo User"
	updateUser, err := repo.UpdateUser(user.ID, newFullName)
	assert.NoError(t, err)
	assert.Equal(t, newFullName, updateUser.FullName)
}

func TestGetUserID(t *testing.T) {
	repo := &InMemoryUserRepository{}
	user := &entities.User{
		ID:       uuid.New(),
		Email:    "test@example.com",
		Password: "password",
		FullName: "Test User",
	}
	repo.users = append(repo.users, user)

	retrievedUser, err := repo.GetUserID(user.ID)
	assert.NoError(t, err)
	assert.Equal(t, user, retrievedUser)
}
