package main

import (
	"employee_management_system/internal/database"
	"employee_management_system/internal/handlers"
	"employee_management_system/internal/middleware"
	"employee_management_system/internal/repositories"
	"employee_management_system/internal/services"
	"log"
	"net/http"
	"os"
)

func main() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	db, err := database.NewConnection(dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	employeeRepo := repositories.NewEmployeeRepository(db)
	departmentRepo := repositories.NewDepartmentRepository(db)

	employeeService := services.NewEmployeeService(employeeRepo)
	departmentService := services.NewDepartmentService(departmentRepo)

	employeeHandler := handlers.NewEmployeeHandler(employeeService)
	departmentHandler := handlers.NewDepartmentHandler(departmentService)

	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/employees", employeeHandler.GetAll)
	mux.HandleFunc("GET /api/employees/search", employeeHandler.Search)
	mux.HandleFunc("POST /api/employees/export", employeeHandler.Export)
	mux.HandleFunc("GET /api/employees/{id}", employeeHandler.GetByID)
	mux.HandleFunc("POST /api/employees", employeeHandler.Create)
	mux.HandleFunc("PUT /api/employees/{id}", employeeHandler.Update)
	mux.HandleFunc("DELETE /api/employees/{id}", employeeHandler.Delete)

	mux.HandleFunc("GET /api/departments", departmentHandler.GetAll)
	mux.HandleFunc("GET /api/departments/{id}", departmentHandler.GetByID)
	mux.HandleFunc("POST /api/departments", departmentHandler.Create)
	mux.HandleFunc("PUT /api/departments/{id}", departmentHandler.Update)
	mux.HandleFunc("DELETE /api/departments/{id}", departmentHandler.Delete)
	mux.HandleFunc("GET /api/departments/{id}/employees", employeeHandler.GetByDepartment)

	handler := middleware.BasicAuth(middleware.Recovery(middleware.Logger(mux)))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on :%s", port)
	if err := http.ListenAndServe(":"+port, handler); err != nil {
		log.Fatal(err)
	}
}
