package tasks

import (
	"github.com/toneres/proa-gol-3-1-web/internal/pkg/commons"
)

type handler struct {
	*commons.RouterHandler
}

func RegisterRoutes(routerHandler *commons.RouterHandler) {

	h := &handler{
		routerHandler,
	}

	routes := h.R.Group("/tasks")

	routes.GET("/", h.GetTasks)
	routes.GET("/add", h.AddTask)
	routes.POST("/create", h.CreateTask)
	routes.GET("/:id/edit", h.EditTask)
	routes.POST("/:id/update", h.UpdateTask)
	routes.POST("/:id/delete", h.DeleteTask)
	routes.POST("/:id/done", h.SetTaskDone)
	routes.POST("/:id/assignee/add", h.AddTaskAssignee)
	routes.POST("/:id/assignee/remove", h.RemoveTaskAssignee)
}
