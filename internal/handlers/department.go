package handlers

import (
	"github.com/thaohs-3758/employee_management_system/internal/models"
	"github.com/thaohs-3758/employee_management_system/internal/services"
	"github.com/thaohs-3758/employee_management_system/internal/utils"
	"net/http"
	"strconv"
)

type DepartmentHandler struct {
	service *services.DepartmentService
}

func NewDepartmentHandler(service *services.DepartmentService) *DepartmentHandler {
	return &DepartmentHandler{service: service}
}

// GetAll handles GET /departments
func (h *DepartmentHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	departments, err := h.service.GetAllDepartments(r.Context())
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondJSON(w, http.StatusOK, departments)

}

// GetByID handles GET /departments/:id
func (h *DepartmentHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, "invalid department ID")
		return
	}

	department, err := h.service.GetDepartmentByID(r.Context(), id)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if department == nil {
		utils.RespondError(w, http.StatusNotFound, "department not found")
		return
	}

	utils.RespondJSON(w, http.StatusOK, department)
}

// Create handles POST /departments
func (h *DepartmentHandler) Create(w http.ResponseWriter, r *http.Request) {
	var department models.Department
	if err := utils.ReadJSON(r, &department); err != nil {
		utils.RespondError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := h.service.CreateDepartment(r.Context(), &department); err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondJSON(w, http.StatusCreated, department)
}

// Update handles PUT /departments/:id
func (h *DepartmentHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, "invalid department ID")
		return
	}

	var department models.Department
	if err := utils.ReadJSON(r, &department); err != nil {
		utils.RespondError(w, http.StatusBadRequest, "invalid request body")
		return
	}
	department.ID = id

	if err := h.service.UpdateDepartment(r.Context(), &department); err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondJSON(w, http.StatusOK, department)
}

// Delete handles DELETE /departments/:id
func (h *DepartmentHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, "invalid department ID")
		return
	}

	if err := h.service.DeleteDepartment(r.Context(), id); err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondJSON(w, http.StatusNoContent, nil)
}
