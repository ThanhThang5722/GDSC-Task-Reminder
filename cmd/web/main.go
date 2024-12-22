package main

import (
	"TaskReminder/pkg/auth"
	"TaskReminder/pkg/database"
	"TaskReminder/pkg/middlewares"
	"TaskReminder/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db := database.GetDbInstance()
	defer db.Close()

	auth.GenerateJWTKey()

	router := gin.Default()
	router.Use(middlewares.CorsMiddleware)
	router.Use(middlewares.CacheBuster)
	api := router.Group("/api")

	routes.Task(api)
	routes.Login(api)

	router.LoadHTMLGlob("templates/**/*.html")
	router.Static("/assets", "./assets")

	web := router.Group("")
	routes.WebRoute(web)
	router.Run()
}
