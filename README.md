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

## API Documentation

- **Postman Collection:** [Employee Management System API](https://ho-sy-thao-2843588.postman.co/workspace/Ho-Sy-Thao's-Workspace~15bbe1dc-3846-453a-8a37-99239cf052a1/collection/54511773-cc46961f-19ba-486a-b522-17b3d1c7deef?action=share&creator=54511773)

## API Authentication
- **Username**: `admin`
- **Password**: `admin123`
