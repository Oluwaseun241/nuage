package repositories

import (
	"errors"
	"nuage/internal/entities"

	"github.com/google/uuid"
)

var (
	ErrEmailExist   = errors.New("email already exists")
	ErrUserNotFound = errors.New("user not found")
)

type InMemoryUserRepository struct {
	users []*entities.User
}

type UserRepository interface {
	CreateUser(email, password, fullName string) (*entities.User, error)
	UpdateUser()
	GetUserID(id uuid.UUID) (*entities.User, error)
}

func (repo *InMemoryUserRepository) CreateUser(email, password, fullName string) (*entities.User, error) {
	// Check if email already exists
	for _, user := range repo.users {
		if user.Email == email {
			return nil, ErrEmailExist
		}
	}

	// Create new user
	newUser := &entities.User{
		ID:       uuid.New(),
		Email:    email,
		Password: password,
		FullName: fullName,
	}
	repo.users = append(repo.users, newUser)
	return newUser, nil
}

// Get user by id
func (repo *InMemoryUserRepository) GetUserID(id uuid.UUID) (*entities.User, error) {
	for _, user := range repo.users {
		if user.ID == id {
			return user, nil
		}
	}
	return nil, ErrUserNotFound
}
