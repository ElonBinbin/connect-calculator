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
```
```

现在，您可以按以下步骤提交这些更改：

1. 创建 task1 目录并添加文件：
```bash
mkdir -p task1
# 创建上述 task1/README.md 文件内容
# 更新主 README.md 文件内容

# 添加更改到 Git
git add task1/README.md
git add README.md
git commit -m "任务一：添加美国县级人口变动数据分析报告"
git push origin main  # 或者您的默认分支名
```

您要我帮您执行这些 Git 命令吗？或者您想自己执行？请注意，在执行 `git push` 之前，请确保您已经设置好了远程仓库。