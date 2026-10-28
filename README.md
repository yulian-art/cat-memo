# 备忘录 · 课前版

一节 90 分钟的后端体验课：Go + Gin + MySQL，做出一个能存数据的备忘录。

这个分支是**上课用的起点**：前端已经写好，后端的 `main.go` 只有骨架和 TODO，课上跟着补完。

## 目录

| 目录 | 内容 |
| --- | --- |
| `memo-web` | 前端页面（Vue 3 + Vite），已完成 |
| `memo-backend` | 后端（Go + Gin + GORM），只有骨架 |

## 课前准备

Go、Node/npm、MySQL 都能在终端里查到版本号，MySQL 里建好库：

```sql
CREATE DATABASE IF NOT EXISTS db1 CHARACTER SET utf8mb4;
```

## 跑起来

```bash
# 终端 1：后端
cd memo-backend
go run .

# 终端 2：前端
cd memo-web
npm install
npm run dev
```

打开终端里给出的地址（默认 http://localhost:5173/）。页面能显示，但读不到数据：因为路由还没登记，接口会返回 404。

## 课上顺序

① `initMySQL()` ② `Todo` ③ `main()` ④ 五个处理函数

完整代码在另一个分支：`post-class`。
