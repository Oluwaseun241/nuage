package repositories

import (
	"errors"
	"nuage/internal/entities"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
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
	UpdateUser(id uuid.UUID, fullname string) (*entities.User, error)
	GetUserID(id uuid.UUID) (*entities.User, error)
}

func hashPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "hashing error"
	}
	return string(hashedPassword)
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
		Password: hashPassword(password),
		FullName: fullName,
	}
	repo.users = append(repo.users, newUser)
	return newUser, nil
}

func (repo *InMemoryUserRepository) UpdateUser(id uuid.UUID, fullname string) (*entities.User, error) {
	for _, user := range repo.users {
		if user.ID == id {
			user.FullName = fullname
			return user, nil
		}
	}
	return nil, ErrUserNotFound
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
