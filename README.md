# 📡 Cloud-Patrol

Cloud-Patrol 是一个面向 **低空经济与智慧校园深度融合** 的全栈系统，聚焦于低空设备在校园场景下的 **巡检、监控、调度与数据可视化管理**，为智慧校园建设提供低空智能化解决方案。

---

## 🚀 项目简介

Cloud-Patrol 是一个现代化的云资源巡检与管理平台，旨在帮助用户：

- **实时监控** 云端服务状态

- **收集与展示** 云资源指标与告警

- **提供可视化界面** 用于操作与查看数据

- 支持扩展多种云平台（如 AWS / GCP / Azure）

本项目使用前后端分离架构，前端负责用户交互和展示，后端负责数据处理和 API 提供。

---

## 📁 项目结构

```
Cloud-Patrol/
├── tmp-backend/           # 后端服务代码
│   ├── src/
│   ├── config/
│   └── ...
├── user-frontend/         # 前端界面代码
│   ├── src/
│   ├── public/
│   └── ...
├── README.md              # 项目说明文档
└── .gitignore
```

---

## ✨ 功能亮点

### 🛩️ 低空巡检管理

- 无人机巡检任务创建与调度

- 巡检路线规划与记录

- 巡检状态实时监控

### 🏫 智慧校园应用

- 校园区域巡检（教学楼 / 宿舍 / 公共区域）

- 校园安防辅助监控

- 应急事件低空快速响应

### 📊 数据可视化

- 巡检数据实时展示

- 历史巡检数据统计分析

- 异常事件可视化标记

### 🔐 系统管理

- 用户与权限管理

- 日志与巡检记录管理

- 系统运行状态监控

---

## 🛠 技术栈

| 层级  | 技术                                 |
| --- | ---------------------------------- |
| 前端  | Vue.js / TypeScript / Pinia / Vite |
| 后端  | Go                                 |
| 数据库 | PostgreSQL / Redis                 |
| API | REST / gRPC                        |
| 其他  | Docker 容器化支持                       |

---

## 🧩 快速开始

### 从源代码开始

#### 1. 克隆仓库

```bash
git clone https://github.com/OrangePanda2022/Cloud-Patrol.git
cd Cloud-Patrol
```

### 2. 后端启动

```bash
# 进入后端目录
cd tmp-backend

# 安装依赖
go mod tidy

# 启动服务
go run server.go
```

### 3. 前端启动

```bash
cd user-frontend
bun install
bun run dev
```

浏览器打开: `http://localhost:5173`

---

### 使用预编译的文件开始

#### 1. 下载<a href="https://github.com/OrangePanda2022/Cloud-Patrol/releases">Release</a>页面中的文件

#### 2. 后端启动

```bash
./server
```

#### 3. 前端启动

- 配置 Nginx
  
  ```nginx
  server {
      listen 80;
      server_name cloud-patrol.local;
  
      # 前端静态文件
      root /var/dist;
      index index.html;
  
      location / {
          try_files $uri $uri/ /index.html;
      }
  }
  ```

- 启动 Nginx
  
  ```bash
  sudo systemctl reload nginx
  ```

---

## 🧾 配置 & 部署

项目通过变量配置关键参数：

| 变量名            | 说明                                    |
| -------------- | ------------------------------------- |
| `BaseURL`      | 大模型 API 的基础访问地址                       |
| `systemPrompt` | 系统级提示词，用于定义系统的基础行为逻辑（如智能巡检策略、低空场景规则等） |
| `accessToken`  | 大模型访问令牌，用于接口鉴权与安全访问控制                 |
| `PORT`         | 后端服务监听端口                              |

部署建议：

- 使用 Docker Compose

- 配置反向代理（Nginx / Traefik）

- 使用 Cloud Provider 提供的 Load Balancer

---

## 💡 贡献指南

欢迎通过 Issue 或 Pull Request 参与改进：

- ⭐ 提交 Bug 报告

- ✨ 提出新功能建议

- 🛠 编写单元与集成测试














