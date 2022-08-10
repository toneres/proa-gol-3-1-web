package tasks

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) SetTaskDone(ctx *gin.Context) {
	var task TaskAttributes
	err := h.DB.QueryRow(`SELECT id, name, assignee, deadline, status FROM tasks WHERE id = ?`, ctx.Param("id")).Scan(&task.ID, &task.Name, &task.Deadline, &task.Status)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	task.Status = "done"

	query := `UPDATE tasks SET status = ? WHERE id = ?`
	_, err = h.DB.Exec(query, task.Status, task.ID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.Redirect(http.StatusMovedPermanently, "/tasks/")
}
