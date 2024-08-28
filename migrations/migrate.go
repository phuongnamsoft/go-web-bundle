package migrations

import (
	"log"

	"github.com/phuongnamsoft/go-web-bundle/app"
	"github.com/phuongnamsoft/go-web-bundle/pkg/models"
)

func Migrate() {
	log.Println("Initiating migration...")
	err := app.Http.Database.DB.Migrator().AutoMigrate(
		&models.Role{},
		&models.RoleAndPermission{},
		&models.User{},
		&models.UserMeta{},
	)
	if err != nil {
		panic(err)
	}
	log.Println("Migration Completed...")
}
