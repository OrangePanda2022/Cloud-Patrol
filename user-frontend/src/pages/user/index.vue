<script setup>
// 模拟用户信息数据
const user = reactive({
  name: '李华',
  studentId: '20230512',
  college: '计算机科学与技术学院',
  avatar: 'https://img.daisyui.com/images/stock/photo-1534528741775-53994a69daeb.jpg',
  level: 'Lv.5',
})

// 模拟统计数据
const stats = reactive([
  {
    label: '我的发布',
    value: 24,
    icon: 'M19 11H5m14 0a2 2 0 012 2v6a2 2 0 01-2 2H5a2 2 0 01-2-2v-6a2 2 0 012-2m14 0V9a2 2 0 00-2-2M5 11V9a2 2 0 012-2m0 0V5a2 2 0 012-2h6a2 2 0 012 2v2M7 7h10',
  },
  {
    label: '获赞',
    value: '1.2k',
    icon: 'M4.318 6.318a4.5 4.5 0 000 6.364L12 20.364l7.682-7.682a4.5 4.5 0 00-6.364-6.364L12 7.636l-1.318-1.318a4.5 4.5 0 00-6.364 0z',
  },
  { label: '收藏', value: 85, icon: 'M5 5a2 2 0 012-2h10a2 2 0 012 2v16l-7-3.5L5 21V5z' },
])

// 核心服务菜单 (针对校园场景)
const coreServices = [
  { name: '我的预约', icon: '📅', color: 'text-primary', desc: '场馆/自习室' },
  { name: '校园树洞', icon: '🌲', color: 'text-success', desc: '匿名发布的贴子' },
  { name: '活动报名', icon: '🎫', color: 'text-warning', desc: '已报名的活动' },
  { name: '失物招领', icon: '👜', color: 'text-secondary', desc: '我的发布' },
]
</script>

<template>
  <div class="min-h-screen bg-base-200 py-6 px-4 md:px-0 font-sans text-base-content">
    <div class="max-w-md mx-auto space-y-6">
      <div class="card bg-base-100 shadow-xl overflow-hidden relative group">
        <div class="card-body pt-16 items-center text-center relative z-1">
          <div class="avatar online">
            <div
              class="w-24 rounded-full ring ring-base-100 ring-offset-base-100 ring-offset-2 shadow-lg"
            >
              <img :src="user.avatar" alt="Avatar" />
            </div>
          </div>

          <div class="mt-2">
            <h2 class="card-title justify-center text-2xl font-bold">
              {{ user.name }}
              <div class="badge badge-accent badge-outline text-xs">{{ user.level }}</div>
            </h2>
            <p class="text-sm opacity-60 mt-1">{{ user.college }}</p>
            <p class="text-xs opacity-50 font-mono">ID: {{ user.studentId }}</p>
          </div>
        </div>
      </div>

      <div class="stats shadow w-full rounded-box bg-base-100">
        <div v-for="(stat, index) in stats" :key="index" class="stat place-items-center p-4">
          <div class="stat-figure text-primary opacity-60">
            <svg
              xmlns="http://www.w3.org/2000/svg"
              class="h-6 w-6"
              fill="none"
              viewBox="0 0 24 24"
              stroke="currentColor"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                :d="stat.icon"
              />
            </svg>
          </div>
          <div class="stat-title text-xs">{{ stat.label }}</div>
          <div class="stat-value text-xl">{{ stat.value }}</div>
        </div>
      </div>

      <div class="card bg-base-100 shadow-lg">
        <div class="card-body p-4">
          <h3 class="font-bold text-sm text-base-content/70 mb-2 px-2">常用服务</h3>
          <div class="grid grid-cols-4 gap-2">
            <button
              v-for="service in coreServices"
              :key="service.name"
              class="flex flex-col items-center justify-center p-2 rounded-xl hover:bg-base-200 transition-colors duration-200 active:scale-95"
            >
              <div :class="`text-2xl mb-1 ${service.color}`">{{ service.icon }}</div>
              <span class="text-xs font-medium">{{ service.name }}</span>
            </button>
          </div>
        </div>
      </div>

      <div class="card bg-base-100 shadow-lg overflow-hidden">
        <ul class="menu bg-base-100 w-full p-0 [&_li>*]:rounded-none">
          <li>
            <a class="py-4 justify-between">
              <div class="flex items-center gap-3">
                <span class="bg-error/10 text-error p-2 rounded-lg">
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    class="h-5 w-5"
                    fill="none"
                    viewBox="0 0 24 24"
                    stroke="currentColor"
                  >
                    <path
                      stroke-linecap="round"
                      stroke-linejoin="round"
                      stroke-width="2"
                      d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z"
                    />
                  </svg>
                </span>
                <span>预警通知设置</span>
              </div>
              <div class="badge badge-sm badge-error">99+</div>
            </a>
          </li>

          <li>
            <a class="py-4 justify-between">
              <div class="flex items-center gap-3">
                <span class="bg-info/10 text-info p-2 rounded-lg">
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    class="h-5 w-5"
                    fill="none"
                    viewBox="0 0 24 24"
                    stroke="currentColor"
                  >
                    <path
                      stroke-linecap="round"
                      stroke-linejoin="round"
                      stroke-width="2"
                      d="M15 17h5l-1.405-1.405A2.032 2.032 0 0118 14.158V11a6.002 6.002 0 00-4-5.659V5a2 2 0 10-4 0v.341C7.67 6.165 6 8.388 6 11v3.159c0 .538-.214 1.055-.595 1.436L4 17h5m6 0v1a3 3 0 11-6 0v-1m6 0H9"
                    />
                  </svg>
                </span>
                <span>消息通知</span>
              </div>
              <svg
                xmlns="http://www.w3.org/2000/svg"
                class="h-4 w-4 opacity-50"
                fill="none"
                viewBox="0 0 24 24"
                stroke="currentColor"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M9 5l7 7-7 7"
                />
              </svg>
            </a>
          </li>

          <li>
            <a class="py-4 justify-between">
              <div class="flex items-center gap-3">
                <span class="bg-base-content/10 text-base-content p-2 rounded-lg">
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    class="h-5 w-5"
                    fill="none"
                    viewBox="0 0 24 24"
                    stroke="currentColor"
                  >
                    <path
                      stroke-linecap="round"
                      stroke-linejoin="round"
                      stroke-width="2"
                      d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z"
                    />
                    <path
                      stroke-linecap="round"
                      stroke-linejoin="round"
                      stroke-width="2"
                      d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"
                    />
                  </svg>
                </span>
                <span>通用设置</span>
              </div>
              <svg
                xmlns="http://www.w3.org/2000/svg"
                class="h-4 w-4 opacity-50"
                fill="none"
                viewBox="0 0 24 24"
                stroke="currentColor"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M9 5l7 7-7 7"
                />
              </svg>
            </a>
          </li>
        </ul>
      </div>

      <div class="px-4">
        <button class="btn btn-outline btn-error w-full">退出登录</button>
      </div>

      <div class="text-center text-xs text-base-content/30 pb-4">Smart Campus v2.0.1</div>
    </div>
  </div>
</template>

<style scoped>
/* 针对移动端的微调，隐藏滚动条但保留功能 */
.no-scrollbar::-webkit-scrollbar {
  display: none;
}
.no-scrollbar {
  -ms-overflow-style: none;
  scrollbar-width: none;
}
</style>
