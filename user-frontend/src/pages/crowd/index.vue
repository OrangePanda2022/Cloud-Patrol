<template>
  <div class="min-h-screen bg-base-200 p-4 md:p-8">
    <!-- Header -->
    <div class="navbar bg-base-100 rounded-box shadow mb-6">
      <div class="flex-1">
        <h1 class="text-xl md:text-2xl font-bold">空中无人机人流检测</h1>
      </div>
      <div class="flex-none gap-2">
        <span class="badge badge-success">系统运行中</span>
      </div>
    </div>

    <!-- Overview Cards -->
    <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-4 mb-6">
      <div class="card bg-base-100 shadow">
        <div class="card-body">
          <h2 class="card-title text-sm">当前监测人数</h2>
          <p class="text-3xl font-bold text-primary">1,284</p>
          <p class="text-xs opacity-70">实时统计</p>
        </div>
      </div>
      <div class="card bg-base-100 shadow">
        <div class="card-body">
          <h2 class="card-title text-sm">人流密度</h2>
          <p class="text-3xl font-bold text-warning">中等</p>
          <p class="text-xs opacity-70">每平方米 2.3 人</p>
        </div>
      </div>
      <div class="card bg-base-100 shadow">
        <div class="card-body">
          <h2 class="card-title text-sm">预警等级</h2>
          <p class="text-3xl font-bold text-error">Ⅱ 级</p>
          <p class="text-xs opacity-70">需重点关注</p>
        </div>
      </div>
      <div class="card bg-base-100 shadow">
        <div class="card-body">
          <h2 class="card-title text-sm">无人机状态</h2>
          <p class="text-3xl font-bold text-success">正常</p>
          <p class="text-xs opacity-70">电量 78%</p>
        </div>
      </div>
    </div>

    <!-- Main Content -->
    <div class="grid grid-cols-1 lg:grid-cols-3 gap-6">
      <!-- Left: Chart / Image -->
      <div class="lg:col-span-2 card bg-base-100 shadow">
        <div class="card-body">
          <h2 class="card-title">监测区域概览</h2>
          <div class="aspect-video rounded-box bg-base-200 flex items-center justify-center">
            <span class="opacity-60">（此处可接入无人机实时画面或人流热力图）</span>
          </div>
        </div>
      </div>

      <!-- Right: Warning List -->
      <div class="card bg-base-100 shadow">
        <div class="card-body">
          <h2 class="card-title">AI 智能预警报告</h2>
          <ul class="space-y-3">
            <li class="alert alert-warning text-sm">
              <span>⚠️ A 区域人流快速上升</span>
            </li>
            <li class="alert alert-error text-sm">
              <span>🚨 B 区域已接近安全阈值</span>
            </li>
            <li class="alert alert-info text-sm">
              <span>ℹ️ 大模型智能建议：调度安保人员</span>
            </li>
          </ul>
        </div>
      </div>
    </div>

    <!-- Report Table -->
    <div class="card bg-base-100 shadow mt-6">
      <div class="card-body">
        <h2 class="card-title">历史监测记录</h2>
        <div class="overflow-x-auto">
          <table class="table table-zebra w-full">
            <thead>
              <tr>
                <th>时间</th>
                <th>区域</th>
                <th>人数</th>
                <th>密度</th>
                <th>预警状态</th>
              </tr>
            </thead>
            <tbody>
              <tr>
                <td>10:30</td>
                <td>A 区</td>
                <td>420</td>
                <td>高</td>
                <td><span class="badge badge-warning">Ⅱ 级</span></td>
              </tr>
              <tr>
                <td>10:00</td>
                <td>B 区</td>
                <td>510</td>
                <td>较高</td>
                <td><span class="badge badge-error">Ⅰ 级</span></td>
              </tr>
              <tr>
                <td>09:30</td>
                <td>C 区</td>
                <td>280</td>
                <td>正常</td>
                <td><span class="badge badge-success">正常</span></td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
// 页面为展示型示例，可在此接入真实 API 或 WebSocket 实时数据
</script>

<style scoped>
/* 可根据需要补充自定义样式 */
</style>
