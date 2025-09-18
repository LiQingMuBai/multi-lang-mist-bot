package repositories

import (
	"context"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func TestUserUsdtSubscriptionsRepository_ListAll2(t *testing.T) {
	dsn := "root:12345678901234567890@(156.251.17.226:6033)/gva"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database: " + err.Error())
	}
	userRepo := NewUserUsdtSubscriptionsRepository(db)
	all, err := userRepo.ListAll(context.Background())
	if err != nil {
		panic("Failed to list user: " + err.Error())
		return
	}
	for _, user := range all {
		fmt.Printf("%+v\n", user)
	}
}
