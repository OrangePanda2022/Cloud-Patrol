<template>
  <div class="min-h-screen bg-base-200">
    <!-- Main -->
    <main class="p-6">
      <!-- Filter -->
      <div class="card bg-base-100 shadow mb-6">
        <div class="card-body">
          <div class="grid grid-cols-1 md:grid-cols-4 gap-4">
            <!-- 活动名称搜索 -->
            <input
              v-model="keyword"
              type="text"
              placeholder="搜索活动名称"
              class="input input-bordered w-full md:col-span-2"
            />

            <!-- 分类筛选 -->
            <select v-model="selectedCategory" class="select select-bordered w-full">
              <option value="">全部分类</option>
              <option v-for="c in categories" :key="c" :value="c">
                {{ c }}
              </option>
            </select>

            <!-- 状态筛选 -->
            <select v-model="selectedStatus" class="select select-bordered w-full">
              <option value="">全部状态</option>
              <option value="ongoing">进行中</option>
              <option value="upcoming">即将开始</option>
              <option value="ended">已结束</option>
            </select>
          </div>
        </div>
      </div>

      <!-- Activity Grid -->
      <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
        <div
          v-for="activity in filteredActivities"
          :key="activity.id"
          class="card bg-base-100 shadow-xl hover:shadow-2xl transition"
        >
          <figure>
            <img :src="activity.image" alt="activity" class="h-48 w-full object-cover" />
          </figure>
          <div class="card-body">
            <h2 class="card-title">
              {{ activity.title }}
              <div class="badge badge-secondary">{{ activity.category }}</div>
            </h2>
            <p class="text-sm text-gray-500">{{ activity.date }}</p>
            <p class="line-clamp-2">{{ activity.description }}</p>
            <div class="card-actions justify-between items-center mt-4">
              <span class="badge" :class="statusClass(activity.status)">
                {{ statusText(activity.status) }}
              </span>
              <button class="btn btn-primary btn-sm" @click="router.push('/community/1')">
                查看详情
              </button>
            </div>
          </div>
        </div>
      </div>
    </main>
  </div>
</template>

<script setup>
import router from '@/router'

const categories = ['文艺', '体育', '学术', '社团']

const activities = ref([
  {
    id: 1,
    title: '校园音乐节',
    category: '文艺',
    status: 'ongoing',
    date: '2026-03-20',
    description: '一年一度的校园音乐盛典，感受青春与旋律。',
    image: 'https://picsum.photos/400/300?1',
  },
  {
    id: 2,
    title: '篮球联赛',
    category: '体育',
    status: 'upcoming',
    date: '2026-04-05',
    description: '各学院篮球队激情对抗，争夺冠军。',
    image: 'https://picsum.photos/400/300?2',
  },
  {
    id: 3,
    title: '人工智能讲座',
    category: '学术',
    status: 'ended',
    date: '2026-03-01',
    description: '探索AI前沿技术与未来发展趋势。',
    image: 'https://picsum.photos/400/300?3',
  },
])

const selectedCategory = ref('')
const selectedStatus = ref('')
const keyword = ref('')

const filteredActivities = computed(() => {
  return activities.value.filter((a) => {
    const matchCategory = !selectedCategory.value || a.category === selectedCategory.value
    const matchStatus = !selectedStatus.value || a.status === selectedStatus.value
    const matchKeyword = !keyword.value || a.title.includes(keyword.value)
    return matchCategory && matchStatus && matchKeyword
  })
})

const statusText = (status) => {
  return status === 'ongoing' ? '进行中' : status === 'upcoming' ? '即将开始' : '已结束'
}

const statusClass = (status) => {
  return status === 'ongoing'
    ? 'badge-success'
    : status === 'upcoming'
      ? 'badge-warning'
      : 'badge-neutral'
}
</script>

<style scoped>
/* 可根据需要扩展自定义样式 */
</style>
