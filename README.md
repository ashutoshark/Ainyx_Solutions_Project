# Ainyx Solutions Project - User API

A RESTful API built with Go (Fiber) + PostgreSQL + SQLC to manage users with dynamic age calculation.

---

## 🛠 Tech Stack

- **Go** - Programming language
- **Fiber** - HTTP framework
- **PostgreSQL** - Database
- **SQLC** - SQL code generator
- **Validator** - Input validation

---

## 📁 Project Structure

```
├── cmd/server/main.go          → Entry point
├── config/config.go            → Environment config
├── db/
│   ├── migrations/             → SQL schema
│   └── sqlc/                   → Generated DB code
├── internal/
│   ├── handler/                → HTTP handlers
│   ├── service/                → Business logic
│   ├── repository/             → Database layer
│   ├── routes/                 → Route definitions
│   └── models/                 → Request/Response structs
├── go.mod
├── README.md
├── EXPLANATION.md              → Interview Q&A
└── reasoning.md                → Approach & decisions
```

---

## 🚀 How to Run

### Prerequisites
- Go 1.21+
- PostgreSQL

### Step 1: Create database
```sql
CREATE DATABASE userdb;

\c userdb

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    dob DATE NOT NULL
);
```

### Step 2: Set environment variables
```bash
export DATABASE_URL="postgres://postgres:yourpassword@localhost:5432/userdb?sslmode=disable"
export PORT=3000
```

### Step 3: Run
```bash
go run cmd/server/main.go
```

Server starts at: **http://localhost:3000**

---

## 📡 API Endpoints

| Method | URL | Description |
|--------|-----|-------------|
| POST | /users | Create user |
| GET | /users | List all users (with age) |
| GET | /users/:id | Get user by ID (with age) |
| PUT | /users/:id | Update user |
| DELETE | /users/:id | Delete user |

---

## 📮 API Examples

### Create User
```http
POST /users
Content-Type: application/json

{
    "name": "Ashutosh",
    "dob": "2000-05-15"
}
```

**Response:**
```json
{
    "id": 1,
    "name": "Ashutosh",
    "dob": "2000-05-15"
}
```

### Get User (age calculated dynamically)
```http
GET /users/1
```

**Response:**
```json
{
    "id": 1,
    "name": "Ashutosh",
    "dob": "2000-05-15",
    "age": 24
}
```

### Update User
```http
PUT /users/1
Content-Type: application/json

{
    "name": "Ashutosh Kumar",
    "dob": "2000-05-15"
}
```

### Delete User
```http
DELETE /users/1
```

---

## 🧪 Run Tests
```bash
go test ./...
```

---

## 📚 Documentation


- **[reasoning.md](reasoning.md)** - Approach and key decisions

---

## 👤 Author

**Ashutosh** - [GitHub](https://github.com/ashutoshark)
