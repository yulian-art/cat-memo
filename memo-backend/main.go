package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)


// 备忘录后端 · 课上练习版
//
// 今天要做的三件事：
//   1. 让服务跑起来：访问 /v1/todo，能回一段 JSON
//   2. 连上 MySQL，把事项存进去
//   3. 添加、查看、修改、删除四个接口
//
// 下面从 main 开始写。忘了下一步做什么，就回来看这几行。
// ================================================================

func main() {
	// 先在这里创建 gin 的路由器，再让它监听 9000 端口
}
