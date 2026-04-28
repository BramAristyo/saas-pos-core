# Point of Sales Core

This repository contains the core components of the Mawish Software-as-a-Service (SaaS) Point of Sale (POS) system. The project is engineered with a primary focus on backend stability and scalability using Go, complemented by a modern web interface built with Vue.

## Project Overview

The Mawish POS ecosystem is divided into two primary modules:
- **Server:** A high-performance RESTful API implementation developed in Go, following strict architectural standards.
- **Web:** A responsive administrative dashboard built with Vue 3 and Tailwind CSS. The frontend development process was significantly accelerated through collaborative implementation with Gemini Agentic.

---

## Backend Architecture

The backend serves as the foundation of the system, designed for high concurrency, data integrity, and long-term maintainability.

### Technical Stack
- **Language:** Go 1.25.7
- **HTTP Framework:** Gin Gonic
- **Persistence Layer:** GORM with PostgreSQL
- **Database Migrations:** Goose
- **Data Validation:** Go-Playground Validator v10
- **Configuration Management:** Viper
- **Security:** JWT (JSON Web Tokens) for stateless authentication

### Architectural Principles
The server implementation strictly adheres to Clean Architecture to ensure separation of concerns and testability:
- `internal/domain/`: Definitions of core entities and interface contracts.
- `internal/usecase/`: Implementation of pure business logic and rules.
- `internal/repository/`: Data persistence logic using GORM implementation.
- `internal/api/`: Request handling, DTO mapping, and route definitions.
- `internal/infrastructure/`: Low-level configurations, database drivers, and logging.
- `internal/dependency/`: Orchestration of dependency injection and component lifecycle.

### Key Implementation Features
- **Standardized Soft Delete:** Consistent application-wide data retention using GORM's `DeletedAt`.
- **Comprehensive Audit Logging:** Automated tracking of state transitions and data modifications.
- **Dynamic Querying:** Shared utilities for robust filtering, sorting, and pagination.
- **Transactional Integrity:** Explicit transaction management within the repository layer for complex operations.

---

## Frontend Implementation

The web interface is designed to provide an efficient user experience for POS management. To optimize development velocity, the UI/UX components and frontend business logic were developed in coordination with Gemini Agentic.

### Technical Stack
- **Framework:** Vue 3 (Composition API)
- **Build System:** Vite
- **Programming Language:** TypeScript
- **Styling Engine:** Tailwind CSS v4
- **State Management:** Pinia
- **Data Visualization:** TanStack Table (Vue Table)

---

## Installation and Setup

### Prerequisites
- Go Runtime (1.25.7 or higher)
- Node.js Environment (v20 or higher)
- PostgreSQL Database Instance

### Backend Configuration
1. Access the server directory:
   ```bash
   cd server
   ```
2. Initialize environment configuration:
   ```bash
   cp .env.example .env
   # Update .env with appropriate PostgreSQL connection strings
   ```
3. Execute database migrations:
   ```bash
   go run cmd/migration/main.go
   ```
4. (Optional) Initialize seed data:
   ```bash
   go run cmd/seeder/main.go
   ```
5. Launch the API service:
   ```bash
   go run cmd/api/main.go
   ```

### Frontend Configuration
1. Access the web directory:
   ```bash
   cd web
   ```
2. Install required packages:
   ```bash
   npm install
   ```
3. Start the development environment:
   ```bash
   npm run dev
   ```

---

## Operational Commands

| Objective | Command |
| :--- | :--- |
| **API Server Execution** | `go run cmd/api/main.go` |
| **Database Migration** | `go run cmd/migration/main.go` |
| **Data Seeding** | `go run cmd/seeder/main.go` |
| **Frontend Development** | `npm run dev` |
| **Frontend Production Build** | `npm run build` |

---

## Confidentiality and License
This software and its source code are private and confidential. Unauthorized distribution or reproduction is prohibited.
