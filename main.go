package main

import (
	"gin/middlewares"
	"gin/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//模板路径
	r.LoadHTMLGlob("templates/*")
	//全局中间件
	r.Use(middlewares.Auth{}.CheckAuthAll, middlewares.Auth{}.CheckAuth1, middlewares.Auth{}.CheckAuth2)
	routers.AdminInit(r) //后台路由
	routers.IndexInit(r) //前台路由
	routers.ApiInit(r)   //api路由
	r.Run(":8088")
}
