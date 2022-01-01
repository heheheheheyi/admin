package v1

import (
	"admin/middleware"
	"admin/model"
	"admin/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
)

//登录
func Login(c *gin.Context) {
	var data model.User
	c.ShouldBindJSON(&data)
	var token string
	var code int
	userInfo, code := model.CheckLogin(data.Username, data.Password)
	if code == errmsg.SUCCESS {
		token, code = middleware.SetToken(data.Username)
	}
	c.JSON(http.StatusOK, gin.H{
		"data":    userInfo,
		"status":  code,
		"message": errmsg.GetErrMsg(code),
		"token":   token,
	})
}
