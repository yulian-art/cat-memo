// 所有请求都以 /v1 开头。
// 开发时由 Vite 代理（vite.config.js）转发到 Go 后端的 9000 端口。
const BASE = '/v1'

// 最多等 6 秒。后端没启动时不要让页面一直转圈。
const TIMEOUT_MS = 6000

// 统一处理一次请求：既检查 HTTP 状态，也检查后端放在响应体里的 error 字段。
async function request(path, options = {}) {
  let response

  try {
    response = await fetch(BASE + path, {
      headers: { 'Content-Type': 'application/json' },
      signal: AbortSignal.timeout(TIMEOUT_MS),
      ...options,
    })
  } catch (err) {
    // 超时和「根本连不上」都发生在拿到响应之前，需要单独说清楚。
    if (err.name === 'TimeoutError') {
      throw new Error('后端没有响应，先确认 memo-backend 已经在 9000 端口运行')
    }
    throw new Error('请求没有发出去，检查 Vite 是否还在运行')
  }

  let data = null
  try {
    data = await response.json()
  } catch {
    // 响应不是 JSON，常见原因是代理没配上，请求被 Vite 当成了页面请求。
    throw new Error('返回内容不是 JSON，先检查 Vite 代理和后端是否在 9000 端口运行')
  }

  if (!response.ok || (data && data.error)) {
    throw new Error((data && data.error) || `请求失败（HTTP ${response.status}）`)
  }
  return data
}

// 查看全部事项
export function listTodos() {
  return request('/todo')
}

// 添加一条事项，只发送标题，id 和 status 由后端决定
export function createTodo(title) {
  return request('/todo', {
    method: 'POST',
    body: JSON.stringify({ title }),
  })
}

// 修改一条事项：本课约定要带上完整的可修改字段
export function updateTodo(id, { title, status }) {
  return request(`/todo/${id}`, {
    method: 'PUT',
    body: JSON.stringify({ title, status }),
  })
}

// 删除一条事项
export function deleteTodo(id) {
  return request(`/todo/${id}`, { method: 'DELETE' })
}