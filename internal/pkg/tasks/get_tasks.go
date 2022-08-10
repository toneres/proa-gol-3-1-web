package tasks

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/toneres/proa-gol-3-1-web/internal/pkg/employees"
)

func (h *handler) GetTasks(ctx *gin.Context) {
	rows, err := h.DB.Query(`SELECT id, name, deadline, status FROM tasks`)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	var tasks []Task
	for rows.Next() {
		var t Task
		rows.Scan(&t.ID, &t.Name, &t.Deadline, &t.Status)

		var allAssignees []employees.Employee
		rowsA, err := h.DB.Query(`SELECT e.id, name, gender FROM employees e JOIN task_assignees ta ON e.id = ta.employee_id WHERE ta.task_id = ? `, t.ID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		for rowsA.Next() {
			var e employees.Employee
			rowsA.Scan(&e.ID, &e.Name, &e.Gender)
			allAssignees = append(allAssignees, e)
		}
		rowsA.Close()

		t.Assignees = allAssignees
		tasks = append(tasks, t)
	}
	rows.Close()

	ctx.HTML(http.StatusOK, "get_tasks.tmpl", gin.H{
		"global": h.GlobalVars,
		"title":  "Tasks List",
		"css":    []string{"tasks/get-tasks.css"},
		"tasks":  tasks,
	})
}
