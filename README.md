# User API - Go REST API Project

A simple REST API built with Go to manage users with automatic age calculation.

---

## 📁 Project Structure

```
user-api/
├── cmd/server/main.go          → Entry point (starts the app)
├── config/config.go            → App settings
├── db/
│   ├── migrations/             → SQL to create table
│   └── sqlc/                   → Database functions
├── internal/
│   ├── handler/                → HTTP request handlers
│   ├── service/                → Business logic
│   ├── repository/             → Database operations
│   ├── routes/                 → URL routes
│   └── models/                 → Data structures
├── go.mod                      → Go dependencies
├── README.md                   → This file
└── EXPLANATION.md              → Detailed explanation + Interview Q&A
```

---

## 🚀 How to Run

### Step 1: Make sure PostgreSQL is running

### Step 2: Create database and table
```sql
CREATE DATABASE userdb;

\c userdb

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    dob DATE NOT NULL
);
```

### Step 3: Run the server
```bash
go run cmd/server/main.go
```

Server runs at: **http://localhost:3000**

---

## 📡 API Endpoints

| Method | URL | Description |
|--------|-----|-------------|
| POST | /users | Create a new user |
| GET | /users | Get all users |
| GET | /users/:id | Get one user by ID |
| PUT | /users/:id | Update a user |
| DELETE | /users/:id | Delete a user |

---

## 📮 Test with Postman

### Create User
```
POST http://localhost:3000/users
Body (JSON):
{
    "name": "Alice",
    "dob": "1990-05-10"
}
```

### Get All Users (returns age!)
```
GET http://localhost:3000/users
```

### Get One User
```
GET http://localhost:3000/users/1
```

### Update User
```
PUT http://localhost:3000/users/1
Body (JSON):
{
    "name": "Alice Updated",
    "dob": "1991-03-15"
}
```

### Delete User
```
DELETE http://localhost:3000/users/1
```

---

## 🧪 Run Tests
```bash
go test ./...
```

---

## 📚 Learn More

See **EXPLANATION.md** for:
- How each file works
- Why we made certain decisions
- Interview questions and answers
