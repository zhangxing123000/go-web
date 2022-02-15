package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web/common"
	"web/dto"
	"web/model"
)

//Register  注册用户
func Register(context *gin.Context) {
	db := common.GetDB()
	name := context.PostForm("name")
	password := context.PostForm("password")
	tel := context.PostForm("tel")

	if len(tel) != 11 {
		context.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": http.StatusUnprocessableEntity,
			"msg":  "手机号必须是11位",
		})
		return
	}
	user := model.User{
		Name:     name,
		Password: password,
		Tel:      tel,
	}
	db.Create(&user)

	context.JSON(http.StatusOK, gin.H{
		"code":    200,
		"msg":     "注册成功",
		"content": dto.ToUserDto(user),
	})

}
