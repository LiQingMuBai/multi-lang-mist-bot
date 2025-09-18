package repositories

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func TestDict_GetDictionaryDetail(t *testing.T) {
	dsn := "root:1234567890@(156.251.17.226:6033)/ushield"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database: " + err.Error())
	}
	userRepo := NewSysDictionariesRepo(db)

	dict, errsg := userRepo.GetDictionaryDetail("server_trx_price")
	if errsg != nil {
		panic(errsg)
		return
	}
	fmt.Printf("dict: %s\n", dict)

}
