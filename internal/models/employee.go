package models

type Employee struct {
	ID           int     `json:"id"`
	Name         string  `json:"name"`
	Age          int     `json:"age"`
	Position     string  `json:"position"`
	DepartmentID int     `json:"departmentId"`
	Salary       float64 `json:"salary"`
	DepartmentName string `json:"departmentName,omitempty"`
}
