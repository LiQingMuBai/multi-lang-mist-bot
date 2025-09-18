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

func TestUserUSDTDepositsRepo_Create(t *testing.T) {
	// 获取当前本地时间
	localTime := time.Now()
	fmt.Println("本地时间:", localTime)

	// 转换为UTC时间
	utcTime := localTime.UTC()
	fmt.Println("UTC时间:", utcTime)

	// 显示时区信息
	fmt.Println("时区:", localTime.Location())

	// 显示时间戳（Unix时间戳与时区无关）
	fmt.Println("Unix时间戳:", localTime.Unix())
}

func TestUserUSDTDdepositsRepo_ListAll(t *testing.T) {
	dsn := "root:1234567890@(156.251.17.226:6033)/ushield?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database: " + err.Error())
	}
	userUSDTDepositsRepo := NewUserUSDTDepositsRepository(db)
	orderNO := "9999999"
	var usdtDeposit domain.UserUSDTDeposits
	usdtDeposit.OrderNO = orderNO
	//usdtDeposit.UserID = callbackQuery.Message.Chat.ID
	usdtDeposit.Status = 0
	usdtDeposit.Placeholder = "111111111111"

	usdtDeposit.Amount = "1232132132132"
	usdtDeposit.CreatedAt = time.Now()

	userUSDTDepositsRepo.Create(context.Background(), &usdtDeposit)

}
