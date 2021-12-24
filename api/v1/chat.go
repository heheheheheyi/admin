package v1

import (
	"admin/model"
	"admin/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 添加用户

func AddChat(c *gin.Context) {
	var data model.Chat
	_ = c.ShouldBindJSON(&data)
	code = model.AddChat(&data)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// 查询聊天记录
func GetChat(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, code := model.GetChat(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// 查询未读聊天记录
func GetNoreadChat(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, code := model.GetNoreadChat(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// 查询未读条数
func GetCount(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	count, code := model.GetCount(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"count":   count,
		"message": errmsg.GetErrMsg(code),
	})
}
