package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func Run() {
	getRoutes()
	err := router.Run(":8000")

	router.Use(cors.Default())

	if err != nil {
		return
	}

}

func getRoutes() {
	api := router.Group("/api")
	parseCode(api)

}
