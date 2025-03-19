package models

type Role struct {
	Name    string
	BitMask int64

	BaseEntity[int16]
}
