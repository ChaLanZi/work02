package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"time"
)

var DB *gorm.DB

func Database() error{
	config := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.addr"),
		viper.GetString("db.name"),
	)

	db, err := gorm.Open(viper.GetString("db.style"), config)
	if err != nil {
		return errors.Wrap(err,"引用第三方库，第一次，warp")
	}

	db.LogMode(viper.GetBool("log"))
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(50)
	db.DB().SetMaxOpenConns(100)
	db.DB().SetConnMaxLifetime(time.Second * 30)

	DB=db

	DB.AutoMigrate(&Goods{})

	return nil
}
