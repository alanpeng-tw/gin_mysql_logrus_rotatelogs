package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"time"

	"gin_mysql_logrus_rotatelogs/config"
)

func LoggerToFile() gin.HandlerFunc{

	logFilePath := config.LOG_FILE_PATH
	logFileName := config.LOG_FILE_NAME

	fileName := path.Join(logFilePath,logFileName )

	//write log to log file
	src , err := os.OpenFile(logFileName, os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0644)
	if err != nil{
		fmt.Println("出事了阿北!!!!")
		fmt.Println("err", err)
	}

	//logrus實例化
	logger := logrus.New()

	//設置輸出
	logger.Out = src

	//設置logrus Level
	logger.SetLevel(logrus.DebugLevel)

	logWriter, err := rotatelogs.New(

		//分割後的文件名稱
		fileName + ".%Y-%m-%d",

		//生成 soft link and pointer to new log file
		rotatelogs.WithLinkName(fileName),

		//最大保存時間(7天)
		rotatelogs.WithMaxAge(7*24*time.Hour),

		//設置log 切割時間間隔(1天)
		rotatelogs.WithRotationTime(24*time.Hour),
	)

	writeMap := lfshook.WriterMap{
		logrus.InfoLevel: logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel: logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}

	lfHook := lfshook.NewHook( writeMap, &logrus.JSONFormatter{
		TimestampFormat : "2006-01-02 15:04:05",
	})

	logger.AddHook(lfHook)

	return func(c *gin.Context) {

		//開始時間
		startTime := time.Now()
		//結束時間
		endTime := time.Now()

		//處理請求
		c.Next()

		//執行時間
		latencyTime := endTime.Sub(startTime)

		//請求方式
		reqMethod := c.Request.Method
		//請求路由
		reqUri := c.Request.RequestURI

		//status code
		statusCode := c.Writer.Status()

		//請求IP
		clientIP := c.ClientIP()

		//日誌格式
		logger.WithField("logger",logrus.Fields{
			"status_code": statusCode,
			"latency_time": latencyTime,
			"client_ip": clientIP,
			"request_method": reqMethod,
			"request_uri": reqUri,
		}).Info()
	}
}

func LoggerToMongo() gin.HandlerFunc{
	return func( c * gin.Context) {}
}

func LoggerToMysql() gin.HandlerFunc{
	return func( c * gin.Context) {}
}

func LoggerToMQ() gin.HandlerFunc{
	return func( c * gin.Context) {}
}
