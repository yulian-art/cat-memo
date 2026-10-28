package routers

import (
	"bubble/controller"
	"bubble/setting"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	if setting.Conf.Release {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()

	// 页面用到的 js/css 在 static/assets 里（和 memo-web 的 dist/assets 同名同内容）
	r.Static("/assets", "static/assets")
	// 页面图标
	r.StaticFile("/favicon.svg", "static/favicon.svg")
	// 告诉 gin 去哪里找页面文件
	r.LoadHTMLGlob("templates/*")
	r.GET("/", controller.IndexHandler)

	// v1
	v1Group := r.Group("v1")
	{
		// 待办事项
		// 添加
		v1Group.POST("/todo", controller.CreateTodo)
		// 查看所有的待办事项
		v1Group.GET("/todo", controller.GetTodoList)
		// 修改某一个待办事项
		v1Group.PUT("/todo/:id", controller.UpdateATodo)
		// 删除某一个待办事项
		v1Group.DELETE("/todo/:id", controller.DeleteATodo)
	}
	return r
}
