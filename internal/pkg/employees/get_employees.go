package employees

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) GetEmployees(ctx *gin.Context) {
	rows, err := h.DB.Query(`SELECT id, name, gender FROM employees`)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	defer rows.Close()

	var employees []Employee
	for rows.Next() {
		var e Employee
		rows.Scan(&e.ID, &e.Name, &e.Gender)
		employees = append(employees, e)
	}

	ctx.HTML(http.StatusOK, "get_employees.tmpl", gin.H{
		"global":    h.GlobalVars,
		"title":     "Employees List",
		"employees": employees,
		"css":       []string{"tasks/get-tasks.css"},
	})
}
