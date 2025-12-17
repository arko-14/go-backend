# Go Backend Task - User Management API

A high-performance RESTful API built with **Go (Golang)**, designed to manage user records and dynamically calculate age based on Date of Birth (DOB). This project adheres to a clean architecture and supports both local execution and **Dockerized** deployment.

## 🚀 Tech Stack

| Component | Technology | Reasoning |
|:----------|:-----------|:----------|
| **Framework** | [GoFiber v2](https://gofiber.io/) | Express-inspired, zero memory allocation, blazing fast. |
| **Database** | PostgreSQL | Robust, relational data integrity. |
| **ORM / SQL** | [SQLC](https://sqlc.dev/) | Compiles SQL to type-safe Go code (no runtime reflection). |
| **Driver** | [pgx/v5](https://github.com/jackc/pgx) | High-performance PostgreSQL driver and toolkit. |
| **Logging** | [Uber Zap](https://github.com/uber-go/zap) | Structured, leveled logging for production environments. |
| **Validation** | [validator](https://github.com/go-playground/validator) | Declarative input validation using struct tags. |

## 📂 Project Structure

```text
go-backend-task/
├── cmd/server/          # Entry point (main.go)
├── db/sqlc/             # Auto-generated Go code for database access
├── internal/
│   ├── handler/         # HTTP Layer (Parse request, send response)
│   ├── service/         # Business Logic (Age calculation)
│   └── models/          # Data Structures
├── Dockerfile           # Docker build instructions
├── docker-compose.yml   # Container orchestration
└── README.md            # Documentation
```

## 🛠️ How to Run (Choose One)

### Method 1: Docker (Recommended 🐳)

This sets up both the API and the Database automatically.

**Start the Application:**

```bash
docker-compose up --build
```

The server will start on port 3000.

**Initialize the Database:** (Run this in a new terminal window once the containers are running)

```bash
docker exec -it postgres_container psql -U postgres -d users_db -c "CREATE TABLE users (id BIGSERIAL PRIMARY KEY, name TEXT NOT NULL, dob DATE NOT NULL);"
```

### Method 2: Manual Setup (Local)

**Prerequisites:** Ensure Go 1.23+ and PostgreSQL are installed.

**Database Setup:**

```sql
CREATE DATABASE users_db;
\c users_db
CREATE TABLE users (id BIGSERIAL PRIMARY KEY, name TEXT NOT NULL, dob DATE NOT NULL);
```

**Run the Server:**

```bash
go mod tidy
go run cmd/server/main.go
```

**Note:** The application automatically detects if it's running locally and uses the default connection string.

## 🧪 API Endpoints

| Method | Endpoint | Description | Example Body |
|:-------|:---------|:------------|:-------------|
| POST | `/users` | Create User | `{"name": "Alice", "dob": "1990-05-10"}` |
| GET | `/users/:id` | Get User | Returns user with calculated Age |
| GET | `/users` | List Users | Returns array of users |
| PUT | `/users/:id` | Update User | `{"name": "Alice Updated", "dob": "1995-01-01"}` |
| DELETE | `/users/:id` | Delete User | Returns 204 No Content |

### Example Request (cURL)

```bash
curl -X POST http://localhost:3000/users \
-H "Content-Type: application/json" \
-d '{"name": "Alice", "dob": "1990-05-10"}'
```

## 📝 Technical Decisions

- **Dual-Mode Configuration:** The application checks for a `DB_SOURCE` environment variable. If present (Docker), it uses that; otherwise, it falls back to local credentials.

- **Dynamic Age Calculation:** Age is calculated on-the-fly in the Service layer to ensure accuracy without needing daily database updates.

- **Clean Architecture:** Separated handler (HTTP) from service (Logic) to allow easier testing and future scalability.

## 📖 Documentation

For detailed information about technical decisions and architectural choices, see [reasoning.md](./reasoning.md).

## 👤 Author

**SP**

---

Built with ❤️ using Go
