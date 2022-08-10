package employees

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

	routes := h.R.Group("/employees")

	routes.GET("/", h.GetEmployees)
	routes.GET("/add", h.AddEmployee)
	routes.POST("/create", h.CreateEmployee)
	routes.GET("/:id/edit", h.EditEmployee)
	routes.POST("/:id/update", h.UpdateEmployee)
	routes.POST("/:id/delete", h.DeleteEmployee)
}
