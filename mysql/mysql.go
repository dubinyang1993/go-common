package mysql

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

var DB *gorm.DB

func Init() {
	addr := viper.GetString("mysql.addr")
	username := viper.GetString("mysql.username")
	password := viper.GetString("mysql.password")
	datebase := viper.GetString("mysql.datebase")
	maxidle := viper.GetInt("mysql.maxidle")
	maxopen := viper.GetInt("mysql.maxoper")
	maxlifetime := viper.GetDuration("mysql.maxlifetime")

	dsn := username + ":" + password + "@tcp(" + addr + ")/" + datebase + "?charset=utf8&loc=Local"
	var err error
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		panic("MySQL conn failed:" + err.Error())
	}

	DB.DB().SetMaxIdleConns(maxidle)
	DB.DB().SetMaxOpenConns(maxopen)
	DB.DB().SetConnMaxLifetime(maxlifetime * time.Second)

	fmt.Println("MySQL conn successful.")
}
