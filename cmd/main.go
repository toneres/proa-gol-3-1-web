package main

import (
	"github.com/toneres/proa-gol-3-1-web/internal/pkg/commons"
	"github.com/toneres/proa-gol-3-1-web/internal/pkg/employees"
	"github.com/toneres/proa-gol-3-1-web/internal/pkg/index"
	"github.com/toneres/proa-gol-3-1-web/internal/pkg/tasks"
)

func main() {
	routerHandler := commons.InitRouterHandler()

	tasks.RegisterRoutes(routerHandler)
	index.RegisterRoutes(routerHandler)
	employees.RegisterRoutes(routerHandler)

	routerHandler.R.Run(":8000")
}
