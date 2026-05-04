# Employee Management System

A scalable Go-based employee management system with a layered architecture, featuring advanced searching, concurrent exports, and robust middleware.

## Features
- **CRUD Operations**: Complete management for Employees and Departments.
- **Advanced Search**: Keyword-based search across name and position.
- **Pagination**: Efficient list retrieval with limit and offset.
- **Concurrent Export**: Parallel export to JSON and CSV using Goroutines.
- **Middleware**: Logging, Panic Recovery, and Basic Authentication.
- **SQL Transactions**: Safe deletion of departments with associated records.

## Tech Stack
- **Language**: Go (Golang)
- **Database**: PostgreSQL
- **Architecture**: Layered (Handler -> Service -> Repository)

## Setup
1. Create PostgreSQL database: `employee_management_system`
2. Set environment variable: `DATABASE_URL=postgres://user:pass@localhost:5432/employee_management_system?sslmode=disable`
3. Run migrations (SQL scripts provided in the documentation).
4. Build and run: `go run cmd/app/main.go`

## API Authentication
- **Username**: `admin`
- **Password**: `admin123`
