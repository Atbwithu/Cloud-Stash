package main

import (
	db2 "UploatImage/com/wei/db"
	"UploatImage/com/wei/pojo"
	_ "database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	r := gin.Default()

	db, err := db2.ConnectDB()
	if err != nil {
		fmt.Println("数据库查询失败:", err)
		return
	}
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/index", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "index",
		})
	})
	// 登录账号
	r.POST("/login", func(context *gin.Context) {
		var user pojo.User
		user.Username = context.Query("username")
		user.Password = context.Query("password")
		tx := db.Find(&user)
		print(" 这个值为", tx)
		context.JSON(200, gin.H{
			"message": tx,
		})
	})

	r.Run()

}
