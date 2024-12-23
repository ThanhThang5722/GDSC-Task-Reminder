package routes

import (
	"TaskReminder/handlers"

	"github.com/gin-gonic/gin"
)

func WebRoute(r *gin.RouterGroup) {
	r.GET("/", handlers.RenderHomePage)
	r.GET("/login", handlers.RenderLoginPage)
	r.GET("/signup", handlers.RenderSignUpPage)
}
