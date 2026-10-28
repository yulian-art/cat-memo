package main

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DB 是全局的数据库连接对象，连上数据库之后，所有接口都通过它读写数据。
var DB *gorm.DB

// Todo 是备忘录里的一条事项。
// 三个字段分别对应数据库表 todos 的三列，也对应前端 JSON 里的三个字段。
type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

// initMySQL 连接本机的 MySQL。
// 返回 error 的原因是：连接失败时要把失败原因交给调用者，让程序决定是否继续启动。
func initMySQL() error {
	dsn := "root:root@tcp(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	var err error

	// 注意这里用 = 给包级变量 DB 赋值；写成 := 会变成局部变量，外面的 DB 仍然是 nil。
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		return err
	}

	// Ping 一下，确认连接真的可用。连不上就关掉已经建立的连接，避免资源泄漏。
	if err = DB.DB().Ping(); err != nil {
		DB.Close()
		return err
	}
	return nil
}

// parseID 从路径 /v1/todo/:id 里取出 id，并检查它是不是正整数。
// 不合法时这里已经写好了 400 响应，调用方直接 return 即可。
func parseID(c *gin.Context) (int, bool) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id 必须是正整数"})
		return 0, false
	}
	return id, true
}

// createTodo 处理「添加一条事项」：POST /v1/todo
func createTodo(c *gin.Context) {
	// 用专门的输入结构体接收请求，只允许客户端传 title。
	// ID 由数据库生成，status 由后端决定初始值，都不能让客户端随便指定。
	var input struct {
		Title string `json:"title"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请求数据格式不正确"})
		return
	}

	// 网页做过检查，后端也必须再检查一次：别人可以绕过网页直接发请求。
	title := strings.TrimSpace(input.Title)
	if title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "标题不能为空"})
		return
	}

	todo := Todo{Title: title, Status: false}
	if err := DB.Create(&todo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "保存失败"})
		return
	}

	// 201 表示「创建成功」，并把带 ID 的完整事项返回给前端。
	c.JSON(http.StatusCreated, todo)
}

// listTodos 处理「查看全部」：GET /v1/todo
func listTodos(c *gin.Context) {
	// make([]Todo, 0) 保证没有数据时返回 []，而不是 null，前端始终按数组处理。
	todoList := make([]Todo, 0)
	if err := DB.Find(&todoList).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "查询失败"})
		return
	}
	c.JSON(http.StatusOK, todoList)
}

// getTodo 处理「查看单条」：GET /v1/todo/:id
func getTodo(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		return
	}

	var todo Todo
	err := DB.First(&todo, id).Error
	if gorm.IsRecordNotFoundError(err) {
		// 编号格式对、但记录不存在，是 404；数据库故障不能也算成 404。
		c.JSON(http.StatusNotFound, gin.H{"error": "事项不存在"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "查询失败"})
		return
	}
	c.JSON(http.StatusOK, todo)
}

// updateTodo 处理「修改一条事项」：PUT /v1/todo/:id
func updateTodo(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		return
	}

	// status 用 *bool 接收：nil 表示请求里没带这个字段（或传了 null）。
	// 指针指向的值可能是 false，而 false 本身是合法取值，不能当成「没传」。
	var input struct {
		Title  string `json:"title"`
		Status *bool  `json:"status"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请求数据格式不正确"})
		return
	}

	title := strings.TrimSpace(input.Title)
	if title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "标题不能为空"})
		return
	}
	if input.Status == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "status 必须是 true 或 false"})
		return
	}

	// 先按 id 找到目标，再修改允许改的字段；不要直接把请求体绑到 Todo 上，否则 ID 也会被改。
	var todo Todo
	err := DB.First(&todo, id).Error
	if gorm.IsRecordNotFoundError(err) {
		c.JSON(http.StatusNotFound, gin.H{"error": "事项不存在"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "查询失败"})
		return
	}

	todo.Title = title
	todo.Status = *input.Status
	// 用 Save 保存整个对象，取消完成时的 false 也会被写进数据库。
	if err = DB.Save(&todo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "保存失败"})
		return
	}
	c.JSON(http.StatusOK, todo)
}

// deleteTodo 处理「删除一条事项」：DELETE /v1/todo/:id
func deleteTodo(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		return
	}

	// 删除必须带条件，不能漏掉 Where 直接删整张表。
	result := DB.Where("id = ?", id).Delete(Todo{})
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "删除失败"})
		return
	}
	// SQL 没报错，不代表真的删掉了记录：RowsAffected 为 0 说明一条都没匹配到。
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "事项不存在"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}

func main() {
	// 先连数据库：连不上，后面的接口都没有意义，直接结束程序并打印原因。
	if err := initMySQL(); err != nil {
		log.Fatalf("连接数据库失败：%v", err)
	}
	// defer 记下一件收尾工作：main 结束时关闭数据库连接。
	defer DB.Close()

	// 把 Todo 结构体和数据库表对应起来；表不存在时自动建表（表名默认是 todos）。
	if err := DB.AutoMigrate(&Todo{}).Error; err != nil {
		log.Fatalf("创建数据表失败：%v", err)
	}

	r := gin.Default()

	// 用分组给一批接口统一加上 /v1 前缀，组内写 /todo 即可。
	v1Group := r.Group("/v1")
	{
		v1Group.POST("/todo", createTodo)       // 添加
		v1Group.GET("/todo", listTodos)         // 查看全部
		v1Group.GET("/todo/:id", getTodo)       // 查看单条
		v1Group.PUT("/todo/:id", updateTodo)    // 修改
		v1Group.DELETE("/todo/:id", deleteTodo) // 删除
	}

	// 启动 HTTP 服务，监听 9000 端口。监听成功后这行会一直阻塞，直到按 Ctrl+C。
	if err := r.Run(":9000"); err != nil {
		log.Fatalf("启动服务失败：%v", err)
	}
}
