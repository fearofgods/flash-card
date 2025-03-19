package models

import (
	"time"
)

type BaseEntity[T any] struct {
	Id T `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`

	CreateBy   *T
	CreateAt   time.Time
	ModifiedBy *T
	ModifiedAt time.Time
}
