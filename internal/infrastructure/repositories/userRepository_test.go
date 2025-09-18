package repositories

import (
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"testing"
)

func TestUserRepository_GetByUserID(t *testing.T) {
	dsn := "root:1234567890@(156.251.17.226:6033)/gva"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database: " + err.Error())
	}
	userRepo := NewUserRepository(db)
	user, err := userRepo.GetByUserID(7347235462)
	if err != nil {
		panic("Failed to get the user")
		return
	}

	log.Printf("userRepo.GetByUserID(12323): %+v\n", user)

}
