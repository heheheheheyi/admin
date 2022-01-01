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

	//登录接口
	r.POST("api/v1/login", v1.Login)
	//注册接口
	r.POST("api/v1/user/add", v1.AddUser)

	//需要token
	router := r.Group("api/v1")
	router.Use(middleware.JwtToken())
	{
		//更新分类
		router.PUT("cate/:id", v1.EditCate)
		//删除分类
		router.DELETE("cate/:id", v1.DeleteCate)
		//获取单个分类信息
		router.GET("cate/:id", v1.GetOneCate)
		//获取全部分类信息
		router.GET("cate", v1.GetCate)
		//添加分类
		router.POST("cate/add", v1.AddCate)

		//更新资源
		router.PUT("res/:id", v1.EditRes)
		//删除资源
		router.DELETE("res/:id", v1.DeleteRes)
		//获取单个资源信息
		router.GET("res/:id", v1.GetOneRes)
		//获取分类下资源信息
		router.GET("cateres/:cid", v1.GetCateRes)
		//获取资源信息
		router.GET("res", v1.GetRes)
		//添加资源
		router.POST("res/add", v1.AddRes)

		//借出资源
		router.PUT("res/borrow/:id/:mid", v1.BorrowRes)
		//获取未借出资源
		router.GET("res/unborrow", v1.GetUnBorrow)
		//获取申请借出资源
		router.GET("res/applyborrow", v1.GetApplyBorrow)
		//拒绝借出资源
		router.DELETE("res/applyborrow/:id/:mid", v1.RejectBorrow)
		//获取已借出资源
		router.GET("res/borrow", v1.GetBorrowRes)
		//获取分类下未借出资源
		router.GET("res/unborrow/:cid", v1.GetCateUnBorrow)
		//获取分类下申请借出资源
		router.GET("res/applyborrow/:cid", v1.GetCateApplyBorrow)
		//获取分类下已借出资源
		router.GET("res/borrow/:cid", v1.GetCateBorrowRes)
		//申请借出资源
		router.POST("res/borrow/:id", v1.ApplyBorrowRes)
		//归还资源
		router.DELETE("res/borrow/:id/:mid", v1.ReturnRes)

		//更新用户
		router.PUT("user/:id", v1.EditUser)
		//获取单个用户信息
		router.GET("user/:id", v1.GetOneUser)
		//获取用户信息
		router.GET("user", v1.GetUser)
		//获取用户id
		router.GET("userid", v1.GetUserId)
		//删除用户
		router.DELETE("user/:id", v1.DeleteUser)

		//获取聊天记录
		router.GET("chat/:id", v1.GetChat)
		//发送信息
		router.POST("chat/add", v1.AddChat)
		//获取用户未读信息
		router.GET("chat/noread/:id", v1.GetNoreadChat)
		//获取用户信息总数
		router.GET("chat/count/:id", v1.GetCount)

		//获取借出记录
		router.GET("record/borrow", v1.GetBorrowRecord)
		//获取归还记录
		router.GET("record/return", v1.GetReturnRecord)

		//上传到七牛云
		router.POST("upload", v1.UpLoad)
	}
	r.Run(utils.HttpPort)
}
