# SaaS POS Core Frontend

## Overview
This repository contains the high-performance, enterprise-grade frontend for the SaaS POS Core ecosystem. Built with Vue 3 and TypeScript, the application is engineered for reliability, scalability, and rapid response times in demanding retail and hospitality environments. It serves as the primary interface for point-of-sale operations, inventory management, and administrative configurations.

## Tech Stack Deep-Dive

### Core Framework
- **Vue 3 (Composition API)**: Leverages the latest reactive paradigms for optimal logic reuse and a clean code architecture.
- **TypeScript**: Implements strict type-safety across the entire codebase to minimize runtime errors and enhance developer productivity through robust IDE support.

### Build and Performance
- **Vite**: Utilized as the build tool and development server, providing ultra-fast Hot Module Replacement (HMR) and highly optimized production builds with efficient code splitting and tree shaking.

### State Management
- **Pinia**: Implements a modular, "store-per-module" architecture. Stores are defined using the setup-store syntax, ensuring consistency with the Composition API and providing a centralized, reactive state for authentication, sales, and configuration data.

### UI/UX Engineering
- **Tailwind CSS**: Utility-first styling framework used for rapid UI development and maintaining a consistent design system across the application.
- **Shadcn Vue**: Accessible, unstyled UI components built on top of Radix UI primitives (via Reka UI), allowing for full design control while adhering to strict accessibility standards (WAI-ARIA).
- **Lucide Vue Next**: A comprehensive library of high-quality, consistent iconography tailored for the POS interface.

### Networking
- **Axios**: Centralized HTTP client configuration located in `src/lib/http.ts`. It features strict TypeScript interceptors for automated Bearer token attachment and synchronized session expiration (401) handling.
- **API Layer**: Decoupled API services ensure that business logic remains independent of the transport layer, facilitating easier testing and maintenance.

## AI Agentic Development
This project utilizes an advanced Agentic Workflow for development, integrating the Gemini CLI to orchestrate complex engineering tasks directly from the terminal.

### Orchestration System
- **GEMINI.md**: Serves as the central project context and instruction persistence layer. It provides the orchestrator with the necessary architectural guidelines, naming conventions, and global rules.
- **Custom Sub-Agents**: Specialized agents are employed to handle domain-specific tasks, ensuring that all code adheres to the project's high standards:
    - **API Integration Agent**: Manages TypeScript definitions and the implementation of service layers.
    - **UI Component Generator**: Scaffolds views and components based on established design tokens and layout patterns.
    - **State Management Agent**: Oversees the creation and modification of Pinia stores and state logic.
    - **Error Handler Agent**: Implements robust error boundaries, try-catch blocks, and user notification logic.
    - **Form Validator Agent**: Orchestrates complex form states, validation schemas, and error messaging.

This agentic approach ensures architectural consistency and significantly accelerates the development lifecycle while maintaining a clean, professional codebase.

## Getting Started

### Prerequisites
- Node.js (Version 20.19.0 or 22.12.0+)
- npm

### Installation
```bash
npm install
```

### Development
Launch the development server with HMR:
```bash
npm run dev
```

### Production Build
Compile and minify the project for production:
```bash
npm run build
```

### Static Analysis and Linting
Run the full linting suite:
```bash
npm run lint
```

## Engineering Standards

### Architecture Principles
- **Clean Architecture**: Strict separation of concerns between views, stores, and API layers.
- **Single Source of Truth**: Centralized state management via Pinia and typed interfaces in `src/types/`.
- **Mobile-First Design**: Responsive layouts engineered to perform across various hardware profiles, from tablets to desktop monitors.

### Quality Assurance
- **Oxlint**: Utilized for high-performance linting to provide near-instant feedback during development.
- **ESLint**: Enforces standardized code quality rules and best practices.
- **vue-tsc**: Performs strict type-checking on both TypeScript files and Vue templates to ensure total type safety.
- **No Code Comments**: The project prioritizes self-documenting code through descriptive naming and clear structural patterns.
