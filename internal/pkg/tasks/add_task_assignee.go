package tasks

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h *handler) AddTaskAssignee(ctx *gin.Context) {
	taskId := strings.TrimSpace(ctx.Param("id"))
	employeeId := strings.TrimSpace(ctx.PostForm("employee-id"))

	var idIfExists int64
	err := h.DB.QueryRow(`SELECT id FROM task_assignees WHERE task_id = ? AND employee_id = ?`, taskId, employeeId).Scan(&idIfExists)
	if err != nil {
		query := `INSERT INTO task_assignees (task_id, employee_id) VALUES (?, ?)`
		_, err = h.DB.Exec(query, taskId, employeeId)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
	}

	ctx.Redirect(http.StatusMovedPermanently, "/tasks/"+taskId+"/edit") // TODO BAHAYA redirect percaya input user
}
