# Go Backend Task - User Management API

A high-performance RESTful API built with **Go (Golang)**, designed to manage user records and dynamically calculate age based on Date of Birth (DOB). This project adheres to a clean architecture, separating concerns between HTTP handlers, business logic, and database interactions.

## 🚀 Tech Stack

| Component | Technology | Reasoning |
|:----------|:-----------|:----------|
| **Framework** | [GoFiber v2](https://gofiber.io/) | Express-inspired, zero memory allocation, blazing fast. |
| **Database** | PostgreSQL | Robust, relational data integrity. |
| **ORM / SQL** | [SQLC](https://sqlc.dev/) | Compiles SQL to type-safe Go code (no runtime reflection). |
| **Driver** | [pgx/v5](https://github.com/jackc/pgx) | High-performance PostgreSQL driver and toolkit. |
| **Logging** | [Uber Zap](https://github.com/uber-go/zap) | Structured, leveled logging for production environments. |
| **Validation** | [go-playground/validator](https://github.com/go-playground/validator) | Declarative input validation using struct tags. |

## 📂 Project Structure

This project follows the [Standard Go Project Layout](https://github.com/golang-standards/project-layout):

```text
go-backend-task/
├── cmd/server/          # Entry point (main.go)
├── db/
│   ├── sqlc/            # Auto-generated Go code for database access
│   └── query.sql        # Raw SQL queries used by SQLC
├── internal/
│   ├── handler/         # HTTP Layer (Parse request, send response)
│   ├── service/         # Business Logic (Age calculation, data transformation)
│   └── models/          # Data Structures (Structs)
├── go.mod               # Dependencies
├── README.md            # Project documentation
└── reasoning.md         # Technical decisions and architecture rationale
```

## 🛠️ Setup & Installation

### 1. Prerequisites

- Go 1.23 or higher
- PostgreSQL installed and running locally

### 2. Database Setup

Run the following SQL commands to create the database and schema:

```sql
CREATE DATABASE users_db;

-- Connect to the database
\c users_db

-- Create the table
CREATE TABLE users (
  id BIGSERIAL PRIMARY KEY,
  name TEXT NOT NULL,
  dob DATE NOT NULL
);
```

### 3. Configuration

Open `cmd/server/main.go` and update the connection string with your PostgreSQL password:

```go
const dbSource = "postgresql://postgres:YOUR_PASSWORD@localhost:5432/users_db?sslmode=disable"
```

### 4. Run the Application

```bash
# Download dependencies
go mod tidy

# Run the server
go run cmd/server/main.go
```

The server will start on port 3000.

## 🧪 API Endpoints

### 1. Create User

- **Endpoint:** `POST /users`
- **Body:** `{"name": "Alice", "dob": "1990-05-10"}`
- **Returns:** Created user object

### 2. Get User (with Age)

- **Endpoint:** `GET /users/:id`
- **Returns:** `{"id": 1, "name": "Alice", "dob": "1990-05-10", "age": 34}`
- **Note:** The `age` field is calculated dynamically on every request

### 3. List All Users

- **Endpoint:** `GET /users`
- **Returns:** Array of user objects

### 4. Update User

- **Endpoint:** `PUT /users/:id`
- **Body:** `{"name": "Alice Updated", "dob": "1992-01-01"}`

### 5. Delete User

- **Endpoint:** `DELETE /users/:id`
- **Returns:** `204 No Content`

## 📖 Documentation

For detailed information about technical decisions and architectural choices, see [reasoning.md](./reasoning.md).

## 👤 Author

**SP**

---

Built with ❤️ using Go
