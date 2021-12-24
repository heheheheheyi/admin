package routes

import (
	v1 "admin/api/v1"
	"admin/middleware"
	"admin/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() {
	//初始化路由

	//设定模式
	gin.SetMode(utils.AppMode)

	//创建gin框架对象
	r := gin.Default()
	//r :=gin.New()
	//r.Use(gin.Recovery())

	//跨域
	r.Use(middleware.Cors())

	r.LoadHTMLFiles("static/admin/index.html")
	r.Static("admin/static", "static/admin/static")
	r.GET("admin", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	r.StaticFile("admin/favicon.ico", "static/admin/favicon.ico")
	//自定义log
	//r.Use(middleware.Logger())

	//设置路由

	r.POST("api/v1/login", v1.Login)
	r.POST("api/v1/user/add", v1.AddUser)
	//需要token
	router := r.Group("api/v1")
	router.Use(middleware.JwtToken())
	{
		router.PUT("cate/:id", v1.EditCate)
		router.DELETE("cate/:id", v1.DeleteCate)
		router.GET("cate/:id", v1.GetOneCate)
		router.GET("cate", v1.GetCate)
		router.POST("cate/add", v1.AddCate)

		router.PUT("res/:id", v1.EditRes)
		router.DELETE("res/:id", v1.DeleteRes)
		router.GET("res/:id", v1.GetOneRes)
		router.GET("cateres/:cid", v1.GetCateRes)
		router.GET("res", v1.GetRes)
		router.POST("res/add", v1.AddRes)

		router.PUT("res/borrow/:id/:mid", v1.BorrowRes)
		router.GET("res/unborrow", v1.GetUnBorrow)
		router.GET("res/applyborrow", v1.GetApplyBorrow)
		router.DELETE("res/applyborrow/:id/:mid", v1.RejectBorrow)
		router.GET("res/borrow", v1.GetBorrowRes)
		router.GET("res/unborrow/:cid", v1.GetCateUnBorrow)
		router.GET("res/applyborrow/:cid", v1.GetCateApplyBorrow)
		router.GET("res/borrow/:cid", v1.GetCateBorrowRes)
		router.POST("res/borrow/:id", v1.ApplyBorrowRes)
		router.DELETE("res/borrow/:id/:mid", v1.ReturnRes)

		router.PUT("user/:id", v1.EditUser)
		router.GET("user/:id", v1.GetOneUser)
		router.GET("user", v1.GetUser)
		router.GET("userid", v1.GetUserId)
		router.DELETE("user/:id", v1.DeleteUser)

		router.GET("chat/:id", v1.GetChat)
		router.POST("chat/add", v1.AddChat)
		router.GET("chat/noread/:id", v1.GetNoreadChat)
		router.GET("chat/count/:id", v1.GetCount)

		router.GET("record/borrow", v1.GetBorrowRecord)
		router.GET("record/return", v1.GetReturnRecord)

		router.POST("upload", v1.UpLoad)
	}
	r.Run(utils.HttpPort)
}
