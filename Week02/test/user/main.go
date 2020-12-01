package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"user/api"
	"user/dao"
)

func main() {

	// 1 初始化数据库
	err := dao.InitMysql("127.0.0.1", "3306", "root", "root", "user")
	if err != nil{
		log.Fatal(err)
	}

	// 启动gin web服务
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// 用户注册
	r.POST("/register",api.RegisterHandler)

	// 用户登录
	//r.POST("/login",api.LoginHandler)


	r.Run() // listen and serve on 0.0.0.0:8080
}
