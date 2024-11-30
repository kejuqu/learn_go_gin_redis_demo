package config

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"localhost/backend/global"
)

func initDB() {
	dsn := AppConfig.Database.Dsn
	// 不会校验账号密码是否正确
	// 注意！！！这里不要使用:=，我们是给全局变量赋值，然后在main函数中使用全局变量db

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	fmt.Print("err:::: ", err)

	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	sqlDb, err := db.DB()

	// 设置数据库中最的的空闲链接数量
	sqlDb.SetMaxIdleConns(AppConfig.Database.MaxIdleConns)
	sqlDb.SetMaxOpenConns(AppConfig.Database.MaxOpenCons)
	sqlDb.SetConnMaxLifetime((time.Hour))

	if err != nil {
		log.Fatalf("Failed to set database connection pool: %v", err)
	}

	global.Db = db

}
