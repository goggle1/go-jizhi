package db

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sirupsen/logrus"
)

var DB *gorm.DB

func InitOrm(database string, maxIdle int, prefix string, logMode bool) {
	var err error
	DB, err = gorm.Open("mysql", database)
	if err != nil {
		panic(err)
	}
	DB.DB().SetMaxIdleConns(100)
	DB.DB().SetMaxOpenConns(16)
	DB.DB().SetConnMaxLifetime(100 * time.Second)
	DB.LogMode(logMode)
	// gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
	// 	return fmt.Sprintf("%s_%s", prefix, defaultTableName)
	// }

	logrus.Info("orm db run")
}
