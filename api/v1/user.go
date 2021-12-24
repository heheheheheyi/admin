package v1

import (
	"admin/model"
	"admin/utils/errmsg"
	"admin/utils/validator"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var code int

// 添加用户

func AddUser(c *gin.Context) {
	var data model.User
	_ = c.ShouldBindJSON(&data)
	var msg string
	msg, code = validator.Validate(&data)
	if code != errmsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": msg,
		})
		return
	}
	code = model.AddUser(&data)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

func GetUserId(c *gin.Context) {
	username := c.Query("username")
	id, code := model.GetUserId(username)
	if id == 0 {
		code = errmsg.ERROR
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"id":      id,
		"message": errmsg.GetErrMsg(code),
	})
}

// 查询单个用户
func GetOneUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, code := model.GetOneUser(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// 查询用户列表

func GetUser(c *gin.Context) {
	username := c.Query("username")
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))

	data, total, code := model.GetUser(username, pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
		"total":   total,
	})
}

// 编辑用户

func EditUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var data model.User
	_ = c.ShouldBindJSON(&data)
	var msg string
	if len(data.Username) < 4 || len(data.Username) > 16 {
		code = errmsg.ERROR
		msg = "用户名必须在4-12个字符之间"
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": msg,
		})
	} else {
		code = model.EditUser(id, &data)
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		})
	}
}

// 删除用户

func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code = model.DeleteUser(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
