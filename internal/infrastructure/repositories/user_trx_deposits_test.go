package repositories

import (
	"context"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strings"
	"testing"
	"ushield_bot/internal/request"
)

//import (
//	"context"
//	_ "github.com/go-sql-driver/mysql"
//	"gorm.io/driver/mysql"
//	"testing"
//	"time"
//	"ushield_bot/internal/domain"
//	tools "ushield_bot/internal/infrastructure/toos"
//
//	"gorm.io/gorm"
//)

func TestUserTRXDdepositsRepo_ListAll(t *testing.T) {
	dsn := "root:1234567890@(156.251.17.226:6033)/ushield"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database: " + err.Error())
	}
	userRepo := NewUserTRXDepositsRepository(db)

	list, errsg := userRepo.ListAll(context.Background(), 6620733754, 1)
	if errsg != nil {
		panic(errsg)
		return
	}
	var builder strings.Builder
	//- [6.29] +3000 TRX（订单 #TOPUP-92308）
	for _, word := range list {
		builder.WriteString("[")
		builder.WriteString(word.CreatedDate)
		builder.WriteString("]")
		builder.WriteString("+")
		builder.WriteString(word.Amount)
		builder.WriteString(" TRX ")
		builder.WriteString(" （订单 #TOPUP- ")
		builder.WriteString(word.OrderNO)
		builder.WriteString("）")

		builder.WriteString("\n") // 添加分隔符
	}

	// 去除最后一个空格
	result := strings.TrimSpace(builder.String())

	fmt.Printf("%+v\n", result)
	//for _, u := range list {
	//	fmt.Printf("%+v\n", u)
	//}
}

func TestUserTRXDdepositsRepo_GetUserTrxDepositsInfoList(t *testing.T) {
	dsn := "root:1234567890@(156.251.17.226:6033)/ushield"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database: " + err.Error())
	}
	userRepo := NewUserTRXDepositsRepository(db)

	var info request.UserTrxDepositsSearch
	info.PageInfo.Page = 1
	info.PageInfo.PageSize = 1
	list, total, _ := userRepo.GetUserTrxDepositsInfoList(context.Background(), info, 6620733754)

	fmt.Printf("total:%d\n", total)
	for i := range list {
		fmt.Printf("%+v\n", list[i])
	}

}
