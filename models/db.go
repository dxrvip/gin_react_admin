package models

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Db *gorm.DB

func DbInit() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Info, // 日志级别
			IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  false,       // 禁用彩色打印
		})
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		viper.GetString("database.DbUser"),
		viper.GetString("database.DbPassword"),
		viper.GetString("database.DbHost"),
		viper.GetString("database.DbPort"),
		viper.GetString("database.DbName"),
	)
	fmt.Println(dns)
	var err error
	Db, err = gorm.Open(mysql.Open(dns), &gorm.Config{Logger: newLogger})
	if err != nil {
		panic("failed to connect database")
	}
	//
	sqlDb, _ := Db.DB()
	sqlDb.SetMaxIdleConns(10)
	sqlDb.SetMaxOpenConns(100)
	sqlDb.SetConnMaxLifetime(10 * time.Hour)

	// 数据迁移
	_ = Db.AutoMigrate(&User{}, &Category{}, &Article{})
}
