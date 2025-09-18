package repositories

import (
	"context"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
	"time"
	"ushield_bot/internal/domain"
)

func TestUserPackageSubscriptionsRepository_Create(t *testing.T) {
	dsn := "root:12345678901234567890@(156.251.17.226:6033)/gva"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database: " + err.Error())
	}
	userRepo := NewUserPackageSubscriptionsRepository(db)

	var pkg domain.UserPackageSubscriptions
	pkg.ChatID = 1
	pkg.Status = 0
	pkg.BundleID = 1
	pkg.CreatedAt = time.Now()

	err = userRepo.Create(context.Background(), &pkg)
	if err != nil {
		panic("Failed to create userPackageSubscriptions: " + err.Error())
		return
	}

}
func TestUserPackageSubscriptionsRepository_Update(t *testing.T) {
	dsn := "root:12345678901234567890@(156.251.17.226:6033)/gva"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database: " + err.Error())
	}
	userRepo := NewUserPackageSubscriptionsRepository(db)

	var pkg domain.UserPackageSubscriptions
	pkg.Id = 1
	pkg.ChatID = 123
	pkg.Status = 0
	pkg.BundleID = 456
	pkg.CreatedAt = time.Now()

	err = userRepo.Update(context.Background(), &pkg)
	if err != nil {
		panic("Failed to create userPackageSubscriptions: " + err.Error())
		return
	}

}
