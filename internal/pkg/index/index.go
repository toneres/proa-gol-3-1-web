package index

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) Index(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.tmpl", gin.H{
		"global": h.GlobalVars,
		"title":  "Beranda",
		"css":    []string{"index.css"},
	})
}
