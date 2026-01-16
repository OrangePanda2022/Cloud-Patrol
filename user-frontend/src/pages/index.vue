<template>
  <div class="min-h-screen bg-base-100">
    <MapView />
    <!-- Main -->
    <main class="container mx-auto px-4 py-6 grid grid-cols-1 lg:grid-cols-4 gap-6">
      <!-- Content -->
      <section class="lg:col-span-3 space-y-6">
        <article
          v-for="item in news"
          :key="item.id"
          class="card bg-base-100 shadow hover:shadow-lg transition"
        >
          <figure class="w-full h-48 md:h-60">
            <img :src="item.image" alt="news" class="object-cover h-full w-full" />
          </figure>
          <div class="card-body">
            <h2 class="card-title text-lg md:text-xl">
              {{ item.title }}
              <div class="badge badge-secondary">热</div>
            </h2>
            <p class="text-sm text-base-content/70 line-clamp-3">
              {{ item.summary }}
            </p>
            <div class="card-actions justify-between items-center">
              <span class="text-xs text-base-content/60">
                {{ item.source }} · {{ item.time }}
              </span>
              <button class="btn btn-sm btn-outline" @click="router.push('/article/1')">
                阅读
              </button>
            </div>
          </div>
        </article>
      </section>

      <!-- Sidebar -->
      <aside class="space-y-4">
        <div class="card bg-base-200 shadow">
          <div class="card-body">
            <h3 class="card-title">热点榜</h3>
            <ul class="space-y-2">
              <li v-for="(hot, index) in hotList" :key="hot" class="flex items-center gap-2">
                <span class="badge badge-error">{{ index + 1 }}</span>
                <span class="text-sm truncate">{{ hot }}</span>
              </li>
            </ul>
          </div>
        </div>
      </aside>
    </main>
  </div>
</template>

<script setup>
import router from '@/router'

const news = ref([
  {
    id: 1,
    title: '一只小猫，成功“占领”了校园',
    summary: '某个午后，一只突然现身的校园小猫在教学楼、草坪和食堂门口来回穿梭。',
    image: 'https://picsum.photos/400/300?1',
    source: '校园趣事',
  },
  {
    id: 2,
    title: '绿茵沸腾的那一刻',
    summary:
      '奔跑的身影、精准的传球、看台上此起彼伏的呐喊声交织在一起，汗水与激情在绿茵场上肆意挥洒。',
    image: 'https://picsum.photos/400/300?2',
    source: '精彩体育',
    time: '4小时前',
  },
])

const hotList = ref([
  '这一幕出现时，所有人都停下了脚步',

  '原本平常的一天，却突然刷爆了朋友圈',

  '没人预料到，这件小事会引发全场围观',

  '当那一刻发生，现场气氛彻底变了',
])
</script>

<style>
/* 可在此扩展自定义样式 */
</style>
