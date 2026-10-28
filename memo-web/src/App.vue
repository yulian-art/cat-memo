<script setup>
import { onMounted, ref } from 'vue'
import { createTodo, deleteTodo, listTodos, updateTodo } from './api'

const todos = ref([]) // 列表数据，来自后端
const newTitle = ref('') // 输入框里的内容
const loading = ref(false) // 是否正在和后端通信
const errorMessage = ref('') // 出错时显示给用户看的文字

// 从后端重新读一次列表。
// 添加、修改、删除成功之后都调用它，让页面显示的就是数据库里的数据。
async function loadTodos() {
  loading.value = true
  try {
    const data = await listTodos()
    todos.value = Array.isArray(data) ? data : []
  } catch (err) {
    errorMessage.value = err.message
  } finally {
    loading.value = false
  }
}

// 添加：先检查输入，再发 POST，成功之后重新读列表
async function handleAdd() {
  const title = newTitle.value.trim()
  if (!title) {
    errorMessage.value = '标题不能为空'
    return
  }

  loading.value = true
  errorMessage.value = ''
  try {
    await createTodo(title)
    newTitle.value = ''
    await loadTodos()
  } catch (err) {
    errorMessage.value = err.message
  } finally {
    loading.value = false
  }
}

// 勾选 / 取消勾选：把新的状态发给后端
async function handleToggle(todo) {
  loading.value = true
  errorMessage.value = ''
  try {
    await updateTodo(todo.id, { title: todo.title, status: !todo.status })
    await loadTodos()
  } catch (err) {
    errorMessage.value = err.message
    await loadTodos() // 失败时也重新读一次，让界面回到数据库里的真实状态
  } finally {
    loading.value = false
  }
}

// 删除
async function handleRemove(todo) {
  loading.value = true
  errorMessage.value = ''
  try {
    await deleteTodo(todo.id)
    await loadTodos()
  } catch (err) {
    errorMessage.value = err.message
  } finally {
    loading.value = false
  }
}

// 组件挂载后先读一次列表
onMounted(loadTodos)
</script>

<template>
  <main class="app">
    <header class="app-header">
      <h1>备忘录</h1>
      <p class="hint">写下一件要做的事，点「添加」。勾选表示完成，右边的按钮可以删除。</p>
    </header>

    <form class="add-form" @submit.prevent="handleAdd">
      <input
        v-model="newTitle"
        class="add-input"
        type="text"
        maxlength="100"
        placeholder="比如：今晚去跑步"
        :disabled="loading"
      />
      <button class="add-button" type="submit" :disabled="loading">添加</button>
    </form>

    <p v-if="errorMessage" class="error">{{ errorMessage }}</p>

    <p v-if="loading && todos.length === 0" class="state">正在读取列表…</p>

    <ul v-else-if="todos.length > 0" class="todo-list">
      <li v-for="todo in todos" :key="todo.id" class="todo-item" :class="{ done: todo.status }">
        <input
          class="todo-check"
          type="checkbox"
          :checked="todo.status"
          :disabled="loading"
          @change="handleToggle(todo)"
        />
        <span class="todo-title">{{ todo.title }}</span>
        <button class="todo-delete" type="button" :disabled="loading" @click="handleRemove(todo)">
          删除
        </button>
      </li>
    </ul>

    <p v-else-if="!errorMessage && !loading" class="state">还没有事项，先添加一条吧。</p>

    <footer class="app-footer">数据来自 Go 后端 + MySQL，刷新页面不会丢。</footer>
  </main>
</template>