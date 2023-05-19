package db

import (
	"fmt"
	"os"
	"time"
	"wkg/conf"

	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	err error
	Orm *gorm.DB
)

func InitDB() {
	dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=%s&parseTime=%s&loc=%s", conf.Cfg.Mysql.User, conf.Cfg.Mysql.Pass, conf.Cfg.Mysql.Host,
		conf.Cfg.Mysql.Port, conf.Cfg.Mysql.Database, conf.Cfg.Mysql.CharSet, conf.Cfg.Mysql.ParseTime, conf.Cfg.Mysql.Loc)

	Orm, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		zap.S().Error(err.Error())
		os.Exit(0)
	}

	sqldb, err := Orm.DB()
	if err != nil {
		zap.S().Error(err.Error())
		panic(err)
	}
	sqldb.SetConnMaxLifetime(time.Hour * 4)
}
