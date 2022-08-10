package commons

import (
	"text/template"
	"time"

	"github.com/gin-gonic/gin"
)

func RegisterFunctions(router *gin.Engine) {
	router.SetFuncMap(template.FuncMap{
		"add":         add,
		"date_format": sqlDateFormat,
	})
}

func add(x, y int) int {
	return x + y
}

func sqlDateFormat(sqlDate, format string) string {
	t, err := time.Parse(time.RFC3339, sqlDate)

	if err != nil {
		return ""
	}

	return t.Format(format)
}
