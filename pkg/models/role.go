package models

import (
	"time"

	"github.com/phuongnamsoft/go-web-bundle/app"
	"github.com/phuongnamsoft/go-web-bundle/pkg/helpers"
	"gorm.io/gorm"
)

type Role struct {
	ID          uint   `gorm:"primarykey"`
	Name        string `json:"name" gorm:"name"`               //nolint:gofmt
	Slug        string `json:"slug" gorm:"slug"`               //nolint:gofmt
	Description string `json:"description" gorm:"description"` //nolint:gofmt
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

func CheckSlug(slug string) bool {
	var role Role
	app.Http.Database.Where("slug = ?", slug).First(&role)
	if role.ID == 0 {
		return false
	}
	return true
}

func CreateRole(name string, description string) *Role {
	slug := helpers.MakeSlug(name)
	role := Role{
		Name:        name,
		Slug:        slug,
		Description: description,
	}

	app.Http.Database.Save(&role)

	return &role
}
