package index

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

	routes := h.R.Group("/")

	routes.GET("/", h.Index)
}
