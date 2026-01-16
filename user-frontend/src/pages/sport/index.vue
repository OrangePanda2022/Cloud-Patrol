<template>
  <div class="min-h-screen bg-base-200 p-4 md:p-8">
    <!-- Header -->
    <div class="max-w-7xl mx-auto mb-8 text-center">
      <h1 class="text-3xl md:text-4xl font-bold text-base-content mb-2">校园体育场馆</h1>
      <p class="text-base-content/70">选择你感兴趣的体育场馆类型，开启健康校园生活</p>
    </div>

    <!-- Category Grid -->
    <div class="max-w-7xl mx-auto grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
      <div
        v-for="item in categories"
        :key="item.id"
        class="card bg-base-100 shadow-xl hover:shadow-2xl transition duration-300"
      >
        <figure class="h-40 overflow-hidden">
          <img :src="item.image" :alt="item.name" class="w-full h-full object-cover" />
        </figure>
        <div class="card-body">
          <h2 class="card-title">
            {{ item.name }}
            <div class="badge badge-primary">{{ item.count }}个</div>
          </h2>
          <p class="text-sm text-base-content/70">
            {{ item.description }}
          </p>
          <div class="card-actions justify-end">
            <button class="btn btn-primary btn-sm" @click="router.push('/sport/1')">
              查看详情
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import router from '@/router'

const categories = ref([
  {
    id: 1,
    name: '足球场',
    count: 3,
    description: '标准11人制与5人制足球场，适合日常训练与比赛。',
    image: 'https://images.unsplash.com/photo-1508098682722-e99c43a406b2',
  },
  {
    id: 2,
    name: '篮球馆',
    count: 5,
    description: '室内外篮球场，配备夜间照明设施。',
    image: 'https://images.unsplash.com/photo-1517649763962-0c623066013b',
  },
  {
    id: 3,
    name: '田径场',
    count: 2,
    description: '400米标准跑道，满足教学与运动会需求。',
    image: 'https://images.unsplash.com/photo-1521412644187-c49fa049e84d',
  },
  {
    id: 4,
    name: '羽毛球馆',
    count: 6,
    description: '专业地胶与灯光，适合课余锻炼。',
    image: 'https://images.unsplash.com/photo-1600508774634-4e11d34730e2',
  },
  {
    id: 5,
    name: '游泳馆',
    count: 1,
    description: '恒温泳池，支持教学与自由游泳。',
    image: 'https://images.unsplash.com/photo-1509228627159-6452c9a5c92f',
  },
  {
    id: 6,
    name: '健身房',
    count: 2,
    description: '配备基础与力量训练器械。',
    image: 'https://images.unsplash.com/photo-1571902943202-507ec2618e8f',
  },
])
</script>

<style scoped>
/* 可在此处补充自定义动画或细节样式 */
</style>
