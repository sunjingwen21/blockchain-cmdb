# Blockchain CMDB Platform

A production-grade Configuration Management Database (CMDB) platform with blockchain integration for immutable audit trails.

## 🏗️ Architecture

```
┌─────────────────────────────────────────────────────────────┐
│                      Frontend (React + TS)                   │
│  ┌─────────────┐  ┌─────────────┐  ┌─────────────────────┐  │
│  │  Dashboard  │  │  Asset Mgmt │  │   Blockchain View   │  │
│  └─────────────┘  └─────────────┘  └─────────────────────┘  │
└─────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────┐
│                    Backend API (Go + Gin)                    │
│  ┌─────────────┐  ┌─────────────┐  ┌─────────────────────┐  │
│  │  REST API   │  │  Auth/JWT   │  │  Business Logic     │  │
│  └─────────────┘  └─────────────┘  └─────────────────────┘  │
└─────────────────────────────────────────────────────────────┘
                              │
              ┌───────────────┼───────────────┐
              ▼               ▼               ▼
┌─────────────────┐  ┌─────────────────┐  ┌─────────────────┐
│   PostgreSQL    │  │     Redis       │  │   Blockchain    │
│   (Primary DB)  │  │   (Cache/Queue) │  │   (Audit Log)   │
└─────────────────┘  └─────────────────┘  └─────────────────┘
```

## 🚀 Tech Stack

### Backend
- **Language**: Go 1.21+
- **Framework**: Gin
- **Database**: PostgreSQL 15
- **Cache**: Redis 7
- **Blockchain**: Ethereum/Polygon SDK
- **Authentication**: JWT + bcrypt

### Frontend
- **Framework**: React 18
- **Language**: TypeScript 5
- **UI Library**: Ant Design 5
- **State Management**: Zustand
- **HTTP Client**: Axios

## 📁 Project Structure

```
blockchain-cmdb/
├── backend/                 # Go backend service
│   ├── api/                # REST API handlers
│   ├── blockchain/         # Blockchain integration
│   ├── models/             # Database models
│   ├── middleware/         # Auth, logging, cors
│   ├── config/             # Configuration
│   ├── utils/              # Utilities
│   └── main.go             # Entry point
├── frontend/               # React frontend
│   ├── src/
│   │   ├── components/     # Reusable UI components
│   │   ├── pages/          # Page components
│   │   ├── hooks/          # Custom React hooks
│   │   ├── services/       # API services
│   │   ├── utils/          # Utilities
│   │   └── types/          # TypeScript types
│   └── package.json
├── docs/                   # Documentation
├── docker/                 # Docker configurations
├── scripts/                # Deployment scripts
└── CHANGELOG.md            # Change log
```

## 🛠️ Development

### Prerequisites
- Go 1.21+
- Node.js 18+
- PostgreSQL 15+
- Redis 7+
- Docker (optional)

### Quick Start

```bash
# Backend
cd backend
go mod init github.com/sunjingwen21/blockchain-cmdb
go get -u github.com/gin-gonic/gin
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgresql
go run main.go

# Frontend
cd frontend
npm install
npm start
```

## 📝 License

MIT License
