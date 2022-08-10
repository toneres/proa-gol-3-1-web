package commons

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type RouterHandler struct {
	R          *gin.Engine
	DB         *sql.DB
	Validate   *validator.Validate
	GlobalVars *gin.H
}
