package tasks

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/toneres/proa-gol-3-1-web/internal/pkg/employees"
)

func (h *handler) EditTask(ctx *gin.Context) {
	var task Task
	err := h.DB.QueryRow(`SELECT id, name, deadline, status FROM tasks WHERE id = ?`, ctx.Param("id")).Scan(&task.ID, &task.Name, &task.Deadline, &task.Status)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	var allAssignees []employees.Employee
	rows, err := h.DB.Query(`SELECT e.id, name, gender FROM employees e JOIN task_assignees ta ON e.id = ta.employee_id WHERE ta.task_id = ? `, ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	for rows.Next() {
		var e employees.Employee
		rows.Scan(&e.ID, &e.Name, &e.Gender)
		allAssignees = append(allAssignees, e)
	}

	rows.Close()

	task.Assignees = allAssignees

	var allEmployees []employees.Employee

	rows, err = h.DB.Query(`SELECT id, name, gender FROM employees`)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	for rows.Next() {
		var e employees.Employee
		rows.Scan(&e.ID, &e.Name, &e.Gender)
		allEmployees = append(allEmployees, e)
	}

	rows.Close()

	ctx.HTML(http.StatusOK, "edit_task.tmpl", gin.H{
		"global":       h.GlobalVars,
		"title":        "Edit Existing Task",
		"task":         task,
		"css":          []string{"tasks/new-edit-task.css"},
		"allEmployees": allEmployees,
	})
}
