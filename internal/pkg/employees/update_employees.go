package employees

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h *handler) UpdateEmployee(ctx *gin.Context) {
	var employee Employee
	err := h.DB.QueryRow(`SELECT id, name, gender FROM employees WHERE id = ?`, ctx.Param("id")).Scan(&employee.ID, &employee.Name, &employee.Gender)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	employee.Name = strings.TrimSpace(ctx.PostForm("name"))
	employee.Gender = strings.TrimSpace(ctx.PostForm("gender"))

	err = h.Validate.Struct(employee)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	query := `UPDATE employees SET name = ?, gender = ? WHERE id = ?`
	_, err = h.DB.Exec(query, employee.Name, employee.Gender, employee.ID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.Redirect(http.StatusMovedPermanently, "/employees/")
}
