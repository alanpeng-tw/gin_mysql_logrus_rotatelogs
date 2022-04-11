package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"gin_mysql_logrus_rotatelogs/db"
	"gin_mysql_logrus_rotatelogs/middleware"
)


func init(){
	fmt.Println(" init start...")
	db.InitMySQLDB()
}

func main(){
	fmt.Println("main start")

	r := setupRouter()
	r.Run()

}

func setupRouter() *gin.Engine{

	//gin
	r := gin.Default()

	//設定log
	r.Use(middleware.LoggerToFile())

	r.GET("/ping", func(c *gin.Context){
		c.String(http.StatusOK,"pong")
	})

	return r
}



