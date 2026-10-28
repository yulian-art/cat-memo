# 备忘录 · 课后版

90 分钟后端体验课的完整代码，三部分都能在本机跑起来。

| 目录 | 说明 | 启动 | 端口 |
| --- | --- | --- | --- |
| `memo-backend` | 本节课的完整后端（Gin + GORM + MySQL） | `go run .` | 9000 |
| `memo-web` | 前端页面（Vue 3 + Vite） | `npm install` 后 `npm run dev` | 5173 |
| `catDemo` | 旧版完整实现：Gin 直接返回页面，页面就是 memo-web 的构建产物 | `go run .` | 8080 |

## 准备

MySQL 建库（两个后端都用 `db1`，表由 AutoMigrate 建）：

```sql
CREATE DATABASE IF NOT EXISTS db1 CHARACTER SET utf8mb4;
```

## 跑起来

本节课的方式（前后端分离）：

```bash
# 终端 1
cd memo-backend && go run .                  # 9000

# 终端 2
cd memo-web && npm install && npm run dev    # 5173
```

旧版的方式（Gin 直接返回页面）：

```bash
cd catDemo && go run .                       # 8080
```

两个后端读写同一张表 `db1.todos`，所以在一边加的事项，另一边也能看到。

## 两个分支怎么切换

| 分支 | 内容 |
| --- | --- |
| `pre-class` | 上课起点：前端 + 后端骨架（默认分支） |
| `post-class` | 课后版：`catDemo` + 完整后端 + 前端 |

第一次拿完整代码：

```bash
git clone git@github.com:yulian-art/cat-memo.git
cd cat-memo
git checkout post-class          # clone 下来默认是 pre-class，要先切过来
```

切回上课起点：

```bash
git checkout pre-class
```

已经 clone 过、但本地没有 `post-class` 时，先拉一次远端分支再切换：

```bash
git fetch origin
git checkout post-class
```

两个分支分别拿最新内容：

```bash
git checkout pre-class && git pull
git checkout post-class && git pull
```

不装 git 也行：在 GitHub 页面左上角的分支下拉框里选 `post-class`，再点 Code → Download ZIP。
