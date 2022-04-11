package db

import (
	"database/sql"
	"log"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"gin_mysql_logrus_rotatelogs/config"
)

func InitMySQLDB() (db *sql.DB) {

	//db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp(34.80.251.57:3306)/"+dbName)
	db, err := sql.Open(config.MYSQL_DRIVER, config.MYSQL_USER+":"+config.MYSQL_PASSWORD+"@tcp("+config.MYSQL_URL+":3306)/"+config.MYSQL_DATABASE_NAME)

	//connection pool
	db.SetMaxOpenConns(10) //用於設定最大開啟的連線數，預設值為0表示不限制
	db.SetMaxIdleConns(10) //設定閒置的連線數則當開啟的一個連線使用完成後可以放在池裡等候下一次使用


	if err != nil {
		fmt.Println(" DB 連線失敗 !!")
		log.Println(err, " DB 連線失敗 !!")

		panic(err.Error())
	}else{
		fmt.Println("DB 連線成功 !!")
	}

	err = db.Ping()
	if err != nil{
		log.Println("Fail to Ping DB Err :[%v]", err.Error())
		return
	}

	log.Println("InitMySQLDB finished")

	return
}
