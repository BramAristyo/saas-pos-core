
# Point of Sales Core

This repository contains the core components of a customized Point of Sale (POS) system developed for **Mawish Coffee**, Semarang.

## Project Overview

The Mawish POS ecosystem is a tailored solution for coffee shop operations, featuring a progressively built backend and a feature-rich frontend.

- **Backend Philosophy:** Hand-written, incrementally developed — starting lean and growing deliberately. No premature abstraction.
- **Frontend Approach:** Approximately 80% of the frontend implementation was executed using an **Agentic Gemini workflow**, enabling rapid prototyping and UI/UX consistency.

---

## Backend Architecture

The backend is written in Go by hand, built incrementally using a **Layered Architecture** pattern. The goal is to keep things straightforward early on — minimal interface abstraction, no over-engineering — and introduce complexity only when the codebase genuinely demands it.

### Design Philosophy

The architecture is intentionally pragmatic. Rather than front-loading abstractions, the codebase grows organically through deliberate iteration:

- **Layered structure from day one** — Handler → Service → Repository — each layer with a clear, single responsibility.
- **Minimal interfaces early** — Interfaces are introduced only when there's a concrete need (e.g., swapping implementations, enabling testing), not as a default pattern.
- **Hand-written everything** — No code generation tools. Every handler, query, and migration is written manually to maintain full understanding of the codebase.
- **Iterative development** — Features are built end-to-end in small cycles: define the route, write the handler, implement the service logic, wire up the repository.

### Layered Structure

```
cmd/
  api/          → Application entrypoint
  migration/    → Schema migration runner
  seeder/       → Development data seeder

internal/
  handler/      → HTTP handlers (request parsing, response formatting)
  service/      → Business logic (pure Go, no framework dependencies)
  repository/   → Data access (SQL queries via GORM)
  domain/       → Shared structs, enums, and error types
  middleware/   → Auth, logging, and request lifecycle hooks

config/         → Environment loading via Viper
```

Interfaces are added to the service and repository layers gradually — only once a module is stable enough that mocking or swapping makes sense.

### Technical Stack

- **Language:** Go 1.25.7
- **HTTP Framework:** Gin Gonic
- **Persistence Layer:** GORM with PostgreSQL
- **Database Migrations:** Goose (SQL files, written by hand)
- **Configuration Management:** Viper
- **Security:** JWT (JSON Web Tokens)

### Development Workflow

Each feature follows this incremental loop:

1. Define the route and HTTP contract.
2. Write the handler, parse input, call service, return response.
3. Implement service logic, business rules, validation, error handling.
4. Write the repository method, direct GORM query, no magic.
5. Test manually via Bruno, then refactor if needed.

This cycle keeps the codebase understandable at every stage of development.

### Key Backend Deliverables

- **Data Seeding Engine:** Located in `cmd/seeder/`, providing consistent state for development and testing environments.
- **Bruno API Documentation:** A complete collection of hand-maintained API requests for manual testing and integration verification.
- **Soft Delete & Audit Fields:** `deleted_at`, `created_at`, `updated_at` applied consistently across entities without a separate audit log service.

---

## Frontend Implementation

The frontend is a modern SPA designed for efficiency and responsiveness.

### Technical Stack

- **Framework:** Vue 3 (Composition API)
- **Build System:** Vite
- **Language:** TypeScript
- **Styling Engine:** Tailwind CSS v4
- **State Management:** Pinia

---

## API Testing (Manual)

Manual API verification is performed using [Bruno](https://www.usebruno.com/).

### Instructions

1. Install the Bruno CLI or Desktop application.
2. Open the collection located at: `server/docs/bruno-api/`.
3. Configure environment variables to match your local server settings.
4. Execute requests against the running Go server to verify endpoint behavior.

---

## Installation and Setup

### Prerequisites

- Go Runtime (1.25.7 or higher)
- Node.js Environment (v20 or higher)
- PostgreSQL Database Instance

### Backend Configuration

1. Navigate to the `server` directory.
2. Copy and configure environment: `cp .env.example .env`
3. Run migrations: `go run cmd/migration/main.go`
4. Seed initial data: `go run cmd/seeder/main.go`
5. Start the API server: `go run cmd/api/main.go`

### Frontend Configuration

1. Navigate to the `web` directory.
2. Install dependencies: `npm install`
3. Start development server: `npm run dev`

---

## Operational Commands

| Objective | Command |
| :--- | :--- |
| **API Server** | `go run cmd/api/main.go` |
| **Database Migration** | `go run cmd/migration/main.go` |
| **Data Seeding** | `go run cmd/seeder/main.go` |
| **Frontend Dev** | `npm run dev` |
| **Frontend Build** | `npm run build` |

---

## Confidentiality and License

This software and its source code are private and confidential. Unauthorized distribution or reproduction is prohibited.
