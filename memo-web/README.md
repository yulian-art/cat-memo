# memo-web

备忘录前端，Vue 3 + Vite，配套后端在 `../memo-backend`（Go + Gin + MySQL）。

## 本地运行

先确认后端已启动（另开一个终端）：

```bash
cd ../memo-backend
go run .          # 监听 9000
```

再启动前端：

```bash
npm install       # 首次运行需要
npm run dev       # 默认 http://localhost:5173/
```

浏览器打开终端给出的地址，页面上可以添加、勾选、删除待办事项。

## 目录

```
src/
├── main.js      # 创建 Vue 应用并挂载到 #app
├── App.vue      # 页面：输入框、列表、勾选、删除、加载与错误提示
├── api.js       # 四个接口的封装，统一处理 HTTP 错误
└── style.css    # 样式
```

## 接口约定

前端只请求相对路径 `/v1/todo`，开发时由 Vite 代理转发到 `http://127.0.0.1:9000`（见 `vite.config.js`）。

| 操作 | 方法与路径 | 请求体 |
| --- | --- | --- |
| 列表 | `GET /v1/todo` | 无 |
| 添加 | `POST /v1/todo` | `{"title":"..."}` |
| 修改 | `PUT /v1/todo/:id` | `{"title":"...","status":true}` |
| 删除 | `DELETE /v1/todo/:id` | 无 |

后端成功返回 `200`（创建为 `201`），失败返回 `400` / `404` / `500` 和 `{"error":"..."}`。

## 生产构建

```bash
npm run build     # 生成 dist/
```

`server.proxy` 只对开发服务器生效。直接双击 `dist/index.html` 不等于部署成功，部署需要单独配置静态服务与接口转发。