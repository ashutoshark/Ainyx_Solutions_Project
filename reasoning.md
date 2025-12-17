# Reasoning & Approach

## Project Overview

This is a RESTful API built with Go to manage users with name and date of birth. The API calculates age dynamically when fetching user details.

## Why I Chose These Technologies

### Go + Fiber
- Fiber is fast and easy to learn
- Syntax is similar to Express.js so its beginner friendly
- Good documentation available

### PostgreSQL
- Most popular relational database
- Good for storing structured data like users
- Free and open source

### SQLC
- Generates type-safe Go code from SQL queries
- No need to write boilerplate database code
- Compile time error checking

## Folder Structure Decision

```
├── cmd/server/          # entry point
├── config/              # environment config
├── db/
│   ├── queries/         # SQL queries
│   └── sqlc/            # generated code
└── internal/
    ├── handler/         # HTTP handlers
    ├── models/          # request/response structs
    ├── repository/      # database operations
    ├── routes/          # URL routes
    └── service/         # business logic
```

I used layered architecture because:
1. **Separation of concerns** - each layer has one job
2. **Easy to test** - can mock dependencies
3. **Easy to understand** - clear flow of data

## Key Decisions

### 1. Dynamic Age Calculation
Instead of storing age in database (which becomes outdated), I calculate it on-the-fly using DOB. This ensures age is always accurate.

```go
func CalculateAge(dob time.Time) int {
    now := time.Now()
    age := now.Year() - dob.Year()
    // adjust if birthday hasn't happened yet
    if now.Month() < dob.Month() {
        age--
    }
    if now.Month() == dob.Month() && now.Day() < dob.Day() {
        age--
    }
    return age
}
```

### 2. Layered Architecture
- **Handler** → receives HTTP request, validates input
- **Service** → contains business logic, processes data
- **Repository** → talks to database

This makes code organized and testable.

### 3. Input Validation
Used `go-playground/validator` for validating request body. Checks if name and dob are provided before processing.

### 4. Error Handling
Each layer handles errors appropriately:
- Repository returns database errors
- Service converts to user-friendly messages
- Handler returns proper HTTP status codes

## What I Would Add With More Time

1. **Authentication** - JWT tokens for secure access
2. **Pagination** - for listing large number of users
3. **Logging** - structured logging with request IDs
4. **Tests** - more unit tests and integration tests
5. **Docker** - containerization for easy deployment

## Challenges Faced

1. **Date parsing** - had to handle different date formats
2. **Age calculation edge cases** - leap years, birthday not yet happened
3. **Null handling** - making sure empty arrays don't return null in JSON

## Conclusion

This project demonstrates basic CRUD operations with clean architecture. The code is simple and readable, suitable for a beginner-level Go project. The dynamic age calculation is the main feature that shows understanding of business logic implementation.
