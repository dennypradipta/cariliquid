package seed

import (
	"github.com/google/uuid"
	"log"

	"github.com/jinzhu/gorm"
	"github.com/dennypradipta/cariliquid/models"
)

var users = []models.User{
	models.User{
		ID:			uuid.New().String(),
		Username:	"Denny Pradipta",
		Email:		"denny@gmail.com",
		Password:	"12345678",
	},
	models.User{
		ID:			uuid.New().String(),
		Username:	"Ardhie Putra Prananta",
		Email:		"ardhie@gmail.com",
		Password:	"12345678",
	},
}

// Load basic data
func Load(db *gorm.DB) {
	// Check if Table "users" Exists
	exists := db.Debug().HasTable("users");
	// If it does not exists, migrate
	if !exists {
		err := db.Debug().AutoMigrate(&models.User{}, ).Error
		if err != nil {
			log.Fatalf("cannot migrate table: %v", err)
		}

		for i := range users {
			err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
			if err != nil {
				log.Fatalf("cannot seed users table: %v", err)
			}
		}
	}
}