<template>
  <div class="min-h-screen flex items-center justify-center p-4">
    <div class="card w-full max-w-md bg-base-100 shadow-xl overflow-hidden">
      <figure class="px-10 pt-10 flex flex-col gap-2">
        <div class="avatar placeholder">
          <div class="bg-primary text-primary-content rounded-full w-16">
            <span class="text-2xl font-bold"></span>
          </div>
        </div>
        <h2 class="card-title text-2xl font-bold">欢迎回来</h2>
        <p class="text-sm text-base-content/60">请登录您的账号以继续</p>
      </figure>

      <div class="card-body pt-6">
        <form @submit.prevent="handleLogin">
          <div class="form-control w-full">
            <label class="label">
              <span class="label-text font-semibold">电子邮箱</span>
            </label>
            <input
              v-model="email"
              type="email"
              placeholder="email@example.com"
              class="input input-bordered focus:input-primary w-full transition-all"
              required
            />
          </div>

          <div class="form-control w-full mt-4">
            <label class="label">
              <span class="label-text font-semibold">密码</span>
              <a href="#" class="label-text-alt link link-hover text-primary">忘记密码？</a>
            </label>
            <input
              v-model="password"
              type="password"
              placeholder="请输入密码"
              class="input input-bordered focus:input-primary w-full transition-all"
              required
            />
          </div>

          <div class="form-control mt-2">
            <label class="label cursor-pointer justify-start gap-3">
              <input type="checkbox" class="checkbox checkbox-primary checkbox-sm" />
              <span class="label-text">记住我</span>
            </label>
          </div>

          <div class="card-actions mt-6">
            <button class="btn btn-primary btn-block text-white" :disabled="loading">
              <span v-if="loading" class="loading loading-bars loading-md"></span>

              <span v-else>立即登录</span>
            </button>
          </div>
        </form>

        <div class="divider text-xs text-base-content/40 my-6">或者使用以下方式登录</div>

        <div class="grid grid-cols-2 gap-4">
          <button class="btn btn-outline btn-sm gap-2">
            <svg class="w-4 h-4" fill="currentColor" viewBox="0 0 24 24">
              <path
                d="M12.48 10.92v3.28h7.84c-.24 1.84-.909 3.264-2.09 4.44-1.12 1.12-2.8 2.32-5.75 2.32-4.57 0-8.21-3.69-8.21-8.26s3.64-8.26 8.21-8.26c2.48 0 4.38.97 5.73 2.23l2.3-2.3C18.66 2.58 15.91 1.5 12.48 1.5 6.13 1.5 1 6.63 1 13s5.13 11.5 11.48 11.5c3.48 0 6.11-1.15 8.16-3.29 2.11-2.11 2.78-5.09 2.78-7.46 0-.71-.06-1.39-.18-1.99h-10.76z"
              />
            </svg>
            Google
          </button>
          <button class="btn btn-outline btn-sm gap-2">
            <svg class="w-4 h-4" fill="currentColor" viewBox="0 0 24 24">
              <path
                d="M12 0c-6.626 0-12 5.373-12 12 0 5.302 3.438 9.8 8.207 11.387.599.111.793-.261.793-.577v-2.234c-3.338.726-4.041-1.412-4.041-1.412-.546-1.387-1.333-1.756-1.333-1.756-1.089-.745.083-.729.083-.729 1.205.084 1.838 1.236 1.838 1.236 1.07 1.835 2.809 1.305 3.495.998.108-.776.417-1.305.76-1.605-2.665-.3-5.466-1.332-5.466-5.93 0-1.31.465-2.38 1.235-3.22-.135-.303-.54-1.523.105-3.176 0 0 1.005-.322 3.3 1.23.96-.267 1.98-.399 3-.405 1.02.006 2.04.138 3 .405 2.28-1.552 3.285-1.23 3.285-1.23.645 1.653.24 2.873.12 3.176.765.84 1.23 1.91 1.23 3.22 0 4.61-2.805 5.625-5.475 5.92.42.36.81 1.096.81 2.22v3.293c0 .319.192.694.805.576 4.765-1.589 8.199-6.086 8.199-11.386 0-6.627-5.373-12-12-12z"
              />
            </svg>
            GitHub
          </button>
        </div>

        <p class="text-center text-sm mt-8">
          还没有账号？
          <a href="#" class="link link-primary font-bold">立即注册</a>
        </p>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import router from '@/router'
const email = ref('')
const password = ref('')
const loading = ref(false)

const handleLogin = () => {
  loading.value = true
  // 模拟 API 请求
  setTimeout(() => {
    console.log('登录信息:', { email: email.value, password: password.value })
    loading.value = false
    alert('登录成功！')
  }, 1000)
  router.push({ path: '/' })
}
</script>
