package config

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
	dsn := os.Getenv("MYSQL_PASSWORD")

	if dsn == "" {
		log.Fatalf("環境変数 'MYSQL_PASSWORD' が設定されていません。")
	}

	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("データベースへの接続に失敗しました: %v", err)
	}
}

func GetDB() *gorm.DB {
	return db
}
