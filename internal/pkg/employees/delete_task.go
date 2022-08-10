package employees

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) DeleteEmployee(ctx *gin.Context) {
	var id int64
	err := h.DB.QueryRow(`SELECT id FROM employees WHERE id = ?`, ctx.Param("id")).Scan(&id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	_, err = h.DB.Exec(`DELETE FROM employees WHERE id = ?`, ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.Redirect(http.StatusMovedPermanently, "/employees/")
}
