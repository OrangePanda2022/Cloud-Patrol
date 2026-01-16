<template>
  <div class="min-h-screen bg-base-200">
    <!-- 内容区域 -->
    <main class="max-w-5xl mx-auto p-6">
      <!-- 活动封面 -->
      <div class="card bg-base-100 shadow-xl">
        <figure>
          <img :src="activity.image" alt="activity cover" class="w-full h-72 object-cover" />
        </figure>

        <div class="card-body">
          <!-- 标题与分类 -->
          <div class="flex flex-col md:flex-row md:items-center md:justify-between gap-4">
            <h2 class="text-3xl font-bold">{{ activity.title }}</h2>
            <div class="flex gap-2">
              <span class="badge badge-secondary">{{ activity.category }}</span>
              <span class="badge" :class="statusClass(activity.status)">
                {{ statusText(activity.status) }}
              </span>
            </div>
          </div>

          <!-- 基本信息 -->
          <div class="grid grid-cols-1 md:grid-cols-3 gap-4 mt-6 text-sm">
            <div class="stat bg-base-200 rounded-box">
              <div class="stat-title">活动时间</div>
              <div class="stat-value text-lg">{{ activity.date }}</div>
            </div>
            <div class="stat bg-base-200 rounded-box">
              <div class="stat-title">活动地点</div>
              <div class="stat-value text-lg">{{ activity.location }}</div>
            </div>
            <div class="stat bg-base-200 rounded-box">
              <div class="stat-title">主办方</div>
              <div class="stat-value text-lg">{{ activity.organizer }}</div>
            </div>
          </div>

          <!-- 活动介绍 -->
          <div class="mt-8">
            <h3 class="text-xl font-semibold mb-2">活动介绍</h3>
            <p class="leading-relaxed text-base-content/80">
              {{ activity.description }}
            </p>
          </div>

          <!-- 报名按钮 -->
          <div class="card-actions justify-end mt-10">
            <a
              :href="activity.signupUrl"
              target="_blank"
              rel="noopener noreferrer"
              class="btn btn-primary btn-lg"
            >
              前往外部应用报名
            </a>
          </div>
        </div>
      </div>
    </main>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'

const router = useRouter()

const activity = ref({
  title: '校园音乐节',
  category: '文艺',
  status: 'ongoing',
  date: '2026-03-20 18:30',
  location: '学校大礼堂',
  organizer: '校学生会',
  description:
    '校园音乐节汇集了校园内外优秀乐队与歌手，是一次展示青春与才华的舞台。欢迎全校师生前来参与，一起感受音乐的魅力。',
  image: 'https://picsum.photos/900/400?music',
  signupUrl: 'https://example.com/signup-app',
})

const goBack = () => {
  router.back()
}

const statusText = (status) =>
  status === 'ongoing' ? '进行中' : status === 'upcoming' ? '即将开始' : '已结束'

const statusClass = (status) =>
  status === 'ongoing' ? 'badge-success' : status === 'upcoming' ? 'badge-warning' : 'badge-neutral'
</script>

<style scoped>
/* 详情页可在此添加动画或自定义样式 */
</style>
