package defaults_tables

import (
	"github.com/Arifkobel/shelf/services/database/schemas"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(schemas.User{}, schemas.File{}, schemas.Provider{})
}

func SetProviders(db *gorm.DB) {
	db.Migrator().DropTable("providers")
	db.AutoMigrate(&schemas.Provider{})
	db.Create(&schemas.Provider{
		Name: "email",
	})
}
