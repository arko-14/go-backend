# Technical Decisions & Architectural Reasoning

This document outlines the architectural decisions, library choices, and design patterns used in the development of the User Management API.

## 1. Architectural Pattern: Clean Architecture

I chose to move away from a "flat" structure to a modular **Layered Architecture**.

- **`cmd/`**: Entry point. Keeps application initialization separate from logic.
- **`internal/handler`**: Handles HTTP transport (JSON parsing, status codes). It knows *nothing* about the database.
- **`internal/service`**: Contains business logic (Age calculation). It bridges the handler and the database.
- **`db/sqlc`**: The data access layer.

**Why?** This ensures separation of concerns. If we switch the database later, the Handler code doesn't change.

## 2. Containerization & Deployment (Docker)

I implemented full Docker support to ensure the application works consistently across different environments ("works on my machine" vs production).

- **Multi-Stage Build:** I used a 2-stage Dockerfile (`builder` -> `alpine`). This compiles the Go binary in a heavy image but deploys it in a tiny Alpine Linux image, keeping the final container size small (~15MB).
- **Orchestration:** `docker-compose` is used to spin up both the **Application** and **Postgres** simultaneously, handling networking automatically.
- **Dual-Mode Configuration:** I modified `main.go` to check for the `DB_SOURCE` environment variable.
  - If present (Docker context), it uses the container network string.
  - If missing (Local context), it gracefully falls back to local `localhost` credentials.
  - This allows developers to run the app using `go run` OR `docker-compose` without changing code.

## 3. Library Choices

### SQLC vs GORM

I chose **SQLC** over a full ORM like GORM.

- **Type Safety:** SQLC generates Go code from raw SQL queries, catching syntax errors at compile time.
- **Performance:** It avoids runtime reflection, making it significantly faster.

### GoFiber

Selected for its performance (based on `fasthttp`) and Express.js-like routing syntax, which simplifies the handler layer.

## 4. Design Decisions

### Dynamic Age Calculation

**Decision:** The `age` field is **not** stored in the database.

**Reasoning:** Age changes automatically over time. Storing it would require daily background updates.

**Solution:** Store `dob` (immutable) and calculate `age` on-the-fly in the Service layer.

### Input Validation

Used `go-playground/validator` to enforce rules at the struct level (e.g., `validate:"required"`). This keeps the handler code DRY and declarative.

## 5. Future Improvements

- **Unit Tests:** Add comprehensive test coverage for the `CalculateAge` service.
- **CI/CD:** Add a GitHub Actions workflow to run linting and build checks on push.
- **Configuration Management:** Use a tool like `viper` to manage complex configurations beyond simple environment variables.

---

This architecture balances performance, maintainability, and scalability while keeping the codebase clean and testable.
