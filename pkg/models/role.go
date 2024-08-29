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

func CheckSlugExists(slug string) bool {
	var role Role
	app.Http.Database.Where("slug = ?", slug).First(&role)
	return role.ID != 0
}

func GenerateSlug(slug string) string {
	if len(slug) == 0 {
		return ""
	}

	slug = helpers.MakeSlug(slug)

	if CheckSlugExists(slug) {
		for i := 1; i < 10; i++ {
			newSlug := slug + "-" + string(i)
			if !CheckSlugExists(newSlug) {
				return newSlug
			}
		}
	}

	return slug
}

func CreateRole(name string, description string) *Role {
	slug := GenerateSlug(name)
	role := &Role{
		Name:        name,
		Slug:        slug,
		Description: description,
	}

	app.Http.Database.Save(&role)

	return role
}

func seed() {
	roles := []Role{
		{Name: "Admin", Description: "Admin Role"},
		{Name: "User", Description: "User Role"},
	}

	for _, role := range roles {
		CreateRole(role.Name, role.Description)
	}
}

