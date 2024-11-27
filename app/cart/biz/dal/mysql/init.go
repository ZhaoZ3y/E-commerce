package mysql

import (
	"fmt"
	"gomall/app/cart/biz/model"
	"gomall/app/cart/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/opentelemetry/tracing"
	"os"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	dsn := fmt.Sprintf(conf.GetConf().MySQL.DSN, os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"))
	DB, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}
	if err := DB.Use(tracing.NewPlugin(tracing.WithoutMetrics())); err != nil {
		panic(err)
	}
	err = DB.AutoMigrate(&model.Cart{})
	if err != nil {
		return
	}
}
