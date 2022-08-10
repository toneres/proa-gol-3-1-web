package tasks

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) AddTask(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "add_task.tmpl", gin.H{
		"global": h.GlobalVars,
		"title":  "Add New Task",
		"css":    []string{"tasks/new-edit-task.css"},
	})
}
