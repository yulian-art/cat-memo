# memo-backend

备忘录后端，Go + Gin + GORM v1 + MySQL，提供 `/v1/todo` 的增删改查接口。

## 本地运行

```bash
# 1. 准备数据库（只需一次；建库，建表由 AutoMigrate 负责）
mysql -h 127.0.0.1 -u root -p -e "CREATE DATABASE IF NOT EXISTS db1 CHARACTER SET utf8mb4"

# 2. 启动
go run .
```

监听 `:9000`。连接信息写在 `main.go` 的 `dsn` 里（`root:root@tcp(127.0.0.1:3306)/db1`），需要与本机一致。

## 接口

| 操作 | 方法与路径 | 请求体 | 成功响应 |
| --- | --- | --- | --- |
| 添加 | `POST /v1/todo` | `{"title":"复习 HTTP"}` | `201` + 事项对象 |
| 列表 | `GET /v1/todo` | 无 | `200` + 数组（空为 `[]`） |
| 单查 | `GET /v1/todo/:id` | 无 | `200` + 事项对象 |
| 修改 | `PUT /v1/todo/:id` | `{"title":"复习 HTTP","status":true}` | `200` + 事项对象 |
| 删除 | `DELETE /v1/todo/:id` | 无 | `200` + `{"message":"deleted"}` |

错误统一为 `{"error":"简短说明"}`：输入错误 `400`，记录不存在 `404`，数据库故障 `500`。

## 快速验证

```bash
curl -i http://127.0.0.1:9000/v1/todo

curl -i -X POST http://127.0.0.1:9000/v1/todo \
  -H 'Content-Type: application/json' \
  -d '{"title":"完成第一次接口练习"}'
```

`main.go` 里的结构：`initMySQL()` 连接数据库、`parseID()` 校验路径参数、五个处理函数分别对应五个接口。