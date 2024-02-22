package repositories

import (
	"errors"
	"nuage/internal/entity"

	"github.com/google/uuid"
)

var (
	ErrEmailExist   = errors.New("email already exists")
	ErrUserNotFound = errors.New("user not found")
)

type InMemoryUserRepository struct {
	users []*entity.User
}

type UserRepository interface {
	CreateUser(email, password, fullName string) (*entity.User, error)
	GetUserID(id uuid.UUID) (*entity.User, error)
}

func (repo *InMemoryUserRepository) CreateUser(email, password, fullName string) (*entity.User, error) {
	// Check if email already exists
	for _, user := range repo.users {
		if user.Email == email {
			return nil, ErrEmailExist
		}
	}

	// Create new user
	newUser := &entity.User{
		ID:       uuid.New(),
		Email:    email,
		Password: password,
		FullName: fullName,
	}
	repo.users = append(repo.users, newUser)
	return newUser, nil
}

// Get user bu id
func (repo *InMemoryUserRepository) GetUserID(id uuid.UUID) (*entity.User, error) {
	for _, user := range repo.users {
		if user.ID == id {
			return user, nil
		}
	}
	return nil, ErrUserNotFound
}
