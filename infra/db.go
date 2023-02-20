package infra

import (
	"fmt"
	"practice-colly/domain/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func Init(path string) *gorm.DB {
	// もし err がnilではないなら、"読み込み出来ませんでした"が出力されます。
	if err != nil {
		fmt.Printf("読み込み出来ませんでした: %v", err)
	}
	db, err = gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		fmt.Println("db init error: ", err)
	}
	fmt.Println("[INFO] db setup done!")
	return db
}

func GetDB() *gorm.DB {
	return db
}

func autoMigrate() {
	db.AutoMigrate(model.BusstopUrl{})
}
