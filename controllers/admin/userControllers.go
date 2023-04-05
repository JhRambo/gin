package admin

import (
	"fmt"
	"gin/models"
	"gin/tools"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	Name  string `json:"name" form:"name"`
	Age   int    `json:"age" form:"age"`
	Phone string `json:"phone" form:"phone"`
}

// 数据绑定
func (con UserController) GetShouldBind(ctx *gin.Context) {
	var user UserController
	err := ctx.ShouldBind(&user)
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": err.Error(),
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

// 查询指定id的数据，动态路由
func (con UserController) GetById(ctx *gin.Context) {
	user := models.User{}
	id, _ := strconv.Atoi(ctx.Param("id"))
	tools.DB.First(&user, id)
	username, _ := ctx.Get("username")
	v, _ := username.(string)
	fmt.Println(v + "：Admin-User-GetById") //获取中间件设置的值
	ctx.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

// 查询指定id的数据
func (con UserController) GetOne(ctx *gin.Context) {
	user := models.User{}
	id, _ := strconv.Atoi(ctx.Query("id"))
	tools.DB.First(&user, id)
	ctx.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

// 查询所有数据
func (con UserController) List(ctx *gin.Context) {
	user := []models.User{}
	tools.DB.Find(&user)
	ctx.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

// 新增
func (con UserController) Add(ctx *gin.Context) {
	age, _ := strconv.Atoi(ctx.PostForm("age"))
	user := models.User{
		Name:     ctx.PostForm("name"),
		Phone:    ctx.PostForm("phone"),
		Age:      age,
		Password: ctx.PostForm("password"),
	}
	tools.DB.Create(&user) // 通过数据的指针来创建
	defer func() {
		if err := recover(); err != nil {
			panic(err)
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "ok",
			})
		}
	}()
}

// 删除
func (con UserController) Del(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.PostForm("id"))
	user := models.User{}
	tools.DB.Where("id = ?", id).Delete(&user)
	defer func() {
		if err := recover(); err != nil {
			panic(err)
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "ok",
			})
		}
	}()
}

// 更新
func (con UserController) Upd(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.PostForm("id"))
	age, _ := strconv.Atoi(ctx.PostForm("age"))
	user := models.User{
		Name:     ctx.PostForm("name"),
		Phone:    ctx.PostForm("phone"),
		Age:      age,
		Password: ctx.PostForm("password"),
	}
	tools.DB.Model(&user).Where("id = ?", id).Updates(user)
	defer func() {
		if err := recover(); err != nil {
			panic(err)
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "ok",
			})
		}
	}()
}
