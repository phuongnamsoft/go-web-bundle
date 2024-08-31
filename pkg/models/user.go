package models

import (
	"fmt"
	"time"

	"github.com/phuongnamsoft/go-web-bundle/config"

	"github.com/phuongnamsoft/go-web-bundle/app"
	"gorm.io/gorm"
)

type User struct {
	ID            uint   `gorm:"primarykey"`
	FirstName     string `json:"first_name" gorm:"first_name"` //nolint:gofmt
	LastName      string `json:"last_name" gorm:"last_name"`   //nolint:gofmt
	Email         string `json:"email" gorm:"email"`
	Password      string `json:"-" gorm:"password"`
	EmailVerified bool   `json:"email_verified" gorm:"email_verified"`
	Currency      string `json:"currency" gorm:"currency"`
	Roles         []Role `gorm:"many2many:user_roles;"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

func AllUsers() []User {
	var users []User
	app.Http.Database.Find(&users)
	return users
}

func GetUserByEmail(email string) (*User, error) {
	var user User
	if err := app.Http.Database.Preload("Metas").Where(&User{Email: email}).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil //nolint:wsl
}

func GetVerifiedUserByEmail(email string) (*User, error) {
	var user User
	if err := app.Http.Database.Preload("Metas").Where(&User{Email: email, EmailVerified: true}).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil //nolint:wsl
}

func GetUserById(id interface{}) (*User, error) {
	var user User
	if err := app.Http.Database.Preload("Files").Where("id = ? ", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *User) Update() error {
	if u.ID != 0 {
		if err := app.Http.Database.Updates(&u).Error; err != nil {
			return err
		}
	} else {
		if err := app.Http.Database.Where(&User{Email: u.Email}).Updates(&u).Error; err != nil {
			return err
		}
	}
	app.Http.Database.First(&u)
	return nil
}

func (u *User) GetUserRoles() []Role {
	u.Roles = []Role{}
	app.Http.Database.Model(&u).Association("Roles").Find(&u.Roles)

	return u.Roles
}

func (u *User) AddRole(role Role) {
	app.Http.Database.Model(&u).Association("Roles").Append(&role)
}

func (u *User) GetUserPermissions() []Permission {
	var permissions []Permission
	roles := u.GetUserRoles()
	for _, role := range roles {
		var rolePermissions []Permission = role.GetRolePermissions()
		permissions = append(permissions, rolePermissions...)
	}
	return permissions
}

func (u *User) Can(permission string) bool {
	return app.Http.Auth.Casbin.Can(fmt.Sprintf("%d", u.ID), permission, config.MatchAll)
}
