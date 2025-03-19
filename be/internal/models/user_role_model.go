package models

import "github.com/google/uuid"

type UserRole struct {
	UId uuid.UUID
	RId int16

	BaseEntity[int32]
}
