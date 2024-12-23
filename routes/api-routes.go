package routes

import (
	"TaskReminder/handlers"

	"github.com/gin-gonic/gin"
)

func Task(r *gin.RouterGroup) {
	taskRoute := r.Group("/Task")
	{
		taskRoute.POST("/", handlers.CreatTask)
		taskRoute.DELETE("/", handlers.RemoveTask)
		taskRoute.PUT("/", handlers.UpdatePriority)
	}
}

func Login(r *gin.RouterGroup) {
	loginRoute := r.Group("/user")
	{
		loginRoute.POST("/", handlers.SignUp)
		loginRoute.POST("/login", handlers.LoginAPI)
		loginRoute.POST("/signup", handlers.SignUp)
	}
}
