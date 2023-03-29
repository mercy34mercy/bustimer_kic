package infra

import (
	"database/sql"
	"fmt"
	_ "modernc.org/sqlite"
)

var (
	db  *sql.DB
	err error
)

func Init(path string) *sql.DB {
	// もし err がnilではないなら、"読み込み出来ませんでした"が出力されます。
	if err != nil {
		fmt.Printf("読み込み出来ませんでした: %v", err)
	}
	db, err = sql.Open("sqlite", path+"?mode=ro")
	if err != nil {
		fmt.Println("db init error: ", err)
	}
	fmt.Println("[INFO] db setup done!")
	return db
}

func GetDB() *sql.DB {
	return db
}
