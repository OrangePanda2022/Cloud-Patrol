<template>
  <div class="min-h-screen bg-base-200">
    <!-- Main Content -->
    <main class="max-w-6xl mx-auto px-4 py-6 grid grid-cols-1 lg:grid-cols-4 gap-6">
      <!-- Article -->
      <article class="lg:col-span-3 bg-base-100 rounded-2xl shadow p-6">
        <!-- Title -->
        <h2 class="text-3xl font-bold mb-3">校园科技文化节圆满举行</h2>

        <!-- Meta -->
        <div class="flex flex-wrap items-center gap-4 text-sm text-base-content/60 mb-6">
          <span>📅 2026-01-10</span>
          <span>✍️ 校园融媒体中心</span>
          <span class="badge badge-primary badge-outline">校园活动</span>
        </div>

        <!-- Cover -->
        <figure class="mb-6">
          <img
            src="https://images.unsplash.com/photo-1523050854058-8df90110c9f1"
            alt="校园活动"
            class="rounded-xl w-full object-cover max-h-[360px]"
          />
        </figure>

        <!-- Content -->
        <div class="prose max-w-none prose-headings:font-bold prose-img:rounded-xl">
          <p>
            为期三天的校园科技文化节在同学们的热情参与中圆满落幕。本次活动以
            <strong>“科技 · 创新 · 青春”</strong> 为主题，展示了我校在科技创新和文化建设方面的成果。
          </p>

          <h3>活动亮点</h3>
          <ul>
            <li>学生科技作品展览</li>
            <li>人工智能体验区</li>
            <li>社团文化舞台表演</li>
          </ul>

          <p>活动不仅丰富了校园文化生活，也激发了同学们对科技创新的浓厚兴趣。</p>
        </div>

        <!-- Actions -->
        <div class="mt-8 flex flex-wrap gap-3">
          <button class="btn btn-primary">👍 点赞</button>
        </div>
      </article>

      <!-- Sidebar -->
      <aside class="space-y-6">
        <!-- Author Card -->
        <div class="bg-base-100 rounded-2xl shadow p-5">
          <div class="flex items-center gap-4">
            <div class="avatar">
              <div class="w-14 rounded-full">
                <img src="https://i.pravatar.cc/150" />
              </div>
            </div>
            <div>
              <h4 class="font-bold">校园融媒体中心</h4>
              <p class="text-sm text-base-content/60">官方账号</p>
            </div>
          </div>
        </div>

        <!-- Related Articles -->
        <div class="bg-base-100 rounded-2xl shadow p-5">
          <h4 class="font-bold mb-4">相关阅读</h4>
          <ul class="space-y-3">
            <li class="link link-hover">🎓 新生开学指南</li>
            <li class="link link-hover">🏀 校园篮球联赛回顾</li>
            <li class="link link-hover">📚 图书馆学习攻略</li>
          </ul>
        </div>
      </aside>
    </main>

    <!-- Footer -->
    <footer class="text-center text-sm text-base-content/60 py-6">© 2026 校园资讯平台</footer>
  </div>
</template>

<script setup>
// 可在此通过 API 获取文章详情数据
</script>

<style scoped>
/* 额外细节可在这里补充 */
</style>
