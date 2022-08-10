package commons

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func InitRouterHandler() *RouterHandler {
	var validate *validator.Validate
	var db *sql.DB

	/*
		Ganti dengan username:password@alamat_server_mysql:3306/nama_database?parseTime=true
		Hati-Hati, jangan upload info database remote di GitHub
	*/
	db, err := sql.Open("mysql", "root@(127.0.0.1:3306)/test_gotodo?parseTime=true")

	if err != nil {
		panic(err)
	}
	validate = validator.New()

	r := gin.Default()

	RegisterFunctions(r)

	r.LoadHTMLGlob("./web/templates/**/*")
	r.Static("/public", "./web/public")
	r.StaticFile("/favicon.ico", "./web/public/favicon.ico")

	return &RouterHandler{
		// Instance router gin
		R: r,
		// Instance database
		DB: db,
		// Instance validator
		Validate: validate,
		// Variabel Global untuk template
		GlobalVars: &gin.H{
			"DocumentTitle": "GoToDo",
		},
	}
}
