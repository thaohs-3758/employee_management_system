package handlers

import (
	"employee_management_system/internal/models"
	"employee_management_system/internal/services"
	"employee_management_system/internal/utils"
	"net/http"
	"strconv"
)

type EmployeeHandler struct {
	service *services.EmployeeService
}

func NewEmployeeHandler(service *services.EmployeeService) *EmployeeHandler {
	return &EmployeeHandler{service: service}
}

// GetAll handles GET /employees?limit=10&offset=0&departmentId=1
func (h *EmployeeHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	limit, _ := strconv.Atoi(query.Get("limit"))
	if limit <= 0 {
		limit = 10
	}

	offset, _ := strconv.Atoi(query.Get("offset"))
	if offset < 0 {
		offset = 0
	}

	var deptID *int
	if val := query.Get("departmentId"); val != "" {
		if id, err := strconv.Atoi(val); err == nil {
			deptID = &id
		}
	}

	employees, totalCount, err := h.service.GetAllEmployees(r.Context(), limit, offset, deptID)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondJSON(w, http.StatusOK, map[string]interface{}{
		"totalCount": totalCount,
		"employees":  employees,
	})
}

// GetByID handles GET /employees/:id
func (h *EmployeeHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, "invalid employee ID")
		return
	}

	employee, err := h.service.GetEmployeeByID(r.Context(), id)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if employee == nil {
		utils.RespondError(w, http.StatusNotFound, "employee not found")
		return
	}

	utils.RespondJSON(w, http.StatusOK, employee)
}

// Create handles POST /employees
func (h *EmployeeHandler) Create(w http.ResponseWriter, r *http.Request) {
	var employee models.Employee
	if err := utils.ReadJSON(r, &employee); err != nil {
		utils.RespondError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := h.service.CreateEmployee(r.Context(), &employee); err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondJSON(w, http.StatusCreated, employee)
}

// Update handles PUT /employees/:id
func (h *EmployeeHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, "invalid employee ID")
		return
	}

	var employee models.Employee
	if err := utils.ReadJSON(r, &employee); err != nil {
		utils.RespondError(w, http.StatusBadRequest, "invalid request body")
		return
	}
	employee.ID = id

	if err := h.service.UpdateEmployee(r.Context(), &employee); err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondJSON(w, http.StatusOK, employee)
}

// Delete handles DELETE /employees/:id
func (h *EmployeeHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, "invalid employee ID")
		return
	}

	if err := h.service.DeleteEmployee(r.Context(), id); err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondJSON(w, http.StatusNoContent, nil)
}

// GetByDepartment handles GET /departments/:id/employees
func (h *EmployeeHandler) GetByDepartment(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, "invalid department ID")
		return
	}

	employees, err := h.service.GetEmployeesByDepartment(r.Context(), id)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondJSON(w, http.StatusOK, employees)
}

// Search handles GET /employees/search?keyword=...
func (h *EmployeeHandler) Search(w http.ResponseWriter, r *http.Request) {
	keyword := r.URL.Query().Get("keyword")
	if keyword == "" {
		utils.RespondError(w, http.StatusBadRequest, "keyword is required")
		return
	}

	employees, err := h.service.SearchEmployees(r.Context(), keyword)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondJSON(w, http.StatusOK, employees)
}

// Export handles POST /employees/export
func (h *EmployeeHandler) Export(w http.ResponseWriter, r *http.Request) {
	files, err := h.service.ExportEmployees(r.Context())
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondJSON(w, http.StatusOK, map[string]interface{}{
		"message": "Export completed successfully",
		"files":   files,
	})
}
