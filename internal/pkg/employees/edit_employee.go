package employees

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) EditEmployee(ctx *gin.Context) {
	var employee Employee
	err := h.DB.QueryRow(`SELECT id, name, gender FROM employees WHERE id = ?`, ctx.Param("id")).Scan(&employee.ID, &employee.Name, &employee.Gender)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.HTML(http.StatusOK, "edit_employee.tmpl", gin.H{
		"global":   h.GlobalVars,
		"title":    "Edit Existing Employee",
		"employee": employee,
		"css":      []string{"tasks/new-edit-task.css"},
	})
}
