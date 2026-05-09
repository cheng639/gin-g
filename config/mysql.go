package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

var gormDB *gorm.DB

func GormDB() *gorm.DB {
	return gormDB
}

func InitGormDB() {
	dsn := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		Config().Mysql.User,
		Config().Mysql.Password,
		Config().Mysql.IP,
		Config().Mysql.Port,
		Config().Mysql.Name,
	)

	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic(err)
	}

	sqldb, err := conn.DB()
	if err != nil {
		panic(err)
	}

	err = sqldb.Ping()
	if err != nil {
		panic(err)
	}

	sqldb.SetConnMaxLifetime(time.Duration(Config().Mysql.MaxLifeTime) * time.Second)
	sqldb.SetMaxIdleConns(Config().Mysql.MaxIdleConns)
	sqldb.SetMaxOpenConns(Config().Mysql.MaxOpenConns)

	gormDB = conn
}

type Mysql struct {
	MaxIdleConns      int    `json:"maxidleconns" yaml:"maxidleconns"`
	MaxOpenConns      int    `json:"maxopenconns" yaml:"maxopenconns"`
	Name              string `json:"name" yaml:"name"`
	User              string `json:"user" yaml:"user"`
	Password          string `json:"password" yaml:"password"`
	IP                string `json:"ip" yaml:"ip"`
	Port              int    `json:"port" yaml:"port"`
	Debug             bool   `json:"debug" yaml:"debug"`
	MaxLifeTime       int    `json:"maxlifetime" yaml:"maxlifetime"`
	AutoMigrate       bool   `json:"automigrate" yaml:"automigrate"`
	InterpolateParams bool   `json:"interpolateparams" yaml:"interpolateparams"`
}
