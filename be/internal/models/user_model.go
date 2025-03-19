package models

import "github.com/google/uuid"

type User struct {
	FirstName string `gorm`
	LastName  string
	Email     string
	Password  string
	Status    int

	BaseEntity[uuid.UUID]
}

const (
	WaitToActive = 0
	Active       = 1
	Block        = 2
)
