package employees

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) AddEmployee(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "add_employee.tmpl", gin.H{
		"global": h.GlobalVars,
		"title":  "Add New Employee",
		"css":    []string{"tasks/new-edit-task.css"},
	})
}
