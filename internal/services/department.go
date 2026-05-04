package services

import (
	"context"
	"errors"
	"employee_management_system/internal/models"
	"employee_management_system/internal/repositories"
)

type DepartmentService struct {
	repo repositories.DepartmentRepository
}

func NewDepartmentService(repo repositories.DepartmentRepository) *DepartmentService {
	return &DepartmentService{repo: repo}
}

func (s *DepartmentService) GetAllDepartments(ctx context.Context) ([]models.Department, error) {
	return s.repo.GetAll(ctx)
}

func (s *DepartmentService) GetDepartmentByID(ctx context.Context, id int) (*models.Department, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *DepartmentService) CreateDepartment(ctx context.Context, department *models.Department) error {
	if err := s.validateDepartment(department); err != nil {
		return err
	}
	return s.repo.Create(ctx, department)
}

func (s *DepartmentService) UpdateDepartment(ctx context.Context, department *models.Department) error {
	if err := s.validateDepartment(department); err != nil {
		return err
	}
	return s.repo.Update(ctx, department)
}

func (s *DepartmentService) validateDepartment(department *models.Department) error {
	if department.Name == "" {
		return errors.New("department name cannot be empty")
	}
	return nil
}

func (s *DepartmentService) DeleteDepartment(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}
