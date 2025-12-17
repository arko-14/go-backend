# Technical Decisions & Architectural Reasoning

This document outlines the architectural decisions, library choices, and design patterns used in the development of the User Management API.

## 1. Architectural Pattern: Clean Architecture

I chose to move away from a "flat" structure (putting everything in `main.go`) to a modular **Layered Architecture**.

- **`cmd/`**: Entry point. Keeps the application initialization separate from logic.
- **`internal/handler`**: Handles HTTP transport (JSON parsing, status codes). It knows *nothing* about the database.
- **`internal/service`**: Contains business logic (Age calculation). It bridges the handler and the database.
- **`db/sqlc`**: The data access layer.

**Why?** This ensures separation of concerns. If we switch the database later, the Handler code doesn't change. If we switch from HTTP to gRPC, the Logic code doesn't change.

## 2. Library Choices

### SQLC vs GORM

I chose **SQLC** over a full ORM like GORM.

- **Type Safety:** SQLC generates Go code from raw SQL queries, catching SQL syntax errors at compile time rather than runtime.
- **Performance:** ORMs use `reflection`, which is slower. SQLC generates raw, optimized code that runs as fast as handwriting it.
- **Control:** It allows full control over the SQL queries (e.g., using `RETURNING` clauses efficiently).

### GoFiber

I selected **Fiber** for the HTTP framework.

- It is built on top of `fasthttp`, offering superior performance compared to the standard `net/http`.
- The API is Express.js-like, making the routing code clean and readable.

### Uber Zap

I used **Zap** for logging instead of the standard `log` package.

- It provides structured logging (JSON), which is essential for parsing logs in production tools (like Datadog or ELK Stack).
- It is zero-allocation and highly performant.

## 3. Design Decisions

### Dynamic Age Calculation

**Decision:** The `age` field is **not** stored in the database.

**Reasoning:** Age is a volatile data point that changes automatically over time. Storing it would require a daily background job (cron) to update every record in the database.

**Solution:** Store `dob` (immutable) and calculate `age` on-the-fly in the Service layer when the data is requested. This ensures 100% data accuracy with zero maintenance.

### Input Validation

I implemented the `go-playground/validator` package to enforce rules at the struct level (e.g., `validate:"required,datetime=2006-01-02"`).

**Reasoning:** This avoids cluttering the handler code with manual `if req.Name == ""` checks, keeping the code DRY (Don't Repeat Yourself) and declarative.

## 4. Future Improvements

Given more time, I would improve the following:

- **Configuration:** Move the database connection string to an `.env` file using `viper` or `godotenv` to avoid hardcoding secrets.
- **Docker:** Add a `Dockerfile` and `docker-compose.yml` to spin up the App and Postgres in one command.
- **Unit Tests:** Add tests for the `CalculateAge` function to handle edge cases (e.g., leap years, timezone considerations).
- **API Documentation:** Integrate Swagger/OpenAPI for interactive API documentation.
- **Rate Limiting:** Implement rate limiting middleware to prevent API abuse.
- **Pagination:** Add pagination support for the list users endpoint to handle large datasets efficiently.

## 5. Performance Considerations

### Why GoFiber over Standard net/http?

GoFiber is built on top of `fasthttp`, which uses a different approach to handling HTTP requests:

- **Zero Memory Allocation:** Reuses objects instead of creating new ones for each request
- **Faster Routing:** Uses a radix tree for route matching
- **Benchmarks:** Consistently shows 2-3x better performance than standard library in high-load scenarios

### Database Connection Pooling

The `pgx` driver provides built-in connection pooling, which:

- Reuses database connections instead of creating new ones for each query
- Reduces latency and database load
- Automatically handles connection lifecycle management

## 6. Security Considerations

### SQL Injection Prevention

By using SQLC with parameterized queries, we automatically prevent SQL injection attacks. All user inputs are properly escaped and sanitized.

### Input Validation

The validator package ensures that:

- Required fields are always present
- Date formats are strictly validated
- Invalid data is rejected before reaching the database layer

---

This architecture balances performance, maintainability, and scalability while keeping the codebase clean and testable.
