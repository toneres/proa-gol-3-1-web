package employees

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h *handler) CreateEmployee(ctx *gin.Context) {
	employee := &Employee{
		Name:   strings.TrimSpace(ctx.PostForm("name")),
		Gender: strings.TrimSpace(ctx.PostForm("gender")),
	}

	err := h.Validate.Struct(employee)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	query := `INSERT INTO employees (name, gender) VALUES (?, ?)`
	result, err := h.DB.Exec(query, employee.Name, employee.Gender)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	employee.ID, err = result.LastInsertId()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.Redirect(http.StatusMovedPermanently, "/employees/")
}
