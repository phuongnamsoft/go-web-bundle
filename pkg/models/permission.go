package models

import (
	"time"
)

type Permission struct {
	ID         uint    `gorm:"primarykey"`
	Name       string  `json:"name" gorm:"name"`
	Slug       string  `json:"slug" gorm:"slug"`
	HttpPath   *string `json:"http_path" gorm:"http_path"`
	HttpMethod *string `json:"http_method" gorm:"http_method"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (Permission) TableName() string {
	return "permissions"
}
