package models

import (
	"time"
)

type UserRole struct {
	ID        uint `gorm:"primarykey"`
	UserID    uint `json:"user_id" gorm:"user_id"` //nolint:gofmt
	RoleID    uint `json:"role_id" gorm:"role_id"` //nolint:gofmt
	User      User `gorm:"foreignKey:UserID"`
	Role      Role `gorm:"foreignKey:RoleID"`
	CreatedAt time.Time
}
