package conf

import (
	"server/model"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func InitMysql() (*gorm.DB, error) {
	logMode := logger.Info
	if !viper.GetBool("mode.debug") {
		logMode = logger.Error
	}

	dsn := viper.GetString("mysql_db.dsn")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// 限定自动生成表名的策略
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "sys_",
			SingularTable: true, // 禁用复数
		},
		Logger: logger.Default.LogMode(logMode),
	})

	if err != nil {
		return nil, err
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(viper.GetInt("mysql_db.MaxIdleConns"))
	sqlDB.SetMaxOpenConns(viper.GetInt("mysql_db.MaxOpenConns"))
	sqlDB.SetConnMaxLifetime(time.Hour)

	// if err := sqlDB.Ping(); err != nil {

	// 	global.Logger.Error("failed to ping database", err)

	// }
	db.AutoMigrate(&model.User{})

	return db, nil
}

// type manager struct {
// 	db *gorm.DB
// }

// var Mgr Manager

// // 这里的 init 函数只会被执行一次

// func init() {

// 	dsn := "root:admin1234@tcp(127.0.0.1:3306)/go_vue_learn?charset=utf8mb4&parseTime=True&loc=Local"

// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

// 	if err != nil {

// 		log.Fatal("failed to connect database", err)

// 	}

// 	sqlDB, _ := db.DB()

// 	if err := sqlDB.Ping(); err != nil {

// 		log.Fatalf("Database ping failed: %v", err)

// 	}

// 	Mgr = &manager{db: db}

// 	if err := db.AutoMigrate(&model.User{}); err != nil {

// 		log.Printf("AutoMigrate for User failed: %v", err)

// 	}

// }
