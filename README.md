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

打开终端里给出的地址（默认 http://localhost:5173/）。页面能显示，但读不到数据：五个接口的处理函数都还是空的，返回的是空响应。

## 课上顺序

① `initMySQL()` ② `Todo` ③ `main()` ④ 五个处理函数

完整代码在另一个分支：`post-class`。

## 两个分支怎么切换

| 分支 | 内容 |
| --- | --- |
| `pre-class` | 上课起点：前端 + 后端骨架（默认分支） |
| `post-class` | 课后版：`catDemo` + 完整后端 + 前端 |

第一次拿代码：

```bash
git clone git@github.com:yulian-art/cat-memo.git
cd cat-memo
git branch -a          # 看远端有哪些分支
```

clone 下来默认就在 `pre-class`。想看课后版：

```bash
git checkout post-class          # 第一次会自动跟踪 origin/post-class
```

再切回上课起点：

```bash
git checkout pre-class
```

已经 clone 过、但本地没有另一个分支时，先拉一次远端分支再切换：

```bash
git fetch origin
git checkout post-class
```

在某个分支上想拿最新内容：

```bash
git pull
```

不想装 git 的话，在 GitHub 页面左上角的分支下拉框里选 `post-class`，再点 Code → Download ZIP，也能直接下到课后版。
