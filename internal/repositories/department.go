package repositories

import (
	"context"
	"database/sql"
	"github.com/thaohs-3758/employee_management_system/internal/models"

	_ "github.com/lib/pq"
)

type DepartmentRepository interface {
	GetAll(ctx context.Context) ([]models.Department, error)
	GetByID(ctx context.Context, id int) (*models.Department, error)
	Create(ctx context.Context, department *models.Department) error
	Update(ctx context.Context, department *models.Department) error
	Delete(ctx context.Context, id int) error
}

type departmentRepo struct {
	db *sql.DB
}

func NewDepartmentRepository(db *sql.DB) DepartmentRepository {
	return &departmentRepo{db: db}
}

func (r *departmentRepo) GetAll(ctx context.Context) ([]models.Department, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT id, name FROM departments")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var departments []models.Department
	for rows.Next() {
		var dept models.Department
		if err := rows.Scan(&dept.ID, &dept.Name); err != nil {
			return nil, err
		}
		departments = append(departments, dept)
	}
	return departments, nil
}

func (r *departmentRepo) GetByID(ctx context.Context, id int) (*models.Department, error) {
	var dept models.Department
	err := r.db.QueryRowContext(ctx, "SELECT id, name FROM departments WHERE id = $1", id).Scan(&dept.ID, &dept.Name)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &dept, nil
}

func (r *departmentRepo) Create(ctx context.Context, department *models.Department) error {
	err := r.db.QueryRowContext(ctx, "INSERT INTO departments (name) VALUES ($1) RETURNING id", department.Name).Scan(&department.ID)
	return err
}

func (r *departmentRepo) Update(ctx context.Context, department *models.Department) error {
	_, err := r.db.ExecContext(ctx, "UPDATE departments SET name = $1 WHERE id = $2", department.Name, department.ID)
	return err
}

func (r *departmentRepo) Delete(ctx context.Context, id int) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, "DELETE FROM employees WHERE department_id = $1", id)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.ExecContext(ctx, "DELETE FROM departments WHERE id = $1", id)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
