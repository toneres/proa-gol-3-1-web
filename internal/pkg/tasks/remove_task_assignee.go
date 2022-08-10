package tasks

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h *handler) RemoveTaskAssignee(ctx *gin.Context) {
	taskId := strings.TrimSpace(ctx.Param("id"))
	employeeId := strings.TrimSpace(ctx.PostForm("employee-id"))
	_, err := h.DB.Exec(`DELETE FROM task_assignees WHERE task_id = ? AND employee_id = ?`, taskId, employeeId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	// TODO error handling bukan karena not found

	ctx.Redirect(http.StatusMovedPermanently, "/tasks/"+taskId+"/edit") // TODO BAHAYA redirect percaya input user
}
