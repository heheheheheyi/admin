package v1

import (
	"admin/model"
	"admin/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 添加分类

func AddCate(c *gin.Context) {
	var data model.Cate
	_ = c.ShouldBindJSON(&data)
	code := model.CreateCate(&data)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// 查询单个分类

func GetOneCate(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, code := model.GetOneCate(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// 查询分类列表

func GetCate(c *gin.Context) {
	catename := c.Query("catename")
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))

	data, total, code := model.GetCate(catename, pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
		"total":   total,
	})
}

// 编辑分类

func EditCate(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var data model.Cate
	_ = c.ShouldBindJSON(&data)
	code = model.EditCate(id, &data)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// 删除分类

func DeleteCate(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code = model.DeleteCate(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
