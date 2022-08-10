package tasks

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h *handler) CreateTask(ctx *gin.Context) {
	task := &TaskAttributes{
		Name:     strings.TrimSpace(ctx.PostForm("name")),
		Deadline: strings.TrimSpace(ctx.PostForm("deadline")),
		Status:   strings.TrimSpace(ctx.PostForm("status")),
	}

	err := h.Validate.Struct(task)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	query := `INSERT INTO tasks (name, deadline, status) VALUES (?, ?, ?)`
	result, err := h.DB.Exec(query, task.Name, task.Deadline, task.Status)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	task.ID, err = result.LastInsertId()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.Redirect(http.StatusMovedPermanently, "/tasks/"+strconv.FormatInt(task.ID, 10)+"/edit")
}
