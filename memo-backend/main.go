package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB *gorm.DB
)
type Todo struct {
    ID     int    `json:"id"`
    Title  string `json:"title"`
    Status bool   `json:"status"`
}

func initMySQL()(err error){
	dsn := "root:root@tcp(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql",dsn)
	if err != nil {
		return
	}
	err = DB.DB().Ping()
	return

}

func main() {
    // 创建数据库
    // sql: create database cat;
    // 连接数据库
	err := initMySQL()
	if err != nil {
		panic(err)
	}
	defer DB.Close()  // 程序退出，关闭数据库连接
	// 模型绑定
	DB.AutoMigrate(Todo{})
    // 创建gin框架
    r := gin.Default()
    
    

    
    
    
    // v1
    v1Group := r.Group("v1")
    {
        // 待办事项
        // 添加
        v1Group.POST("/todo", func(c *gin.Context){
            // 前端页面填写待办事项 点击提交 会发请求到这里
            // 1. 从请求中把数据拿出来
            var todo Todo
            c.BindJSON(&todo)
            
            // 2. 存入数据库
            if err = DB.Create(&todo).Error; err != nil {
                // 3. 返回响应
                c.JSON(http.StatusOK, gin.H{"error": err.Error()})
                
            }else {
                // 3. 返回响应
                c.JSON(http.StatusOK, todo)
            }
            

        })
        // 查看
        v1Group.GET("/todo", func(c *gin.Context){
            // 查询todo表里的所有数据
            var todoList []Todo
            if err = DB.Find(&todoList).Error; err != nil {
                c.JSON(http.StatusOK, gin.H{"error": err.Error()})
            }else {
                c.JSON(http.StatusOK, todoList)
            }
        })
        // 查看单个可以在后面跟id，此id对应数据库中存的ID
        v1Group.GET("/todo/:id", func(c *gin.Context){
            
        })
        // 修改
        v1Group.PUT("/todo/:id", func(c *gin.Context){
            id := c.Params.ByName("id")
            var todo Todo
            if err = DB.Where("id=?",id).First(&todo).Error; err != nil {
                c.JSON(http.StatusOK,gin.H{"error": err.Error()})
                return
            }
            c.BindJSON(&todo)
            if err =  DB.Save(&todo).Error; err != nil {
                c.JSON(http.StatusOK, gin.H{"error":err.Error()})
            } else {
                c.JSON(http.StatusOK, todo)
            }
        })
        // 删除
        v1Group.DELETE("/todo/:id", func(c *gin.Context){
            id := c.Params.ByName("id")
            if err = DB.Where("id=?",id).Delete(Todo{}).Error; err != nil {
                c.JSON(http.StatusOK,gin.H{"error": err.Error()})
            }else {
                c.JSON(http.StatusOK, gin.H{id : "deleted"})
            }
            
        })
    }
	// 启动 HTTP 服务，监听 8080 端口
    r.Run(":9000")
}