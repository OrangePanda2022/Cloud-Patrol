<template>
  <div class="min-h-screen bg-base-200 p-4 md:p-8">
    <div class="max-w-6xl mx-auto">
      <!-- Header -->
      <div class="mb-6">
        <h1 class="text-3xl font-bold mb-2">{{ venue.name }}</h1>
        <div class="flex flex-wrap gap-2">
          <span class="badge badge-primary">{{ venue.type }}</span>
          <span class="badge badge-outline">{{ venue.location }}</span>
        </div>
      </div>

      <!-- Main Info -->
      <div class="grid grid-cols-1 lg:grid-cols-3 gap-6">
        <!-- Left: Image & Map -->
        <div class="lg:col-span-2 space-y-6">
          <div class="card bg-base-100 shadow">
            <figure class="h-64">
              <img :src="venue.image" alt="场馆图片" class="w-full h-full object-cover" />
            </figure>
          </div>

          <!-- Distribution / Map -->
          <div class="card bg-base-100 shadow">
            <div class="card-body">
              <h2 class="card-title">场馆分布</h2>
              <p class="text-base-content/70">
                {{ venue.distribution }}
              </p>
              <div class="mt-2">
                <div class="mockup-window bg-base-300">
                  <div class="flex justify-center items-center h-40 text-base-content/50">
                    校园地图占位区域
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- Right: Status -->
        <div class="space-y-6">
          <!-- Open Time -->
          <div class="card bg-base-100 shadow">
            <div class="card-body">
              <h2 class="card-title">开放时间</h2>
              <ul class="space-y-2 text-sm">
                <li
                  v-for="(time, index) in venue.openTime"
                  :key="index"
                  class="flex justify-between"
                >
                  <span>{{ time.day }}</span>
                  <span class="font-medium">{{ time.time }}</span>
                </li>
              </ul>
            </div>
          </div>

          <!-- Field Condition -->
          <div class="card bg-base-100 shadow">
            <div class="card-body">
              <h2 class="card-title">场地状态</h2>

              <div class="flex items-center gap-3">
                <span class="badge" :class="fieldConditionInfo.badge">
                  {{ fieldConditionInfo.text }}
                </span>
              </div>

              <p class="text-sm text-base-content/70 mt-2">
                {{ fieldConditionInfo.desc }}
              </p>
            </div>
          </div>

          <!-- Usage -->
          <div class="card bg-base-100 shadow">
            <div class="card-body">
              <h2 class="card-title">当前使用情况</h2>
              <div class="space-y-3">
                <div>
                  <div class="flex justify-between text-sm mb-1">
                    <span>使用率</span>
                    <span>{{ venue.usage }}%</span>
                  </div>
                  <progress
                    class="progress progress-primary"
                    :value="venue.usage"
                    max="100"
                  ></progress>
                </div>
                <div class="text-sm text-base-content/70">
                  {{ venue.usageDesc }}
                </div>
              </div>
            </div>
          </div>

          <!-- Action -->
          <div class="card bg-base-100 shadow">
            <div class="card-body">
              <button class="btn btn-accent w-full" @click="showAiAdvice(advice)">AI 建议</button>
              <button class="btn btn-primary w-full">预约使用</button>
              <button class="btn btn-outline w-full mt-2">查看使用记录</button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import router from '@/router'
const conditionMap = {
  dry: {
    text: '干燥（适合使用）',
    badge: 'badge-success',
    desc: '场地干燥，适合训练和比赛。',
  },
  wet: {
    text: '潮湿（谨慎使用）',
    badge: 'badge-warning',
    desc: '场地略微潮湿，注意防滑。',
  },
  closed: {
    text: '积水 / 暂停使用',
    badge: 'badge-error',
    desc: '场地存在积水，暂不开放。',
  },
}

const fieldConditionInfo = computed(() => {
  return conditionMap[venue.fieldCondition]
})

// 示例数据（后续可通过路由 + API 动态获取）
const venue = reactive({
  name: '东区标准足球场',
  type: '足球场',
  location: '校园东区体育中心',
  image: 'https://images.unsplash.com/photo-1508098682722-e99c43a406b2',
  distribution: '位于校园东区体育中心核心区域，靠近学生宿舍与主干道，交通便利。',
  openTime: [
    { day: '周一至周五', time: '08:00 - 22:00' },
    { day: '周六', time: '09:00 - 22:00' },
    { day: '周日', time: '09:00 - 20:00' },
  ],
  usage: 68,
  usageDesc: 'AI 出行建议：当前为中等使用状态，晚间时段较为繁忙。',
  fieldCondition: 'wet', // dry | wet | closed
})

const advice =
  '东区标准足球场坐落于校园东区体育中心核心区域，因紧邻学生宿舍与主干道而具备极佳的交通便利性。目前场馆处于中等使用状态（占用率68%），但受环境影响地面较为潮湿，晚间时段人流较为繁忙。我希望去这里运动，你能给我一些建议吗？'

const showAiAdvice = (question) => {
  router.push({
    path: '/chat',
    query: { question },
  })
}
</script>

<style scoped>
/* 可扩展地图、状态动画等 */
</style>
