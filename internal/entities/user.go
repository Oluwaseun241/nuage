package entities

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID
	Email    string
	FullName string
	Password string
}

