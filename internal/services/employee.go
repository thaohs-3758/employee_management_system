package services

import (
	"context"
	"github.com/thaohs-3758/employee_management_system/internal/models"
	"github.com/thaohs-3758/employee_management_system/internal/repositories"
	"github.com/thaohs-3758/employee_management_system/internal/utils"
	"errors"
	"os"
	"sync"
)

type EmployeeService struct {
	repo repositories.EmployeeRepository
}

func NewEmployeeService(repo repositories.EmployeeRepository) *EmployeeService {
	return &EmployeeService{repo: repo}
}

func (s *EmployeeService) GetAllEmployees(ctx context.Context, limit, offset int, deptID *int) ([]models.Employee, int, error) {
	return s.repo.GetAll(ctx, limit, offset, deptID)
}

func (s *EmployeeService) GetEmployeeByID(ctx context.Context, id int) (*models.Employee, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *EmployeeService) CreateEmployee(ctx context.Context, employee *models.Employee) error {
	if err := s.validateEmployee(employee); err != nil {
		return err
	}
	return s.repo.Create(ctx, employee)
}

func (s *EmployeeService) UpdateEmployee(ctx context.Context, employee *models.Employee) error {
	if err := s.validateEmployee(employee); err != nil {
		return err
	}
	return s.repo.Update(ctx, employee)
}

func (s *EmployeeService) validateEmployee(employee *models.Employee) error {
	if employee.Name == "" {
		return errors.New("employee name cannot be empty")
	}
	if employee.Age <= 0 {
		return errors.New("employee age must be greater than 0")
	}
	if employee.Salary < 0 {
		return errors.New("employee salary cannot be negative")
	}
	return nil
}

func (s *EmployeeService) DeleteEmployee(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}

func (s *EmployeeService) GetEmployeesByDepartment(ctx context.Context, deptID int) ([]models.Employee, error) {
	return s.repo.GetByDepartmentID(ctx, deptID)
}

func (s *EmployeeService) SearchEmployees(ctx context.Context, keyword string) ([]models.Employee, error) {
	return s.repo.Search(ctx, keyword)
}

func (s *EmployeeService) ExportEmployees(ctx context.Context) ([]string, error) {
	employees, _, err := s.repo.GetAll(ctx, 1000, 0, nil)
	if err != nil {
		return nil, err
	}

	var wg sync.WaitGroup
	var mu sync.Mutex
	var exportedFiles []string
	var exportErrors []error

	os.MkdirAll("exports", 0755)

	wg.Add(2)
	go func() {
		defer wg.Done()
		filename := "exports/employees.json"
		if err := utils.ExportToJSON(filename, employees); err != nil {
			mu.Lock()
			exportErrors = append(exportErrors, err)
			mu.Unlock()
			return
		}
		mu.Lock()
		exportedFiles = append(exportedFiles, filename)
		mu.Unlock()
	}()

	go func() {
		defer wg.Done()
		filename := "exports/employees.csv"
		if err := utils.ExportToCSV(filename, employees); err != nil {
			mu.Lock()
			exportErrors = append(exportErrors, err)
			mu.Unlock()
			return
		}
		mu.Lock()
		exportedFiles = append(exportedFiles, filename)
		mu.Unlock()
	}()

	wg.Wait()

	if len(exportErrors) > 0 {
		return exportedFiles, exportErrors[0]
	}

	return exportedFiles, nil
}
