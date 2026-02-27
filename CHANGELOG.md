# CHANGELOG

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added (Day 2 - 2026-02-27)
- ✅ PostgreSQL database integration with GORM
  - Database connection management
  - Auto-migration for User, Asset, AssetHistory models
  - Connection pooling and graceful shutdown
- ✅ JWT authentication system
  - User login with email/password
  - User registration with role assignment
  - JWT token generation and validation
  - Password change functionality
  - Protected route middleware
  - Role-based access control (admin/user)
- ✅ Blockchain connection module (Ethereum)
  - Ethereum client initialization
  - Chain ID detection
  - Balance checking
  - Transaction status tracking
  - Asset registration placeholder (smart contract ready)
  - Key pair generation
- ✅ Authentication API endpoints
  - POST /api/v1/auth/login
  - POST /api/v1/auth/register
  - GET /api/v1/auth/me (protected)
  - POST /api/v1/auth/change-password (protected)
- ✅ Enhanced main.go with all integrations
  - Database initialization
  - Blockchain client setup
  - JWT middleware integration
  - Protected route groups

### Planned (Day 3)
- [ ] Frontend Ant Design integration
- [ ] Login page UI
- [ ] Asset dashboard page
- [ ] Asset CRUD with database persistence
- [ ] Docker Compose setup with PostgreSQL
- [ ] Environment configuration template

### Added (Day 1 - 2026-02-26)
- ✅ Backend API foundation with Go + Gin framework
  - RESTful API structure with versioning (/api/v1)
  - Health check endpoint
  - Configuration management (environment-based)
  - User CRUD operations (List, Get, Create, Update, Delete)
  - Asset CRUD operations with history tracking
- ✅ Frontend React + TypeScript project structure
  - Type definitions for User and Asset
  - API service layer with Axios
  - Unit test structure with Jest
- ✅ Database Models (base structure)
  - User model with validation
  - Asset model with status tracking and history
- ✅ Unit Tests
  - Handler tests for API endpoints
  - Model validation tests
  - Config loading tests
- ✅ Git repository setup with SSH key authentication

### Planned (Future)
- [ ] Blockchain asset registration/on-chain storage
- [ ] Asset transfer history on blockchain
- [ ] Frontend asset management interface
- [ ] Real-time asset monitoring
- [ ] CI/CD pipeline with GitHub Actions
- [ ] API documentation (Swagger/OpenAPI)

## Technical Decisions

### Backend: Go + Gin
- **Rationale**: High performance, low memory footprint, excellent concurrency support
- **Alternative considered**: Node.js/NestJS (faster development but higher resource usage)

### Frontend: React + TypeScript + Ant Design
- **Rationale**: Type safety, mature ecosystem, enterprise-grade UI components
- **Alternative considered**: Vue 3 (simpler but less enterprise adoption)

### Blockchain: Ethereum/Polygon
- **Rationale**: Mature ecosystem, wide adoption, cost-effective on Polygon
- **Alternative considered**: Solana (faster but less mature tooling)

### Database: PostgreSQL + Redis
- **Rationale**: ACID compliance, JSON support, proven reliability
- **Redis**: Caching, session storage, rate limiting

## Known Issues
- None yet (project just initialized)

## Security Considerations
- All database credentials must use environment variables
- JWT secrets must be rotated regularly
- Blockchain private keys must be stored in secure vault (e.g., AWS KMS, HashiCorp Vault)
- API rate limiting to be implemented
- Input validation on all endpoints

