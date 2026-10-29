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
// 这个文件只有骨架，课上跟着把 TODO 一个一个补完，补完就能跑通增删改查。
//
// 补的顺序：
//   ① initMySQL()   连接 MySQL
//   ② Todo          一条待办事项长什么样（已经写好了，先看懂）
//   ③ main()        连库 → 建表 → 注册路由 → 启动
//   ④ 五个处理函数   增、查全部、查单条、改、删
//
// 跑起来：在本目录执行 go run . ，监听 9000。


// DB 是全局的数据库连接对象。initMySQL() 写好后，所有接口都用它读写数据。
var DB *gorm.DB

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
//  3. Ping 一下，确认真的连上；连不上就把已经建立的连接关掉
//  4. 出错时把错误 return 出去
func initMySQL() error {
	// TODO ①：连接数据库
	return nil
}

// createTodo 添加一条：POST /v1/todo
//
// ① 从请求体里取出 title（用 c.ShouldBindJSON）
// ② 去掉首尾空格，空的就返回 400
// ③ 存进数据库：ID 由数据库生成，status 一开始是 false
// ④ 成功返回 201 和这条事项
func createTodo(c *gin.Context) {
	// TODO ④-1：添加一条
	c.JSON(http.StatusNotImplemented, gin.H{"error": "createTodo 还没实现"})
}

// listTodos 查看全部：GET /v1/todo
//
// 提示：make([]Todo, 0) 可以让空结果返回 [] 而不是 null
func listTodos(c *gin.Context) {
	// TODO ④-2：查询列表
	c.JSON(http.StatusNotImplemented, gin.H{"error": "listTodos 还没实现"})
}

// getTodo 查看单条：GET /v1/todo/:id
//
// 提示：c.Param("id") 拿到的是字符串，用 strconv.Atoi 转成数字
func getTodo(c *gin.Context) {
	// TODO ④-3：查询单条
	c.JSON(http.StatusNotImplemented, gin.H{"error": "getTodo 还没实现"})
}

// updateTodo 修改一条：PUT /v1/todo/:id
//
// 提示：status 用 *bool 接收，才能分清「传了 false」和「没传」
func updateTodo(c *gin.Context) {
	// TODO ④-4：修改
	c.JSON(http.StatusNotImplemented, gin.H{"error": "updateTodo 还没实现"})
}

// deleteTodo 删除一条：DELETE /v1/todo/:id
//
// 提示：除了检查 .Error，还要看 RowsAffected，确认真的删到了记录
func deleteTodo(c *gin.Context) {
	// TODO ④-5：删除
	c.JSON(http.StatusNotImplemented, gin.H{"error": "deleteTodo 还没实现"})
}

func main() {
	// TODO ③-1：先连数据库，失败就 panic 退出
	// TODO ③-2：defer DB.Close()，程序结束时关掉连接
	// TODO ③-3：DB.AutoMigrate(&Todo{}) 建表，并检查返回的 .Error

	r := gin.Default()

	// TODO ③-4：注册 /v1 分组和五个接口
	// v1Group := r.Group("/v1")
	// {
	//     v1Group.POST("/todo", createTodo)
	//     v1Group.GET("/todo", listTodos)
	//     v1Group.GET("/todo/:id", getTodo)
	//     v1Group.PUT("/todo/:id", updateTodo)
	//     v1Group.DELETE("/todo/:id", deleteTodo)
	// }

	if err := r.Run(":9000"); err != nil {
		log.Fatal(err)
	}
}
