package core

import (
	"BlogServer/global"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
	"time"
)

func InitDB() *gorm.DB {
	dc := global.Config.DB   // 读库
	dc1 := global.Config.DB1 // 写库

	// TODO:pgSql的支持

	db, err := gorm.Open(mysql.Open(dc.DSN()), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true, // 不生产外键约束
	})
	if err != nil {
		logrus.Fatal("数据库连接失败：%s", err.Error())
	}
	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	logrus.Infof("数据库连接成功")

	if !dc1.IsEmpty() {
		// 读写库不为空 就注册读写分离的配置
		err = db.Use(dbresolver.Register(dbresolver.Config{
			Sources:  []gorm.Dialector{mysql.Open(dc1.DSN())}, // 写数据库
			Replicas: []gorm.Dialector{mysql.Open(dc.DSN())},  // 读数据库
			Policy:   dbresolver.RandomPolicy{},
		}))
		if err != nil {
			logrus.Fatal("数据库读写分离配置错误:%s", err.Error())
		}
	}
	return db
}
