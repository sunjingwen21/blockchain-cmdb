# Blockchain CMDB

基于区块链技术的配置管理数据库（Configuration Management Database）平台。

## 🏗️ 架构

```
┌─────────────────┐     ┌──────────────────┐     ┌─────────────────┐
│   React Frontend │────▶│   Go Backend API │────▶│   PostgreSQL    │
│   (TypeScript)   │     │   (Gin + GORM)   │     │   (主数据库)      │
└─────────────────┘     └──────────────────┘     └─────────────────┘
                               │
                               ▼
                        ┌──────────────────┐
                        │   Blockchain     │
                        │   (Ethereum)     │
                        └──────────────────┘
```

## 🚀 技术栈

- **后端**: Go + Gin + GORM
- **前端**: React 18 + TypeScript + Ant Design
- **区块链**: Ethereum (ethers.js)
- **数据库**: PostgreSQL + Redis
- **部署**: Docker + Docker Compose

## 📁 项目结构

```
blockchain-cmdb/
├── backend/           # Go 后端服务
│   ├── api/          # RESTful API 路由
│   ├── blockchain/   # 区块链交互层
│   ├── models/       # 数据模型
│   ├── tests/        # 单元测试
│   └── config/       # 配置文件
├── frontend/         # React 前端
│   ├── src/
│   │   ├── components/  # UI组件
│   │   ├── pages/       # 页面
│   │   ├── utils/       # 工具函数
│   │   └── styles/      # 样式
│   ├── public/       # 静态资源
│   └── tests/        # 测试文件
├── docker/           # Docker配置
├── scripts/          # 部署脚本
├── docs/             # 文档
├── CHANGELOG.md      # 变更日志
└── README.md         # 项目说明
```

## 🛠️ 开发

### 后端启动
```bash
cd backend
go mod init blockchain-cmdb
go get -u github.com/gin-gonic/gin gorm.io/gorm gorm.io/driver/postgres
```

### 前端启动
```bash
cd frontend
npx create-react-app . --template typescript
npm install antd ethers
```

### Docker 启动
```bash
docker-compose up -d
```

## 📜 许可证

MIT License
