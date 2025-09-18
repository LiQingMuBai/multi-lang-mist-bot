package repositories

import (
	"context"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
	"time"
	"ushield_bot/internal/domain"
)

func TestUserEnergyOrdersRepo_Create(t *testing.T) {
	dsn := "root:seven7ushield@(156.251.17.226:6033)/ushield_xyz"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database: " + err.Error())
	}
	userRepo := NewUserEnergyOrdersRepo(db)

	var pkg domain.UserEnergyOrders
	pkg.ChatId = "1111111"
	pkg.Amount = 65000
	pkg.OrderNo = "ADDDDJDDDDD"
	pkg.FromAddress = "TXXXXX"
	pkg.ToAddress = "TXXXXX123213213"
	pkg.CreatedAt = time.Now()

	err = userRepo.Create(context.Background(), &pkg)
	if err != nil {
		panic("Failed to create userPackageSubscriptions: " + err.Error())
		return
	}

}
func TestUserEnergyOrdersRepo_Count(t *testing.T) {
	dsn := "root:seven7ushield@(156.251.17.226:6033)/ushield_xyz"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database: " + err.Error())
	}
	energyRepo := NewUserEnergyOrdersRepo(db)

	count, err := energyRepo.Count(context.Background(), 1111111)

	if err != nil {
		panic("Failed to count userEnergyOrders: " + err.Error())
	}
	fmt.Println(count)
}
