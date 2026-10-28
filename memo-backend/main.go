package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// ================================================================
// 备忘录后端 · 课上练习版
//
// 这个文件只有骨架，课上跟着 TODO 一个个补完。
//
// 补的顺序：
//   ① initMySQL()  连接 MySQL
//   ② main()       连库 → 建表 → 注册路由 → 启动
//   ③ 五个接口      添加、查看全部、查看单个、修改、删除
//
// 跑起来：在本目录执行 go run . ，监听 9000。
// ================================================================

// DB 是全局的数据库连接对象。initMySQL() 写好后，所有接口都用它读写数据。
var (
	DB *gorm.DB
)

// Todo 是一条待办事项：编号、标题、完成没有。
// json 标签决定前端看到的字段名；表名不用自己取，GORM 会把 Todo 变成 todos。
type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

// initMySQL 连接本机 MySQL。
//
// 要补的：
//  1. dsn 写清 用户名:密码@tcp(地址:端口)/库名，本机是 root:root@tcp(127.0.0.1:3306)/db1
//  2. DB, err = gorm.Open("mysql", dsn)
//     注意这里是 = 不是 :=，写成 := 会新开一个局部变量，外面的 DB 还是空的
//  3. Ping 一下，确认真的连上
//  4. 出错时把错误 return 出去
func initMySQL() (err error) {
	// TODO ①：连接数据库
	return
}

func main() {
	// TODO ②-1：先连数据库，失败就 panic
	// TODO ②-2：defer DB.Close()，程序结束时关掉连接
	// TODO ②-3：DB.AutoMigrate(Todo{}) 建表

	r := gin.Default()

	// v1
	v1Group := r.Group("v1")
	{
		// 待办事项
		// 添加
		v1Group.POST("/todo", func(c *gin.Context) {
			// TODO ③-1：从请求里拿数据 → 存进数据库 → 回 JSON
		})
		// 查看全部
		v1Group.GET("/todo", func(c *gin.Context) {
			// TODO ③-2：查出 todos 表里的所有数据，回给前端
		})
		// 查看单个，此 id 对应数据库中存的 ID
		v1Group.GET("/todo/:id", func(c *gin.Context) {
			// TODO ③-3（课后）：按 id 查一条
		})
		// 修改
		v1Group.PUT("/todo/:id", func(c *gin.Context) {
			// TODO ③-4：按 id 查出来 → 绑定请求体 → Save
		})
		// 删除
		v1Group.DELETE("/todo/:id", func(c *gin.Context) {
			// TODO ③-5：按 id 删除
		})
	}

	// 启动 HTTP 服务，监听 9000 端口
	r.Run(":9000")
}