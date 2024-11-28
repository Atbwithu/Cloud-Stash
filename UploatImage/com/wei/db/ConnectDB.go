package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//var DB *gorm.DB
//var err error

func ConnectDB() (*gorm.DB, error) {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:Wzl673008!@tcp(47.98.127.137)/cloud_db?charset=utf8mb4&parseTime=True&loc=Local"

	// 使用gorm.Open建立数据库连接
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %v", err)
	}

	return db, nil
}
