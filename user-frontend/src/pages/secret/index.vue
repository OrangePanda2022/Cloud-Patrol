<template>
  <div class="min-h-screen bg-base-200">
    <!-- Main Content -->
    <main class="max-w-5xl mx-auto p-4 grid grid-cols-1 lg:grid-cols-3 gap-6">
      <!-- Post Input -->
      <section class="lg:col-span-1 card bg-base-100 shadow-xl">
        <div class="card-body">
          <h2 class="card-title">发布心情</h2>
          <textarea
            v-model="newPost"
            class="textarea textarea-bordered w-full"
            placeholder="匿名说点什么吧…"
            rows="4"
          />
          <div class="card-actions justify-end">
            <button class="btn btn-primary" @click="addPost">发布</button>
          </div>
        </div>
      </section>

      <!-- Posts List -->
      <section class="lg:col-span-2 space-y-4">
        <div
          v-for="post in displayedPosts"
          :key="post.id"
          class="card bg-base-100 shadow-md hover:shadow-xl transition"
        >
          <div class="card-body">
            <div class="flex justify-between items-center">
              <span class="badge badge-outline">匿名</span>
              <span class="text-sm opacity-60">{{ post.time }}</span>
            </div>
            <p class="mt-2">{{ post.content }}</p>
            <div class="card-actions justify-end mt-2">
              <button class="btn btn-ghost btn-sm">👍 {{ post.likes }}</button>
            </div>
          </div>
        </div>

        <div v-if="displayedPosts.length === 0" class="text-center opacity-60">
          还没有内容，来发布第一条吧 🌱
        </div>

        <div ref="loadMoreTrigger" class="h-6"></div>
      </section>
    </main>
  </div>
</template>

<script setup>
import { ref, reactive, computed, onMounted } from 'vue'

const newPost = ref('')
const pageSize = 2
const page = ref(1)
const posts = ref([
  { id: 1, content: '今天下雨，在图书馆坐了一下午，感觉很治愈。', time: '10 分钟前', likes: 12 },
  { id: 2, content: '期末周好累，希望大家都能顺利通过考试。', time: '1 小时前', likes: 20 },
  { id: 3, content: '终于完成了实验报告，轻松一下。', time: '2 小时前', likes: 5 },
  { id: 4, content: '图书馆的人好多啊，找不到座位。', time: '3 小时前', likes: 8 },
  { id: 5, content: '明天有考试，好紧张...', time: '5 小时前', likes: 3 },
])

const displayedPosts = ref([])

function loadMore() {
  const start = displayedPosts.value.length
  const end = start + pageSize
  displayedPosts.value.push(...posts.value.slice(start, end))
}

function addPost() {
  if (!newPost.value.trim()) return
  posts.value.unshift({ id: Date.now(), content: newPost.value, time: '刚刚', likes: 0 })
  newPost.value = ''
  displayedPosts.value = []
  page.value = 1
  loadMore()
}

const loadMoreTrigger = ref(null)

onMounted(() => {
  loadMore()
  const observer = new IntersectionObserver(
    (entries) => {
      if (entries[0].isIntersecting) {
        loadMore()
      }
    },
    { threshold: 1 },
  )

  if (loadMoreTrigger.value) {
    observer.observe(loadMoreTrigger.value)
  }
})
</script>

<style>
/* 可根据需要添加自定义动画或样式 */
</style>
