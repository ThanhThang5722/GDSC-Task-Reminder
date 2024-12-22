package routes

import (
	"TaskReminder/handlers"

	"github.com/gin-gonic/gin"
)

func Task(r *gin.RouterGroup) {
	r.POST("/addTask", handlers.CreatTask)
}

func Login(r *gin.RouterGroup) {
	loginRoute := r.Group("/user")
	{
		loginRoute.POST("", handlers.SignUp)
		loginRoute.POST("/login", handlers.LoginAPI)
	}
}
