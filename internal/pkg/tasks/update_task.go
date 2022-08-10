package tasks

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h *handler) UpdateTask(ctx *gin.Context) {
	var task TaskAttributes
	taskId := ctx.Param("id")
	err := h.DB.QueryRow(`SELECT id, name, deadline, status FROM tasks WHERE id = ?`, taskId).Scan(&task.ID, &task.Name, &task.Deadline, &task.Status)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	task.Name = strings.TrimSpace(ctx.PostForm("name"))
	task.Deadline = strings.TrimSpace(ctx.PostForm("deadline"))
	task.Status = strings.TrimSpace(ctx.PostForm("status"))

	err = h.Validate.Struct(task)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	query := `UPDATE tasks SET name = ?, deadline = ?, status = ? WHERE id = ?`
	_, err = h.DB.Exec(query, task.Name, task.Deadline, task.Status, task.ID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.Redirect(http.StatusMovedPermanently, "/tasks")
}
