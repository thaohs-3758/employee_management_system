package utils

import (
	"github.com/thaohs-3758/employee_management_system/internal/models"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
)

func ExportToJSON(filename string, employees []models.Employee) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(employees)
}

func ExportToCSV(filename string, employees []models.Employee) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	header := []string{"ID", "Name", "Age", "Position", "DepartmentID", "Salary"}
	if err := writer.Write(header); err != nil {
		return err
	}

	for _, emp := range employees {
		row := []string{
			fmt.Sprintf("%d", emp.ID),
			emp.Name,
			fmt.Sprintf("%d", emp.Age),
			emp.Position,
			fmt.Sprintf("%d", emp.DepartmentID),
			fmt.Sprintf("%.2f", emp.Salary),
		}
		if err := writer.Write(row); err != nil {
			return err
		}
	}

	return nil
}
