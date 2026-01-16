<template>
  <div class="min-h-screen bg-base-200 p-4 md:p-8">
    <!-- Header -->
    <div class="flex flex-col md:flex-row md:items-center md:justify-between gap-4 mb-6">
      <div>
        <h1 class="text-2xl md:text-3xl font-bold">固定资产管理</h1>
        <p class="text-base-content/70">AI 大模型智能检测中</p>
      </div>
      <div class="flex gap-2">
        <button class="btn btn-primary btn-sm md:btn-md">新增资产</button>
        <button class="btn btn-outline btn-sm md:btn-md">导出数据</button>
      </div>
    </div>

    <!-- Stats -->
    <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-4 mb-8">
      <div class="stat bg-base-100 shadow rounded-2xl">
        <div class="stat-title">资产总数</div>
        <div class="stat-value text-primary">128</div>
        <div class="stat-desc">全部固定资产</div>
      </div>
      <div class="stat bg-base-100 shadow rounded-2xl">
        <div class="stat-title">路灯数量</div>
        <div class="stat-value text-secondary">86</div>
        <div class="stat-desc">正常运行中</div>
      </div>
      <div class="stat bg-base-100 shadow rounded-2xl">
        <div class="stat-title">配电箱数量</div>
        <div class="stat-value">42</div>
        <div class="stat-desc">含备用设备</div>
      </div>
      <div class="stat bg-base-100 shadow rounded-2xl">
        <div class="stat-title">异常设备</div>
        <div class="stat-value text-error">3</div>
        <div class="stat-desc">需尽快处理</div>
      </div>
    </div>

    <!-- Filter -->
    <div class="card bg-base-100 shadow mb-6 rounded-2xl">
      <div class="card-body">
        <div class="flex flex-col md:flex-row gap-4">
          <select class="select select-bordered w-full md:max-w-xs">
            <option>全部类型</option>
            <option>路灯</option>
            <option>配电箱</option>
          </select>
          <select class="select select-bordered w-full md:max-w-xs">
            <option>全部状态</option>
            <option>正常</option>
            <option>故障</option>
            <option>维护中</option>
          </select>
          <input
            type="text"
            placeholder="搜索资产编号 / 位置"
            class="input input-bordered w-full md:flex-1"
          />
        </div>
      </div>
    </div>

    <!-- Table -->
    <div class="card bg-base-100 shadow rounded-2xl overflow-x-auto">
      <table class="table table-zebra">
        <thead>
          <tr>
            <th>资产编号</th>
            <th>类型</th>
            <th>位置</th>
            <th>状态</th>
            <th>安装日期</th>
            <th class="text-right">操作</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="item in assets" :key="item.id">
            <td>{{ item.code }}</td>
            <td>
              <span
                class="badge"
                :class="item.type === '路灯' ? 'badge-secondary' : 'badge-accent'"
              >
                {{ item.type }}
              </span>
            </td>
            <td>{{ item.location }}</td>
            <td>
              <span
                class="badge"
                :class="{
                  'badge-success': item.status === '正常',
                  'badge-error': item.status === '故障',
                  'badge-warning': item.status === '维护中',
                }"
              >
                {{ item.status }}
              </span>
            </td>
            <td>{{ item.installDate }}</td>
            <td class="text-right gap-2 flex flex-wrap">
              <button class="btn btn-xs btn-outline">详情</button>
              <button class="btn btn-xs btn-error btn-outline">删除</button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'

const assets = ref([
  {
    id: 1,
    code: 'LD-001',
    type: '路灯',
    location: '人民路与建设路交叉口',
    status: '正常',
    installDate: '2022-05-12',
  },
  {
    id: 2,
    code: 'PDX-014',
    type: '配电箱',
    location: '滨河公园南门',
    status: '维护中',
    installDate: '2021-11-03',
  },
  {
    id: 3,
    code: 'LD-045',
    type: '路灯',
    location: '科技大道西段',
    status: '故障',
    installDate: '2020-08-19',
  },
])
</script>

<style scoped>
/* 可在此扩展个性化动画或细节样式 */
</style>
