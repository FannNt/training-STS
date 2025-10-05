package database

import (
	"log"

	"book-api/models"

	"gorm.io/gorm"
)

// SeedUsers creates initial users if they don't exist
func SeedUsers(db *gorm.DB) error {
	users := []struct {
		username string
		password string
	}{
		{"admin", "admin123"},
		{"user1", "password1"},
		{"testuser", "test123"},
	}

	for _, userData := range users {
		var existingUser models.User
		result := db.Where("username = ?", userData.username).First(&existingUser)

		if result.Error == gorm.ErrRecordNotFound {
			user := models.User{
				Username: userData.username,
			}
			
			// Hash password before storing
			if err := user.SetPassword(userData.password); err != nil {
				return err
			}

			if err := db.Create(&user).Error; err != nil {
				log.Printf("Failed to seed user %s: %v", userData.username, err)
				continue
			}
			log.Printf("Seeded user: %s", userData.username)
		}
	}

	return nil
}
