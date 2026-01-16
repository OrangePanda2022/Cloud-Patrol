<template>
  <div class="min-h-screen bg-base-200 p-4 md:p-8">
    <!-- Header -->
    <div class="mb-6 flex flex-col gap-4 md:flex-row md:items-center md:justify-between">
      <div>
        <h1 class="text-2xl md:text-3xl font-bold">空中无人机地面垃圾巡检报告</h1>
        <p class="text-base-content/70">Drone-based Ground Garbage Inspection Report</p>
      </div>
    </div>

    <!-- Summary Cards -->
    <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-6 gap-4 mb-6">
      <div class="card bg-base-100 shadow">
        <div class="card-body">
          <h2 class="card-title text-sm">巡检区域</h2>
          <p class="text-3xl font-bold">12</p>
          <p class="text-sm text-success">+2 新增区域</p>
        </div>
      </div>
      <div class="card bg-base-100 shadow">
        <div class="card-body">
          <h2 class="card-title text-sm">地面散落垃圾</h2>
          <p class="text-3xl font-bold">31</p>
          <p class="text-sm text-warning">需人工清理</p>
        </div>
      </div>
      <div class="card bg-base-100 shadow">
        <div class="card-body">
          <h2 class="card-title text-sm">垃圾桶垃圾</h2>
          <p class="text-3xl font-bold">17</p>
          <p class="text-sm text-info">可统一转运</p>
        </div>
      </div>
      <div class="card bg-base-100 shadow">
        <div class="card-body">
          <h2 class="card-title text-sm">高风险区域</h2>
          <p class="text-3xl font-bold">5</p>
          <p class="text-sm text-error">重点关注</p>
        </div>
      </div>
      <div class="card bg-base-100 shadow">
        <div class="card-body">
          <h2 class="card-title text-sm">巡检完成率</h2>
          <p class="text-3xl font-bold">92%</p>
          <progress class="progress progress-primary w-full" value="92" max="100"></progress>
        </div>
      </div>
    </div>

    <!-- Main Content -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
      <!-- 地面散落垃圾表 -->
      <div class="card bg-base-100 shadow">
        <div class="card-body">
          <h2 class="card-title">地面散落垃圾巡检明细</h2>
          <div class="overflow-x-auto">
            <table class="table table-zebra">
              <thead>
                <tr>
                  <th>编号</th>
                  <th>位置</th>
                  <th>风险等级</th>
                  <th>状态</th>
                </tr>
              </thead>
              <tbody>
                <tr v-for="item in groundInspections" :key="item.id">
                  <td>{{ item.id }}</td>
                  <td>{{ item.location }}</td>
                  <td>
                    <span class="badge" :class="riskClass(item.risk)">{{ item.risk }}</span>
                  </td>
                  <td>
                    <span class="badge badge-outline">{{ item.status }}</span>
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>
      </div>

      <!-- 垃圾桶垃圾表 -->
      <div class="card bg-base-100 shadow">
        <div class="card-body">
          <h2 class="card-title">垃圾桶垃圾巡检明细</h2>
          <div class="overflow-x-auto">
            <table class="table table-zebra">
              <thead>
                <tr>
                  <th>编号</th>
                  <th>位置</th>
                  <th>风险等级</th>
                  <th>状态</th>
                </tr>
              </thead>
              <tbody>
                <tr v-for="item in binInspections" :key="item.id">
                  <td>{{ item.id }}</td>
                  <td>{{ item.location }}</td>
                  <td>
                    <span class="badge" :class="riskClass(item.risk)">{{ item.risk }}</span>
                  </td>
                  <td>
                    <span class="badge badge-outline">{{ item.status }}</span>
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>
      </div>

      <!-- Side Panel -->
      <div class="card bg-base-100 shadow">
        <div class="card-body">
          <h2 class="card-title">巡检信息</h2>
          <div class="space-y-3 text-sm">
            <div class="flex justify-between">
              <span>无人机型号</span>
              <span class="font-medium">DJI Mavic 3</span>
            </div>
            <div class="flex justify-between">
              <span>巡检时间</span>
              <span class="font-medium">2026-01-15</span>
            </div>
            <div class="flex justify-between">
              <span>飞行时长</span>
              <span class="font-medium">38 分钟</span>
            </div>
            <div class="flex justify-between">
              <span>操作人员</span>
              <span class="font-medium">张三</span>
            </div>
          </div>

          <div class="divider"></div>

          <h3 class="font-semibold mb-2">处理建议</h3>
          <ul class="list-disc list-inside text-sm space-y-1 text-base-content/80">
            <li>优先清理高风险垃圾点</li>
            <li>加强河道及绿化带巡检频率</li>
            <li>结合AI识别提高分类准确率</li>
          </ul>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'

const groundInspections = ref([
  { id: 'G-01', location: '东区公园草坪', risk: '高', status: '待处理' },
  { id: 'G-02', location: '河道沿线', risk: '中', status: '待处理' },
  { id: 'G-03', location: '道路交叉口', risk: '低', status: '已处理' },
])

const binInspections = ref([
  { id: 'B-01', location: '居民小区垃圾桶', risk: '中', status: '处理中' },
  { id: 'B-02', location: '商业街垃圾桶', risk: '高', status: '待处理' },
])

const riskClass = (risk) => {
  if (risk === '高') return 'badge-error'
  if (risk === '中') return 'badge-warning'
  return 'badge-success'
}
</script>

<style scoped></style>
