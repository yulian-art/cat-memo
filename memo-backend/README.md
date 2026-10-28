# memo-backend

备忘录后端，Go + Gin + GORM v1 + MySQL，提供 `/v1/todo` 的增删改查接口。

## 本地运行

```bash
# 1. 准备数据库（只需一次；建表由 AutoMigrate 负责）
mysql -h 127.0.0.1 -u root -p -e "CREATE DATABASE IF NOT EXISTS db1 CHARACTER SET utf8mb4"

# 2. 启动
go run .
```

监听 `:9000`。连接信息写在 `main.go` 的 `dsn` 里（`root:root@tcp(127.0.0.1:3306)/db1`），需要与本机一致。

## 接口

| 操作 | 方法与路径 | 请求体 | 响应 |
| --- | --- | --- | --- |
| 添加 | `POST /v1/todo` | `{"title":"复习 HTTP"}` | `200` + 带 `id` 的事项 |
| 列表 | `GET /v1/todo` | 无 | `200` + 数组（空库为 `null`） |
| 单查 | `GET /v1/todo/:id` | 无 | 还没实现，处理函数是空的 |
| 修改 | `PUT /v1/todo/:id` | `{"title":"复习 HTTP","status":true}` | `200` + 修改后的事项 |
| 删除 | `DELETE /v1/todo/:id` | 无 | `200` + `{"<id>":"deleted"}` |

出错时 HTTP 仍然是 `200`，响应体里给 `{"error":"原因"}`，前端会同时检查这两样。

## 快速验证

```bash
curl -i http://127.0.0.1:9000/v1/todo

curl -i -X POST http://127.0.0.1:9000/v1/todo \
  -H 'Content-Type: application/json' \
  -d '{"title":"完成第一次接口练习"}'
```

`main.go` 里的结构：`initMySQL()` 连接数据库，`main()` 里连库、建表、注册五个接口，每个接口的处理函数直接写在路由后面。