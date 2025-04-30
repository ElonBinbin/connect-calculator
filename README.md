# 最终面试挑战

本仓库包含两个面试挑战任务的完整实现。

## 任务一：美国县级人口变动数据分析
- 详细分析报告位于 [task1/README.md](task1/README.md)
- 包含数据来源、分析方法和主要发现

## 任务二：Connect Calculator 实现
一个基于 ConnectRPC 的全栈计算器应用。

### 目录结构
- backend/ - Go 实现的 ConnectRPC 后端
- frontend/ - Next.js 实现的前端页面
- task1/ - 任务一的数据分析报告和相关资料

### 启动方式

#### 后端
```bash
cd backend
go run main.go
```

#### 前端
```bash
cd frontend
npm install
npm run dev

### Envyo代理
envoy -c envoy.yaml

```
###  测试后端
```bash
cd backend
go run server_test.go