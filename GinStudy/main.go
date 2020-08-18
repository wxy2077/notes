package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"path"
	"time"
)

type UserInfo struct {
	Name string `json:"name" binding:"required"`
	Age  int    `json:"age" binding:"required"`
}

// 自定义一个中间件
func TestMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		fmt.Println("请求前")
		context.Next()
		fmt.Println("请求后")
	}
}

// 处理跨域请求中间价,支持options访问
// 参考 https://stackoverflow.com/questions/29418478/go-gin-framework-cors
// https://juejin.im/post/6844903589870043149
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 允许所有的域名访问
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		// 通过所有OPTIONS方法
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

// 日志文件
var (
	logFilePath = "./"
	logFileName = "system.log"
)

func LoggerMiddleware() gin.HandlerFunc {
	// 日志文件
	fileName := path.Join(logFilePath, logFileName)
	// 写入文件
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err)
	}
	// 实例化
	logger := logrus.New()
	//设置日志级别
	logger.SetLevel(logrus.DebugLevel)
	//设置输出
	logger.Out = src

	// 设置 rotatelogs
	logWriter, err := rotatelogs.New(
		// 分割后的文件名称
		fileName+".%Y%m%d.log",
		// 生成软链，指向最新日志文件
		rotatelogs.WithLinkName(fileName),
		// 设置最大保存时间(7天)
		rotatelogs.WithMaxAge(7*24*time.Hour),
		// 设置日志切割时间间隔(1天)
		rotatelogs.WithRotationTime(24*time.Hour),
	)

	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}

	logger.AddHook(lfshook.NewHook(writeMap, &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	}))

	return func(c *gin.Context) {
		//开始时间
		startTime := time.Now()
		//处理请求
		c.Next()
		//结束时间
		endTime := time.Now()
		// 执行时间
		latencyTime := endTime.Sub(startTime)
		//请求方式
		reqMethod := c.Request.Method
		//请求路由
		reqUrl := c.Request.RequestURI
		//状态码
		statusCode := c.Writer.Status()
		//请求ip
		clientIP := c.ClientIP()
		//请求UA
		clientUA := c.Request.UserAgent()

		// 日志格式
		logger.WithFields(logrus.Fields{
			"status_code":  statusCode,
			"latency_time": latencyTime,
			"client_ip":    clientIP,
			"req_method":   reqMethod,
			"req_uri":      reqUrl,
			"UA":           clientUA,
		}).Info()

	}
}

func main() {
	router := gin.Default()

	router.Use(TestMiddleware())
	router.Use(CORSMiddleware())
	router.Use(LoggerMiddleware())

	/*
		GET 请求测试
	*/
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	/*
		get请求示例
	*/
	router.GET("/info", func(c *gin.Context) {
		user := c.DefaultQuery("user", "root")
		password := c.Query("password")
		logrus.Infof("记录日志")
		//logrus.WithFields(logrus.Fields{
		//	"animal": "walrus",
		//}).Info("A walrus appears")

		c.JSON(200, gin.H{
			"code":    200,
			"message": "success",
			"data": gin.H{
				"user":     user,
				"password": password,
			},
		})
	})

	/*
		POST form表单请求示例
	*/
	router.POST("/auth", func(c *gin.Context) {
		// 获取form表单的数据
		user := c.DefaultPostForm("user", "post_user")
		password := c.PostForm("password")

		c.JSON(200, gin.H{
			"code":    200,
			"message": "success",
			"data": gin.H{
				"user":     user,
				"password": password,
			},
		})
	})

	/*
		POST 绑定结构体参数 获取json数据
	*/
	router.POST("/user", func(c *gin.Context) {
		// 声明一个user结构体
		var u UserInfo

		// 绑定结构体Json格式
		err := c.ShouldBindJSON(&u)
		if err != nil {
			log.Println("错误原因:", err.Error())

			c.JSON(200, gin.H{
				"code":    400,
				"message": err.Error(),
			})
		} else {
			c.JSON(200, gin.H{
				"code":    200,
				"message": "success",
				"data":    u,
			})
		}

	})

	_ = router.Run("127.0.0.1:8070") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
