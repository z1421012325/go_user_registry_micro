package account

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"os"
	"time"
)

var (
	DB *gorm.DB
)

const (
	DEFAULT_ACCOUNT_MYSQL_URL = "root:zyms90bdcs@tcp(localhost:3306)/micros_account?charset=utf8&parseTime=True&loc=Local"
)



func init() {

	var mysqladdress string

	if len(os.Getenv("ACCOUNT_MYSQL_URL")) != 0{
		mysqladdress = os.Getenv("ACCOUNT_MYSQL_URL")
	} else {
		mysqladdress = DEFAULT_ACCOUNT_MYSQL_URL
	}

	db,err  := gorm.Open("mysql",mysqladdress)
	if err != nil {
		log.Println("mysql connect is error : ",err)
		os.Exit(1)
	}

	if len(os.Getenv("generating_mode")) != 0 {
		db.LogMode(true)
	}

	err = db.DB().Ping()
	if err != nil {
		log.Println("db mysql ping result err : ",err)
		os.Exit(1)
	}

	db.DB().SetConnMaxLifetime(time.Second*5)
	db.DB().SetMaxIdleConns(5)
	db.DB().SetMaxOpenConns(50)

	DB = db
}


