package repositories

import (
	"context"
	"database/sql"
	"github.com/thaohs-3758/employee_management_system/internal/models"
	"fmt"

	_ "github.com/lib/pq"
)

type EmployeeRepository interface {
	GetAll(ctx context.Context, limit, offset int, deptID *int) ([]models.Employee, int, error)
	GetByID(ctx context.Context, id int) (*models.Employee, error)
	Create(ctx context.Context, employee *models.Employee) error
	Update(ctx context.Context, employee *models.Employee) error
	Delete(ctx context.Context, id int) error
	GetByDepartmentID(ctx context.Context, deptID int) ([]models.Employee, error)
	Search(ctx context.Context, keyword string) ([]models.Employee, error)
}

type employeeRepo struct {
	db *sql.DB
}

func NewEmployeeRepository(db *sql.DB) EmployeeRepository {
	return &employeeRepo{db: db}
}

func (r *employeeRepo) GetAll(ctx context.Context, limit, offset int, deptID *int) ([]models.Employee, int, error) {
	var employees []models.Employee
	var totalCount int

	query := `SELECT id, name, age, position, department_id, salary FROM employees`
	countQuery := `SELECT COUNT(*) FROM employees`
	args := []interface{}{}
	argPos := 1

	if deptID != nil {
		query += fmt.Sprintf(" WHERE department_id = $%d", argPos)
		countQuery += fmt.Sprintf(" WHERE department_id = $%d", argPos)
		args = append(args, *deptID)
		argPos++
	}

	err := r.db.QueryRowContext(ctx, countQuery, args...).Scan(&totalCount)
	if err != nil {
		return nil, 0, err
	}
	query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", argPos, argPos+1)
	args = append(args, limit, offset)

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	for rows.Next() {
		var emp models.Employee
		if err := rows.Scan(
			&emp.ID,
			&emp.Name,
			&emp.Age,
			&emp.Position,
			&emp.DepartmentID,
			&emp.Salary,
		); err != nil {
			return nil, 0, err
		}
		employees = append(employees, emp)
	}
	return employees, totalCount, nil
}

func (r *employeeRepo) GetByID(ctx context.Context, id int) (*models.Employee, error) {
	var emp models.Employee
	err := r.db.QueryRowContext(ctx, `
		SELECT id, name, age, position, department_id, salary
		FROM employees
		WHERE id = $1
	`, id).Scan(
		&emp.ID,
		&emp.Name,
		&emp.Age,
		&emp.Position,
		&emp.DepartmentID,
		&emp.Salary,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &emp, nil
}

func (r *employeeRepo) Create(ctx context.Context, employee *models.Employee) error {
	err := r.db.QueryRowContext(ctx,
		`INSERT INTO employees (name, age, position, department_id, salary)
		 VALUES ($1, $2, $3, $4, $5) RETURNING id`,
		employee.Name,
		employee.Age,
		employee.Position,
		employee.DepartmentID,
		employee.Salary,
	).Scan(&employee.ID)
	return err
}

func (r *employeeRepo) Update(ctx context.Context, employee *models.Employee) error {
	_, err := r.db.ExecContext(ctx,
		`UPDATE employees SET name = $1, age = $2, position = $3, department_id = $4, salary = $5 WHERE id = $6`,
		employee.Name,
		employee.Age,
		employee.Position,
		employee.DepartmentID,
		employee.Salary,
		employee.ID,
	)
	return err
}

func (r *employeeRepo) Delete(ctx context.Context, id int) error {
	_, err := r.db.ExecContext(ctx, "DELETE FROM employees WHERE id = $1", id)
	return err
}

func (r *employeeRepo) GetByDepartmentID(ctx context.Context, deptID int) ([]models.Employee, error) {
	rows, err := r.db.QueryContext(ctx, `
		SELECT e.id, e.name, e.age, e.position, e.department_id, e.salary, d.name as department_name
		FROM employees e
		JOIN departments d ON e.department_id = d.id
		WHERE e.department_id = $1
	`, deptID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var employees []models.Employee
	for rows.Next() {
		var emp models.Employee
		if err := rows.Scan(
			&emp.ID,
			&emp.Name,
			&emp.Age,
			&emp.Position,
			&emp.DepartmentID,
			&emp.Salary,
			&emp.DepartmentName,
		); err != nil {
			return nil, err
		}
		employees = append(employees, emp)
	}
	return employees, nil
}
func (r *employeeRepo) Search(ctx context.Context, keyword string) ([]models.Employee, error) {
	rows, err := r.db.QueryContext(ctx, `
		SELECT id, name, age, position, department_id, salary
		FROM employees
		WHERE name ILIKE $1 OR position ILIKE $1
	`, "%"+keyword+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var employees []models.Employee
	for rows.Next() {
		var emp models.Employee
		if err := rows.Scan(
			&emp.ID,
			&emp.Name,
			&emp.Age,
			&emp.Position,
			&emp.DepartmentID,
			&emp.Salary,
		); err != nil {
			return nil, err
		}
		employees = append(employees, emp)
	}
	return employees, nil
}
