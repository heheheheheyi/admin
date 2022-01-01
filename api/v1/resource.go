package v1

import (
	"admin/model"
	"admin/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 添加资源
func AddRes(c *gin.Context) {
	var data model.Res
	_ = c.ShouldBindJSON(&data)
	var msg string
	if data.Resname == "" || data.Cid == 0 || data.Img == "" {
		code = errmsg.ERROR
		msg = "参数不完整"
	} else {
		code = errmsg.SUCCESS
	}
	if code != errmsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": msg,
		})
		return
	}
	code = model.CreateRes(&data)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// 查询单个资源

func GetOneRes(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, code := model.GetOneRes(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

func GetCateRes(c *gin.Context) {
	cid, _ := strconv.Atoi(c.Param("cid"))
	resname := c.Query("resname")
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))

	data, total, code := model.GetCateRes(cid, resname, pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
		"total":   total,
	})
}

// 查询资源列表

func GetRes(c *gin.Context) {
	resname := c.Query("resname")
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))

	data, total, code := model.GetRes(resname, pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
		"total":   total,
	})
}

// 编辑资源

func EditRes(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var data model.Res
	_ = c.ShouldBindJSON(&data)
	code = model.EditRes(id, &data)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// 删除资源

func DeleteRes(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code = model.DeleteRes(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

//确认借出
func BorrowRes(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	mid, _ := strconv.Atoi(c.Param("mid"))
	code = model.BorrowRes(id, mid)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

//拒绝借出
func RejectBorrow(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	mid, _ := strconv.Atoi(c.Param("mid"))
	code = model.RejectBorrow(id, mid)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

//未借出列表
func GetUnBorrow(c *gin.Context) {
	resname := c.Query("resname")
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))

	data, total, code := model.GetUnBorrow(resname, pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
		"total":   total,
	})
}

//待确认列表
func GetApplyBorrow(c *gin.Context) {
	resname := c.Query("resname")
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))

	data, total, code := model.GetApplyBorrow(resname, pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
		"total":   total,
	})
}

//已借出列表
func GetBorrowRes(c *gin.Context) {
	resname := c.Query("resname")
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))

	data, total, code := model.GetBorrowRes(resname, pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
		"total":   total,
	})
}

//未借出列表(分类)
func GetCateUnBorrow(c *gin.Context) {
	cid, _ := strconv.Atoi(c.Param("cid"))
	resname := c.Query("resname")
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))

	data, total, code := model.GetCateUnBorrow(cid, resname, pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
		"total":   total,
	})
}

//待确认列表(分类)
func GetCateApplyBorrow(c *gin.Context) {
	cid, _ := strconv.Atoi(c.Param("cid"))
	resname := c.Query("resname")
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))

	data, total, code := model.GetCateApplyBorrow(cid, resname, pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
		"total":   total,
	})
}

//已借出列表(分类)
func GetCateBorrowRes(c *gin.Context) {
	cid, _ := strconv.Atoi(c.Param("cid"))
	resname := c.Query("resname")
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))

	data, total, code := model.GetCateBorrowRes(cid, resname, pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
		"total":   total,
	})
}

//申请借出资源
func ApplyBorrowRes(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var data model.Res
	_ = c.ShouldBindJSON(&data)
	code = model.ApplyBorrowRes(id, data.Uid, data.BorrowTime, data.ReturnTime)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

//归还资源
func ReturnRes(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	mid, _ := strconv.Atoi(c.Param("mid"))
	code = model.ReturnRes(id, mid)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
