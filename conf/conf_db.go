package conf

import "fmt"

type DB struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	DataBase string `yaml:"db"`
	Debug    bool   `yaml:"debug"`  // 打印全部日志
	Source   string `yaml:"source"` // 数据库的源 mysql pgsql
}

func (db DB) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		db.User, db.Password, db.Host, db.Port, db.DataBase)
}

func (db DB) IsEmpty() bool {
	return db.User == "" && db.Password == "" && db.Host == "" && db.Port == 0
}
